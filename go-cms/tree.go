package main

// │ ├ └ ─ ╴

// .
// ├─╴blog
// │  ├─╴wikipedia.txt
// │  └─╴wikipedia2.txt
// ├─╴tags
// │  ├─╴misc.txt
// │  └─╴encyclopedia.txt
// ├─╴index.txt
// ├─╴blog.txt
// └─╴tags.txt

func DrawTree(root, blog, tags []string) string {
	// Assumption: the root, blog and tags slices are sorted properly (LWFS)
	var tree_txt string = ".\n├─╴blog\n"
	for i, entry := range blog {
		if i != (len(blog) - 1) {
			tree_txt += "│  ├─╴" + entry + "\n"
		} else {
			tree_txt += "│  └─╴" + entry + "\n"
		}
	}
	tree_txt += "├─╴tags\n"
	for i, entry := range tags {
		if i != (len(tags) - 1) {
			tree_txt += "│  ├─╴" + entry + "\n"
		} else {
			tree_txt += "│  └─╴" + entry + "\n"
		}
	}
	for i, entry := range root {
		if i != (len(root) - 1) {
			tree_txt += "├─╴" + entry + "\n"
		} else {
			tree_txt += "└─╴" + entry + "\n"
		}
	}
	return tree_txt
}
