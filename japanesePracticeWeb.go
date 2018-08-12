package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	sheets "google.golang.org/api/sheets/v4"

	s "strings"
)

type japanesePracticeWeb struct {
	settings      Settings
	sheetsService *sheets.Service
	db            wordDb
}

type WordData struct {
	WordType   int
	English    string
	Japanese   string
	Kanji      string
	Subtype    string
	JlptLevel  int
	TestAnswer string
}

type adjectiveData struct {
	English  string
	Japanese string
}

type verbFormData struct {
	form           string
	usesPoliteness bool
}

type wordDb struct {
	wordData           []WordData
	verbPolitenessData []string
	verbFormData       []verbFormData
	adjectiveFormData  []string
}

type OptionButton struct {
	Name       string
	Value      string
	IsDisabled bool
	IsChecked  bool
	Text       string
}

type TestVariables struct {
	PageTitle                  string
	DisplayOptionButtons       []OptionButton
	WordOptionButtons          []OptionButton
	PolitenessOptionButtons    []OptionButton
	FormOptionButtons          []OptionButton
	AdjectiveFormOptionButtons []OptionButton
	TestButtonDisabled         bool
	TestWords                  []WordData
	TestForms                  []string
}

const (
	verb           = iota
	adjective      = iota
	verbForms      = iota
	adjectiveForms = iota
)

const (
	verbSheet          = "verbs"
	verbFormSheet      = "verbforms"
	adjectiveSheet     = "adjectives"
	adjectiveFormSheet = "adjectiveforms"
)

const (
	casual   = "Casual"
	polite   = "Polite"
	agnostic = "None"
)

const (
	present         = "Present"
	presentNegative = "Present Negative"
	past            = "Past"
	pastNegative    = "Past Negative"
	te              = "Te"
	potential       = "Potential"
	volitional      = "Volitional"
	want            = "Want"

	adjectiveFormPrefix = "adj-"
)

const (
	irregularVerb = "Irregular"
	ichidanVerb   = "Ichidan"
	godanVerb     = "Godan"
	regularVerb   = "Regular"

	iAdjective  = "い-adjective"
	naAdjective = "な-adjective"
)

const (
	kana  = "Kana"
	kanji = "Kanji"
)

var jp japanesePracticeWeb
var Title string

func main() {

	jp = japanesePracticeWeb{}
	if !jp.Setup() {
		return
	}

	fmt.Print("Reading input spreadsheet... ")

	jp.db = wordDb{}
	jp.db.PopulateWordDatabase(jp)

	fmt.Print("done\n\n")

	Title = "Japanese Vocab Practice"

	rand.Seed(time.Now().UTC().UnixNano())

	http.HandleFunc("/", DisplayOptionButtons)
	http.HandleFunc("/controls", WordTypeSelected)

	log.Fatal(http.ListenAndServe(getPort(), nil))

	// runtime.Goexit()
}

func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":8080"
}

func (db *wordDb) PopulateWordDatabase(jp japanesePracticeWeb) bool {

	var numSheets int
	jp.sheetsService, numSheets = SheetsSetupJWT(jp.settings.GoogleSheetID)

	db.wordData = make([]WordData, 0)
	db.verbPolitenessData = make([]string, 0)
	db.verbFormData = make([]verbFormData, 0)
	db.adjectiveFormData = make([]string, 0)

	for sheetIndex := 0; sheetIndex < numSheets; sheetIndex++ {

		// Skip the first (header) row in all sheets, start with A2
		sheetData := GetDataForSheetIndex(jp.sheetsService, jp.settings.GoogleSheetID, sheetIndex, "A2", "E")
		sheetName := GetSheetName(jp.sheetsService, jp.settings.GoogleSheetID, sheetIndex)
		sheetName = s.ToLower(sheetName)

		var wordType int
		if s.Contains(sheetName, verbFormSheet) {
			wordType = verbForms
		} else if s.Contains(sheetName, adjectiveFormSheet) {
			wordType = adjectiveForms
		} else if s.Contains(sheetName, verbSheet) {
			wordType = verb
		} else if s.Contains(sheetName, adjectiveSheet) {
			wordType = adjective
		}

		for i, row := range sheetData.Values {

			// Skip if no data in this row
			var foundData bool
			for i := 0; i < 5; i++ {
				if len(row[i].(string)) != 0 {
					foundData = true
					break
				}
			}

			if foundData != true {
				break
			}

			// Decide which type of sheet this is and populate data accordingly
			switch wordType {

			case verb:
				var newWord WordData
				newWord.WordType = verb
				newWord.English = row[0].(string)
				newWord.Subtype = row[1].(string)
				newWord.Japanese = row[2].(string)
				newWord.Kanji = row[3].(string)
				newWord.JlptLevel, _ = strconv.Atoi(row[4].(string))
				db.wordData = append(db.wordData, newWord)
				//fmt.Printf("Found verb: %s\n", newWord.English)
				break

			case adjective:
				var newWord WordData
				newWord.WordType = adjective
				newWord.English = row[0].(string)
				newWord.Subtype = row[1].(string)
				newWord.Japanese = row[2].(string)
				newWord.Kanji = row[3].(string)
				newWord.JlptLevel, _ = strconv.Atoi(row[4].(string))
				db.wordData = append(db.wordData, newWord)
				break

			case verbForms:
				if len(row[0].(string)) > 0 {
					db.verbPolitenessData = append(db.verbPolitenessData, row[0].(string))
				}

				if len(row[1].(string)) > 0 {
					var formData verbFormData
					formData.form = row[1].(string)
					formData.usesPoliteness = s.ToLower(row[2].(string)) == "y"
					db.verbFormData = append(db.verbFormData, formData)
				}
				break

			case adjectiveForms:
				if len(row[0].(string)) > 0 {
					db.adjectiveFormData = append(db.adjectiveFormData, adjectiveFormPrefix+row[0].(string))
				}
				break
			}

			// Failsafe
			if i == 10000 {
				break
			}
		}
	}

	return true
}

func (jp *japanesePracticeWeb) Setup() bool {

	err := jp.settings.Setup()
	if err != nil {
		fmt.Printf("%s", err)
		return false
	}

	return true
}

func DisplayOptionButtons(w http.ResponseWriter, r *http.Request) {

	// Display some radio buttons to the user
	wordOptionButtons := SetupWordOptions(false, false)
	displayOptionButtons := SetupDisplayOptions("kana")

	MyTestVariables := TestVariables{
		PageTitle:            Title,
		DisplayOptionButtons: displayOptionButtons,
		WordOptionButtons:    wordOptionButtons,
		TestButtonDisabled:   true,
	}

	t, err := template.ParseFiles("jpMain.html") //parse the html file homepage.html
	if err != nil {                              // if there is an error
		log.Print("template parsing error: ", err) // log it
	}

	err = t.Execute(w, MyTestVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {                     // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

func SanitizeHtmlName(input string) string {
	return (s.Replace(s.ToLower(input), " ", "", -1))
}

func SetupDisplayOptions(checkedName string) []OptionButton {

	displayOptionButtons := []OptionButton{
		OptionButton{"displayRadio", "kana", false, checkedName == "kana", "Kana"},
		OptionButton{"displayRadio", "kanji", false, checkedName == "kanji", "Kanji"},
	}

	return displayOptionButtons
}

func SetupWordOptions(hasVerbs bool, hasAdjectives bool) []OptionButton {

	wordOptionButtons := []OptionButton{
		OptionButton{"wordVerb", "verb", false, hasVerbs, "Verbs"},
		OptionButton{"wordAdjective", "adjective", false, hasAdjectives, "Adjectives"},
	}

	return wordOptionButtons
}

func SetupVerbOptions(politenessData []string, formData []string) ([]OptionButton, []OptionButton) {

	numVerbPoliteness := len(jp.db.verbPolitenessData)
	politenessOptionButtons := make([]OptionButton, 0)
	for i := 0; i < numVerbPoliteness; i++ {
		newButton := OptionButton{SanitizeHtmlName(jp.db.verbPolitenessData[i]), jp.db.verbPolitenessData[i], false, politenessData[i] != "", jp.db.verbPolitenessData[i]}
		politenessOptionButtons = append(politenessOptionButtons, newButton)
	}

	numVerbForms := len(jp.db.verbFormData)
	formOptionButtons := make([]OptionButton, 0)
	for i := 0; i < numVerbForms; i++ {
		formName := jp.db.verbFormData[i].form
		newButton := OptionButton{SanitizeHtmlName(formName), formName, false, formData[i] != "", formName}
		formOptionButtons = append(formOptionButtons, newButton)
	}

	return politenessOptionButtons, formOptionButtons
}

func SetupAdjectiveOptions(formData []string) []OptionButton {

	numAdjectiveForms := len(jp.db.adjectiveFormData)
	formOptionButtons := make([]OptionButton, 0)
	for i := 0; i < numAdjectiveForms; i++ {
		formName := jp.db.adjectiveFormData[i]
		displayFormName := s.Replace(formName, adjectiveFormPrefix, "", 1)
		newButton := OptionButton{SanitizeHtmlName(formName), formName, false, formData[i] != "", displayFormName}
		formOptionButtons = append(formOptionButtons, newButton)
	}

	return formOptionButtons
}

// RunTest builds the word list from the requested options
func RunTest(form url.Values) ([]WordData, []string) {

	verbOption := form.Get("wordVerb")
	adjectiveOption := form.Get("wordAdjective")

	// Build word list
	outputWords := make([]WordData, 0)
	outputForms := make([]string, 0)

	for i := 0; i < len(jp.db.wordData)-1; i++ {

		if (verbOption != "" && jp.db.wordData[i].WordType == verb) || (adjectiveOption != "" && jp.db.wordData[i].WordType == adjective) {
			outputWords = append(outputWords, jp.db.wordData[i])
		}
	}

	// Randomize the word list
	for i := len(outputWords) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		outputWords[i], outputWords[j] = outputWords[j], outputWords[i]
	}

	numVerbPoliteness := 0
	numVerbForms := 0
	allowedVerbPoliteness := make([]string, 0)
	allowedVerbForms := make([]verbFormData, 0)
	verbPolitenessAgnosticAllowed := false

	if verbOption != "" {
		// Verb politeness options
		numPolitenessOptions := len(jp.db.verbPolitenessData)

		for i := 0; i < numPolitenessOptions; i++ {
			option := form.Get(s.ToLower(jp.db.verbPolitenessData[i]))
			if option != "" && option != agnostic {
				allowedVerbPoliteness = append(allowedVerbPoliteness, option)
			}

			if option == agnostic {
				verbPolitenessAgnosticAllowed = true
			}
		}

		// Politeness failsafe
		numVerbPoliteness = len(allowedVerbPoliteness)
		if numVerbPoliteness == 0 && !verbPolitenessAgnosticAllowed {
			// If agnostic is not chosen, give us the default politeness level
			allowedVerbPoliteness = append(allowedVerbPoliteness, jp.db.verbPolitenessData[0])
			numVerbPoliteness = 1
		}

		// Verb form options
		numFormsOptions := len(jp.db.verbFormData)

		for j := 0; j < numFormsOptions; j++ {
			option := form.Get(SanitizeHtmlName(jp.db.verbFormData[j].form))
			if option != "" &&
				((verbPolitenessAgnosticAllowed && !jp.db.verbFormData[j].usesPoliteness) ||
					(numVerbPoliteness > 0 && jp.db.verbFormData[j].usesPoliteness)) {
				allowedVerbForms = append(allowedVerbForms, jp.db.verbFormData[j])
			}
		}
		numVerbForms = len(allowedVerbForms)

		// Form failsafe
		if numVerbForms == 0 {
			// Give us one form if none have been selected
			if numVerbPoliteness > 0 {
				allowedVerbForms = append(allowedVerbForms, jp.db.verbFormData[0])
			} else {
				// Will only deliver Te form if agnostic politeness is all that is picked
				for i := 0; i < numFormsOptions; i++ {
					if jp.db.verbFormData[i].form == te {
						allowedVerbForms = append(allowedVerbForms, jp.db.verbFormData[i])
						break
					}
				}
			}
			numVerbForms = 1
		}
	}

	numAdjectiveForms := 0
	allowedAdjectiveForms := make([]string, 0)
	if adjectiveOption != "" {

		// Adjective form options
		numFormsOptions := len(jp.db.adjectiveFormData)

		for j := 0; j < numFormsOptions; j++ {
			option := form.Get(SanitizeHtmlName(jp.db.adjectiveFormData[j]))
			if option != "" {
				displayForm := s.Replace(jp.db.adjectiveFormData[j], adjectiveFormPrefix, "", 1)
				allowedAdjectiveForms = append(allowedAdjectiveForms, displayForm)
			}
		}
		numAdjectiveForms = len(allowedAdjectiveForms)

		// Form failsafe
		if numAdjectiveForms == 0 {
			// Give us one form if none have been selected
			displayForm := s.Replace(jp.db.adjectiveFormData[0], adjectiveFormPrefix, "", 1)
			allowedAdjectiveForms = append(allowedAdjectiveForms, displayForm)
			numAdjectiveForms = 1
		}
	}

	// Pass Kanji flag into answer routines
	displayOption := form.Get("displayRadio")
	var showKanji bool
	if displayOption == "kanji" {
		showKanji = true
	}

	// Loop over the list and setup the politeness/form strings and answers
	for i := 0; i < len(outputWords); i++ {

		wordInfo := outputWords[i]

		var question string

		// Pick verb info
		if wordInfo.WordType == verb {

			formData := allowedVerbForms[rand.Intn(numVerbForms)]

			var politeness string
			if numVerbPoliteness > 0 {
				politeness = allowedVerbPoliteness[rand.Intn(numVerbPoliteness)]
			}

			if formData.usesPoliteness {
				question += politeness
				question += " "
			}

			question += formData.form
			question += " form"

			createVerbAnswer(&outputWords[i], politeness, formData, showKanji)

		} else if wordInfo.WordType == adjective {

			formData := allowedAdjectiveForms[rand.Intn(numAdjectiveForms)]

			question += formData
			question += " form"

			createAdjectiveAnswer(&outputWords[i], formData, showKanji)
		}

		outputForms = append(outputForms, question)
	}

	return outputWords, outputForms
}

func WordTypeSelected(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	verbOption := r.Form.Get("wordVerb")
	adjectiveOption := r.Form.Get("wordAdjective")
	runTestState := r.Form.Get("runTest")

	var runTest bool
	var testWords []WordData
	var testForms []string
	if len(runTestState) > 0 {
		runTest = true
		testWords, testForms = RunTest(r.Form)
	}

	hasVerbs := false
	var politenessOptionButtons []OptionButton
	var formOptionButtons []OptionButton

	if verbOption != "" {

		hasVerbs = true

		numPolitenessOptions := len(jp.db.verbPolitenessData)
		politenessOptionData := make([]string, 0)
		for i := 0; i < numPolitenessOptions; i++ {
			politenessOptionData = append(politenessOptionData, r.Form.Get(SanitizeHtmlName(jp.db.verbPolitenessData[i])))
		}

		numFormOptions := len(jp.db.verbFormData)
		formOptionData := make([]string, 0)
		for i := 0; i < numFormOptions; i++ {
			formOptionData = append(formOptionData, r.Form.Get(SanitizeHtmlName(jp.db.verbFormData[i].form)))
		}

		politenessOptionButtons, formOptionButtons = SetupVerbOptions(politenessOptionData, formOptionData)
	}

	hasAdjectives := false
	var adjectiveFormOptionButtons []OptionButton

	if adjectiveOption != "" {

		hasAdjectives = true

		numAdjectiveFormOptions := len(jp.db.adjectiveFormData)
		adjectiveFormOptionData := make([]string, 0)
		for i := 0; i < numAdjectiveFormOptions; i++ {
			adjectiveFormOptionData = append(adjectiveFormOptionData, r.Form.Get(SanitizeHtmlName(jp.db.adjectiveFormData[i])))
		}

		adjectiveFormOptionButtons = SetupAdjectiveOptions(adjectiveFormOptionData)
	}

	wordOptionButtons := SetupWordOptions(hasVerbs, hasAdjectives)
	checkedRadio := r.Form.Get("displayRadio")

	displayOptionButtons := SetupDisplayOptions(checkedRadio)

	var MyTestVariables TestVariables

	if !hasVerbs && !hasAdjectives {
		MyTestVariables.TestButtonDisabled = true
	}

	MyTestVariables.PageTitle = Title
	MyTestVariables.DisplayOptionButtons = displayOptionButtons
	MyTestVariables.WordOptionButtons = wordOptionButtons

	if hasVerbs {
		MyTestVariables.PolitenessOptionButtons = politenessOptionButtons
		MyTestVariables.FormOptionButtons = formOptionButtons
	}

	if hasAdjectives {
		MyTestVariables.AdjectiveFormOptionButtons = adjectiveFormOptionButtons
	}

	if runTest {

		MyTestVariables.TestWords = testWords
		MyTestVariables.TestForms = testForms
	}

	// generate page by passing page variables into template
	t, err := template.ParseFiles("jpMain.html") //parse the html file homepage.html
	if err != nil {                              // if there is an error
		log.Print("template parsing error: ", err) // log it
	}

	err = t.Execute(w, MyTestVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {                     // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}
