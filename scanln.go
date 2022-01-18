package main
import "fmt"
import "time"

func main(){
  var str [5]string
  t:=time.Now()
  for i:=0;i<5;i++{
    fmt.Println("Enter the string")
    fmt.Scanln(&str[i])
  }
  for arr:=range str{
    fmt.Println(arr)
  }
  fmt.Println(time.Since(t))
}