package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	l := Lambda{
		HandlerName:    "helloWorld",
		LambdaFileName: "hello-world",
		DomainName:     "@hello-world",
		SpecName:       "hello world",
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
