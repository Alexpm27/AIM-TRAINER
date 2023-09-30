package game

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"goAIM/GUI"
	"goAIM/models"
	"goAIM/utilities"
)

func NewGame() *models.Game {
	game := &models.Game{}
	game.MyApp = app.New()
	game.MyWindow = game.MyApp.NewWindow("AIM TRAINER")
	game.ContentTop = container.NewVBox()
	game.ButtonContainer = container.NewWithoutLayout()
	GUI.InitUI(game)
	initEventHandlers(game)
	go utilities.GenerateButtons(game)
	go utilities.Level(game)
	return game
}

func initEventHandlers(g *models.Game) {
	g.ContentTop.Add(widget.NewButton("Comenzar", func() {
		StartGame(g)
	}))
}

func StartGame(g *models.Game) {
	if !g.TimerRunning {
		g.Started = true
		g.Points = 0
		g.Timer = 40
		g.ButtonInterval = 2 * time.Second
		g.MyWindow.Canvas().Refresh(g.MyWindow.Content())

		g.TimerRunning = true
		go utilities.RunTimer(g)
	}
}
