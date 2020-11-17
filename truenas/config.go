package truenas

import (
	"context"

	"github.com/adarobin/gotruenas"
)

// Config holds the provider configuration
type Config struct {
	TrueNASHost string
	Username    string
	Password    string
}

// Client returns a new client for accessing TrueNAS
func (c *Config) Client() (*gotruenas.APIClient, context.Context, error) {
	config := gotruenas.NewConfiguration()

	config.Host = c.TrueNASHost
	config.Scheme = "https"

	auth := context.WithValue(context.Background(), gotruenas.ContextBasicAuth, gotruenas.BasicAuth{
		UserName: c.Username,
		Password: c.Password,
	})

	return gotruenas.NewAPIClient(config), auth, nil
}
