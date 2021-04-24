package main

import (
	"flag"
	"path/filepath"
	"sync"

	"github.com/sititou70/frums-credits-cli/csf"
	"github.com/sititou70/frums-credits-cli/routines"
)

var (
	scoreDir    = flag.String("i", "", "path to CSF(credits score format) dir")
	skipSec     = flag.Int("s", 0, "time to skip play (sec)")
	verboseMode = flag.Bool("v", false, "print extra information")
)

func main() {
	// parse options
	flag.Parse()
	if len(*scoreDir) == 0 {
		panic("path to score dir is required")
	}

	// parse CSF(credits score format)
	var csf = csf.NewCSF(*scoreDir)

	// launch routines
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	go routines.PlayAudio(&waitGroup, filepath.Join(*scoreDir, csf.Score.Meta.AudioFilePath), *skipSec)
	go routines.DisplayScreen(&waitGroup, &csf, *verboseMode)
	waitGroup.Wait()
}
