package main
import "fmt"
import "time"

func main(){
  t:=time.Now()
  fmt.Println(t)
  sum:=0
  arr:=[10]int{1,2,3,4,5,6,7,8,9,10}
  for i:=0;i<len(arr);i++{
    sum+=arr[i]
  }
  fmt.Println(sum)
  fmt.Println(time.Since(t))
}