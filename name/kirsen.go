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
 * @file kirsen.go
 * @package name
 * @author Dr.NP <np@corp.herewetech.com>
 * @since 07/04/2019
 */

package name

import (
	"fmt"
	"math"
	"yixuan_naming/texts"
	"yixuan_naming/utils"

	"yixuan_naming/calendar"
	"yixuan_naming/list"
	"yixuan_naming/unihan"
)

const (
	// MaxRank : Maxinum rank score of name
	MaxRank = 100

	// MaxNames : Maxinum names return by kirsen
	MaxNames = 10000
)

var (
	strokeTable [][][]int
)

// KirsenConditions : Conditions of name list generation
type KirsenConditions struct {
	FamilyNameRunes []rune
	MiddleNameRunes []rune
	PrefixNameRunes []rune
	SuffixNameRunes []rune
	Gender          int
	NeedMiddleName  bool
	NeedBirthTime   bool
	GivenNameLength int
	QueryNums       int
	CharacterLevel  int
}

// Traditionalize : Traditionalize conditions
func (c *KirsenConditions) Traditionalize() {
	var (
		rt  rune
		u   *unihan.HanCharacter
		err error
	)

	for i, r := range c.FamilyNameRunes {
		u, err = unihan.Query(r)
		if err == nil {
			rt, err = u.QueryTraditionalLazy()
			if err == nil {
				c.FamilyNameRunes[i] = rt
			}
		}
	}

	for i, r := range c.MiddleNameRunes {
		u, err = unihan.Query(r)
		if err == nil {
			rt, err = u.QueryTraditionalLazy()
			if err == nil {
				c.MiddleNameRunes[i] = rt
			}
		}
	}

	for i, r := range c.PrefixNameRunes {
		u, err = unihan.Query(r)
		if err == nil {
			rt, err = u.QueryTraditionalLazy()
			if err == nil {
				c.PrefixNameRunes[i] = rt
			}
		}
	}

	for i, r := range c.SuffixNameRunes {
		u, err = unihan.Query(r)
		if err == nil {
			rt, err = u.QueryTraditionalLazy()
			if err == nil {
				c.SuffixNameRunes[i] = rt
			}
		}
	}

	return
}

// KirsenData : Kirsen result
type KirsenData struct {
	language           int
	List               []*Name                `json:"list"`
	Calendar           *calendar.Calendar     `json:"calendar"`
	GanzhiFiveElements GanzhiFiveElementsSpec `json:"ganzhi_five_elements"`
	SoundFiveElements  SoundFiveElements      `json:"sound_five_elements"`
	EightCharacters    eightCharacters        `json:"eight_characters"`
	Animal             animal                 `json:"animal"`
	Total              int                    `json:"total"`
}

func (kirsen *KirsenData) calculateGanzhi() {
	kirsen.GanzhiFiveElements = GanzhiFiveElements(kirsen.Calendar)
}

func (kirsen *KirsenData) calculateSounds() {
	kirsen.SoundFiveElements.YearSound.id, kirsen.SoundFiveElements.YearSound.Name, kirsen.SoundFiveElements.YearSound.Description = GanzhiSoundAlias(kirsen.Calendar.Ganzhi.Year, kirsen.language)
	kirsen.SoundFiveElements.MonthSound.id, kirsen.SoundFiveElements.MonthSound.Name, kirsen.SoundFiveElements.MonthSound.Description = GanzhiSoundAlias(kirsen.Calendar.Ganzhi.Month, kirsen.language)
	kirsen.SoundFiveElements.DaySound.id, kirsen.SoundFiveElements.DaySound.Name, kirsen.SoundFiveElements.DaySound.Description = GanzhiSoundAlias(kirsen.Calendar.Ganzhi.Day, kirsen.language)
	kirsen.SoundFiveElements.HourSound.id, kirsen.SoundFiveElements.HourSound.Name, kirsen.SoundFiveElements.HourSound.Description = GanzhiSoundAlias(kirsen.Calendar.Ganzhi.Hour, kirsen.language)
}

func (kirsen *KirsenData) calculateAnimal() {
	kirsen.Animal.Radicals = getAnimalRadicals(kirsen.Calendar.Lunar.AnimalSign, kirsen.language)
	kirsen.Animal.Years = texts.GetMessage(texts.MessageAnimalYear, kirsen.Calendar.Lunar.AnimalSign, kirsen.language)
}

func (kirsen *KirsenData) calculateEightCharacters() {
	kirsen.EightCharacters.Year = &kirsen.Calendar.Ganzhi.Year
	kirsen.EightCharacters.Month = &kirsen.Calendar.Ganzhi.Month
	kirsen.EightCharacters.Day = &kirsen.Calendar.Ganzhi.Day
	kirsen.EightCharacters.Hour = &kirsen.Calendar.Ganzhi.Hour
	kirsen.EightCharacters.complete()
}

// CalcCommonStrokes : Get strokes of common characters
func CalcCommonStrokes(level int) {
	var (
		charList      map[rune]int32
		c             *unihan.HanCharacter
		stroke        int
		strokeCounter = make(map[int]int)
		count         int
	)

	switch level {
	case 2:
		charList = list.GetCommonL2Traditional()
	case 1:
		charList = list.GetCommonL1Traditional()
	}

	for r := range charList {
		c, _ = unihan.Query(r)
		if c != nil {
			stroke = c.QueryStrokePrefer()
			strokeCounter[stroke]++
		}
	}

	for stroke, count = range strokeCounter {
		fmt.Println("Stroke: ", stroke, "\tCount: ", count)
	}

	return
}

func calcRank(f0, f1, g0, g1 int) int {
	var (
		tianGe, diGe, renGe, zongGe, waiGe                int
		tianCai, diCai, renCai                            int
		scoreTian, scoreDi, scoreRen, scoreZong, scoreWai int
		scoreThreeElement                                 int
		rank                                              int
	)

	_g81 := func(i int) int {
		if i > 81 {
			return i - 80
		}
		return i
	}

	if f1 > 0 {
		tianGe = f0 + f1
		if g1 > 0 {
			renGe = f1 + g0
			waiGe = f0 + g1
		} else {
			renGe = f1
			waiGe = f0 + 1
		}
	} else {
		tianGe = f0 + 1
		if g1 > 0 {
			renGe = f0 + g0
			waiGe = 1 + g1
		} else {
			renGe = f0
			waiGe = 2
		}
	}

	if g1 > 0 {
		diGe = g0 + g1
	} else {
		diGe = g0 + 1
	}

	zongGe = f0 + f1 + g0 + g1

	tianCai = ((tianGe - 1) % 10) / 2
	diCai = ((diGe - 1) % 10) / 2
	renCai = ((renGe - 1) % 10) / 2

	scores := []int{0, 0, 25, 50, 75, 100}
	scoreTian = scores[getRule81Rank(_g81(tianGe))]
	scoreDi = scores[getRule81Rank(_g81(diGe))]
	scoreRen = scores[getRule81Rank(_g81(renGe))]
	scoreZong = scores[getRule81Rank(_g81(zongGe))]
	scoreWai = scores[getRule81Rank(_g81(waiGe))]
	scoreThreeElement = scores[getRuleThreeElementRank(tianCai*25+renCai*5+diCai)]

	rank = int(
		math.Ceil(float64(scoreRen)*0.21) +
			math.Ceil(float64(scoreZong)*0.2) +
			math.Ceil(float64(scoreTian)*0.13) +
			math.Ceil(float64(scoreDi)*0.13) +
			math.Ceil(float64(scoreWai)*0.13) +
			math.Ceil(float64(scoreThreeElement)*0.20))
	if rank > 100 {
		rank = 100
	}

	return rank
}

// FillRankTable : Calculate rank scores and fill into table
func FillRankTable() error {
	var (
		familyNameStroke0, familyNameStroke1 int
		givenNameStroke0, givenNameStroke1   int
		familyNameStroke, givenNameStroke    int
		rank                                 int
	)

	if strokeTable == nil {
		return fmt.Errorf("Fill stroke table failed")
	}

	for familyNameStroke0 = 1; familyNameStroke0 <= list.MaxStroke; familyNameStroke0++ {
		for familyNameStroke1 = 0; familyNameStroke1 <= list.MaxStroke; familyNameStroke1++ {
			familyNameStroke = familyNameStroke1*list.MaxStroke + familyNameStroke0
			for givenNameStroke0 = 1; givenNameStroke0 <= list.MaxStroke; givenNameStroke0++ {
				for givenNameStroke1 = 0; givenNameStroke1 <= list.MaxStroke; givenNameStroke1++ {
					givenNameStroke = givenNameStroke1*list.MaxStroke + givenNameStroke0
					rank = calcRank(familyNameStroke0, familyNameStroke1, givenNameStroke0, givenNameStroke1)
					strokeTable[familyNameStroke][rank] = append(strokeTable[familyNameStroke][rank], givenNameStroke)
				}
			}
		}
	}

	return nil
}

// GetRanksFromTable : Get ranks from table
func GetRanksFromTable(familyNameStroke0, familyNameStroke1 int) [][]int {
	familyNameStroke := familyNameStroke1*list.MaxStroke + familyNameStroke0
	if familyNameStroke < len(strokeTable) {
		return strokeTable[familyNameStroke]
	}

	return nil
}

// Rune to string
func kirsenSingle(list []rune) [][]rune {
	var ret [][]rune
	for _, r := range list {
		c, _ := unihan.Query(r)
		if c != nil {
			ret = append(ret, []rune{c.Unicode})
		}
	}

	return ret
}

func kirsenDouble(list1, list2 []rune) [][]rune {
	var ret [][]rune
	for _, r1 := range list1 {
		c1, _ := unihan.Query(r1)
		if c1 != nil {
			for _, r2 := range list2 {
				c2, _ := unihan.Query(r2)
				if c2 != nil {
					ret = append(ret, []rune{c1.Unicode, c2.Unicode})
				}
			}
		}
	}

	return ret
}

func kirsenTriple(list1, list2, list3 []rune) [][]rune {
	return nil
}

// Kirsen : Fetch name list
func Kirsen(language int, c *KirsenConditions, birthTime int64, loc utils.Location) (*KirsenData, error) {
	var (
		h              *unihan.HanCharacter
		err            error
		f0, f1, g0, g1 int
		g, p, s        int
		rank           int
		sList          [][]int
		cg0, cg1       []rune
		givenNameRunes [][]rune
		total          int
		topRank        int
		name           *Name
		nameList       []*Name
		kirsen         = &KirsenData{}
	)

	kirsen.Calendar = calendar.New(birthTime, loc)
	kirsen.Calendar.Ganzhi.YearString = kirsen.Calendar.Ganzhi.Year.String(kirsen.language)
	kirsen.Calendar.Ganzhi.MonthString = kirsen.Calendar.Ganzhi.Month.String(kirsen.language)
	kirsen.Calendar.Ganzhi.DayString = kirsen.Calendar.Ganzhi.Day.String(kirsen.language)
	kirsen.Calendar.Ganzhi.HourString = kirsen.Calendar.Ganzhi.Hour.String(kirsen.language)

	kirsen.calculateGanzhi()
	kirsen.calculateSounds()
	kirsen.calculateAnimal()
	kirsen.calculateEightCharacters()

	if c.CharacterLevel != 2 {
		c.CharacterLevel = 1
	}

	if c.QueryNums < 0 || c.QueryNums > MaxNames {
		c.QueryNums = MaxNames
	}

	if len(c.FamilyNameRunes) > 1 {
		h, err = unihan.Query(c.FamilyNameRunes[1])
		if err == nil {
			f1 = h.QueryStrokePrefer()
		} else {
			return nil, err
		}
	}

	if len(c.FamilyNameRunes) > 0 {
		h, err = unihan.Query(c.FamilyNameRunes[0])
		if err == nil {
			f0 = h.QueryStrokePrefer()
		} else {
			return nil, err
		}
	}

	if len(c.PrefixNameRunes) > 0 {
		h, err = unihan.Query(c.PrefixNameRunes[0])
		if err == nil {
			p = h.QueryStrokePrefer()
		} else {
			return nil, err
		}
	}

	if len(c.SuffixNameRunes) > 0 {
		h, err = unihan.Query(c.SuffixNameRunes[0])
		if err == nil {
			s = h.QueryStrokePrefer()
		} else {
			return nil, err
		}
	}

	// Fetch table
	sList = GetRanksFromTable(f0, f1)
	for rank = MaxRank; rank > 0; rank-- {
		if sList[rank] == nil {
			continue
		}

		if c.QueryNums == 0 && rank < topRank {
			break
		}

		for _, g = range sList[rank] {
			g0 = g % (list.MaxStroke)
			g1 = g / (list.MaxStroke)

			if g0 > 0 {
				if p > 0 && g0 != p {
					continue
				}

				if c.CharacterLevel == 2 {
					cg0 = list.GetCommonL2ByStrokeTraditional(g0)
				} else {
					cg0 = list.GetCommonL1ByStrokeTraditional(g0)
				}

				if cg0 != nil && len(cg0) > 0 {
					if g1 > 0 && c.GivenNameLength > 1 {
						// Double
						if s > 0 && g1 != s {
							continue
						}

						if c.CharacterLevel == 2 {
							cg1 = list.GetCommonL2ByStrokeTraditional(g1)
						} else {
							cg1 = list.GetCommonL1ByStrokeTraditional(g1)
						}

						if cg1 != nil && len(cg1) > 0 {
							givenNameRunes = kirsenDouble(cg0, cg1)
						}
					} else {
						// Single
						givenNameRunes = kirsenSingle(cg0)
					}

					for _, v := range givenNameRunes {
						if p > 0 && c.PrefixNameRunes[0] != v[0] {
							continue
						}

						if s > 0 && c.SuffixNameRunes[0] != v[len(v)-1] {
							continue
						}

						if topRank == 0 {
							topRank = rank
						}

						name = NewNameRunes(c.FamilyNameRunes, nil, v)
						name.Rank = rank
						nameList = append(nameList, name)
						total++
					}
				}
			}
		}

		if c.QueryNums > 0 && total > c.QueryNums {
			break
		}
	}

	if c.QueryNums > 0 && len(nameList) > c.QueryNums {
		nameList = nameList[:c.QueryNums]
		total = c.QueryNums
	}

	for _, name = range nameList {
		name.Normalize()
		name.RemoveUnihan()
		v := list.QueryCommonNames(fmt.Sprintf("%s%s", name.Simplified.FamilyName.Str, name.Simplified.GivenName.Str))
		if v > 0 {
			name.IsCommon = true
		}
	}

	kirsen.List = nameList
	kirsen.Total = total

	return kirsen, nil
}

func genTable() {
	var (
		maxM = list.MaxStroke * (list.MaxStroke + 1)
	)

	strokeTable = make([][][]int, maxM+1)
	for i := 0; i <= maxM; i++ {
		strokeTable[i] = make([][]int, MaxRank+1)
	}
}

func init() {
	genTable()
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
