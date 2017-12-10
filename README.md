# Goober

Tired of none-sense Makefile syntax and I want to build and configure my application using golang not making
complex Makefile and bash scripts. This project is created out of THE frusration of dealing with bash and
makefile syntax.

The fun part is you can `Yum` boring build process and `Burp` the result

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

	peanut.
		Yum("touch").
		Yum("./awesome.txt").
    Burp()
}
```

At the moment there are 2 options available

1 - adding system environment variables

2 - adding custom environment variables

```go
package main

import (
	"log"

	"github.com/nulloop/goober"
)

func main() {
	peanut, err := goober.New(goober.OptSystemEnv())
	if err != nil {
		log.Fatal(err)
	}

	err = peanut.
		Yum("").
		Yum("./awesome.txt").
		Burp()

	if err != nil {
		log.Fatal(err)
	}

	peanut.Yum(`


		rm awesome.txt

	`).Burp()
}
```
