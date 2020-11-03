package main

//GetTrimmedNeighbors takes in a string pattern (read), an adjacency list of an overlap network and an integer maxK.
//It returns the set of neighbors of pattern in the updated adjacency list after trimming all edges x -> y
//such that we can reach y from x via a path of length between 2 and maxK.
func GetTrimmedNeighbors(pattern string, adjList map[string][]string, maxK int) []string {
	neighbors := adjList[pattern] // current neighbors
	//need extended neighborhood too

	extendedNeighbors := GetExtendedNeighbors(pattern, adjList, maxK)

	/*
		//we also could have created a new list of things that don't appear in extended neighborhood
		newNbd := make([]string, 0)
		for i := range neighbors {
			//do I see current neighbor in extended neighborhood? If NO, add it to new neighborhood
			if Contains(extendedNeighbors, neighbors[i]) == false {
				newNbd = append(newNbd, neighbors[i])
			}
		}
	*/

	//range over neighbors and throw out anything that occurs in extended neighborhood
	for i := len(neighbors) - 1; i >= 0; i-- {
		// do we see current neighbor in extended neighborhood?
		if Contains(extendedNeighbors, neighbors[i]) {
			neighbors = append(neighbors[:i], neighbors[i+1:]...)
		}
	}
	return neighbors
}
