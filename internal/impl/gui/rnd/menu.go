package rnd

import (
	"github.com/geniot/jigsawpuzzle/internal/api"
	"github.com/geniot/jigsawpuzzle/internal/ctx"
	"github.com/geniot/jigsawpuzzle/internal/glb"
	"github.com/veandco/go-sdl2/sdl"
	"math"
	"time"
)

type Menu struct {
	menuItems              []*MenuItem
	isVisible              bool
	processedMenuTimeStamp int64
	scene                  *Scene
}

func NewMenu(scn *Scene) *Menu {
	ahs := [2]api.ActionHandler{scn.NewGame, ctx.LoopIns.Stop}
	items := [2]string{"New Game", "Quit"}
	menuWidth, menuHeight := 0, 0
	for i := 0; i < len(items); i++ {
		width, height, _ := ctx.FontIns.SizeUTF8(items[i])
		menuWidth = int(math.Max(float64(menuWidth), float64(width)))
		menuHeight += height
	}

	offsetX := (glb.SCREEN_LOGICAL_WIDTH - menuWidth) / 2
	offsetY := (glb.SCREEN_LOGICAL_HEIGHT - menuHeight) / 2
	mItems := [len(items)]*MenuItem{}
	for i := 0; i < len(items); i++ {
		mItems[i] = NewMenuItem(int32(offsetX), int32(offsetY), glb.If(i == 0, true, false), items[i], ahs[i])
		offsetY += menuHeight / len(items)
	}

	s := time.Now().UnixNano()
	return &Menu{mItems[:], true, s, scn}
}

func (menu Menu) Render() {
	if menu.isVisible {
		ctx.RendererIns.SetDrawColor(192, 192, 192, 150)
		ctx.RendererIns.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
		ctx.RendererIns.FillRect(&sdl.Rect{0, 0, glb.SCREEN_LOGICAL_WIDTH, glb.SCREEN_LOGICAL_HEIGHT})
		for i := range menu.menuItems {
			menu.menuItems[i].Render()
		}
	}
}

func (menu *Menu) Step(n uint64) {
	if p, ok := ctx.PressedKeysCodesSetIns[glb.GCW_BUTTON_DOWN]; ok && p != menu.processedMenuTimeStamp {
		menu.processedMenuTimeStamp = p
		for i := range menu.menuItems {
			if menu.menuItems[i].isSelected && i < len(menu.menuItems)-1 {
				menu.menuItems[i].isSelected = false
				menu.menuItems[i+1].isSelected = true
			}
		}
	} else if p, ok = ctx.PressedKeysCodesSetIns[glb.GCW_BUTTON_UP]; ok && p != menu.processedMenuTimeStamp {
		menu.processedMenuTimeStamp = p
		for i := range menu.menuItems {
			if menu.menuItems[i].isSelected && i > 0 {
				menu.menuItems[i].isSelected = false
				menu.menuItems[i-1].isSelected = true
			}
		}
	}
	for i := range menu.menuItems {
		menu.menuItems[i].Step(n)
	}
}
