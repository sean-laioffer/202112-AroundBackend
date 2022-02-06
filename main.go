package main

import (
    "fmt"
    "log"
    "net/http" 

    "around/handler"   
)
func main() {
    fmt.Println("started-service")
    log.Fatal(http.ListenAndServe(":8080", handler.InitRouter()))
}