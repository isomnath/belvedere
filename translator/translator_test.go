package translator

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/isomnath/belvedere/config"
	"github.com/isomnath/belvedere/log"

	"github.com/stretchr/testify/suite"
)

type TranslatorTestSuite struct {
	suite.Suite
}

func (suite *TranslatorTestSuite) SetupSuite() {
	config.LoadBaseConfig()
	config.LoadTranslationsConfig()
	log.Setup()
	suite.prepareTranslationFiles()
	Initialize()
}

func (suite *TranslatorTestSuite) TearDownSuite() {
	os.RemoveAll(config.GetTranslationConfig().Path())
}

func (suite *TranslatorTestSuite) TestTranslateShouldReturnUntranslatedMessage() {
	Kill()
	translatedMessage := Translate("some message", "id")
	suite.Equal("some message", translatedMessage)
}

func (suite *TranslatorTestSuite) TestLoadTranslationsShouldPanicIfTranslationFileHasInvalidData() {
	_ = ioutil.WriteFile(fmt.Sprintf("%s/th.json", config.GetTranslationConfig().Path()), []byte("{\"ERROR_INTERNAL_SERVER\": \"something went wrong\""), 0644)
	suite.Panics(func() {
		loadTranslations()
	})
	_ = os.Remove(fmt.Sprintf("%s/th.json", config.GetTranslationConfig().Path()))
}

func (suite *TranslatorTestSuite) TestLoadTranslationsPathShouldPanicIfWalkOfTranslationsPathFails() {
	validPath := config.GetTranslationConfig().Path()
	invalidPath := "./invalidPath"
	_ = os.Setenv("TRANSLATIONS_PATH", invalidPath)
	config.LoadTranslationsConfig()
	suite.Panics(func() {
		loadTranslations()
	})
	_ = os.Setenv("TRANSLATIONS_PATH", validPath)
	config.LoadTranslationsConfig()
}

func (suite *TranslatorTestSuite) TestTranslateShouldReturnTranslatedMessageBasedOnInputLanguage() {
	translatedMessage := Translate("ERROR_INTERNAL_SERVER", "id")
	suite.Equal("ada yang salah", translatedMessage)
}

func (suite *TranslatorTestSuite) TestTranslateShouldReturnUntranslatedMessageWhenTranslationKeyNotFoundInBundle() {
	translatedMessage := Translate("SOME_KEY", "id")
	suite.Equal("SOME_KEY", translatedMessage)
}

func (suite *TranslatorTestSuite) prepareTranslationFiles() {
	path := config.GetTranslationConfig().Path()
	_ = os.Mkdir(config.GetTranslationConfig().Path(), os.ModePerm)
	_ = ioutil.WriteFile(fmt.Sprintf("%s/en.json", path), []byte("{\"ERROR_INTERNAL_SERVER\": \"something went wrong\"}"), 0644)
	_ = ioutil.WriteFile(fmt.Sprintf("%s/id.json", path), []byte("{\"ERROR_INTERNAL_SERVER\": \"ada yang salah\"}"), 0644)
}

func TestTranslatorTestSuite(t *testing.T) {
	suite.Run(t, new(TranslatorTestSuite))
}
