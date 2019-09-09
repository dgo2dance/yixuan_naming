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
 * @file folkways.go
 * @package dict
 * @author Dr.NP <np@corp.herewetech.com>
 * @since 08/03/2019
 */

package dict

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// FolkwayItem : Item of TW dictionary
type FolkwayItem struct {
	Unicode     rune   `json:"unicode"`
	Utf8Str     string `json:"utf8_str"`
	Explanation string `json:"explanation,omitempty"`
}

var twM map[rune]*FolkwayItem

// QueryFolkway : Query TW dictionary
func QueryFolkway(r rune) *FolkwayItem {
	if twM != nil {
		return twM[r]
	}

	return nil
}

// LoadFolkways : Load TW dictionary from origin file
func LoadFolkways(dir string) (int, error) {
	var (
		fullPath string
		f        *os.File
		scanner  *bufio.Scanner
		line     string
		parts    []string
		total    int
		err      error
		r        rune
		rcode    int
	)

	twM = make(map[rune]*FolkwayItem)
	fullPath = fmt.Sprintf("%s/dict/Folkways.dict", dir)
	f, err = os.Open(fullPath)
	if err != nil {
		twM = nil
		return 0, fmt.Errorf("Load dictionary file <%s> failed", fullPath)
	}

	scanner = bufio.NewScanner(f)
	for scanner.Scan() == true {
		line = scanner.Text()
		parts = strings.Split(line, "||")
		if 3 == len(parts) {
			rcode, err = strconv.Atoi(parts[0])
			if err == nil {
				r = rune(rcode)
				if twM[r] == nil {
					twM[r] = &FolkwayItem{
						Unicode:     r,
						Utf8Str:     parts[1],
						Explanation: parts[2],
					}
				}

				total++
			}
		}
	}

	f.Close()

	return total, nil
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
