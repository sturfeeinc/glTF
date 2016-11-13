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
// The material appearance of a primitive.
type Material struct {
	GlTFChildOfRootProperty
	Technique GlTFid                    `json:"technique,omitempty"`
	Values    map[string]MaterialValues `json:"values"`
}
