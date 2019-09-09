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
 * @file radicals.go
 * @package utils
 * @author Dr.NP <np@corp.herewetech.com>
 * @since 05/29/2019
 */

package utils

// ChineseRadical : Radical from KangXi indexes
type ChineseRadical struct {
	ID        int    `json:"id"`
	Unicode   rune   `json:"unicode"`
	Str       string `json:"str"`
	Stroke    int    `json:"stroke"`
	Meaning   string `json:"meaning"`
	Frequency int    `json:"frequency"`
}

var radicalList = []ChineseRadical{
	{0, 0, "", 0, "", 0},
	{1, 21033, "一", 1, "one", 42},
	{2, 21034, "丨", 1, "line", 21},
	{3, 21035, "丶", 1, "dot", 10},
	{4, 21036, "丿(乚、乛)", 1, "slash", 33},
	{5, 21037, "乙(乀)", 1, "second", 42},
	{6, 21038, "亅", 1, "hook", 19},
	{7, 21039, "二", 2, "two", 29},
	{8, 21040, "亠", 2, "lid", 38},
	{9, 21041, "人(亻)", 2, "man", 794},
	{10, 21042, "儿", 2, "son, legs", 52},
	{11, 21043, "入", 2, "enter", 28},
	{12, 21044, "八(丷)", 2, "eight", 44},
	{13, 21045, "冂", 2, "wide", 50},
	{14, 21046, "冖", 2, "cloth cover", 30},
	{15, 21047, "冫", 2, "ice", 115},
	{16, 21048, "几", 2, "table", 38},
	{17, 21049, "凵", 2, "receptacle", 23},
	{18, 21050, "刀(刂、⺈)", 2, "knife", 377},
	{19, 21051, "力", 2, "power", 163},
	{20, 21052, "勹", 2, "wrap", 64},
	{21, 21053, "匕", 2, "spoon", 19},
	{22, 21054, "匚", 2, "box", 64},
	{23, 21055, "匸", 2, "hiding enclosure", 17},
	{24, 21056, "十", 2, "ten", 55},
	{25, 21057, "卜", 2, "divination", 45},
	{26, 21058, "卩(㔾)", 2, "seal (device)", 40},
	{27, 21059, "厂", 2, "cliff", 129},
	{28, 21060, "厶", 2, "private", 40},
	{29, 21061, "又", 2, "again", 91},
	{30, 21062, "口", 3, "mouth", 1146},
	{31, 21063, "囗", 3, "enclosure", 118},
	{32, 21064, "土", 3, "earth", 580},
	{33, 21065, "士", 3, "scholar", 24},
	{34, 21066, "夂", 3, "go", 11},
	{35, 21067, "夊", 3, "go slowly", 23},
	{36, 21068, "夕", 3, "evening", 34},
	{37, 21069, "大", 3, "big", 132},
	{38, 21070, "女", 3, "woman", 681},
	{39, 21071, "子", 3, "child", 83},
	{40, 21072, "宀", 3, "roof", 246},
	{41, 21073, "寸", 3, "inch", 40},
	{42, 21074, "小(⺌、⺍)", 3, "small", 41},
	{43, 21075, "尢(尣)", 3, "lame", 66},
	{44, 21076, "尸", 3, "corpse", 148},
	{45, 21077, "屮", 3, "sprout", 38},
	{46, 21078, "山", 3, "mountain", 636},
	{47, 21079, "巛(川)", 3, "river", 26},
	{48, 21080, "工", 3, "work", 17},
	{49, 21081, "己", 3, "oneself", 20},
	{50, 21082, "巾", 3, "turban", 295},
	{51, 21083, "干", 3, "dry", 9},
	{52, 21084, "幺(么)", 3, "short thread", 50},
	{53, 21085, "广", 3, "dotted cliff", 15},
	{54, 21086, "廴", 3, "long stride", 9},
	{55, 21087, "廾", 3, "arch", 50},
	{56, 21088, "弋", 3, "shoot", 15},
	{57, 21089, "弓", 3, "bow", 165},
	{58, 21090, "彐(彑)", 3, "snout", 25},
	{59, 21091, "彡", 3, "bristle", 62},
	{60, 21092, "彳", 3, "step", 215},
	{61, 21093, "心(忄、⺗)", 4, "heart", 1115},
	{62, 21094, "戈", 4, "halberd", 116},
	{63, 21095, "戶(户、戸)", 4, "door", 44},
	{64, 21096, "手(扌、龵)", 4, "hand", 1203},
	{65, 21097, "支", 4, "branch", 26},
	{66, 21098, "攴(攵)", 4, "rap, tap", 296},
	{67, 21099, "文", 4, "script", 26},
	{68, 21100, "斗", 4, "dipper", 32},
	{69, 21101, "斤", 4, "axe", 55},
	{70, 21102, "方", 4, "square", 92},
	{71, 21103, "无(旡)", 4, "not", 12},
	{72, 21104, "日", 4, "sun", 453},
	{73, 21105, "曰", 4, "say", 37},
	{74, 21106, "月", 4, "moon", 69},
	{75, 21107, "木", 4, "tree", 1369},
	{76, 21108, "欠", 4, "lack", 235},
	{77, 21109, "止", 4, "stop", 99},
	{78, 21110, "歹(歺)", 4, "death", 231},
	{79, 21111, "殳", 4, "weapon", 93},
	{80, 21112, "毋(母)", 4, "do not", 16},
	{81, 21113, "比", 4, "compare", 21},
	{82, 21114, "毛", 4, "fur", 211},
	{83, 21115, "氏", 4, "clan", 10},
	{84, 21116, "气", 4, "steam", 17},
	{85, 21117, "水(氵、氺)", 4, "water", 1595},
	{86, 21118, "火(灬)", 4, "fire", 639},
	{87, 21119, "爪(爫)", 4, "claw", 36},
	{88, 21120, "父", 4, "father", 10},
	{89, 21121, "爻", 4, "Trigrams", 16},
	{90, 21122, "爿(丬)", 4, "split wood", 48},
	{91, 21123, "片", 4, "slice", 77},
	{92, 21124, "牙", 4, "fang", 9},
	{93, 21125, "牛(牜、⺧)", 4, "cow", 233},
	{94, 21126, "犬(犭)", 4, "dog", 444},
	{95, 21127, "玄", 5, "profound", 6},
	{96, 21128, "玉(王、玊)", 5, "jade", 473},
	{97, 21129, "瓜", 5, "melon", 55},
	{98, 21130, "瓦", 5, "tile", 174},
	{99, 21131, "甘", 5, "sweet", 22},
	{100, 21132, "生", 5, "life", 22},
	{101, 21133, "用", 5, "use", 10},
	{102, 21134, "田", 5, "field", 192},
	{103, 21135, "疋(⺪)", 5, "bolt of cloth", 15},
	{104, 21136, "疒", 5, "sickness", 526},
	{105, 21137, "癶", 5, "footsteps", 15},
	{106, 21138, "白", 5, "white", 109},
	{107, 21139, "皮", 5, "skin", 94},
	{108, 21140, "皿", 5, "dish", 129},
	{109, 21141, "目(⺫)", 5, "eye", 647},
	{110, 21142, "矛", 5, "spear", 65},
	{111, 21143, "矢", 5, "arrow", 64},
	{112, 21144, "石", 5, "stone", 499},
	{113, 21145, "示(礻)", 5, "spirit", 213},
	{114, 21146, "禸", 5, "track", 12},
	{115, 21147, "禾", 5, "grain", 431},
	{116, 21148, "穴", 5, "cave", 298},
	{117, 21149, "立", 5, "stand", 101},
	{118, 21150, "竹(⺮)", 6, "bamboo", 953},
	{119, 21151, "米", 6, "rice", 318},
	{120, 21152, "糸(糹)", 6, "silk", 823},
	{121, 21153, "缶", 6, "jar", 77},
	{122, 21154, "网(⺲、罓、⺳)", 6, "net", 163},
	{123, 21155, "羊(⺶、⺷)", 6, "sheep", 156},
	{124, 21156, "羽", 6, "feather", 220},
	{125, 21157, "老(耂)", 6, "old", 22},
	{126, 21158, "而", 6, "and", 22},
	{127, 21159, "耒", 6, "plow", 84},
	{128, 21160, "耳", 6, "ear", 172},
	{129, 21161, "聿(⺺、⺻)", 6, "brush", 19},
	{130, 21162, "肉(⺼)", 6, "meat", 674},
	{131, 21163, "臣", 6, "minister", 16},
	{132, 21164, "自", 6, "self", 34},
	{133, 21165, "至", 6, "arrive", 24},
	{134, 21166, "臼", 6, "mortar", 71},
	{135, 21167, "舌", 6, "tongue", 31},
	{136, 21168, "舛", 6, "oppose", 10},
	{137, 21169, "舟", 6, "boat", 197},
	{138, 21170, "艮", 6, "stopping", 5},
	{139, 21171, "色", 6, "color", 21},
	{140, 21172, "艸(⺿)", 6, "grass", 1902},
	{141, 21173, "虍", 6, "tiger", 114},
	{142, 21174, "虫", 6, "insect", 1067},
	{143, 21175, "血", 6, "blood", 60},
	{144, 21176, "行", 6, "walk enclosure", 53},
	{145, 21177, "衣(⻂)", 6, "clothes", 607},
	{146, 21178, "襾(西、覀)", 6, "cover", 29},
	{147, 21179, "見", 7, "see", 161},
	{148, 21180, "角(⻇)", 7, "horn", 158},
	{149, 21181, "言(訁)", 7, "speech", 861},
	{150, 21182, "谷", 7, "valley", 54},
	{151, 21183, "豆", 7, "bean", 68},
	{152, 21184, "豕", 7, "pig", 148},
	{153, 21185, "豸", 7, "badger", 140},
	{154, 21186, "貝", 7, "shell", 277},
	{155, 21187, "赤", 7, "red", 31},
	{156, 21188, "走", 7, "run", 285},
	{157, 21189, "足(⻊)", 7, "foot", 580},
	{158, 21190, "身", 7, "body", 97},
	{159, 21191, "車", 7, "cart", 361},
	{160, 21192, "辛", 7, "bitter", 36},
	{161, 21193, "辰", 7, "morning", 15},
	{162, 21194, "辵(⻌、⻍、⻎}})", 7, "walk", 381},
	{163, 21195, "邑(⻏)", 7, "city", 350},
	{164, 21196, "酉", 7, "wine", 290},
	{165, 21197, "釆", 7, "distinguish", 14},
	{166, 21198, "里", 7, "village", 14},
	{167, 21199, "金(釒)", 8, "gold", 806},
	{168, 21200, "長(镸)", 8, "long", 55},
	{169, 21201, "門", 8, "gate", 246},
	{170, 21202, "阜(⻖)", 8, "mound", 348},
	{171, 21203, "隶", 8, "slave", 12},
	{172, 21204, "隹", 8, "short-tailed bird", 233},
	{173, 21205, "雨", 8, "rain", 298},
	{174, 21206, "青(靑)", 8, "blue", 17},
	{175, 21207, "非", 8, "wrong", 25},
	{176, 21208, "面(靣)", 9, "face", 66},
	{177, 21209, "革", 9, "leather", 305},
	{178, 21210, "韋", 9, "tanned leather", 100},
	{179, 21211, "韭", 9, "leek", 20},
	{180, 21212, "音", 9, "sound", 43},
	{181, 21213, "頁", 9, "leaf", 372},
	{182, 21214, "風", 9, "wind", 182},
	{183, 21215, "飛", 9, "fly", 92},
	{184, 21216, "食(飠)", 9, "eat", 403},
	{185, 21217, "首", 9, "head", 20},
	{186, 21218, "香", 9, "fragrant", 37},
	{187, 21219, "馬", 10, "horse", 472},
	{188, 21220, "骨", 10, "bone", 185},
	{189, 21221, "高(髙)", 10, "tall", 34},
	{190, 21222, "髟", 10, "hair", 243},
	{191, 21223, "鬥", 10, "fight", 23},
	{192, 21224, "鬯", 10, "sacrificial wine", 8},
	{193, 21225, "鬲", 10, "cauldron", 73},
	{194, 21226, "鬼", 10, "ghost", 141},
	{195, 21227, "魚", 11, "fish", 571},
	{196, 21228, "鳥", 11, "bird", 750},
	{197, 21229, "鹵", 11, "salt", 44},
	{198, 21230, "鹿", 11, "deer", 104},
	{199, 21231, "麥", 11, "wheat", 131},
	{200, 21232, "麻", 11, "hemp", 34},
	{201, 21233, "黃", 12, "yellow", 42},
	{202, 21234, "黍", 12, "millet", 46},
	{203, 21235, "黑", 12, "black", 172},
	{204, 21236, "黹", 12, "embroidery", 8},
	{205, 21237, "黽", 13, "frog", 40},
	{206, 21238, "鼎", 13, "tripod", 14},
	{207, 21239, "鼓", 13, "drum", 46},
	{208, 21240, "鼠", 13, "rat", 92},
	{209, 21241, "鼻", 14, "nose", 49},
	{210, 21242, "齊(斉)", 14, "even", 18},
	{211, 21243, "齒", 15, "tooth", 162},
	{212, 21244, "龍", 16, "dragon", 14},
	{213, 21245, "龜", 16, "turtle", 24},
	{214, 21246, "龠", 17, "flute", 19},
}

// GetRadical : Get radical from list
func GetRadical(index int) *ChineseRadical {
	if index < 0 || index > 214 {
		return nil
	}

	return &radicalList[index]
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
