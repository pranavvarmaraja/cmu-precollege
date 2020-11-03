package Io

import "EvolutionaryTrees/Datatypes"

type IO interface {

	/* ReadMatrixFromFile
	   REQUIRES : Valid file path. File provided is a tab-separated matrix.
	              First line is number of samples. Left column are row names.
	              All numbers in matrix can be parsed as floats.
	   ENSURES  : Returns file matrix into float matrix and list of row names.
	*/
	ReadMatrixFromFile(string) (Datatypes.DistanceMatrix, []string)

	/* ReadStringsFromFile
	   REQUIRES : Valid file path. File of strings, one string per line.
		 ENSURES  : String array corresponding to strings in file.
	*/
	ReadStringsFromFile(string) ([]string)

	/* ReadDNAStringsFromFile
	   REQUIRES : Valid file path. File provided annotated as outlined in
		 						README. (each label MUST BE UNIQUE...see ensures below)
		 ENSURES  : String dictionary, where keys are annotation labels and values
		            are right column.
	*/
	ReadDNAStringsFromFile(string) (map[string]string)


	/* PrintGraphViz
	   REQUIRES : Tree is completed. (a tree is considered "completed" if
	                                      all fields are populated)
     ENSURES  : Prints a visualization of the given tree.
	*/
  PrintGraphViz(Datatypes.Tree) // void

  /* ToNewick
     REQUIRES : Tree is completed. (a tree is considered "completed" if
                                    all fields are populated)
     ENSURES  : Returns formatted string corresponding to tree.
                (more specifically Newick format, which is a popular medium
                 used for data visualization software)
  */
  ToNewick(Datatypes.Tree) (string)

	/* CreateCSV
     REQUIRES : Tree is completed. (a tree is considered "completed" if
                                    all fields are populated)
								Labels MUST be FASTA annotated.
								List of categories for annotation table. For SARS-Cov-2,
								use ["Wuhan","Italy","USA"].
     ENSURES  : \result is string annotation table.
		 DESCRIP  : Creates CSV annotation table for Newick tree. Use ONLY for
		            data visualization.
  */
  CreateCSV(Datatypes.Tree, []string) (string)

	/* CreateDistanceMatrix
     REQUIRES : File name is valid and setting in {W, F}
	 	              W: File is FASTA format or annotated strings.
	 							  F: File is strings, one string per line.
     ENSURES  : \result is a valid distance matrix, slice of labels.
		 DESCRIP  : Given raw DNA strings, annotated or unannotated, produces a
		            symmetric pairwise distance matrix. Returns a slice of annotations
								where labels[i] is the label for matrix[i]. If unannotated, produces
								dummy labels for each string.
  */
	CreateDistanceMatrix(string, string) (Datatypes.DistanceMatrix, []string)

}
