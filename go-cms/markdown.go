package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

type MarkdownFile struct {
	Path        string
	Content     string
	Frontmatter FrontMatter
}

type FrontMatter struct {
	Title       string    `yaml:"title"`
	Description string    `yaml:"desc"`
	Date        time.Time `yaml:"date"`
	Tags        []string  `yaml:"tags"`
}

func ReadFrontMatter(data []byte) (*FrontMatter, error) {
	r := regexp.MustCompile("^---(.|\n)*---")
	frontmatterdata := r.Find(data)
	if frontmatterdata == nil {
		return nil, fmt.Errorf("no frontmatter")
	}
	frontmatterstring := strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(string(frontmatterdata), "---"), "---"))
	var fmtr FrontMatter
	err := yaml.Unmarshal([]byte(frontmatterstring), &fmtr)
	if err != nil {
		return nil, err
	}
	for i := range fmtr.Tags {
		fmtr.Tags[i] = strings.ToLower(strings.ReplaceAll(strings.TrimSpace(fmtr.Tags[i]), " ", "-"))
	}
	return &fmtr, nil
}

func FindMarkdownFiles(path string) []MarkdownFile {
	var matches []string
	err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(path, ".md") {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	if matches == nil {
		return nil
	}

	var out []MarkdownFile
	for _, match := range matches {
		file, err := os.Open(match)
		if err != nil {
			fmt.Printf("error while opening %s\n", match)
			continue
		}
		defer file.Close()

		data, err := io.ReadAll(file)
		if err != nil {
			fmt.Printf("error while reading %s\n", match)
			continue
		}
		fmtr, err := ReadFrontMatter(data)
		if err != nil {
			fmt.Printf("error while extracting frontmatter %s\n", match)
			continue
		}
		r := regexp.MustCompile("^---(.|\n)*---")
		loc := r.FindIndex(data)
		out = append(out, MarkdownFile{match, string(data[:loc[0]]) + string(data[loc[1]:]), *fmtr})
	}

	if len(out) == 0 {
		return nil
	}
	return out
}

func (md MarkdownFile) AsPage(target string) Page {
	reg := regexp.MustCompile("[^a-zA-Z0-9]")
	url := target + strings.ToLower(reg.ReplaceAllString(md.Frontmatter.Title, ""))
	return CreatePage(md.Frontmatter.Title, url, md.Content, md.Frontmatter.Description, md.Frontmatter.Date, md.Frontmatter.Tags)
}
