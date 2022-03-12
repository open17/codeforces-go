// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_b(t *testing.T) {
	examples := [][]string{
		{
			`[0,3,1,0,4,5,2,0]`, 
			`[4,11]`,
		},
		{
			`[0,1,0,3,0,2,2,0]`, 
			`[1,3,4]`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, mergeNodes, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-281/problems/merge-nodes-in-between-zeros/