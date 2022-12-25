package handler

import (
	"fmt"

	"github.com/tbistr/atcoder-go/atcodergo"
)

type Handler struct {
	atcoder     *atcodergo.Client
	sessionFile string
}

func New(sessionFile string) (*Handler, error) {
	a, err := atcodergo.NewClient()
	if err != nil {
		return nil, fmt.Errorf("failed to init atcoder-go library: %w", err)
	}

	return &Handler{
		atcoder:     a,
		sessionFile: sessionFile,
	}, nil
}
