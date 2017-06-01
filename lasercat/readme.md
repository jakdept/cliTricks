lasercat
========

STOP USING THIS START USING JQ

[DOWNLOAD HERE](https://stedolan.github.io/jq/)

[LASERCATS](http://www.nbc.com/sites/nbcunbc/files/files/images/2014/9/06/140228_2750569_SNL_Digital_Short__Laser_Cats_anvver_1.jpg)

Ok I'm done now.

`lasercat` pulls a specific target from JSON input.

Usage
-----

```bash
$ lasercat --help
Usage of lasercat:
  -target value
    	locations to pluck from input (default [])
```

Simple Example
--------------

```bash
$ echo '{"everything":"awesome","team":{"everything":"cool"}}'|go run lasercat.go -target "everything" -target "team","everything"
awesome cool
```

Complex Example
---------------

```bash
{
  "menu": {
    "id": "file",
    "value": "File",
    "popup": {
      "menuitem": [
        {
          "value": "New",
          "onclick": "CreateNewDoc()"
        },
        {
          "value": "Open",
          "onclick": "OpenDoc()"
        },
        {
          "value": "Close",
          "onclick": "CloseDoc()"
        }
      ]
    }
  }
}
```

```bash
$  cat exampleData|lasercat -target "menu","id" -target "menu","popup","menuitem",1,"onclick"
file OpenDoc()
```

Invalid Position
----------------

If you give a target that does not exist within given JSON input, you will get an error.

```bash
⇒  cat exampleData|lasercat -target "menu","id" -target "menu","popup","menuitem",1,"oncllick"
2015/10/27 05:48:29 non-existant map position
```
