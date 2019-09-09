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
 * @file apis.go
 * @package main
 * @author Dr.NP <np@corp.herewetech.com>
 * @since 05/20/2019
 */

package main

import (
	"errors"
	"net/url"
	"strconv"
	"unicode/utf8"

	"github.com/herewetech/yixuan_naming/unihan"

	"github.com/valyala/fasthttp"
)

/*
func apiSolar(ctx *fasthttp.RequestCtx) {
	defer func() {
		if x := recover(); x != nil {
			ctx.SetUserValue("_envelope_code", 20011)
			ctx.SetUserValue("_envelope_message", "Timestamp overflow")
			ctx.SetUserValue("_envelope_data", nil)
		}
	}()

	timestamp, _ := strconv.ParseInt(ctx.UserValue("time").(string), 10, 64)
	c := calendar.ByTimestamp(timestamp)
	type retSolar struct {
		Animal    string `json:"animal"`
		Year      int64  `json:"year"`
		Month     int64  `json:"month"`
		Day       int64  `json:"day"`
		Hour      int64  `json:"hour"`
		Minute    int64  `json:"minute"`
		Second    int64  `json:"second"`
		Week      int64  `json:"week"`
		WeekAlias string `json:"week_alias"`
		Leap      bool   `json:"leap"`
	}

	r := &retSolar{
		Animal:    c.Solar.Animal().Alias(),
		Year:      c.Solar.GetYear(),
		Month:     c.Solar.GetMonth(),
		Day:       c.Solar.GetDay(),
		Hour:      c.Solar.GetHour(),
		Minute:    c.Solar.GetMinute(),
		Second:    c.Solar.GetSecond(),
		Week:      c.Solar.WeekNumber(),
		WeekAlias: c.Solar.WeekAlias(),
		Leap:      c.Solar.IsLeep(),
	}

	ctx.SetUserValue("_envelope_data", r)

	return
}

func apiLunar(ctx *fasthttp.RequestCtx) {
	defer func() {
		if x := recover(); x != nil {
			ctx.SetUserValue("_envelope_code", 20011)
			ctx.SetUserValue("_envelope_message", "Timestamp overflow")
			ctx.SetUserValue("_envelope_data", nil)
		}
	}()

	timestamp, _ := strconv.ParseInt(ctx.UserValue("time").(string), 10, 64)
	c := calendar.ByTimestamp(timestamp)
	type retLunar struct {
		Animal     string `json:"animal"`
		Year       int64  `json:"year"`
		YearAlias  string `json:"year_str"`
		Month      int64  `json:"month"`
		MonthAlias string `json:"month_str"`
		Day        int64  `json:"day"`
		DayAlias   string `json:"day_str"`
		Leap       bool   `json:"leap"`
		LeapMonth  bool   `json:"leap_month"`
	}

	r := &retLunar{
		Animal:     c.Lunar.Animal().Alias(),
		Year:       c.Lunar.GetYear(),
		Month:      c.Lunar.GetMonth(),
		Day:        c.Lunar.GetDay(),
		YearAlias:  c.Lunar.YearAlias(),
		MonthAlias: c.Lunar.MonthAlias(),
		DayAlias:   c.Lunar.DayAlias(),
		Leap:       c.Lunar.IsLeap(),
		LeapMonth:  c.Lunar.IsLeapMonth(),
	}

	ctx.SetUserValue("_envelope_data", r)

	return
}

func apiGanzhi(ctx *fasthttp.RequestCtx) {
	defer func() {
		if x := recover(); x != nil {
			ctx.SetUserValue("_envelope_code", 20011)
			ctx.SetUserValue("_envelope_message", "Timestamp overflow")
			ctx.SetUserValue("_envelope_data", nil)
		}
	}()

	timestamp, _ := strconv.ParseInt(ctx.UserValue("time").(string), 10, 64)
	c := calendar.ByTimestamp(timestamp)
	type retGanzhi struct {
		Animal     string `json:"animal"`
		Year       string `json:"year"`
		YearOrder  int64  `json:"year_order"`
		Month      string `json:"month"`
		MonthOrder int64  `json:"month_order"`
		Day        string `json:"day"`
		DayOrder   int64  `json:"day_order"`
		Hour       string `json:"hour"`
		HourOrder  int64  `json:"hour_order"`
	}

	r := &retGanzhi{
		Animal:     c.Ganzhi.Animal().Alias(),
		Year:       c.Ganzhi.YearGanzhiAlias(),
		Month:      c.Ganzhi.MonthGanzhiAlias(),
		Day:        c.Ganzhi.DayGanzhiAlias(),
		Hour:       c.Ganzhi.HourGanzhiAlias(),
		YearOrder:  c.Ganzhi.YearGanzhiOrder(),
		MonthOrder: c.Ganzhi.MonthGanzhiOrder(),
		DayOrder:   c.Ganzhi.DayGanzhiOrder(),
		HourOrder:  c.Ganzhi.HourGanzhiOrder(),
	}

	ctx.SetUserValue("_envelope_data", r)

	return
}
*/

func getRune(mode, input string) (rune, error) {
	var (
		u    int64
		r    rune
		size int
		err  error
	)

	switch mode {
	case "u":
		u, err = strconv.ParseInt(input, 16, 32)
		r = rune(u)
	case "d":
		u, err = strconv.ParseInt(input, 10, 32)
		r = rune(u)
	case "c":
		u, _ := url.QueryUnescape(input)
		r, size = utf8.DecodeRuneInString(u)
		if size == 0 {
			err = errors.New("Invalid input")
		}
	default:
		err = errors.New("Invalid mode")
	}

	return r, err
}

func apiUnihan(ctx *fasthttp.RequestCtx) {
	var (
		mode  = ctx.UserValue("mode").(string)
		input = ctx.UserValue("input").(string)
		r     rune
		h     *unihan.HanCharacter
		err   error
	)

	r, err = getRune(mode, input)
	h, err = unihan.Query(r)
	if err != nil || h == nil {
		ctx.SetUserValue("_envelope_code", 10404)
		ctx.SetUserValue("_envelope_message", "Unihan does not exists")
		ctx.SetStatusCode(fasthttp.StatusNotFound)

		return
	}

	ctx.SetUserValue("_envelope_data", h)

	return
}

func apiStroke(ctx *fasthttp.RequestCtx) {
	var (
		mode  = ctx.UserValue("mode").(string)
		input = ctx.UserValue("input").(string)
		r     rune
		h     *unihan.HanCharacter
		err   error
	)

	r, err = getRune(mode, input)
	h, err = unihan.Query(r)
	if err != nil || h == nil {
		ctx.SetUserValue("_envelope_code", 10404)
		ctx.SetUserValue("_envelope_message", "Unihan does not exists")
		ctx.SetStatusCode(fasthttp.StatusNotFound)

		return
	}

	ks, us, s, _ := h.QueryStroke()

	ctx.SetUserValue("_envelope_data", map[string]int{
		"stroke":         s,
		"unicode_stroke": us,
		"kangxi_stroke":  ks,
	})

	return
}

func apiTraditional(ctx *fasthttp.RequestCtx) {
	var (
		mode  = ctx.UserValue("mode").(string)
		input = ctx.UserValue("input").(string)
		r     rune
		h     *unihan.HanCharacter
		err   error
	)

	r, err = getRune(mode, input)
	h, err = unihan.Query(r)
	if err != nil || h == nil {
		ctx.SetUserValue("_envelope_code", 10404)
		ctx.SetUserValue("_envelope_message", "Unihan does not exists")
		ctx.SetStatusCode(fasthttp.StatusNotFound)

		return
	}

	tsi, _ := h.QueryTraditional()

	ctx.SetUserValue("_envelope_data", tsi)
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
