package main

import (
	"log"
)

func main() {
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
