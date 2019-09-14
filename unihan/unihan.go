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
 * @file unihan.go
 * @package unihan
 * @author Dr.NP <np@corp.herewetech.com>
 * @since 05/27/2019
 */

package unihan

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"yixuan_naming/utils"
)

// Indices

type hanDictionaryIndex struct {
	IndexType        string `json:"index_type"`
	Index            string `json:"index"`
	IndexDescription string `json:"index_description,omitempty"`
}

// Like-datas
type hanDictionaryLikeData struct {
	DataType        string `json:"data_type"`
	Data            string `json:"data"`
	DataDescription string `json:"data_description,omitempty"`
}

// IRG
type hanIRGSource struct {
	SourceType        string `json:"source_type"`
	Source            string `json:"source"`
	SourceDescription string `json:"source_description,omitempty"`
}

// Numeric
type hanNumericValue struct {
	ValueType        string `json:"value_type"`
	Value            int    `json:"value"`
	ValueDescription string `json:"value_description"`
}

// Other-Mappings
type hanMapping struct {
	MappingType         string `json:"mapping_type"`
	Mapping             string `json:"mapping"`
	MappingDescrription string `json:"mapping_description,omitempty"`
}

// RadicalStrokeCount
type hanRSCount struct {
	RSType                       string `json:"radical_stroke_type"`
	RadicalStroke                string `json:"radical_stroke"`
	Radical                      int    `json:"radical"`
	RadicalStr                   string `json:"radical_str"`
	RadicalStrokeCount           int    `json:"radical_stroke_count"`
	RadicalFrequency             int    `json:"radical_frequency"`
	RadicalAdditionalStrokeCount int    `json:"radical_additional_stroke_count"`
	RadicalStrokeDescription     string `json:"radical_stroke_description,omitempty"`
}

// Reading
type hanReading struct {
	ReadingType        string `json:"reading_type"`
	Reading            string `json:"reading"`
	ReadingDescription string `json:"reading_description,omitempty"`
}

// Variant
type hanVariant struct {
	VariantType        string `json:"variant_type"`
	Variant            string `json:"variant"`
	VariantDescription string `json:"variant_description,omitempty"`
}

// HanCharacter : Structure of unihan item
type HanCharacter struct {
	Unicode             rune                              `json:"unicode"`
	Utf8Str             string                            `json:"utf8_str"`
	DictionaryIndecies  map[string]*hanDictionaryIndex    `json:"dictionary_indices,omitempty"`
	DictionaryLikeDatas map[string]*hanDictionaryLikeData `json:"dictionary_like_datas,omitempty"`
	IRGSources          map[string]*hanIRGSource          `json:"irg_sourcess,omitempty"`
	NumericValues       map[string]*hanNumericValue       `json:"numeric_values,omitempty"`
	OtherMappings       map[string]*hanMapping            `json:"other_mappings,omitempty"`
	RadicalStrokeCounts map[string]*hanRSCount            `json:"radical_stroke_counts,omitempty"`
	Readings            map[string]*hanReading            `json:"readings,omitempty"`
	Variants            map[string]*hanVariant            `json:"variants,omitempty"`
}

// Map for unihan
var unihanM map[rune]*HanCharacter

func (c *HanCharacter) appendProperty(property, value string) {
	switch property {
	case "kCheungBauerIndex", "kCowles", "kDaeJaweon", "kFennIndex", "kGSR", "kHanYu", "kIRGDaeJaweon", "kIRGDaiKanwaZiten", "kIRGHanyuDaZidian", "kIRGKangXi", "kKangXi", "kKarlgren", "kLau", "kMatthews", "kMeyerWempe", "kMorohashi", "kNelson", "kSBGY":
		// DictionaryIndices
		if c.DictionaryIndecies == nil {
			c.DictionaryIndecies = make(map[string]*hanDictionaryIndex)
		}

		if c.DictionaryIndecies[property] == nil {
			c.DictionaryIndecies[property] = &hanDictionaryIndex{
				IndexType: property,
				Index:     value,
			}
		}
	case "kCangjie", "kCheungBauer", "kCihaiT", "kFenn", "kFourCornerCode", "kFrequency", "kGradeLevel", "kHDZRadBreak", "kHKGlyph", "kPhonetic", "kTotalStrokes":
		// DictionaryLikeDatas
		if c.DictionaryLikeDatas == nil {
			c.DictionaryLikeDatas = make(map[string]*hanDictionaryLikeData)
		}

		if c.DictionaryLikeDatas[property] == nil {
			c.DictionaryLikeDatas[property] = &hanDictionaryLikeData{
				DataType: property,
				Data:     value,
			}
		}
	case "kCompatibilityVariant", "kIICore", "kIRG_GSource", "kIRG_HSource", "kIRG_JSource", "kIRG_KPSource", "kIRG_KSource", "kIRG_TSource", "kIRG_USource", "kIRG_VSource", "kIRG_MSource":
		// IRG
		if c.IRGSources == nil {
			c.IRGSources = make(map[string]*hanIRGSource)
		}

		if c.IRGSources[property] == nil {
			c.IRGSources[property] = &hanIRGSource{
				SourceType: property,
				Source:     value,
			}
		}
	case "kAccountingNumeric", "kOtherNumeric", "kPrimaryNumeric":
		if c.NumericValues == nil {
			c.NumericValues = make(map[string]*hanNumericValue)
		}

		// Only single
		if c.NumericValues[property] == nil {
			nv, _ := strconv.ParseInt(value, 10, 64)
			c.NumericValues[property] = &hanNumericValue{
				ValueType: property,
				Value:     int(nv),
			}
		}
	case "kBigFive", "kCCCII", "kCNS1986", "kCNS1992", "kEACC", "kGB0", "kGB1", "kGB3", "kGB5", "kGB7", "kGB8", "kHKSCS", "kIBMJapan", "kJa", "kJinmeiyoKanji", "kJis0", "kJis1", "kJIS0213", "kJoyoKanji", "kKoreanEducationHanja", "kKoreanName", "kKPS0", "kKPS1", "kKSC0", "kKSC1", "kMainlandTelegraph", "kPseudoGB1", "kTaiwanTelegraph", "kTGH", "kXerox":
		if c.OtherMappings == nil {
			c.OtherMappings = make(map[string]*hanMapping)
		}

		if c.OtherMappings[property] == nil {
			c.OtherMappings[property] = &hanMapping{
				MappingType: property,
				Mapping:     value,
			}
		}
	case "kRSAdobe_Japan1_6", "kRSJapanese", "kRSKangXi", "kRSKanWa", "kRSKorean", "kRSUnicode":
		if c.RadicalStrokeCounts == nil {
			c.RadicalStrokeCounts = make(map[string]*hanRSCount)
		}

		var (
			radical                 int
			radicalAdditionalStroke int
		)

		if property != "kRSAdobe_Japan1_6" {
			// Parse RS
			rs := regexp.MustCompile(`([0-9]+)\.([0-9]+)`).FindStringSubmatch(value)
			if rs != nil {
				radical, _ = strconv.Atoi(rs[1])
				radicalAdditionalStroke, _ = strconv.Atoi(rs[2])

				if radical > 214 || radical < 0 {
					radical = 0
				}
			}
		}

		radicalObj := utils.GetRadical(radical)
		if c.RadicalStrokeCounts[property] == nil {
			c.RadicalStrokeCounts[property] = &hanRSCount{
				RSType:                       property,
				RadicalStroke:                value,
				Radical:                      radical,
				RadicalStr:                   radicalObj.Str,
				RadicalStrokeCount:           radicalObj.Stroke,
				RadicalFrequency:             radicalObj.Frequency,
				RadicalAdditionalStrokeCount: radicalAdditionalStroke,
			}
		}
	case "kCantonese", "kDefinition", "kHangul", "kHanyuPinlu", "kHanyuPinyin", "kJapaneseKun", "kJapaneseOn", "kKorean", "kMandarin", "kTang", "kVietnamese", "kXHC1983":
		if c.Readings == nil {
			c.Readings = make(map[string]*hanReading)
		}

		if c.Readings[property] == nil {
			c.Readings[property] = &hanReading{
				ReadingType: property,
				Reading:     value,
			}
		}
	case "kSemanticVariant", "kSimplifiedVariant", "kSpecializedSemanticVariant", "kTraditionalVariant", "kZVariant":
		if c.Variants == nil {
			c.Variants = make(map[string]*hanVariant)
		}

		if c.Variants[property] == nil {
			c.Variants[property] = &hanVariant{
				VariantType: property,
				Variant:     value,
			}
		}

		/*
			if property == "kTraditionalVariant" {
				ts := regexp.MustCompile(`U\+(\w+)`).FindAllStringSubmatch(value, -1)
				if len(ts) > 1 {
					// Multi
					fmt.Print(c.Unicode, " ", c.Utf8Str, " => ")
					for _, t := range ts {
						r, _ := strconv.ParseInt(t[1], 16, 32)
						fmt.Print(r, " ", string(r), " ")
					}

					fmt.Println()
				}
			}
		*/
	}

	return
}

// LoadUnihanLibraries : Load unihan txt files (12.0+)
func LoadUnihanLibraries(dir string) (int, int, error) {
	fileList := []string{
		"Unihan_DictionaryIndices.txt",
		"Unihan_DictionaryLikeData.txt",
		"Unihan_IRGSources.txt",
		"Unihan_NumericValues.txt",
		"Unihan_OtherMappings.txt",
		"Unihan_RadicalStrokeCounts.txt",
		"Unihan_Readings.txt",
		"Unihan_Variants.txt",
	}

	var (
		filename string
		fullPath string
		f        *os.File
		err      error
		line     string
		reLine   = regexp.MustCompile(`^U\+([0-9A-Fa-f]+)\t([0-9A-Za-z_]+)\t(.+)`)
		rcode    int64
		r        rune
		parts    []string
		total    int
		scanner  *bufio.Scanner
	)

	unihanM = make(map[rune]*HanCharacter)
	for _, filename = range fileList {
		fullPath = fmt.Sprintf("%s/unihan/%s", dir, filename)
		f, err = os.Open(fullPath)
		if err != nil {
			// Open file failed
			unihanM = nil
			return 0, 0, fmt.Errorf("Read unihan library file <%s> failed", fullPath)
		}

		scanner = bufio.NewScanner(f)
		for scanner.Scan() == true {
			line = scanner.Text()
			if parts = reLine.FindStringSubmatch(line); parts != nil {
				// Valid line
				rcode, err = strconv.ParseInt(parts[1], 16, 32)
				if err == nil {
					r = rune(rcode)
					if unihanM[r] == nil {
						unihanM[r] = &HanCharacter{
							Unicode: r,
							Utf8Str: string(r),
						}
					}

					unihanM[r].appendProperty(parts[2], parts[3])
					total++
				}
			}
		}

		f.Close()
	}

	return total, len(unihanM), nil
}

// Query : query unihan character by given unicode
func Query(unicode rune) (*HanCharacter, error) {
	c := unihanM[unicode]

	return c, nil
}

// QueryStroke : Query stroke of character
func (c *HanCharacter) QueryStroke() (int, int, int, error) {
	var (
		stroke        int
		unicodeStroke int
		kangxiStroke  int
	)

	// Total stroke
	if c.DictionaryLikeDatas != nil &&
		c.DictionaryLikeDatas["kTotalStrokes"] != nil {
		stroke, _ = strconv.Atoi(c.DictionaryLikeDatas["kTotalStrokes"].Data)
	}

	if c.RadicalStrokeCounts != nil {
		// Unicode RS
		if c.RadicalStrokeCounts["kRSUnicode"] != nil {
			unicodeStroke = c.RadicalStrokeCounts["kRSUnicode"].RadicalStrokeCount + c.RadicalStrokeCounts["kRSUnicode"].RadicalAdditionalStrokeCount
		}

		// KangxiRS
		if c.RadicalStrokeCounts["kRSKangXi"] != nil {
			kangxiStroke = c.RadicalStrokeCounts["kRSKangXi"].RadicalStrokeCount + c.RadicalStrokeCounts["kRSKangXi"].RadicalAdditionalStrokeCount
		}
	}

	return kangxiStroke, unicodeStroke, stroke, nil
}

// QueryStrokePrefer : Query stroke singly
func (c *HanCharacter) QueryStrokePrefer() int {
	ks, us, s, _ := c.QueryStroke()
	if us > 0 {
		return us
	}

	if ks > 0 {
		return ks
	}

	return s
}

// QueryTraditional : Query traditional character
func (c *HanCharacter) QueryTraditional() ([]rune, error) {
	var (
		t   []string
		ts  [][]string
		tsi []rune
	)

	if c.Variants != nil &&
		c.Variants["kTraditionalVariant"] != nil {
		v := c.Variants["kTraditionalVariant"].Variant
		ts = regexp.MustCompile(`U\+(\w+)`).FindAllStringSubmatch(v, -1)
		for _, t = range ts {
			r, _ := strconv.ParseInt(t[1], 16, 32)
			//if rune(r) != c.Unicode {
			tsi = append(tsi, rune(r))
			//}
		}
	}

	return tsi, nil
}

// QueryTraditionalPrefer : Query preferer traditional character
func (c *HanCharacter) QueryTraditionalPrefer() (rune, error) {
	tsi, err := c.QueryTraditional()
	if err != nil || tsi == nil {
		return 0, err
	}

	return tsi[0], nil
}

// QueryTraditionalLazy : Lazy query traditional character, assert KangXi dictionary
func (c *HanCharacter) QueryTraditionalLazy() (rune, error) {
	var (
		p        rune
		pc       *HanCharacter
		ci, pi   string
		ckx, pkx bool
		err      error
	)
	p, err = c.QueryTraditionalPrefer()
	if p == 0 || err != nil {
		// No traditional
		return c.Unicode, nil
	}

	pc, _ = Query(p)
	if c.DictionaryIndecies != nil &&
		c.DictionaryIndecies["kIRGKangXi"] != nil {
		ci = c.DictionaryIndecies["kIRGKangXi"].Index
		if '0' == ci[len(ci)-1] {
			ckx = true
		} else {
			ckx = false
		}
	}

	if pc != nil && pc.DictionaryIndecies != nil &&
		pc.DictionaryIndecies["kIRGKangXi"] != nil {
		pi = pc.DictionaryIndecies["kIRGKangXi"].Index
		if '0' == pi[len(pi)-1] {
			pkx = true
		} else {
			pkx = false
		}
	}

	if ci != "" && ci != pi && ckx == true {
		return c.Unicode, nil
	} else if pkx == true {
		return pc.Unicode, nil
	}

	return c.Unicode, nil
}

// QuerySimplified : Query traditional character
func (c *HanCharacter) QuerySimplified() ([]rune, error) {
	var (
		t   []string
		ts  [][]string
		tsi []rune
	)

	if c.Variants != nil &&
		c.Variants["kSimplifiedVariant"] != nil {
		v := c.Variants["kSimplifiedVariant"].Variant
		ts = regexp.MustCompile(`U\+(\w+)`).FindAllStringSubmatch(v, -1)
		for _, t = range ts {
			r, _ := strconv.ParseInt(t[1], 16, 32)
			//if rune(r) != c.Unicode {
			tsi = append(tsi, rune(r))
			//}
		}
	}

	return tsi, nil
}

// QuerySimplifiedPrefer : Query prefered simplified character
func (c *HanCharacter) QuerySimplifiedPrefer() (rune, error) {
	tsi, err := c.QuerySimplified()
	if err != nil || tsi == nil {
		return 0, err
	}

	return tsi[0], nil
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
