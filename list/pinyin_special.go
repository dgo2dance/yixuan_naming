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
 * @file pinyin_special.go
 * @package list
 * @author Dr.NP <np@corp.herewetech.com>
 * @since 07/25/2019
 */

package list

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var pinyinSpecialM map[rune]string

// QueryPinyinSpecial : Check and query pinyin speicial value by given rune
func QueryPinyinSpecial(input rune) string {
	if pinyinSpecialM != nil {
		return pinyinSpecialM[input]
	}

	return "_"
}

// LoadPinyinSpecial : Special pinyin values
func LoadPinyinSpecial(dir string) (int, error) {
	var (
		fullPath string
		f        *os.File
		scanner  *bufio.Scanner
		line     string
		parts    []string
		key      rune
		keyv     int
		total    int
		err      error
	)

	pinyinSpecialM = make(map[rune]string)
	fullPath = fmt.Sprintf("%s/list/PinyinSpecial.txt", dir)
	f, err = os.Open(fullPath)
	if err != nil {
		pinyinSpecialM = nil
		return 0, fmt.Errorf("Load pinyin special list <%s> failed", fullPath)
	}

	scanner = bufio.NewScanner(f)
	for scanner.Scan() == true {
		line = scanner.Text()
		parts = strings.Split(line, ":")
		if 2 == len(parts) {
			keyv, err = strconv.Atoi(parts[0])
			if err == nil {
				key = rune(keyv)
				pinyinSpecialM[key] = parts[1]
				total++
			}
		}
	}

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
