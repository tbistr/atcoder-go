package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	"github.com/tbistr/atcoder-go/atcodergo"
)

type GlobalConfig struct {
	SessionFile          string             `json:"session_file"`
	TemplateCmdName      string             `json:"template_cmd_name"`
	TemplateCmdArgs      []string           `json:"template_cmd_args"`
	TemplateCmdJsonInput bool               `json:"template_cmd_json_input"`
	TemplateFile         string             `json:"template_file"`
	RunCmdName           string             `json:"run_cmd_name"`
	RunCmdArgs           []string           `json:"run_cmd_args"`
	MainFileName         string             `json:"main_file_name"`
	DefaultLanguage      atcodergo.Language `json:"default_language"`
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
		fmt.Printf("%s: %#v\n", f.Tag.Get("json"), v.Elem().FieldByName(f.Name))
	}

	return nil
}

// touchReadConfig reads config from configFile.
// If doesn't exist, write out default value.
func touchReadConfig(configFile string, defaultC *GlobalConfig) (*GlobalConfig, error) {
	c, err := readConfig(configFile)
	if errors.Is(err, os.ErrNotExist) {
		f, _ := os.OpenFile(defaultC.TemplateFile, os.O_CREATE, 0644)
		f.Close()
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
