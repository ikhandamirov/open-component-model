package dir_test

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"ocm.software/open-component-model/bindings/go/blob"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ocm.software/open-component-model/bindings/go/input/dir"
	v1 "ocm.software/open-component-model/bindings/go/input/dir/spec/v1"
	"ocm.software/open-component-model/bindings/go/runtime"
)

func TestGetV1DirBlob_Success(t *testing.T) {
	tests := []struct {
		name           string
		mediaType      string
		compress       bool
		preserveDir    bool
		followSymlinks bool
		excludeFiles   []string
		includeFiles   []string
	}{
		{
			name:           "default dir spec",
			mediaType:      "application/vnd.gardener.landscaper.blueprint.v1+tar",
			compress:       false,
			preserveDir:    false,
			followSymlinks: false,
			excludeFiles:   []string{},
			includeFiles:   []string{},
		},
		{
			name:           "compressed dir",
			mediaType:      "application/vnd.gardener.landscaper.blueprint.v1+tar",
			compress:       true,
			preserveDir:    false,
			followSymlinks: false,
			excludeFiles:   []string{},
			includeFiles:   []string{},
		},
		{
			name:           "preserveDir set to true",
			mediaType:      "application/vnd.gardener.landscaper.blueprint.v1+tar",
			compress:       false,
			preserveDir:    true,
			followSymlinks: false,
			excludeFiles:   []string{},
			includeFiles:   []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create directory to test with.
			tempDir := t.TempDir()
			dirBase := "input-dir"
			dirAbs := filepath.Join(tempDir, dirBase)
			err := os.Mkdir(dirAbs, 0755)
			require.NoError(t, err)
			fileName1 := filepath.Join(dirAbs, "blueprint.yaml")
			fileData1 := "blueprint"
			fileName2 := filepath.Join(dirAbs, "deploy-execution.yaml")
			fileData2 := "deploy-execution"
			fileName3 := filepath.Join(dirAbs, "export-execution.yaml")
			fileData3 := "export-execution"
			err = os.WriteFile(fileName1, []byte(fileData1), 0644)
			require.NoError(t, err)
			err = os.WriteFile(fileName2, []byte(fileData2), 0644)
			require.NoError(t, err)
			err = os.WriteFile(fileName3, []byte(fileData3), 0644)
			require.NoError(t, err)

			// Create v1.File spec.
			dirSpec := v1.Dir{
				Type:           runtime.NewUnversionedType("dir"),
				Path:           dirAbs,
				MediaType:      tt.mediaType,
				Compress:       tt.compress,
				PreserveDir:    tt.preserveDir,
				FollowSymlinks: tt.followSymlinks,
				ExcludeFiles:   tt.excludeFiles,
				IncludeFiles:   tt.includeFiles,
			}

			// Get blob.
			b, err := dir.GetV1DirBlob(dirSpec)
			require.NoError(t, err)
			require.NotNil(t, b)

			// Test blob properties.
			if sizeAware, ok := b.(blob.SizeAware); ok {
				size := sizeAware.Size()
				assert.GreaterOrEqual(t, size, int64(0))
			}

			if digestAware, ok := b.(blob.DigestAware); ok {
				digest, ok := digestAware.Digest()
				assert.True(t, ok)
				assert.NotEmpty(t, digest)
			}

			// Test reading data
			reader, err := b.ReadCloser()
			require.NoError(t, err)
			defer reader.Close()

			data, err := io.ReadAll(reader)
			require.NoError(t, err)

			if tt.compress {
				// Decompress gzipped data
				gzReader, err := gzip.NewReader(bytes.NewReader(data))
				require.NoError(t, err)
				defer gzReader.Close()

				data, err = io.ReadAll(gzReader)
				require.NoError(t, err)

				// Test media type for compressed blob
				if mediaTypeAware, ok := b.(blob.MediaTypeAware); ok {
					mediaType, known := mediaTypeAware.MediaType()
					assert.True(t, known)
					assert.Equal(t, tt.mediaType+"+gzip", mediaType)
				}
			} else {
				// Test media type for uncompressed blob
				if mediaTypeAware, ok := b.(blob.MediaTypeAware); ok {
					mediaType, known := mediaTypeAware.MediaType()
					assert.True(t, known)
					assert.Equal(t, tt.mediaType, mediaType)
				}
			}

			// Extract files from tar archive
			fileToExtract := filepath.Base(fileName1)
			if tt.preserveDir {
				fileToExtract = filepath.Join(dirBase, fileToExtract)
			}
			untarredData1, err := extractFileFromTar(data, fileToExtract)
			require.NoError(t, err)
			assert.Equal(t, fileData1, string(untarredData1))

			fileToExtract = filepath.Base(fileName2)
			if tt.preserveDir {
				fileToExtract = filepath.Join(dirBase, fileToExtract)
			}
			untarredData2, err := extractFileFromTar(data, fileToExtract)
			require.NoError(t, err)
			assert.Equal(t, fileData2, string(untarredData2))

			fileToExtract = filepath.Base(fileName3)
			if tt.preserveDir {
				fileToExtract = filepath.Join(dirBase, fileToExtract)
			}
			untarredData3, err := extractFileFromTar(data, fileToExtract)
			require.NoError(t, err)
			assert.Equal(t, fileData3, string(untarredData3))
		})
	}
}

func TestGetV1DirBlob_EmptyPath(t *testing.T) {
	// Create v1.Dir spec with empty path.
	dirSpec := v1.Dir{
		Type: runtime.NewUnversionedType("file"),
		Path: "",
	}

	// Get blob should fail.
	dirBlob, err := dir.GetV1DirBlob(dirSpec)
	assert.Error(t, err)
	assert.Nil(t, dirBlob)
}

func TestGetV1DirBlob_EmptyDir(t *testing.T) {
	// Create v1.Dir spec pointing to folder with no files.
	tempDir := t.TempDir()
	dirSpec := v1.Dir{
		Type: runtime.NewUnversionedType("file"),
		Path: tempDir,
	}

	// Get blob should fail
	dirBlob, err := dir.GetV1DirBlob(dirSpec)
	assert.Error(t, err)
	assert.Nil(t, dirBlob)
}

// extractFileFromTar extracts a specific file from a tar archive and returns its content
func extractFileFromTar(tarData []byte, fileName string) ([]byte, error) {
	// Create a reader from the byte data
	reader := bytes.NewReader(tarData)

	// Create a tar reader
	tr := tar.NewReader(reader)

	// Normalize the file name for comparison
	normalizedFileName := filepath.Clean(fileName)

	// Iterate through the files in the tar archive
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			return nil, fmt.Errorf("error reading tar header: %w", err)
		}

		// Normalize the header name for comparison
		normalizedHeaderName := filepath.Clean(header.Name)

		// Check if this is the file we're looking for
		if normalizedHeaderName == normalizedFileName {
			// Make sure it's a regular file
			if header.Typeflag != tar.TypeReg {
				return nil, fmt.Errorf("'%s' is not a regular file (type: %c)", fileName, header.Typeflag)
			}

			// Read the file content
			content, err := io.ReadAll(tr)
			if err != nil {
				return nil, fmt.Errorf("error reading file content: %w", err)
			}

			return content, nil
		}
	}

	// File not found
	return nil, fmt.Errorf("file '%s' not found in tar archive", fileName)
}
