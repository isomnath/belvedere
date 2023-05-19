package utilities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type DateTimeUtilitiesTestSuite struct {
	suite.Suite
}

func (suite *DateTimeUtilitiesTestSuite) TestGetCurrentDateTimeInUTC() {
	t := GetCurrentDateTimeInUTC()
	suite.Equal(time.UTC, t.Location())
}

func (suite *DateTimeUtilitiesTestSuite) TestGetCurrentDateTimeInTimezone() {
	tUTC := GetCurrentDateTimeInUTC()
	tWIB := GetCurrentDateTimeInTimezone("Asia/Jakarta")
	suite.Equal(tUTC, tWIB.UTC())
}

func (suite *DateTimeUtilitiesTestSuite) TestGetCurrentDateTimeInTimezoneWhenTimezoneIsIncorrect() {
	t := GetCurrentDateTimeInTimezone("invalid time zone")
	suite.Equal(time.UTC, t.Location())
}

func TestDateTimeUtilitiesTestSuite(t *testing.T) {
	suite.Run(t, new(DateTimeUtilitiesTestSuite))
}
