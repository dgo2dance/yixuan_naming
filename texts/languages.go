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
 * @file texts.go
 * @package texts
 * @author Dr.NP <np@corp.herewetech.com>
 * @since 07/16/2019
 */

package texts

import "strings"

const (
	// LanguageSimplified : Simplified Chinese
	LanguageSimplified = iota
	// LanguageTraditional : Traditional Chinese
	LanguageTraditional
	// LanguageEnglish : English
	LanguageEnglish
)

// LanguageDefault : Equal to LangaugeSimplified
var LanguageDefault int

// AssertLanguage : Determine language code by given string
func AssertLanguage(lang string) int {
	switch strings.ToLower(lang) {
	case "0", "s", "simplified":
		return LanguageSimplified
	case "1", "t", "traditional":
		return LanguageTraditional
	case "2", "e", "english":
		return LanguageEnglish
	default:
		return LanguageDefault
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
