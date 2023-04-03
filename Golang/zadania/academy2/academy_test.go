package academy

import (
	"testing"
)

type mock struct {
	students map[string]Student
}

func (m *mock) List(year uint8) ([]string, error) {
	var names []string
	for name, student := range m.students {
		if student.Year() == year {
			names = append(names, name)
		}
	}
	return names, nil
}

func (m *mock) Get(name string) (Student, error) {
	s, err := m.students[name]
	if !err {
		return Sophomore{}, ErrStudentNotFound
	}
	return s, nil
}

func (m *mock) Save(name string, _ uint8) error {
	s := m.students[name]
	m.students[name] = s
	return nil
}

func (m *mock) Graduate(name string) error {
	delete(m.students, name)
	return nil
}

func TestGradeYear(t *testing.T) {
	r := &mock{
		students: map[string]Student{
			"John": Sophomore{name: "John"},
			"Jane": Sophomore{name: "Jane"},
			"Bob":  Sophomore{name: "Bob"},
		},
	}

	err := GradeYear(r, 1)
	if err != nil {
		t.Fatal(err)
	}

	if len(r.students) != 3 {
		t.Fatalf("expected 3 students, got %d", len(r.students))
	}

	if r.students["John"].Year() != 2 {
		t.Errorf("expected John to be in year 2, got %d", r.students["John"].Year())
	}
}

func TestGradeStudent(t *testing.T) {
	r := &mock{
		students: map[string]Student{
			"John": Sophomore{name: "John"},
			"Jane": Sophomore{name: "Jane"},
			"Bob":  Sophomore{name: "Bob"},
		},
	}

	err := GradeStudent(r, "John")
	if err != nil {
		t.Fatal(err)
	}

	if r.students["John"].Year() != 2 {
		t.Errorf("expected John to be in year 2, got %d", r.students["John"].Year())
	}
}
