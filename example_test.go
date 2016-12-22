/*
Package logHandlerWraper is a handler wraper for package net/http,
which provide the http access log in os.StdOut
*/
package logHandlerWraper

import (
	"fmt"
	"log"
	"net/http"
)

func ExampleWrap() {
	addr := "127.0.0.1:8080"
	loggedFileServer := Wrap(http.FileServer(http.Dir("/usr/share/doc")))
	fmt.Println("Starting up, serving:", addr)
	log.Fatal(http.ListenAndServe(addr, loggedFileServer))
}
