package handler

import (
	"fmt"

	"github.com/tbistr/atcoder-go/atcodergo"
)

type Handler struct {
	atcoder *atcodergo.Client
}

func New() (*Handler, error) {
	a, err := atcodergo.NewClient()
	if err != nil {
		return nil, fmt.Errorf("failed to init atcoder-go library: %w", err)
	}
	return &Handler{
		atcoder: a,
	}, nil
}
