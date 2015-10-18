package main

import (
	"bufio"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

func TemplateBuild(in bufio.Reader, out io.Writer, t template.Template) (err error) {
	var parts []string
	line, notDone, err := in.ReadLine()
	if err != nil {
		return err
	}
	for notDone {
		parts = strings.Fields(string(line))
		err = t.Execute(out, parts)
		if err != nil {
			return
		}
		line, notDone, err = in.ReadLine()
	}
	return
}

func main() {
	templateStringPointer := flag.String("template", "", "raw template to build with")
	templateFile := flag.String("templateFile", "", "file containing template to use")

	flag.Parse()
	templateString := *templateStringPointer

	if templateString == "" {
		templateBytes, err := ioutil.ReadFile(*templateFile)
		if err != nil {
			log.Fatal("bad file to read template from")
		}
		templateString = string(templateBytes)
	}
	if templateString == "" {
		log.Fatal("no template provided")
	}

	t := template.New("t")
	template, err := t.Parse(templateString)
	if err != nil {
		log.Fatal(err)
	}
	input := bufio.NewReader(os.Stdin)
	log.Fatal(TemplateBuild(*input, os.Stdout, *template))
}
