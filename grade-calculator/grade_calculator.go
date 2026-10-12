package esepunittests

type GradeCalculator struct {
	grades   []Grade
	passFail bool
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator(passFail ...bool) *GradeCalculator {
	gc := &GradeCalculator{
		grades: make([]Grade, 0),
	}

	if len(passFail) > 0 {
		gc.passFail = passFail[0]
	}

	return gc
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()
	if gc.passFail {
		if numericalGrade >= 70 {
			return "Pass"
		}
		return "Fail"
	}

	if numericalGrade >= 90 {
		return "A"
	} else if numericalGrade >= 80 {
		return "B"
	} else if numericalGrade >= 70 {
		return "C"
	} else if numericalGrade >= 60 {
		return "D"
	}

	return "F"
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	switch gradeType {
	case Assignment, Exam, Essay:
		gc.grades = append(gc.grades, Grade{
			Name:  name,
			Grade: grade,
			Type:  gradeType,
		})
	}
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	assignments := make([]Grade, 0)
	exams := make([]Grade, 0)
	essays := make([]Grade, 0)

	for _, grade := range gc.grades {
		switch grade.Type {
		case Assignment:
			assignments = append(assignments, grade)
		case Exam:
			exams = append(exams, grade)
		case Essay:
			essays = append(essays, grade)
		}
	}

	assignmentAverage := computeAverage(assignments)
	examAverage := computeAverage(exams)
	essayAverage := computeAverage(essays)

	weightedGrade := float64(assignmentAverage)*.5 +
		float64(examAverage)*.35 +
		float64(essayAverage)*.15

	return int(weightedGrade)
}

func computeAverage(grades []Grade) int {
	sum := 0

	for _, grade := range grades {
		sum += grade.Grade
	}

	return sum / len(grades)
}
