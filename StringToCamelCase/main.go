package kata
import (
  //"fmt"
  "strings"
  "unicode"
)



func ToCamelCase(s string) string {
	// your code
  
  f := func(c rune) bool{
    return c == '_' || c == '-' 
  }
  
  //fmt.Printf("%q", strings.FieldsFunc(s, f))
  
  var result string
  
  for index, elem := range strings.FieldsFunc(s, f){
    
    if index == 0 && !(unicode.IsUpper([]rune(elem)[0])) {
      result += strings.ToLower(elem)
      continue
    }
    result += strings.Title(elem)
  }
  
	return result
}