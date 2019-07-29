package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	nf, err2 := os.Create("logFile.txt")
	fmt.Println("TEST")
	if err2 != nil {
		fmt.Println("err happened", err2)
	}
	log.SetOutput(nf)

}

func main() {
	_, err := os.Open("tester.txt")
	if err != nil {
		log.Println("err happened", err)
		//fmt.Println("err happened again", err)
	}

}
