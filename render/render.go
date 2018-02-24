package render

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"os"
	"strings"
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
	RenderInputTypes []RenderInputType
	destDir          string
}

func NewBagRender(conf *BagRenderConf) *BagRender {
	if conf == nil || len(conf.RenderInputTypes) == 0 {
		panic("fuck u")
	}
	ret := &BagRender{
		conf:   conf,
		render: NewRenderImpl(conf.RenderInputTypes),
	}
	for _, filename := range tplFiles {
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

func (r *BagRender) Render() (io.Reader, error) {
	buf := new(bytes.Buffer)
	if err := tplEngine.ExecuteTemplate(buf, "header", r.render); err != nil {
		return nil, err
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return nil, err
	}

	buf.Reset()

	_, err = buf.Write(formatted)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (r *BagRender) RenderToFile(filename string) error {
	lastSlash := strings.LastIndex(filename, "/")

	if err := os.MkdirAll(string(filename[:lastSlash+1]), os.ModePerm); err != nil {
		return fmt.Errorf("os.MkdirAll %s", err)
	}

	file, err := os.Create(filename)

	if err != nil {
		return fmt.Errorf("os.Create %s", err)
	}
	defer file.Close()
	reader, err := r.Render()
	if err != nil {
		return err
	}
	_, err = io.Copy(file, reader)

	return err
}

// render func

func (r *BagRender) UpperType() string {
	return "String"
}

func (r *BagRender) Type() string {
	return "string"
}
