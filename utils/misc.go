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
 * @file misc.go
 * @package utils
 * @author Dr.NP <np@corp.herewetech.com>
 * @since 06/27/2019
 */

package utils

import "bytes"

// Location : Geographic location
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

const (
	// GenderMale : Male
	GenderMale = 1

	// GenderFemale : Female
	GenderFemale = 2

	// GenderUnknown : Unkonwn
	GenderUnknown = 0
)

// RunesToStr : Convert rune slice to string
func RunesToStr(runes []rune) string {
	var (
		b bytes.Buffer
		r rune
	)

	for _, r = range runes {
		b.WriteRune(r)
	}

	return b.String()
}

// StrToRunes : Convert string to rune slice
func StrToRunes(str string) []rune {
	return []rune(str)
}

// FixMod : Fix mod
func FixMod(i, m int) int {
	r := i % m
	if r == 0 {
		r = m
	}

	return r
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
