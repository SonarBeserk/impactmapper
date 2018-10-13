package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/BurntSushi/toml"
)

var (
	tmpl        = string(MustAsset("files/dot.tmpl"))
	defaultToml = MustAsset("files/map.toml")

	templateName = "map"
)

// Goal is an outcome
type Goal struct {
	Name string
}

// Actor is who accomplishes the goal
type Actor struct {
	Name string
	Goal int
}

// Impact is the what changes the actor could make
type Impact struct {
	Name  string
	Actor int
}

// Deliverable is something measurable as progress towards the goal
type Deliverable struct {
	Name   string
	Impact int
}

// Data is used to write the template
type Data struct {
	Name         string
	Mode         string
	Goals        map[string]Goal
	Actors       map[string]Actor
	Impacts      map[string]Impact
	Deliverables map[string]Deliverable
}

func main() {
	if _, err := os.Stat(templateName + ".toml"); os.IsNotExist(err) {
		ioutil.WriteFile(templateName+".default.toml", defaultToml, 0644)
		log.Fatalln("Failed to find map.toml file, created example file")
	}

	var data Data
	if _, err := toml.DecodeFile(templateName+".toml", &data); err != nil {
		log.Fatalln("Failed to decode toml file")
	}

	tmpl, err := template.New("dot").Parse(tmpl)
	if err != nil {
		log.Fatalln("Failed to parse template file")
	}

	f, err := os.Create("impact.dot")
	if err != nil {
		log.Fatalln("Failed to create dot file")
	}

	defer f.Close()

	w := bufio.NewWriter(f)

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Fatalln("Failed to run template")
	}

	err = w.Flush()
	if err != nil {
		log.Fatalln("Failed to write dot file")
	}
}
