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
 * @file aliases.go
 * @package texts
 * @author Dr.NP <np@corp.herewetech.com>
 * @since 07/17/2019
 */

package texts

const (
	// AliasRank : 0
	AliasRank = iota
	// AliasEightTrigram : 1
	AliasEightTrigram
	// AliasEightTrigramNature : 2
	AliasEightTrigramNature
	// AliasFiveElement : 3
	AliasFiveElement
	// AliasFiveElementCompare : 4
	AliasFiveElementCompare
	// AliasGan : 5
	AliasGan
	// AliasZhi : 6
	AliasZhi
	// AliasGanFiveElement : 7
	AliasGanFiveElement
	// AliasTenGod : 8
	AliasTenGod
	// AliasTenGodRepresentative : 9
	AliasTenGodRepresentative
	// AliasYinYang : 10
	AliasYinYang
	// AliasAnimal : 11
	AliasAnimal
	// AliasLunarMonth : 12
	AliasLunarMonth
	// AliasLunarDay : 13
	AliasLunarDay
	// AliasSolarterm : 14
	AliasSolarterm
	// AliasSoundFiveElement : 15
	AliasSoundFiveElement
)

// Aliases
var (
	rankAliases = [][]string{
		{"非运", "大凶", "凶", "半吉", "吉", "大吉"},
		{"非運", "大兇", "兇", "半吉", "吉", "大吉"},
	}
	eightTrigramAliases = [][]string{
		{"坤", "艮", "坎", "巽", "震", "离", "兑", "乾"},
		{"坤", "艮", "坎", "巽", "震", "離", "兌", "乾"},
	}
	eightTrigramNatureAliases = [][]string{
		{"地", "山", "水", "风", "雷", "火", "泽", "天"},
		{"地", "山", "水", "風", "雷", "火", "澤", "天"},
	}
	fiveElementAliases = [][]string{
		{"木", "火", "土", "金", "水"},
		{"木", "火", "土", "金", "水"},
	}
	fiveElementCompareAliases = [][]string{
		{"比劫", "食伤", "才财", "杀官", "枭印"},
		{"比劫", "食傷", "才財", "殺官", "梟印"},
	}
	ganAliases = [][]string{
		{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"},
		{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"},
	}
	zhiAliases = [][]string{
		{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"},
		{"子", "醜", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"},
	}
	ganFiveElementAliases = [][]string{
		{"甲木", "乙木", "丙火", "丁火", "戊土", "己土", "庚金", "辛金", "壬水", "癸水"},
		{"甲木", "乙木", "丙火", "丁火", "戊土", "己土", "庚金", "辛金", "壬水", "癸水"},
	}
	tenGodAliases = [][]string{
		{"比肩", "劫财", "食神", "伤官", "偏财", "正财", "七杀", "正官", "偏印", "正印"},
		{"比肩", "劫財", "食神", "傷官", "偏財", "正財", "七殺", "正官", "偏印", "正印"},
	}
	tenGodRpresentativeAliases = [][]string{
		{"兄弟", "兄弟", "子孙", "子孙", "妻财", "妻财", "官鬼", "官鬼", "父母", "父母"},
		{"兄弟", "兄弟", "子孫", "子孫", "妻財", "妻財", "官鬼", "官鬼", "父母", "父母"},
	}
	yinYangAliases = [][]string{
		{"阳", "阴"},
		{"陽", "陰"},
	}
	animalAliases = [][]string{
		{"鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊", "猴", "鸡", "狗", "猪"},
		{"鼠", "牛", "虎", "兔", "龍", "蛇", "馬", "羊", "猴", "雞", "狗", "豬"},
	}
	lunarMonthAliases = [][]string{
		{"正", "二", "三", "四", "五", "六", "七", "八", "九", "十", "冬", "腊", ""},
		{"正", "二", "三", "四", "五", "六", "七", "八", "九", "十", "冬", "臘", ""},
	}
	lunarDayAliases = [][]string{
		{
			"初一", "初二", "初三", "初四", "初五", "初六", "初七", "初八", "初九", "初十",
			"十一", "十二", "十三", "十四", "十五", "十六", "十七", "十八", "十九", "二十",
			"廿一", "廿二", "廿三", "廿四", "廿五", "廿六", "廿七", "廿八", "廿九", "三十",
		},
		{
			"初一", "初二", "初三", "初四", "初五", "初六", "初七", "初八", "初九", "初十",
			"十一", "十二", "十三", "十四", "十五", "十六", "十七", "十八", "十九", "二十",
			"廿一", "廿二", "廿三", "廿四", "廿五", "廿六", "廿七", "廿八", "廿九", "三十",
		},
	}
	solartermAliases = [][]string{
		{
			"小寒", "大寒", "立春", "雨水", "惊蛰", "春分",
			"清明", "谷雨", "立夏", "小满", "芒种", "夏至",
			"小暑", "大暑", "立秋", "处暑", "白露", "秋分",
			"寒露", "霜降", "立冬", "小雪", "大雪", "冬至",
		},
		{
			"小寒", "大寒", "立春", "雨水", "驚蟄", "春分",
			"清明", "穀雨", "立夏", "小滿", "芒種", "夏至",
			"小暑", "大暑", "立秋", "處暑", "白露", "秋分",
			"寒露", "霜降", "立冬", "小雪", "大雪", "冬至",
		},
	}
	soundFiveElementAliases = [][]string{
		{
			"海中金", "炉中火", "大林木", "路旁土", "剑锋金", "山头火",
			"涧下水", "城头土", "白蜡金", "杨柳木", "井泉水", "屋上土",
			"霹雳火", "松柏木", "长流水", "砂中金", "山下火", "平地木",
			"壁上土", "金箔金", "覆灯火", "天河水", "大驿土", "钗钏金",
			"桑柘木", "大溪水", "砂中土", "天上火", "石榴木", "大海水",
		},
		{
			"海中金", "爐中火", "大林木", "路旁土", "劍鋒金", "山頭火",
			"澗下水", "城頭土", "白蠟金", "楊柳木", "井泉水", "屋上土",
			"霹靂火", "松柏木", "長流水", "砂中金", "山下火", "平地木",
			"壁上土", "金箔金", "覆燈火", "天河水", "大驛土", "釵釧金",
			"桑柘木", "大溪水", "砂中土", "天上火", "石榴木", "大海水",
		},
	}
)

// GetAlias : Get aliases text
func GetAlias(alias int, index int, language int) string {
	var aliases [][]string

	switch alias {
	case AliasRank:
		aliases = rankAliases
	case AliasEightTrigram:
		aliases = eightTrigramAliases
	case AliasEightTrigramNature:
		aliases = eightTrigramNatureAliases
	case AliasFiveElement:
		aliases = fiveElementAliases
	case AliasFiveElementCompare:
		aliases = fiveElementCompareAliases
	case AliasGan:
		aliases = ganAliases
	case AliasZhi:
		aliases = zhiAliases
	case AliasGanFiveElement:
		aliases = ganFiveElementAliases
	case AliasTenGod:
		aliases = tenGodAliases
	case AliasTenGodRepresentative:
		aliases = tenGodRpresentativeAliases
	case AliasYinYang:
		aliases = yinYangAliases
	case AliasAnimal:
		aliases = animalAliases
	case AliasLunarMonth:
		aliases = lunarMonthAliases
	case AliasLunarDay:
		aliases = lunarDayAliases
	case AliasSolarterm:
		aliases = solartermAliases
	case AliasSoundFiveElement:
		aliases = soundFiveElementAliases
	}

	if aliases == nil || len(aliases) < 1 {
		return ""
	}

	if language < 0 || language >= len(aliases) {
		language = LanguageDefault
	}

	if index < 0 || index >= len(aliases[language]) {
		return ""
	}

	return aliases[language][index]
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
