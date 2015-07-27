# giny
a tiny go framework (learning purpose)
usage
```go
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
```
## Application
```go
  giny.NewApplication ()
```
### Register a route to application

```go
  giny.RegisterRouter (path, router)
```

## Router
```go
  giny.NewRouter ()
```
