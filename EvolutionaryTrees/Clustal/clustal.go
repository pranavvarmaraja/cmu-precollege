package Clustal

import (
	// "fmt"
	"EvolutionaryTrees/Datatypes"
)

/*
 MultiAlign : Tree * []string -> Alignment
 REQUIRES : non-nil. strings in {A, T, C, G}*.
            len(guideTree) > len(dnaStrings).
 ENSURES  : estimated multiple sequence alignment.
 DESCRIP  : Given DNA strings and their phylogenetic tree structure,
            performs a progressive multiple sequence alignment between
            these strings.
*/
func MultiAlign(guideTree Datatypes.Tree, dnaStrings []string, match, mismatch, gap, supergap float64) Datatypes.Alignment {

	for i, gene := range dnaStrings {
		//label the i-th leaf node with the i-th string given
		guideTree[i].Algnmt = Datatypes.Alignment{gene}
	}

	//range through the internal nodes and use the progressive alignment to assign an alignment
	//to each internal node
	for p := len(dnaStrings); p < len(guideTree); p++ {
		//align the two children of guideTree[p]
		align1 := guideTree[p].Child1.Algnmt
		align2 := guideTree[p].Child2.Algnmt
		guideTree[p].Algnmt = ProgressiveAlign(align1, align2, match, mismatch, gap, supergap)
	}

	//where is the multiple alignment hiding? At the root!
	//root's label in the tree is ... the max index
	return guideTree[len(guideTree)-1].Algnmt
}
