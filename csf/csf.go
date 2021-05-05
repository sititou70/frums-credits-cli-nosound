package csf

import (
	"io/fs"
	"math"

	"github.com/gdamore/tcell"
)

type CSF struct {
	Score Score
}

// main
func NewCSF(rootFS fs.FS) CSF {
	var csf CSF
	csf.Score = parseScoreDir(rootFS)
	return csf
}

func (csf CSF) DrawFrame(screen tcell.Screen, musicPlayTimeSec float64) {
	var currentPlayedBars = musicPlayTimeSec / (60 / (float64)(csf.Score.Meta.BPM) * 4)
	for _, part := range csf.Score.Parts {
		var barIndex = int(currentPlayedBars)
		if len(part.Bars) <= barIndex {
			continue
		}

		var bar = part.Bars[barIndex]
		if len(bar.Items) == 0 {
			continue
		}

		var itemIndex = int((float64)(len(bar.Items)) * (currentPlayedBars - math.Floor(currentPlayedBars)))
		var item = bar.Items[itemIndex]
		SetContentToScreen(screen, item.Position.X, item.Position.Y, []rune(*item.Content))
	}
}

// utils
func SetContentToScreen(screen tcell.Screen, x int, y int, data []rune) {
	var dataLength = len(data)
	var currentPosition DisplayPosition = DisplayPosition{0, 0}
	for charIndex := 0; charIndex < dataLength; charIndex++ {
		if data[charIndex] == '\n' {
			currentPosition.Y++
			currentPosition.X = 0
			continue
		}
		if data[charIndex] != ' ' {
			screen.SetContent(x+currentPosition.X, y+currentPosition.Y, data[charIndex], nil, tcell.StyleDefault)
		}
		currentPosition.X++
	}
}
