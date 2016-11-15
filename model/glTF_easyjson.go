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

func easyjson45d8642bDecodeGithubComSturfeeincGlTF(in *jlexer.Lexer, out *GlTF) {
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
		case "asset":
			(out.Asset).UnmarshalEasyJSON(in)
		case "cameras":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Cameras = make(map[string]Camera)
				} else {
					out.Cameras = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 Camera
					(v1).UnmarshalEasyJSON(in)
					(out.Cameras)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
		case "images":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Images = make(map[string]Image)
				} else {
					out.Images = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v2 Image
					(v2).UnmarshalEasyJSON(in)
					(out.Images)[key] = v2
					in.WantComma()
				}
				in.Delim('}')
			}
		case "scene":
			out.Scene = GlTFid(in.String())
		case "accessors":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Accessors = make(map[string]Accessor)
				} else {
					out.Accessors = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v3 Accessor
					(v3).UnmarshalEasyJSON(in)
					(out.Accessors)[key] = v3
					in.WantComma()
				}
				in.Delim('}')
			}
		case "animations":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Animations = make(map[string]Animation)
				} else {
					out.Animations = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v4 Animation
					(v4).UnmarshalEasyJSON(in)
					(out.Animations)[key] = v4
					in.WantComma()
				}
				in.Delim('}')
			}
		case "meshes":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Meshes = make(map[string]Mesh)
				} else {
					out.Meshes = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v5 Mesh
					(v5).UnmarshalEasyJSON(in)
					(out.Meshes)[key] = v5
					in.WantComma()
				}
				in.Delim('}')
			}
		case "scenes":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Scenes = make(map[string]Scene)
				} else {
					out.Scenes = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v6 Scene
					(v6).UnmarshalEasyJSON(in)
					(out.Scenes)[key] = v6
					in.WantComma()
				}
				in.Delim('}')
			}
		case "techniques":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Techniques = make(map[string]Technique)
				} else {
					out.Techniques = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v7 Technique
					(v7).UnmarshalEasyJSON(in)
					(out.Techniques)[key] = v7
					in.WantComma()
				}
				in.Delim('}')
			}
		case "textures":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Textures = make(map[string]Texture)
				} else {
					out.Textures = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v8 Texture
					(v8).UnmarshalEasyJSON(in)
					(out.Textures)[key] = v8
					in.WantComma()
				}
				in.Delim('}')
			}
		case "extensionsUsed":
			in.Delim('[')
			if !in.IsDelim(']') {
				out.ExtensionsUsed = make([]string, 0, 4)
			} else {
				out.ExtensionsUsed = nil
			}
			for !in.IsDelim(']') {
				var v9 string
				v9 = string(in.String())
				out.ExtensionsUsed = append(out.ExtensionsUsed, v9)
				in.WantComma()
			}
			in.Delim(']')
		case "bufferViews":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.BufferViews = make(map[string]BufferView)
				} else {
					out.BufferViews = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v10 BufferView
					(v10).UnmarshalEasyJSON(in)
					(out.BufferViews)[key] = v10
					in.WantComma()
				}
				in.Delim('}')
			}
		case "materials":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Materials = make(map[string]Material)
				} else {
					out.Materials = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v11 Material
					(v11).UnmarshalEasyJSON(in)
					(out.Materials)[key] = v11
					in.WantComma()
				}
				in.Delim('}')
			}
		case "nodes":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Nodes = make(map[string]Node)
				} else {
					out.Nodes = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v12 Node
					(v12).UnmarshalEasyJSON(in)
					(out.Nodes)[key] = v12
					in.WantComma()
				}
				in.Delim('}')
			}
		case "programs":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Programs = make(map[string]Program)
				} else {
					out.Programs = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v13 Program
					(v13).UnmarshalEasyJSON(in)
					(out.Programs)[key] = v13
					in.WantComma()
				}
				in.Delim('}')
			}
		case "shaders":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Shaders = make(map[string]Shader)
				} else {
					out.Shaders = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v14 Shader
					(v14).UnmarshalEasyJSON(in)
					(out.Shaders)[key] = v14
					in.WantComma()
				}
				in.Delim('}')
			}
		case "skins":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Skins = make(map[string]Skin)
				} else {
					out.Skins = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v15 Skin
					(v15).UnmarshalEasyJSON(in)
					(out.Skins)[key] = v15
					in.WantComma()
				}
				in.Delim('}')
			}
		case "buffers":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Buffers = make(map[string]Buffer)
				} else {
					out.Buffers = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v16 Buffer
					(v16).UnmarshalEasyJSON(in)
					(out.Buffers)[key] = v16
					in.WantComma()
				}
				in.Delim('}')
			}
		case "samplers":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Samplers = make(map[string]Sampler)
				} else {
					out.Samplers = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v17 Sampler
					(v17).UnmarshalEasyJSON(in)
					(out.Samplers)[key] = v17
					in.WantComma()
				}
				in.Delim('}')
			}
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
func easyjson45d8642bEncodeGithubComSturfeeincGlTF(out *jwriter.Writer, in GlTF) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"asset\":")
	(in.Asset).MarshalEasyJSON(out)
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"cameras\":")
	if in.Cameras == nil {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v18First := true
		for v18Name, v18Value := range in.Cameras {
			if !v18First {
				out.RawByte(',')
			}
			v18First = false
			out.String(string(v18Name))
			out.RawByte(':')
			(v18Value).MarshalEasyJSON(out)
		}
		out.RawByte('}')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"images\":")
	if in.Images == nil {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v19First := true
		for v19Name, v19Value := range in.Images {
			if !v19First {
				out.RawByte(',')
			}
			v19First = false
			out.String(string(v19Name))
			out.RawByte(':')
			(v19Value).MarshalEasyJSON(out)
		}
		out.RawByte('}')
	}
	if in.Scene != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"scene\":")
		out.String(string(in.Scene))
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"accessors\":")
	if in.Accessors == nil {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v20First := true
		for v20Name, v20Value := range in.Accessors {
			if !v20First {
				out.RawByte(',')
			}
			v20First = false
			out.String(string(v20Name))
			out.RawByte(':')
			(v20Value).MarshalEasyJSON(out)
		}
		out.RawByte('}')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"animations\":")
	if in.Animations == nil {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v21First := true
		for v21Name, v21Value := range in.Animations {
			if !v21First {
				out.RawByte(',')
			}
			v21First = false
			out.String(string(v21Name))
			out.RawByte(':')
			(v21Value).MarshalEasyJSON(out)
		}
		out.RawByte('}')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"meshes\":")
	if in.Meshes == nil {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v22First := true
		for v22Name, v22Value := range in.Meshes {
			if !v22First {
				out.RawByte(',')
			}
			v22First = false
			out.String(string(v22Name))
			out.RawByte(':')
			(v22Value).MarshalEasyJSON(out)
		}
		out.RawByte('}')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"scenes\":")
	if in.Scenes == nil {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v23First := true
		for v23Name, v23Value := range in.Scenes {
			if !v23First {
				out.RawByte(',')
			}
			v23First = false
			out.String(string(v23Name))
			out.RawByte(':')
			(v23Value).MarshalEasyJSON(out)
		}
		out.RawByte('}')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"techniques\":")
	if in.Techniques == nil {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v24First := true
		for v24Name, v24Value := range in.Techniques {
			if !v24First {
				out.RawByte(',')
			}
			v24First = false
			out.String(string(v24Name))
			out.RawByte(':')
			(v24Value).MarshalEasyJSON(out)
		}
		out.RawByte('}')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"textures\":")
	if in.Textures == nil {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v25First := true
		for v25Name, v25Value := range in.Textures {
			if !v25First {
				out.RawByte(',')
			}
			v25First = false
			out.String(string(v25Name))
			out.RawByte(':')
			(v25Value).MarshalEasyJSON(out)
		}
		out.RawByte('}')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"extensionsUsed\":")
	out.RawByte('[')
	for v26, v27 := range in.ExtensionsUsed {
		if v26 > 0 {
			out.RawByte(',')
		}
		out.String(string(v27))
	}
	out.RawByte(']')
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"bufferViews\":")
	if in.BufferViews == nil {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v28First := true
		for v28Name, v28Value := range in.BufferViews {
			if !v28First {
				out.RawByte(',')
			}
			v28First = false
			out.String(string(v28Name))
			out.RawByte(':')
			(v28Value).MarshalEasyJSON(out)
		}
		out.RawByte('}')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"materials\":")
	if in.Materials == nil {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v29First := true
		for v29Name, v29Value := range in.Materials {
			if !v29First {
				out.RawByte(',')
			}
			v29First = false
			out.String(string(v29Name))
			out.RawByte(':')
			(v29Value).MarshalEasyJSON(out)
		}
		out.RawByte('}')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"nodes\":")
	if in.Nodes == nil {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v30First := true
		for v30Name, v30Value := range in.Nodes {
			if !v30First {
				out.RawByte(',')
			}
			v30First = false
			out.String(string(v30Name))
			out.RawByte(':')
			(v30Value).MarshalEasyJSON(out)
		}
		out.RawByte('}')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"programs\":")
	if in.Programs == nil {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v31First := true
		for v31Name, v31Value := range in.Programs {
			if !v31First {
				out.RawByte(',')
			}
			v31First = false
			out.String(string(v31Name))
			out.RawByte(':')
			(v31Value).MarshalEasyJSON(out)
		}
		out.RawByte('}')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"shaders\":")
	if in.Shaders == nil {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v32First := true
		for v32Name, v32Value := range in.Shaders {
			if !v32First {
				out.RawByte(',')
			}
			v32First = false
			out.String(string(v32Name))
			out.RawByte(':')
			(v32Value).MarshalEasyJSON(out)
		}
		out.RawByte('}')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"skins\":")
	if in.Skins == nil {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v33First := true
		for v33Name, v33Value := range in.Skins {
			if !v33First {
				out.RawByte(',')
			}
			v33First = false
			out.String(string(v33Name))
			out.RawByte(':')
			(v33Value).MarshalEasyJSON(out)
		}
		out.RawByte('}')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"buffers\":")
	if in.Buffers == nil {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v34First := true
		for v34Name, v34Value := range in.Buffers {
			if !v34First {
				out.RawByte(',')
			}
			v34First = false
			out.String(string(v34Name))
			out.RawByte(':')
			(v34Value).MarshalEasyJSON(out)
		}
		out.RawByte('}')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"samplers\":")
	if in.Samplers == nil {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v35First := true
		for v35Name, v35Value := range in.Samplers {
			if !v35First {
				out.RawByte(',')
			}
			v35First = false
			out.String(string(v35Name))
			out.RawByte(':')
			(v35Value).MarshalEasyJSON(out)
		}
		out.RawByte('}')
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
func (v GlTF) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson45d8642bEncodeGithubComSturfeeincGlTF(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GlTF) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson45d8642bEncodeGithubComSturfeeincGlTF(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GlTF) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson45d8642bDecodeGithubComSturfeeincGlTF(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GlTF) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson45d8642bDecodeGithubComSturfeeincGlTF(l, v)
}