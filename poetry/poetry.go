/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2019 HereweTech Co.LTD
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of
 * this software and associated documentation files (the "Software"), to deal in
 * the Software without restriction, including without limitation the rights to
 * use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
 * the Software, and to permit persons to whom the Software is furnished to do so,
 * subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
 * FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
 * COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
 * IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
 * CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

/**
 * @file poetry.go
 * @package poetry
 * @author Dr.NP <np@corp.herewetech.com>
 * @since 07/21/2019
 */

package poetry

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Poet types
const (
	// PoetShijing : Shijing
	PoetShijing = iota
	// PoetChuci : Chuci
	PoetChuci
	// PoetFourbooks : Fourbooks
	PoetFourbooks
	// PoetZhouyi : Zhouyi
	PoetZhouyi
	// PoetFlower : HuaJianJi
	PoetFlower
	// PoetSouthTang : NanTangHouZhu
	PoetSouthTang
	// PoetTang : TangShi
	PoetTang
	// PoetSong : SongShi
	PoetSong
	// PoetSongci : SongCi
	PoetSongci
)

// Poetry : Struct of poetries
type Poetry struct {
	Type       int    `json:"type"`
	Author     string `json:"author"`
	Title      string `json:"title"`
	Paragraphs string `json:"paragraphs"`
}

// word : Words in poetries
type word struct {
	word     string
	poetries []int
}

var (
	poetriesASimplified     []*Poetry
	poetriesATraditional    []*Poetry
	poetryWordsMSimplified  map[string][]int
	poetryWordsMTraditional map[string][]int
)

// QueryPoetries : Check and query poetries by given word
func QueryPoetries(word string) []*Poetry {
	var ret []*Poetry
	tags, ok := poetryWordsMSimplified[word]
	if ok {
		for _, tag := range tags {
			if tag >= 0 && tag < len(poetryWordsMSimplified) {
				ret = append(ret, poetriesASimplified[tag])
			}
		}
	}

	return ret
}

// InPoetries : Whether word in poetries
func InPoetries(word string) bool {
	tags, ok := poetryWordsMSimplified[word]
	if ok {
		if len(tags) > 0 {
			return true
		}
	}

	return false
}

// LoadPoetries : Load poetry from list
func LoadPoetries(dir string) (int, int, error) {
	var (
		fullPath      string
		f             *os.File
		scanner       *bufio.Scanner
		line          string
		parts         []string
		tag           string
		tags          []string
		tagN          int
		pType         int
		totalWords    int
		totalPoetries int
		err           error
	)

	poetryWordsMSimplified = make(map[string][]int)
	poetryWordsMTraditional = make(map[string][]int)

	fullPath = fmt.Sprintf("%s/poetry/PoetriesS.txt", dir)
	f, err = os.Open(fullPath)
	if err != nil {
		poetriesASimplified = nil
		poetriesATraditional = nil
		poetryWordsMSimplified = nil
		poetryWordsMTraditional = nil

		return 0, 0, fmt.Errorf("Load poetries from <%s> failed", fullPath)
	}

	scanner = bufio.NewScanner(f)
	for scanner.Scan() == true {
		line = scanner.Text()
		parts = strings.Split(line, "$")
		if len(parts) == 4 {
			switch parts[0] {
			case "ShiJing":
				pType = PoetShijing
			case "ChuCi":
				pType = PoetChuci
			case "4.Books":
				pType = PoetFourbooks
			case "ZhouYi":
				pType = PoetZhouyi
			case "Flower":
				pType = PoetFlower
			case "SouthTang.Poet":
				pType = PoetSouthTang
			case "Tang.Poet":
				pType = PoetTang
			case "Song.Poet":
				pType = PoetSong
			case "Song.Ci":
				pType = PoetSongci
			default:
				pType = -1
			}

			poetriesASimplified = append(poetriesASimplified, &Poetry{
				Type:       pType,
				Author:     parts[1],
				Title:      parts[2],
				Paragraphs: parts[3],
			})
			totalPoetries++
		}
	}

	f.Close()

	fullPath = fmt.Sprintf("%s/poetry/PoetryWordsS.txt", dir)
	f, err = os.Open(fullPath)
	if err != nil {
		poetriesASimplified = nil
		poetriesATraditional = nil
		poetryWordsMSimplified = nil
		poetryWordsMTraditional = nil

		return 0, 0, fmt.Errorf("Load poetry words from <%s> failed", fullPath)
	}

	scanner = bufio.NewScanner(f)
	for scanner.Scan() == true {
		line = scanner.Text()
		parts = strings.Split(line, "$")
		if len(parts) == 2 {
			tags = strings.Split(parts[1], "|")
			for _, tag = range tags {
				tagN, err = strconv.Atoi(tag)
				if err == nil {
					poetryWordsMSimplified[parts[0]] = append(poetryWordsMSimplified[parts[0]], tagN)
				}
			}

			totalWords++
		}
	}

	f.Close()

	return totalPoetries, totalWords, nil
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
