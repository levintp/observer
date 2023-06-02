package main

import (
	"log"

	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/oj"

	"github.com/levintp/observer/internal/config"
)

func main() {
	configuration := config.Get("/etc/observer/observer.toml")
	log.Printf("Configuration: %+v\n", configuration)
	log.Println("---")

	var data interface{}
	content := []byte(`
	[
		{
			"slot": "1",
			"value": 300
		},
		{
			"slot": "2",
			"value": 176
		},
		{
			"slot": "3",
			"value": 426
		},
		{
			"slot": "4",
			"value": 32
		}
	]
	`)

	data, err := oj.Parse(content)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully parsed")

	expr, err := jp.ParseString("[?(@.value > 400)].slot")
	if err != nil {
		log.Fatal(err)
	}
	ys := expr.Get(data)
	for k, v := range ys {
		log.Println(k, "=>", v)
	}
}
