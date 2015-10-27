prettycat
=========

Used to pretty up some JSON output. It will take any valid JSON input - one by one - format it, and output it.

Usage
-----

```bash
$ prettycat --help
Usage of prettycat:
  -indent string
    	indentation to use (default "  ")
```

Simple Example
--------------

```bash
$ echo '{"everything":"awesome","team":{"everything":"cool"}}'|prettycat
{
  "everything": "awesome",
  "team": {
    "everything": "cool"
  }
}
```

Complex Example
---------------

Multiple JSON inputs can be fed in - they will be cleaned.

```
$ cat input.txt 
{"everything":"awesome","team":{"everything":"cool"}}
{"employees":[
{"firstName":"John", "lastName":"Doe"},
{"firstName":"Anna", "lastName":"Smith"},
{"firstName":"Peter", "lastName":"Jones"}
]}{
"glossary": {
"title": "example glossary",
"GlossDiv": {
"title": "S",
"GlossList": {
"GlossEntry": {
"ID": "SGML",
"SortAs": "SGML",
"GlossTerm": "Standard Generalized Markup Language",
"Acronym": "SGML",
"Abbrev": "ISO 8879:1986",
"GlossDef": {
"para": "A meta-markup language, used to create markup languages such as DocBook.",
"GlossSeeAlso": ["GML", "XML"]
},
"GlossSee": "markup"
}}}}}
```

```bash
⇒  cat input.txt|prettycat
{
  "everything": "awesome",
  "team": {
    "everything": "cool"
  }
}
{
  "employees": [
    {
      "firstName": "John",
      "lastName": "Doe"
    },
    {
      "firstName": "Anna",
      "lastName": "Smith"
    },
    {
      "firstName": "Peter",
      "lastName": "Jones"
    }
  ]
}
{
  "glossary": {
    "GlossDiv": {
      "GlossList": {
        "GlossEntry": {
          "Abbrev": "ISO 8879:1986",
          "Acronym": "SGML",
          "GlossDef": {
            "GlossSeeAlso": [
              "GML",
              "XML"
            ],
            "para": "A meta-markup language, used to create markup languages such as DocBook."
          },
          "GlossSee": "markup",
          "GlossTerm": "Standard Generalized Markup Language",
          "ID": "SGML",
          "SortAs": "SGML"
        }
      },
      "title": "S"
    },
    "title": "example glossary"
  }
}
```

