package logHandlerWraper_test

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cyrusn/logHandlerWraper"
)

func ExampleWrap() {
	addr := "127.0.0.1:8080"
	loggedFileServer := logHandlerWraper.Wrap(http.FileServer(http.Dir("/usr/share/doc")))
	fmt.Println("Starting up, serving:", addr)
	log.Fatal(http.ListenAndServe(addr, loggedFileServer))
}
