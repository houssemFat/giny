package giny
import (
  "fmt"
)

var GlobalLog string

func init() {
    GlobalLog = ""
}
/**
 *
 */
func Log(message string)  {
    GlobalLog = GlobalLog + message + "\n"
    fmt.Printf(message + "\n")
}
/**
 *
 */
func GetLogStack () string {
    return GlobalLog
}

/**
 *
 */
func CleanLogStack () {
   GlobalLog = ""
}
