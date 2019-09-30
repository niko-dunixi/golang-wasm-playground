package main

import (
	"github.com/gorilla/mux"
	base "github.com/paul-nelson-baker/base-rando-project"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	Pages = []WasmPage{
		{
			Name:     "Test",
			WasmFile: "File.wasm",
		},
	}
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/resources/javascript/{javascript}", JavaScriptHandler)
	router.HandleFunc("/resources/css/{css}", CssHandler)
	router.HandleFunc("/resources/wasm/{wasm}", WasmHandler)
	router.HandleFunc("/page/{page}", ExperimentHandler)
	router.HandleFunc("/", HomeHandler)

	//http.Handle("/", router)

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

type HomePage struct {
	Title string
	Pages []WasmPage
}

type WasmPage struct {
	Name     string
	WasmFile string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := loadTemplate("pages/home.gohtml")
	if err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte(err.Error()))
	}
	if t == nil {
		return
	}

	homePage := HomePage{
		Title: "Type Safe Home",
		Pages: Pages,
	}
	if err = t.Execute(w, homePage); err != nil {
		log.Println(err)
		w.WriteHeader(500)
		_, _ = w.Write([]byte(err.Error()))
	}
}

func CssHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	file := vars["css"]
	passAssetToRequest("css/"+file, "text/css", w)
}

func JavaScriptHandler(w http.ResponseWriter, r *http.Request) {
	passAssetToRequest("javascript/wasm_exec.js", "application/javascript", w)

}

func WasmHandler(w http.ResponseWriter, r *http.Request) {
}

func ExperimentHandler(w http.ResponseWriter, r *http.Request) {

}

func passAssetToRequest(name, mime string, w http.ResponseWriter) {
	file, err := base.Assets.Open(name)
	if err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	_, _ = w.Write(bytes)
	w.Header().Add("Content-Type", mime)
}

func loadTemplate(name string) (*template.Template, error) {
	file, err := base.Assets.Open(name)
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return template.New(name).Parse(string(bytes))
}
