//+build ignore

package main

import (
	"github.com/shurcooL/vfsgen"
	"net/http"
)

func main() {
	var fs http.FileSystem = http.Dir("assets")
	vfsgen.Generate(fs, vfsgen.Options{
		PackageName:  "base_rando_project",
		VariableName: "Assets",
	})
}
