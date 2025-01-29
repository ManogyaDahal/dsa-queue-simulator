
package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

/* Defining constants */
const (
	screenWidth  = 640 
	screenHeight = 480 
	roadWidth    = 90
	laneWidth    = 30
	lineWidth    = 2
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
  drawGrass(screen)
	drawRoads(screen)
	drawLaneMarkings(screen)
}

/*      Function to Draw grass, Roads and Lane lines      */
func drawGrass(screen *ebiten.Image){
  vector.DrawFilledRect(screen, 0, 0, float32(screenWidth),float32(screenHeight),blue, false)
}

func drawRoads(screen *ebiten.Image) {
  /* Draw Vertical and Horizontal road  respectively */
	vector.DrawFilledRect(screen, float32((screenWidth-roadWidth)/2), 0, float32(roadWidth), float32(screenHeight), gray, false)
	vector.DrawFilledRect(screen, 0, float32((screenHeight-roadWidth)/2), float32(screenWidth), float32(roadWidth), gray, false)
}

func drawLaneMarkings(screen *ebiten.Image) {
	// Vertical lane markings
	for i := 0; i < screenHeight; i += 40 {
		vector.DrawFilledRect(screen, float32((screenWidth-roadWidth)/2+laneWidth), float32(i), float32(lineWidth), 20, white, false)
		vector.DrawFilledRect(screen, float32((screenWidth-roadWidth)/2+2*laneWidth), float32(i), float32(lineWidth), 20, white, false)
	}

	// Horizontal lane markings
	for i := 0; i < screenWidth; i += 40 {
		vector.DrawFilledRect(screen, float32(i), float32((screenHeight-roadWidth)/2+laneWidth), 20, float32(lineWidth), white, false)
		vector.DrawFilledRect(screen, float32(i), float32((screenHeight-roadWidth)/2+2*laneWidth), 20, float32(lineWidth), white, false)
	}
}
/*                Ends              */

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

var (
	gray  = color.RGBA{100, 100, 100, 255} 
	white = color.RGBA{255, 255, 255, 255} 
  blue = color.RGBA{0, 0, 128, 1 }
)

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("dsa-queue-simulator")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
