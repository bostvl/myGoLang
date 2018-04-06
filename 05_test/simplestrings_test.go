package simplestrings

import "testing"

const weekdays = "Monday Tuesday Wednesday Thursday Friday"

// test that Tuesday is a weekday
func TestTuesdayIsWeekday(t *testing.T) {
	actualResult := Index(weekdays, "Tuesday")
	if actualResult == -1 {
		t.Fatalf("Expected number >= 0 but got -1")
	}
}

// test that Sunday is not a weekday
func TestSundayIsNotWeekday(t *testing.T) {
	actualResult := Index(weekdays, "Sunday")
	var expectedResult = -1
	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", string(expectedResult), string(actualResult))
	}
}

// test that an empty search string returns 0
func TestEmptyStringSearch(t *testing.T) {
	actualResult := Index(weekdays, "")
	var expectedResult = 0
	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", string(expectedResult), string(actualResult))
	}
}

// test that the string Monday is not found in the empty string
func TestMondayIsNotInEmptyString(t *testing.T) {
	actualResult := Index("", "Mondey")
	var expectedResult = -1
	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", string(expectedResult), string(actualResult))
	}
}
