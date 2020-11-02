package app

import (
	"proxy-server/common"
	"proxy-server/pipes"
	"sync"
)

var once sync.Once

// app abstract interface and struct
//
//

// Interface export
type Interface interface {
	Run(mp pipes.MsgPipe) (ok bool, runner func())
}

type App struct {
	name    string
	config  common.Config
	clsName string
	depApps []string
}

//
func (a *App) Reference() string {
	return a.name
}

//
func (a *App) JavaClassName() string {
	return a.clsName
}

func (a *App) GetDepApps() []string {
	return a.depApps
}

func init() {
	once.Do(func() {
		//  init registery

	})
}


func Load()  {
	
}