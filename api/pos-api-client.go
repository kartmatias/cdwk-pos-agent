package api

import (
	"github.com/kartmatias/cdwk-pos-agent/cfg"
	"go.uber.org/zap"
)

const path_get_product = "products"

var urlBase string

func setup(logger *zap.Logger) {
	mycfg := cfg.GetInstance()
	urlBase = mycfg.BaseUrl
}
