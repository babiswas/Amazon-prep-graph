package main
import "sync"
import "fmt"

func main(){
  var wg sync.WaitGroup
  wg.Add(1)
  go func(wg *sync.WaitGroup){
    defer wg.Done()
    for i:=0;i<35;i++{
       fmt.Println("hello")
    }
  }(&wg)
  wg.Add(1)
  go func(wg *sync.WaitGroup){
    defer wg.Done()
    for i:=0;i<35;i++{
      fmt.Println("bello")
    }
  }(&wg)
  wg.Wait()
}

