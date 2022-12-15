package rnd

import (
	"container/list"
	"github.com/geniot/jigsawpuzzle/internal/api"
)

type Scene struct {
	renderables *list.List
}

func NewScene() *Scene {

	scn := &Scene{}

	scn.renderables = list.New()

	scn.renderables.PushBack(NewField(scn))
	//scn.renderables.PushBack(NewDebugGrid())
	//l.PushBack(NewFpsCounter())

	return scn
}

func (scene *Scene) Render() {
	for e := scene.renderables.Front(); e != nil; e = e.Next() {
		e.Value.(api.IRenderable).Render()
	}
}

func (scene *Scene) Step(n uint64) {
	for e := scene.renderables.Front(); e != nil; e = e.Next() {
		e.Value.(api.IRenderable).Step(n)
	}
}
