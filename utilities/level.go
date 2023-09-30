package utilities

import (
	"goAIM/models"
	"time"
)

func Level(g *models.Game) {
	for {
		if g.Started {
			if g.Points < 9 {
				g.TextWidget.Text = "Principiante"
			} else if g.Points < 15 {
				g.TextWidget.Text = "Promedio"
			} else if g.Points < 20 {
				g.TextWidget.Text = "Avanzado"
			} else {
				g.TextWidget.Text = "Veterano"
			}

			g.TextWidget.Refresh()
			time.Sleep(1 * time.Second)
		} else {
			time.Sleep(1 * time.Second)
		}
	}
}
