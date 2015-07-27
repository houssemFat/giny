package giny
// import package
import (
  "fmt"
  //"io/ioutil"
  "net/http"
)

func DefineHandler (route string, ha func(r *http.Request/*, */))  {
    http.HandleFunc(route, func (_w http.ResponseWriter, _r *http.Request) {
        fmt.Println("serving route %s", route)
        ha (_r)
    })
}

type Router struct  {
    basePattern string
    routes map[string]*Router
    callbacks map[string]func(w http.ResponseWriter, request *http.Request, params map[string]string)
}
/**
 * Handling route
 */
func (this Router) handle(w http.ResponseWriter, request *http.Request, params map[string]string){
   // remove base route
   // start handling routes
   Log("start evaluating " + this.basePattern)
    // url := request.URL.Path[len(this.basePattern):]
    url := request.URL.Path
     // check for callbacks
     for pattern, callback := range this.callbacks {
       Log("checking " + this.basePattern +  pattern)
       match , _params := CheckPathAgainstUrl( this.basePattern +  pattern,  url, true)
       if match == true {
         Log("found handler on " + pattern)
         callback(w, request, _params)
         return
       }
    }
    Log("no rooute found, search nested routes ")
    // check for nested routes
    FindRouter (w, request, url, this.routes)
}
/**
 *
 */
func (this Router) On(pattern string, callback func(w http.ResponseWriter, request *http.Request, params map[string]string)){
    this.callbacks[pattern] = callback
}
/**
 *
 */
func (this Router) RegisterRouter(pattern string, route *Router){
    this.routes[pattern] = route
    route.basePattern = this.basePattern + pattern
    // route.base
}
/**
 *
 */
func NewRouter () *Router{
  _route := new (Router)
  _route.callbacks = make(map[string]func(w http.ResponseWriter, request *http.Request, params map[string]string))
  _route.routes = make(map[string]*Router)
  return _route
}
