package giny
/**
 * Base application code
 */
import(
  "net/http"
  //"strings"
)
/**
 * Define the application structure
 */
type Application struct {
  routes map[string]*Router
}
/**
 * Start the application
 */
func (this Application) Run (){
    StartServer(8000, &this);
}
/**
 * RegisterRouter
 */
func (this Application) RegisterRouter (pattern string,  route *Router){
    this.routes[pattern] = route
    route.basePattern = pattern
}
/**
 * Handling route
 */
func (this Application) HandleRequest(w http.ResponseWriter, request *http.Request){
    CleanLogStack ()
    Log ("handling route " + request.URL.Path)
    url := request.URL.Path
    FindRouter  (w, request, url, this.routes)
}
/**
 * Create and return a new application
 */
func NewApplication () *Application{
  var _app = new (Application)
  _app.routes = make(map[string]*Router)
  return _app
}
