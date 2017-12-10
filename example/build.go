package main

import (
	"log"

	"github.com/nulloop/goober"
)

func main() {
	peanut, err := goober.New()
	if err != nil {
		log.Fatal(err)
	}

	err = peanut.
		Yum("touch").
		Yum("./awesome.txt").
		Burp()

	if err != nil {
		log.Fatal(err)
	}

	peanut.Yum(`
		
		
		rm awesome.txt
		
	`).Burp()
}
