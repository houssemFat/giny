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
### initalisation
```go
  giny.NewRouter ()
```
### attach event
#### router.On(path, func (w http.ResponseWriter, request *http.Request, params map[string]string))

```go

/**
 * get item based in id
 * @param w
 * @param request
 * @param params map[string]string , dictionary of key, value
 */
var GetItem = func (w http.ResponseWriter, request *http.Request, params map[string]string){
  url := "http://jsonplaceholder.typicode.com/comments?postId=" + params ["id"]
  response := giny.GetJsonServerResponse(url)
  // id, err := strconv.Atoi(params["id"])
  giny.SendJsonResponse (w,  response)
}
// attach the event
// here the path will match base/item/50
  router.On ("/item/:id", GetItem)
```

### Utils
#### sending json response

```go
  giny.SendJsonResponse (w,  data string)
```
