package Datatypes

type DistanceMatrix [][]float64
type Tree []*Node
type Alignment []string

// we also think of a cluster as a *Node
type Node struct {
	Algnmt		     Alignment
	Num						 int
	Age            float64
	Label          string
	Child1, Child2 *Node // if at leaf, both will be nil
}
