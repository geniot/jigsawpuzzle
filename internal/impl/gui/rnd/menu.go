package rnd

import (
	"github.com/geniot/jigsawpuzzle/internal/ctx"
	"github.com/geniot/jigsawpuzzle/internal/glb"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

type Menu struct {
	isVisible     bool
	processedMenu int64
	scene         *Scene
}

func NewMenu(scn *Scene) *Menu {
	s := time.Now().UnixNano()
	return &Menu{false, s, scn}
}

func (menu Menu) Render() {
	if menu.isVisible {
		ctx.RendererIns.SetDrawColor(192, 192, 192, 150)
		ctx.RendererIns.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
		ctx.RendererIns.FillRect(&sdl.Rect{0, 0, glb.SCREEN_LOGICAL_WIDTH, glb.SCREEN_LOGICAL_HEIGHT})
	}
}

func (menu *Menu) Step(n uint64) {
	if p, ok := ctx.PressedKeysCodesSetIns[glb.GCW_BUTTON_MENU]; ok && p != menu.processedMenu {
		menu.processedMenu = p
		menu.isVisible = !menu.isVisible
	}
}
