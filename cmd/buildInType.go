package main

import (
	"github.com/JodeZer/bag/render"
)

var buildInTypes = []string{
	"string",
	"int",
	"int64",
	"int32",
	"int16",
	"int8",
	"uint",
	"uint8",
	"uint16",
	"uint32",
	"uint64",
	"float32",
	"float64",
	"bool",
	"rune",
	"complex64",
	"complex128",
}

var allBuiltInRenderType []render.RenderInputType

func init() {
	pointerType := make([]string, len(buildInTypes))
	for i, t := range buildInTypes {
		pointerType[i] = "*" + t
	}
	buildInTypes = append(buildInTypes, pointerType...)

	for _, t := range buildInTypes {
		allBuiltInRenderType = append(allBuiltInRenderType, render.RenderInputType{
			Type: t,
		})
	}
}
