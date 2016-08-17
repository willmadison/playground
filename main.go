package main

import (
	"encoding/json"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/playground/types"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)
		response.Write([]byte(`{"greeting": "hello, world!"}`))
	})

	r.HandleFunc("/order/{orderNumber}", func(response http.ResponseWriter, request *http.Request) {
		pathParams := mux.Vars(request)

		orderNumber := pathParams["orderNumber"]

		rawResponse, _ := json.Marshal(types.Order{orderNumber})

		response.Header().Set("Content-Type", "application/json")

		response.WriteHeader(http.StatusOK)
		response.Write(rawResponse)
	}).Methods(http.MethodGet)

	hostPort := net.JoinHostPort("", os.Getenv("PORT"))

	http.ListenAndServe(hostPort, r)
}
