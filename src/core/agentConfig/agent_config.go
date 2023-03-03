package agentConfig

import (
	"PerfmonGo/base/configBase"
	"errors"
	"fmt"
	"strconv"
)

type AgentConfig struct {
	configBase.ConfigBase
}

func New(filePath string) (*AgentConfig, error) {
	c, err := configBase.New(filePath)
	if err != nil {
		return nil, err
	}
	ac := &AgentConfig{*c}

	if _, err := ac.GetAgentName(); err != nil {
		return nil, errors.New(fmt.Sprintf("configBase need 'agent_name' key: %s", err.Error()))
	}
	if _, err := ac.GetPerfmonItems(); err != nil {
		return nil, errors.New(fmt.Sprintf("configBase need 'perfmon' key: %s", err.Error()))
	}
	return ac, nil
}

func (c *AgentConfig) GetAgentName() (string, error) {
	val, err := c.FindKey("agent_name")
	if err != nil {
		return "", err
	}
	if name, ok := val.(string); ok {
		return name, nil
	}
	return "", errors.New("configBase key 'agent_name' not a string")
}

func (c *AgentConfig) GetPerfmonItems() ([]any, error) {
	val, err := c.FindKey("perfmon")
	if err != nil {
		return nil, err
	}
	if items, ok := val.([]any); ok {
		return items, nil
	}
	return nil, errors.New("configBase key 'perfmon' not a list")
}

func (c *AgentConfig) GetSubmitConfig() (map[string]any, error) {
	val, err := c.FindKey("submit")
	if err != nil {
		return nil, err
	}
	if submit, ok := val.(map[string]any); ok {
		return submit, nil
	}
	return nil, errors.New("configBase key 'submit' not a dict<string>")
}

func (c *AgentConfig) GetReportUrl() (string, error) {
	val, err := c.FindKey("report")
	if err != nil {
		return "", err
	}
	if report, ok := val.(string); ok {
		return report, nil
	}
	return "", errors.New("configBase key 'report' not a string")
}

func (c *AgentConfig) GetProcessCount() (int, error) {
	val, err := c.FindKey("process")
	if err != nil {
		return 0, err
	}
	if process, ok := val.(string); ok {
		i, err := strconv.Atoi(process)
		if err == nil {
			return i, nil
		}
		return 0, err
	}
	return 0, errors.New(fmt.Sprintf("cannot parse '%s' as integer with key 'process'", val))
}
