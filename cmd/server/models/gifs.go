package models

import (
	"errors"
	"image/gif"
)

type GifStoreInterface interface {
	Add(*gif.GIF, int) error
	Get(int) (*gif.GIF, bool)
	Delete(int)
}

type GifStore struct {
	data map[int]*gif.GIF
}

func GifStoreInit() *GifStore {
	return &GifStore{data: map[int]*gif.GIF{}}
}

func (gs *GifStore) Add(g *gif.GIF, index int) error {
	if _, ok := gs.Get(index); ok {
		return errors.New("index taken")
	}
	gs.data[index] = g
	return nil
}

func (gs *GifStore) Get(index int) (*gif.GIF, bool) {
	g, ok := gs.data[index]
	return g, ok
}

func (gs *GifStore) Delete(index int) {
	if _, ok := gs.Get(index); ok {
		delete(gs.data, index)
	}
}
