package main

import (
  "giny"
  "fmt"
  //"net/http"
  "./api"
)

func main() {
    fmt.Println("starting server on port 8000")
    //
    app := giny.NewApplication ()
    /**
     * Register route
     */
    base := giny.NewRouter ()
    // register
    app.RegisterRouter ("/api", api.GetUserRouter())

    app.RegisterRouter ("/", base);

    // finally start your application
    app.Run ()
}
