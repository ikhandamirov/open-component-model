//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen-v0.32. DO NOT EDIT.

package runtime

import (
	json "encoding/json"

	goruntime "ocm.software/open-component-model/bindings/go/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentMeta) DeepCopyInto(out *ComponentMeta) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentMeta.
func (in *ComponentMeta) DeepCopy() *ComponentMeta {
	if in == nil {
		return nil
	}
	out := new(ComponentMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Digest) DeepCopyInto(out *Digest) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Digest.
func (in *Digest) DeepCopy() *Digest {
	if in == nil {
		return nil
	}
	out := new(Digest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElementMeta) DeepCopyInto(out *ElementMeta) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.ExtraIdentity != nil {
		in, out := &in.ExtraIdentity, &out.ExtraIdentity
		*out = make(goruntime.Identity, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElementMeta.
func (in *ElementMeta) DeepCopy() *ElementMeta {
	if in == nil {
		return nil
	}
	out := new(ElementMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Label) DeepCopyInto(out *Label) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Label.
func (in *Label) DeepCopy() *Label {
	if in == nil {
		return nil
	}
	out := new(Label)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalBlob) DeepCopyInto(out *LocalBlob) {
	*out = *in
	out.Type = in.Type
	if in.GlobalAccess != nil {
		out.GlobalAccess = in.GlobalAccess.DeepCopyTyped()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalBlob.
func (in *LocalBlob) DeepCopy() *LocalBlob {
	if in == nil {
		return nil
	}
	out := new(LocalBlob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyTyped is an autogenerated deepcopy function, copying the receiver, creating a new goruntime.Typed.
func (in *LocalBlob) DeepCopyTyped() goruntime.Typed {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Meta) DeepCopyInto(out *Meta) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Meta.
func (in *Meta) DeepCopy() *Meta {
	if in == nil {
		return nil
	}
	out := new(Meta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectMeta) DeepCopyInto(out *ObjectMeta) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]Label, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectMeta.
func (in *ObjectMeta) DeepCopy() *ObjectMeta {
	if in == nil {
		return nil
	}
	out := new(ObjectMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Reference) DeepCopyInto(out *Reference) {
	*out = *in
	in.ElementMeta.DeepCopyInto(&out.ElementMeta)
	out.Digest = in.Digest
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Reference.
func (in *Reference) DeepCopy() *Reference {
	if in == nil {
		return nil
	}
	out := new(Reference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resource) DeepCopyInto(out *Resource) {
	*out = *in
	in.ElementMeta.DeepCopyInto(&out.ElementMeta)
	if in.SourceRefs != nil {
		in, out := &in.SourceRefs, &out.SourceRefs
		*out = make([]SourceRef, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Access != nil {
		out.Access = in.Access.DeepCopyTyped()
	}
	if in.Digest != nil {
		in, out := &in.Digest, &out.Digest
		*out = new(Digest)
		**out = **in
	}
	in.CreationTime.DeepCopyInto(&out.CreationTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resource.
func (in *Resource) DeepCopy() *Resource {
	if in == nil {
		return nil
	}
	out := new(Resource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Signature) DeepCopyInto(out *Signature) {
	*out = *in
	out.Digest = in.Digest
	out.Signature = in.Signature
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Signature.
func (in *Signature) DeepCopy() *Signature {
	if in == nil {
		return nil
	}
	out := new(Signature)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignatureInfo) DeepCopyInto(out *SignatureInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignatureInfo.
func (in *SignatureInfo) DeepCopy() *SignatureInfo {
	if in == nil {
		return nil
	}
	out := new(SignatureInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Source) DeepCopyInto(out *Source) {
	*out = *in
	in.ElementMeta.DeepCopyInto(&out.ElementMeta)
	if in.Access != nil {
		out.Access = in.Access.DeepCopyTyped()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Source.
func (in *Source) DeepCopy() *Source {
	if in == nil {
		return nil
	}
	out := new(Source)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceRef) DeepCopyInto(out *SourceRef) {
	*out = *in
	if in.IdentitySelector != nil {
		in, out := &in.IdentitySelector, &out.IdentitySelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]Label, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceRef.
func (in *SourceRef) DeepCopy() *SourceRef {
	if in == nil {
		return nil
	}
	out := new(SourceRef)
	in.DeepCopyInto(out)
	return out
}
