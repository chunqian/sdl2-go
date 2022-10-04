package sdl

// typedef
type WindowShapeMode = int32

// struct
const (
	ShapeModeDefault              WindowShapeMode = 0
	ShapeModeBinarizeAlpha        WindowShapeMode = 1
	ShapeModeReverseBinarizeAlpha WindowShapeMode = 2
	ShapeModeColorKey             WindowShapeMode = 3
)

// struct
type SDL_WindowShapeParams struct {
	ColorKey SDL_Color
}

type SDL_WindowShapeMode struct {
	Mode       int32
	Parameters SDL_WindowShapeParams
}
