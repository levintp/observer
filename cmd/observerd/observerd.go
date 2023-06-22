package main

import (
	"log"

	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/oj"

	"github.com/levintp/observer/internal/config"
)

func main() {
	configuration := config.Get()
	log.Printf("Configuration:\n\n%v\n", configuration)
	log.Println("---")

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
	content = []byte(`{"value": 4}`)

	data, err := oj.Parse(content)
	if err != nil {
		log.Fatal(err)
	}

	//expr, err := jp.ParseString("[?(@.value > 400)].slot")
	expr, err := jp.ParseString("[?(@ > 3)]")
	if err != nil {
		log.Fatal(err)
	}
	ys := expr.Get(data)
	for k, v := range ys {
		log.Println(k, "=>", v)
	}
}
