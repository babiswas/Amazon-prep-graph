package main
import "fmt"
import "encoding/json"


type Person struct{
   Name string `json:"name"`
   Id int `json:"id"`
}

func main(){
  var person [5]Person
  var mydata [5]interface{}
  person[0]=Person{"bapan",1}
  person[1]=Person{"tapan",2}
  person[2]=Person{"madan",3}
  person[3]=Person{"chandan",4}
  person[4]=Person{"salim",5}
  for index,p:=range person{
     data,_:=json.Marshal(p)
     mydata[index]=string(data)
     fmt.Println(mydata[index])
  }
  for _,data:= range mydata{
    p1:=Person{}
    bytedata:=json.RawMessage(data.(string))
    bdata,err:=bytedata.MarshalJSON()
    if err!=nil{
      fmt.Println("Error occured")
    }
    err=json.Unmarshal(bdata,&p1)
    fmt.Println(p1)
  }
}


