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

func TemplateBuild(in io.Reader, out io.Writer, t *template.Template) (err error) {
	r := bufio.NewReader(in)
	var parts []string
	line, err := r.ReadString('\n')
	parts = strings.Fields(string(line))
	// log.Print(line)
	// log.Print(err)
	for err == nil {
		// log.Print(parts)
		if len(parts) > 0 {
			err = t.Execute(out, parts)
			if err != nil {
				return
			}
			out.Write([]byte("\n"))
		}
		line, err = r.ReadString('\n')
		parts = strings.Fields(string(line))
		// log.Print(line)
	}
	if err == io.EOF {
		return nil
	}
	return err
}

func main() {
	templateStringPointer := flag.String("template", "", "raw template to build with")
	templateFile := flag.String("templateFile", "", "file containing template to use")

	flag.Parse()
	templateString := *templateStringPointer

	if templateString == "" {
		if *templateFile == "" {
			log.Fatal("no template provided")
		}
		templateBytes, err := ioutil.ReadFile(*templateFile)
		if err != nil {
			log.Fatal("bad file to read template from")
		}
		templateString = string(templateBytes)
	}

	if !strings.HasSuffix(templateString, "\n") {
		templateString = templateString + "\n"
	}

	t := template.New("t")
	template, err := t.Parse(templateString)
	if err != nil {
		log.Fatal(err)
	}
	// input := bufio.NewReader(os.Stdin)
	// output := bufio.NewWriter(os.Stdout)
	if err := TemplateBuild(os.Stdin, os.Stdout, template); err != nil {
		log.Print(err)
	}
}
