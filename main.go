package main

import (
	"embed"
	"flag"
	"io/fs"
	"sync"

	"github.com/sititou70/frums-credits-cli/csf"
	"github.com/sititou70/frums-credits-cli/routines"
)

var (
	skipSec     = flag.Int("s", 0, "time to skip play (sec)")
	verboseMode = flag.Bool("v", false, "print extra information")
)

//go:embed assets/csf-root
var csfFS embed.FS
var csfRootDirPath = "assets/csf-root"

func main() {
	// parse options
	flag.Parse()

	// parse CSF(credits score format)
	csfRootFS, _ := fs.Sub(csfFS, csfRootDirPath)
	var csf = csf.NewCSF(csfRootFS)

	// launch routines
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	go routines.PlayAudio(&waitGroup, csfRootFS, csf.Score.Meta.AudioFilePath, *skipSec)
	go routines.DisplayScreen(&waitGroup, &csf, *verboseMode)
	waitGroup.Wait()
}
