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
 * @file descriptions.go
 * @package unihan
 * @author Dr.NP <np@corp.herewetech.com>
 * @since 05/29/2019
 */

package unihan

var propertyDescriptions = map[string]string{
	"kAccountingNumeric":          "会计学数字",
	"kBigFive":                    "Big5编码",
	"kCangjie":                    "仓颉输入码",
	"kCantonese":                  "粤语读音",
	"kCCCII":                      "中文资讯交换码（十六进制）",
	"kCheungBauer":                "以汉字写粤语",
	"kCheungBauerIndex":           "以汉字写粤语索引",
	"kCihaiT":                     "辞海",
	"kCNS1986":                    "CNS 1986（十六进制）",
	"kCNS1992":                    "CNS 1992（十六进制）",
	"kCompatibilityVariant":       "",
	"kCowles":                     "Roy T Cowles字典",
	"kDaeJaweon":                  "Dae Jaweon韩文字典",
	"kDefinition":                 "英文释义",
	"kEACC":                       "",
	"kFenn":                       "",
	"kFennIndex":                  "",
	"kFourCornerCode":             "四角号码",
	"kFrequency":                  "使用频率等级（繁体）",
	"kGB0":                        "GB 2312-80",
	"kGB1":                        "GB 12345-90",
	"kGB3":                        "GB 7589-87",
	"kGB5":                        "GB 7590-87",
	"kGB7":                        "现代汉语通用汉字列表",
	"kGB8":                        "GB/T 8565.2-1988",
	"kGradeLevel":                 "朗文初级中文词典定义香港学校系统认知年级",
	"kGSR":                        "Bernhard Karlgren’s Grammata Serica Recensa (1957)",
	"kHangul":                     "Hangul现代韩语读音",
	"kHanYu":                      "汉语大字典",
	"kHanyuPinlu":                 "汉语频率",
	"kHanyuPinyin":                "汉语拼音",
	"kHDZRadBreak":                "汉语大字典部首检字",
	"kHKGlyph":                    "香港常用字字形表",
	"kHKSCS":                      "香港增补字Big5映射",
	"kIBMJapan":                   "IBM日文映射",
	"kIICore":                     "IICore表义符号",
	"kIRG_GSource":                "国标中文",
	"kIRG_HSource":                "香港中文",
	"kIRG_JSource":                "日文",
	"kIRG_KPSource":               "朝鲜韩文",
	"kIRG_KSource":                "韩国韩文",
	"kIRG_MSource":                "澳门中文",
	"kIRG_TSource":                "台湾中文",
	"kIRG_USource":                "",
	"kIRG_VSource":                "越南文",
	"kIRGDaeJaweon":               "Dae Jaweon韩文字典索引",
	"kIRGDaiKanwaZiten":           "Dai Kanwa Ziten日文字典索引",
	"kIRGHanyuDaZidian":           "汉语大字典索引",
	"kIRGKangXi":                  "康熙字典索引",
	"kJa":                         "",
	"kJapaneseKun":                "日文读音",
	"kJapaneseOn":                 "中国日文读音",
	"kJinmeiyoKanji":              "人名用汉字（日本）",
	"kJis0":                       "JIS X 0208-1990",
	"kJis1":                       "JIS X 0212-1990",
	"kJIS0213":                    "JIS X 0213-2004",
	"kJoyoKanji":                  "常用汉字（日本）",
	"kKangXi":                     "康熙字典",
	"kKarlgren":                   "中文与中国日文分析字典",
	"kKorean":                     "韩文读音",
	"kKoreanEducationHanja":       "汉文教育用基础汉字（韩国）",
	"kKoreanName":                 "人名用汉字（韩国）",
	"kKPS0":                       "KPS 9566-97",
	"kKPS1":                       "KPS 10721-2000",
	"kKSC0":                       "KS X 1001:1992",
	"kKSC1":                       "KS X 1002:1991",
	"kLau":                        "实用粤英字典",
	"kMainlandTelegraph":          "电报明码",
	"kMandarin":                   "习惯拼音读音",
	"kMatthews":                   "马修斯汉英词典",
	"kMeyerWempe":                 "学生粤英词典",
	"kMorohashi":                  "Dai Kanwa Ziten修订版",
	"kNelson":                     "现代读者日英字典",
	"kOtherNumeric":               "其它数字",
	"kPhonetic":                   "万字注音索引",
	"kPrimaryNumeric":             "标准写法数字",
	"kPseudoGB1":                  "",
	"kRSAdobe_Japan1_6":           "Adobe-Japan1-6 CID部首笔画",
	"kRSJapanese":                 "日文部首笔画",
	"kRSKangXi":                   "康熙字典部首笔画",
	"kRSKanWa":                    "",
	"kRSKorean":                   "韩文部首笔画",
	"kRSUnicode":                  "Unicode部首笔画",
	"kSBGY":                       "宋本广韵索引",
	"kSemanticVariant":            "语意变体",
	"kSimplifiedVariant":          "简体变体",
	"kSpecializedSemanticVariant": "专业语意变体",
	"kTaiwanTelegraph":            "台湾电报码",
	"kTang":                       "唐代发音",
	"kTGH":                        "通用规范汉字表",
	"kTotalStrokes":               "总笔画数",
	"kTraditionalVariant":         "繁体变体",
	"kVietnamese":                 "越南语读音",
	"kXerox":                      "施乐编码",
	"kXHC1983":                    "现代汉语词典读音",
	"kZVariant":                   "Z变体",
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
