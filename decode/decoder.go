package decode

import (
	"encoding/base64"
	"io/ioutil"
	"strings"

	"github.com/mcnijman/go-emailaddress"
	"github.com/realpamisa/gmail-extractor/pkg"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

var AppendEmail []string
var C = 1

func HandleEmailDecode(data string) {
	// --- Decoding: convert encS from ShiftJIS to UTF8
	// declare a decoder which reads from the string we have just encoded
	rInUTF8 := transform.NewReader(strings.NewReader(data), japanese.ShiftJIS.NewDecoder())
	// decode our string
	decBytes, _ := ioutil.ReadAll(rInUTF8)

	var decodedByte, _ = base64.StdEncoding.DecodeString(string(decBytes))
	var decodedString = string(decodedByte)
	text := []byte(decodedString)
	validateHost := false

	emails := emailaddress.Find(text, validateHost)
	for _, val := range emails {
		checkIfValid := ContainsString(val.String())
		if checkIfValid != true {
			AppendEmail = append(AppendEmail, val.String())
		}
		if len(AppendEmail) == 1000 {
			pkg.CsvSave(AppendEmail, C)
			C++
		}

	}

}

func ContainsString(val string) bool { //check email already exist
	for _, a := range AppendEmail {
		if a == val {
			return true
		}
	}
	return false

}
