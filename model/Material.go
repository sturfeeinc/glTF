/*
 * glTF JSON model
 *
 * Do not modify this file. It is automatically generated
 * with JsonModelGen (https://github.com/sturfeeinc/JsonModelGen)
 * Copyright (c) 2016 Sturfee inc. - http://sturfee.com
 */

package model

type Material struct {
	GlTFChildOfRootProperty
	Technique GlTFid                     `json:"technique"`
	Values    map[string]*MaterialValues `json:"values"`
}
