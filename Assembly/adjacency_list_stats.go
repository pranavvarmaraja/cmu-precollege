package main

//AverageOutDegree takes the adjacency list of an overlap network as input. It returns
//the average number of elements to which a given string is "connected"
//in the network.
func AverageOutDegree(adjList map[string][]string) float64 {
	s := SumOutDegree(adjList)
	return float64(s) / float64(len(adjList))
}

func SumOutDegree(adjList map[string][]string) int {
	c := 0

	for _, val := range adjList {
		c += len(val)
	}

	return c
}
