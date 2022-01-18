package main
import "fmt"
import "os"
import "path/filepath"
import "encoding/csv"


func get_all_paths(path string)([]string,[]string){
   path_list:=make([]string,1)
   directory_list:=make([]string,1)
   filepath.Walk(path,func(path string,info os.FileInfo,err error)error{
      if info.IsDir(){
         directory_list=append(directory_list,path)
      }else{
         path_list=append(path_list,path)
      }
      return nil
   })
   return path_list,directory_list
}

func create_records(record []string,filename string){
   file,err:=os.OpenFile(filename,os.O_CREATE|os.O_APPEND|os.O_WRONLY,0777)
   defer file.Close()
   if err!=nil{
      fmt.Println("Error Occured")
      return
   }
   header:=[]string{"Filenames\n"}
   w:=csv.NewWriter(file)
   w.Flush()
   defer w.Flush()
   w.Write(header)
   for _,line:=range record{
      _,err:=file.WriteString(line+"\n")
      if err!=nil{
         fmt.Println("Error Occured")
      }
   }
}

func main(){
   if len(os.Args)<1{
      fmt.Println("Please provide file path")
      return
   }
   path_list,directory_list:=get_all_paths(os.Args[1])
   create_records(path_list,"Filenames.csv")
   create_records(directory_list,"Directories.csv")
}