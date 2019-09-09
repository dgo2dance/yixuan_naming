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
 * @file ganzhi.go
 * @package utils
 * @author Dr.NP <np@corp.herewetech.com>
 * @since 06/27/2019
 */

package utils

import (
	"fmt"

	"github.com/herewetech/yixuan_naming/texts"
)

// TianGan * 10
const (
	// GanJia : TianGan(0)
	GanJia = iota
	// GanYi : TianGan(1)
	GanYi
	// GanBing : TianGan(2)
	GanBing
	// GanDing : TianGan(3)
	GanDing
	// GanWu : TianGan(4)
	GanWu
	// GanJi : TianGan(5)
	GanJi
	// GanGeng : TianGan(6)
	GanGeng
	// GanXin : TianGan(7)
	GanXin
	// GanRen : TianGan(8)
	GanRen
	// GanGui : TianGan(9)
	GanGui
)

// DiZhi * 12
const (
	// ZhiZi : DiZhi(0)
	ZhiZi = iota
	// ZhiChou : DiZhi(1)
	ZhiChou
	// ZhiYin : DiZhi(2)
	ZhiYin
	// ZhiMao : DiZhi(3)
	ZhiMao
	// ZhiChen : DiZhi(4)
	ZhiChen
	// ZhiSi : DiZhi(5)
	ZhiSi
	// ZhiWu : DiZhi(6)
	ZhiWu
	// ZhiWei : DiZhi(7)
	ZhiWei
	// ZhiShen : DiZhi(8)
	ZhiShen
	// ZhiYou : DiZhi(9)
	ZhiYou
	// ZhiXu : DiZhi(10)
	ZhiXu
	// ZhiHai : DiZhi(11)
	ZhiHai
)

// GanFiveElement : Five-element of tiangan
func GanFiveElement(gan int) int {
	switch gan {
	case GanJia, GanYi:
		return ElementWood
	case GanBing, GanDing:
		return ElementFire
	case GanWu, GanJi:
		return ElementEarth
	case GanGeng, GanXin:
		return ElementMetal
	case GanRen, GanGui:
		return ElementWater
	}

	return -1
}

// ZhiFiveElement : Five-element of dizhi
func ZhiFiveElement(zhi int) int {
	switch zhi {
	case ZhiYin, ZhiMao:
		return ElementWood
	case ZhiSi, ZhiWu:
		return ElementFire
	case ZhiChen, ZhiXu, ZhiChou, ZhiWei:
		return ElementEarth
	case ZhiShen, ZhiYou:
		return ElementMetal
	case ZhiHai, ZhiZi:
		return ElementWater
	}

	return -1
}

// GanYinYang : YinYang of tiangan
func GanYinYang(gan int) int {
	if gan%2 == 0 {
		return YinYangYang
	}

	return YinYangYin
}

// ZhiYinYang : YinYang of dizhi
func ZhiYinYang(zhi int) int {
	if zhi%2 == 0 {
		return YinYangYang
	}

	return YinYangYin
}

// GanzhiPair : TianGan0DiZhi
type GanzhiPair struct {
	TianGan int `json:"tian_gan"`
	DiZhi   int `json:"di_zhi"`
}

// valid : If GanzhiPair valid
func (gz *GanzhiPair) valid() bool {
	if gz.TianGan >= 0 &&
		gz.TianGan < 10 &&
		gz.DiZhi >= 0 &&
		gz.DiZhi < 12 &&
		gz.TianGan%2 == gz.DiZhi%2 {
		return true
	}

	return false
}

// Value : Get value of GanzhiPair
func (gz *GanzhiPair) Value() int {
	if !gz.valid() {
		return -1
	}

	for v := gz.DiZhi; v < 60; v += 12 {
		if v%10 == gz.TianGan {
			return v
		}
	}

	return -1
}

// String : Get string (aliases) of GanzhiPair
func (gz *GanzhiPair) String(language int) string {
	if !gz.valid() {
		return "_"
	}

	return fmt.Sprintf("%s%s",
		texts.GetAlias(texts.AliasGan, gz.TianGan, language),
		texts.GetAlias(texts.AliasZhi, gz.DiZhi, language))
}

// ParseGanzhi : Parse int to GanzhiPair
func ParseGanzhi(v int) *GanzhiPair {
	return &GanzhiPair{
		TianGan: v % 10,
		DiZhi:   v % 12,
	}
}

// CompareGan : Compare tiangan -> ten gods
func CompareGan(self, other int) int {
	selfFE := self / 2
	otherFE := other / 2
	selfYY := self % 2
	otherYY := other % 2
	diff := 0
	if selfYY != otherYY {
		diff = 1
	}

	compareFE := CompareFiveElements(selfFE, otherFE)
	switch compareFE {
	case FiveElementEqual:
		return diff
	case FiveElementKill:
		return 4 + diff
	case FiveElementKilled:
		return 6 + diff
	case FiveElementBirth:
		return 2 + diff
	case FiveElementBirthed:
		return 8 + diff
	}

	return 0
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
