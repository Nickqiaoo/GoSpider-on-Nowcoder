package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"nowcoder"
	"os"
	"strconv"
	"time"
)
func main() {
	url := "https://www.nowcoder.com/"
	cli := http.Client{}
	file, err := os.OpenFile("nowcoder.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	t := time.Now()
	num, err := file.WriteString(t.String()+"\n")
	//if num != len(t.String()) {
		//log.Fatal(err)
	//}
	for i := 1; i < 200; i++ {
		if i==404{
			continue
		}
		response, _ := cli.Get(url + strconv.Itoa(i))
		if response == nil{
			continue
		}
		fmt.Println(response.Status)
		body, _ := ioutil.ReadAll(response.Body)
		if response.StatusCode == 200 {
			nowcoder.Wg.Add(1)
			nowcoder.Matchpage(&body,i)
		}
	}
	t = time.Now()
	num, err = file.WriteString(t.String())
	if num != len(t.String()) {
		log.Fatal(err)
	}
	nowcoder.Wg.Wait()

}
