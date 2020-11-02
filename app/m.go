package app

import (
	"sort"
	"sync"
)

type appstate int

// APP states
const (
	AppStateLoad appstate = iota
	AppStateInit
	AppStatePending
	AppStateRunning
	AppStateClose
)

//
type AppItem struct {
	App     *App
	state   appstate
	weight  int
	depApps []*AppItem
}

//
type Registry struct {
	ordered bool
	lock    sync.Mutex
	apps    appList
}

type appList []*AppItem

func (l appList) Len() int {
	return len(l)
}

func (l appList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l appList) Less(i, j int) bool {
	if l[i].calWeigh() < l[j].calWeigh() {
		return true
	}
	return false
}

//
func (r *Registry) Order() bool {
	r.lock.Lock()
	defer r.lock.Unlock()

	for _, i := range r.apps {
		depAppNames := i.App.GetDepApps()
		count := 0
		for _, n := range depAppNames {
			for _, ii := range r.apps {
				if i == ii {
					continue
				}
				if n == ii.App.name {
					i.depApps = append(i.depApps, ii)
					count++
				}
			}
		}
		if count != len(depAppNames) {
			// TODO || add log
			r.ordered = false
			return false
		}
	}
	defer func() { r.ordered = true }()
	// all ready
	for _, i := range r.apps {
		i.weight = i.calWeigh()
	}
	sort.Sort(r.apps)
	return true
}

func (ai *AppItem) calWeigh() int {
	w := 0
	if len(ai.depApps) == 0 {
		return w
	}

	for _, i := range ai.depApps {
		w += i.calWeigh() + 1
	}
	return w
}

func (l *appList) add(a *AppItem) {
	*l = append(*l, a)
}

//
func (r *Registry) Add(a *App) {
	r.lock.Lock()
	r.lock.Unlock()
	defer func() { r.ordered = false }()

	ai := &AppItem{
		App: a, state: AppStateLoad,
	}
	r.apps = append(r.apps, ai)
}
