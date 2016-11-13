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
// Metadata about the glTF asset.
type Asset struct {
	GlTFProperty
	Copyright          string       `json:"copyright,omitempty"`
	Generator          string       `json:"generator,omitempty"`
	PremultipliedAlpha bool         `json:"premultipliedAlpha"`
	Profile            AssetProfile `json:"profile"`
	Version            string       `json:"version,omitempty" validator:"required"`
}
