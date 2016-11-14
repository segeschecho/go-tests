package main

import (
	"net/http"
	"fmt"
	"github.com/segeschecho/go-tests/mocking/service_defs"
)

func main(){
	http.HandleFunc("/service", service_handler)
	http.ListenAndServe(":8080", nil)
}


func service_handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query()["nombre"][0]
	surname := r.URL.Query()["apellido"][0]

	res := service_defs.Render_service(name, surname)

	fmt.Fprintf(w, res)
}
