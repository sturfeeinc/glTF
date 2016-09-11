/*
 * glTF JSON model
 *
 * Do not modify this file. It is automatically generated
 * with glTFModelGenerator (https://github.com/sturfeeinc/glTFModelGenerator)
 * Copyright (c) 2016 Sturfee inc. - http://sturfee.com
 * generated 2016-09-11 18:55:47.106260962 +0300 MSK
 */

package glTF

// http://json-schema.org/draft-03/schema
// A camera's projection.  A node can reference a camera ID to apply a transform to place the camera in the scene.
type Camera struct {
	GlTFChildOfRootProperty
	Type         string             `json:"type,omitempty" validator:"required"`
	Orthographic CameraOrthographic `json:"orthographic,omitempty"`
	Perspective  CameraPerspective  `json:"perspective,omitempty"`
}
