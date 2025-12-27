package main

import (
	"fmt"
	wordwrap "nrissot/cms/word-wrap"
	"strings"
)

var txt_header string = "Nathan Rissot                                    Github  Blog  Contact\n::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::"

var txt_footer string = "Copyright (c) 2025 Nathan Rissot · curl this page:\n$ curl https://nrissot.github.io%s.txt\nexplore: $ curl https://nrissot.github.io/map.txt"

func (p *Page) GenerateTXT() {
	var txt string = txt_header + "\n"

	txt += "\n# " + p.Title + "\n"

	txt_details := p.Date.Format("02/01/2006") + " ·"
	for _, t := range p.Tags {
		txt_details += " [" + t + "]"
	}
	txt += wordwrap.Wrap(txt_details, 70)
	txt += "\n\n"

	// TODO : in-house wordwrap
	txt += wordwrap.Wrap(strings.TrimSpace(p.Content), 70)

	txt += "\n\n" + fmt.Sprintf(txt_footer, p.URL)
	p.TXT = txt
}

func (lp *ListPage) GenerateTXT() {
	var txt string = txt_header + "\n"

	txt += "\n# " + lp.Content + "\n"

	for _, article := range lp.Articles {
		txt_details := article.Date.Format("02/01/2006") + " ·"
		for _, t := range article.Tags {
			txt_details += " [" + t + "]"
		}
		txt += "\n-> " + article.Title + "\n" + wordwrap.WrapString(txt_details, 70) + "\n" + wordwrap.WrapString(article.Description, 70) + "\n\n$ curl https://nrissot.github.io" + article.URL + ".txt\n\n"
	}

	txt += "\n" + fmt.Sprintf(txt_footer, lp.URL)

	lp.TXT = txt
}
