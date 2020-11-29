package go_koans

import "io/ioutil"
import "strings"

func aboutFiles() { //5
	filename := "about_files.go"
	contents, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(contents), "\n")
	assert(lines[5] == "func aboutFiles() { //5") // handling files is too trivial
}
