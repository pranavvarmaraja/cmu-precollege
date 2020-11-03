package UPGMA

import (
	"EvolutionaryTrees/Datatypes"
)

/*
  DelClusters : []*Node * int * int -> []*Node
  REQUIRES    : row < len(clusters) ; col < len(clusters), row < col
  ENSURES     : len(\result) == len(clusters) - 1
  DESCRIP     : Given a slice of Node pointers along with a row/col index,
                deletes the clusters in the slice corresponding to these
                indices.

*/
func DelClusters(clusters []*Datatypes.Node, row, col int) []*Datatypes.Node {
	clusters = append(clusters[:col], clusters[col+1:]...)
	clusters = append(clusters[:row], clusters[row+1:]...)
	return clusters
}
