package main

import (
	"flag"
	"fmt"
	"lorem-generator/utils"
)

func main() {
	_paragraphs_count := flag.Int("p", 1, "number of paragraphs to generate")
	_lang := flag.String("t", "en", "language type, en or cn")
	flag.Parse()
	utils.LANG = *_lang
	cnt := *_paragraphs_count

	var s string
	for i := 0; i < cnt; i++ {
		s += utils.Generate()
	}
	fmt.Println(s)
}
