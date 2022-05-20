package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("024前缀树.go")
	if err != nil {
		log.Fatal(err)
	}

	scaner := bufio.NewScanner(file)
	count := 0
	for scaner.Scan() {
		count++
		fmt.Println(scaner.Text())
	}

	fmt.Println("一共", count, "行")

}
