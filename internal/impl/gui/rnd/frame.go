package rnd

import (
	"github.com/geniot/jigsawpuzzle/internal/ctx"
	"github.com/geniot/jigsawpuzzle/internal/glb"
	"time"
)

type Frame struct {
	cellX          int32
	cellY          int32
	processedRight int64
	processedLeft  int64
	processedUp    int64
	processedDown  int64
	scene          *Scene
}

func NewFrame(scn *Scene) *Frame {
	s := time.Now().UnixNano()
	return &Frame{0, 0, s, s, s, s, scn}
}

func (frame Frame) Render() {
	ctx.RendererIns.SetDrawColor(0, 0, 0, 150)
	x := frame.cellX * glb.CELL_WIDTH
	y := frame.cellY * glb.CELL_HEIGHT
	drawThickRect(x, y, 3)
}

func (frame *Frame) Step(n uint64) {
	if p, ok := ctx.PressedKeysCodesSetIns[glb.GCW_BUTTON_RIGHT]; ok && p != frame.processedRight {
		frame.processedRight = p
		frame.cellX += 1
		if frame.cellX >= glb.CELLS_HORIZONTAL {
			frame.cellX = glb.CELLS_HORIZONTAL - 1
		}
	} else if p, ok = ctx.PressedKeysCodesSetIns[glb.GCW_BUTTON_LEFT]; ok && p != frame.processedLeft {
		frame.processedLeft = p
		frame.cellX -= 1
		if frame.cellX <= 0 {
			frame.cellX = 0
		}
	} else if p, ok = ctx.PressedKeysCodesSetIns[glb.GCW_BUTTON_UP]; ok && p != frame.processedUp {
		frame.processedUp = p
		frame.cellY -= 1
		if frame.cellY <= 0 {
			frame.cellY = 0
		}
	} else if p, ok = ctx.PressedKeysCodesSetIns[glb.GCW_BUTTON_DOWN]; ok && p != frame.processedDown {
		frame.processedDown = p
		frame.cellY += 1
		if frame.cellY >= glb.CELLS_VERTICAL {
			frame.cellY = glb.CELLS_VERTICAL - 1
		}
	}
}
