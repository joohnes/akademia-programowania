package academy

import (
	"math"
)

type Student struct {
	Name       string
	Grades     []int
	Project    int
	Attendance []bool
}

// AverageGrade returns an average grade given a
// slice containing all grades received during a
// semester, rounded to the nearest integer.
func AverageGrade(grades []int) int {
	var sum float64
	for grade := range grades {
		sum = sum + float64(grades[grade])
	}
	if len(grades) == 0 {
		return 0
	}
	return int(math.Round(sum / float64(len(grades))))
}

// AttendancePercentage returns a percentage of class
// attendance, given a slice containing information
// whether a student was present (true) of absent (false).
//
// The percentage of attendance is represented as a
// floating-point number ranging from 0 to 1.
func AttendancePercentage(attendance []bool) float64 {
	var countTruths int

	for _, i := range attendance {
		if i {
			countTruths += 1
		}
	}
	if countTruths == 0 {
		return 0
	}
	return float64(countTruths) / float64(len(attendance))
}

// FinalGrade returns a final grade achieved by a student,
// ranging from 1 to 5.
//
// The final grade is calculated as the average of a project grade
// and an average grade from the semester, with adjustments based
// on the student's attendance. The final grade is rounded
// to the nearest integer.

// If the student's attendance is below 80%, the final grade is
// decreased by 1. If the student's attendance is below 60%, average
// grade is 1 or project grade is 1, the final grade is 1.

func FinalGrade(s Student) int {
	avgGrade := AverageGrade(s.Grades)
	attd := AttendancePercentage(s.Attendace)

	if attd < 0.6 || (avgGrade == 1 || s.Project == 1) {
		return 1
	}
	var projectGrade = float64(avgGrade+s.Project) / 2
	if attd < 0.8 {
		projectGrade -= 1
	}

	return int(math.Round(projectGrade))
}

// GradeStudents returns a map of final grades for a given slice of
// Student structs. The key is a student's name and the value is a
// final grade.
func GradeStudents(students []Student) map[string]uint8 {
	gradedStudents := map[string]uint8{}
	if len(students) == 0 {
		return gradedStudents
	}
	for _, student := range students {
		gradedStudents[student.Name] = uint8(FinalGrade(student))
	}
	return gradedStudents
}
