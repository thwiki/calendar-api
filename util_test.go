package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanitizeDateValid(t *testing.T) {
	var date string
	var err error

	// Standard
	date, err = SanitizeDate("2010-01-01")
	assert.Equal(t, "2010-01-01", date)
	assert.Nil(t, err)

	date, err = SanitizeDate("2020-10-24")
	assert.Equal(t, "2020-10-24", date)
	assert.Nil(t, err)

	// Without Padding Zero
	date, err = SanitizeDate("2010-1-1")
	assert.Equal(t, "2010-01-01", date)
	assert.Nil(t, err)

	// Any Year
	date, err = SanitizeDate("0001-12-31")
	assert.Equal(t, "0001-12-31", date)
	assert.Nil(t, err)
}

func TestSanitizeDateInvalid(t *testing.T) {
	var date string
	var err error

	// Empty
	date, err = SanitizeDate("")
	assert.Equal(t, "", date)
	assert.Error(t, &InvalidDateError{}, err)

	// Extra Spaces
	date, err = SanitizeDate("2010-01-01 ")
	assert.Equal(t, "", date)
	assert.Error(t, &InvalidDateError{}, err)

	date, err = SanitizeDate(" 2010-01-01")
	assert.Equal(t, "", date)
	assert.Error(t, &InvalidDateError{}, err)

	// Gibberish
	date, err = SanitizeDate("123cat")
	assert.Equal(t, "", date)
	assert.Error(t, &InvalidDateError{}, err)

	date, err = SanitizeDate("123-cat-dog")
	assert.Equal(t, "", date)
	assert.Error(t, &InvalidDateError{}, err)

	// Missing Date
	date, err = SanitizeDate("2010-01")
	assert.Equal(t, "", date)
	assert.Error(t, &InvalidDateError{}, err)

	// Missing Section
	date, err = SanitizeDate("2010-01-")
	assert.Equal(t, "", date)
	assert.Error(t, &InvalidDateError{}, err)

	// Out of Range Month
	date, err = SanitizeDate("2010-13-01")
	assert.Equal(t, "", date)
	assert.Error(t, &InvalidDateError{}, err)

	date, err = SanitizeDate("2010-00-01")
	assert.Equal(t, "", date)
	assert.Error(t, &InvalidDateError{}, err)

	// Out of Range Day
	date, err = SanitizeDate("2010-12-32")
	assert.Equal(t, "", date)
	assert.Error(t, &InvalidDateError{}, err)

	date, err = SanitizeDate("2010-01-00")
	assert.Equal(t, "", date)
	assert.Error(t, &InvalidDateError{}, err)
}
