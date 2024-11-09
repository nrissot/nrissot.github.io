import re

input = """
<pre><code>
elif line.startswith(" "):
    event["DESCRIPTION"] += line.removeprefix(" ").removesuffix("\n")
else:
    for prefix in ("DTSTART:", "DTEND:", "SUMMARY:", "LOCATION:", "DESCRIPTION:", "UID:") :
        if line.startswith(prefix):
            event[prefix.removesuffix(":")] = line.removeprefix(prefix).removesuffix("\n")
            break
</code></pre>
"""

reserved_words = ["elif ", "if ", "else", "for ", " in ", " as ", "def ", "class ", "break", "continue", "case", "self", "try", "except ", "not ", "True", "False", "with ", "return", "global ", "None"]
symbols = ["(",")","[","]","{","}"]

def spanify_string(match : re.Match):
    return "<span class=\"code-string\">" + match.string[match.start():match.end()] + "</span>"
def spanify_comment(match : re.Match):
    return "" + match.string[match.start():match.end()] + "</span>"
def spanify_word(word : str):
    return "<span class=\"code-function\">" + word + "</span>"
def spanify_symbol(symbol : str):
    return "<span class=\"code-symbol\">" + symbol + "</span>"


def main():
    replaced = re.sub(r"#.*(\n|$)", spanify_comment, input)
    replaced = re.sub(r"\".*\"", spanify_string, replaced)
    for word in reserved_words :
        replaced = replaced.replace(word, spanify_word(word))
    for symbol in symbols :
        replaced = replaced.replace(symbol, spanify_symbol(symbol))
    with open("codified.txt", "w", encoding="utf-8") as f:
        print(replaced, file=f)

if __name__ == "__main__" :
    main()