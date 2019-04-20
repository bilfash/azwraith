package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

type (
	Config interface {
		GetEntry() []entry
		RegisterEntry(name string, mail string, pattern string)
		DeleteEntry(index int)
	}
	entry struct {
		Name    string
		Email   string
		Pattern string
	}
	config struct {
		filename string
		Entries  []entry
	}
)

func Conf(file string) Config {
	c := config{
		filename: file,
		Entries:  make([]entry, 0),
	}
	c.readConfig()
	return &c
}

func (c *config) readConfig() {
	yamlFile, err := ioutil.ReadFile(c.filename)
	if err != nil && !strings.Contains(err.Error(), "no such file or directory") {
		fmt.Printf("yamlFile.Get err : %v ", err)
	} else if err != nil && strings.Contains(err.Error(), "no such file or directory") {
		ioutil.WriteFile(c.filename, []byte(""), 0644)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Printf("Unmarshal: %v", err)
	}
}

func (c *config) saveToConfig() {
	d, err := yaml.Marshal(&c)
	if err != nil {
		fmt.Printf("error saving to config: %v", err)
	}
	err = ioutil.WriteFile(c.filename, d, 0644)
	if err != nil {
		fmt.Printf("error saving to config: %v", err)
	}
}

func (c *config) GetEntry() []entry {
	return c.Entries
}

func (c *config) RegisterEntry(name string, mail string, pattern string) {
	ent := entry{
		Name:    name,
		Email:   mail,
		Pattern: pattern,
	}
	c.Entries = append(c.Entries, ent)
	c.saveToConfig()
}

func (c *config) DeleteEntry(index int) {
	if index < len(c.Entries) {
		c.Entries = append(c.Entries[:index], c.Entries[index+1:]...)
		c.saveToConfig()
	}
}
