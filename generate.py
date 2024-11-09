from data.text import langs
from data.articles import articles_list
from pathlib import Path
import time

# TEMPLATES_DIR = "templates"
# templates = Path(TEMPLATES_DIR).glob("*.html")

def main():
    script = ""
    with open("templates\\script.html", "r", encoding="utf-8") as script_source:
        script = script_source.read()
    for lang in langs:
        with open(Path("templates\\index_template.html"), "r", encoding="utf-8") as source:
            content = source.read()
        for keyword in lang[1].keys():
            content = content.replace(f"%{keyword}%", lang[1][keyword])
        content = content.replace("%Script%", script)
        content += f"<!-- last generated at: {time.strftime("%Y/%m/d %H:%M:%S")} -->"
        with open(f"{lang[0]}/index.html", "w", encoding="utf-8") as target:
            print(content,file=target)
            print(f"generated {lang[0]}/index.html at: {time.strftime("%Y/%m/d %H:%M:%S")}")

    blog_article_list_string = """"""
    for article in articles_list:
        with open("templates\\blog_article_template.html", "r", encoding="utf-8") as article_template:
            article_page_content = article_template.read()

        for keyword in article.keys():
            article_page_content = article_page_content.replace(f"%{keyword}%", article[keyword])
        
        article_page_content = article_page_content.replace("%Script%", script)    
        article_page_content += f"<!-- last generated at: {time.strftime("%Y/%m/d %H:%M:%S")} -->"

        with open(article["Article-Source"], "r", encoding="utf-8") as content_source :
            article_page_content = article_page_content.replace("%Article-Content%", content_source.read())

        blog_article_list_string += (
            f"""<article>
                <header>
                    <h3 class="no-hash"><a href="blog/{article["Article-URL"]}.html">{article["Article-Title"]}</a></h3>
                    <span>{article["Article-Date"]}, {article["Article-Lang-Version"]}, {article["Article-Reading-Time"]}</span>
                </header>
                <p class="details">{article["Article-Details"]}</p>
                <p>
                    {article["Article-Summary"]}
                </p>
            </article>\n"""
        )

        with open(f"blog\\{article["Article-URL"]}.html", "w", encoding="utf-8") as target :
            print(article_page_content,file=target)

        print(f"generated blog/{article["Article-URL"]}.html at: {time.strftime("%Y/%m/d %H:%M:%S")}")
    with open("templates\\blog_template.html", "r", encoding="utf-8") as blog_template :
        blog_content = blog_template.read()

    blog_content = blog_content.replace("%Content%", blog_article_list_string)
    blog_content = blog_content.replace("%Script%", script)

    with open("blog.html", "w", encoding="utf-8") as target :
        print(blog_content, file=target)
    print(f"generated blog.html at: {time.strftime("%Y/%m/d %H:%M:%S")}")

if __name__ == "__main__":
    main()