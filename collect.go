package nowcoder

import ("regexp"
"fmt"
"os"
"log")

var idre ,_=regexp.Compile("-[0-9]\">.+")
var notfound,_=regexp.Compile("<title>牛客网-用户不存在</title>")
//Matchpage s
func Matchpage(respose *[]byte) {
	if notfound.Match(*respose){
			fmt.Println("404")
		}else {
			id:=idre.Find(*respose)
			fmt.Printf("%s",id)
			file, err := os.OpenFile("nowcoder.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			num, err := file.Write(id)
			if num != len(id) {
				log.Fatal(err)
			}	
		}
}