// AUTOGENERATED FILE: easyjson marshaller/unmarshallers.

package glTF

import (
	json "encoding/json"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ = json.RawMessage{}
	_ = jlexer.Lexer{}
	_ = jwriter.Writer{}
)

func easyjson6abf4620DecodeGithubComSturfeeincGlTF(in *jlexer.Lexer, out *AnimationChannelTarget) {
	if in.IsNull() {
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = GlTFid(in.String())
		case "path":
			out.Path = string(in.String())
		case "extensions":
			out.Extensions = in.Interface()
		case "extras":
			out.Extras = in.Interface()
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
}
func easyjson6abf4620EncodeGithubComSturfeeincGlTF(out *jwriter.Writer, in AnimationChannelTarget) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Id != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"id\":")
		out.String(string(in.Id))
	}
	if in.Path != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"path\":")
		out.String(string(in.Path))
	}
	if in.Extensions != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"extensions\":")
		out.Raw(json.Marshal(in.Extensions))
	}
	if in.Extras != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"extras\":")
		out.Raw(json.Marshal(in.Extras))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AnimationChannelTarget) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6abf4620EncodeGithubComSturfeeincGlTF(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AnimationChannelTarget) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6abf4620EncodeGithubComSturfeeincGlTF(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AnimationChannelTarget) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6abf4620DecodeGithubComSturfeeincGlTF(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AnimationChannelTarget) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6abf4620DecodeGithubComSturfeeincGlTF(l, v)
}
