package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{}
}

func (h *Hello) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("hello world")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Oppps", http.StatusBadRequest)
		return
	}
	// log.Printf("data %s\n", d)
	fmt.Fprintf(w, "Hello %s:\n", d)
}
