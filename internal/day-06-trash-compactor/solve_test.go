package day06trashcompactor_test

import (
	"testing"

	. "github.com/ewoutquax/advent-of-code-2025/internal/day-06-trash-compactor"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ParseInputTestSuite struct {
	suite.Suite
	grid Grid
}

type ColumnCalculateTotalTestSuite struct {
	suite.Suite
	grid Grid
}

type GridCalculateTotalTestSuite struct {
	suite.Suite
	grid Grid
}

func TestRunSuiteParseInput(t *testing.T) {
	suite.Run(t, new(ParseInputTestSuite))
}

func (s *ParseInputTestSuite) SetupSuite() {
	s.grid = ParseInput(testInput())
}

func (s *ParseInputTestSuite) TestParseInputSizes() {
	// Sizes
	s.Len(s.grid.Columns, 4)
	s.Len(s.grid.Columns[0].IntRows, 3)
}

func (s *ParseInputTestSuite) TestParseInputNumericValues() {
	s.Equal(123, s.grid.Columns[0].IntRows[0])
	s.Equal(328, s.grid.Columns[1].IntRows[0])
	s.Equal(387, s.grid.Columns[2].IntRows[1])
}

func (s *ParseInputTestSuite) TestParseInputStringValues() {
	s.Equal("123", s.grid.Columns[0].StringRows[0])
	s.Equal(" 45", s.grid.Columns[0].StringRows[1])
	s.Equal("  6", s.grid.Columns[0].StringRows[2])
	s.Equal("98 ", s.grid.Columns[1].StringRows[2])
	s.Equal("23 ", s.grid.Columns[3].StringRows[1])
}

func (s *ParseInputTestSuite) TestParseInputOperator() {
	s.Equal(OperatorMultiply, s.grid.Columns[0].Operator)
	s.Equal(OperatorSum, s.grid.Columns[1].Operator)
}

func TestRunSuiteColumnCalculateTotal(t *testing.T)  { suite.Run(t, new(ColumnCalculateTotalTestSuite)) }
func (s *ColumnCalculateTotalTestSuite) SetupSuite() { s.grid = ParseInput(testInput()) }

func (s *ColumnCalculateTotalTestSuite) TestColumnCalculateTotalTestSuiteInt() {
	testCases := map[int]int{
		0: 33210,
		1: 490,
		2: 4243455,
		3: 401,
	}

	for inputColumnIdx, expectedResult := range testCases {
		actualResult := s.grid.Columns[inputColumnIdx].CalculateTotalInt()
		s.Equal(expectedResult, actualResult)
	}
}

func (s *ColumnCalculateTotalTestSuite) TestColumnCalculateTotalTestSuiteString() {
	expectedResults := []int{
		1058,
		3253600,
		625,
		8544,
	}

	var actualResults []int = make([]int, 4)
	for inputColumnIdx := range 4 {
		actualResults[inputColumnIdx] = s.grid.Columns[inputColumnIdx].CalculateTotalString()
	}

	for _, expectedResult := range expectedResults {
		s.Contains(actualResults, expectedResult)
	}
}

func TestRunSuiteGridCalculateTotal(t *testing.T)  { suite.Run(t, new(GridCalculateTotalTestSuite)) }
func (s *GridCalculateTotalTestSuite) SetupSuite() { s.grid = ParseInput(testInput()) }

func (s *GridCalculateTotalTestSuite) TestGridCalculateTotalTestSuiteInt() {
	total := s.grid.CalculateGrandTotal()

	s.Equal(4277556, total)
}

func (s *GridCalculateTotalTestSuite) TestGridCalculateTotalTestSuiteString() {
	total := s.grid.CalculateGrandTotalString()

	s.Equal(3263827, total)
}

func TestOperatorLocations(t *testing.T) {
	expectedLocations := []int{0, 4, 8, 12}
	lines := testInput()

	locations := OperatorLocations(lines[len(lines)-1])

	assert.Equal(t, expectedLocations, locations)
}

func testInput() []string {
	return []string{
		"123 328  51 64 ",
		" 45 64  387 23 ",
		"  6 98  215 314",
		"*   +   *   +  ",
	}
}
