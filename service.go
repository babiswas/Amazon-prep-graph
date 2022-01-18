package main
import "fmt"
import "bufio"
import "os"

type Service interface{
   service() string
}

type User struct{
   name string
}

type Employee struct{
   role string
}

func (u *User)service()string{
   return u.name
}

func (e *Employee)service()string{
   return e.role
}

func main(){
  m:=make(map[string]Service)
  m["user"]=&User{"bapan"}
  m["employee"]=&Employee{"engineer"}
  scanner:=bufio.NewScanner(os.Stdin)
  for{
      fmt.Println("Enter the key")
      scanner.Scan()
      key:=scanner.Text()
      fmt.Println(m[key].service())
  }
}