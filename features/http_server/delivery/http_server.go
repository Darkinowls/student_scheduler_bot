package delivery

import (
	"fmt"
	"io"
	"net/http"
)

func ServeHttpServer(port *string) {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, HTTP!\n")
	})

	http.ListenAndServe(fmt.Sprintf(":%s", *port), nil)
}
