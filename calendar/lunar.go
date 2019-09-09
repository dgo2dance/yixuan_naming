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
 * @file lunar.go
 * @package calendar
 * @author Dr.NP <np@corp.herewetech.com>
 * @since 06/06/2019
 */

package calendar

import (
	"time"
)

type lunarYear struct {
	year      int
	days      []int
	totalDays int
	leapMonth int
	leapType  int
}

var (
	lunarVars = []int64{
		0x04bd8, 0x04ae0, 0x0a570, 0x054d5, 0x0d260, 0x0d950, 0x16554, 0x056a0, 0x09ad0, 0x055d2, // 1900-1909
		0x04ae0, 0x0a5b6, 0x0a4d0, 0x0d250, 0x1d255, 0x0b540, 0x0d6a0, 0x0ada2, 0x095b0, 0x14977, // 1910-1919
		0x04970, 0x0a4b0, 0x0b4b5, 0x06a50, 0x06d40, 0x1ab54, 0x02b60, 0x09570, 0x052f2, 0x04970, // 1920-1929
		0x06566, 0x0d4a0, 0x0ea50, 0x06e95, 0x05ad0, 0x02b60, 0x186e3, 0x092e0, 0x1c8d7, 0x0c950, // 1930-1939
		0x0d4a0, 0x1d8a6, 0x0b550, 0x056a0, 0x1a5b4, 0x025d0, 0x092d0, 0x0d2b2, 0x0a950, 0x0b557, // 1940-1949
		0x06ca0, 0x0b550, 0x15355, 0x04da0, 0x0a5b0, 0x14573, 0x052b0, 0x0a9a8, 0x0e950, 0x06aa0, // 1950-1959
		0x0aea6, 0x0ab50, 0x04b60, 0x0aae4, 0x0a570, 0x05260, 0x0f263, 0x0d950, 0x05b57, 0x056a0, // 1960-1969
		0x096d0, 0x04dd5, 0x04ad0, 0x0a4d0, 0x0d4d4, 0x0d250, 0x0d558, 0x0b540, 0x0b6a0, 0x195a6, // 1970-1979
		0x095b0, 0x049b0, 0x0a974, 0x0a4b0, 0x0b27a, 0x06a50, 0x06d40, 0x0af46, 0x0ab60, 0x09570, // 1980-1989
		0x04af5, 0x04970, 0x064b0, 0x074a3, 0x0ea50, 0x06b58, 0x055c0, 0x0ab60, 0x096d5, 0x092e0, // 1990-1999
		0x0c960, 0x0d954, 0x0d4a0, 0x0da50, 0x07552, 0x056a0, 0x0abb7, 0x025d0, 0x092d0, 0x0cab5, // 2000-2009
		0x0a950, 0x0b4a0, 0x0baa4, 0x0ad50, 0x055d9, 0x04ba0, 0x0a5b0, 0x15176, 0x052b0, 0x0a930, // 2010-2019
		0x07954, 0x06aa0, 0x0ad50, 0x05b52, 0x04b60, 0x0a6e6, 0x0a4e0, 0x0d260, 0x0ea65, 0x0d530, // 2020-2029
		0x05aa0, 0x076a3, 0x096d0, 0x04afb, 0x04ad0, 0x0a4d0, 0x1d0b6, 0x0d250, 0x0d520, 0x0dd45, // 2030-2039
		0x0b5a0, 0x056d0, 0x055b2, 0x049b0, 0x0a577, 0x0a4b0, 0x0aa50, 0x1b255, 0x06d20, 0x0ada0, // 2040-2049
		0x14b63, 0x09370, 0x049f8, 0x04970, 0x064b0, 0x168a6, 0x0ea50, 0x06b20, 0x1a6c4, 0x0aae0, // 2050-2059
		0x0a2e0, 0x0d2e3, 0x0c960, 0x0d557, 0x0d4a0, 0x0da50, 0x05d55, 0x056a0, 0x0a6d0, 0x055d4, // 2060-2069
		0x052d0, 0x0a9b8, 0x0a950, 0x0b4a0, 0x0b6a6, 0x0ad50, 0x055a0, 0x0aba4, 0x0a5b0, 0x052b0, // 2070-2079
		0x0b273, 0x06930, 0x07337, 0x06aa0, 0x0ad50, 0x14b55, 0x04b60, 0x0a570, 0x054e4, 0x0d160, // 2080-2089
		0x0e968, 0x0d520, 0x0daa0, 0x16aa6, 0x056d0, 0x04ae0, 0x0a9d4, 0x0a2d0, 0x0d150, 0x0f252, // 2090-2099
		0x0d520, // 2100
	}

	lunarStart = time.Date(1900, 1, 30, 16, 0, 0, 0, time.UTC)
	lunarYears []*lunarYear

	animalSignAliases = []string{"鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊", "猴", "鸡", "狗", "猪", ""}
	lunarMonthAliases = []string{"正", "二", "三", "四", "五", "六", "七", "八", "九", "十", "冬", "腊", ""}
	lunarDayAliases   = []string{
		"初一", "初二", "初三", "初四", "初五", "初六", "初七", "初八", "初九", "初十",
		"十一", "十二", "十三", "十四", "十五", "十六", "十七", "十八", "十九", "二十",
		"廿一", "廿二", "廿三", "廿四", "廿五", "廿六", "廿七", "廿八", "廿九", "三十",
	}
)

func parseVar(v int64) *lunarYear {
	var (
		leapMonth int
		leapType  int
		m         int
		mt        int
		mDays     []int
		totalDays int
	)

	leapMonth = int(v & int64(0x0f))
	leapType = int(v >> 16)
	for m = 1; m <= 12; m++ {
		mt = int((v >> uint(16-m)) & int64(1))
		if mt == 1 {
			mDays = append(mDays, 30)
			totalDays += 30
		} else {
			mDays = append(mDays, 29)
			totalDays += 29
		}

		if m == leapMonth {
			if leapType > 0 {
				mDays = append(mDays, 30)
				totalDays += 30
			} else {
				mDays = append(mDays, 29)
				totalDays += 29
			}
		}
	}

	return &lunarYear{
		days:      mDays,
		totalDays: totalDays,
		leapMonth: leapMonth,
		leapType:  leapType,
	}
}

type lunar struct {
	t                *time.Time
	Year             int    `json:"year"`
	YearString       string `json:"year_alias"`
	Month            int    `json:"month"`
	MonthString      string `json:"month_alias"`
	Day              int    `json:"day"`
	DayString        string `json:"day_alias"`
	AnimalSign       int    `json:"animal_sign"`
	AnimalSignString string `json:"animal_sign_alias"`
	YearMonths       int    `json:"year_months"`
	YearDays         int    `json:"year_days"`
	LeapMonth        bool   `json:"leap_month"`
}

// complete : Calculate the full lunar
func (l *lunar) complete() {
	var (
		d          = l.t.Sub(lunarStart)
		subDays    int
		remainDays int
		m          int
		y          = lunarYears[0]
	)

	subDays = int(d.Hours()) / 24
	remainDays = subDays
	for m = 0; m <= 200; m++ {
		y = lunarYears[m]
		if y != nil {
			remainDays = subDays
			subDays -= y.totalDays
		}

		if subDays < 0 {
			break
		}
	}

	l.Year = y.year
	l.YearMonths = len(y.days)
	l.YearDays = y.totalDays
	subDays = remainDays
	for m = 0; m < len(y.days); m++ {
		remainDays = subDays
		subDays -= y.days[m]
		if subDays < 0 {
			break
		}
	}

	l.Month = m + 1
	if y.leapMonth > 0 {
		if y.leapMonth == m {
			l.LeapMonth = true
		}

		if m > y.leapMonth {
			l.Month = m
		}
	}

	if remainDays > 29 {
		// Out of range
		remainDays = 29
	}

	if remainDays < 0 {
		remainDays = 0
	}

	l.MonthString = lunarMonthAliases[l.Month-1]
	l.Day = remainDays + 1
	l.DayString = lunarDayAliases[l.Day-1]
	l.AnimalSign = (l.Year - 4) % 12
	l.AnimalSignString = animalSignAliases[l.AnimalSign]

	return
}

// Parse all vars
func init() {
	for m := 0; m <= 200; m++ {
		y := parseVar(lunarVars[m])
		y.year = 1900 + m
		lunarYears = append(lunarYears, y)
	}

	return
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
