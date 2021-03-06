package main

import (
	"fmt"
	"math/rand"
	"unicode/utf8"
)

func isVerbIrregular(base string, subType string) bool {

	if base == "ある" || base == "いる" || base == "くる" || base == "いく" || subType == "する" {
		return true
	}

	return false
}

func addVerbSuffix(inputString string, outputString *string, numRunesToRemove int, suffix string) {

	if numRunesToRemove == 0 {

		*outputString = inputString
		*outputString += suffix
		return
	}

	numRunesInWord := utf8.RuneCountInString(inputString)

	var numRunes int
	for _, runeValue := range inputString {
		if numRunes < numRunesInWord-numRunesToRemove {
			*outputString += string(runeValue)
			numRunes++
		} else {
			*outputString += suffix
			break
		}
	}
}

func createAruIruAnswer(wordInfo *WordData, inputString string, politeness string, formData verbFormData) {

	switch formData.form {

	case present:
		if politeness == polite {
			if wordInfo.Japanese == "ある" {
				wordInfo.TestAnswer = "あります"
			} else {
				wordInfo.TestAnswer = "います"
			}
		} else if politeness == casual {
			wordInfo.TestAnswer = inputString
		}
		break

	case presentNegative:
		if politeness == polite {
			if wordInfo.Japanese == "ある" {
				wordInfo.TestAnswer = "ありません"
			} else {
				wordInfo.TestAnswer = "いません"
			}
		} else if politeness == casual {
			if wordInfo.Japanese == "ある" {
				wordInfo.TestAnswer = "ない"
			} else {
				wordInfo.TestAnswer = "いない"
			}
		}
		break

	case past:
		if politeness == polite {
			if wordInfo.Japanese == "ある" {
				wordInfo.TestAnswer = "ありました"
			} else {
				wordInfo.TestAnswer = "いました"
			}
		} else if politeness == casual {
			if wordInfo.Japanese == "ある" {
				wordInfo.TestAnswer = "あった"
			} else {
				wordInfo.TestAnswer = "いた"
			}
		}
		break

	case pastNegative:
		if politeness == polite {
			if wordInfo.Japanese == "ある" {
				wordInfo.TestAnswer = "ありませんでした"
			} else {
				wordInfo.TestAnswer = "いませんでした"
			}
		} else if politeness == casual {
			if wordInfo.Japanese == "ある" {
				wordInfo.TestAnswer = "なかった"
			} else {
				wordInfo.TestAnswer = "いなかった"
			}
		}
		break

	case te:
		if wordInfo.Japanese == "ある" {
			wordInfo.TestAnswer = "あって"
		} else {
			wordInfo.TestAnswer = "いて"
		}
		break

	case potential:
		if wordInfo.Japanese == "ある" {
			wordInfo.TestAnswer = "あれる"
		} else {
			wordInfo.TestAnswer = "いられる"
		}
		break

	case volitional:
		if politeness == polite {
			if wordInfo.Japanese == "ある" {
				wordInfo.TestAnswer = "ありましょう"
			} else {
				wordInfo.TestAnswer = "いましょう"
			}
		} else if politeness == casual {
			if wordInfo.Japanese == "ある" {
				wordInfo.TestAnswer = "あろう"
			} else {
				wordInfo.TestAnswer = "いよう"
			}
		}
		break

	case want:
		if wordInfo.Japanese == "ある" {
			wordInfo.TestAnswer = "ありたい"
		} else {
			wordInfo.TestAnswer = "いたい"
		}
		break

	default:
		// Unknown form!
		break
	}
}

func createKuruAnswer(wordInfo *WordData, inputString string, politeness string, formData verbFormData) {

	switch formData.form {

	case present:
		if politeness == polite {
			wordInfo.TestAnswer = "きます"
		} else if politeness == casual {
			wordInfo.TestAnswer = inputString
		}
		break

	case presentNegative:
		if politeness == polite {
			wordInfo.TestAnswer = "きません"
		} else if politeness == casual {
			wordInfo.TestAnswer = "こない"
		}
		break

	case past:
		if politeness == polite {
			wordInfo.TestAnswer = "きました"
		} else if politeness == casual {
			wordInfo.TestAnswer = "きた"
		}
		break

	case pastNegative:
		if politeness == polite {
			wordInfo.TestAnswer = "きませんでした"
		} else if politeness == casual {
			wordInfo.TestAnswer = "こなかった"
		}
		break

	case te:
		wordInfo.TestAnswer = "きて"
		break

	case potential:
		wordInfo.TestAnswer = "こられる"
		break

	case volitional:
		if politeness == polite {
			wordInfo.TestAnswer = "きましょう"
		} else if politeness == casual {
			wordInfo.TestAnswer = "こよう"
		}
		break

	case want:
		wordInfo.TestAnswer = "きたい"
		break

	default:
		// Unknown form!
		break
	}
}

func createIkuAnswer(wordInfo *WordData, inputString string, politeness string, formData verbFormData) {

	switch formData.form {

	case present:
		if politeness == polite {
			wordInfo.TestAnswer = "いきます"
		} else if politeness == casual {
			wordInfo.TestAnswer = inputString
		}
		break

	case presentNegative:
		if politeness == polite {
			wordInfo.TestAnswer = "いきません"
		} else if politeness == casual {
			wordInfo.TestAnswer = "いかない"
		}
		break

	case past:
		if politeness == polite {
			wordInfo.TestAnswer = "いきました"
		} else if politeness == casual {
			wordInfo.TestAnswer = "いった"
		}
		break

	case pastNegative:
		if politeness == polite {
			wordInfo.TestAnswer = "いきませんでした"
		} else if politeness == casual {
			wordInfo.TestAnswer = "いかなかった"
		}
		break

	case te:
		wordInfo.TestAnswer = "いって"
		break

	case potential:
		wordInfo.TestAnswer = "いける"
		break

	case volitional:
		if politeness == polite {
			wordInfo.TestAnswer = "いきましょう"
		} else if politeness == casual {
			wordInfo.TestAnswer = "いこう"
		}
		break

	case want:
		wordInfo.TestAnswer = "いきたい"
		break

	default:
		// Unknown form!
		break
	}
}

func createSuruAnswer(wordInfo *WordData, inputString string, politeness string, formData verbFormData) {

	switch formData.form {

	case present:
		if politeness == polite {
			addVerbSuffix(inputString, &wordInfo.TestAnswer, 2, "します")
		} else if politeness == casual {
			wordInfo.TestAnswer = inputString
		}
		break

	case presentNegative:
		if politeness == polite {
			addVerbSuffix(inputString, &wordInfo.TestAnswer, 2, "しません")
		} else if politeness == casual {
			addVerbSuffix(inputString, &wordInfo.TestAnswer, 2, "しない")
		}
		break

	case past:
		if politeness == polite {
			addVerbSuffix(inputString, &wordInfo.TestAnswer, 2, "しました")
		} else if politeness == casual {
			addVerbSuffix(inputString, &wordInfo.TestAnswer, 2, "した")
		}
		break

	case pastNegative:
		if politeness == polite {
			addVerbSuffix(inputString, &wordInfo.TestAnswer, 2, "しませんでした")
		} else if politeness == casual {
			addVerbSuffix(inputString, &wordInfo.TestAnswer, 2, "しなかった")
		}
		break

	case te:
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 2, "して")
		break

	case potential:
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 2, "できる")
		break

	case volitional:
		if politeness == polite {
			addVerbSuffix(inputString, &wordInfo.TestAnswer, 2, "しましょう")
		} else if politeness == casual {
			addVerbSuffix(inputString, &wordInfo.TestAnswer, 2, "しよう")
		}
		break

	case want:
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 2, "したい")
		break

	default:
		// Unknown form!
		break
	}
}

func createIrregularVerbAnswer(wordInfo *WordData, inputString string, politeness string, formData verbFormData) {

	switch wordInfo.Japanese {

	case "ある":
		createAruIruAnswer(wordInfo, inputString, politeness, formData)
		break

	case "いる":
		createAruIruAnswer(wordInfo, inputString, politeness, formData)
		break

	case "くる":
		createKuruAnswer(wordInfo, inputString, politeness, formData)
		break

	case "いく":
		createIkuAnswer(wordInfo, inputString, politeness, formData)
		break

	default:
		if wordInfo.Subtype != "する" {
			// Unknown irregular verb!
		}
		break
	}

	if wordInfo.Subtype == "する" {
		createSuruAnswer(wordInfo, inputString, politeness, formData)
	}
}

func createVerbPolitePresentForm(wordInfo *WordData, inputString string, politeness string) {

	lastChar, _ := utf8.DecodeLastRuneInString(inputString)

	if lastChar == 'る' && wordInfo.Subtype == ichidanVerb {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "ます")
	} else if lastChar == 'る' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "ります")
	} else if lastChar == 'う' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "います")
	} else if lastChar == 'む' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "みます")
	} else if lastChar == 'く' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "きます")
	} else if lastChar == 'ぐ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "ぎます")
	} else if lastChar == 'つ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "ちます")
	} else if lastChar == 'ぶ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "びます")
	} else if lastChar == 'ぬ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "にます")
	} else if lastChar == 'す' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "します")
	}
}

func createVerbCasualPresentForm(wordInfo *WordData, inputString string, politeness string) {

	wordInfo.TestAnswer = inputString
}

func createVerbPolitePresentNegativeForm(wordInfo *WordData, inputString string, politeness string) {

	lastChar, _ := utf8.DecodeLastRuneInString(inputString)

	if lastChar == 'る' && wordInfo.Subtype == ichidanVerb {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "ません")
	} else if lastChar == 'る' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "りません")
	} else if lastChar == 'う' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "いません")
	} else if lastChar == 'む' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "みません")
	} else if lastChar == 'く' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "きません")
	} else if lastChar == 'ぐ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "ぎません")
	} else if lastChar == 'つ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "ちません")
	} else if lastChar == 'ぶ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "びません")
	} else if lastChar == 'ぬ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "にません")
	} else if lastChar == 'す' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "しません")
	}
}

func createVerbCasualPresentNegativeForm(wordInfo *WordData, inputString string, politeness string) {

	lastChar, _ := utf8.DecodeLastRuneInString(inputString)

	if lastChar == 'る' && wordInfo.Subtype == ichidanVerb {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "ない")
	} else if lastChar == 'る' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "らない")
	} else if lastChar == 'う' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "わない")
	} else if lastChar == 'む' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "まない")
	} else if lastChar == 'く' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "かない")
	} else if lastChar == 'ぐ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "がない")
	} else if lastChar == 'つ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "たない")
	} else if lastChar == 'ぶ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "ばない")
	} else if lastChar == 'ぬ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "なない")
	} else if lastChar == 'す' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "さない")
	}
}

func createVerbPolitePastForm(wordInfo *WordData, inputString string, politeness string) {

	lastChar, _ := utf8.DecodeLastRuneInString(inputString)

	if lastChar == 'る' && wordInfo.Subtype == ichidanVerb {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "ました")
	} else if lastChar == 'る' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "りました")
	} else if lastChar == 'う' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "いました")
	} else if lastChar == 'む' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "みました")
	} else if lastChar == 'く' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "きました")
	} else if lastChar == 'ぐ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "ぎました")
	} else if lastChar == 'つ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "ちました")
	} else if lastChar == 'ぶ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "びました")
	} else if lastChar == 'ぬ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "にました")
	} else if lastChar == 'す' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "しました")
	}
}

func createVerbCasualPastForm(wordInfo *WordData, inputString string, politeness string) {

	createVerbTeForm(wordInfo, inputString)

	lastChar, _ := utf8.DecodeLastRuneInString(wordInfo.TestAnswer)

	answerBytes := []byte(wordInfo.TestAnswer)
	answerCopy := fmt.Sprintf("%s", answerBytes)
	wordInfo.TestAnswer = ""

	if lastChar == 'て' {
		addVerbSuffix(answerCopy, &wordInfo.TestAnswer, 1, "た")
	} else if lastChar == 'で' {
		addVerbSuffix(answerCopy, &wordInfo.TestAnswer, 1, "だ")
	}
}

func createVerbPolitePastNegativeForm(wordInfo *WordData, inputString string, politeness string) {

	lastChar, _ := utf8.DecodeLastRuneInString(inputString)

	if lastChar == 'る' && wordInfo.Subtype == ichidanVerb {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "ませんでした")
	} else if lastChar == 'る' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "りませんでした")
	} else if lastChar == 'う' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "いませんでした")
	} else if lastChar == 'む' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "みませんでした")
	} else if lastChar == 'く' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "きませんでした")
	} else if lastChar == 'ぐ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "ぎませんでした")
	} else if lastChar == 'つ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "ちませんでした")
	} else if lastChar == 'ぶ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "びませんでした")
	} else if lastChar == 'ぬ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "にませんでした")
	} else if lastChar == 'す' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "しませんでした")
	}
}

func createVerbCasualPastNegativeForm(wordInfo *WordData, inputString string, politeness string) {

	lastChar, _ := utf8.DecodeLastRuneInString(inputString)

	if lastChar == 'る' && wordInfo.Subtype == ichidanVerb {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "なかった")
	} else if lastChar == 'る' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "らなかった")
	} else if lastChar == 'う' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "わなかった")
	} else if lastChar == 'む' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "まなかった")
	} else if lastChar == 'く' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "かなかった")
	} else if lastChar == 'ぐ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "がなかった")
	} else if lastChar == 'つ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "たなかった")
	} else if lastChar == 'ぶ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "ばなかった")
	} else if lastChar == 'ぬ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "ななかった")
	} else if lastChar == 'す' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "さなかった")
	}
}

func createVerbTeForm(wordInfo *WordData, inputString string) {

	lastChar, _ := utf8.DecodeLastRuneInString(inputString)

	if lastChar == 'る' && wordInfo.Subtype == ichidanVerb {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "て")
	} else if lastChar == 'る' || lastChar == 'う' || lastChar == 'つ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "って")
	} else if lastChar == 'む' || lastChar == 'ぶ' || lastChar == 'ぬ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "んで")
	} else if lastChar == 'く' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "いて")
	} else if lastChar == 'ぐ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "いで")
	} else if lastChar == 'す' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "して")
	} else if wordInfo.Subtype == "する" {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 2, "して")
	}
}

func createVerbPotentialForm(wordInfo *WordData, inputString string) {

	lastChar, _ := utf8.DecodeLastRuneInString(inputString)

	if lastChar == 'る' && wordInfo.Subtype == ichidanVerb {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "られる")
	} else if lastChar == 'る' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "れる")
	} else if lastChar == 'う' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "える")
	} else if lastChar == 'む' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "める")
	} else if lastChar == 'く' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "ける")
	} else if lastChar == 'ぐ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "げる")
	} else if lastChar == 'つ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "てる")
	} else if lastChar == 'ぶ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "べる")
	} else if lastChar == 'ぬ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "ねる")
	} else if lastChar == 'す' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "せる")
	}
}

func createVerbPoliteVolitionalForm(wordInfo *WordData, inputString string, politeness string) {

	lastChar, _ := utf8.DecodeLastRuneInString(inputString)

	if lastChar == 'る' && wordInfo.Subtype == ichidanVerb {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "ましょう")
	} else if lastChar == 'る' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "りましょう")
	} else if lastChar == 'う' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "いましょう")
	} else if lastChar == 'む' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "みましょう")
	} else if lastChar == 'く' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "きましょう")
	} else if lastChar == 'ぐ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "ぎましょう")
	} else if lastChar == 'つ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "ちましょう")
	} else if lastChar == 'ぶ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "びましょう")
	} else if lastChar == 'ぬ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "にましょう")
	} else if lastChar == 'す' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "しましょう")
	}
}

func createVerbCasualVolitionalForm(wordInfo *WordData, inputString string, politeness string) {

	lastChar, _ := utf8.DecodeLastRuneInString(inputString)

	if lastChar == 'る' && wordInfo.Subtype == ichidanVerb {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "よう")
	} else if lastChar == 'る' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "ろう")
	} else if lastChar == 'う' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "おう")
	} else if lastChar == 'む' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "もう")
	} else if lastChar == 'く' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "こう")
	} else if lastChar == 'ぐ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "ごう")
	} else if lastChar == 'つ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "とう")
	} else if lastChar == 'ぶ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "ぼう")
	} else if lastChar == 'ぬ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "のう")
	} else if lastChar == 'す' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "そう")
	}
}

func createVerbWantForm(wordInfo *WordData, inputString string) {

	lastChar, _ := utf8.DecodeLastRuneInString(inputString)

	if lastChar == 'る' && wordInfo.Subtype == ichidanVerb {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "たい")
	} else if lastChar == 'る' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "りたい")
	} else if lastChar == 'う' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "いたい")
	} else if lastChar == 'む' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "みたい")
	} else if lastChar == 'く' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "きたい")
	} else if lastChar == 'ぐ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "ぎたい")
	} else if lastChar == 'つ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "ちたい")
	} else if lastChar == 'ぶ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "びたい")
	} else if lastChar == 'ぬ' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "にたい")
	} else if lastChar == 'す' {
		addVerbSuffix(inputString, &wordInfo.TestAnswer, 1, "したい")
	}
}

func createVerbAnswer(wordInfo *WordData, politeness string, formData verbFormData, showKanji bool, language string) {

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

	if isVerbIrregular(wordInfo.Japanese, wordInfo.Subtype) {
		createIrregularVerbAnswer(wordInfo, inputString, politeness, formData)
	} else {

		switch formData.form {

		case present:
			if politeness == polite {
				createVerbPolitePresentForm(wordInfo, inputString, politeness)
			} else if politeness == casual {
				createVerbCasualPresentForm(wordInfo, inputString, politeness)
			}
			break

		case presentNegative:
			if politeness == polite {
				createVerbPolitePresentNegativeForm(wordInfo, inputString, politeness)
			} else if politeness == casual {
				createVerbCasualPresentNegativeForm(wordInfo, inputString, politeness)
			}
			break

		case past:
			if politeness == polite {
				createVerbPolitePastForm(wordInfo, inputString, politeness)
			} else if politeness == casual {
				createVerbCasualPastForm(wordInfo, inputString, politeness)
			}
			break

		case pastNegative:
			if politeness == polite {
				createVerbPolitePastNegativeForm(wordInfo, inputString, politeness)
			} else if politeness == casual {
				createVerbCasualPastNegativeForm(wordInfo, inputString, politeness)
			}
			break

		case te:
			createVerbTeForm(wordInfo, inputString)
			break

		case potential:
			createVerbPotentialForm(wordInfo, inputString)
			break

		case volitional:
			if politeness == polite {
				createVerbPoliteVolitionalForm(wordInfo, inputString, politeness)
			} else if politeness == casual {
				createVerbCasualVolitionalForm(wordInfo, inputString, politeness)
			}
			break

		case want:
			createVerbWantForm(wordInfo, inputString)
			break
		}
	}
}
