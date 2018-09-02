package main

import "math/rand"

func createPresentNegativeForm(wordInfo *WordData, inputString string) {

	if wordInfo.Subtype == iAdjective {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "くない")
	} else {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 0, "じゃない")
	}
}

func createPastForm(wordInfo *WordData, inputString string) {

	if wordInfo.Subtype == iAdjective {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "かった")
	} else {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 0, "でした")
	}
}

func createPastNegativeForm(wordInfo *WordData, inputString string) {

	if wordInfo.Subtype == iAdjective {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "くなかった")
	} else {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 0, "じゃなかった")
	}
}

func createTeForm(wordInfo *WordData, inputString string) {

	if wordInfo.Subtype == iAdjective {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "くて")
	} else {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 0, "で")
	}
}

func createAdjectiveAnswer(wordInfo *WordData, form string, showKanji bool, language string) {

	wordInfo.TestAnswer = ""
	wordInfo.AnswerLanguage = language

	if language == langMix {
		langRand := rand.Intn(2)
		if langRand == 0 {
			wordInfo.AnswerLanguage = langJapanese
		} else {
			wordInfo.AnswerLanguage = langEnglish
		}
	}

	if wordInfo.AnswerLanguage == langEnglish {
		wordInfo.TestAnswer = wordInfo.English
		return
	}

	var inputString string
	if showKanji {
		inputString = wordInfo.Kanji
	} else {
		inputString = wordInfo.Japanese
	}

	switch form {

	case present:
		wordInfo.TestAnswer = inputString
		break

	case presentNegative:
		createPresentNegativeForm(wordInfo, inputString)
		break

	case past:
		createPastForm(wordInfo, inputString)
		break

	case pastNegative:
		createPastNegativeForm(wordInfo, inputString)
		break

	case te:
		createTeForm(wordInfo, inputString)
		break

	default:
		break
	}
}
