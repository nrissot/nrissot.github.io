from data.text import english, french, langs
from pathlib import Path
import time

TEMPLATES_DIR = "templates"
templates = Path(TEMPLATES_DIR).glob("*.html")

def main():
    for template in templates :
        for lang in langs:
            with open(template, "r", encoding="utf-8") as source:
                content = source.read()
                for keyword in lang[1].keys():
                    content = content.replace(f"%{keyword}%", lang[1][keyword])
                content += f"<!-- last generated at: {time.strftime("%Y/%m/d %H:%M:%S")} -->"
                with open(f"{lang[0]}/{template.name.replace("_template", "")}", "w", encoding="utf-8") as target:
                    print(f"generated {lang[0]}/{template.name.replace("_template", "")} at: {time.strftime("%Y/%m/d %H:%M:%S")}")
                    print(content,file=target)

if __name__ == "__main__":
    main()