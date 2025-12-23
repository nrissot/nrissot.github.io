package main

import "fmt"

// TODO : make it a pretty box
func GenerateArticleCardHTML(p Page) string {
	pills := ""
	for _, t := range p.Tags {
		if pills != "" {
			pills += " "
		}
		pills += GenerateTagPillHTML(t)
	}
	return fmt.Sprintf("<article><span>-> </span><a href=\"%s\">%s</a>\n<span> | </span><i>%s</i>\n<span> | </span><nav>%s</nav></article>", p.URL, p.Title, p.Description, pills)
}

func GenerateHeadHTML() string {
	return "<!DOCTYPE html>\n<html lang=\"fr\">\n<head>\n\t<meta charset=\"UTF-8\">\n\t<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n\t<title>Nathan Rissot [HOME]</title>\n\t<link rel=\"stylesheet\" href=\"/static/style.css\">\n\t<link rel=\"icon\" type=\"image/x-icon\" href=\"/static/assets/images/favicon.ico\">\n</head>\n<script src=\"/script.js\"></script>"
}

func GenerateHeaderHTML() string {
	return "<header><h1><a href=\"/index.html\">Nathan Rissot</a></h1>                                    <nav><a href=\"https://github.com/nrissot\">Github</a>  <a href=\"/blog.html\">Blog</a>  <a href=\"/index.html#contact\">Contact</a></nav>\n<span aria-hidden=\"true\" class=\"gray\">::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::</span></header>\n"
}

// TODO : add proper curl
func GenerateFooterHTML() string {
	return "<footer>Copyright (c) 2025 Nathan Rissot</footer>"
}

func GenerateTagPillHTML(tag string) string {
	return "<a href=\"/tags/" + tag + ".html\" class=\"tag\">[" + tag + "]</a>"
}
