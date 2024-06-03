package webserver

import (
	"fmt"
	"net/http"
)

func MyWebServer() {
	fs := http.FileServer(http.Dir("./webserver/assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", home)

	fmt.Println("Server started at port 3000")
	http.ListenAndServe(":3000", nil)
}

func home(writer http.ResponseWriter, req *http.Request) {
	http.ServeFile(writer, req, "./webserver/index.html")
}
