/*
package logHandlerWraper is a handler wraper for package net/http,
which provide the http access log in os.StdOut
*/
package logHandlerWraper

import (
	"fmt"
	"net/http"
	"time"
)

type loggedResposeWriter struct {
	http.ResponseWriter
	status int
	size   int
}

func (w *loggedResposeWriter) Status() int {
	return w.status
}

func (w *loggedResposeWriter) Size() int {
	return w.size
}

func (w *loggedResposeWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *loggedResposeWriter) Write(b []byte) (int, error) {
	w.size = len(b)
	return w.ResponseWriter.Write(b)
}

func Wrap(handler http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		lw := &loggedResposeWriter{
			ResponseWriter: w,
			status:         200,
			size:           0,
		}
		handler.ServeHTTP(lw, req)

		method := req.Method
		start := time.Now()
		statusCode := lw.Status()
		size := lw.Size()
		requestURI := req.RequestURI
		logTimeFormat := "2006-01-02 15:04:05"

		content := fmt.Sprintf(
			"[%s] [%s](%d %s) \"%s\" in %v",
			time.Now().Format(logTimeFormat),
			method,
			statusCode,
			http.StatusText(statusCode),
			requestURI,
			time.Since(start),
		)

		if size != 0 {
			content = fmt.Sprintf("%s size: %d bytes", content, size)
		}

		switch statusCode {
		case 200, 201, 202:
			content = fmt.Sprintf("\033[1;32m%s\033[0m", content)
		case 301, 302:
			content = fmt.Sprintf("\033[1;37m%s\033[0m", content)
		case 304:
			content = fmt.Sprintf("\033[1;33m%s\033[0m", content)
		case 401, 403:
			content = fmt.Sprintf("\033[4;31m%s\033[0m", content)
		case 404:
			content = fmt.Sprintf("\033[1;31m%s\033[0m", content)
		case 500:
			content = fmt.Sprintf("\033[1;36m%s\033[0m", content)
		}
		fmt.Println(content)
	}
}
