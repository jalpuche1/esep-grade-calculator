# ESEP Grade Calculator

A Go library that calculates a final letter grade from assignments,
exams, and essays.

## Requirements

- Go 1.24.4 or newer
- Git to clone the repository

## Setup

Clone the repository and enter the Go module directory:

```powershell
git clone https://github.com/jalpuche1/esep-grade-calculator.git
cd esep-grade-calculator
cd grade-calculator
```

## Using the library

Create a calculator with NewGradeCalculator(), add scores using
AddGrade(name, score, gradeType), and obtain the letter grade with
GetFinalGrade().

Grade types are Assignment, Exam, and Essay. The intended weights are
50% for assignments, 35% for exams, and 15% for essays.

## Run tests

From the grade-calculator directory:

```powershell
go test -v ./...
```

The original test suite contains A, B, and F cases. At the initial
repository state, the A and B tests fail, while the F test passes
despite using high scores. Later assignment stages correct these issues.

## Measure coverage

```powershell
go test -coverprofile=coverage.txt ./...
go tool cover -func=coverage.txt
go tool cover -html=coverage.txt
```

The HTML command opens a view showing which statements the tests execute.

## Repository files

- README.MD: setup, usage, and testing instructions.
- grade-calculator/go.mod: module name and required Go version.
- grade-calculator/grade_calculator.go: grade storage and calculation logic.
- grade-calculator/grade_calculator_test.go: automated unit tests.