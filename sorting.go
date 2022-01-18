package main
import "fmt"
import "sort"

func main(){
   nums:=[]int{6,4,1,2,5,3}
   fmt.Println(nums)
   sort.Ints(nums)
   fmt.Println(nums)
}