package main

import (
	"log"
	"os"
	"text/template"
)

type Lambda struct {
	HandlerName    string
	LambdaFileName string
	DomainName     string
	SpecName       string
	HasApiGateway  bool
	HasSqs         bool
	HasEventBridge bool
}

func (l *Lambda) initLambda() error {
	if err := l.CreateFile(PackageFile, "", "package.json"); err != nil {
		return err
	}

	if err := l.CreateSrcFile(IndexFile, "index.ts"); err != nil {
		return err
	}

	if err := l.CreateSrcFile(HandlerFile, l.LambdaFileName+".ts"); err != nil {
		return err
	}

	if err := l.CreateSrcFile(SpecFile, l.LambdaFileName+".spec.ts"); err != nil {
		return err
	}

	return nil
}

func (l *Lambda) CreateSrcFile(t, fileName string) error {
	return l.CreateFile(t, "/src/", fileName)
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
