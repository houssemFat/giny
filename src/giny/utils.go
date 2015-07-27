package giny

import (
    "fmt"
    "regexp"
    "strings"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

/**
 * check the current pattern with
 */
func CheckPathAgainstUrl (path string, currentUrl string, exact bool ) (bool, map[string]string) {
  pattern := FormatRegexpPattern (path, exact)
  Log ("CheckPathAgainstUrl")
  Log ("pattern composed " + pattern )
  Log ("search in " + currentUrl )
  r, err := regexp.Compile(pattern)
  if err != nil {
      fmt.Printf(" %q", err)
    }
    // Will print 'Match'
    if r.MatchString(currentUrl) == true {
      match  := r.FindStringSubmatch(currentUrl)
      result := make(map[string]string)
      for i, name := range r.SubexpNames() {
         result[name] = match[i]
      }
      return true, result
      } else {
      return false , nil
      }

}
/**
 * return a build regexp pattern from a string
 * app/:id/:test
 * to app/(?P<id>[\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?/(?P<test>[\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?
 */
func FormatRegexpPattern (path string, excat bool) string{
   // api/me/photo/:id
   // api/me/photo/55
   any := `>[\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?`
   paths := strings.Split(path, "/")
   regexpPattern := ""
   slag := ""
  	for i:= range paths {
  		slag = paths [i]
      // ignore empty string
      if slag == "" {
  		   continue
  		}
      // check if it's a params
  		if strings.HasPrefix(slag , ":") == true {
  			slag = `(?P<` + slag[1:] + any
  		}
  		regexpPattern +=  slag + "/"
  	}
    // check if path starts with '/'
    if strings.HasPrefix(path, "/") == true {
      regexpPattern = "/" + regexpPattern
    }
    // check if we need exact match
    // used for inner router route
    if (excat == true){
      regexpPattern = "^" + regexpPattern
      if strings.HasSuffix(regexpPattern, "/") == false {
        regexpPattern += "/"
      }

      regexpPattern += "{0,1}$"
    }
    //Log ("composed " + regexpPattern )
    return regexpPattern
}

/**
 * Handling no route
 */
func NoRouteFoundException (w http.ResponseWriter){
    w.WriteHeader(404)
    w.Write([]byte(GetLogStack()))
}
/**
 *
 */
func FindRouter (w http.ResponseWriter, request *http.Request, url string, routes map[string]*Router){
  //logErrors :=  "stack :\n"
  // initialize
  for pattern, route  := range routes {
     Log ("evaluating : " +  pattern )
     Log ("formatted regxp " + FormatRegexpPattern (pattern, false))
     match , params :=  CheckPathAgainstUrl(pattern, url, false)
      if match == true {
        Log ("founded route " + pattern )
        route.handle(w, request, params)
        return
      }
   }
   Log ("evaluating all routes but no route match ")
   // any way handle error
   NoRouteFoundException (w)
}
/**
 *
 */
func GetJsonServerResponse (url string) string{
  response, err := http.Get(url)
  if err != nil {
        panic(err.Error())
  }
  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
      panic(err.Error())
  }
  var data interface{} // TopTracks
  err = json.Unmarshal(body, &data)
  if err != nil {
      panic(err.Error())
  }
  dataResponse , err := json.Marshal (data)
  if err != nil {
      panic(err.Error())
  }
  return string (dataResponse)
}
