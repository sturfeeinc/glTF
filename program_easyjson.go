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

func easyjson1c6ddb42DecodeGithubComSturfeeincGlTF(in *jlexer.Lexer, out *Program) {
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
		case "attributes":
			in.Delim('[')
			if !in.IsDelim(']') {
				out.Attributes = make([]string, 0, 4)
			} else {
				out.Attributes = nil
			}
			for !in.IsDelim(']') {
				var v1 string
				v1 = string(in.String())
				out.Attributes = append(out.Attributes, v1)
				in.WantComma()
			}
			in.Delim(']')
		case "fragmentShader":
			out.FragmentShader = GlTFid(in.String())
		case "vertexShader":
			out.VertexShader = GlTFid(in.String())
		case "name":
			out.Name = string(in.String())
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
func easyjson1c6ddb42EncodeGithubComSturfeeincGlTF(out *jwriter.Writer, in Program) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"attributes\":")
	out.RawByte('[')
	for v2, v3 := range in.Attributes {
		if v2 > 0 {
			out.RawByte(',')
		}
		out.String(string(v3))
	}
	out.RawByte(']')
	if in.FragmentShader != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"fragmentShader\":")
		out.String(string(in.FragmentShader))
	}
	if in.VertexShader != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"vertexShader\":")
		out.String(string(in.VertexShader))
	}
	if in.Name != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"name\":")
		out.String(string(in.Name))
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
func (v Program) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1c6ddb42EncodeGithubComSturfeeincGlTF(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Program) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1c6ddb42EncodeGithubComSturfeeincGlTF(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Program) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1c6ddb42DecodeGithubComSturfeeincGlTF(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Program) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1c6ddb42DecodeGithubComSturfeeincGlTF(l, v)
}
