package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	data, err := ParsePackage()

	if err != nil {
		log.Fatal(err)
	}

	value, ok := data["name"].(string)
	if !ok {
		log.Fatal("Package name is invalid or not exists")
	}

	domainName := "@" + value
	var differentDomain string
	var lambdaName string
	var hasApiGateway string
	var hasSqs string
	var hasEventBridge string

	fmt.Printf("Domain name [%s] ", domainName)
	fmt.Scanln(&differentDomain)
	if differentDomain != "" {
		domainName = differentDomain
	}
	fmt.Print("Lambda name ")
	fmt.Scanln(&lambdaName)
	fmt.Print("Gateway handler [Y/n] ")
	fmt.Scanln(&hasApiGateway)
	fmt.Print("SQS handler [Y/n] ")
	fmt.Scanln(&hasSqs)
	fmt.Print("EventBridge handler [Y/n] ")
	fmt.Scanln(&hasEventBridge)

	sName := strings.Split(lambdaName, "-")

	l := Lambda{
		HandlerName:    CreateHandlerName(sName),
		LambdaFileName: lambdaName,
		DomainName:     domainName,
		SpecName:       strings.Join(sName, " "),
		HasApiGateway:  ParseBooleanInput(hasApiGateway),
		HasSqs:         ParseBooleanInput(hasSqs),
		HasEventBridge: ParseBooleanInput(hasEventBridge),
	}

	if err = l.initLambda(); err != nil {
		log.Fatal(err)
	}
}

func ParsePackage() (map[string]interface{}, error) {
	packageFile, err := os.ReadFile("package.json")
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	if err := json.Unmarshal(packageFile, &data); err != nil {
		log.Fatal(err)
	}

	return data, err
}

func ParseBooleanInput(input string) bool {
	if input == "Y" || input == "y" || input == "" {
		return true
	}

	return false
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
