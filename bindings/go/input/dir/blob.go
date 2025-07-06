package dir

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"ocm.software/open-component-model/bindings/go/blob"
	"ocm.software/open-component-model/bindings/go/blob/compression"
	"ocm.software/open-component-model/bindings/go/blob/filesystem"
	"ocm.software/open-component-model/bindings/go/blob/inmemory"
	v1 "ocm.software/open-component-model/bindings/go/input/dir/spec/v1"
)

type InputDirBlob struct {
	*filesystem.Blob
	DirMediaType string
}

func (i InputDirBlob) MediaType() (mediaType string, known bool) {
	return i.DirMediaType, i.DirMediaType != ""
}

var _ interface {
	blob.MediaTypeAware
	blob.SizeAware
	blob.DigestAware
} = (*InputDirBlob)(nil)

func GetV1DirBlob(dir v1.Dir) (blob.ReadOnlyBlob, error) {
	if dir.Path == "" {
		return nil, fmt.Errorf("dir path must not be empty")
	}

	// TODO:
	// - Handle FollowSymlinks, ExcludeFiles and IncludeFiles options.

	reader, err := packDirToTar(dir.Path, &dir)
	if err != nil {
		return nil, fmt.Errorf("error producing tar archive: %w", err)
	}

	var dirBlob blob.ReadOnlyBlob = inmemory.New(reader, inmemory.WithMediaType(dir.MediaType))

	if dir.Compress {
		dirBlob = compression.Compress(dirBlob)
	}

	return dirBlob, nil
}

func packDirToTar(path string, dir *v1.Dir) (_ io.Reader, err error) {
	// Determine the base directory for relative paths in the tar archive.
	baseDir := path
	if dir.PreserveDir {
		baseDir = filepath.Dir(path)
	}

	// Create a new virtual FileSystem instance based on the provided directory path.
	fs, err := filesystem.NewFS(baseDir, os.O_RDONLY)
	if err != nil {
		return nil, fmt.Errorf("failed to create filesystem while trying to access %v: %w", path, err)
	}

	// Create a new tar writer
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)
	defer func() {
		closeErr := tw.Close()
		if closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	// Walk recursively through directory contents and add it to the tar.
	err = walkDirContents(path, baseDir, fs, tw)
	if err != nil {
		return nil, fmt.Errorf("failed to add directory contents to tar: %w", err)
	}

	// Close the tar writer
	if err := tw.Close(); err != nil {
		return nil, fmt.Errorf("failed to close tar writer: %w", err)
	}

	return bytes.NewReader(buf.Bytes()), nil
}

func walkDirContents(currentDir string, baseDir string, fs filesystem.FileSystem, tw *tar.Writer) (err error) {
	// Read directory contents.
	dirRelPath, err := filepath.Rel(baseDir, currentDir)
	dirEntries, err := fs.ReadDir(dirRelPath)
	if err != nil {
		return err
	}

	// Iterate over directory entries.
	for _, entry := range dirEntries {
		// Get FileInfo for the entry.
		info, err := entry.Info()
		if err != nil {
			return err
		}

		// Create tar header.
		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}

		// Set the name in the tar archive (relative to the folder being archived).
		entryPath := filepath.Join(currentDir, entry.Name())
		relPath, err := filepath.Rel(baseDir, entryPath)
		if err != nil {
			return err
		}
		header.Name = relPath

		// Write the header to the tar archive.
		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		// If the entry is a file, copy its content to the tar archive.
		if entry.Type().IsRegular() {
			file, err := fs.OpenFile(relPath, os.O_RDONLY, 0644)
			if err != nil {
				return err
			}
			defer func() {
				closeErr := file.Close()
				if closeErr != nil && err == nil {
					err = closeErr
				}
			}()

			if _, err := io.Copy(tw, file); err != nil {
				return err
			}
		}

		// If the entry is a directory, recursively process its subfolders.
		if entry.IsDir() {
			if err := walkDirContents(entryPath, baseDir, fs, tw); err != nil {
				return err
			}
		}
	}

	return nil
}
