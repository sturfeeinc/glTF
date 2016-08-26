/*
 * glTF JSON model
 *
 * Do not modify this file. It is automatically generated
 * with JsonModelGen (https://github.com/sturfeeinc/JsonModelGen)
 * Copyright (c) 2016 Sturfee inc. - http://sturfee.com
 */

package model

type Camera struct {
	GlTFChildOfRootProperty
	Orthographic CameraOrthographic `json:"orthographic"`
	Perspective  CameraPerspective  `json:"perspective"`
	Type         *string            `json:"type" validator:"required"`
}
