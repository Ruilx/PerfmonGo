package configBase

import (
	"errors"
	"fmt"
	jsonIter "github.com/json-iterator/go"
	"os"
	"strconv"
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
		return errors.New(fmt.Sprintf("file path '%s' not a configBase file", c.filePath))
	}
	data, err := os.ReadFile(c.filePath)
	json := jsonIter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal(data, &c.cfg)
	if err != nil {
		return errors.New(fmt.Sprintf("file '%s' cannot parse by json: '%s'", c.filePath, err.Error()))
	}
	//if _, ok := c.cfg.(map[string]any); !ok {
	//	return errors.New(fmt.Sprintf("file '%s' cannot parse by json object"), c.filePath)
	//}
	return nil
}

func (c *ConfigBase) FindKey(keys ...string) (value any, err error) {
	currentNodeMap := &c.cfg
	var currentNodeList *[]any
	var currentNodeValue *any
	currentType := 'o' // o: object; l: list; v: value
	for _, key := range keys {
		switch currentType {
		case 'o':
			if node, ok := (*currentNodeMap)[key]; ok {
				switch node.(type) {
				case map[string]any:
					temp := node.(map[string]any)
					currentNodeMap = &temp
					currentType = 'o'
				case []any:
					temp := node.([]any)
					currentNodeList = &temp
					currentType = 'l'
				default:
					currentNodeValue = &node
					currentType = 'v'
				}
			} else {
				err = errors.New(fmt.Sprintf("a dict value has no key named '%s'", key))
				value = nil
				return
			}
		case 'l':
			if i, e := strconv.Atoi(key); e == nil {
				if len(*currentNodeList) <= i {
					err = errors.New(fmt.Sprintf("list index '%d' out of range '%d'", i, len(*currentNodeList)))
					value = nil
					return
				}
				node := &(*currentNodeList)[i]
				switch (*node).(type) {
				case map[string]any:
					temp := (*node).(map[string]any)
					currentNodeMap = &temp
					currentType = 'o'
				case []any:
					temp := (*node).([]any)
					currentNodeList = &temp
					currentType = 'l'
				default:
					currentNodeValue = node
					currentType = 'v'
				}
			} else {
				err = errors.New(fmt.Sprintf("a list value has no index named '%s'", key))
				value = nil
				return
			}
		case 'v':
			err = errors.New(fmt.Sprintf("a value '%s' not have any keys '%s'", *currentNodeValue, key))
			value = nil
			return
		}
	}
	err = nil
	switch currentType {
	case 'o':
		value = *currentNodeMap
	case 'l':
		value = *currentNodeList
	case 'v':
		value = *currentNodeValue
	}
	return
}
