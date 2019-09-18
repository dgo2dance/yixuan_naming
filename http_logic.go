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
 * @file http_logic.go
 * @package main
 * @author Dr.NP <np@corp.herewetech.com>
 * @since 05/20/2019
 */

package main

import (
	"fmt"
	"strconv"
	"time"
	"unicode/utf8"

	"yixuan_naming/common"
	"yixuan_naming/name"
	"yixuan_naming/texts"
	"yixuan_naming/utils"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/valyala/fasthttp"
)

func index(ctx *fasthttp.RequestCtx) {
	ctx.SetUserValue("_envelope_data", "Hello naming")

	return
}

func clientSign(ctx *fasthttp.RequestCtx) {
	r := ctx.UserValue("_g").(*common.GlobalRuntime)
	claims := &common.JwtClaims{}
	claims.UserType = ctx.UserValue("client_type").(string)
	claims.ExpiresAt = time.Now().Add(time.Minute * 5).Unix()
	key := []byte(r.Config.GetString("jwt_key"))
	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := tkn.SignedString(key)
	if err != nil {
		fmt.Print(err.Error())
		ctx.Error("Sign JWT failed", fasthttp.StatusInternalServerError)
	} else {
		ctx.SetUserValue("_envelope_data", tokenString)
	}

	return
}

func clientInfo(ctx *fasthttp.RequestCtx) {
	ctx.SetUserValue("_envelope_data", ctx.UserValue("jwt_claims"))

	return
}

func nameRank(ctx *fasthttp.RequestCtx) {
	var (
		args            = ctx.QueryArgs()
		familyName      []byte
		familyNameRunes []rune
		middleName      []byte
		middleNameRunes []rune
		givenName       []byte
		givenNameRunes  []rune
		birthTime       int64
		longitude       float64
		latitude        float64
		language        []byte
		languageCode    int
	)

	familyName = args.Peek("family")
	if familyName != nil {
		for len(familyName) > 0 {
			r, size := utf8.DecodeRune(familyName)
			if size > 0 && r != utf8.RuneError {
				familyNameRunes = append(familyNameRunes, r)
			} else {
				break
			}

			familyName = familyName[size:]
		}
	}

	middleName = args.Peek("middle")
	if middleName != nil {
		for len(middleName) > 0 {
			r, size := utf8.DecodeRune(middleName)
			if size > 0 && r != utf8.RuneError {
				middleNameRunes = append(middleNameRunes, r)
			} else {
				break
			}

			middleName = middleName[size:]
		}
	}

	givenName = args.Peek("given")
	if givenName != nil {
		for len(givenName) > 0 {
			r, size := utf8.DecodeRune(givenName)
			if size > 0 && r != utf8.RuneError {
				givenNameRunes = append(givenNameRunes, r)
			} else {
				break
			}

			givenName = givenName[size:]
		}
	}

	b := args.Peek("birth")
	if b != nil {
		birthTime, _ = strconv.ParseInt(string(b), 10, 64)
	}

	longitude = args.GetUfloatOrZero("longitude")
	latitude = args.GetUfloatOrZero("latitude")
	language = args.Peek("lang")
	languageCode = texts.AssertLanguage(string(language))

	r := ctx.UserValue("_g").(*common.GlobalRuntime)
	r.Logger.Printf("Name rank from %s with name <%v.%v.%v>, birth timestamp <%d>, location <%f:%f>, language <%d>",
		ctx.RemoteIP().String(),
		familyNameRunes,
		middleNameRunes,
		givenNameRunes,
		birthTime,
		latitude,
		longitude,
		languageCode)
	n := name.NewNameRunes(familyNameRunes, middleNameRunes, givenNameRunes)
	n.Normalize()
	ret, _ := name.Rank(languageCode, n, birthTime, utils.Location{Latitude: latitude, Longitude: longitude})

	ctx.SetUserValue("_envelope_data", ret)

	return
}

func nameKirsen(ctx *fasthttp.RequestCtx) {
	var (
		args            = ctx.QueryArgs()
		familyName      []byte
		familyNameRunes []rune
		prefixName      []byte
		prefixNameRunes []rune
		suffixName      []byte
		suffixNameRunes []rune
		birthTime       int64
		longitude       float64
		latitude        float64
		givenNameLength int
		gender          int
		queryNums       int
		characterLevel  int
		language        []byte
		languageCode    int
	)

	familyName = args.Peek("family")
	if familyName != nil {
		for len(familyName) > 0 {
			r, size := utf8.DecodeRune(familyName)
			if size > 0 && r != utf8.RuneError {
				familyNameRunes = append(familyNameRunes, r)
			} else {
				break
			}

			familyName = familyName[size:]
		}
	}

	prefixName = args.Peek("prefix")
	if prefixName != nil {
		for len(prefixName) > 0 {
			r, size := utf8.DecodeRune(prefixName)
			if size > 0 && r != utf8.RuneError {
				prefixNameRunes = append(prefixNameRunes, r)
			} else {
				break
			}

			prefixName = prefixName[size:]
		}
	}

	suffixName = args.Peek("suffix")
	if suffixName != nil {
		for len(suffixName) > 0 {
			r, size := utf8.DecodeRune(suffixName)
			if size > 0 && r != utf8.RuneError {
				suffixNameRunes = append(suffixNameRunes, r)
			} else {
				break
			}

			suffixName = suffixName[size:]
		}
	}

	b := args.Peek("birth")
	if b != nil {
		birthTime, _ = strconv.ParseInt(string(b), 10, 64)
	}

	longitude = args.GetUfloatOrZero("longitude")
	latitude = args.GetUfloatOrZero("latitude")
	givenNameLength = args.GetUintOrZero("length")
	gender = args.GetUintOrZero("gender")
	if gender != utils.GenderFemale && gender != utils.GenderMale {
		gender = utils.GenderUnknown
	}

	queryNums = args.GetUintOrZero("nums")
	language = args.Peek("lang")
	languageCode = texts.AssertLanguage(string(language))
	characterLevel = args.GetUintOrZero("character_level")

	if 0 == longitude && 0 == latitude {
		longitude = 120.0
		latitude = 45.0
	}

	r := ctx.UserValue("_g").(*common.GlobalRuntime)
	conditions := &name.KirsenConditions{
		FamilyNameRunes: familyNameRunes,
		PrefixNameRunes: prefixNameRunes,
		SuffixNameRunes: suffixNameRunes,
		Gender:          gender,
		NeedMiddleName:  false,
		NeedBirthTime:   false,
		GivenNameLength: givenNameLength,
		QueryNums:       queryNums,
		CharacterLevel:  characterLevel,
	}

	conditions.Traditionalize()
	r.Logger.Printf("Name kirsen from %s with family name <%v>, prefix <%v> and suffix <%v>, birth timestamp <%d>, location <%f:%f>, given name length <%d>, gender <%d>, query numbers <%d>, language <%d>",
		ctx.RemoteIP().String(),
		conditions.FamilyNameRunes,
		conditions.PrefixNameRunes,
		conditions.SuffixNameRunes,
		birthTime,
		latitude,
		longitude,
		givenNameLength,
		gender,
		queryNums,
		languageCode)

	ret, _ := name.Kirsen(languageCode, conditions, birthTime, utils.Location{Latitude: latitude, Longitude: longitude})
	ctx.SetUserValue("_envelope_data", ret)

	return
}

// HTTP CORS Options request
func cors(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	ctx.Response.Header.Set("Access-Control-Allow-Headers", "Content-Type")
	ctx.Response.Header.Set("Access-Control-Max-Age", "86400")

	return
}

// Full-stack routes
func f(c fasthttp.RequestHandler, authorization string, s *common.HTTPServer) fasthttp.RequestHandler {
	return common.HTTPEnvelope(
		common.HTTPGlobalRuntime(
			common.HTTPAuthorization(c, authorization), s.Runtime))
}

func taskCommonChars(ctx *fasthttp.RequestCtx) {
	var (
		args  = ctx.QueryArgs()
		level int
	)

	level = args.GetUintOrZero("level")
	if level != 2 {
		level = 1
	}

	name.CalcCommonStrokes(level)

	return
}

// Logic routers
func svc(s *common.HTTPServer) {
	s.Router.GET("/", f(index, "none", s))
	s.Router.GET("/client_sign/:client_type", f(clientSign, "basic", s))
	s.Router.GET("/client_info", f(clientInfo, "jwt", s))
	s.Router.OPTIONS("/*_all", cors)

	s.Router.GET("/api/unihan/:mode/:input", f(apiUnihan, "none", s))
	s.Router.GET("/api/stroke/:mode/:input", f(apiStroke, "none", s))
	s.Router.GET("/api/traditional/:mode/:input", f(apiTraditional, "none", s))

	// Logics
	s.Router.GET("/name/rank", f(nameRank, "none", s))
	s.Router.GET("/name/kirsen", f(nameKirsen, "none", s))

	// Tasks
	s.Router.GET("/task/common_chars_length", taskCommonChars)

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
