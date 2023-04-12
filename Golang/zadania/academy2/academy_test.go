package academy

import (
	"github.com/grupawp/akademia-programowania/Golang/zadania/academy2/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGradeYear(t *testing.T) {
	want := ""
	mock := mocks.NewRepository(t)
	mock.On("List").Return([]string{"a"}, nil)
	got := GradeYear(mock, 2)
	assert.Equal(t, want, got)
}

func TestGradeStudent(t *testing.T) {
	want := ""
	mock := mocks.NewRepository(t)
	mockStudent := mocks.NewStudent(t)
	mock.On("Get").Return(mockStudent, nil)
	got := GradeStudent(mock, "John")
	assert.Equal(t, want, got)
}
