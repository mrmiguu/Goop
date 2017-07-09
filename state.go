package goop

type state struct {
	controller controller
	assets     Assets
}

func State(c controller, a Assets) state {
	return state{c, a}
}

type controller interface {
	New(*Global, Assets)
}

type States map[string]state

// type state struct {
// 	assets Assets
// 	g      *Global
// }

// type State Asset

// func newState(a Assets) *state {
// 	return &state{assets: a}
// }

// func (s *state) preload(g *Global) {
// 	s.g = g
// 	for _, asset := range s.assets {
// 		asset.preload(g)
// 	}
// }

// func (s *state) create() {
// 	for _, asset := range s.assets {
// 		asset.create()
// 	}
// }

// func (s *state) Enable(b bool) {
// 	for _, asset := range s.assets {
// 		asset.Enable(b)
// 	}
// }

// type states struct {
// 	states States
// 	g      *Global
// }

// type States map[string]State

// func newStates(s States) *states {
// 	return &states{states: s}
// }

// func (s *states) preload(g *Global) {
// 	s.g = g
// 	for _, state := range s.states {
// 		state.preload(g)
// 	}
// }

// func (s *states) create() {
// 	for _, state := range s.states {
// 		state.create()
// 	}
// }
