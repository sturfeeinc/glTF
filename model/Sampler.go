/*
 * glTF JSON model
 *
 * Do not modify this file. It is automatically generated
 * with JsonModelGen (https://github.com/sturfeeinc/JsonModelGen)
 * Copyright (c) 2016 Sturfee inc. - http://sturfee.com
 */

package model

type Sampler struct {
	GlTFChildOfRootProperty
	MagFilter *int `json:"magFilter"`
	MinFilter *int `json:"minFilter"`
	WrapS     *int `json:"wrapS"`
	WrapT     *int `json:"wrapT"`
}
