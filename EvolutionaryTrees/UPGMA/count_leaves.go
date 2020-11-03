package UPGMA

import (
	"EvolutionaryTrees/Datatypes"
)

/* CountLeaves : *Node -> int
   REQUIRES    : v is non-nil.
	 ENSURES     : \result == number of leaves in the subtree of v
   DESCRIP     : CountLeaves is a recursive Node function that counts
	               the number of leaves in the tree rooted at the node.
								 It returns 1 at a leaf.
*/
func CountLeaves(v *Datatypes.Node) int {
	if v.Child1 == nil || v.Child2 == nil {
		return 1
	}
	//default: return sum of leaves of each child.
	return CountLeaves(v.Child1) + CountLeaves(v.Child2)
}
