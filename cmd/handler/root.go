package handler

import (
	"fmt"

	"github.com/tbistr/atcoder-go/atcodergo"
)

type Handler struct {
	atcoder    *atcodergo.Client
	configFile string
	config     *GlobalConfig
}

func New(configFile string, defauldConfig *GlobalConfig) (*Handler, error) {
	a, err := atcodergo.NewClient()
	if err != nil {
		return nil, fmt.Errorf("failed to init atcoder-go library: %w", err)
	}

	c, err := touchReadConfig(configFile, defauldConfig)
	if err != nil {
		return nil, err
	}

	a.LoginWithSession(c.SessionFile)
	return &Handler{
		atcoder:    a,
		configFile: configFile,
		config:     c,
	}, nil
}
