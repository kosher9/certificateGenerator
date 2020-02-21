package cert

import "testing"

func TestValidCertData(t *testing.T) {
	c, err := New("Golang", "Bob", "2014-07-27")
	if err!=nil{
		t.Errorf("Cert data should be valid. err=%v", err)
	}
	if c == nil {
		t.Errorf("Cert should be a valid reference")
	}

	if c.Course != "GOLANG COURSE"{
		t.Errorf("Course name is not valid. Expected = 'GOLANG COURSE'. Got = %v", c.Course)
	}
}

func TestCourseEmptyValue(t *testing.T) {
	_, err := New("", "Bob", "2014-07-27")
	if err == nil {
		t.Errorf("Course is empty. err = %v", err)
	}
}

func TestCourseTooLong(t *testing.T) {
	course := "dvjkenrvienvoernvowvneokvnepkvnepkveovkeovweonovwvow"
	_, err := New(course, "Bob", "2014-07-27")
	if err == nil {
		t.Errorf("Error should be returned on a too long course name (course = %s)", course)
	}
}

