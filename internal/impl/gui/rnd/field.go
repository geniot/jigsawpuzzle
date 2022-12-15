package rnd

import (
	"github.com/geniot/jigsawpuzzle/internal/ctx"
	. "github.com/geniot/jigsawpuzzle/internal/glb"
	"github.com/geniot/jigsawpuzzle/resources"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"math"
	"math/rand"
	"time"
)

type Field struct {
	cellWidth  int32
	cellHeight int32
	tiles      []int
	rotations  []int
	background *sdl.Texture
	scene      *Scene
}

func NewField(scn *Scene) *Field {
	rand.Seed(time.Now().UnixNano())
	tls := [CELLS_VERTICAL * CELLS_HORIZONTAL]int{}
	rtn := [CELLS_VERTICAL * CELLS_HORIZONTAL]int{}
	for i := 0; i < len(tls); i++ {
		tls[i] = i
		rtn[i] = rand.Intn(4)
	}
	rand.Shuffle(len(tls), func(i, j int) { tls[i], tls[j] = tls[j], tls[i] })

	surface, _ := img.LoadRW(resources.GetResource("image2.png"), true)
	defer surface.Free()
	txt, _ := ctx.RendererIns.CreateTextureFromSurface(surface)
	cW := surface.W / CELLS_HORIZONTAL
	cH := surface.H / CELLS_VERTICAL
	return &Field{cW, cH, tls[:], rtn[:], txt, scn}
}

func (field Field) Render() {
	for i := 0; i < len(field.tiles); i++ {
		tile := field.tiles[i]
		angle := float64(0) //float64(field.rotations[i] * 90)

		sourceX := int32(math.Mod(float64(tile), CELLS_VERTICAL+1))
		sourceY := int32(tile / CELLS_HORIZONTAL)
		targetX := int32(math.Mod(float64(i), CELLS_VERTICAL+1))
		targetY := int32(i / CELLS_HORIZONTAL)

		srcRect := &sdl.Rect{sourceX * field.cellWidth, sourceY * field.cellHeight, field.cellWidth, field.cellHeight}
		trgRect := &sdl.Rect{targetX * CELL_WIDTH, targetY * CELL_HEIGHT, CELL_WIDTH, CELL_HEIGHT}
		flip := sdl.FLIP_NONE

		ctx.RendererIns.CopyEx(field.background, srcRect, trgRect, angle, &sdl.Point{CELL_WIDTH / 2, CELL_HEIGHT / 2}, flip)
	}
}

func (field *Field) Step(n uint64) {

}
