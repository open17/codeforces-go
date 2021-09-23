// Code generated by generator_test.
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [a]")
	exampleIns := [][]string{{`2`}, {`11`}, {`10000`}, {`69`}, {`1010`}}
	exampleOuts := [][]string{{`[1,1]`}, {`[2,9]`}, {`[1,9999]`}, {`[1,68]`}, {`[11,999]`}}
	// custom test cases or WA cases.
	//exampleIns = append(exampleIns, []string{``})
	//exampleOuts = append(exampleOuts, []string{``})
	if err := testutil.RunLeetCodeFunc(t, getNoZeroIntegers, exampleIns, exampleOuts); err != nil {
		t.Fatal(err)
	}
}