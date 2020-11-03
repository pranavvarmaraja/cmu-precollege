package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	electoralVotes := ReadElectoralVotes("electoralVotes.txt")

	fileName := "debates.txt"

	polls := ReadPollingData(fileName)

	numTrials := 1000
	margin := 0.05

	prob1, prob2, prob3 := (SimulateMultipleElections(polls, electoralVotes, numTrials, margin))

	fmt.Print(prob1, prob2, prob3)
}

func SimulateMultipleElections(polls map[string]float64,
	electoralVotes map[string]uint, numTrials int, margin float64) (float64, float64, float64) {

	winCount1 := 0
	winCount2 := 0
	tieCount := 0

	for i := 0; i < numTrials; i++ {
		votes1, votes2 := SimulateOneElection(polls, electoralVotes, margin)

		if votes1 > votes2 {
			winCount1++
		} else if votes2 > votes1 {
			winCount2++
		} else {
			tieCount++
		}
	}

	return float64(winCount1) / float64(numTrials), float64(winCount2) / float64(numTrials), float64(tieCount) / float64(numTrials)

}

func SimulateOneElection(polls map[string]float64, electoralVotes map[string]uint, margin float64) (uint, uint) {

	var collegeVotes1 uint = 0
	var collegeVotes2 uint = 0
	for state := range polls {
		poll := polls[state]
		numVotes := electoralVotes[state]

		adjustedPoll := AddNoise(poll, margin)
		if adjustedPoll >= 0.5 {
			collegeVotes1 += numVotes
		} else {
			collegeVotes2 += numVotes
		}
	}

	return collegeVotes1, collegeVotes2
}

func AddNoise(poll float64, margin float64) float64 {

	return rand.NormFloat64()/2*margin + poll

}
