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
 * @file animal.go
 * @package name
 * @author Dr.NP <np@corp.herewetech.com>
 * @since 07/03/2019
 */

package name

import (
	"yixuan_naming/texts"
	"yixuan_naming/utils"
)

type radicalsMeaning struct {
	RadicalIndexes []int                   `json:"radical_indexes"`
	Radicals       []*utils.ChineseRadical `json:"radicals"`
	Meaning        string                  `json:"meaning"`
}

type animalRadicals struct {
	Lucky   []radicalsMeaning `json:"lucky"`
	Ominous []radicalsMeaning `json:"ominous"`
}

type animalRadicalsTpl struct {
	lucky   [][]int
	ominous [][]int
}

type animal struct {
	Radicals *animalRadicals `json:"radicals"`
	Years    string          `json:"years"`
}

var (
	animalRadicalsTplList = []animalRadicalsTpl{
		// 0
		{
			lucky: [][]int{
				{12, 40},
				{119, 151, 195},
				{140, 167, 96},
				{9, 75, 74},
				{102},
			},
			ominous: [][]int{
				{46},
				{18, 19, 57},
				{32},
				{61},
				{112},
				{107, 85, 187, 164, 86},
			},
		},
		// 1
		{
			lucky: [][]int{
				{85},
				{9, 75},
			},
			ominous: [][]int{
				{74},
				{86},
				{102, 159, 187},
				{112, 46},
				{143, 120, 18, 19, 16},
			},
		},
		// 2
		{
			lucky: [][]int{
				{46},
				{96},
				{167, 75, 145, 85},
				{74, 94, 187},
			},
			ominous: [][]int{
				{72, 86},
				{102, 30, 10},
				{120, 112, 18, 19, 143, 57, 88, 157},
			},
		},
		// 3
		{
			lucky: [][]int{
				{74},
				{9, 115, 75},
				{11, 40},
				{167, 106, 96, 151},
				{94},
			},
			ominous: [][]int{
				{187, 164},
				{112, 19, 18},
				{107, 85},
				{85, 47},
			},
		},
		// 4
		{
			lucky: [][]int{
				{85},
				{167, 96, 106, 155},
				{74},
				{195, 164, 9},
			},
			ominous: [][]int{
				{32, 102, 115, 145},
				{32, 61, 72},
				{112, 140},
				{19, 18},
				{120, 94},
				{86},
			},
		},
		// 5
		{
			lucky: [][]int{
				{140},
				{142, 195},
				{75, 115, 102, 46},
				{167, 96},
				{74, 32},
			},
			ominous: [][]int{
				{61},
				{112, 18, 143, 57},
				{86, 9, 120},
			},
		},
		// 6
		{
			lucky: [][]int{
				{140, 167},
				{96, 75, 115},
				{142, 151, 119},
				{9, 74},
				{32},
			},
			ominous: [][]int{
				{102, 86, 85},
				{159, 112, 19, 164, 187},
			},
		},
		// 7
		{
			lucky: [][]int{
				{167, 106, 96, 140},
				{74, 102, 151, 119},
				{187, 115, 75, 9},
			},
			ominous: [][]int{
				{61, 94, 120},
				{159, 85, 46, 72, 86},
			},
		},
		// 8
		{
			lucky: [][]int{
				{75, 115},
				{167, 96, 151, 119},
				{102, 46, 74},
				{85, 9},
				{46},
			},
			ominous: [][]int{
				{86, 112},
				{30, 14},
				{120, 18, 19, 107, 94},
			},
		},
		// 9
		{
			lucky: [][]int{
				{119, 151, 142},
				{75, 115, 96, 102},
				{74, 9, 14},
				{46, 140, 72, 167},
			},
			ominous: [][]int{
				{112, 94, 18, 19, 164, 143, 57, 120, 159, 187},
			},
		},
		// 10
		{
			lucky: [][]int{
				{195, 151, 119},
				{9, 14, 187},
				{167, 96, 140, 102, 75, 115, 74},
				{85},
				{9},
			},
			ominous: [][]int{
				{86},
				{112, 120, 46, 72},
				{164, 159, 18, 88, 149},
			},
		},
		// 11
		{
			lucky: [][]int{
				{151, 119, 195},
				{85, 167, 96},
				{74, 75, 115},
				{9, 46, 32, 140},
			},
			ominous: [][]int{
				{120, 112, 18, 19, 143, 57, 10, 107, 88},
			},
		},
	}
)

func getAnimalRadicals(index int, language int) *animalRadicals {
	var (
		ret = &animalRadicals{}
		ii  int
		rs  []*utils.ChineseRadical
	)

	if index < 0 || index > 11 {
		return nil
	}

	for before := 0; before < index; before++ {
		ii += len(animalRadicalsTplList[before].lucky) + len(animalRadicalsTplList[before].ominous)
	}

	for _, r := range animalRadicalsTplList[index].lucky {
		rs = nil
		for _, rc := range r {
			rs = append(rs, utils.GetRadical(rc))
		}

		ret.Lucky = append(ret.Lucky, radicalsMeaning{
			RadicalIndexes: r,
			Radicals:       rs,
			Meaning:        texts.GetMessage(texts.MessageAnimalRadicalsDescription, ii, language),
		})

		ii++
	}

	for _, r := range animalRadicalsTplList[index].ominous {
		rs = nil
		for _, rc := range r {
			rs = append(rs, utils.GetRadical(rc))
		}

		ret.Ominous = append(ret.Ominous, radicalsMeaning{
			RadicalIndexes: r,
			Radicals:       rs,
			Meaning:        texts.GetMessage(texts.MessageAnimalRadicalsDescription, ii, language),
		})

		ii++
	}

	return ret
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
