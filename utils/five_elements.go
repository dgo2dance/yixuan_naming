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
 * @package utils
 * @author Dr.NP <np@corp.herewetech.com>
 * @since 06/27/2019
 */

package utils

// Five elements
const (
	// ElementWood : Mu
	ElementWood = iota
	// ElementFire : Huo
	ElementFire
	// ElementEarth : Tu
	ElementEarth
	// ElementMetal : Jin
	ElementMetal
	// ElementWater : Shui
	ElementWater

	// ElementUnknown : Nothing here
	ElementUnknown = -1
)

// FiveElementCompare
const (
	// FiveElementEqual : 0
	FiveElementEqual = 0

	// FiveElementKill : 1
	FiveElementKill = 1

	// FiveElementKilled : 2
	FiveElementKilled = 2

	// FiveElementBirth : 3
	FiveElementBirth = 3

	// FiveElementBirthed : 4
	FiveElementBirthed = 4

	// FiveElementCompareInvalid : -1
	FiveElementCompareInvalid = -1
)

// FiveElementsCount : Count stat of five elements
type FiveElementsCount struct {
	Wood  int `json:"wood"`
	Fire  int `json:"fire"`
	Earth int `json:"earth"`
	Metal int `json:"metal"`
	Water int `json:"water"`
}

// CompareFiveElements : Compare five-elements - Wood->Fire->Earth->Metal->Water
func CompareFiveElements(fe1, fe2 int) int {
	if fe1 < fe2 {
		fe1 += 5
	}

	switch fe1 - fe2 {
	case 0:
		// Equal
		return FiveElementEqual
	case 1:
		// Birthed
		return FiveElementBirthed
	case 2:
		// Killed
		return FiveElementKilled
	case 3:
		// Kill
		return FiveElementKill
	case 4:
		// Birth
		return FiveElementBirth
	default:
		// Invalid
		return FiveElementCompareInvalid
	}
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
