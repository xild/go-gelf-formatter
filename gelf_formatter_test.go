package xild

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/Sirupsen/logrus"
)

// GELF Format Specification
// Version 1.1 (11/2013)

// A GELF message is a GZIP’d or ZLIB’d JSON string with the following fields:

// version string (UTF-8)
// short_message string (UTF-8)
// full_message string (UTF-8)
// timestamp number
// level number
// _[additional field] string (UTF-8) or number
//more information: http://docs.graylog.org/en/2.1/pages/gelf.html
func TestLogFormart(t *testing.T) {
	expectedFields := []string{"version", "short_message", "full_message", "timestamp", "level"}
	formatter := &GELFFormatter{}

	b, err := formatter.Format(logrus.WithField("body", "o body"))
	if err != nil {
		t.Fatal("Unable to format entry: ", err)
	}

	entry := make(map[string]interface{})
	err = json.Unmarshal(b, &entry)
	if err != nil {
		t.Fatal("Unable to unmarshal formatted entry: ", err)
	}

	if ok, value := validExpectedField(expectedFields, entry); !ok {
		t.Fatalf("[GreyLog required field not found %s the values founds are %s]", value, entry)
	}
}

func TestWithLogAddFieldsAsAdditionalFields(t *testing.T) {
	expectedFields := []string{"_body", "_err"}

	formatter := &GELFFormatter{}

	b, err := formatter.Format(logrus.WithFields(logrus.Fields{
		"body": "value of body",
		"err":  "value of err",
	}))
	if err != nil {
		t.Fatal("Unable to format entry: ", err)
	}

	entry := make(map[string]interface{})
	err = json.Unmarshal(b, &entry)
	if err != nil {
		t.Fatal("Unable to unmarshal formatted entry: ", err)
	}

	if ok, value := validExpectedField(expectedFields, entry); !ok {
		t.Fatalf("[Additional field not found %s the values founds are %s]", value, entry)
	}

}

func TestWithLogSetFacilityWithEnvSetted(t *testing.T) {
	expectedFields := []string{"_body", "_err", "_appName"}
	os.Setenv("APPLICATION_NAME", "TEST-XILD")

	formatter := &GELFFormatter{}

	b, err := formatter.Format(logrus.WithFields(logrus.Fields{
		"body": "value of body",
		"err":  "value of err",
	}))
	if err != nil {
		t.Fatal("Unable to format entry: ", err)
	}

	entry := make(map[string]interface{})
	err = json.Unmarshal(b, &entry)
	if err != nil {
		t.Fatal("Unable to unmarshal formatted entry: ", err)
	}

	if ok, value := validExpectedField(expectedFields, entry); !ok {
		t.Fatalf("[Additional field not found %s the values founds are %s]", value, entry)
	}

}

func validExpectedField(expectedFields []string, entry map[string]interface{}) (bool, string) {
	for _, value := range expectedFields {
		if _, ok := entry[value]; !ok {
			return false, value
		}
	}
	return true, ""
}
