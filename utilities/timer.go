package utilities

import (
	"fmt"
	"goAIM/GUI"
	"goAIM/models"
	"time"
)

func RunTimer(g *models.Game) {
	for g.Timer >= 0 {
		if !g.Started {
			break
		}
		g.TimerLabel.SetText(fmt.Sprintf("Tiempo restante: %ds", g.Timer))
		g.ProgressBar.SetValue(float64(g.Timer) / 40.0)
		time.Sleep(1 * time.Second)
		g.Timer--
	}
	g.Started = false
	g.TimerRunning = false
	g.MyWindow.Canvas().Refresh(g.MyWindow.Content())
	GUI.ShowScoreDialog(g)
}
