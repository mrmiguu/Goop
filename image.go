package goop

import "github.com/gopherjs/gopherjs/js"

type image struct {
	file   string
	g      *Global
	sprite *js.Object
}

func Image(file string) *image {
	return &image{file: "assets/" + file}
}

func (i *image) preload(g *Global) {
	i.g = g
	i.g.game.Get("load").Call("image", i.file, i.file)
}

func (i *image) create() {
	i.sprite = i.g.game.Get("add").Call("sprite", 0, 0, i.file)
}

func (i *image) Enable(tf bool) {
	i.sprite.Set("visible", tf)
}
