package giny
// import package
import (
  "net/http"
)
type User struct  {
    cookie string
}

type Session struct  {
    user User
}
/**
 * Start the session
 */
func (s Session) start (){

}
/**
 * log user in session
 */
func (s Session) LoginUser (r *http.Request, data  map[string]int){
  s.user = User {"cookieId"}
}


/**
 * log user in session
 */
func (s Session) LogoutUser (r *http.Request, data  map[string]int){
  // s.user = {}
}
