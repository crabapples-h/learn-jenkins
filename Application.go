package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":9999", nil)
}
func index(response http.ResponseWriter, request *http.Request) {
	data := fmt.Sprintf("%s:%s", "hello world1", time.Now().Format("2006-01-02 15:04:05"))
	response.Write([]byte(data))
}
