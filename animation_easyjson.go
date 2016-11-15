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

func easyjson7b4e1d00DecodeGithubComSturfeeincGlTF(in *jlexer.Lexer, out *Animation) {
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
		case "samplers":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Samplers = make(map[string]AnimationSampler)
				} else {
					out.Samplers = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 AnimationSampler
					(v1).UnmarshalEasyJSON(in)
					(out.Samplers)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
		case "channels":
			in.Delim('[')
			if !in.IsDelim(']') {
				out.Channels = make([]AnimationChannel, 0, 1)
			} else {
				out.Channels = nil
			}
			for !in.IsDelim(']') {
				var v2 AnimationChannel
				(v2).UnmarshalEasyJSON(in)
				out.Channels = append(out.Channels, v2)
				in.WantComma()
			}
			in.Delim(']')
		case "parameters":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Parameters = make(map[string]AnimationParameter)
				} else {
					out.Parameters = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v3 AnimationParameter
					v3 = in.Interface()
					(out.Parameters)[key] = v3
					in.WantComma()
				}
				in.Delim('}')
			}
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
func easyjson7b4e1d00EncodeGithubComSturfeeincGlTF(out *jwriter.Writer, in Animation) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"samplers\":")
	if in.Samplers == nil {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v4First := true
		for v4Name, v4Value := range in.Samplers {
			if !v4First {
				out.RawByte(',')
			}
			v4First = false
			out.String(string(v4Name))
			out.RawByte(':')
			(v4Value).MarshalEasyJSON(out)
		}
		out.RawByte('}')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"channels\":")
	out.RawByte('[')
	for v5, v6 := range in.Channels {
		if v5 > 0 {
			out.RawByte(',')
		}
		(v6).MarshalEasyJSON(out)
	}
	out.RawByte(']')
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"parameters\":")
	if in.Parameters == nil {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v7First := true
		for v7Name, v7Value := range in.Parameters {
			if !v7First {
				out.RawByte(',')
			}
			v7First = false
			out.String(string(v7Name))
			out.RawByte(':')
			out.Raw(json.Marshal(v7Value))
		}
		out.RawByte('}')
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
func (v Animation) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7b4e1d00EncodeGithubComSturfeeincGlTF(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Animation) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7b4e1d00EncodeGithubComSturfeeincGlTF(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Animation) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7b4e1d00DecodeGithubComSturfeeincGlTF(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Animation) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7b4e1d00DecodeGithubComSturfeeincGlTF(l, v)
}
