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
 * @file rules.go
 * @package name
 * @author Dr.NP <np@corp.herewetech.com>
 * @since 06/06/2019
 */

package name

import (
	"github.com/herewetech/yixuan_naming/texts"
)

type rule81 struct {
	Summary string `json:"summary"`
	Details string `json:"details"`
	Rank    int    `json:"rank"`
}

type ruleThreeElement struct {
	Details string `json:"details"`
	Rank    int    `json:"rank"`
}

const (
	// RankNone : Not a valid rank
	RankNone int = iota
	// RankDaXiong : 1
	RankDaXiong
	// RankXiong : 2
	RankXiong
	// RankBanJi : 3
	RankBanJi
	// RankJi : 4
	RankJi
	// RankDaJi : 5
	RankDaJi
)

var (
	rule81Ranks = []int{
		RankNone, RankDaJi, RankDaXiong, RankDaJi, RankXiong, RankDaJi,
		RankBanJi, RankJi, RankJi, RankXiong, RankDaXiong, RankDaJi,
		RankXiong, RankJi, RankDaXiong, RankDaJi, RankDaJi, RankJi,
		RankDaJi, RankDaXiong, RankDaXiong, RankJi, RankDaXiong, RankDaJi,
		RankDaJi, RankJi, RankBanJi, RankBanJi, RankDaXiong, RankJi,
		RankBanJi, RankDaJi, RankDaJi, RankDaJi, RankDaXiong, RankDaJi,
		RankDaXiong, RankDaJi, RankBanJi, RankDaJi, RankXiong, RankDaJi,
		RankBanJi, RankDaXiong, RankDaXiong, RankDaJi, RankDaXiong, RankDaJi,
		RankDaJi, RankBanJi, RankXiong, RankBanJi, RankDaJi, RankDaXiong,
		RankDaXiong, RankBanJi, RankDaXiong, RankDaJi, RankBanJi, RankDaXiong,
		RankDaXiong, RankDaJi, RankDaXiong, RankDaJi, RankDaXiong, RankDaJi,
		RankDaXiong, RankDaJi, RankDaJi, RankDaXiong, RankXiong, RankJi,
		RankXiong, RankBanJi, RankDaXiong, RankBanJi, RankDaXiong, RankBanJi,
		RankDaXiong, RankDaXiong, RankDaXiong, RankDaJi,
	}

	ruleThreeElementRanks = []int{
		RankDaJi, RankDaJi, RankDaJi, RankDaXiong, RankDaXiong, RankDaJi,
		RankBanJi, RankDaJi, RankDaXiong, RankDaXiong, RankDaXiong, RankBanJi,
		RankDaXiong, RankDaXiong, RankDaXiong, RankDaXiong, RankDaXiong,
		RankDaXiong, RankDaXiong, RankDaXiong, RankDaJi, RankDaXiong, RankDaXiong,
		RankJi, RankBanJi, RankDaJi, RankDaJi, RankDaJi, RankDaXiong, RankDaXiong,
		RankDaJi, RankBanJi, RankDaXiong, RankDaXiong, RankDaXiong, RankBanJi,
		RankDaJi, RankDaJi, RankDaXiong, RankDaXiong, RankDaXiong, RankDaXiong,
		RankDaXiong, RankDaXiong, RankDaXiong, RankDaXiong, RankXiong, RankDaXiong,
		RankDaXiong, RankDaXiong, RankBanJi, RankBanJi, RankDaXiong, RankDaXiong,
		RankDaXiong, RankDaJi, RankBanJi, RankDaJi, RankXiong, RankDaXiong,
		RankDaXiong, RankDaJi, RankDaJi, RankDaJi, RankDaJi, RankDaXiong, RankBanJi,
		RankDaJi, RankDaJi, RankDaXiong, RankDaXiong, RankDaXiong, RankDaXiong,
		RankDaXiong, RankDaXiong, RankDaXiong, RankDaXiong, RankDaXiong,
		RankDaXiong, RankDaXiong, RankDaXiong, RankDaXiong, RankDaXiong,
		RankDaXiong, RankDaXiong, RankBanJi, RankDaJi, RankDaJi, RankDaJi,
		RankDaXiong, RankDaXiong, RankDaXiong, RankDaJi, RankDaXiong, RankDaXiong,
		RankBanJi, RankDaXiong, RankDaXiong, RankDaJi, RankDaXiong, RankDaJi,
		RankBanJi, RankDaJi, RankDaXiong, RankBanJi, RankDaXiong, RankDaXiong,
		RankDaXiong, RankDaXiong, RankDaXiong, RankDaXiong, RankDaXiong,
		RankDaXiong, RankDaXiong, RankDaXiong, RankDaXiong, RankBanJi, RankDaJi,
		RankBanJi, RankDaXiong, RankBanJi, RankDaXiong, RankDaXiong, RankBanJi,
		RankDaXiong,
	}
)

func getRule81(index int, language int) *rule81 {
	if index < 0 && index > 81 {
		return nil
	}

	ret := &rule81{Rank: rule81Ranks[index]}
	ret.Summary = texts.GetMessage(texts.MessageFiveRuleSummary, index, language)
	ret.Details = texts.GetMessage(texts.MessageFiveRuleDescription, index, language)

	return ret
}

func getRuleThreeElement(index int, language int) *ruleThreeElement {
	if index < 0 && index > 124 {
		return nil
	}

	ret := &ruleThreeElement{Rank: ruleThreeElementRanks[index]}
	ret.Details = texts.GetMessage(texts.MessageThreeElementDescription, index, language)

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
