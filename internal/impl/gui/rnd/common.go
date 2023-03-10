package rnd

import (
	"github.com/geniot/jigsawpuzzle/internal/ctx"
	"github.com/geniot/jigsawpuzzle/internal/glb"
	"github.com/geniot/jigsawpuzzle/resources"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func drawThickRect(x int32, y int32, thickness int32) {
	for i := int32(0); i < thickness; i++ {
		drawRect(x+i, y+i, glb.CELL_WIDTH-i*2, glb.CELL_HEIGHT-i*2)
	}

}

func drawRect(x int32, y int32, width int32, height int32) {
	ctx.RendererIns.DrawLine(x, y, x+width, y)
	ctx.RendererIns.DrawLine(x+width, y, x+width, y+height)
	ctx.RendererIns.DrawLine(x+width, y+width, x, y+height)
	ctx.RendererIns.DrawLine(x, y+width, x, y)
}

func drawText(txt string, x int32, y int32, color sdl.Color) (int32, int32) {
	textSurface, _ := ctx.FontIns.RenderUTF8Blended(txt, color)
	defer textSurface.Free()
	textTexture, _ := ctx.RendererIns.CreateTextureFromSurface(textSurface)
	ctx.RendererIns.Copy(textTexture, nil,
		&sdl.Rect{X: x, Y: y, W: textSurface.W, H: textSurface.H})
	defer textTexture.Destroy()
	return textSurface.W, textSurface.H
}

func loadTexture(fileName string) *sdl.Texture {
	surface, _ := img.LoadRW(resources.GetResource(fileName), true)
	defer surface.Free()
	txt, err := ctx.RendererIns.CreateTextureFromSurface(surface)
	if err != nil {
		println(err.Error())
	}
	return txt
}
