package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type ConfigBase struct{
	filePath string
	cfg map[]any
}

func New(filePath string) (*ConfigBase, err){
	c := &ConfigBase{
		cfg: make(map[]any),
	}
	err = c.setConfigPath(filePath)
	return c, err
}

func (c *ConfigBase) setConfigPath(filePath string){
	c.filePath = filePath
	c.loadConfig()
}

func (c *ConfigBase) loadConfig() err{
	if c.filePath == ""{
		return errors.New(fmt.Sprintf("file path '%s' not set.", c.filePath))
	}
	stat, err := os.Stat(c.filePath)
	if err != nil {
		return errors.New(fmt.Sprintf("file path '%s' not exist.", c.filePath))
	}
	if stat.IsDir() {
		return errors.New(fmt.Sprintf("file path '%s' not a config file.", c.filePath))
	}
}
