package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello is simple handler
type Hello struct {
	logger *log.Logger
}

// NewHello creates a new hello handler with the given logger
func NewHello(logger *log.Logger) *Hello {
	return &Hello{logger: logger}
}

// ServeHTTP implements the go http.Handler interface
// https://golang.org/pkg/net/http/#Handler
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// read the body
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.logger.Println("Error reading request body", err)
		rw.WriteHeader(http.StatusBadRequest)
		_, _ = rw.Write([]byte("Error reading request body"))
		return
	}
	h.logger.Println(reqBody)

	// write the response
	_, _ = fmt.Fprintf(rw, "Hello %s\n", reqBody)
}