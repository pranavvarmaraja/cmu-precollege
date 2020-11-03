package main

//StringSliceEquals takes two slice strings and returns true if they have
//the same elements and false otherwise.
func StringSliceEquals(patterns1, patterns2 []string) bool {
	return StrSliceEq(patterns1, patterns2) // forgot this was in helper code :)
}
