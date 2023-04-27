package utils

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)


func StrToNumTheObj(s string)interface{}{

	if s == "" {
		return 0
	}

	if strings.Contains(s, "-") || strings.Contains(s, ",") || strings.Contains(s, "") {
		str := strings.ReplaceAll(s, ",", "")
		if len(str) == 1 {
			str = strings.ReplaceAll(str, "-", "0")
		}

		if strings.Contains(str, ".") {
			d, err := strconv.ParseFloat(str, 32)
			if err != nil {
				log.Fatal(err)
			}
			return d
		}
		d, err := strconv.ParseInt(str, 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		return d
	}
	if strings.Contains(s, ".") {
		d, err := strconv.ParseFloat(s, 32)
		if err != nil {
			log.Fatal(err)
		}
		return d
	} else {
		d, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		return d
	}


}

//  print reverse after getting index of td
func PrintReverseDataByte(str string,idx int)string{

	if(string(str[idx])==">"){

		return ""
	}
	return PrintReverseDataByte(str,idx-1)+string(str[idx])
	
}

func SplitRawStr(splitChar string,rawtext string,putToList *[]string){

	*putToList=strings.Split(rawtext, splitChar)
	
}

// calculate how much td is there
func CalcHowMuchTd(str string)int{
	
	substr:="</td>"
	idx:=strings.Index(str,substr)
	if idx>0{
		return 1+CalcHowMuchTd(str[idx+4:])
	}else{
		return 0
	}
}



func ExecuteMongoScript(mongoScriptArgs []string){


	cmd := exec.Command("mongosh",mongoScriptArgs...)
	 err := cmd.Run()
	if err != nil {
		fmt.Println("mongosh script not executed")
		fmt.Print(err.Error())
		return
	} else {
		fmt.Print("mongosh script executed\n")
	}
	


}