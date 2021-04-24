package csf

type DisplayPosition struct {
	X int
	Y int
}

type DisplayItem struct {
	Position DisplayPosition
	Content  *string
}

type Bar struct {
	Items []DisplayItem
}

type Part struct {
	Bars []Bar
	Zindex int
}

type MetaInfo struct {
	BPM            int     `yaml:"BPM"`
	AudioFilePath  string  `yaml:"AudioFilePath"`
	AudioOffsetSec float64 `yaml:"AudioOffsetSec"`
}

type Score struct {
	Parts []Part
	Meta  MetaInfo
}
