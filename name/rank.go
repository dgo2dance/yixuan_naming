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
 * @file rank.go
 * @package name
 * @author Dr.NP <np@corp.herewetech.com>
 * @since 06/05/2019
 */

package name

import (
	"fmt"
	"math"
	"strings"

	"yixuan_naming/calendar"
	"yixuan_naming/dict"
	"yixuan_naming/list"
	"yixuan_naming/poetry"
	"yixuan_naming/texts"
	"yixuan_naming/utils"
)

/*
type loc struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
*/

type fiveRules struct {
	TianGe                       int     `json:"tian_ge"`
	TianGeFiveElement            string  `json:"tian_ge_five_element"`
	TianGeFiveElementDescription string  `json:"tian_ge_five_element_description"`
	TianGeGod                    *tenGod `json:"tian_ge_god"`
	TianGeRule                   *rule81 `json:"tian_ge_rule"`
	TianGeRuleRank               string  `json:"tian_ge_rank"`

	DiGe                       int     `json:"di_ge"`
	DiGeFiveElement            string  `json:"di_ge_five_element"`
	DiGeFiveElementDescription string  `json:"di_ge_five_element_description"`
	DiGeGod                    *tenGod `json:"di_ge_god"`
	DiGeRule                   *rule81 `json:"di_ge_rule"`
	DiGeRuleRank               string  `json:"di_ge_rank"`

	RenGe                       int     `json:"ren_ge"`
	RenGeFiveElement            string  `json:"ren_ge_five_element"`
	RenGeFiveElementDescription string  `json:"ren_ge_five_element_description"`
	RenGeRule                   *rule81 `json:"ren_ge_rule"`
	RenGeRuleRank               string  `json:"ren_ge_rank"`

	ZongGe                       int     `json:"zong_ge"`
	ZongGeFiveElement            string  `json:"zong_ge_five_element"`
	ZongGeFiveElementDescription string  `json:"zong_ge_five_element_description"`
	ZongGeGod                    *tenGod `json:"zong_ge_god"`
	ZongGeRule                   *rule81 `json:"zong_ge_rule"`
	ZongGeRuleRank               string  `json:"zong_ge_rank"`

	WaiGe                       int     `json:"wai_ge"`
	WaiGeFiveElement            string  `json:"wai_ge_five_element"`
	WaiGeFiveElementDescription string  `json:"wai_ge_five_element_description"`
	WaiGeGod                    *tenGod `json:"wai_ge_god"`
	WaiGeRule                   *rule81 `json:"wai_ge_rule"`
	WaiGeRuleRank               string  `json:"wai_ge_rank"`

	ThreeElement     *ruleThreeElement `json:"three_element"`
	ThreeElementRank string            `json:"three_element_rank"`
}

type soundFiveElement struct {
	id          int
	Name        string `json:"name"`
	Description string `json:"description"`
}

type dictXinhua struct {
	FamilyNameXinhua []*dict.XinhuaItem `json:"family_name_xinhua,omitempty"`
	MiddleNameXinhua []*dict.XinhuaItem `json:"middle_name_xinhua,omitempty"`
	GivenNameXinhua  []*dict.XinhuaItem `json:"given_name_xinhua,omitempty"`
}

type dictFolkways struct {
	FamilyNameFolkway []*dict.FolkwayItem `json:"family_name_folkway,omitempty"`
	MiddleNameFolkway []*dict.FolkwayItem `json:"middle_name_folkway,omitempty"`
	GivenNameFolkway  []*dict.FolkwayItem `json:"given_name_folkway,omitempty"`
}

type rank struct {
	RankFiveRules       int    `json:"rank_five_rules"`
	RankFiveElements    int    `json:"rank_five_elements"`
	RankEightCharacters int    `json:"rank_eight_characters"`
	RankTotal           int    `json:"rank_total"`
	RankDescription     string `json:"rank_description"`
}

// RankData : struct of name ranking result
type RankData struct {
	language           int
	Name               *Name                  `json:"name"`
	DictXinhua         dictXinhua             `json:"dict_xinhua"`
	DictFolkways       dictFolkways           `json:"dict_folkways"`
	BaiJiaXing         *list.BaiJiaXing       `json:"bai_jia_xing,omitempty"`
	Poetries           []*poetry.Poetry       `json:"poetries,omitempty"`
	FiveRules          fiveRules              `json:"five_rules"`
	EightCharacters    eightCharacters        `json:"eight_characters"`
	Calendar           *calendar.Calendar     `json:"calendar"`
	GanzhiFiveElements GanzhiFiveElementsSpec `json:"ganzhi_five_elements"`
	SoundFiveElements  SoundFiveElements      `json:"sound_five_elements"`
	Animal             animal                 `json:"animal"`
	Rank               rank                   `json:"rank"`
	Homonyms           []string               `json:"homonyms"`
	FamilyNameScore    int                    `json:"family_name_score"`
	CommonName         bool                   `json:"common_name"`
	Illegal            bool                   `json:"illegal"`
}

func (rank *RankData) calculateFiveRules() {
	// Traditional name prefered
	n := &rank.Name.Traditional
	if n.FamilyName.Len >= 2 {
		// hyphenated
		rank.FiveRules.TianGe = n.FamilyName.Strokes[0] + n.FamilyName.Strokes[1]
		if n.GivenName.Len >= 1 {
			rank.FiveRules.RenGe = n.FamilyName.Strokes[1] + n.GivenName.Strokes[0]
		} else {
			rank.FiveRules.RenGe = n.FamilyName.Strokes[1]
		}
	} else if n.FamilyName.Len == 1 {
		rank.FiveRules.TianGe = n.FamilyName.Strokes[0] + 1
		if n.GivenName.Len >= 1 {
			rank.FiveRules.RenGe = n.FamilyName.Strokes[0] + n.GivenName.Strokes[0]
		} else {
			rank.FiveRules.RenGe = n.FamilyName.Strokes[0]
		}
	}

	if n.GivenName.Len >= 2 {
		rank.FiveRules.DiGe = n.GivenName.Strokes[0] + n.GivenName.Strokes[1]
	} else if n.GivenName.Len == 1 {
		rank.FiveRules.DiGe = n.GivenName.Strokes[0] + 1
	}

	if n.FamilyName.Len >= 2 {
		if n.GivenName.Len >= 2 {
			rank.FiveRules.WaiGe = n.FamilyName.Strokes[0] + n.GivenName.Strokes[1]
		} else if n.GivenName.Len == 1 {
			rank.FiveRules.WaiGe = n.FamilyName.Strokes[0] + 1
		}
	} else if n.FamilyName.Len == 1 {
		if n.GivenName.Len >= 2 {
			rank.FiveRules.WaiGe = 1 + n.GivenName.Strokes[1]
		} else if n.GivenName.Len == 1 {
			rank.FiveRules.WaiGe = 2
		}
	}

	for k, h := range n.FamilyName.Strokes {
		if k < 2 {
			rank.FiveRules.ZongGe = rank.FiveRules.ZongGe + h
		}
	}
	for k, h := range n.GivenName.Strokes {
		if k < 2 {
			rank.FiveRules.ZongGe = rank.FiveRules.ZongGe + h
		}
	}

	_mod := func(i, m int) int {
		r := i % m
		if r == 0 {
			r = m
		}

		return r
	}

	tf := _mod(rank.FiveRules.TianGe, 10) - 1
	df := _mod(rank.FiveRules.DiGe, 10) - 1
	rf := _mod(rank.FiveRules.RenGe, 10) - 1
	zf := _mod(rank.FiveRules.ZongGe, 10) - 1
	wf := _mod(rank.FiveRules.WaiGe, 10) - 1
	rank.FiveRules.TianGeFiveElement = texts.GetAlias(texts.AliasGanFiveElement, tf, rank.language)
	rank.FiveRules.TianGeFiveElementDescription = texts.GetMessage(texts.MessageGanDescription, tf, rank.language)
	rank.FiveRules.DiGeFiveElement = texts.GetAlias(texts.AliasGanFiveElement, df, rank.language)
	rank.FiveRules.DiGeFiveElementDescription = texts.GetMessage(texts.MessageGanDescription, df, rank.language)
	rank.FiveRules.RenGeFiveElement = texts.GetAlias(texts.AliasGanFiveElement, rf, rank.language)
	rank.FiveRules.RenGeFiveElementDescription = texts.GetMessage(texts.MessageGanDescription, rf, rank.language)
	rank.FiveRules.ZongGeFiveElement = texts.GetAlias(texts.AliasGanFiveElement, zf, rank.language)
	rank.FiveRules.ZongGeFiveElementDescription = texts.GetMessage(texts.MessageGanDescription, zf, rank.language)
	rank.FiveRules.WaiGeFiveElement = texts.GetAlias(texts.AliasGanFiveElement, wf, rank.language)
	rank.FiveRules.WaiGeFiveElementDescription = texts.GetMessage(texts.MessageGanDescription, wf, rank.language)

	tg := utils.CompareGan(rf, tf)
	dg := utils.CompareGan(rf, df)
	zg := utils.CompareGan(rf, zf)
	wg := utils.CompareGan(rf, wf)
	rank.FiveRules.TianGeGod = getTenGod(tg, rank.language)
	rank.FiveRules.DiGeGod = getTenGod(dg, rank.language)
	rank.FiveRules.ZongGeGod = getTenGod(zg, rank.language)
	rank.FiveRules.WaiGeGod = getTenGod(wg, rank.language)

	_g81 := func(i int) int {
		if i > 81 {
			return i - 80
		}
		return i
	}

	rank.FiveRules.TianGeRule = getRule81(_g81(rank.FiveRules.TianGe), rank.language)
	rank.FiveRules.DiGeRule = getRule81(_g81(rank.FiveRules.DiGe), rank.language)
	rank.FiveRules.RenGeRule = getRule81(_g81(rank.FiveRules.RenGe), rank.language)
	rank.FiveRules.ZongGeRule = getRule81(_g81(rank.FiveRules.ZongGe), rank.language)
	rank.FiveRules.WaiGeRule = getRule81(_g81(rank.FiveRules.WaiGe), rank.language)

	rank.FiveRules.TianGeRuleRank = texts.GetAlias(texts.AliasRank, rank.FiveRules.TianGeRule.Rank, rank.language)
	rank.FiveRules.DiGeRuleRank = texts.GetAlias(texts.AliasRank, rank.FiveRules.DiGeRule.Rank, rank.language)
	rank.FiveRules.RenGeRuleRank = texts.GetAlias(texts.AliasRank, rank.FiveRules.RenGeRule.Rank, rank.language)
	rank.FiveRules.ZongGeRuleRank = texts.GetAlias(texts.AliasRank, rank.FiveRules.ZongGeRule.Rank, rank.language)
	rank.FiveRules.WaiGeRuleRank = texts.GetAlias(texts.AliasRank, rank.FiveRules.WaiGeRule.Rank, rank.language)
}

func (rank *RankData) calculateEightCharacters() {
	rank.EightCharacters.Year = &rank.Calendar.Ganzhi.Year
	rank.EightCharacters.Month = &rank.Calendar.Ganzhi.Month
	rank.EightCharacters.Day = &rank.Calendar.Ganzhi.Day
	rank.EightCharacters.Hour = &rank.Calendar.Ganzhi.Hour
	rank.EightCharacters.complete()
}

func (rank *RankData) calculateGanzhi() {
	rank.GanzhiFiveElements = GanzhiFiveElements(rank.Calendar)
}

func (rank *RankData) calculateSounds() {
	rank.SoundFiveElements.YearSound.id, rank.SoundFiveElements.YearSound.Name, rank.SoundFiveElements.YearSound.Description = GanzhiSoundAlias(rank.Calendar.Ganzhi.Year, rank.language)
	rank.SoundFiveElements.MonthSound.id, rank.SoundFiveElements.MonthSound.Name, rank.SoundFiveElements.MonthSound.Description = GanzhiSoundAlias(rank.Calendar.Ganzhi.Month, rank.language)
	rank.SoundFiveElements.DaySound.id, rank.SoundFiveElements.DaySound.Name, rank.SoundFiveElements.DaySound.Description = GanzhiSoundAlias(rank.Calendar.Ganzhi.Day, rank.language)
	rank.SoundFiveElements.HourSound.id, rank.SoundFiveElements.HourSound.Name, rank.SoundFiveElements.HourSound.Description = GanzhiSoundAlias(rank.Calendar.Ganzhi.Hour, rank.language)
}

func (rank *RankData) calculateAnimal() {
	rank.Animal.Radicals = getAnimalRadicals(rank.Calendar.Lunar.AnimalSign, rank.language)
	rank.Animal.Years = texts.GetMessage(texts.MessageAnimalYear, rank.Calendar.Lunar.AnimalSign, rank.language)
}

func (rank *RankData) calculateRankFiveRules() {
	scores := []int{0, 0, 25, 50, 75, 100}

	// FiveRules
	scoreTian := scores[rank.FiveRules.TianGeRule.Rank]
	scoreRen := scores[rank.FiveRules.RenGeRule.Rank]
	scoreDi := scores[rank.FiveRules.DiGeRule.Rank]
	scoreZong := scores[rank.FiveRules.ZongGeRule.Rank]
	scoreWai := scores[rank.FiveRules.WaiGeRule.Rank]

	// ThreeRules
	tianCai := ((rank.FiveRules.TianGe - 1) % 10) / 2
	renCai := ((rank.FiveRules.RenGe - 1) % 10) / 2
	diCai := ((rank.FiveRules.DiGe - 1) % 10) / 2
	threeElement := tianCai*25 + renCai*5 + diCai
	rank.FiveRules.ThreeElement = getRuleThreeElement(threeElement, rank.language)
	rank.FiveRules.ThreeElementRank = texts.GetAlias(texts.AliasRank, rank.FiveRules.ThreeElement.Rank, rank.language)
	scoreThreeElement := scores[rank.FiveRules.ThreeElement.Rank]

	rank.Rank.RankFiveRules = int(
		math.Ceil(float64(scoreRen)*0.21) +
			math.Ceil(float64(scoreZong)*0.2) +
			math.Ceil(float64(scoreTian)*0.13) +
			math.Ceil(float64(scoreDi)*0.13) +
			math.Ceil(float64(scoreWai)*0.13) +
			math.Ceil(float64(scoreThreeElement)*0.20))
	if rank.Rank.RankFiveRules > 100 {
		rank.Rank.RankFiveRules = 100
	}
}

func (rank *RankData) calculateRankEightElements() {
}

func (rank *RankData) calculateRanks() {
	rank.calculateRankFiveRules()
	rank.calculateRankEightElements()
}

func (rank *RankData) queryXinhua() {
	n := &rank.Name.Simplified
	for _, r := range n.FamilyName.Runes {
		rank.DictXinhua.FamilyNameXinhua = append(rank.DictXinhua.FamilyNameXinhua, dict.QueryXinhua(r))
	}

	for _, r := range n.MiddleName.Runes {
		rank.DictXinhua.MiddleNameXinhua = append(rank.DictXinhua.MiddleNameXinhua, dict.QueryXinhua(r))
	}

	for _, r := range n.GivenName.Runes {
		rank.DictXinhua.GivenNameXinhua = append(rank.DictXinhua.GivenNameXinhua, dict.QueryXinhua(r))
	}
}

func (rank *RankData) queryFolkways() {
	n := &rank.Name.Simplified
	for _, r := range n.FamilyName.Runes {
		rank.DictFolkways.FamilyNameFolkway = append(rank.DictFolkways.FamilyNameFolkway, dict.QueryFolkway(r))
	}

	for _, r := range n.MiddleName.Runes {
		rank.DictFolkways.MiddleNameFolkway = append(rank.DictFolkways.MiddleNameFolkway, dict.QueryFolkway(r))
	}

	for _, r := range n.GivenName.Runes {
		rank.DictFolkways.GivenNameFolkway = append(rank.DictFolkways.GivenNameFolkway, dict.QueryFolkway(r))
	}
}

func (rank *RankData) queryDictionaries() {
	rank.queryXinhua()
	rank.queryFolkways()
}

func (rank *RankData) queryBaiJiaXing() {
	rank.BaiJiaXing = list.QueryBaiJiaXing(rank.Name.Simplified.FamilyName.Str)
}

func (rank *RankData) queryCommonName() {
	v := list.QueryCommonNames(fmt.Sprintf("%s%s", rank.Name.Simplified.FamilyName.Str, rank.Name.Simplified.GivenName.Str))
	if v > 0 {
		rank.CommonName = true
		rank.Name.IsCommon = true
	}
}

func (rank *RankData) queryFamilyName() {
	rank.FamilyNameScore = list.QueryFamilyName(rank.Name.Simplified.FamilyName.Str)
}

func (rank *RankData) queryPoetry() {
	rank.Poetries = poetry.QueryPoetries(rank.Name.Simplified.GivenName.Str)
	return
}

func groupPinyin(pinyin []string) [][]string {
	var (
		l, i  int
		t     = len(pinyin)
		group [][]string
	)
	for l = 2; l <= len(pinyin); l++ {
		for i = 0; i <= (t - l); i++ {
			group = append(group, pinyin[i:i+l])
		}
	}

	return group
}

// Rank : Rank name with birth time
func Rank(language int, name *Name, birthTime int64, loc utils.Location) (*RankData, error) {
	var (
		rank        = &RankData{language: language, Name: name, Illegal: false}
		pinyinGroup [][]string
		pinyin      string
		sensitives  []string
		commons     []string
	)

	pinyinGroup = groupPinyin(name.Pinyin)
	for _, p := range pinyinGroup {
		pinyin = strings.Join(p, ",")
		sensitives = list.QuerySensitive(pinyin)
		commons = list.QueryCommon(pinyin)

		if sensitives != nil {
			rank.Illegal = true
			return rank, fmt.Errorf("Illegal name")
		}

		if commons != nil {
			rank.Homonyms = append(rank.Homonyms, commons...)
		}
	}

	rank.Calendar = calendar.New(birthTime, loc)
	rank.Calendar.Ganzhi.YearString = rank.Calendar.Ganzhi.Year.String(rank.language)
	rank.Calendar.Ganzhi.MonthString = rank.Calendar.Ganzhi.Month.String(rank.language)
	rank.Calendar.Ganzhi.DayString = rank.Calendar.Ganzhi.Day.String(rank.language)
	rank.Calendar.Ganzhi.HourString = rank.Calendar.Ganzhi.Hour.String(rank.language)

	// For Chinese, ignore middle name now
	rank.calculateFiveRules()
	rank.calculateEightCharacters()
	rank.calculateGanzhi()
	rank.calculateSounds()
	rank.calculateAnimal()
	rank.queryDictionaries()
	rank.queryBaiJiaXing()
	rank.queryCommonName()
	rank.queryFamilyName()
	rank.queryPoetry()
	rank.calculateRanks()

	return rank, nil
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
