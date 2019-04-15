package krutf

import (
	"bytes"
	"io/ioutil"

	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
)

type code int

// UTF8 - 0, EUCKR - 1
const (
	UTF8  code = 0
	EUCKR code = 1
)

// ConvertByte convert encoding of byte slices.
// Takes 2 parameters [byte slices to convert] and [expecting encoding (UTF8 or EUCKR)].
// Return converted byte slices and error.
func ConvertByte(b []byte, c code) ([]byte, error) {
	if c == 0 {
		// To UTF8
		ret, err := byteToUTF(b)
		if err == nil {
			return ret, nil
		}
		return []byte{}, err
	} else if c == 1 {
		// To EUCKR
		return byteToKR(b), nil
	} else {
		// Unknown Code
		panic("Please use either UTF8 or EUCKR")
	}
}

// ConvertString convert encoding of string.
// Takes 2 parameters [string to convert] and [expecting encoding (UTF8 or EUCKR)].
// Return converted string and error.
func ConvertString(s string, c code) (string, error) {

	if c == 0 {
		// To UTF8
		ret, err := stringToUTF(s)
		if err == nil {
			return ret, nil
		}
		return "", err
	} else if c == 1 {
		// To EUCKR
		return stringToKR(s), nil
	} else {
		panic("Please use either UTF8 or EUCKR")
	}
}

func stringToUTF(s string) (string, error) {
	newByte, err := byteToUTF([]byte(s))
	return string(newByte), err
}

func byteToUTF(b []byte) ([]byte, error) {
	eucReader := bytes.NewReader(b)
	utf8Decoder := korean.EUCKR.NewDecoder()
	utf8Reader := transform.NewReader(eucReader, utf8Decoder)
	d, err := ioutil.ReadAll(utf8Reader)
	return d, err
}

func stringToKR(s string) string {
	newByte := byteToKR([]byte(s))
	return string(newByte)
}

func byteToKR(b []byte) []byte {
	var buffer bytes.Buffer
	krEncoder := korean.EUCKR.NewEncoder()
	utf8Writer := transform.NewWriter(&buffer, krEncoder)
	utf8Writer.Write(b)
	defer utf8Writer.Close()
	return buffer.Bytes()
}
