package main

import (
	"flag"
	"log"
	"os"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	var name string
	var domain string

	ParseFlags(&name, &domain)

	sName := strings.Split(name, "-")
	handlerName := CreateHandlerName(sName)

	l := Lambda{
		HandlerName:    handlerName,
		LambdaFileName: name,
		DomainName:     domain,
		SpecName:       strings.Join(sName, " "),
	}

	if err := l.CreateFile(PackageFile, "", "package.json"); err != nil {
		log.Fatal(err)
	}

	if err := l.CreateSrcFile(IndexFile, "index.ts"); err != nil {
		log.Fatal(err)
	}

	if err := l.CreateSrcFile(HandlerFile, l.LambdaFileName+".ts"); err != nil {
		log.Fatal(err)
	}

	if err := l.CreateSrcFile(SpecFile, l.LambdaFileName+".spec.ts"); err != nil {
		log.Fatal(err)
	}
}

func CreateHandlerName(sName []string) string {
	caser := cases.Title(language.AmericanEnglish)
	firstName := sName[0]

	names := []string{}
	for _, v := range sName[1:] {
		names = append(names, caser.String(v))
	}

	return firstName + strings.Join(names, "")
}

func ParseFlags(name *string, domain *string) {
	flag.StringVar(name, "name", "", "Lambda name")
	flag.StringVar(domain, "domain", "", "Lambda domain name")
	flag.Parse()

	if *name == "" {
		log.Fatal("Lambda name is required")
	}

	if *domain == "" {
		log.Fatal("Lambda domain name is required")
	}
}

type Lambda struct {
	HandlerName    string
	LambdaFileName string
	DomainName     string
	SpecName       string
}

func (l *Lambda) CreateFile(t, path, fileName string) error {
	lambdaPath := "packages/" + l.LambdaFileName + "/" + path
	tml, err := template.New(fileName).Parse(t)
	if err != nil {
		log.Printf("Error parsing %s file\n", fileName)
		return err
	}

	os.MkdirAll(lambdaPath, os.ModePerm)

	f, err := os.OpenFile(
		lambdaPath+fileName,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		log.Printf("Error opening %s file\n", fileName)
		return err
	}
	defer f.Close()

	return tml.Execute(f, l)
}

func (l *Lambda) CreateSrcFile(t, fileName string) error {
	return l.CreateFile(t, "/src/", fileName)
}
