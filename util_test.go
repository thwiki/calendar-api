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

func TestFormatSMWDateValid(t *testing.T) {
	var date string
	var err error

	// Standard
	date, err = SanitizeSMWDate("1/2010/01/01")
	assert.Equal(t, "2010-01-01", date)
	assert.Nil(t, err)

	// Standard with Time
	date, err = SanitizeSMWDate("1/2010/01/01/08/12/20/0")
	assert.Equal(t, "2010-01-01", date)
	assert.Nil(t, err)

	// Non Standard
	date, err = SanitizeSMWDate("2010/01/01")
	assert.Equal(t, "2010-01-01", date)
	assert.Nil(t, err)

	// Missing Time Part
	date, err = SanitizeSMWDate("1/2010/01/01/08")
	assert.Equal(t, "2010-01-01", date)
	assert.Nil(t, err)
}

func TestFormatSMWDateInalid(t *testing.T) {
	var date string
	var err error

	// Missing Date Part
	date, err = SanitizeSMWDate("1/2010/01")
	assert.Equal(t, "", date)
	assert.Error(t, &InvalidDateError{}, err)

	date, err = SanitizeSMWDate("2010/01")
	assert.Equal(t, "", date)
	assert.Error(t, &InvalidDateError{}, err)
}

func TestSanitizeWikiText(t *testing.T) {
	var result string

	// Normal Text
	result = SanitizeWikiText("abcd efgh")
	assert.Equal(t, "abcd efgh", result)

	// Internal Links
	result = SanitizeWikiText("abcd [[cat]] efgh")
	assert.Equal(t, "abcd cat efgh", result)

	// Internal Links with Text
	result = SanitizeWikiText("abcd [[cat|dog]] efgh")
	assert.Equal(t, "abcd dog efgh", result)

	// Internal Links with Text and Spaces
	result = SanitizeWikiText("abcd[[ cat | dog ]]efgh")
	assert.Equal(t, "abcd dog efgh", result)

	// External Links
	result = SanitizeWikiText("abcd [https://thwiki.cc/] efgh")
	assert.Equal(t, "abcd  efgh", result)

	// External Links with Text
	result = SanitizeWikiText("abcd [https://thwiki.cc/ dog] efgh")
	assert.Equal(t, "abcd dog efgh", result)

	// External Links with Text and Spaces
	result = SanitizeWikiText("abcd[https://thwiki.cc/  dog ]efgh")
	assert.Equal(t, "abcd dog efgh", result)

	// Normal Quotes
	result = SanitizeWikiText("abcd 'cat' efgh")
	assert.Equal(t, "abcd 'cat' efgh", result)

	// Bold
	result = SanitizeWikiText("abcd '''cat''' efgh")
	assert.Equal(t, "abcd <b>cat</b> efgh", result)

	// Bold with Spaces
	result = SanitizeWikiText("abcd''' cat '''efgh")
	assert.Equal(t, "abcd<b> cat </b>efgh", result)

	// Italic
	result = SanitizeWikiText("abcd ''cat'' efgh")
	assert.Equal(t, "abcd <i>cat</i> efgh", result)

	// Italic with Spaces
	result = SanitizeWikiText("abcd'' cat ''efgh")
	assert.Equal(t, "abcd<i> cat </i>efgh", result)

	// Bold and Italic
	result = SanitizeWikiText("abcd '''''cat''''' efgh")
	assert.Equal(t, "abcd <i><b>cat</b></i> efgh", result)

	// Bold and Italic with Spaces
	result = SanitizeWikiText("abcd''''' cat '''''efgh")
	assert.Equal(t, "abcd<i><b> cat </b></i>efgh", result)
}
