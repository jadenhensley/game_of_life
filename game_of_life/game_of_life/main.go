package main

import (
	"image/color"
	"log"
	"math"
	"time"

	"game_of_life/game_of_life/mywidgets"

	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	CELL_COUNT_WIDTH = 32
	SCREEN_RATIO     = 8
	SCALE_FACTOR     = 80
	WIDTH            = SCREEN_RATIO * SCALE_FACTOR
	HEIGHT           = WIDTH // 8 * 80 = 640 by 640
	CELL_SIZE        = WIDTH / CELL_COUNT_WIDTH
	CELL_COUNT_TOTAL = CELL_COUNT_WIDTH * CELL_COUNT_WIDTH
)

var PAUSED = true
var pastTime time.Time = time.Now()
var loopRan bool = false
var image *ebiten.Image = ebiten.NewImage(WIDTH, HEIGHT)

type Game struct {
	ui    *ebitenui.UI
	grids [2][CELL_COUNT_WIDTH][CELL_COUNT_WIDTH]bool
}

func RGBAFromPercent(r, g, b, a float64) color.RGBA {
	iR := uint8(math.Round(r * 255))
	iG := uint8(math.Round(g * 255))
	iB := uint8(math.Round(b * 255))
	iA := uint8(math.Round(a * 255))
	return color.RGBA{iR, iG, iB, iA}
}

func getNeighborCount(g *Game, x int, y int) int {
	count := 0
	if (x - 1) >= 0 {
		if g.grids[0][y][x-1] {
			count += 1
		}
		if (y - 1) >= 0 {
			if g.grids[0][y-1][x-1] {
				count += 1
			}
		}
		if (y + 1) <= CELL_COUNT_WIDTH-1 {
			if g.grids[0][y+1][x-1] {
				count += 1
			}
		}
	}

	if (x + 1) <= CELL_COUNT_WIDTH-1 {
		if g.grids[0][y][x+1] {
			count += 1
		}
		if (y - 1) >= 0 {
			if g.grids[0][y-1][x+1] {
				count += 1
			}
		}
		if (y + 1) <= CELL_COUNT_WIDTH-1 {
			if g.grids[0][y+1][x+1] {
				count += 1
			}
		}
	}

	if (y - 1) >= 0 {
		if g.grids[0][y-1][x] {
			count += 1
		}
	}
	if (y + 1) <= CELL_COUNT_WIDTH-1 {
		if g.grids[0][y+1][x] {
			count += 1
		}
	}

	return count

}

func (g *Game) Update() error {
	g.ui.Update()

	currentTime := time.Now()

	duration := currentTime.Sub(pastTime)
	secondsSince := int(duration.Seconds())
	// fmt.Printf("%d\n", secondsSince)

	// fmt.Printf("%d\n", seconds)

	mX, mY := ebiten.CursorPosition()
	// fmt.Printf("%d,%d MOUSE POS\n", mX, mY)
	pressed := inpututil.PressedKeys()
	for _, key := range pressed {
		if key == ebiten.KeyEnter {
			PAUSED = !PAUSED
		}
	}

	if PAUSED && (0 <= mX && mX <= WIDTH) && (0 <= mY && mY <= HEIGHT) {
		cX := mX / CELL_SIZE
		cY := mY / CELL_SIZE
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
			g.grids[0][cY][cX] = true
			// fmt.Printf("%d,%d SET TO TRUE\n", cX, cY)
		}
	}

	if !PAUSED {
		if secondsSince%2 == 0 && !loopRan {
			for y := 0; y < CELL_COUNT_WIDTH; y++ {
				for x := 0; x < CELL_COUNT_WIDTH; x++ {
					count := getNeighborCount(g, x, y)
					if !g.grids[0][y][x] {
						if count == 3 {
							g.grids[1][y][x] = true
						}
					} else {
						if count <= 1 {
							g.grids[1][y][x] = false
						}
						if count >= 4 {
							g.grids[1][y][x] = false
						}
						if count == 2 || count == 3 {
							g.grids[1][y][x] = true
						}
					}
				}
			}
			for y := 0; y < CELL_COUNT_WIDTH; y++ {
				for x := 0; x < CELL_COUNT_WIDTH; x++ {
					g.grids[0][y][x] = g.grids[1][y][x]
					g.grids[1][y][x] = false
				}
			}
			loopRan = true
		}
	}

	if secondsSince%2 != 0 {
		loopRan = false
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.ui.Draw(screen)

	image.Fill(RGBAFromPercent(.9, .9, .9, 1))

	// Draw grid lines
	for i := 0; i <= CELL_COUNT_WIDTH; i++ {
		x := i * CELL_SIZE
		vector.StrokeLine(image, float32(x), 0, float32(x), float32(HEIGHT), float32(1), RGBAFromPercent(0.5, 0.5, 0.5, 1), false)
		vector.StrokeLine(image, 0, float32(x), float32(WIDTH), float32(x), float32(1), RGBAFromPercent(0.5, 0.5, 0.5, 1), false)
	}

	// Draw yellow squares for selected cells
	for y := 0; y < CELL_COUNT_WIDTH; y++ {
		for x := 0; x < CELL_COUNT_WIDTH; x++ {
			if g.grids[0][y][x] {
				// fmt.Printf("%d,%d IS TRUE\n", x, y)
				vector.DrawFilledRect(image, float32(x*CELL_SIZE), float32(y*CELL_SIZE), float32(CELL_SIZE), float32(CELL_SIZE), RGBAFromPercent(0.9, 0.9, 0, 1), false)
			}
		}
	}

	screen.DrawImage(image, &ebiten.DrawImageOptions{})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WIDTH, HEIGHT
}

func main() {
	ebiten.SetWindowSize(WIDTH, HEIGHT)
	ebiten.SetWindowTitle("Game of Life (made by Jaden Hensley)")
	// rand.Seed(time.Now().UnixNano())

	rootContainer := mywidgets.GetContainer()

	ui := ebitenui.UI{
		Container: rootContainer,
	}

	// g := Game{
	// 	ui:   &ui,
	// 	grid: [CELL_COUNT_TOTAL]bool{},
	// }

	var g Game = Game{ui: &ui,
		grids: [2][CELL_COUNT_WIDTH][CELL_COUNT_WIDTH]bool{
			{},
			{},
		},
	}

	for y := 0; y < CELL_COUNT_WIDTH; y++ {
		for x := 0; x < CELL_COUNT_WIDTH; x++ {
			g.grids[0][y][x] = false
			g.grids[0][y][x] = false
		}
	}

	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
