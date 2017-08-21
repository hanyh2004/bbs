package views

import (
	"github.com/skycoin/bbs/src/store/state/pack"
	"github.com/skycoin/cxo/skyobject"
	"sync"
	"github.com/skycoin/bbs/src/store/state/views/content_view"
)


type View interface {

	// Init initiates the view.
	Init(pack *skyobject.Pack, headers *pack.Headers, mux *sync.Mutex) error

	// Update updates the view.
	Update(pack *skyobject.Pack, headers *pack.Headers, mux *sync.Mutex) error

	// Get obtains information from the view.
	Get(id string, a ...interface{}) (interface{}, error)
}

type Adder func() (string, View)

func Add(viewsMap map[string]View, add Adder) {
	id, view := add()
	viewsMap[id] = view
}

const (
	NameContent = "content"
)

func Content() Adder {
	return func() (string, View) {
		return NameContent, new(content_view.ContentView)
	}
}
