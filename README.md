# Goober

Tired of none-sense Makefile syntax and I want to build and configure my application using golang not making
complex Makefile and bash scripts. This project is created out of THE frusration of dealing with bash and
makefile syntax.

The fun part is you can `Yum` boring build process and `Burp` the result :P

```go
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

	err = peanut.
		Yum(`

      touch ./awesome2.txt

    `).
		Burp()

	if err != nil {
		log.Fatal(err)
	}
}

```
