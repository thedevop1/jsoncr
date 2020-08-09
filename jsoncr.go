// Package jsoncr provides method to remove comments and indentations in JSON.
package jsoncr

import (
	"bytes"
	"errors"
	"io"
	"text/scanner"
)

// Remove comments and indentations from r.
func Remove(r io.Reader) ([]byte, error) {
	var err error
	s := scanner.Scanner{Error: func(s *scanner.Scanner, msg string) {
		err = errors.New(msg)
	}}
	s.Init(r)

	var b bytes.Buffer
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		if err != nil {
			return nil, err
		}
		b.WriteString(s.TokenText())
	}
	return b.Bytes(), nil
}
