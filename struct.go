package main
import "fmt"

type Journal struct{
  articles [5]string
}

func (j *Journal)add_article(arr ...string){
   index:=0
   for _,val:=range arr{
     j.articles[index]=val
     index=index+1
   }
}

func (j Journal)display(){
   for _,str:=range j.articles{
      fmt.Println(str)
   }
}

func main(){
  j:=&Journal{}
  j.add_article("hello","bello","mello","tello","kello")
  (*j).display()
}