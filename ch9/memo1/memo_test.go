package memo_test

import (
	"gopl.io/ch9/memo1"
	"gopl.io/ch9/memotest"
	"testing"
)

//  go test -run=TestConcurrent -race -v gopl.io/ch9/memo1
var httpGetBody = memotest.HttpGetBody

func Test(t *testing.T)  {
	m := memo.New(httpGetBody)
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T)  {
	m := memo.New(httpGetBody)
	memotest.Concurrent(t, m)
}