package googleworkspace_client

import (
	"context"
	"errors"
	"strings"

	"github.com/selefra/selefra-provider-googleworkspace/constants"
)

type Client struct {
	Config *Config
}

func NewClients(config Config) ([]*Client, error) {
	return []*Client{&Client{Config: &config}}, nil
}

func IsCancelled(ctx context.Context) bool {
	err := ctx.Err()
	return err != nil && (errors.Is(err, context.Canceled) || strings.Contains(err.Error(), constants.Contextcanceled))
}
