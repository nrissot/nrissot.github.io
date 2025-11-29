window.onload = function() {
    // set the theme
    // check if we already have a value
    var theme = localStorage.getItem("theme");
    if (theme == null) {
        theme = "light dark"
    }
    setColorScheme(theme);

    // create the theme switcher
    var container = document.createElement("div");
    container.setAttribute("class", "themeswitcher");
    
    var pre = document.createElement("pre");

    pre.append("*thème-------------*\n| ");
    addCustomRadio("navigateur", "light dark", pre, (theme=="light dark"));
    pre.append("   |\n| ");
    addCustomRadio("thème clair", "light", pre, (theme=="light"));
    pre.append("  |\n| ");
    addCustomRadio("thème sombre", "dark", pre, (theme=="dark"));
    pre.append(" |\n*------------------*");

    container.append(pre);
    document.querySelector('body>pre').append(container);
}

function addCustomRadio(text, id, parent, checked) {
    var radio = document.createElement("input");
    var label = document.createElement("label");

    radio.setAttribute("type", "radio");
    radio.setAttribute("name", "themeswitcherinput");
    radio.setAttribute("id", id);
    radio.setAttribute("class", "custom-radio");
    if (checked) {
        radio.setAttribute("checked", "checked");
    }
    radio.setAttribute("onclick", "onThemeSwitcherInput(this);")
    label.setAttribute("for", id);
    label.append(text);
    parent.append(radio);
    parent.append(label);
}

function onThemeSwitcherInput(el) {
    if (el.id != "light dark" && el.id != "light" && el.id != "dark") {
        throw "Invalid theme id, what is happening ?";
        return;
    }
    setColorScheme(el.id);
    localStorage.setItem("theme", el.id);
}

function setColorScheme(scheme) {
    document.querySelector(':root').style.setProperty("--color-scheme", scheme);
}