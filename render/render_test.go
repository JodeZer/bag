package render

import (
	"io"
	"os"
	"testing"
)

func TestRender(t *testing.T) {
	render := NewBagRender(&BagRenderConf{
		RenderInputTypes: []RenderInputType{
			// RenderInputType{
			// 	Type: "int",
			// },
			// RenderInputType{
			// 	Type: "string",
			// },
			RenderInputType{
				Package: "github.com/xxx/pkg",
				Type:    "StructA",
			},
		},
	})

	reader, e := render.Render()

	t.Logf("%q", e)
	io.Copy(os.Stdout, reader)

	t.Log(render.RenderToFile("a.go"))
}
