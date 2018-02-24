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
var builtIn bool

var vomit = &cobra.Command{
	Use: "vomit",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO handle type config
		var renderTypeArray []render.RenderInputType
		if builtIn {
			renderTypeArray = allBuiltInRenderType
		} else {
			renderTypeArray = []render.RenderInputType{
				render.RenderInputType{
					Package: "github.com/xxx/test",
					Type:    "StructA",
				},
			}
		}

		renderIns := render.NewBagRender(&render.BagRenderConf{
			RenderInputTypes: renderTypeArray,
		})
		err := renderIns.RenderToFile(processFullPathDir(dir) + "/bag/" + "gen_bag.go")
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
	vomit.Flags().BoolVarP(&builtIn, "onlyBuiltIn", "b", false, "gen all builtin types")
}

func main() {
	Root.AddCommand(vomit)
	if err := Root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
