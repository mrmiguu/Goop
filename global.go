package goop

import "github.com/gopherjs/gopherjs/js"

var (
	phaser *js.Object
)

func init() {
	style := js.Global.Get("document").Get("body").Get("style")
	style.Set("background", "#000000")
	style.Set("margin", 0)

	phaser = js.Global.Get("Phaser")
}

type Global struct {
	game   *js.Object
	assets Assets
	states States
}

func New(width, height int, a Assets, s States) {
	g := &Global{phaser.Get("Game").New(width, height), a, s}

	g.game.Get("state").Call("add", "main", js.M{
		"preload": g.preload,
		"create":  g.create,
	}, true)
}

func (G *Global) preload() {
	G.game.Get("stage").Set("background", "#000000")
	G.game.Get("canvas").Set("oncontextmenu", func(e *js.Object) {
		e.Call("preventDefault")
	})

	scale := G.game.Get("scale")
	scale.Set("scaleMode", phaser.Get("ScaleManager").Get("SHOW_ALL"))
	scale.Set("fullScreenScaleMode", phaser.Get("ScaleManager").Get("SHOW_ALL"))
	scale.Set("pageAlignHorizontally", true)
	scale.Set("pageAlignVertically", true)

	for _, asset := range G.assets {
		asset.preload(G)
	}
	for _, state := range G.states {
		for _, asset := range state.assets {
			asset.preload(G)
		}
	}
}

func (G *Global) create() {
	for _, asset := range G.assets {
		asset.create()
	}
	for _, state := range G.states {
		for _, asset := range state.assets {
			asset.create()
		}
	}
	for _, state := range G.states {
		state.controller.New(G, state.assets)
	}
}
