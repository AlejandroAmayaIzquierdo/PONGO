package game

import (
	"fmt"
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const DRAW_INTERVAL float64 = 1000000000.0 / float64(FPS)
const FPS int = 60

type Game struct {
	window    *sdl.Window
	renderer  *sdl.Renderer
	deltaTime float64
	delta     float64
	timer     float64
	drawCount int
	lastTime  int
	running   bool

	xPos int
	yPos int
}

func NewGame(win *sdl.Window) *Game {
	g := new(Game)

	g.running = true

	g.window = win

	g.renderer, _ = sdl.CreateRenderer(win, -1, sdl.RENDERER_ACCELERATED)

	g.xPos = (1280 / 2) - 50
	g.yPos = (720 / 2) - 50

	return g
}

func (g *Game) Clean() {
	g.running = false
}

func (g *Game) Handle() {

	g.lastTime = time.Now().Nanosecond()

	for g.running {
		currentTime := time.Now().Nanosecond()

		g.deltaTime = math.Abs(float64(currentTime-g.lastTime)) / 1000000000.0
		g.delta += math.Abs(float64(currentTime-g.lastTime)) / DRAW_INTERVAL
		g.timer += float64(math.Abs(float64(currentTime - g.lastTime)))

		g.Update()

		if g.delta >= 1.0 {
			g.Draw()

			g.delta--
			g.drawCount++
		}

		if g.timer >= 1000000000.0 {
			fmt.Printf("FPS: %d\n", g.drawCount)
			g.timer = 0
			g.drawCount = 0
		}

		g.lastTime = currentTime
	}
}

func (g *Game) Update() {
	for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
		switch e.GetType() {
		case sdl.QUIT:
			g.Clean()
		case sdl.KEYDOWN:
			keyEvent := e.(*sdl.KeyboardEvent)
			if keyEvent.Keysym.Sym == sdl.GetKeyFromName("d") {
				g.xPos += 10

			} else if keyEvent.Keysym.Sym == sdl.GetKeyFromName("a") {
				// fmt.Printf("Key pressed: %d\n", keyEvent.Keysym.Sym)
				g.xPos -= 10
			}
			// case sdl.KEYUP:
			// 	keyEvent := e.(*sdl.KeyboardEvent)
			// 	fmt.Printf("Key released: %d\n", keyEvent.Keysym.Sym)
		}
	}
}

func (g *Game) Draw() {
	g.renderer.SetDrawColor(0, 0, 0, 255)
	g.renderer.Clear()

	g.renderer.SetDrawColor(255, 255, 255, 255)
	g.renderer.FillRect(&sdl.Rect{X: int32(g.xPos), Y: int32(g.yPos), W: 100, H: 100})

	g.renderer.Present()
}
