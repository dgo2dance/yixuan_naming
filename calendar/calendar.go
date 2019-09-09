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
 * @file calendar.go
 * @package calendar
 * @author Dr.NP <np@corp.herewetech.com>
 * @since 06/06/2019
 */

package calendar

import (
	"time"

	"yixuan_naming/utils"
)

type timeSpec struct {
	Year   int    `json:"year"`
	Month  int    `json:"month"`
	Day    int    `json:"day"`
	Hour   int    `json:"hour"`
	Minute int    `json:"minute"`
	Second int    `json:"second"`
	Zone   string `json:"zone"`
	Offset int    `json:"offset"`
}

func (s *timeSpec) parse(t time.Time) {
	s.Year = t.Year()
	s.Month = int(t.Month())
	s.Day = t.Day()
	s.Hour, s.Minute, s.Second = t.Clock()
	s.Zone, s.Offset = t.Zone()

	return
}

// Calendar : Main struct
type Calendar struct {
	t           time.Time
	Location    utils.Location `json:"location"`
	GeneralTime timeSpec       `json:"general_time"`
	UTCTime     timeSpec       `json:"utc_time"`
	ChinaTime   timeSpec       `json:"china_time"`
	LocalTime   timeSpec       `json:"local_time"`
	RealTime    timeSpec       `json:"real_time"`
	Solar       solar          `json:"solar"`
	Lunar       lunar          `json:"lunar"`
	Ganzhi      ganzhi         `json:"ganzhi"`
}

// New : Create new calendar
func New(timestamp int64, loc utils.Location) *Calendar {
	ret := &Calendar{
		t:        time.Unix(timestamp, 0),
		Location: loc,
	}

	var (
		tChina = ret.China()
		tReal  = ret.Real()
	)

	ret.GeneralTime.parse(ret.General())
	ret.UTCTime.parse(ret.UTC())
	ret.ChinaTime.parse(tChina)
	ret.LocalTime.parse(ret.Local())
	ret.RealTime.parse(tReal)

	ret.Solar = solar{t: &tChina}
	ret.Lunar = lunar{t: &tChina}
	ret.Ganzhi = ganzhi{t: &tReal}

	ret.Solar.complete()
	ret.Lunar.complete()
	ret.Ganzhi.complete()

	return ret
}

// General : Genaral time
func (c *Calendar) General() time.Time {
	return c.t
}

// UTC : UTC time
func (c *Calendar) UTC() time.Time {
	return c.t.UTC()
}

// China : CST/CDT time
func (c *Calendar) China() time.Time {
	l, _ := time.LoadLocation("Asia/Shanghai")
	return c.t.In(l)
}

// Local : Local time
func (c *Calendar) Local() time.Time {
	offset := int(c.Location.Longitude * 240)
	l := time.FixedZone("LocalTime", offset)
	return c.t.In(l)
}

var (
	realSunFixLeap = []int{
		0,
		-189, -218, -246, -273, -301, -327, -354, -380, -405, -430,
		-455, -479, -502, -525, -547, -568, -589, -609, -628, -647,
		-665, -682, -698, -714, -728, -742, -755, -779, -790, -799,
		-817, -824, -830, -836, -841, -845, -849, -851, -853, -854,
		-855, -854, -853, -851, -848, -845, -841, -836, -831, -824,
		-818, -810, -802, -793, -664, -774, -763, -752, -741, -728,
		-716, -703, -689, -675, -661, -647, -632, -616, -601, -585,
		-568, -552, -535, -518, -501, -484, -466, -449, -431, -413,
		-395, -377, -358, -340, -322, -304, -285, -267, -249, -231,
		-213, -196, -178, -161, -144, -127, -110, -93, -77, -61,
		46, 30, 16, 1, 13, 27, 41, 54, 66, 79,
		91, 102, 113, 124, 134, 143, 153, 161, 169, 177,
		184, 70, 196, 201, 206, 210, 217, 216, 219, 220,
		222, 222, 222, 222, 221, 219, 217, 214, 211, 207,
		203, 198, 193, 187, 181, 174, 167, 159, 151, 142,
		133, 124, 114, 104, 94, 83, 72, 60, 48, 36,
		24, 12, 1, 14, 39, 52, -65, -78, -91, -105,
		-117, -130, -143, -156, -168, -181, -193, -205, -217, -229,
		-240, -251, -262, -273, -283, -293, -302, -311, -320, -328,
		-336, -343, -350, -356, -362, -368, -372, -376, -380, -383,
		-385, -387, -389, -389, -389, -389, -388, -386, -384, -381,
		-377, -373, -368, -363, -357, -351, -344, -336, -328, -319,
		-310, -300, -290, -279, -267, -255, -242, -229, -216, -201,
		-187, -171, -156, -140, -123, -107, -89, -72, 54, 35,
		17, 2, 21, 41, 60, 80, 100, 121, 141, 162,
		183, 183, 204, 225, 246, 267, 288, 310, 331, 353,
		374, 395, 417, 438, 459, 480, 501, 522, 542, 562,
		582, 602, 621, 640, 659, 678, 696, 696, 713, 731,
		748, 764, 780, 796, 796, 811, 825, 839, 853, 866,
		878, 890, 901, 912, 921, 931, 940, 948, 955, 961,
		967, 972, 976, 980, 982, 984, 985, 985, 984, 983,
		981, 977, 973, 969, 963, 956, 949, 941, 932, 922,
		911, 900, 887, 874, 860, 846, 830, 814, 797, 779,
		760, 741, 721, 700, 678, 656, 633, 609, 585, 561,
		535, 509, 483, 456, 429, 402, 374, 346, 317, 288,
		259, 230, 201, 171, 142, 112, 82, 52, 23, 7,
		37, -66, -96, -125, -154, -183,
	}
	realSunFixNormal = []int{
		0,
		-189, -218, -246, -273, -301, -327, -354, -380, -405, -430,
		-455, -479, -502, -525, -547, -568, -589, -609, -628, -647,
		-665, -682, -698, -714, -728, -742, -755, -779, -790, -799,
		-817, -824, -830, -836, -841, -845, -849, -851, -853, -854,
		-855, -854, -853, -851, -848, -845, -841, -836, -831, -824,
		-818, -810, -802, -793, -664, -774, -763, -752, -741,
		-716, -703, -689, -675, -661, -647, -632, -616, -601, -585,
		-568, -552, -535, -518, -501, -484, -466, -449, -431, -413,
		-395, -377, -358, -340, -322, -304, -285, -267, -249, -231,
		-213, -196, -178, -161, -144, -127, -110, -93, -77, -61,
		46, 30, 16, 1, 13, 27, 41, 54, 66, 79,
		91, 102, 113, 124, 134, 143, 153, 161, 169, 177,
		184, 70, 196, 201, 206, 210, 217, 216, 219, 220,
		222, 222, 222, 222, 221, 219, 217, 214, 211, 207,
		203, 198, 193, 187, 181, 174, 167, 159, 151, 142,
		133, 124, 114, 104, 94, 83, 72, 60, 48, 36,
		24, 12, 1, 14, 39, 52, -65, -78, -91, -105,
		-117, -130, -143, -156, -168, -181, -193, -205, -217, -229,
		-240, -251, -262, -273, -283, -293, -302, -311, -320, -328,
		-336, -343, -350, -356, -362, -368, -372, -376, -380, -383,
		-385, -387, -389, -389, -389, -389, -388, -386, -384, -381,
		-377, -373, -368, -363, -357, -351, -344, -336, -328, -319,
		-310, -300, -290, -279, -267, -255, -242, -229, -216, -201,
		-187, -171, -156, -140, -123, -107, -89, -72, 54, 35,
		17, 2, 21, 41, 60, 80, 100, 121, 141, 162,
		183, 183, 204, 225, 246, 267, 288, 310, 331, 353,
		374, 395, 417, 438, 459, 480, 501, 522, 542, 562,
		582, 602, 621, 640, 659, 678, 696, 696, 713, 731,
		748, 764, 780, 796, 796, 811, 825, 839, 853, 866,
		878, 890, 901, 912, 921, 931, 940, 948, 955, 961,
		967, 972, 976, 980, 982, 984, 985, 985, 984, 983,
		981, 977, 973, 969, 963, 956, 949, 941, 932, 922,
		911, 900, 887, 874, 860, 846, 830, 814, 797, 779,
		760, 741, 721, 700, 678, 656, 633, 609, 585, 561,
		535, 509, 483, 456, 429, 402, 374, 346, 317, 288,
		259, 230, 201, 171, 142, 112, 82, 52, 23, 7,
		37, -66, -96, -125, -154, -183,
	}
)

// Real : Real-sun time
func (c *Calendar) Real() time.Time {
	// Find index
	leap := false
	y := c.t.Year()
	if y%400 == 0 {
		// Leap
		leap = true
	} else if y%4 == 0 && y%100 != 0 {
		// Leap
		leap = true
	}

	d := c.t.YearDay()
	fix := 0
	if leap {
		fix = realSunFixLeap[d]
	} else {
		fix = realSunFixNormal[d]
	}

	offset := int(c.Location.Longitude*240) + fix
	l := time.FixedZone("RealSunTime", offset)

	return c.t.In(l)
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
