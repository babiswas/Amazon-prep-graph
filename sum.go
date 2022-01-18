package main
import "fmt"
import "time"

func add(arr ...int){
  sum:=0
  for _,num:=range arr{
    sum+=num
  }
  fmt.Println(sum)
}

func main(){
  t:=time.Now()
  arr:=make([]int,10)
  for i:=0;i<10;i++{
    arr[i]=10+i
  }
  add(arr ...)
  fmt.Println(arr)
  fmt.Println(time.Since(t))
}