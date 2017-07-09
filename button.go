package goop

type button struct {
	down, up *image
	shadow   *image // might not exist
	g        *Global
}

func Button(file string) *button {
	return &button{
		down: Image(wildcards(file, "down")),
		up:   Image(wildcards(file, "up")),
	}
}

func (b *button) preload(g *Global) {
	b.g = g
	b.down.preload(g)
	b.up.preload(g)
}

func (b *button) create() {
	b.down.create()
	b.up.create()
}

func (b *button) Enable(tf bool) {
	b.down.Enable(tf)
	b.up.Enable(tf)
}
