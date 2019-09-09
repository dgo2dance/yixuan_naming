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
)

var (
	commonCharactersL1            map[rune]int32
	commonCharactersL2            map[rune]int32
	commonCharactersL1Traditional map[rune]int32
	commonCharactersL2Traditional map[rune]int32
)

// LoadCommonL1 : Load CommonChineseNamesCharacters.txt into commonL1 map
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
	fullPath = fmt.Sprintf("%s/list/CommonChineseNameCharacters.txt", dir)
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

// AppendCommonL2 : Append rune to common L2 map
func AppendCommonL2(r rune) {
	commonCharactersL2[r]++
}

// CountCommonL1 : Size of common L1
func CountCommonL1() int {
	return len(commonCharactersL1)
}

// CountCommonL2 : Size of common L2
func CountCommonL2() int {
	return len(commonCharactersL2)
}

// TraditionalizeCommonCharacters : Traditionalize common list
func TraditionalizeCommonCharacters() int {
	return 0
}

// GetCommonL1 : Get rune list of L1
func GetCommonL1() map[rune]int32 {
	return commonCharactersL1
}

// GetCommonL2 : Get rune list of L2
func GetCommonL2() map[rune]int32 {
	return commonCharactersL2
}

func init() {
	commonCharactersL2 = make(map[rune]int32)
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
