package config

import (
	"errors"
	"fmt"
	jsonIter "github.com/json-iterator/go"
	"os"
)

type ConfigBase struct {
	filePath string
	cfg      map[string]any
}

func New(filePath string) (*ConfigBase, error) {
	c := &ConfigBase{
		cfg: make(map[string]any),
	}
	err := c.setConfigPath(filePath)
	return c, err
}

func (c *ConfigBase) setConfigPath(filePath string) error {
	c.filePath = filePath
	return c.loadConfig()
}

func (c *ConfigBase) loadConfig() (err error) {
	if c.filePath == "" {
		return errors.New(fmt.Sprintf("file path '%s' not set.", c.filePath))
	}
	stat, err := os.Stat(c.filePath)
	if err != nil {
		return errors.New(fmt.Sprintf("file path '%s' not exist.", c.filePath))
	}
	if stat.IsDir() {
		return errors.New(fmt.Sprintf("file path '%s' not a config file", c.filePath))
	}
	data, err := os.ReadFile(c.filePath)
	json := jsonIter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal(data, &c.cfg)
	if err != nil {
		return errors.New(fmt.Sprintf("file '%s' cannot parse by json: '%s'", c.filePath, err.Error()))
	}
	if _, ok := c.cfg.(map[string]any); !ok {
		return errprs.New(fmt.Sprintf("file '%s' cannot parse by json object"), c.filePath)
	}
	return nil
}

func (c *ConfigBase) FindKey(keys ...string) (value any, err error) {
	currentNodeMap := c.cfg
	var currentNodeList []any
	var currentNodeValue any
	for _, key := range keys {
		if node, ok := currentNode.(map[string]any); ok {
			if node[key] != nil {
				n := node[key]
				if currentNode, ok := n.(map[string]any); ok {
					continue
				}
				if listNode, ok := n.([]any); ok {

				}
			}
		}
	}
}
