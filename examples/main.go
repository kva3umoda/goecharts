package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"path"
	"runtime"

	"github.com/kva3umoda/goecharts/examples/examples"
)

type Exampler interface {
	Examples(workdir string)
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	workdir := workDir()

	examplers := []Exampler{
		examples.CandlestickExamples{},
		examples.LineExamples{},
	}

	for _, e := range examplers {
		e.Examples(workdir)
	}

	fs := http.FileServer(http.Dir(workdir))
	log.Println("running server at http://localhost:8089")
	log.Fatal(http.ListenAndServe("localhost:8089", logRequest(fs)))
}

func workDir() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic(errors.New("No caller information "))
	}

	curdir := path.Join(path.Dir(filename), "html")
	err := os.Mkdir(curdir, 0777)
	if err != nil && !errors.Is(err, os.ErrExist) {
		panic(err)
	}
	return curdir
}
