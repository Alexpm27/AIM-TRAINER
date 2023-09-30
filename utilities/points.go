package utilities

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"goAIM/models"
	"math/rand"
	"time"
)

func CreatePoint(g *models.Game) *widget.Button {
	btn := widget.NewButton("", func() {
		g.Points++
		g.Score.SetText(fmt.Sprintf("Puntuación: %d", g.Points))
	})
	btn.Resize(fyne.NewSize(20, 20))

	return btn
}

func CreateTarget(g *models.Game) *canvas.Image {
	target := canvas.NewImageFromFile("./images/target.PNG")
	target.Resize(fyne.NewSize(20, 20))

	return target
}

func GenerateButtons(g *models.Game) {
	minX := float32(200)
	maxX := float32(1000)
	minY := float32(100)
	maxY := float32(500)
	for {
		if g.Started {
			button := CreatePoint(g)
			target := CreateTarget(g)
			x := rand.Float32()*(maxX-minX) + minX
			y := rand.Float32()*(maxY-minY) + minY
			button.Move(fyne.NewPos(x, y))
			target.Move(fyne.NewPos(x, y))
			g.ButtonContainer.Add(button)
			g.ButtonContainer.Add(target)
			time.Sleep(g.ButtonInterval)
			button.Hide()
			target.Hide()

			if g.ButtonInterval > 500*time.Millisecond {
				g.ButtonInterval -= 51 * time.Millisecond
			}
		} else {
			time.Sleep(1 * time.Second)
		}
	}
}
