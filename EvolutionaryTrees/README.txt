/*********************************************
 README (EvolutionaryTrees)
*********************************************/

PROGRAMMING LANGUAGE: ---------------------------------------------------------

GOLANG. (v 1.11)

DESCRIPTION: ------------------------------------------------------------------

This folder contains starter code, test cases, and sample datasets for
Ultrametric Tree Generation (UPGMA) and Progressive Multiple Alignment (Clustal).
It also contains useful helper functions for reading and writing data
to files, producing frequency maps, creating pairwise distance matrices,
and dictionary operations. Details are below.

!(IMPORTANT)!
DO NOT tinker with the file structure of this package. Dependencies may fail and
we do not guarantee any compilation/function behavior post modification.

PACKAGES: ---------------------------------------------------------------------

Clustal:
Contains the code for progressive multiple alignment.

  - functions.go contains the shell function which the user should call
  - progressive_align.go contains the core logic
  - updated_align.go contains the aligner, suited to align alignments themselves

Some of this code is still IN PROGRESS (needs to be cleaned up for user use),
so use at your own discretion.

Datatypes: !(IMPORTANT)!
Contains all relevant datatypes required for building and storing tree
information. Most folders MUST have this package imported, and this has
already been done for you.

To access a datatype in a different package, call: Datatype.x, where x is the
datatype you want to use

Datasets: !(IMPORTANT)!
Contains RAW and processed datasets for the SARS-Cov-2 virus and 16S genome region
of river water bacteria from Western Pennsylvania. Input datasets are in the
"Input" directory, and output datasets are/will be placed in the "Output" 
directory.

We recommend ONLY working with processed data (because processing functions
take hours to run). But if you would like to peek at the raw data, you can
find this in Input/x/RAW.

DO NOT tinker with the file-paths here. This will break functions that search
for input files and write to output files.

Io:
Contains all completed code for input-ouput, file-reading and writing, and data
processing.
  - process.go contains functions for file processing and writing
  - functions.go contains functions for file reading and string representations

An interface for all relevant functions can be found in
functions_h.go; we recommend you only use functions provided in the interface
according to their specification. We do not guarantee the behavior of other
functions in Io and misused interface functions.

Lib:
Contains extraneous helper functions used in Clustal and data processing, NOT
UPGMA! Most of this code has been imported from other modules.

  - dict.go contains various dictionary operations
  - distance.go creates distance matrices from string slices
    - Uses Levenshtein Distance (match = 1, sigma = 5, mu = 1)
  - freq.go creates frequency maps from string slices

We mention again: this code is not necessary for building Ultrametric Trees.

UPGMA: !(IMPORTANT)!
Contains starter code and a file structure for building ultrametric trees
between strings. The work is spread amongst multiple files.

  - cluster_operations.go contains all function(s) that manipulate clusters.
    (see definition of clusters in upgma.go)

  - count_leaves.go contains function(s) that count the number of
    leaves at the base of a given parent node.

  - init.go contains functions that initialize trees, clusters, and matrices.
    These functions have already been implemented for you; use the interface
    provided in init_h.go for directions on how to use these functions.

  - matrix_operations.go contains all functions that manipulate distance matrices.

  - upgma.go contains the driver engine for constructing evolutionary trees. Will
    utilize almost all the functions provided in the files above.

A testing suite is also provided in this folder. The tests themselves are located
in folder Tests, and the tester functions are located in functions_test.go. To run
a test, type the command (when inside UPGMA, without the "%"):

  % go test -run x

where x is the name of the test you would like to run. The tests that are provided
are:

  TestFindMinElement
  TestAddRowCol
  TestDelRowCol
  TestCountLeaves
  TestUPGMA

FILES: ------------------------------------------------------------------------

driver.go: !(IMPORTANT)!
In the case your testing suite malfunctions or you would like to manually test
your own code, we provide a main() method in this file.

All relevant packages already have been imported; simply comment out packages
in the header before using them. To run this file, type the following
into the command line (when inside EvolutionaryTrees, without the "%"):

  % go run driver.go

create.go: !(IMPORTANT)!
Runs your functions against the datasets provided in Datasets/Input...
Places the output of these files in the corresponding Datasets/Output... folders.

To run this file, type the following commands (when inside EvolutionaryTrees,
without the "%"):

  % go build create.go
  % ./create

ANNOTATIONS/DATASETS: ---------------------------------------------------------
!(IMPORTANT)!

You can create your own datasets to test your functions against; these datasets
can contain either annotated or unannotated strings. An unannotated file, should
have the following format:

   string1
   string2
   string3
   ...
   stringN

One string per line. The processing functions assign dummy labels each unannotated
strings, according to the order they appeared in.

An annotated file should have the following format:

  annotation1
  string1
  annotation2
  string2
  ...
  annotationN
  stringN

Each annotation string should have the following format:

  category|description

NO TWO STRINGS MAY HAVE THE SAME ANNOTATION. This may cause dictionary errors
and will not produce the intended result. Strings may have the same category, but
must have unique descriptions.

To test these datasets against your code, run the following set of commands in the
driver.go file (ensure you've uncommented each respective package):

  UNANNOTATED:
  strs   := Io.ReadStringsFromFile(filePath)
  dnaMap := Lib.CreateFrequencyDNAMap(strs)
  S, DS  := Lib.GetKeyValues(dnaMap)
  M      := Lib.CalcDistanceMatrix(DS)
  T      := UPGMA.UPGMA(M, S)

  ANNOTATED:
  dnaMap := Io.ReadDNAStringsFromFile(filePath)
  S, DS  := Lib.GetKeyValues(dnaMap)
  M      := Lib.CalcDistanceMatrix(DS)
  T      := UPGMA.UPGMA(M, S)
