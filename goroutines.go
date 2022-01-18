package main
import "fmt"
import "sync"

func main(){
  ch1:=make(chan int,10)
  ch2:=make(chan int,10)
  task1:=false
  task2:=false
  var wg sync.WaitGroup
  wg.Add(1)
  go func(wg *sync.WaitGroup,ch1 chan<- int,task1 *bool){
     fmt.Println("Executor1 is sending")
     defer wg.Done()
     for i:=0;i<20;i++{
       ch1<-i
     }
     *task1=true
     close(ch1)
  }(&wg,ch1,&task1)
  wg.Add(1)
  go func(wg *sync.WaitGroup,ch2 chan<- int,task2 *bool){
    fmt.Println("Executor 2 is sending")
    defer wg.Done()
    for i:=0;i<20;i++{
      ch2<-i
    }
    *task2=true
    close(ch2)
  }(&wg,ch2,&task2)
  wg.Add(1)
  go func(wg *sync.WaitGroup,task1 *bool,task2 *bool){
     defer wg.Done()
     for{
         select{
            case mg1:=<-ch1:
               fmt.Println(mg1)
               for a:=range ch1{
                  fmt.Println(a)
               }
            case mg2:=<-ch2:
               fmt.Println(mg2)
               for a:=range ch2{
                  fmt.Println(a)
               }
            default:
              fmt.Println("No message recieved")
         }
         if (*task1==true)&& (*task2==true){
            break
         }
     }
  }(&wg,&task1,&task2)
  wg.Wait()
}