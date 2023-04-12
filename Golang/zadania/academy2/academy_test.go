package academy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGradeYear(t *testing.T) {
	mock := NewMockRepository(t)
	mockStudent := NewMockStudent(t)
	mock.On("List", uint8(2)).Return([]string{"John"}, nil)
	mock.On("Get", "John").Return(mockStudent, nil)
	mockStudent.On("FinalGrade").Return(5)
	mockStudent.On("Year").Return(uint8(2))
	mockStudent.On("Name").Return("John")
	mock.On("Save", "John", uint8(3)).Return(nil)
	got := GradeYear(mock, 2)
	assert.Nil(t, got)
}

func TestGradeStudent(t *testing.T) {
	mock := NewMockRepository(t)
	mockStudent := NewMockStudent(t)
	mock.On("Get", "John").Return(mockStudent, nil)
	mockStudent.On("FinalGrade").Return(5)
	mockStudent.On("Year").Return(uint8(2))
	mockStudent.On("Name").Return("John")
	mock.On("Save", "John", uint8(3)).Return(nil)
	got := GradeStudent(mock, "John")
	assert.Nil(t, got)
}
