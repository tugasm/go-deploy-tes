package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}
	http.HandleFunc("/", helloHandler)
	log.Println("Listing for" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

//package main
//
//import (
//"io"
//"log"
//"net/http"
//"os"
//)
//
//func main() {
//	//helper.Sum(5, 5)
//}
