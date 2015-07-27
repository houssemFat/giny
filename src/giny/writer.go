package giny

import (
  "net/http"
)
/**
 *
 */
func SendJsonResponse (w http.ResponseWriter, data string){
  w.Header().Set("Content-Type", "application/json")
  w.Write([]byte(data))
  //w.Write([]byte("********************LOG*******************\n"))
  //w.Write([]byte(GetLogStack()))
}
