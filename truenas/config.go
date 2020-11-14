package truenas

import (
	"context"

	"github.com/adarobin/gotruenas"
)

// Config holds the provider configuration
type Config struct {
	TrueNASHost string
}

// Client returns a new client for accessing Red Hat Satellite
func (c *Config) Client() (*gotruenas.APIClient, context.Context, error) {
	config := gotruenas.NewConfiguration()

	// token, err := gorhsm.GenerateAccessToken(c.RefreshToken)
	// if err != nil {
	// 	return nil, nil, err
	// }

	// auth := context.WithValue(context.Background(), gorhsm.ContextAPIKey, gorhsm.APIKey{
	// 	Key:    token.AccessToken,
	// 	Prefix: "Bearer", // Omit if not necessary.
	// })

	return gotruenas.NewAPIClient(config), auth, nil
}
