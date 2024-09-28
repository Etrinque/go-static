# Go-Static

A static site generator written in GO.

## Pupose

Go-Static is a simple, general purpose Markdown to HTML conversion tool written in GO and utilizing [Templ](https://templ.guide/).

The creation of Go-Static is strictly for my education and usage in my own personal website project.  Primarily for blog content. Markdown has a nice, simple and concise syntax and allows for great creativity in writing.  It is a great fit for simple pages, blogs, news articles, general information pages, etc.

There are many static site generators which have greater functionality, and if you need to adhear to certain specifications I would advise checking out other projects for now.

Projects like:
    - [Makdown-it](https://github.com/markdown-it)
    - [Goldmark](https://github.com/yuin/goldmark/)
    - [Golang-Commonmark](https://gitlab.com/golang-commonmark/markdown)
    - etc...

## TechStack

- GO 1.23+
- [TEMPL](https://templ.guide/)

## Mechanism

 Go-Static is a simple lexical anaylsis and parser that creates an AST[^1] and renders HTML out.

1. The lexer steps through the file, creates, and organizes Tokens into Tree Nodes.

2. The Tree Nodes are then stepped through and used to populate the HTML template backed by TEMPL.

3. The template is then rendered into output of HTML code file that follows site guidelines for style and functionality.

4. Finally the file is distributed to the website directory to be used immediately in hosting.

[^1]: [Abstract Syntax Tree - Wiki](https://en.wikipedia.org/wiki/Abstract_syntax_tree)
