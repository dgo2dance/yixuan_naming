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
 * @package name
 * @author Dr.NP <np@corp.herewetech.com>
 * @since 06/06/2019
 */

package name

import (
	"yixuan_naming/calendar"
	"yixuan_naming/texts"
	"yixuan_naming/utils"
)

// GanzhiFiveElementsSpec : Counter of five-elements of ganzhi calendar
type GanzhiFiveElementsSpec struct {
	FiveElements      utils.FiveElementsCount `json:"five_elements"`
	FiveElementsZhi   utils.FiveElementsCount `json:"five_elements_zhi"`
	FiveElementsTotal utils.FiveElementsCount `json:"five_elements_total"`
}

// GanzhiSound : Get sound five-elements by given gan-zhi
func GanzhiSound(ganzhi utils.GanzhiPair) int {
	return ganzhi.Value()
}

// GanzhiSoundAlias : Get sound five-elements id / name / description by given gan-zhi
func GanzhiSoundAlias(ganzhi utils.GanzhiPair, language int) (int, string, string) {
	id := ganzhi.Value()
	if id >= 0 && id < 60 {
		return id,
			texts.GetAlias(texts.AliasSoundFiveElement, id/2, language),
			texts.GetMessage(texts.MessageSoundFiveElementDescription, id, language)
	}

	return 0, "", ""
}

// GanzhiFiveElements : Count five-elements of ganzhi calendar
func GanzhiFiveElements(c *calendar.Calendar) GanzhiFiveElementsSpec {
	var (
		v   int
		ret = GanzhiFiveElementsSpec{}
	)

	// TianGan
	for _, v = range []int{c.Ganzhi.Year.TianGan, c.Ganzhi.Month.TianGan, c.Ganzhi.Day.TianGan, c.Ganzhi.Hour.TianGan} {
		switch v {
		case utils.GanJia, utils.GanYi:
			ret.FiveElements.Wood++
		case utils.GanBing, utils.GanDing:
			ret.FiveElements.Fire++
		case utils.GanWu, utils.GanJi:
			ret.FiveElements.Earth++
		case utils.GanGeng, utils.GanXin:
			ret.FiveElements.Metal++
		case utils.GanRen, utils.GanGui:
			ret.FiveElements.Water++
		}
	}

	// DiZhi & ZhiCang
	for _, v = range []int{c.Ganzhi.Year.DiZhi, c.Ganzhi.Month.DiZhi, c.Ganzhi.Day.DiZhi, c.Ganzhi.Hour.DiZhi} {
		switch v {
		case utils.ZhiZi:
			ret.FiveElements.Water++
			ret.FiveElementsZhi.Water++ // 癸
		case utils.ZhiChou:
			ret.FiveElements.Earth++
			ret.FiveElementsZhi.Earth++ // 己
			ret.FiveElementsZhi.Metal++ // 辛
			ret.FiveElementsZhi.Water++ // 癸
		case utils.ZhiYin:
			ret.FiveElements.Wood++
			ret.FiveElementsZhi.Wood++  // 甲
			ret.FiveElementsZhi.Fire++  // 丙
			ret.FiveElementsZhi.Earth++ // 戊
		case utils.ZhiMao:
			ret.FiveElements.Wood++
			ret.FiveElementsZhi.Wood++ // 乙
		case utils.ZhiChen:
			ret.FiveElements.Earth++
			ret.FiveElementsZhi.Earth++ // 戊
			ret.FiveElementsZhi.Water++ // 癸
			ret.FiveElementsZhi.Wood++  // 乙
		case utils.ZhiSi:
			ret.FiveElements.Fire++
			ret.FiveElementsZhi.Fire++  // 丙
			ret.FiveElementsZhi.Earth++ // 戊
			ret.FiveElementsZhi.Metal++ // 庚
		case utils.ZhiWu:
			ret.FiveElements.Fire++
			ret.FiveElementsZhi.Fire++  // 丁
			ret.FiveElementsZhi.Earth++ // 己
		case utils.ZhiWei:
			ret.FiveElements.Earth++
			ret.FiveElementsZhi.Earth++ // 己
			ret.FiveElementsZhi.Wood++  // 乙
			ret.FiveElementsZhi.Fire++  // 丁
		case utils.ZhiShen:
			ret.FiveElements.Metal++
			ret.FiveElementsZhi.Metal++ // 庚
			ret.FiveElementsZhi.Water++ // 壬
			ret.FiveElementsZhi.Earth++ // 戊
		case utils.ZhiYou:
			ret.FiveElements.Metal++
			ret.FiveElementsZhi.Metal++ // 辛
		case utils.ZhiXu:
			ret.FiveElements.Earth++
			ret.FiveElementsZhi.Earth++ // 戊
			ret.FiveElementsZhi.Metal++ // 辛
			ret.FiveElementsZhi.Fire++  // 丁
		case utils.ZhiHai:
			ret.FiveElements.Water++
			ret.FiveElementsZhi.Water++ // 壬
			ret.FiveElementsZhi.Wood++  // 甲
		}
	}

	ret.FiveElementsTotal.Wood = ret.FiveElements.Wood + ret.FiveElementsZhi.Wood
	ret.FiveElementsTotal.Fire = ret.FiveElements.Fire + ret.FiveElementsZhi.Fire
	ret.FiveElementsTotal.Earth = ret.FiveElements.Earth + ret.FiveElementsZhi.Earth
	ret.FiveElementsTotal.Metal = ret.FiveElements.Metal + ret.FiveElementsZhi.Metal
	ret.FiveElementsTotal.Water = ret.FiveElements.Water + ret.FiveElementsZhi.Water

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
