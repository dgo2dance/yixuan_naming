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
 * @file eight_characters.go
 * @package name
 * @author Dr.NP <np@corp.herewetech.com>
 * @since 06/27/2019
 */

package name

import (
	"github.com/herewetech/yixuan_naming/utils"
)

type eightCharacters struct {
	Year      *utils.GanzhiPair `json:"year"`
	Month     *utils.GanzhiPair `json:"month"`
	Day       *utils.GanzhiPair `json:"day"`
	Hour      *utils.GanzhiPair `json:"hour"`
	Ling      int               `json:"ling"`
	Shi       int               `json:"shi"`
	ShiYi     int               `json:"shi_yi"`
	Di        bool              `json:"di"`
	Self      int               `json:"self"`
	Like      int               `json:"like"`
	LikeYi    int               `json:"like_yi"`
	Stretch   bool              `json:"stretch"`
	StretchYi bool              `json:"stretch_yi"`
}

func (ec *eightCharacters) calculateLing() int {
	var (
		feDay   = utils.GanFiveElement(ec.Day.TianGan)
		feMonth = utils.ZhiFiveElement(ec.Month.DiZhi)
		c, l    int
	)

	c = utils.CompareFiveElements(feDay, feMonth)
	switch c {
	case utils.FiveElementEqual:
		l = 50
	case utils.FiveElementBirth:
		l = 40
	case utils.FiveElementBirthed:
		l = 30
	case utils.FiveElementKill:
		l = 20
	case utils.FiveElementKilled:
		l = 10
	}

	return l
}

func (ec *eightCharacters) calculateShi() (int, int) {
	var (
		feDay      = utils.GanFiveElement(ec.Day.TianGan)
		ret, retYi int
	)

	for _, c := range []int{
		utils.GanFiveElement(ec.Year.TianGan),
		utils.ZhiFiveElement(ec.Year.DiZhi),
		utils.GanFiveElement(ec.Month.TianGan),
		//utils.ZhiFiveElement(ec.Month.DiZhi),
		utils.ZhiFiveElement(ec.Day.DiZhi),
		utils.GanFiveElement(ec.Hour.TianGan),
		utils.ZhiFiveElement(ec.Hour.DiZhi),
	} {
		//fmt.Println("Shi", c, feDay)
		switch utils.CompareFiveElements(c, feDay) {
		case utils.FiveElementEqual, utils.FiveElementBirth:
			retYi += 10
		default:
			retYi -= 10
		}
	}

	//fmt.Println("ShiYi", utils.ZhiFiveElement(ec.Month.DiZhi), feDay)
	switch utils.CompareFiveElements(utils.ZhiFiveElement(ec.Month.DiZhi), feDay) {
	case utils.FiveElementEqual, utils.FiveElementBirth:
		ret = retYi + 10
	default:
		ret = retYi - 10
	}

	return ret, retYi
}

func (ec *eightCharacters) calculateDi() bool {
	var (
		feDay = utils.GanFiveElement(ec.Day.TianGan)
		ret   = false
	)

	switch feDay {
	case utils.ElementWood:
		switch ec.Day.DiZhi {
		case utils.ZhiHai, utils.ZhiYin, utils.ZhiMao, utils.ZhiWei, utils.ZhiChen:
			ret = true
		}
	case utils.ElementFire:
		switch ec.Day.DiZhi {
		case utils.ZhiYin, utils.ZhiSi, utils.ZhiWu, utils.ZhiWei, utils.ZhiXu:
			ret = true
		}
	case utils.ElementEarth:
		switch ec.Day.DiZhi {
		case utils.ZhiChen, utils.ZhiXu, utils.ZhiChou, utils.ZhiWei, utils.ZhiSi, utils.ZhiWu:
			ret = true
		}
	case utils.ElementMetal:
		switch ec.Day.DiZhi {
		case utils.ZhiSi, utils.ZhiShen, utils.ZhiYou, utils.ZhiXu, utils.ZhiChou:
			ret = true
		}
	case utils.ElementWater:
		switch ec.Day.DiZhi {
		case utils.ZhiShen, utils.ZhiHai, utils.ZhiZi, utils.ZhiChou, utils.ZhiChen:
			ret = true
		}
	}

	return ret
}

func (ec *eightCharacters) complete() {
	ec.Ling = ec.calculateLing()
	ec.Shi, ec.ShiYi = ec.calculateShi()
	ec.Di = ec.calculateDi()
	ec.Stretch = false
	ec.StretchYi = false
	ec.Self = utils.GanFiveElement(ec.Day.TianGan)

	// Normal
	c := 0
	if ec.Ling >= 50 {
		c++
	}

	if ec.Shi >= 10 {
		c++
	}

	if ec.Di {
		c++
	}

	if c >= 2 {
		ec.Stretch = true
	}

	if ec.Stretch {
		ec.Like = ec.Self + 2
	} else {
		ec.Like = ec.Self - 1
	}

	if ec.Like < 0 {
		ec.Like += 5
	}

	if ec.Like >= 5 {
		ec.Like -= 5
	}

	// YiXuan
	if ec.Ling+ec.ShiYi >= 50 {
		ec.StretchYi = true
	}

	feDay := utils.GanFiveElement(ec.Day.TianGan)
	if utils.CompareFiveElements(feDay, utils.ZhiFiveElement(ec.Month.DiZhi)) == utils.FiveElementKilled {
		if utils.CompareFiveElements(feDay, utils.GanFiveElement(ec.Month.TianGan)) == utils.FiveElementBirthed &&
			utils.CompareFiveElements(feDay, utils.ZhiFiveElement(ec.Day.DiZhi)) == utils.FiveElementBirthed {
			ec.StretchYi = true
		}
	}

	yi2 := 0
	yi4 := 0
	if utils.CompareFiveElements(feDay, utils.ZhiFiveElement(ec.Month.DiZhi)) == utils.FiveElementBirthed {
		if utils.CompareFiveElements(feDay, utils.GanFiveElement(ec.Month.TianGan)) == utils.FiveElementKill {
			yi2++
		}
		if utils.CompareFiveElements(feDay, utils.ZhiFiveElement(ec.Day.DiZhi)) == utils.FiveElementKill {
			yi2++
		}
		if utils.CompareFiveElements(feDay, utils.ZhiFiveElement(ec.Year.DiZhi)) == utils.FiveElementKill {
			yi2++
		}
		if utils.CompareFiveElements(feDay, utils.GanFiveElement(ec.Month.TianGan)) == utils.FiveElementBirthed {
			yi4++
		}
		if utils.CompareFiveElements(feDay, utils.ZhiFiveElement(ec.Day.DiZhi)) == utils.FiveElementKilled {
			yi4++
		}
		if utils.CompareFiveElements(feDay, utils.ZhiFiveElement(ec.Year.DiZhi)) == utils.FiveElementKilled {
			yi4++
		}
	}

	yi3 := 0
	if ec.Ling == 50 {
		if utils.CompareFiveElements(feDay, utils.GanFiveElement(ec.Month.TianGan)) == utils.FiveElementKilled {
			yi3++
		}
		if utils.CompareFiveElements(feDay, utils.ZhiFiveElement(ec.Day.DiZhi)) == utils.FiveElementKilled {
			yi3++
		}
		if utils.CompareFiveElements(feDay, utils.ZhiFiveElement(ec.Year.DiZhi)) == utils.FiveElementKilled {
			yi3++
		}
	}

	if yi2 > 2 || yi3 > 2 {
		ec.StretchYi = false
	}

	if ec.StretchYi {
		ec.LikeYi = ec.Self + 2
	} else {
		ec.LikeYi = ec.Self - 1
	}

	if ec.LikeYi < 0 {
		ec.LikeYi += 5
	}

	if ec.LikeYi >= 5 {
		ec.LikeYi -= 5
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
