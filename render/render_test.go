package render

import "testing"

func TestRender(t *testing.T) {
	render := NewBagRender(&BagRenderConf{
		TemplateFiles: tplFiles,
	})
	t.Logf("%q", render.Render())
}
