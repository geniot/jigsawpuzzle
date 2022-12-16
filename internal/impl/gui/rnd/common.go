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
		ctx.RendererIns.DrawRect(&sdl.Rect{x + i, y + i, glb.CELL_WIDTH - i*2, glb.CELL_HEIGHT - i*2})
	}

}

func getTextWidth(text string) int32 {
	width, _, _ := ctx.FontIns.SizeUTF8(text)
	return int32(width)
}

func drawText(txt string, x int32, y int32, color sdl.Color) int32 {
	textSurface, _ := ctx.FontIns.RenderUTF8Blended(txt, color)
	defer textSurface.Free()
	textTexture, _ := ctx.RendererIns.CreateTextureFromSurface(textSurface)
	ctx.RendererIns.Copy(textTexture, nil,
		&sdl.Rect{X: x, Y: y, W: textSurface.W, H: textSurface.H})
	defer textTexture.Destroy()
	return textSurface.W
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
