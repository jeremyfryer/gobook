package math

import "testing"

type testpair struct {
	values   []float64
	expected float64
}

var testsAverage = []testpair{
	{ []float64{1,2}, 1.5 },
	{ []float64{1,1,1,1,1,1}, 1 },
	{ []float64{-1,1}, 0 },
	{ []float64{}, 0 },
}

func TestAverage(t *testing.T) {
	for _, pair := range testsAverage {
		v := Average(pair.values)
		if v != pair.expected {
			t.Error("For", pair.values,
				"expected", pair.expected,
				"got", v)
		}
	}
}

var testsMin = []testpair{
	{ []float64{1,2,3,4}, 1 },
	{ []float64{99,88,77,66}, 66 },
	{ []float64{-99,8,-77,6}, -99 },
}

func TestMin(t *testing.T) {
	for _, pair := range testsMin {
		v := Min(pair.values)
		if v != pair.expected {
			t.Error("For", pair.values,
				"expected", pair.expected,
				"got", v)
		}
	}
}

var testsMax = []testpair{
	{ []float64{1,2,3,4}, 4 },
	{ []float64{99,88,77,66}, 99 },
	{ []float64{-5,-88,-7,-66}, -5 },
}

func TestMax(t *testing.T) {
	for _, pair := range testsMax {
		v := Max(pair.values)
		if v != pair.expected {
			t.Error("For", pair.values,
				"expected", pair.expected,
				"got", v)
		}
	}
}