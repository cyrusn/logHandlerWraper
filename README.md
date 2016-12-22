# A logging handler wrapper for Golang net/http package

## Usage

``` go
package main

import (
  "log"
  "net/http"

  "github.com/cyrusn/logHandlerWraper"
)

func main() {
  addr := "127.0.0.1:8080"
  loggedFileServer := logHandlerWraper.Wrap(http.FileServer(http.Dir("/usr/share/doc")))
  fmt.Println("Starting up, serving:", addr)
  log.Fatal(http.ListenAndServe(addr, loggedFileServer))
}
```
