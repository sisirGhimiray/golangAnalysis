package main

import (

	// "os/exec"

	"fmt"
	"os"
	"strings"
	"time"

	"example.com/golang/db"
	makestockList "example.com/golang/makeStockList"
)

// There is two sign "pre" and "nif" one for premarket Data and another for current nifty
func ProcessFile(fileName string,sign string){

	makestockList.ReadFile(fileName)
	
	makestockList.Process(sign)
}

func main() {

// Watch the directory if it has file or new file added
// process the file if added





	
		files,err:=os.ReadDir("./rawData")
		if err!=nil{
			panic(err)
		}

	
			for _,file:=range files{
				
				fileNameDir:=fmt.Sprintf("./rawData/%s",file.Name())
				// This is the function where we process file raw Text and make object
				ProcessFile(fileNameDir,file.Name()[0:3])
				time.Sleep(time.Second*1)

				
				// this is taking the file path name
				dirPlusFileName:=fmt.Sprintf("./result/%s.%s",strings.Split(file.Name(), ".")[0],"json")
				// This function takes file path as first argument and second the prefix of file
				db.ConnectToDatabaseAddFile(dirPlusFileName,file.Name()[0:3])
				fileOpForDelTxt,err:=os.OpenFile(fileNameDir,os.O_WRONLY|os.O_TRUNC,0644)
				if err!=nil{
					panic(err)
				}
				defer fileOpForDelTxt.Close()

				fileOpForDelTxt.Truncate(0)
				}	
			
	}



