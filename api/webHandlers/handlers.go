package webHandlers

import (
	"fmt"
	"net/http"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Ha entrado en handler", r.Method)
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("<h1>Hello, World!</h1>"))

}
