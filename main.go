package main

import (
	"bytes"
	"github.com/docker/docker/builder/dockerfile/command"
	"github.com/docker/docker/builder/dockerfile/parser"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var debug = false

// DockerfileFromPath reads a Dockerfiler from a oath
func DockerfileFromPath(input string) (*Dockerfile, error) {
	payload, err := ioutil.ReadFile(input)
	if err != nil {
		return nil, err
	}
	if debug {
		log.Println(string(payload))
	}
	return DockerfileRead(bytes.NewReader(payload))
}

// DockerfileRead reads a Dockerfile as io.Reader
func DockerfileRead(input io.Reader) (*Dockerfile, error) {
	dockerfile := Dockerfile{}

	root, err := parser.Parse(input)
	if err != nil {
		return nil, err
	}
	dockerfile.root = root

	return &dockerfile, nil
}

// GetFrom returns the current FROM
func (d *Dockerfile) From() string {
	for _, node := range d.root.Children {
		if node.Value == command.From {
			from := strings.Split(node.Original, " ")[1]
			if debug {
				log.Println(from)
			}
			return from
		}
	}

	return ""
}

// String returns a docker-readable Dockerfile
func (d *Dockerfile) String() string {
	lines := []string{}
	for _, child := range d.root.Children {
		lines = append(lines, child.Original)
	}
	return strings.Join(lines, "\n")
}

// load rules file
func loadRules(ruleFile string) (Rules, error) {
	rulesdata, err := ioutil.ReadFile(ruleFile)

	if err != nil {
		log.Panic("failed to read rules file")
	}
	var rules Rules

	err = yaml.Unmarshal(rulesdata, &rules)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	if debug {
		log.Printf("--- rules:\n%v\n\n", rules)
	}

	return rules, err
}

func validFrom(rules Rules, dfile *Dockerfile) bool {
	for _, entry := range rules.From {
		if entry == dfile.From() {
			return true
		}
	}
	return false
}

func main() {
	debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	dockerfile := os.Getenv("DOCKERFILE")
	ruleFile := os.Getenv("RULESFILE")
	rules, _ := loadRules(ruleFile)
	dfile, err := DockerfileFromPath(dockerfile)
	if err != nil {
		log.Panic(err)
	}
	from := dfile.From()
	log.Println(from)
	if !validFrom(rules, dfile) {
		log.Panic("FROM not valid")
	}

	os.Exit(0)
}
