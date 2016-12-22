package logHandlerWraper

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestWrap(t *testing.T) {
	addr := "127.0.0.1:8080"
	loggedFileServer := Wrap(http.FileServer(http.Dir("/usr/share/doc")))
	fmt.Println("Starting up, serving:", addr)
	log.Fatal(http.ListenAndServe(addr, loggedFileServer))
}
