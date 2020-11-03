package UPGMA

import (
	"EvolutionaryTrees/Datatypes"
)

type Init interface {

  /* InitializeMatrix : int * int -> DistanceMatrix
     REQUIRES         : m >= 0, n >= 0
  	 ENSURES          : len(\result) == m ; len(\result[0]) == n ; \result[i][j] == 0.0
     DESCRIP          : Creates empty matrix of size m x n.
   */
  InitializeMatrix(int, int) Datatypes.DistanceMatrix

  /*
    InitializeClusters : Tree -> []*Node
    REQUIRES           : true
    ENSURES            : \result is a subtree of t
    DESCRIP            : A Tree method that returns a slice of pointers to the leaves
                         of the Tree.
  */
  InitializeClusters(Datatypes.Tree) []*Datatypes.Node

  /*
    InitializeTree : []string -> Tree
    REQUIRES       : true
    ENSURES        : \result rooted binary tree with leaves populated.
    DESCRIP        : Takes the n names of our present-day species (leaves) and
                     returns a rooted binary tree with 2n-1 total nodes, where
                     the leaves are the first n and have the associated species names.
  */
  InitializeTree([]string) Datatypes.Tree

}
