package ctx

import (
	"github.com/geniot/jigsawpuzzle/internal/api"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var (
	ApplicationIns api.IApplication
	WindowIns      api.IWindow
	DeviceIns      api.IDevice
	ConfigIns      api.IConfig

	LoopIns        api.IStartable
	EventLoopIns   api.IRunnable
	PhysicsLoopIns api.IRunnable
	RenderLoopIns  api.IRunnable

	SceneIns api.IScene

	RendererIns            *sdl.Renderer
	FontIns                *ttf.Font
	PressedKeysCodesSetIns map[sdl.Keycode]int64 = make(map[sdl.Keycode]int64)
)
