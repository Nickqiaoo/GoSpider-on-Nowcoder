package nowcoder

import (
	"fmt"
	"regexp"
	//"os"
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //sql
	"log"
	"sync"
)
//Wg var
var Wg sync.WaitGroup
var idre, _ = regexp.Compile("-[0-9]\">.+")
var notfound, _ = regexp.Compile("<title>牛客网-用户不存在</title>")
var cjre, _ = regexp.Compile("font-green\">[0-9]+")
var db, err = sql.Open("mysql", "root:6055965@tcp(127.0.0.1:3306)/nowcoder?charset=utf8")
var stmt, _ = db.Prepare(`INSERT per (uid,username,level,chengjiu) VALUES (?,?,?,?)`)
var lock sync.Mutex 
//Matchpage s
func Matchpage(respose *[]byte,i int) {
	if notfound.Match(*respose) {
		fmt.Println("404")
	} else {
		id := idre.Find(*respose)
		if len(id)<1{
			return
		}
		level := int(id[1]) - 48
		id = id[4:]
		chengjiu := cjre.Find(*respose)
		chengjiu = chengjiu[12:]
		sum := 0
		for i := range chengjiu {
			sum = sum*10 + (int(chengjiu[i]) - 48)
		}
		/*fmt.Printf("%s",id)
		file, err := os.OpenFile("nowcoder.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		num, err := file.Write(id)
		if num != len(id) {
			log.Fatal(err)
		}*/
		stmt.Exec(i, id, level, sum)
		fmt.Printf("%d %d %s %d\n",i,level,id,sum)
	}
	defer Wg.Done()
}
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func convert(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}
