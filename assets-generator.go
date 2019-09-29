//+build ignore

package main

import (
	"bytes"
	"github.com/shurcooL/vfsgen"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"
)

func main() {
	goroot := getGoRoot()
	wasmFile := path.Join(goroot, "misc", "wasm", "wasm_exec.js")
	bytes, err := ioutil.ReadFile(wasmFile)
	if err != nil {
		log.Fatalln(err)
	}
	err = ioutil.WriteFile(path.Join("assets", "wasm_exec.js"), bytes, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	var fs http.FileSystem = http.Dir("assets")
	vfsgen.Generate(fs, vfsgen.Options{
		PackageName:  "base_rando_project",
		VariableName: "Assets",
	})
}

func getGoRoot() string {
	goroot := os.Getenv("GOROOT")
	if goroot == "" {
		command := exec.Command("go", "env", "GOROOT")
		b := bytes.Buffer{}
		command.Stdout = &b
		if err := command.Run(); err != nil {
			log.Fatalln(err)
		}
		goroot = strings.TrimSpace(b.String())
	}
	return goroot
}
