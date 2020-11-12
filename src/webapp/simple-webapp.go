// simple webapp
package main

import (
	"fmt"
	"net/http"
)

func enGreetings(responseWriter http.ResponseWriter, req *http.Request) {
	responseWriter.Write([]byte("<h1>hello there!</h1>"))
}

func frGreetings(responseWriter http.ResponseWriter, req *http.Request) {
	responseWriter.Write([]byte("<h1>bonjour! ça va bien?</h1>"))
}

func pbGreetings(responseWriter http.ResponseWriter, req *http.Request) {
	// inbuilt UTF8 support
	responseWriter.Write([]byte("<h1>ਸਤ ਸ੍ਰੀ ਅਕਾਲ</h1>"))
}

func main() {
	http.HandleFunc("/english", enGreetings)
	http.HandleFunc("/french", frGreetings)
	http.HandleFunc("/punjabi", pbGreetings)

	fmt.Println("serving @localhost:1514")
	http.ListenAndServe("localhost:1514", nil)
}
