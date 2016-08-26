/*
 * glTF JSON model
 *
 * Do not modify this file. It is automatically generated
 * with JsonModelGen (https://github.com/sturfeeinc/JsonModelGen)
 * Copyright (c) 2016 Sturfee inc. - http://sturfee.com
 */

package model

type TechniqueStatesFunctions struct {
	GlTFProperty
	BlendColor            []*float64 `json:"blendColor"`
	BlendEquationSeparate []*int     `json:"blendEquationSeparate"`
	ColorMask             []*bool    `json:"colorMask"`
	DepthMask             []*bool    `json:"depthMask"`
	PolygonOffset         []*float64 `json:"polygonOffset"`
	BlendFuncSeparate     []*int     `json:"blendFuncSeparate"`
	CullFace              []*int     `json:"cullFace"`
	DepthFunc             []*int     `json:"depthFunc"`
	DepthRange            []*float64 `json:"depthRange"`
	FrontFace             []*int     `json:"frontFace"`
	LineWidth             []*float64 `json:"lineWidth"`
	Scissor               []*float64 `json:"scissor"`
}
