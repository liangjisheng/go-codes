package main

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

// Yaml struct of yaml
type Yaml struct {
	Mysql struct {
		User     string `yaml:"user"`
		Host     string `yaml:"host"`
		Password string `yaml:"password"`
		Port     string `yaml:"port"`
		Name     string `yaml:"name"`
	}
	Cache struct {
		Enable bool     `yaml:"enable"`
		List   []string `yaml:"list,flow"`
	}
}

// Yaml1 struct of yaml
type Yaml1 struct {
	SQLConf   Mysql `yaml:"mysql"`
	CacheConf Cache `yaml:"cache"`
}

// Yaml2 struct of yaml
type Yaml2 struct {
	Mysql `yaml:"mysql,inline"`
	Cache `yaml:"cache,inline"`
}

// Mysql struct of mysql conf
type Mysql struct {
	User     string `yaml:"user"`
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
}

// Cache struct of cache conf
type Cache struct {
	Enable bool     `yaml:"enable"`
	List   []string `yaml:"list,flow"`
}

func demoTest() {
	conf := new(Yaml)
	yamlFile, err := ioutil.ReadFile("test.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	log.Printf("yamlFile:\n%s\n", string(yamlFile))
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	log.Printf("conf:\n%v\n\n", conf)

	conf1 := new(Yaml1)
	err = yaml.Unmarshal(yamlFile, conf1)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	log.Printf("conf1:\n%v\n\n", conf1)

	conf2 := new(Yaml2)
	yamlFile, err = ioutil.ReadFile("test1.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	log.Printf("yamlFile:\n%s\n", string(yamlFile))
	err = yaml.Unmarshal(yamlFile, &conf2)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	log.Printf("conf2:\n%v\n\n", conf2)

	resultMap := make(map[string]interface{})
	err = yaml.Unmarshal(yamlFile, &resultMap)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	log.Printf("resultMap:\n%v\n", resultMap)
}
