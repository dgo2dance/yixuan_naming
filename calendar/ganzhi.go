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
 * @package calendar
 * @author Dr.NP <np@corp.herewetech.com>
 * @since 06/12/2019
 */

package calendar

import (
	"time"

	"github.com/herewetech/yixuan_naming/utils"
)

type ganzhi struct {
	t           *time.Time
	YearOrder   int              `json:"year_order"`
	Year        utils.GanzhiPair `json:"year"`
	YearString  string           `json:"year_alias"`
	Month       utils.GanzhiPair `json:"month"`
	MonthString string           `json:"month_alias"`
	Day         utils.GanzhiPair `json:"day"`
	DayString   string           `json:"day_alias"`
	Hour        utils.GanzhiPair `json:"hour"`
	HourString  string           `json:"hour_alias"`
}

// GetSolartermsGanzhi : Get solarterms from LiChun to DaHan for a whole ganzhi year
func GetSolartermsGanzhi(year int) []time.Time {
	solartermsCurr := GetSolarterms(year)
	solartermsNext := GetSolarterms(year + 1)
	ret := make([]time.Time, 24)
	copy(ret, solartermsCurr[2:])
	ret[22] = solartermsNext[0]
	ret[23] = solartermsNext[1]

	return ret
}

func (g *ganzhi) complete() {
	var (
		idx, year, month       int
		solarterm, spring      time.Time
		solarterms             []time.Time
		dH, dD, dM, dY, dC, dI int
	)

	year = g.t.Year()
	solarterms = GetSolarterms(year)
	spring = solarterms[2]
	if g.t.Before(spring) {
		year--
	}

	g.YearOrder = year
	solarterms = GetSolartermsGanzhi(year)

	g.Year.TianGan = (year - 4) % 10
	g.Year.DiZhi = (year - 4) % 12
	//g.YearString = g.Year.String()

	for idx, solarterm = range solarterms {
		if g.t.Before(solarterm) {
			idx--
			break
		}
	}

	month = (idx / 2) + 1
	g.Month.TianGan = (month + g.Year.TianGan*2 + 1) % 10
	g.Month.DiZhi = (month + 1) % 12
	//g.MonthString = g.Month.String()

	year = g.t.Year()
	dM = int(g.t.Month())
	if dM < 3 {
		dM += 12
		year--
	}
	dY = year % 100
	dC = year / 100
	dD = g.t.Day()
	if dM%2 == 0 {
		dI = 6
	}
	dH = g.t.Hour()
	if dH >= 23 {
		dD++
	}

	g.Day.TianGan = ((4 * dC) + (dC / 4) + (5 * dY) + (dY / 4) + ((3 * (dM + 1)) / 5) + dD - 4) % 10
	g.Day.DiZhi = ((8 * dC) + (dC / 4) + (5 * dY) + (dY / 4) + ((3 * (dM + 1)) / 5) + dD + 6 + dI) % 12
	//g.DayString = g.Day.String()

	g.Hour.DiZhi = ((g.t.Hour() + 1) / 2) % 12
	g.Hour.TianGan = (g.Hour.DiZhi + g.Day.TianGan*2) % 10
	//g.HourString = g.Hour.String()

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
