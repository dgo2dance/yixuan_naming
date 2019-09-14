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
 * @file common_chars.go
 * @package list
 * @author Dr.NP <np@corp.herewetech.com>
 * @since 07/04/2019
 */

package list

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"yixuan_naming/unihan"
)

const (
	// MaxStroke : Max strokes of common characters
	MaxStroke int = 40
)

var (
	commonCharactersL1                  map[rune]int32
	commonCharactersL2                  map[rune]int32
	commonCharactersL1Traditional       map[rune]int32
	commonCharactersL2Traditional       map[rune]int32
	commonCharactersStrokeL1            [][]rune
	commonCharactersStrokeL2            [][]rune
	commonCharactersStrokeL1Traditional [][]rune
	commonCharactersStrokeL2Traditional [][]rune
)

// LoadCommonL1 : Load CommonChineseNamesCharactersL1.txt into commonL1 map
func LoadCommonL1(dir string) (int, error) {
	var (
		fullPath string
		f        *os.File
		err      error
		scanner  *bufio.Scanner
		line     string
		r        rune
		rcode    int64
		total    int
	)

	commonCharactersL1 = make(map[rune]int32)
	fullPath = fmt.Sprintf("%s/list/CommonChineseNameCharactersL1.txt", dir)
	f, err = os.Open(fullPath)
	if err != nil {
		commonCharactersL1 = nil
		return 0, fmt.Errorf("Load list file <%s> failed", fullPath)
	}

	scanner = bufio.NewScanner(f)
	for scanner.Scan() == true {
		line = scanner.Text()
		rcode, err = strconv.ParseInt(line, 10, 64)
		if err == nil {
			r = rune(rcode)
			commonCharactersL1[r]++
			total++
		}
	}

	f.Close()

	return total, nil
}

// LoadCommonL2 : Load CommonChineseNamesCharactersL2.txt into commonL2 map
func LoadCommonL2(dir string) (int, error) {
	var (
		fullPath string
		f        *os.File
		err      error
		scanner  *bufio.Scanner
		line     string
		r        rune
		rcode    int64
		total    int
	)

	commonCharactersL2 = make(map[rune]int32)
	fullPath = fmt.Sprintf("%s/list/CommonChineseNameCharactersL2.txt", dir)
	f, err = os.Open(fullPath)
	if err != nil {
		commonCharactersL1 = nil
		return 0, fmt.Errorf("Load list file <%s> failed", fullPath)
	}

	scanner = bufio.NewScanner(f)
	for scanner.Scan() == true {
		line = scanner.Text()
		rcode, err = strconv.ParseInt(line, 10, 64)
		if err == nil {
			r = rune(rcode)
			commonCharactersL2[r]++
			total++
		}
	}

	f.Close()

	return total, nil
}

// CountCommonL1 : Size of common L1
func CountCommonL1() int {
	return len(commonCharactersL1)
}

// CountCommonL2 : Size of common L2
func CountCommonL2() int {
	return len(commonCharactersL2)
}

// PrepareCommonCharacters : Traditionalize common list and count strokes
func PrepareCommonCharacters() int {
	var (
		r, rt  rune
		ct     int32
		c      *unihan.HanCharacter
		stroke int
		total  int
	)

	commonCharactersL1Traditional = make(map[rune]int32)
	for r, ct = range commonCharactersL1 {
		c, _ = unihan.Query(r)
		if c != nil {
			stroke = c.QueryStrokePrefer()
			if stroke <= MaxStroke {
				commonCharactersStrokeL1[stroke] = append(commonCharactersStrokeL1[stroke], r)
			}

			rt, _ = c.QueryTraditionalLazy()
			if rt > 0 {
				c, _ = unihan.Query(rt)
				if c != nil {
					commonCharactersL1Traditional[rt] = ct
					stroke = c.QueryStrokePrefer()
					if stroke <= MaxStroke {
						commonCharactersStrokeL1Traditional[stroke] = append(commonCharactersStrokeL1Traditional[stroke], rt)
					}

					total++
				}
			}
		}
	}

	commonCharactersL2Traditional = make(map[rune]int32)
	for r, ct = range commonCharactersL2 {
		c, _ = unihan.Query(r)
		if c != nil {
			stroke = c.QueryStrokePrefer()
			if stroke <= MaxStroke {
				commonCharactersStrokeL2[stroke] = append(commonCharactersStrokeL2[stroke], r)
			}

			rt, _ = c.QueryTraditionalLazy()
			if rt > 0 {
				c, _ = unihan.Query(rt)
				if c != nil {
					commonCharactersL2Traditional[rt] = ct
					stroke = c.QueryStrokePrefer()
					if stroke <= MaxStroke {
						commonCharactersStrokeL2Traditional[stroke] = append(commonCharactersStrokeL2Traditional[stroke], rt)
					}

					total++
				}
			}
		}
	}

	return total
}

// GetCommonL1 : Get rune list of L1
func GetCommonL1() map[rune]int32 {
	return commonCharactersL1
}

// GetCommonL2 : Get rune list of L2
func GetCommonL2() map[rune]int32 {
	return commonCharactersL2
}

// GetCommonL1Traditional : Get traditional rune list of L1
func GetCommonL1Traditional() map[rune]int32 {
	return commonCharactersL1Traditional
}

// GetCommonL2Traditional : Get traditional rune list of L2
func GetCommonL2Traditional() map[rune]int32 {
	return commonCharactersL2Traditional
}

// GetCommonL1ByStroke : Get characters by given stroke (L1)
func GetCommonL1ByStroke(stroke int) []rune {
	if stroke < 1 || stroke > MaxStroke {
		return nil
	}

	return commonCharactersStrokeL1[stroke]
}

// GetCommonL2ByStroke : Get characters by given stroke (L2)
func GetCommonL2ByStroke(stroke int) []rune {
	if stroke < 1 || stroke > MaxStroke {
		return nil
	}

	return commonCharactersStrokeL2[stroke]
}

// GetCommonL1ByStrokeTraditional : Get characters by given stroke (L1)
func GetCommonL1ByStrokeTraditional(stroke int) []rune {
	if stroke < 1 || stroke > MaxStroke {
		return nil
	}

	return commonCharactersStrokeL1Traditional[stroke]
}

// GetCommonL2ByStrokeTraditional : Get characters by given stroke (L2)
func GetCommonL2ByStrokeTraditional(stroke int) []rune {
	if stroke < 1 || stroke > MaxStroke {
		return nil
	}

	return commonCharactersStrokeL2Traditional[stroke]
}

func init() {
	commonCharactersStrokeL1 = make([][]rune, MaxStroke+1)
	commonCharactersStrokeL1Traditional = make([][]rune, MaxStroke+1)
	commonCharactersStrokeL2 = make([][]rune, MaxStroke+1)
	commonCharactersStrokeL2Traditional = make([][]rune, MaxStroke+1)

	return
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
