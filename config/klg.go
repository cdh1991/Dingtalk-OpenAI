package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Klb struct {
	Title  string `yaml:"title"`
	Prefix string `yaml:"prefix"`
	Suffix string `yaml:"suffix"`
}


var klgTmp *[]Klb

// LoadPrompt 加载Prompt
func LoadKlb() *[]Klb {
	data, err := ioutil.ReadFile("Klb.yml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(data, &klbTmp)
	if err != nil {
		log.Fatal(err)
	}
	return klbTmp
}
