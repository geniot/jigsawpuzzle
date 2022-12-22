package rnd

import (
	"container/list"
	"github.com/geniot/jigsawpuzzle/internal/ctx"
	"github.com/geniot/jigsawpuzzle/internal/glb"
)

type Scene struct {
	renderables        *list.List
	field              *Field
	frame              *Frame
	menu               *Menu
	processedTimeStamp int64
}

func NewScene() *Scene {

	scn := &Scene{}

	scn.renderables = list.New()

	scn.field = NewField(scn)
	scn.frame = NewFrame(scn)
	scn.menu = NewMenu(scn)

	scn.renderables.PushBack(scn.field)
	scn.renderables.PushBack(scn.frame)
	scn.renderables.PushBack(scn.menu)

	//scn.renderables.PushBack(NewDebugGrid())
	//scn.renderables.PushBack(NewFpsCounter())

	return scn
}

func (scene *Scene) Render() {
	scene.field.Render()
	if !scene.menu.isVisible && !scene.field.isPuzzleComplete {
		scene.frame.Render()
	}
	scene.menu.Render()
}

func (scene *Scene) Step(n uint64) {
	if p, ok := ctx.PressedKeysCodesSetIns[glb.GCW_BUTTON_MENU]; ok && p != scene.processedTimeStamp {
		scene.processedTimeStamp = p
		scene.menu.isVisible = !scene.menu.isVisible
	}
	if scene.menu.isVisible {
		scene.menu.Step(n)
	} else {
		if !scene.field.isPuzzleComplete {
			scene.frame.Step(n)
		}
	}
}

func (scene *Scene) NewGame() {
	scene.menu.isVisible = false
	scene.field.Reset()
}
