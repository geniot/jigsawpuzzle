package gui

import (
	"github.com/geniot/jigsawpuzzle/internal/ctx"
	"github.com/geniot/jigsawpuzzle/internal/glb"
	"github.com/geniot/jigsawpuzzle/internal/impl/dev"
	"github.com/geniot/jigsawpuzzle/internal/impl/gui/loop"
	"github.com/geniot/jigsawpuzzle/internal/impl/gui/rnd"
	"github.com/geniot/jigsawpuzzle/resources"
	"github.com/veandco/go-sdl2/ttf"
)

type ApplicationImpl struct {
}

func NewApplication() *ApplicationImpl {
	return &ApplicationImpl{}
}

func (app *ApplicationImpl) Start() {
	ctx.DeviceIns = dev.NewDevice()
	ctx.FontIns, _ = ttf.OpenFontRW(resources.GetResource(glb.FONT_FILE_NAME), 1, glb.FONT_SIZE)
	ctx.ConfigIns = NewConfig()
	ctx.WindowIns = NewWindow()

	ctx.LoopIns = loop.NewLoop()
	ctx.EventLoopIns = loop.NewEventLoop()
	ctx.PhysicsLoopIns = loop.NewPhysicsLoop()
	ctx.RenderLoopIns = loop.NewRenderLoop()

	ctx.SceneIns = rnd.NewScene()

	ctx.LoopIns.Start()

	//graceful shutdown :) we let the loop finish all rendering/processing
	app.Stop()
}

func (app *ApplicationImpl) Stop() {
	ctx.FontIns.Close()
	ctx.DeviceIns.Stop()
}
