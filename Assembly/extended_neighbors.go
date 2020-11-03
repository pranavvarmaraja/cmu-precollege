package main

//GetExtendedNeighbors takes in a pattern (read), the overlap graph and maxK.
//It returns the extendedNeighbors list. For each neighbor *n* in this list,
//distance between n and pattern is between 2 to maxK.
func GetExtendedNeighbors(pattern string, adjList map[string][]string, maxK int) []string {
	extendedNeighbors := []string{}
	currentNeighbors := adjList[pattern] // where you can reach in one step

	// let's take one step at a time, and use a subroutine to do so
	for j := 2; j <= maxK; j++ {
		currentNeighbors = AdjacentStrings(currentNeighbors, adjList) // update current neighbors to make one step
		currentNeighbors = RemovePattern(currentNeighbors, pattern)   // throw out original pattern if you see it
		extendedNeighbors = append(extendedNeighbors, currentNeighbors...)
	}

	return extendedNeighbors
}

//RemovePattern searches through a list of strings for a given string and removes
//it where it occurs.
func RemovePattern(patterns []string, text string) []string {
	for i := len(patterns) - 1; i >= 0; i-- {
		if patterns[i] == text {
			// remove this pattern
			patterns = append(patterns[:i], patterns[i+1:]...)
		}
	}
	return patterns
}

//AdjacentStrings takes a list of strings and finds all neighbors of these strings
//using an underyling adjacency list.
func AdjacentStrings(currentNeighbors []string, adjList map[string]([]string)) []string {
	reachableNeighbors := make([]string, 0)

	//range over current neighbors and add all neighbors to the list we are building.
	for i := 0; i < len(currentNeighbors); i++ {
		reachableNeighbors = append(reachableNeighbors, adjList[currentNeighbors[i]]...)
	}

	return reachableNeighbors
}
