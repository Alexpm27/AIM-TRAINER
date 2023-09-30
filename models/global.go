package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"sync"
	"time"
)

type Game struct {
	MyApp           fyne.App
	MyWindow        fyne.Window
	Score           *widget.Label
	TimerLabel      *widget.Label
	ProgressBar     *widget.ProgressBar
	TextWidget      *canvas.Text
	ContentTop      *fyne.Container
	ButtonContainer *fyne.Container
	mu              sync.Mutex
	Points          int
	Timer           int
	ButtonInterval  time.Duration
	Started         bool
	TimerRunning    bool
}
