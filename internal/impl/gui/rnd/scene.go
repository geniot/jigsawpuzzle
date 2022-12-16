package rnd

import (
	"container/list"
	"github.com/geniot/jigsawpuzzle/internal/api"
)

type Scene struct {
	renderables *list.List
	field       *Field
	frame       *Frame
}

func NewScene() *Scene {

	scn := &Scene{}

	scn.renderables = list.New()

	scn.field = NewField(scn)
	scn.frame = NewFrame(scn)

	scn.renderables.PushBack(scn.field)
	scn.renderables.PushBack(scn.frame)

	//scn.renderables.PushBack(NewDebugGrid())
	//scn.renderables.PushBack(NewFpsCounter())

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
