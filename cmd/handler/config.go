package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
)

type GlobalConfig struct {
	SessionFile  string `json:"session_file"`
	TemplateFile string `json:"template_file"`
	MainFileName string `json:"main_file_name"`
}

// ShowGlobalConfig shows global config.
func (h *Handler) ShowGlobalConfig() error {
	c, err := readConfig(h.configFile)
	if err != nil {
		return err
	}

	t := reflect.TypeOf(*c)
	v := reflect.ValueOf(c)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("%s: %s\n", f.Tag.Get("json"), v.Elem().FieldByName(f.Name))
	}

	return nil
}

// touchReadConfig reads config from configFile.
// If doesn't exist, write out default value.
func touchReadConfig(configFile string, defaultC *GlobalConfig) (*GlobalConfig, error) {
	c, err := readConfig(configFile)
	if errors.Is(err, os.ErrNotExist) {
		return defaultC, writeConfig(configFile, defaultC)
	}
	return c, err
}

func readConfig(configFile string) (*GlobalConfig, error) {
	b, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	c := &GlobalConfig{}
	return c, json.Unmarshal(b, c)
}

func writeConfig(configFile string, config *GlobalConfig) error {
	if err := os.MkdirAll(filepath.Dir(configFile), 0755); err != nil {
		return err
	}
	b, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(configFile, b, 0755)
}
