package main
import "encoding/json"
import "fmt"
import "bufio"
import "os"

type Student struct{
   Name string `json:"name"`
   Id int   `json:"id"`
   Email string `json:"email"`
}

func main(){
   var students [5]Student
   for i:=0;i<len(students);i++{
      scanner:=bufio.NewScanner(os.Stdin)
      fmt.Println("Enter the student name:")
      scanner.Scan()
      name:=scanner.Text()
      fmt.Println("Enter the email")
      scanner.Scan()
      email:=scanner.Text()
      students[i]=Student{name,i,email}
   }
   for i:=0;i<5;i++{
      fmt.Println(students[i])
   }
   for index,stu:=range students{
     data,err:=json.Marshal(stu)
     fmt.Println(index)
     if err!=nil{
        print("Error occured")
     }
     fmt.Println(string(data))
   }
}