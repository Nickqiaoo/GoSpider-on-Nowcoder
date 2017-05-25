package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"nowcoder"
	"strconv"
)

func main(){
	url:="https://www.nowcoder.com/"
	cli:=http.Client{}
	for i:=0;i<1000;i++{	
		response,_:=cli.Get(url+strconv.Itoa(i))
		fmt.Println(response.Status+"\n")
		body,_:=ioutil.ReadAll(response.Body)
		nowcoder.Matchpage(&body)
	}
	db,err:=sql.Open("mysql", "root:6055965@/nowcoder")
	err = db.Ping()
	if err != nil {
	    log.Fatal(err)
	}
}