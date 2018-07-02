package application

import (
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// Service ...
type Service struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Endpoint    string `yaml:"endpoint"`
}

// RemoteServices ...
type RemoteServices struct {
	Services []Service `yaml:"services"`
}

// GetApps ...
func (c *RemoteServices) GetApps(filePATH string) *RemoteServices {
	yamlFile, err := ioutil.ReadFile(filePATH)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	ymlContent := []byte(os.ExpandEnv(string(yamlFile)))
	err = yaml.Unmarshal(ymlContent, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}
