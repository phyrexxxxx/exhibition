package handlers

import (
	"github.com/phyrexxxxx/exhibition/config"
)

// use Embedding to prevent "cannot define new methods on non-local type config.ApiConfig" error
type HandlerApiConfig struct {
	*config.ApiConfig
}
