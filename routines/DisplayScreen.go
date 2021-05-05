package routines

import (
	"fmt"
	"log"
	"math"
	"os"
	"sync"
	"time"

	"github.com/gdamore/tcell"
	"github.com/sititou70/frums-credits-cli/config"
	CSF "github.com/sititou70/frums-credits-cli/csf"
)

// routine main
func DisplayScreen(waitGroup *sync.WaitGroup, csf *CSF.CSF, verbose bool) {
	// setup tcell screen
	var screen, err = tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}
	if err = screen.Init(); err != nil {
		log.Fatal(err)
	}

	//// quit when press any key
	go func() {
		for {
			ev := screen.PollEvent()
			switch ev.(type) {
			case *tcell.EventKey:
				screen.Fini()
				os.Exit(0)
			}
		}
	}()

	// display loop
	frameRateMS := (time.Duration)(math.Floor(config.FRAME_RATE * 1000 * float64(time.Millisecond)))
	for now := range time.Tick(frameRateMS) {
		var playTime = now.Sub(PlayState.playStartTime)
		var musicPlayTimeSec = playTime.Seconds() - csf.Score.Meta.AudioOffsetSec

		if PlayState.playStarted && musicPlayTimeSec > 0 {
			screen.Clear()

			csf.DrawFrame(screen, musicPlayTimeSec)

			if verbose {
				var bars = musicPlayTimeSec / (60 / (float64)(csf.Score.Meta.BPM) * 4)
				var verboseInfo = fmt.Sprintf(
					"bar / beat: %d / %f\nplaytime: %s",
					int(bars)+1, (bars-math.Floor(bars))*4+1,
					playTime.String())
				CSF.SetContentToScreen(screen, 0, 18, ([]rune)(verboseInfo))
			}

			screen.Show()
		}

		if PlayState.playFinished {
			break
		}
	}

	// finish
	screen.Fini()
	waitGroup.Done()
}
