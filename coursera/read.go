package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)
const (
    maxLength = 20
)
type fullname struct{
	fname string
	lname string
}
func main(){
	var filename string
	sli := make([]fullname,1)
	fmt.Println("Enter the name of the file to be read")
	fmt.Scan(&filename)
	f,_ := os.Open(filename)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)    
	
	for scanner.Scan() {
        
		fn := strings.Split(scanner.Text()," ")[0]
		ln := strings.Split(scanner.Text()," ")[1]
		var f fullname
		if len(fn) > maxLength {
			f.fname = fn[:maxLength]
        }else{
			f.fname = fn
		}
    	if len(ln) > maxLength {
        	f.lname = ln[:maxLength]
    	}else{
			f.lname = ln
		}	
		sli = append(sli,f)
    }
	for _,x := range sli{
		fmt.Println(x)
	}
}