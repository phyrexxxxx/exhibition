package auth

import (
	"github.com/phyrexxxxx/exhibition/config"
)

// use Embedding to prevent "cannot define new methods on non-local type config.ApiConfig" error
type AuthApiConfig struct {
	*config.ApiConfig
}
