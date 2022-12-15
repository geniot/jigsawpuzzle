package rnd

import (
	"github.com/geniot/jigsawpuzzle/internal/ctx"
	"github.com/geniot/jigsawpuzzle/internal/glb"
	"github.com/geniot/jigsawpuzzle/resources"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Field struct {
	background *sdl.Texture

	scene *Scene
}

func NewField(scn *Scene) *Field {
	surface, _ := img.LoadRW(resources.GetResource("image2.png"), true)
	defer surface.Free()
	txt, _ := ctx.RendererIns.CreateTextureFromSurface(surface)
	return &Field{txt, scn}
}

func (field Field) Render() {
	ctx.RendererIns.Copy(field.background, nil, &sdl.Rect{0, 0, glb.SCREEN_LOGICAL_WIDTH, glb.SCREEN_LOGICAL_HEIGHT})
}

func (field *Field) Step(n uint64) {

}
