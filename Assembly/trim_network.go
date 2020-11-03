package main

//TrimNetwork takes in an overlap network adjList and a max iteration maxK
//It removes all n-transitivity edges in adjList, n <= maxK.
//It returns a trimmed copy of the original network.
func TrimNetwork(adjList map[string][]string, maxK int) map[string][]string {
	adjList2 := make(map[string][]string)

	//fill in your code here :)
	//range over old adjacency list and update the neighbors of each key
	for pattern := range adjList {
		adjList2[pattern] = GetTrimmedNeighbors(pattern, adjList, maxK)
	}

	return adjList2
}
