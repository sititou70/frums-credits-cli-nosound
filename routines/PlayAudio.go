package routines

import (
	"io/ioutil"
	"log"
	"sync"
	"time"

	"github.com/hajimehoshi/oto"
	"github.com/sititou70/frums-credits-cli/config"
	"github.com/tosone/minimp3"
)

var PlayState = struct {
	isPlaying     bool
	playStarted   bool
	playFinished  bool
	playStartTime time.Time
}{
	isPlaying:    false,
	playStarted:  false,
	playFinished: false,
}

func PlayAudio(waitGroup *sync.WaitGroup, path string, skipSec int) {
	var err error

	var file []byte
	if file, err = ioutil.ReadFile(path); err != nil {
		log.Fatal(err)
	}

	var dec *minimp3.Decoder
	var data []byte
	if dec, data, err = minimp3.DecodeFull(file); err != nil {
		log.Fatal(err)
	}

	var context *oto.Context
	if context, err = oto.NewContext(dec.SampleRate, dec.Channels, 2, config.AUDIO_PLAY_BUFSIZE); err != nil {
		log.Fatal(err)
	}

	PlayState.isPlaying = true
	PlayState.playStarted = true
	PlayState.playStartTime = time.Now().Add(time.Duration(-skipSec * int(time.Second)))

	var player = context.NewPlayer()
	var sampleParByte = 2
	var skipBufSizeByte = skipSec * dec.SampleRate * sampleParByte * dec.Channels
	player.Write(data[skipBufSizeByte:])

	PlayState.isPlaying = false
	PlayState.playFinished = true

	<-time.After(time.Second)

	dec.Close()
	if err = player.Close(); err != nil {
		log.Fatal(err)
	}

	waitGroup.Done()
}
