package main
import "fmt"
import "encoding/json"
import "strconv"

type Person struct{
  Name string `json:"name"`
  Id int   `json:"id"`
}

type Human struct{
  Name string `json:"name"`
  Profession string `json:"profession"`
}

func (p Person)get_person()string{
   return p.Name+" "+strconv.Itoa(p.Id)
}

func (h Human)get_human()string{
   return h.Name+" "+h.Profession
}

func get_object(i interface{}){
   if obj,ok:=i.(Person);ok{
      fmt.Println(obj.get_person())
   }
   if obj,ok:=i.(Human);ok{
      fmt.Println(obj.get_human())
   }
}


func main(){
   p:=Person{"bapan",12}
   h:=Human{"bapa","Engineer"}
   fmt.Println(p.get_person())
   fmt.Println(h.get_human())
   obj,_:=json.Marshal(p)
   fmt.Println(string(obj))
   obj,_=json.Marshal(h)
   fmt.Println(string(obj))
   get_object(p)
   get_object(h)
}