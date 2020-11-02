package bootstrap

import (
	"proxy-server/app"
	"proxy-server/common"
)

type Option struct {
	Name    string
	Value   string
	AppName string
}

type Component interface {
	Load()
	Run()
	Stop()
}

// BootStrap export
type BootStrap struct {
	loaderConfig common.Config
	appConfig    common.Config
	apps         []app.App
}

//
func (b *BootStrap) Load() {

}

func (b *BootStrap) orderLoadApps() {

}

//
func (b *BootStrap) Perpare(op ...Option) (bool, func()) {
	// Load config
	configMap := make(map[string]map[string]*Option)

	for _, o := range op {
		if v, e := configMap[o.AppName]; e == true {
			v[o.Name] = &o
		} else {
			configMap[o.AppName] = make(map[string]*Option)
		}
	}

	// init components

	// init app

	return true, func() {

	}
}
