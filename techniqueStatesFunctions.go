/*
 * glTF JSON model
 *
 * Do not modify this file. It is automatically generated
 * with glTFModelGenerator (https://github.com/sturfeeinc/glTFModelGenerator)
 * Copyright (c) 2016 Sturfee inc. - http://sturfee.com
 * generated 2016-09-17 21:21:17.912222485 +0300 MSK
 */

package glTF

// http://json-schema.org/draft-03/schema
// Arguments for fixed-function rendering state functions other than `enable()`/`disable()`.
type TechniqueStatesFunctions struct {
	GlTFProperty
	BlendColor            []float64 `json:"blendColor"`
	BlendEquationSeparate []int     `json:"blendEquationSeparate"`
	BlendFuncSeparate     []int     `json:"blendFuncSeparate"`
	CullFace              []int     `json:"cullFace"`
	FrontFace             []int     `json:"frontFace"`
	PolygonOffset         []float64 `json:"polygonOffset"`
	Scissor               []float64 `json:"scissor"`
	ColorMask             []bool    `json:"colorMask"`
	DepthFunc             []int     `json:"depthFunc"`
	DepthMask             []bool    `json:"depthMask"`
	DepthRange            []float64 `json:"depthRange"`
	LineWidth             []float64 `json:"lineWidth"`
}
