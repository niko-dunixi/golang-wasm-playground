package main

import (
	"github.com/gorilla/mux"
	playground "github.com/paul-nelson-baker/golang-wasm-playground"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	// I should figure out if there's an easy way to fill out this data that isn't overly complicated
	Pages = []string{"simple-cat-example.wasm"}
	// Stub this out to prevent false compiler errors. It's actually initialized in `assets_vfsdata.go`.
	Assets http.FileSystem
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/resources/javascript/{javascript}", JavaScriptHandler)
	router.HandleFunc("/resources/css/{css}", CssHandler)
	router.HandleFunc("/wasm/{wasm}", WasmHandler)
	router.HandleFunc("/page/{page}", PageHandler)
	router.HandleFunc("/", HomeHandler)

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
	Pages []string
}

type WasmPage struct {
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
		Pages: Pages,
	}
	if err = t.Execute(w, homePage); err != nil {
		log.Println(err)
		w.WriteHeader(500)
		_, _ = w.Write([]byte(err.Error()))
	}
}

func PageHandler(w http.ResponseWriter, r *http.Request) {
	t, err := loadTemplate("pages/wasm-page.gohtml")
	if err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte(err.Error()))
	}
	if t == nil {
		return
	}
	vars := mux.Vars(r)
	wasmFile := vars["page"]
	page := WasmPage{
		WasmFile: wasmFile,
	}
	if err = t.Execute(w, page); err != nil {
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
	vars := mux.Vars(r)
	file := vars["javascript"]
	passAssetToRequest("javascript/"+file, "application/javascript", w)
}

func WasmHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	file := vars["wasm"]
	passAssetToRequest("wasm/"+file, "application/wasm", w)
}

func passAssetToRequest(name, mime string, w http.ResponseWriter) {
	file, err := playground.Assets.Open(name)
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
	w.Header().Add("Content-Type", mime)
	_, _ = w.Write(bytes)
}

func loadTemplate(name string) (*template.Template, error) {
	file, err := playground.Assets.Open(name)
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return template.New(name).Parse(string(bytes))
}
