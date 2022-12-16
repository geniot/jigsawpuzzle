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
	if _, ok := ctx.PressedKeysCodesSetIns[glb.GCW_BUTTON_A]; ok {
		ctx.RendererIns.SetDrawColor(255, 0, 0, 150)
	} else {
		ctx.RendererIns.SetDrawColor(0, 0, 0, 150)
	}

	x := frame.cellX * glb.CELL_WIDTH
	y := frame.cellY * glb.CELL_HEIGHT
	drawThickRect(x, y, 3)
}

func (frame *Frame) Step(n uint64) {
	if p, ok := ctx.PressedKeysCodesSetIns[glb.GCW_BUTTON_RIGHT]; ok && p != frame.processedRight {
		frame.processedRight = p
		if frame.cellX < glb.CELLS_HORIZONTAL-1 {
			frame.scene.field.switchTiles(frame.cellX, frame.cellY, frame.cellX+1, frame.cellY)
			frame.cellX += 1
		}
	} else if p, ok = ctx.PressedKeysCodesSetIns[glb.GCW_BUTTON_LEFT]; ok && p != frame.processedLeft {
		frame.processedLeft = p
		if frame.cellX > 0 {
			frame.scene.field.switchTiles(frame.cellX, frame.cellY, frame.cellX-1, frame.cellY)
			frame.cellX -= 1
		}
	} else if p, ok = ctx.PressedKeysCodesSetIns[glb.GCW_BUTTON_UP]; ok && p != frame.processedUp {
		frame.processedUp = p
		if frame.cellY > 0 {
			frame.scene.field.switchTiles(frame.cellX, frame.cellY, frame.cellX, frame.cellY-1)
			frame.cellY -= 1
		}
	} else if p, ok = ctx.PressedKeysCodesSetIns[glb.GCW_BUTTON_DOWN]; ok && p != frame.processedDown {
		frame.processedDown = p
		if frame.cellY < glb.CELLS_VERTICAL-1 {
			frame.scene.field.switchTiles(frame.cellX, frame.cellY, frame.cellX, frame.cellY+1)
			frame.cellY += 1
		}
	}
}
