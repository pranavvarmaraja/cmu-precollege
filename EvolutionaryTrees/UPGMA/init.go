package UPGMA

import (
	"EvolutionaryTrees/Datatypes"
	"strconv"
)

/* InitializeMatrix : int * int -> DistanceMatrix
   REQUIRES         : m >= 0, n >= 0
	 ENSURES          : len(\result) == 0 ; len(\result[0]) == 0 ; \result[i][j] == 0.0
   DESCRIP          : Creates empty matrix of size m x n.
 */
func InitializeMatrix(m int, n int) Datatypes.DistanceMatrix {
	mtx := make([][]float64, m)

	for i := range mtx {
		mtx[i] = make([]float64, n)
	}
	return mtx
}

/*
  InitializeClusters : Tree -> []*Node
  REQUIRES           : true
  ENSURES            : \result is a subtree of t
  DESCRIP            : A Tree method that returns a slice of pointers to the leaves
                       of the Tree.
*/
func InitializeClusters(t Datatypes.Tree) []*Datatypes.Node {
	numNodes := len(t)
	numLeaves := (numNodes + 1) / 2

	clusters := make([]*Datatypes.Node, numLeaves)
	// clusters[i] should point to the i-th leaf node of t
	for i := range clusters {
		clusters[i] = t[i]
	}

	return clusters
}

/*
  InitializeTree : []string -> Tree
  REQUIRES       : true
  ENSURES        : \result rooted binary tree with leaves populated.
  DESCRIP        : Takes the n names of our present-day species (leaves) and
                   returns a rooted binary tree with 2n-1 total nodes, where
                   the leaves are the first n and have the associated species names.
*/
func InitializeTree(speciesNames []string) Datatypes.Tree {
	numLeaves := len(speciesNames)
	var t Datatypes.Tree // a Tree is []*Node

	t = make([]*Datatypes.Node, 2*numLeaves-1)
	// all of these pointers have default value of nil; we need to point them at nodes

	// we should create 2n-1 nodes.
	for i := range t {
		var vx Datatypes.Node
		// let's label the first numLeaves nodes with the appropriate species name.
		// by default, vx.age = 0.0, and its children are nil.
		if i < numLeaves {
			//at a leaf ... let's assign its label.
			vx.Label = speciesNames[i]
		} else {
			// let's just give it an unspecific name
			vx.Label = "Ancestor species " + strconv.Itoa(i)
		}
		// indexing the node
		vx.Num = i
		// one thing to do: point t[i] at vx
		t[i] = &vx
	}

	return t
}
