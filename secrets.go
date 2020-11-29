package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Creds struct {
	Creds []Cred `yaml:"creds"`
}

type Cred struct {
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Secret   string `yaml:"secret"`
	URL      string `yaml:"url"`
}

func (c *Creds) getCreds() *Creds {
	credsFile, err := ioutil.ReadFile("secrets.yaml")
	if err != nil {
		log.Printf("Unable to open creds file %v", err)
	}

	err = yaml.Unmarshal(credsFile, c)

	if err != nil {
		log.Printf("File parse err %v", err)
	}

	return c
}

func getCredbyName(cs Creds, n string) (Cred, error) {
	var c Cred
	fmt.Println("Getting cred by name " + n)
	for _, c = range cs.Creds {
		if c.Name == n {
			return c, nil
		}
	}

	return c, errors.New("Could not parse or load a valid cred")
}
