package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type (
	Config interface {
		GetEntry() []entry
	}
	entry struct {
		Name    string
		Email   string
		Pattern string
	}
	config struct {
		file    string
		Entries []entry
	}
)

var conf Config

func Conf(file string) Config {
	if conf == nil || len(conf.GetEntry()) == 0 {
		c := config{
			file:    file,
			Entries: make([]entry, 0),
		}
		c.readConfig(file)
		conf = &c
	}
	return conf
}

func (c *config) readConfig(file string) {
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("yamlFile.Get err : %v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Printf("Unmarshal: %v", err)
	}
}

func (c *config) GetEntry() []entry {
	return c.Entries
}
