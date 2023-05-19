package utilities

import "time"

func GetCurrentDateTimeInUTC() time.Time {
	return formatTimeToISO8601(time.Now().UTC())
}

func GetCurrentDateTimeInTimezone(timezone string) time.Time {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return formatTimeToISO8601(time.Now().UTC())
	}
	return formatTimeToISO8601(time.Now().In(loc))
}

func formatTimeToISO8601(t time.Time) time.Time {
	formattedTime, _ := time.Parse(time.RFC3339, t.Format(time.RFC3339))
	return formattedTime
}
