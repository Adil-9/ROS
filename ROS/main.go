package main

import (
	"fmt"
	"net/http"
	"ros/web"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", web.Home)

	mux.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates/"))))
	fmt.Printf("Server runnig on: http://localhost:6060")

	http.ListenAndServe(":6060", mux)
}
