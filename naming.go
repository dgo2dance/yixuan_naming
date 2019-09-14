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
 * @file naming.go
 * @package main
 * @author Dr.NP <np@corp.herewetech.com>
 * @since 05/20/2019
 */

package main

import (
	"yixuan_naming/common"
	"yixuan_naming/dict"
	"yixuan_naming/list"
	"yixuan_naming/name"
	"yixuan_naming/poetry"
	"yixuan_naming/texts"
	"yixuan_naming/unihan"
)

const (
	// DefaultHTTPListenAddr : String value of default HTTP server address
	DefaultHTTPListenAddr string = "127.0.0.1:7788"
	// DefaultLibraryPath : String value of default storage path locally
	DefaultLibraryPath string = "/usr/share/naming"
	// DefaultLanguage : Default message language
	DefaultLanguage = texts.LanguageSimplified
)

var (
	g = common.NewRuntime("naming")
	s = common.NewHTTPServer(g)
)

func main() {
	g.Config.SetDefault("HTTP_Listen_Address", DefaultHTTPListenAddr)
	g.Config.SetDefault("Library_Path", DefaultLibraryPath)
	g.Config.SetDefault("Default_language", DefaultLanguage)
	texts.LanguageDefault = g.Config.GetInt("Default_language")

	g.Logger.Printf("Start server")

	var (
		lines, linePoetries, lineWords int
		chars                          int
		err                            error
	)

	// Unihan dictionary
	lines, chars, err = unihan.LoadUnihanLibraries(g.Config.GetString("Library_Path"))
	if err != nil {
		g.Logger.Fatal(err)
	} else {
		g.Logger.Printf("Load %d lines from unihan libraries, %d characters total", lines, chars)
	}

	// Unihan common characters
	lines, err = list.LoadCommonL1(g.Config.GetString("Library_Path"))
	if err != nil {
		g.Logger.Fatal(err)
	} else {
		g.Logger.Printf("Load %d lines from common character list to common characters L1", lines)
	}

	lines, err = list.LoadCommonL2(g.Config.GetString("Library_Path"))
	if err != nil {
		g.Logger.Fatal(err)
	} else {
		g.Logger.Printf("Load %d lines from common character list to common characters L2", lines)
	}

	lines = list.PrepareCommonCharacters()
	g.Logger.Printf("Prepared %d characters from common character lists", lines)

	// TraditionalSpecial
	lines, err = list.LoadTraditionalSpecial(g.Config.GetString("Library_Path"))
	if err != nil {
		g.Logger.Fatal(err)
	} else {
		g.Logger.Printf("Load %d lines from traditional special list", lines)
	}

	// PinyinSpecial
	lines, err = list.LoadPinyinSpecial(g.Config.GetString("Library_Path"))
	if err != nil {
		g.Logger.Fatal(err)
	} else {
		g.Logger.Printf("Load %d lines from pinyin special list", lines)
	}

	// Common words
	lines, err = list.LoadCommon(g.Config.GetString("Library_Path"))
	if err != nil {
		g.Logger.Fatal(err)
	} else {
		g.Logger.Printf("Load %d lines from common words", lines)
	}

	// Sensitive words
	lines, err = list.LoadSensitive(g.Config.GetString("Library_Path"))
	if err != nil {
		g.Logger.Fatal(err)
	} else {
		g.Logger.Printf("Load %d lines from sensitive words", lines)
	}

	// Common names
	lines, err = list.LoadCommonNames(g.Config.GetString("Library_Path"))
	if err != nil {
		g.Logger.Fatal(err)
	} else {
		g.Logger.Printf("Load %d lines from common names", lines)
	}

	// Common names
	lines, err = list.LoadFamilyNames(g.Config.GetString("Library_Path"))
	if err != nil {
		g.Logger.Fatal(err)
	} else {
		g.Logger.Printf("Load %d lines from family names", lines)
	}

	// Character five elements
	lines, err = list.LoadCharacterFiveElements(g.Config.GetString("Library_Path"))
	if err != nil {
		g.Logger.Fatal(err)
	} else {
		g.Logger.Printf("Load %d lines from character five elements list", lines)
	}

	// BaiJiaXing
	lines, err = list.LoadBaiJiaXing(g.Config.GetString("Library_Path"))
	if err != nil {
		g.Logger.Fatal(err)
	} else {
		g.Logger.Printf("Load %d lines from baijiaxing", lines)
	}

	// Character dictionaries
	lines, err = dict.LoadXinhua(g.Config.GetString("Library_Path"))
	if err != nil {
		g.Logger.Fatal(err)
	} else {
		g.Logger.Printf("Load %d lines from dictionary Xinhua", lines)
	}

	lines, err = dict.LoadFolkways(g.Config.GetString("Library_Path"))
	if err != nil {
		g.Logger.Fatal(err)
	} else {
		g.Logger.Printf("Load %d lines from dictionary Folkways", lines)
	}

	// Messages
	lines, err = texts.LoadMessages(g.Config.GetString("Library_Path"))
	if err != nil {
		g.Logger.Fatal(err)
	} else {
		g.Logger.Printf("Load %d lines from message files", lines)
	}

	// Poetries
	linePoetries, lineWords, err = poetry.LoadPoetries(g.Config.GetString("Library_Path"))
	if err != nil {
		g.Logger.Fatal(err)
	} else {
		g.Logger.Printf("Load %d poetries, %d words", linePoetries, lineWords)
	}

	err = name.FillRankTable()
	if err != nil {
		g.Logger.Fatal(err)
	} else {
		g.Logger.Println("Rank table filled")
	}

	svc(s)
	s.Start()
	g.Wait()

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
