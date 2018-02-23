package render

import (
	"bytes"
	"fmt"
	"go/format"
	"text/template"

	"github.com/JodeZer/bag/tpl"
)

type BagRender struct {
	conf   *BagRenderConf
	render *RenderImpl
}

var tplFiles = []string{
	"tpl/func.gogo",
	"tpl/header.gogo",
	"tpl/slice.gogo",
	"tpl/include.gogo",
}

var tplEngine = template.New("bag/tpl")

type BagRenderConf struct {
	TemplateFiles    []string
	SliceRenderTypes []string
	FuncRenderTypes  []string
	destDir          string
}

func NewBagRender(conf *BagRenderConf) *BagRender {
	if conf == nil {
		panic("fuck u")
	}
	ret := &BagRender{
		conf:   conf,
		render: NewRenderImpl(nil, nil),
	}
	for _, filename := range conf.TemplateFiles {
		data, err := tpl.Asset(filename)
		if err != nil {
			panic(err)
		}

		if _, err = tplEngine.Parse(string(data)); err != nil {
			panic(err)
		}
		//fmt.Println(string(data))
	}

	return ret
}

func (r *BagRender) Render() error {
	buf := new(bytes.Buffer)
	if err := tplEngine.ExecuteTemplate(buf, "header", r.render); err != nil {
		return err
	}
	//fmt.Println(buf.String())
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	// _, err = io.Copy(wr, bytes.NewBuffer(formatted))
	// return err
	fmt.Println(string(formatted))
	return nil
}

// render func

func (r *BagRender) UpperType() string {
	return "String"
}

func (r *BagRender) Type() string {
	return "string"
}
