package main

import (
	"fmt"
	"html/template"
	"os"
	"time"
)

var deployedBy = os.Getenv("DEPLOYED_BY")
var commit = os.Getenv("COMMIT")
var branch = os.Getenv("BRANCH")

func init() {
	if (deployedBy == "") || (commit == "") || (branch == "") {
		fmt.Println("DEPLOYED_BY, COMMIT and ENV variables are required and their value can't be empty")
		os.Exit(1)
	}
}

type HTML struct {
	DeployedBy string
	CommitID   string
	Branch     string
	Date       string
}

func generateHTML(inFile string, outFile string) error {

	tmpl, err := template.ParseFiles(inFile)
	if err != nil {
		fmt.Sprintf("There reading the html template: %v", err)
		return err
	}

	file, err := os.Create(outFile)

	defer file.Close()
	if err != nil {
		fmt.Sprintf("Error opening output file: %v", err)
		return err
	}

	tplErr := tmpl.Execute(file, HTML{
		DeployedBy: deployedBy,
		CommitID:   commit,
		Branch:     branch,
		Date:       time.Now().Format("2006-01-02 15:04:05"),
	})
	if tplErr != nil {
		return tplErr
	}

	return nil
}

func main() {

	generateHTML("templates/index.html.tpl", "index.html")
}
