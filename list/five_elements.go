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
 * @file five_elements.go
 * @package list
 * @author Dr.NP <np@corp.herewetech.com>
 * @since 07/10/2019
 */

package list

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"yixuan_naming/utils"
)

var feM map[rune]int

// QueryFiveElement : Check and return five element by given rune
func QueryFiveElement(r rune) int {
	if feM == nil {
		return utils.ElementUnknown
	}

	v, exists := feM[r]
	if !exists {
		return utils.ElementUnknown
	}

	return v
}

// LoadCharacterFiveElements : Load five elements list of Chinese characters
func LoadCharacterFiveElements(dir string) (int, error) {
	var (
		fullPath string
		f        *os.File
		err      error
		scanner  *bufio.Scanner
		line     string
		parts    []string
		r        rune
		rcode    int
		fe       int
		total    int
	)

	feM = make(map[rune]int)
	fullPath = fmt.Sprintf("%s/list/CharacterFiveElements.txt", dir)
	f, err = os.Open(fullPath)
	if err != nil {
		feM = nil
		return 0, fmt.Errorf("Load five elements list file <%s> failed", fullPath)
	}

	scanner = bufio.NewScanner(f)
	for scanner.Scan() == true {
		line = scanner.Text()
		parts = strings.Split(line, ",")
		if 2 == len(parts) {
			rcode, _ = strconv.Atoi(parts[0])
			fe, _ = strconv.Atoi(parts[1])
			if rcode > 0 && fe < 5 && fe >= 0 {
				r = rune(rcode)
				feM[r] = fe
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
