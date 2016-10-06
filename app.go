package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "gopkg.in/resty.v0"
)

func main() {
    http.HandleFunc("/", hello)
    http.HandleFunc("/javafortune", javafortune)
    http.HandleFunc("/rubyfortune", rubyfortune)
    http.HandleFunc("/pythonfortune", pythonfortune)
    http.HandleFunc("/nodefortune", nodefortune)
    http.HandleFunc("/dockerfortune", dockerfortune)
    http.HandleFunc("/health", health)
    err := http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(w, "hello, world!")
}

func health(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("content-type", "application/json")
    fmt.Fprintln(w, "{\"status\": \"UP\"}")
}


func fortune(w http.ResponseWriter, url string) {
    resp, err := resty.R().Get(url)
    if err != nil {
        fmt.Fprintln(w, err)
    } else {
        w.Header().Set("content-type", "application/json")
        fmt.Fprintln(w, resp)
    }
}


func javafortune(w http.ResponseWriter, req *http.Request) {
    fortune(w, "http://localhost:8087/fortunes/random")
}

func rubyfortune(w http.ResponseWriter, req *http.Request) {
    fortune(w, "http://localhost:8087/ruby-demo/javafortune")
}

func pythonfortune(w http.ResponseWriter, req *http.Request) {
    fortune(w, "http://localhost:8087/python-demo/javafortune")
}

func nodefortune(w http.ResponseWriter, req *http.Request) {
    fortune(w, "http://localhost:8087/node-demo/javafortune")
}

func dockerfortune(w http.ResponseWriter, req *http.Request) {
    fortune(w, "http://localhost:8087/docker-demo/javafortune")
}
