package main
import "fmt"
import "sync"

func main(){
  tasklist:=[10]string{"task1","task2","task3","task4","task5","task6","task7","task8","task9","task10"}
  var wg sync.WaitGroup
  var mu sync.Mutex
  wg.Add(1)
  go executor(tasklist,0,5,&wg,&mu)
  wg.Add(1)
  go executor(tasklist,5,10,&wg,&mu)
  wg.Wait()
}

func executor(task [10]string,index1 int,index2 int,wg *sync.WaitGroup,mu *sync.Mutex){
   defer wg.Done()
   for i:=index1;i<index2;i++{
      mu.Lock()
      fmt.Println(task[i])
      mu.Unlock()
   }
}