templcat
========

Used to build text output from field input using a given template.

Templating engine is [Go's template engine](https://golang.org/pkg/text/template).

Usage
-----

```bash
$  templcat --help
Usage of templcat:
  -template string
    	raw template to build with
  -templateFile string
    	file containing template to use
```

Simple Example
--------------

```bash
$ cat names
bob top
mary top
ray middle
```

```bash
$ cat names|templcat --template '{"employee":{"name":{{index . 0}}, "level":{{index . 1}}}}'
{"employee":{"name":bob, "level":top}}
{"employee":{"name":mary, "level":top}}
{"employee":{"name":ray, "level":middle}}
```

File Based Example
------------------

```bash
$ cat fruits
apple banana cherry date eggplant
fig
grape honeydew
jack@jack-mobile:~/working/golang/src/github.com/JackKnifed/cliTricks/templcat|master⚡
$ cat fruit_template
{
  "fruits":[{{range $index, $element := .}}
	{{if $index}},{{end}}"{{.}}"{{end}}
  ]
}
```

```bash
$ cat fruits|templcat --templateFile fruit_template
{
  "fruits":[
	"apple"
	,"banana"
	,"cherry"
	,"date"
	,"eggplant"
  ]
}

{
  "fruits":[
	"fig"
  ]
}

{
  "fruits":[
	"grape"
	,"honeydew"
  ]
}
```