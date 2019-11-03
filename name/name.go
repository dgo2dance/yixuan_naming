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
 * @file name.go
 * @package name
 * @author Dr.NP <np@corp.herewetech.com>
 * @since 06/30/2019
 */

package name

import (
	"fmt"
	"regexp"
	"strings"

	"yixuan_naming/list"
	"yixuan_naming/unihan"
	"yixuan_naming/utils"
)

type nameSpec struct {
	Runes        []rune                 `json:"runes"`
	Strokes      []int                  `json:"strokes,omitempty"`
	Characters   []*unihan.HanCharacter `json:"characters,omitempty"`
	FiveElements []int                  `json:"five_elements"`
	Str          string                 `json:"string"`
	Len          int                    `json:"length"`
}

type nameDef struct {
	FamilyName  nameSpec `json:"family_name"`
	MiddleName  nameSpec `json:"middle_name,omitempty"`
	GivenName   nameSpec `json:"given_name"`
	FullNameStr string   `json:"full_name"`
}

// Name : Name defination
type Name struct {
	Original    nameDef  `json:"original,omitempty"`
	Simplified  nameDef  `json:"simplified,omitempty"`
	Traditional nameDef  `json:"traditional,omitempty"`
	PinyinTone  []string `json:"pinyin_tone"`
	Pinyin      []string `json:"pinyin"`
	Rank        int      `json:"rank,omitempty"`
	IsCommon    bool     `json:"is_common"`
}

// NewName : Create name from string
func NewName(familyName, middleName, givenName string) *Name {
	name := &Name{}
	name.Original.FamilyName.Str = familyName
	name.Original.MiddleName.Str = middleName
	name.Original.GivenName.Str = givenName
	name.Original.FamilyName.Runes = []rune(familyName)
	name.Original.MiddleName.Runes = []rune(middleName)
	name.Original.GivenName.Runes = []rune(givenName)

	return name
}

// NewNameRunes : Create name from rune arrays
func NewNameRunes(familyName, middleName, givenName []rune) *Name {
	name := &Name{}
	name.Original.FamilyName.Runes = familyName
	name.Original.MiddleName.Runes = middleName
	name.Original.GivenName.Runes = givenName
	name.Original.FamilyName.Str = string(familyName)
	name.Original.MiddleName.Str = string(middleName)
	name.Original.GivenName.Str = string(givenName)

	return name
}

// assignUnihan : Assign unihan characters of name specifications
func (ns *nameSpec) assignUnihan() {
	var (
		r   rune
		c   *unihan.HanCharacter
		err error
	)

	for _, r = range ns.Runes {
		c, err = unihan.Query(r)
		if err == nil {
			// Strokes
			ns.Characters = append(ns.Characters, c)
		}
	}

	ns.Len = len(ns.Runes)
}

// assignSpec : Assign properties of name specification (known unihan)
func (ns *nameSpec) assignSpec() {
	var (
		c      *unihan.HanCharacter
		fe     int
		stroke int
	)

	ns.Runes = nil
	for _, c = range ns.Characters {
		if c != nil {
			ns.Runes = append(ns.Runes, c.Unicode)
			stroke = list.QueryStrokeSpecial(c.Unicode)
			if stroke <= 0 {
				stroke = c.QueryStrokePrefer()
			}
			ns.Strokes = append(ns.Strokes, stroke)
			fe = list.QueryFiveElement(c.Unicode)
			ns.FiveElements = append(ns.FiveElements, fe)
		}
	}

	ns.Len = len(ns.Runes)
	ns.Str = string(ns.Runes)
}

// simplified : Simplify name specifications
func (ns *nameSpec) simplify() []*unihan.HanCharacter {
	var (
		c, cs *unihan.HanCharacter
		r     rune
		ret   []*unihan.HanCharacter
	)

	for _, c = range ns.Characters {
		// Check simplified
		r, _ = c.QuerySimplifiedPrefer()
		cs, _ = unihan.Query(r)
		if cs == nil {
			cs = c
		}

		ret = append(ret, cs)
	}

	return ret
}

// traditionalized : Traditionalized name specifications
func (ns *nameSpec) traditionalized() []*unihan.HanCharacter {
	var (
		c, cs *unihan.HanCharacter
		r     rune
		ret   []*unihan.HanCharacter
	)

	for _, c = range ns.Characters {
		// Check special
		//r = special[c.Unicode]
		r = list.QueryTraditionalSpecial(c.Unicode)
		if r == 0 {
			r, _ = c.QueryTraditionalLazy()
		}
		cs, _ = unihan.Query(r)
		if cs == nil {
			cs = c
		}

		ret = append(ret, cs)
	}

	return ret
}

// Normalize : Normalize name (simplifed & traditional)
func (name *Name) Normalize() {
	name.Original.FamilyName.assignUnihan()
	name.Original.MiddleName.assignUnihan()
	name.Original.GivenName.assignUnihan()
	name.Original.FamilyName.assignSpec()
	name.Original.MiddleName.assignSpec()
	name.Original.GivenName.assignSpec()
	name.Original.FullNameStr = fmt.Sprintf("%s %s", name.Original.FamilyName.Str, name.Original.GivenName.Str)

	// Simplified
	name.Simplified.FamilyName.Characters = name.Original.FamilyName.simplify()
	name.Simplified.FamilyName.assignSpec()
	name.Simplified.MiddleName.Characters = name.Original.MiddleName.simplify()
	name.Simplified.MiddleName.assignSpec()
	name.Simplified.GivenName.Characters = name.Original.GivenName.simplify()
	name.Simplified.GivenName.assignSpec()
	name.Simplified.FullNameStr = fmt.Sprintf("%s %s", name.Simplified.FamilyName.Str, name.Simplified.GivenName.Str)

	// Traditional
	name.Traditional.FamilyName.Characters = name.Original.FamilyName.traditionalized()
	name.Traditional.FamilyName.assignSpec()
	name.Traditional.MiddleName.Characters = name.Original.MiddleName.traditionalized()
	name.Traditional.MiddleName.assignSpec()
	name.Traditional.GivenName.Characters = name.Original.GivenName.traditionalized()
	name.Traditional.GivenName.assignSpec()
	name.Traditional.FullNameStr = fmt.Sprintf("%s %s", name.Traditional.FamilyName.Str, name.Traditional.GivenName.Str)

	_stripTone := func(pinyin string) string {
		var (
			re  *regexp.Regexp
			ret string
		)

		re = regexp.MustCompile(`ā|á|ǎ|à`)
		ret = re.ReplaceAllString(pinyin, "a")
		re = regexp.MustCompile(`ō|ó|ǒ|ò`)
		ret = re.ReplaceAllString(ret, "o")
		re = regexp.MustCompile(`ê|ē|é|ě|è`)
		ret = re.ReplaceAllString(ret, "e")
		re = regexp.MustCompile(`ī|í|ǐ|ì`)
		ret = re.ReplaceAllString(ret, "i")
		re = regexp.MustCompile(`ū|ú|ǔ|ù`)
		ret = re.ReplaceAllString(ret, "u")
		re = regexp.MustCompile(`ǖ|ǘ|ǚ|ǜ|ü`)
		ret = re.ReplaceAllString(ret, "yu")

		return ret
	}

	_getPinyin := func(c *unihan.HanCharacter) (string, string) {
		var (
			pinyin     = "_"
			pinyinTone = "_"
			parts      []string
		)

		pinyinTone = list.QueryPinyinSpecial(c.Unicode)
		if pinyinTone != "" && pinyinTone != "_" {
			pinyin = _stripTone(pinyinTone)
		} else {
			if c != nil && c.Readings != nil {
				if c.Readings["kMandarin"] != nil {
					pinyin = _stripTone(c.Readings["kMandarin"].Reading)
					pinyinTone = c.Readings["kMandarin"].Reading
				} else {
					if c.Readings["kXHC1983"] != nil {
						// XianDaiHanYuCiDian
						parts = strings.Split(c.Readings["kXHC1983"].Reading, " ")
						parts = strings.Split(parts[0], ":")
					} else if c.Readings["kHanyuPinyin"] != nil {
						parts = strings.Split(c.Readings["kHanyuPinyin"].Reading, ":")
					}

					if parts != nil && len(parts) == 2 {
						parts = strings.Split(parts[1], ",")
						pinyin = _stripTone(parts[0])
						pinyinTone = parts[0]
					}
				}
			}
		}

		return pinyin, pinyinTone
	}

	for _, v := range name.Simplified.FamilyName.Characters {
		p, pt := _getPinyin(v)
		name.Pinyin = append(name.Pinyin, p)
		name.PinyinTone = append(name.PinyinTone, pt)
	}
	for _, v := range name.Simplified.MiddleName.Characters {
		p, pt := _getPinyin(v)
		name.Pinyin = append(name.Pinyin, p)
		name.PinyinTone = append(name.PinyinTone, pt)
	}
	for _, v := range name.Simplified.GivenName.Characters {
		p, pt := _getPinyin(v)
		name.Pinyin = append(name.Pinyin, p)
		name.PinyinTone = append(name.PinyinTone, pt)
	}

	for i, v := range name.Original.FamilyName.FiveElements {
		if v == utils.ElementUnknown {
			name.Original.FamilyName.FiveElements[i] = name.Simplified.FamilyName.FiveElements[i]
		}
	}
	for i, v := range name.Original.MiddleName.FiveElements {
		if v == utils.ElementUnknown {
			name.Original.MiddleName.FiveElements[i] = name.Simplified.MiddleName.FiveElements[i]
		}
	}
	for i, v := range name.Original.GivenName.FiveElements {
		if v == utils.ElementUnknown {
			name.Original.GivenName.FiveElements[i] = name.Simplified.GivenName.FiveElements[i]
		}
	}

	for i, v := range name.Traditional.FamilyName.FiveElements {
		if v == utils.ElementUnknown {
			name.Traditional.FamilyName.FiveElements[i] = name.Simplified.FamilyName.FiveElements[i]
		}
	}
	for i, v := range name.Traditional.MiddleName.FiveElements {
		if v == utils.ElementUnknown {
			name.Traditional.MiddleName.FiveElements[i] = name.Simplified.MiddleName.FiveElements[i]
		}
	}
	for i, v := range name.Traditional.GivenName.FiveElements {
		if v == utils.ElementUnknown {
			name.Traditional.GivenName.FiveElements[i] = name.Simplified.GivenName.FiveElements[i]
		}
	}
}

// RemoveUnihan : Remove unihan defination from name
func (name *Name) RemoveUnihan() {
	// Original
	name.Original.FamilyName.Characters = nil
	name.Original.MiddleName.Characters = nil
	name.Original.GivenName.Characters = nil

	// Simplified
	name.Simplified.FamilyName.Characters = nil
	name.Simplified.MiddleName.Characters = nil
	name.Simplified.GivenName.Characters = nil

	// Traditional
	name.Traditional.FamilyName.Characters = nil
	name.Traditional.MiddleName.Characters = nil
	name.Traditional.GivenName.Characters = nil

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
