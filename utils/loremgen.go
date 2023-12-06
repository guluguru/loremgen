package utils

import (
	"lorem-generator/config"
	"math/rand"
)

var LANG string

type paragraph struct {
	content   string
	sentences []sentence
}
type sentence struct {
	content  string
	snippets []snippet
}
type snippet struct {
	content string
}

/**
 * Lorem ipsum generate
 */
func Generate() string {
	p := paragraph{}
	p.generate()
	return p.content
}

func (o *paragraph) generate() (ok bool) {
	/** sentences count */
	var sc int
	sc = rand.Intn(4) + 3

	for i := 0; i < sc; i++ {
		o.sentences = append(o.sentences, sentence{})
		o.sentences[i].generate()
		if i != sc-1 {
			o.content += o.sentences[i].content
		} else {

			o.content += o.sentences[i].content + "\n\n"
		}
	}
	return true
}

/**
 * After:
 * sentence.content => "alias accusantium consequatur aut, perferendis sit voluptatem accusantium consequatur."
 */
func (o *sentence) generate() (ok bool) {
	/** snippets count */
	var sc int
	var sep string
	var tail string

	if LANG == "en" {
		sc = rand.Intn(7) + 2
		sep = ", "
		tail = ". "
	} else if LANG == "cn" {
		sc = rand.Intn(7) + 2
		sep = "，"
		tail = "。"
	}

	for i := 0; i < sc; i++ {
		o.snippets = append(o.snippets, snippet{})
		o.snippets[i].generate()
		if i != sc-1 {
			o.content += o.snippets[i].content + sep
		} else {
			o.content += o.snippets[i].content + tail
		}
	}
	return true
}

/**
 * After:
 * snippet.content => "nisi in sed in esse aliqua consectetur"
 */
func (o *snippet) generate() (ok bool) {
	/** words count */
	var wc int
	var words []string
	var sep string

	if LANG == "en" {
		wc = rand.Intn(7) + 6
		words = config.En
		sep = " "
	} else if LANG == "cn" {
		wc = rand.Intn(4) + 3
		words = config.Cn
	}

	for i := 0; i < wc; i++ {
		if i != wc-1 {
			o.content += words[rand.Intn(len(words))] + sep
		} else {
			o.content += words[rand.Intn(len(words))]
		}
	}

	return true
}
