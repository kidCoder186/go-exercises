package phonenumber

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

// LocalCode structure
type LocalCode struct {
	exCode  string
	subCode string
}

// PhoneNumber structure
type PhoneNumber struct {
	intCode  byte
	areaCode string
	LocalCode
}

func (p PhoneNumber) String() string {
	return fmt.Sprintf("%s%s%s", p.areaCode, p.exCode, p.subCode)
}

func (p PhoneNumber) Format() string {
	return fmt.Sprintf("(%s) %s-%s", p.areaCode, p.exCode, p.subCode)
}

func newPhoneNumber(s string) (PhoneNumber, error) {
	fields := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsDigit(r)
	})
	num := strings.Join(fields, "")
	p := PhoneNumber{intCode: '1'}
	if len(num) == 11 {
		p.intCode = num[0]
		p.areaCode = num[1:4]
		p.LocalCode = LocalCode{num[4:7], num[7:11]}
	} else if len(num) == 10 {
		p.areaCode = num[:3]
		p.LocalCode = LocalCode{num[3:6], num[6:10]}
	} else {
		return p, fmt.Errorf("invalid input")
	}
	if p.intCode != '1' {
		return p, errors.New("invalid international code")
	}
	if p.areaCode[0] <= '1' {
		return p, errors.New("invalid area code")
	}
	if p.exCode[0] <= '1' {
		return p, errors.New("invalid exchange code")
	}
	return p, nil
}

// Number returns phone number
func Number(s string) (string, error) {
	p, err := newPhoneNumber(s)
	return fmt.Sprintf("%s", p), err
}

// AreaCode returns area code of a phone number
func AreaCode(s string) (string, error) {
	p, err := newPhoneNumber(s)
	return p.areaCode, err
}

// Format returns a formatted phone number
func Format(s string) (string, error) {
	p, err := newPhoneNumber(s)
	return p.Format(), err
}
