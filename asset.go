package goop

type Asset interface {
	preload(*Global)
	create()
	Enable(bool)
}

type Assets map[string]Asset
