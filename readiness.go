package main

import (
	"net/http"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
  w.Header().Add("Content-Type", "text/html; charset=utf8")
  w.WriteHeader(http.StatusOK)
  w.Write([]byte(http.StatusText(http.StatusOK)))
}
