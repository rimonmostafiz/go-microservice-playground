package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	logger *log.Logger
}

func NewHello(logger *log.Logger) *Hello {
	return &Hello{logger: logger}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.logger.Println("Error reading request body", err)
		rw.WriteHeader(http.StatusBadRequest)
		_, _ = rw.Write([]byte("Error reading request body"))
		return
	}
	h.logger.Println(reqBody)
	_, _ = fmt.Fprintf(rw, "Hello %s\n", reqBody)
}