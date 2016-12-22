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
			"[%s](%d %s) \"%s\" in %v",
			method,
			statusCode,
			http.StatusText(statusCode),
			requestURI,
			time.Since(start),
		)

		if size != 0 {
			content = fmt.Sprintf("%s size: %d bytes", content, size)
		}

		switch {
		case statusCode < 200:
			content = fmt.Sprintf("\037[1;30m%s\033[0m", content)
		case statusCode < 300:
			content = fmt.Sprintf("\033[1;32m%s\033[0m", content)
		case statusCode < 400:
			content = fmt.Sprintf("\033[1;35m%s\033[0m", content)
		case statusCode < 500:
			content = fmt.Sprintf("\033[1;31m%s\033[0m", content)
		case statusCode < 600:
			content = fmt.Sprintf("\033[1;36m%s\033[0m", content)
		}

		fmt.Sprintf("\033[1;30m[%s]\033[0m]", time.Now().Format(logTimeFormat), content)
		fmt.Println(content)
	}
}
