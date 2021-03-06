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

func easyjson3b94576aDecodeGithubComSturfeeincGlTF(in *jlexer.Lexer, out *Asset) {
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
		case "copyright":
			out.Copyright = string(in.String())
		case "generator":
			out.Generator = string(in.String())
		case "premultipliedAlpha":
			out.PremultipliedAlpha = bool(in.Bool())
		case "profile":
			(out.Profile).UnmarshalEasyJSON(in)
		case "version":
			out.Version = string(in.String())
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
func easyjson3b94576aEncodeGithubComSturfeeincGlTF(out *jwriter.Writer, in Asset) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Copyright != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"copyright\":")
		out.String(string(in.Copyright))
	}
	if in.Generator != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"generator\":")
		out.String(string(in.Generator))
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"premultipliedAlpha\":")
	out.Bool(bool(in.PremultipliedAlpha))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"profile\":")
	(in.Profile).MarshalEasyJSON(out)
	if in.Version != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"version\":")
		out.String(string(in.Version))
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
func (v Asset) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3b94576aEncodeGithubComSturfeeincGlTF(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Asset) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3b94576aEncodeGithubComSturfeeincGlTF(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Asset) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3b94576aDecodeGithubComSturfeeincGlTF(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Asset) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3b94576aDecodeGithubComSturfeeincGlTF(l, v)
}
