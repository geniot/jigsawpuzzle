package rnd

import (
	"github.com/geniot/jigsawpuzzle/internal/api"
	"github.com/geniot/jigsawpuzzle/internal/ctx"
	"github.com/geniot/jigsawpuzzle/internal/glb"
	"time"
)

type MenuItem struct {
	offsetX            int32
	offsetY            int32
	selectorWidth      int32
	isSelected         bool
	text               string
	actionHandler      api.ActionHandler
	processedTimeStamp int64
}

func NewMenuItem(oX int32, oY int32, isS bool, txt string, ah api.ActionHandler) *MenuItem {
	width, _, _ := ctx.FontIns.SizeUTF8(">")
	return &MenuItem{oX, oY, int32(width), isS, txt, ah, time.Now().UnixNano()}
}

func (menuItem MenuItem) Render() {
	drawText(menuItem.text, menuItem.offsetX, menuItem.offsetY, glb.COLOR_BLACK)
	if menuItem.isSelected {
		drawText(">", menuItem.offsetX-menuItem.selectorWidth, menuItem.offsetY, glb.COLOR_BLACK)
	}
}

func (menuItem *MenuItem) Step(n uint64) {
	if menuItem.isSelected {
		if p, ok := ctx.PressedKeysCodesSetIns[glb.GCW_BUTTON_A]; ok && p != menuItem.processedTimeStamp {
			menuItem.processedTimeStamp = p
			menuItem.actionHandler()
		}
	}
}
