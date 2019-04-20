package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"strings"
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
		Entries []entry
	}
)

var conf Config

func Conf(file string) Config {
	if conf == nil || len(conf.GetEntry()) == 0 {
		c := config{
			Entries: make([]entry, 0),
		}
		c.readConfig(file)
		conf = &c
	}
	return conf
}

func (c *config) readConfig(file string) {
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil && !strings.Contains(err.Error(), "no such file or directory") {
		log.Printf("yamlFile.Get err : %v ", err)
	} else if err != nil && strings.Contains(err.Error(), "no such file or directory") {
		ioutil.WriteFile(file, []byte(""), 0644)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Printf("Unmarshal: %v", err)
	}
}

func (c *config) GetEntry() []entry {
	return c.Entries
}
