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
 * @file messages.go
 * @package texts
 * @author Dr.NP <np@corp.herewetech.com>
 * @since 07/16/2019
 */

package texts

import (
	"bufio"
	"fmt"
	"os"
)

const (
	// MessageFiveRuleDescription : 0
	MessageFiveRuleDescription = iota
	// MessageFiveRuleSummary : 1
	MessageFiveRuleSummary
	// MessageGanDescription : 2
	MessageGanDescription
	// MessageSoundFiveElementDescription : 3
	MessageSoundFiveElementDescription
	// MessageTenGodDescription : 4
	MessageTenGodDescription
	// MessageTenGodSoul : 5
	MessageTenGodSoul
	// MessageThreeElementDescription : 6
	MessageThreeElementDescription
	// MessageAnimalYear : 7
	MessageAnimalYear
	// MessageAnimalRadicalsDescription : 8
	MessageAnimalRadicalsDescription
)

var (
	fiveRuleDescriptionMessages         [][]string
	fiveRuleSummaryMessages             [][]string
	ganDescriptionMessages              [][]string
	soundFiveElementDescriptionMessages [][]string
	tenGodDescriptionMessages           [][]string
	tenGodSoulMessages                  [][]string
	threeElementDescriptionMessages     [][]string
	animalYearMessages                  [][]string
	animalRadicalsDescriptions          [][]string
)

// GetMessage : Get message from list
func GetMessage(message int, index int, language int) string {
	var messages [][]string

	switch message {
	case MessageFiveRuleDescription:
		messages = fiveRuleDescriptionMessages
	case MessageFiveRuleSummary:
		messages = fiveRuleSummaryMessages
	case MessageGanDescription:
		messages = ganDescriptionMessages
	case MessageSoundFiveElementDescription:
		messages = soundFiveElementDescriptionMessages
	case MessageTenGodDescription:
		messages = tenGodDescriptionMessages
	case MessageTenGodSoul:
		messages = tenGodSoulMessages
	case MessageThreeElementDescription:
		messages = threeElementDescriptionMessages
	case MessageAnimalYear:
		messages = animalYearMessages
	case MessageAnimalRadicalsDescription:
		messages = animalRadicalsDescriptions
	}

	if messages == nil || len(messages) < 1 {
		return ""
	}

	if language < 0 || language >= len(messages) {
		language = LanguageDefault
	}

	if index < 0 || index >= len(messages[language]) {
		return ""
	}

	return messages[language][index]
}

// LoadMessages : Load message files
func LoadMessages(dir string) (int, error) {
	var (
		total  int
		cLines int
	)

	cLines, fiveRuleDescriptionMessages = loadMessage(dir, "FiveRuleDescriptions")
	total += cLines
	cLines, fiveRuleSummaryMessages = loadMessage(dir, "FiveRuleSummaries")
	total += cLines
	cLines, ganDescriptionMessages = loadMessage(dir, "GanDescriptions")
	total += cLines
	cLines, soundFiveElementDescriptionMessages = loadMessage(dir, "SoundFiveElementDescriptions")
	total += cLines
	cLines, tenGodDescriptionMessages = loadMessage(dir, "TenGodDescriptions")
	total += cLines
	cLines, tenGodSoulMessages = loadMessage(dir, "TenGodSouls")
	total += cLines
	cLines, threeElementDescriptionMessages = loadMessage(dir, "ThreeElementDescriptions")
	total += cLines
	cLines, animalYearMessages = loadMessage(dir, "AnimalYears")
	total += cLines
	cLines, animalRadicalsDescriptions = loadMessage(dir, "AnimalRadicalsDescriptions")
	total += cLines

	return total, nil
}

func loadMessage(dir, name string) (int, [][]string) {
	var (
		fullPath string
		lang     string
		lines    []string
		ret      [][]string
		total    int
	)

	for _, lang = range []string{"S", "T", "E"} {
		fullPath = fmt.Sprintf("%s/message/%s%s.txt", dir, name, lang)
		lines = loadLines(fullPath)
		if lines != nil {
			total += len(lines)
		}
		ret = append(ret, lines)
	}

	return total, ret
}

func loadLines(file string) []string {
	var (
		f       *os.File
		err     error
		scanner *bufio.Scanner
		ret     []string
	)

	f, err = os.Open(file)
	if err != nil {
		return nil
	}

	scanner = bufio.NewScanner(f)
	for scanner.Scan() == true {
		ret = append(ret, scanner.Text())
	}

	f.Close()

	return ret
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
