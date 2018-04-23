package toml

import (
	"io"

	"git.containerum.net/ch/cherry/pkg/models"
	shushiToml "github.com/BurntSushi/toml"
)

func ParseService(re io.Reader) (models.Service, error) {
	var service models.Service
	_, err := shushiToml.DecodeReader(re, &service)
	return service, err
}
