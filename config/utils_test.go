package config

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type UtilsTestSuite struct {
	suite.Suite
}

func (suite *UtilsTestSuite) SetupTest() {
	viper.New()
	viper.AutomaticEnv()
}

func (suite *UtilsTestSuite) TestCheckKeyDoesNotPanicWhenKeyIsSet() {
	key := "TEST_KEY"
	value := 12

	os.Setenv(key, fmt.Sprintf("%d", value))
	defer os.Unsetenv(key)

	suite.NotPanics(func() {
		checkKey(key)
	})
}

func (suite *UtilsTestSuite) TestCheckKeyPanicsWhenKeyIsNotSet() {
	key := "TEST_KEY"

	suite.PanicsWithValue("Missing key: TEST_KEY", func() {
		checkKey(key)
	})
}

func (suite *UtilsTestSuite) TestAllUtilFunction() {
	intKey := "INT_KEY"
	intValue := 10
	stringKey := "STRING_KEY"
	stringValue := "TEST_STRING"
	boolKey := "BOOL_KEY"
	stringSliceKey := "STRING_SLICE_KEY"
	stringSliceValue := "TEST_STRING_1,TEST_STRING_2,TEST_STRING_3"
	intSliceKey := "INT_SLICE_KEY"
	intSliceValue := "10,12,13,14,53"
	stringInterfaceMapKey := "STRING_INTERFACE_MAP_KEY"
	stringInterfaceMapValue := `{"TEST_PARENT_KEY_1": {"TEST_CHILD_KEY": "11"}, "TEST_PARENT_KEY_2": {"TEST_CHILD_KEY": "12"}}`
	stringStringMapKey := "STRING_STRING_MAP_KEY"
	stringStringMapValue := `{"TEST_KEY_1": "VALUE_1", "TEST_KEY_2": "VALUE_2"}`

	os.Setenv(intKey, fmt.Sprintf("%d", intValue))
	defer os.Unsetenv(intKey)
	os.Setenv(stringKey, stringValue)
	defer os.Unsetenv(stringKey)
	os.Setenv(boolKey, strconv.FormatBool(true))
	defer os.Unsetenv(boolKey)
	os.Setenv(stringSliceKey, stringSliceValue)
	defer os.Unsetenv(stringSliceKey)
	os.Setenv(intSliceKey, intSliceValue)
	defer os.Unsetenv(intSliceKey)
	os.Setenv(stringInterfaceMapKey, stringInterfaceMapValue)
	defer os.Unsetenv(stringInterfaceMapKey)
	os.Setenv(stringStringMapKey, stringStringMapValue)
	defer os.Unsetenv(stringStringMapKey)

	suite.Equal(intValue, getInt(intKey, true))
	suite.Equal(stringValue, getString(stringKey, true))
	suite.Equal(true, getBool(boolKey, true))
	suite.Equal([]string{"TEST_STRING_1", "TEST_STRING_2", "TEST_STRING_3"}, getStringSlice(stringSliceKey, true))
	suite.Equal([]int{10, 12, 13, 14, 53}, getIntSlice(intSliceKey, true))
	suite.Equal(map[string]interface{}{
		"TEST_PARENT_KEY_1": map[string]interface{}{"TEST_CHILD_KEY": "11"},
		"TEST_PARENT_KEY_2": map[string]interface{}{"TEST_CHILD_KEY": "12"},
	}, getStringMapInterface(stringInterfaceMapKey, true))
	suite.Equal(map[string]string{"TEST_KEY_1": "VALUE_1", "TEST_KEY_2": "VALUE_2"},
		getStringMapString(stringStringMapKey, true))
}

func TestUtilsTestSuite(t *testing.T) {
	suite.Run(t, new(UtilsTestSuite))
}
