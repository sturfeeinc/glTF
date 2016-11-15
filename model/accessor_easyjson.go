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

func easyjsonD39932edDecodeGithubComSturfeeincGlTF(in *jlexer.Lexer, out *Accessor) {
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
		case "min":
			in.Delim('[')
			if !in.IsDelim(']') {
				out.Min = make([]float64, 0, 8)
			} else {
				out.Min = nil
			}
			for !in.IsDelim(']') {
				var v1 float64
				v1 = float64(in.Float64())
				out.Min = append(out.Min, v1)
				in.WantComma()
			}
			in.Delim(']')
		case "bufferView":
			out.BufferView = GlTFid(in.String())
		case "byteOffset":
			out.ByteOffset = int(in.Int())
		case "byteStride":
			out.ByteStride = int(in.Int())
		case "componentType":
			out.ComponentType = int(in.Int())
		case "count":
			out.Count = int(in.Int())
		case "type":
			out.Type = string(in.String())
		case "max":
			in.Delim('[')
			if !in.IsDelim(']') {
				out.Max = make([]float64, 0, 8)
			} else {
				out.Max = nil
			}
			for !in.IsDelim(']') {
				var v2 float64
				v2 = float64(in.Float64())
				out.Max = append(out.Max, v2)
				in.WantComma()
			}
			in.Delim(']')
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
func easyjsonD39932edEncodeGithubComSturfeeincGlTF(out *jwriter.Writer, in Accessor) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Min) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"min\":")
		out.RawByte('[')
		for v3, v4 := range in.Min {
			if v3 > 0 {
				out.RawByte(',')
			}
			out.Float64(float64(v4))
		}
		out.RawByte(']')
	}
	if in.BufferView != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"bufferView\":")
		out.String(string(in.BufferView))
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"byteOffset\":")
	out.Int(int(in.ByteOffset))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"byteStride\":")
	out.Int(int(in.ByteStride))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"componentType\":")
	out.Int(int(in.ComponentType))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"count\":")
	out.Int(int(in.Count))
	if in.Type != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"type\":")
		out.String(string(in.Type))
	}
	if len(in.Max) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"max\":")
		out.RawByte('[')
		for v5, v6 := range in.Max {
			if v5 > 0 {
				out.RawByte(',')
			}
			out.Float64(float64(v6))
		}
		out.RawByte(']')
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
func (v Accessor) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD39932edEncodeGithubComSturfeeincGlTF(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Accessor) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD39932edEncodeGithubComSturfeeincGlTF(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Accessor) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD39932edDecodeGithubComSturfeeincGlTF(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Accessor) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD39932edDecodeGithubComSturfeeincGlTF(l, v)
}