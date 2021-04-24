package csf

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/wesovilabs/koazee"
	"github.com/yargevad/filepathx"
	yaml "gopkg.in/yaml.v2"
)

type ContentsMap map[string]string

// parsers
func parseScoreDir(dirPath string) Score {
	var score Score

	// meta
	bytes, err := ioutil.ReadFile(filepath.Join(dirPath, "./meta.yaml"))
	if err != nil {
		panic(err)
	}
	score.Meta = parseMetaInfo(bytes)

	// load data items
	contentsMap := make(ContentsMap)
	dataFilePaths, err := filepathx.Glob(filepath.Join(dirPath, "./data/**/*"))
	for _, path := range dataFilePaths {
		file, err := os.Stat(path)
		if err != nil {
			panic(err)
		}
		if file.IsDir() {
			continue
		}

		bytes, err := ioutil.ReadFile(path)

		var removePath = filepath.Join(dirPath, "./data") + "/"
		var key = strings.Replace(path, removePath, "", -1)
		contentsMap[key] = string(bytes)
	}

	// parse part file
	partFilePaths, err := filepathx.Glob(filepath.Join(dirPath, "./scores/**/*.part"))
	if err != nil {
		panic(err)
	}

	score.Parts = make([]Part, 0)
	for _, path := range partFilePaths {
		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}

		var part = parsePartFile(bytes, &contentsMap)
		score.Parts = append(score.Parts, part)
	}
	sort.Slice(score.Parts, func(i, j int) bool {
		return score.Parts[i].Zindex < score.Parts[j].Zindex
	})

	return score
}

func parseMetaInfo(metaFileContent []byte) MetaInfo {
	var meta MetaInfo
	err := yaml.UnmarshalStrict(metaFileContent, &meta)
	if err != nil {
		panic(fmt.Sprint("invalid meta file: ", err))
	}

	return meta
}

func parsePartFile(partFileContent []byte, contentsMap *ContentsMap) Part {
	var part Part = Part{Bars: make([]Bar, 0), Zindex: 0}
	var currentPosition DisplayPosition = DisplayPosition{X: 0, Y: 0}
	var flipMode = struct{ vertical bool }{false}

	var bars = strings.Split(string(partFileContent), "---")
	for _, barContent := range bars {
		part.Bars = append(part.Bars, Bar{Items: make([]DisplayItem, 0)})
		var lines = strings.Split(barContent, "\n")
		for _, line := range lines {
			if len(line) == 0 {
				continue
			}

			// interpret line
			var contentKey string
			switch line[0] {
			case '"':
				contentKey = line[1 : len(line)-1]
			case '#':
				var args = strings.Split(line[2:], " ")
				switch args[0] {
				case "MOVETO":
					if len(args) < 3 {
						panic("partfile parse error: MOVETO: a few arguments")
					}
					x, errx := strconv.Atoi(args[1])
					y, erry := strconv.Atoi(args[2])
					if errx != nil || erry != nil {
						panic("partfile parse error: MOVETO: invalid syntax")
					}
					currentPosition.X = x
					currentPosition.Y = y
				case "ZINDEX":
					if len(args) < 2 {
						panic("partfile parse error: ZINDEX: a few arguments")
					}
					zIndex, err := strconv.Atoi(args[1])
					if err != nil {
						panic("partfile parse error: ZINDEX: invalid syntax")
					}
					part.Zindex = zIndex
				case "FLIP":
					if len(args) < 3 {
						panic("partfile parse error: FLIP: a few arguments")
					}
					switch args[1] {
					case "vertical":
						flipMode.vertical = args[2] == "on"
					}
				}
				continue
			case '/':
				continue
			default:
				contentKey = line
			}

			// create DisplayItem
			//// create original content
			if _, hasContent := (*contentsMap)[contentKey]; !hasContent {
				(*contentsMap)[contentKey] = contentKey
			}
			//// create flipped content
			if flipMode.vertical {
				var originalContent = (*contentsMap)[contentKey]
				contentKey = contentKey + ":vertical-flipped"
				if _, hasContent := (*contentsMap)[contentKey]; !hasContent {
					(*contentsMap)[contentKey] = flipContentVertical(originalContent)
				}
			}

			//// create item
			var currentItems *[]DisplayItem = &part.Bars[len(part.Bars)-1].Items
			var content = (*contentsMap)[contentKey]
			*currentItems = append(*currentItems, DisplayItem{Position: currentPosition, Content: &content})
		}
	}

	return part
}

// utils
func flipContentVertical(content string) string {
	var lines = strings.Split(content, "\n")
	var flipedLines = koazee.StreamOf(lines).
		Map(func(x string) string { return reverseString(x) }).Out().Val().([]string)
	return strings.Join(flipedLines, "\n")
}

// from: https://kodify.net/go/reverse-string/
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
