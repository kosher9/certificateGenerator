package cert

import (
	"fmt"
	"strings"
	"time"
)

type Cert struct {
	Course string
	Name   string
	Date   time.Time

	LabelTitle         string
	LabelCompletion    string
	LabelPresented     string
	LabelParticipation string
	LabelDate          string
}

type Saver interface {
	Save(c Cert) error
}

var MaxLenghtCourse = 20

func New(course, name, date string) (*Cert, error){
	c, err := validateCourse(course)
	if err != nil {
		return nil, err
	}
	n := name
	d := date
	
	cert := &Cert{
		Course:             c,
		Name:               n,
		LabelTitle:         fmt.Sprintf("%v Certificate - %v", c, n),
		LabelCompletion:    "Certificate of completion",
		LabelPresented:     "This Certificate is Presente To",
		LabelParticipation: fmt.Sprintf("For participation in the %v", c),
		LabelDate:          fmt.Sprintf("Date: %v", d),
	}
	return cert, nil
}

func validateCourse(course string) (string, error) {
	c, err := validateStr(course, MaxLenghtCourse)
	if err != nil {
		return "", err
	}
	if !strings.HasSuffix(c, " course"){
		c += " course"
	}
	return strings.ToUpper(c), nil
}

func validateStr(str string, maxLenght int) (string, error) {
	c := strings.TrimSpace(str)
	if len(c) <= 0 {
		return c, fmt.Errorf("Invalid string. got = '%v', len = %d", c, len(c))
	} else if len(c) >= maxLenght {
		return c, fmt.Errorf("Invalid string. got = '%v', len = %d", c, len(c))
	}
	return str, nil
}