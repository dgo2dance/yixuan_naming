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
 * @file baijiaxing.go
 * @package list
 * @author Dr.NP <np@corp.herewetech.com>
 * @since 07/15/2019
 */

package list

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// BaiJiaXing : BaiJiaXing structure
type BaiJiaXing struct {
	Sort       int    `json:"sort"`
	FamilyName string `json:"family_name"`
	Place      string `json:"place"`
}

var baiJiaXingMSimplified map[string]*BaiJiaXing
var baiJiaXingMTraditional map[string]*BaiJiaXing

// QueryBaiJiaXing : Check and query baijiaxing
func QueryBaiJiaXing(familyName string) *BaiJiaXing {
	if baiJiaXingMSimplified != nil {
		return baiJiaXingMSimplified[familyName]
	}

	return nil
}

// LoadBaiJiaXing : Load BaiJiaXing from list
func LoadBaiJiaXing(dir string) (int, error) {
	var (
		fullPath string
		f        *os.File
		err      error
		scanner  *bufio.Scanner
		line     string
		parts    []string
		sort     int
		total    int
	)

	baiJiaXingMSimplified = make(map[string]*BaiJiaXing)
	baiJiaXingMTraditional = make(map[string]*BaiJiaXing)
	fullPath = fmt.Sprintf("%s/list/BaiJiaXing.txt", dir)
	f, err = os.Open(fullPath)
	if err != nil {
		baiJiaXingMSimplified = nil
		baiJiaXingMTraditional = nil
		return 0, fmt.Errorf("Load list file <%s> failed", fullPath)
	}

	scanner = bufio.NewScanner(f)
	for scanner.Scan() == true {
		line = scanner.Text()
		parts = strings.Split(line, ",")
		if 3 == len(parts) {
			sort, err = strconv.Atoi(parts[0])
			if err == nil {
				baiJiaXingMSimplified[parts[1]] = &BaiJiaXing{
					Sort:       sort,
					FamilyName: parts[1],
					Place:      parts[2],
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
