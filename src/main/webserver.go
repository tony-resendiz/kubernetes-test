package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server now listening on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, getJson())
}

type ExampleJson struct {
    UserId int `json:"userId"`
    Id int `json:"id"`
    Title string `json:"title"`
    Body string `json:"body"`
}

func getJson() string {

    resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
    if err != nil { 
        panic(err) 
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    var p ExampleJson 
    err = json.Unmarshal(body, &p)
    if err != nil {
        panic(err)
    }

    return fmt.Sprintf("body: %v", p.Body)
}