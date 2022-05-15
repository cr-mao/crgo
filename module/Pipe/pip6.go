package main

import (
	"encoding/csv"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	var err error
	db, err = gorm.Open("mysql",
		"root:@tcp(localhost:3306)/huan?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	db.SingularTable(true)
	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(20)
	return db
}

type Book struct {
	UserId   int    `gorm:"column:user_id"`
	Nickname string `gorm:"column:nickname"`
}
type BookList struct {
	Data []*Book
	Page int
}

const sql = "select * from user_weixin order by user_id limit ? offset ? "

func ReadData() InChan {
	page := 1
	pagesize := 200
	var in = make(chan *BookList)
	go func() {
		defer close(in)
		for {
			booklist := &BookList{make([]*Book, 0), page}
			db := GetDB().Raw(sql, pagesize, (page-1)*pagesize).Find(&booklist.Data)
			if db.Error != nil || db.RowsAffected == 0 {
				break
			}
			page++
			in <- booklist
		}
	}()
	return in
}

type Result struct {
	Page int
	Err  error
}
type InChan chan *BookList
type OutChan chan *Result

//管道函数
func WriteData(in InChan) OutChan {
	out := make(OutChan)
	go func() {
		defer close(out)
		for d := range in {
			out <- &Result{Page: d.Page, Err: SaveData(d)}
		}
	}()
	return out
}

type DataPipeCmd func(in InChan) OutChan
type DataCmd func() InChan

//多路复用
func Pipe3(c1 DataCmd, cs ...DataPipeCmd) OutChan {
	in := c1()
	out := make(OutChan)
	wg := sync.WaitGroup{}
	for _, c := range cs {
		getChan := c(in)
		wg.Add(1)
		go func(input OutChan) {
			defer wg.Done()
			for v := range input {
				out <- v
			}
		}(getChan)
	}
	go func() {
		defer close(out)
		wg.Wait()
	}()
	return out
}

//写入到csv文件
func SaveData(data *BookList) error {
	file := fmt.Sprintf("csv/%d.csv", data.Page)
	csvFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer csvFile.Close()
	w := csv.NewWriter(csvFile) //创建一个新的写入文件流
	header := []string{"user_id", "book_name"}
	export := [][]string{
		header,
	}
	for _, d := range data.Data {
		cnt := []string{
			strconv.Itoa(d.UserId),
			d.Nickname,
		}
		export = append(export, cnt)
	}
	err = w.WriteAll(export)
	if err != nil {
		return err
	}
	w.Flush()
	return nil
}

func testData() {
	start := time.Now().Unix()
	for item := range Pipe3(ReadData, WriteData, WriteData) {
		fmt.Println(item.Page, "ok")
	}
	end := time.Now().Unix()
	fmt.Printf("测试--用时:%d秒\r\n", end-start)
}

func main() {
	testData()
}
