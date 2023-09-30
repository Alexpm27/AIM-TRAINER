package GUI

import (
	"fmt"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	models "goAIM/models"
)

func InitUI(g *models.Game) {
	g.Score = widget.NewLabel("Puntuación: 0")
	g.TimerLabel = widget.NewLabel("Tiempo restante: 40s")
	g.ProgressBar = widget.NewProgressBar()
	g.TextWidget = canvas.NewText(" ", theme.PrimaryColor())

	g.ContentTop = container.NewVBox(
		canvas.NewText("Haz clic en los targets para ganar puntos en 40 segundos.", theme.PrimaryColor()),
		g.Score,
		g.TimerLabel,
		g.ProgressBar,
		g.TextWidget,
	)

	g.ButtonContainer = container.NewWithoutLayout()

	content := container.NewVBox(g.ContentTop, g.ButtonContainer)
	g.MyWindow.SetContent(content)
}

func ShowScoreDialog(g *models.Game) {
	dialog.ShowInformation("Puntuación", fmt.Sprintf("Tu puntuación es: %d", g.Points), g.MyWindow)
}
