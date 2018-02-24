package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/JodeZer/bag/render"
	"github.com/spf13/cobra"
)

var Root = &cobra.Command{
	Use: "bag",
}

var dir string
var conf string

var vomit = &cobra.Command{
	Use: "vomit",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO handle type config
		renderIns := render.NewBagRender(&render.BagRenderConf{
			RenderInputTypes: []render.RenderInputType{
				render.RenderInputType{
					Type: "int",
				},
				render.RenderInputType{
					Type: "string",
				},
				render.RenderInputType{
					Package: "github.com/xxx/pkg",
					Type:    "StructA",
				},
			},
		})
		err := renderIns.RenderToFile(processFullPathDir(dir) + "/bag/" + "bag.go")
		if err != nil {
			panic(err)
		}
	},
}

func processFullPathDir(dir string) string {
	curDir, err := filepath.Abs(dir)
	if err != nil {
		panic(err)
	}
	return curDir
}

func init() {
	vomit.Flags().StringVarP(&dir, "to", "t", "", "write file to")
	vomit.Flags().StringVarP(&conf, "config", "c", "", "type config")
}

func main() {
	Root.AddCommand(vomit)
	if err := Root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
