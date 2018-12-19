package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"golang.org/x/image/colornames"
)

type sprite struct {
	image *ebiten.Image
	x     float64
	y     float64
}

var sp *sprite

func init() {
	img, _ := ebiten.NewImage(60, 60, ebiten.FilterNearest)
	img.Fill(colornames.Blue)
	sp = &sprite{
		image: img,
		x:     100,
		y:     100,
	}
}

func (s *sprite) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(s.x, s.y)
	screen.DrawImage(s.image, op)
}

func getTouchPosition() (int, int) {
	var x, y int
	ts := ebiten.TouchIDs()
	if len(ts) == 1 {
		x, y = ebiten.TouchPosition(ts[0])
	}

	return x, y
}

func (s *sprite) IsJustPressed() bool {
	rect := s.image.Bounds()
	// x, y := getTouchPosition()
	cursorX, cursorY := ebiten.CursorPosition()

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if cursorX >= int(s.x) && cursorX <= int(s.x)+rect.Dx() && cursorY >= int(s.y) && cursorY <= int(s.y)+rect.Dy() {
		// if float64(cursorX) >= s.x && float64(cursorX) <= s.x+float64(rect.Dx()) && float64(cursorY) >= s.y && float64(cursorY) <= s.y+float64(rect.Dy()) {
			return true
		}
	}

	touchX, touchY := getTouchPosition()
	if touchX >= int(s.x) && touchX <= int(s.x)+rect.Dx() && touchY >= int(s.y) && touchY <= int(s.y)+rect.Dy() {
	// if float64(touchX) >= s.x && float64(touchX) <= s.x+float64(rect.Dx()) && float64(touchY) >= s.y && float64(touchY) <= s.y+float64(rect.Dy()) {
		return true
	}

	return false
}

func (s *sprite) IsJustReleased() bool {
	return true
}

func (s *sprite) IsPressed() bool {
	return true
}

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	
	screen.Fill(colornames.Black)
	
	ebitenutil.DebugPrint(screen, strconv.Itoa(int(ebiten.CurrentFPS())))

	sp.draw(screen)

	if sp.IsJustPressed() {
		sp.image.Fill(colornames.Red)
	}

	ebitenutil.DebugPrint(screen, fmt.Sprint(sp.IsJustPressed()))

	return nil
}

func main() {
	if err := ebiten.Run(update, 380, 640, 1, "スプライト"); err != nil {
		log.Fatal(err)
	}
}
