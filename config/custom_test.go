package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type CustomConfigTestSuite struct {
	suite.Suite
}

func (suite *CustomConfigTestSuite) SetupTest() {
	_ = os.Setenv("TEST_KEY_ONE", fmt.Sprintf("%d", 20))
	_ = os.Setenv("TEST_KEY_TWO", fmt.Sprintf("%s", "test_value"))
	_ = os.Setenv("TEST_KEY_THREE", fmt.Sprintf("%t", true))
	_ = os.Setenv("TEST_KEY_FOUR", fmt.Sprintf("%s", "val_1,val_2,val_3"))
	_ = os.Setenv("TEST_KEY_FIVE", fmt.Sprintf("%s", "1,2,3"))

	viper.New()
	viper.AutomaticEnv()
}

func (suite *CustomConfigTestSuite) TearDownTest() {
	_ = os.Unsetenv("TEST_KEY_ONE")
	_ = os.Unsetenv("TEST_KEY_TWO")
	_ = os.Unsetenv("TEST_KEY_THREE")
	_ = os.Unsetenv("TEST_KEY_FOUR")
	_ = os.Unsetenv("TEST_KEY_FIVE")
}

func (suite *CustomConfigTestSuite) TestCustomConfigUnmarshalSuccess() {
	type TestStruct struct {
		TestKeyOne   int      `mapstructure:"TEST_KEY_ONE"`
		TestKeyTwo   string   `mapstructure:"TEST_KEY_TWO"`
		TestKeyThree bool     `mapstructure:"TEST_KEY_THREE"`
		TestKeyFour  []string `mapstructure:"TEST_KEY_FOUR"`
		TestKeyFive  []int    `mapstructure:"TEST_KEY_FIVE"`
	}
	var ts TestStruct
	customConfig(&ts)

	suite.Equal(TestStruct{
		TestKeyOne:   20,
		TestKeyTwo:   "test_value",
		TestKeyThree: true,
		TestKeyFour:  []string{"val_1", "val_2", "val_3"},
		TestKeyFive:  []int{1, 2, 3},
	}, ts)
}

func TestCustomConfigTestSuite(t *testing.T) {
	suite.Run(t, new(CustomConfigTestSuite))
}
