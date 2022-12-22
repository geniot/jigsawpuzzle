package rnd

import (
	"github.com/geniot/jigsawpuzzle/internal/ctx"
	. "github.com/geniot/jigsawpuzzle/internal/glb"
	"github.com/geniot/jigsawpuzzle/resources"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"math"
	"math/rand"
	"strconv"
	"time"
)

type Field struct {
	cellWidth        int32
	cellHeight       int32
	tiles            []int
	rotations        []int
	background       *sdl.Texture
	isPuzzleComplete bool
	scene            *Scene
}

func NewField(scn *Scene) *Field {
	fld := &Field{}
	fld.Reset()
	return fld
}

func (field *Field) Reset() {
	rand.Seed(time.Now().UnixNano())
	tls := [CELLS_VERTICAL * CELLS_HORIZONTAL]int{}
	rtn := [CELLS_VERTICAL * CELLS_HORIZONTAL]int{}
	for i := 0; i < len(tls); i++ {
		tls[i] = i
		rtn[i] = rand.Intn(4)
	}

	for i := 0; i < 1000; i++ {
		rand.Shuffle(len(tls), func(i, j int) { tls[i], tls[j] = tls[j], tls[i] })
		if IsShufflePerfect(tls[:], CELLS_HORIZONTAL) {
			break
		}
	}

	randomImageIndex := strconv.FormatInt(int64(rand.Intn(3)+1), 10)
	surface, _ := img.LoadRW(resources.GetResource("image"+randomImageIndex+".png"), true)
	defer surface.Free()

	field.tiles = tls[:]
	field.rotations = rtn[:]
	field.background, _ = ctx.RendererIns.CreateTextureFromSurface(surface)
	field.cellWidth = surface.W / CELLS_HORIZONTAL
	field.cellHeight = surface.H / CELLS_VERTICAL
	field.isPuzzleComplete = false
}

func (field *Field) Render() {
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

func (field *Field) switchTiles(x1 int32, y1 int32, x2 int32, y2 int32) {
	if _, ok := ctx.PressedKeysCodesSetIns[GCW_BUTTON_A]; ok {
		index1 := y1*CELLS_HORIZONTAL + x1
		index2 := y2*CELLS_HORIZONTAL + x2

		tmpTile := field.tiles[index1]
		field.tiles[index1] = field.tiles[index2]
		field.tiles[index2] = tmpTile

		tmpRotation := field.rotations[index1]
		field.rotations[index1] = field.rotations[index2]
		field.rotations[index2] = tmpRotation
	}
}

func (field *Field) checkComplete() {
	for i := 0; i < len(field.tiles); i++ {
		if field.tiles[i] != i {
			field.isPuzzleComplete = false
			return
		}
	}
	field.isPuzzleComplete = true
}
