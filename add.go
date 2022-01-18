package main
import "fmt"
import "bufio"
import "strconv"
import "os"

func main(){
   scanner:=bufio.NewScanner(os.Stdin)
   fmt.Println("Enter the first number")
   scanner.Scan()
   num1,_:=strconv.ParseInt(scanner.Text(),10,64)
   fmt.Println("Enter the second number")
   scanner.Scan()
   num2,_:=strconv.ParseInt(scanner.Text(),10,64)
   fmt.Println(num1+num2)
}