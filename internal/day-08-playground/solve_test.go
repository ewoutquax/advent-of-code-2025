package day08playground_test

import (
	"fmt"
	"testing"

	. "github.com/ewoutquax/advent-of-code-2025/internal/day-08-playground"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	boxes := ParseInput(testInput())

	assert := assert.New(t)
	assert.Equal("[]day08playground.JunctionBox", fmt.Sprintf("%T", boxes))
	assert.Len(boxes, 20)

	assert.Equal(162, boxes[0].X)
	assert.Equal(817, boxes[0].Y)
	assert.Equal(812, boxes[0].Z)
	assert.Equal(57, boxes[1].X)
	assert.Equal(689, boxes[len(boxes)-1].Z)
}

func TestDistances(t *testing.T) {
	boxes := ParseInput(testInput())

	distances := DistanceBetweenBoxes(boxes)

	assert := assert.New(t)
	assert.Len(distances.BetweenBoxes, 20*19/2)
	assert.Len(distances.Lengths, 20*19/2)

	shortestLength := distances.Lengths[0]
	assert.Equal("162,817,812", distances.BetweenBoxes[shortestLength].From.ToS())
	assert.Equal("425,690,689", distances.BetweenBoxes[shortestLength].To.ToS())

	shortestLength2 := distances.Lengths[1]
	assert.Equal("162,817,812", distances.BetweenBoxes[shortestLength2].From.ToS())
	assert.Equal("431,825,988", distances.BetweenBoxes[shortestLength2].To.ToS())
}

func TestFindCircuits(t *testing.T) {
	boxes := ParseInput(testInput())
	distances := DistanceBetweenBoxes(boxes)

	result := FindCircuitSizes(distances, 10)
	fmt.Printf("result: %v\n", result)

	assert.Equal(t, 40, result)
}

/**
func TestFindMinNeededConnections(t *testing.T) {
	boxes := ParseInput(testInput())

	result := FindMinNeededConnections(boxes)

	assert.Equal(t, 40, result)
}
*/

func testInput() []string {
	return []string{
		"162,817,812",
		"57,618,57",
		"906,360,560",
		"592,479,940",
		"352,342,300",
		"466,668,158",
		"542,29,236",
		"431,825,988",
		"739,650,466",
		"52,470,668",
		"216,146,977",
		"819,987,18",
		"117,168,530",
		"805,96,715",
		"346,949,466",
		"970,615,88",
		"941,993,340",
		"862,61,35",
		"984,92,344",
		"425,690,689",
	}
}
