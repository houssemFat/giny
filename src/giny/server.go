package giny
// import package
import (
  "fmt"
  "net/http"
  "strconv"
)
/**
 * Start  a server at port
 * @port
 */
func StartServer (port int, app *Application){
      fmt.Println("extracting %s", "me- and others"[:3])
      http.HandleFunc("/", app.HandleRequest)
      http.ListenAndServe(":" + strconv.Itoa(port), nil)
}
