package main

import (
    "fmt"
    "os"

    "github.com/JunichiMitsunaga/dagger_example/hello/greeting"
)

func main() {
    name := os.Getenv("NAME")
    if name == "" {
        name = "Hoge"
    }
    fmt.Printf(greeting.Greeting(name))
}
