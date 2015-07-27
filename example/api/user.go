package api

import (
  "giny"
  "net/http"
  "../models"
  "encoding/json"
   //"strconv"
)
/**
 * login
 */
func UserLogin  (w http.ResponseWriter, request *http.Request, params map[string]string){
    params["response"] = "login"
    data,  err := json.Marshal(params)
    if (err != nil){

    }
    giny.SendJsonResponse (w,  string(data))
}
/**
 * informations
 */
func UserInformations (w http.ResponseWriter, request *http.Request, params map[string]string){
  user := &models.User {Name :"houssem", Username :"houssemFat", Id : 50 }
  data,  err := json.Marshal(user)
  if (err != nil){

  }
  giny.SendJsonResponse (w, string(data))
}
/**
 * register
 */
var UserRegister = func (w http.ResponseWriter, request *http.Request, params map[string]string){
  params["response"] = "register"
  data,  err := json.Marshal(params)
  if (err != nil){

  }
  giny.SendJsonResponse (w,  string(data))
}
/**
 * get list
 */
var GetList = func (w http.ResponseWriter, request *http.Request, params map[string]string){
  url := "http://jsonplaceholder.typicode.com/users"
  response := giny.GetJsonServerResponse(url)
  // id, err := strconv.Atoi(params["id"])
  giny.SendJsonResponse (w,  response)
}
/**
 * get item based in id
 */
var GetItem = func (w http.ResponseWriter, request *http.Request, params map[string]string){
  url := "http://jsonplaceholder.typicode.com/comments?postId=" + params ["id"]
  response := giny.GetJsonServerResponse(url)
  // id, err := strconv.Atoi(params["id"])
  giny.SendJsonResponse (w,  response)
}
/**
 *
 */
func GetUserRouter () *giny.Router {
    router := giny.NewRouter ()
    router.On ("/login", UserLogin)
    router.On ("/register", UserRegister)
    router.On ("/info", UserInformations)
    router.On ("/item/:id", GetList)
    router.On ("/list", GetList)
    router.On ("/", UserInformations)
    return router
}
