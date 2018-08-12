package main

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

func createAdjectiveAnswer(wordInfo *WordData, form string, showKanji bool) {

	wordInfo.TestAnswer = ""

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
