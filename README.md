# Goober v2

Tired of none-sense Makefile syntax and I want to build and configure my application using golang not making
complex Makefile and bash scripts. This project is created out of THE frusration of dealing with bash and
makefile syntax.

`goober` has only one method `Yum`. Yum has similar signature as `Printf`. Just write the bash script and let it run.

```go
func buildScript() {
	err := goober.Yum(`
		go build -o %s ./cmd/awesome/main.go
	`, "./bin/awesome")

	if err != nil {
		panic(err)
	}
}
```
