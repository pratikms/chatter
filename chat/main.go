package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/pratikms/chatter/trace"
)

type templateHandler struct {
	once     sync.Once
	filename string
	// templ represents a single template
	templ *template.Template
}

// ServeHTTP handles the HTTP request
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		log.Println("Setting template once")
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	log.Println("Executing template")
	t.templ.Execute(w, r)
}

func main() {
	var addr = flag.String("addr", ":8080", "The address of the application")
	flag.Parse()

	r := newRoom()
	r.tracer = trace.New(os.Stdout)

	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)

	// run the room
	go r.run()

	log.Printf("Starting server on %s\n...", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
