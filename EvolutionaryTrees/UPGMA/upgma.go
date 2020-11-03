package UPGMA

import (
	"EvolutionaryTrees/Datatypes"
)

/*
  UPGMA    : DistanceMatrix * []string -> Tree
  REQUIRES : len(mtx) == len(mtx[len(mtx) - 1]) ; mtx[i][j] == m[j][i]
             len(mtx) == len(speciesNames)
  ENSURES  : \result is valid tree
  DESCRIP  : Given species names and pairwise distances between species,
             constructs an ultrametic evolutionary tree. Tree is represented
             as an array of nodes. (see main.go for declaration)
*/
func UPGMA(mtx Datatypes.DistanceMatrix, speciesNames []string) Datatypes.Tree {
	t := InitializeTree(speciesNames)

	clusters := InitializeClusters(t)

	numLeaves := len(mtx)

	//range over all internal nodes and create each one by one step of UPGMA
	for p := numLeaves; p < 2*numLeaves-1; p++ {
		row, col, minVal := FindMinElement(mtx)

		//I know the minVal, so I know the age of t[p]
		t[p].Age = minVal / 2.0

		//I also know the children, they are clusters[row] and clusters[col]
		t[p].Child1 = clusters[row]
		t[p].Child2 = clusters[col]

		//node has been added to tree. next, update matrix.

		mtx = AddRowCol(mtx, clusters, row, col)

		mtx = DelRowCol(mtx, row, col)

		//next, update clusters.

		clusters = append(clusters, t[p]) // point final element of clusters at new node

		//hack out clusters[row] and clusters[col]

		clusters = DelClusters(clusters, row, col)

	}

	// the tree has been generated!

	return t
}

/*
func AssertSameSize(mtx Datatypes.DistanceMatrix, speciesNames []string) {
	if len(mtx) != len(speciesNames) {
		panic("Error: matrix has different number of species than species given.")
	}
}
*/
