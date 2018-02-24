package render

import "testing"

func TestProcessBuiltInType(t *testing.T) {
	t.Log(processBuiltInType("int"))
	t.Log(processBuiltInType("*int"))
	t.Log(processBuiltInType("*string"))
	t.Log(processBuiltInType("string"))
}

func TestProcessImportType(t *testing.T) {
	// pkg.Fuck PkgFuck
	t.Log(processImportType(&RenderInputType{
		Package: "github.com/xxx/pkg",
		Type:    "Fuck",
	}))

	// *pkg.Fuck PkgFuckPtr
	t.Log(processImportType(&RenderInputType{
		Package: "github.com/xxx/pkg",
		Type:    "*Fuck",
	}))

	// **pkg.Fuck PkgFuckPtrPtr
	t.Log(processImportType(&RenderInputType{
		Package: "github.com/xxx/pkg",
		Type:    "**Fuck",
	}))

}
