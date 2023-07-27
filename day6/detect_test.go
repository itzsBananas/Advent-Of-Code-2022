package main

import (
	"testing"
)

type res struct {
	r string
	i int
}

func helper(s string, markerSize int) res {
	rec := NewRecent(markerSize)
	i := 0
	for _, ru := range s {
		i += 1
		if rec.Has(ru) {
			rec = rec.Clear(ru)
		}
		rec = rec.Add(ru)
		if rec.IsFull() {
			break
		}
	}
	result := res{rec.String(), i}
	return result
}

func TestDetect0(t *testing.T) {
	r := helper("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 4)
	expectOutput := res{"jpqm", 7}
	if r.r != expectOutput.r {
		t.Errorf("Expected string output, %s, but got %s", expectOutput.r, r.r)
	}

	if r.i != expectOutput.i {
		t.Errorf("Expected index output, %d, but got %d", expectOutput.i, r.i)
	}
}

func TestDetect1(t *testing.T) {
	r := helper("bvwbjplbgvbhsrlpgdmjqwftvncz", 4)
	expectOutput := res{"vwbj", 5}
	if r.r != expectOutput.r {
		t.Errorf("Expected string output, %s, but got %s", expectOutput.r, r.r)
	}

	if r.i != expectOutput.i {
		t.Errorf("Expected index output, %d, but got %d", expectOutput.i, r.i)
	}
}

func TestDetect2(t *testing.T) {
	r := helper("nppdvjthqldpwncqszvftbrmjlhg", 4)
	expectOutput := res{"pdvj", 6}
	if r.r != expectOutput.r {
		t.Errorf("Expected string output, %s, but got %s", expectOutput.r, r.r)
	}

	if r.i != expectOutput.i {
		t.Errorf("Expected index output, %d, but got %d", expectOutput.i, r.i)
	}
}

func TestDetect3(t *testing.T) {
	r := helper("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 4)
	expectOutput := res{"rfnt", 10}
	if r.r != expectOutput.r {
		t.Errorf("Expected string output, %s, but got %s", expectOutput.r, r.r)
	}

	if r.i != expectOutput.i {
		t.Errorf("Expected index output, %d, but got %d", expectOutput.i, r.i)
	}
}

func TestDetect4(t *testing.T) {
	r := helper("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 4)
	expectOutput := res{"zqfr", 11}
	if r.r != expectOutput.r {
		t.Errorf("Expected string output, %s, but got %s", expectOutput.r, r.r)
	}

	if r.i != expectOutput.i {
		t.Errorf("Expected index output, %d, but got %d", expectOutput.i, r.i)
	}
}

func TestDetect5(t *testing.T) {
	r := helper("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 14)
	expectOutput := res{"qmgbljsphdztnv", 19}
	if r.r != expectOutput.r {
		t.Errorf("Expected string output, %s, but got %s", expectOutput.r, r.r)
	}

	if r.i != expectOutput.i {
		t.Errorf("Expected index output, %d, but got %d", expectOutput.i, r.i)
	}
}

func TestDetect6(t *testing.T) {
	r := helper("bvwbjplbgvbhsrlpgdmjqwftvncz", 14)
	expectOutput := res{"vbhsrlpgdmjqwf", 23}
	if r.r != expectOutput.r {
		t.Errorf("Expected string output, %s, but got %s", expectOutput.r, r.r)
	}

	if r.i != expectOutput.i {
		t.Errorf("Expected index output, %d, but got %d", expectOutput.i, r.i)
	}
}

func TestDetect7(t *testing.T) {
	r := helper("nppdvjthqldpwncqszvftbrmjlhg", 14)
	expectOutput := res{"ldpwncqszvftbr", 23}
	if r.r != expectOutput.r {
		t.Errorf("Expected string output, %s, but got %s", expectOutput.r, r.r)
	}

	if r.i != expectOutput.i {
		t.Errorf("Expected index output, %d, but got %d", expectOutput.i, r.i)
	}
}

func TestDetect8(t *testing.T) {
	r := helper("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 14)
	expectOutput := res{"wmzdfjlvtqnbhc", 29}
	if r.r != expectOutput.r {
		t.Errorf("Expected string output, %s, but got %s", expectOutput.r, r.r)
	}

	if r.i != expectOutput.i {
		t.Errorf("Expected index output, %d, but got %d", expectOutput.i, r.i)
	}
}

func TestDetect9(t *testing.T) {
	r := helper("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 14)
	expectOutput := res{"jwzlrfnpqdbhtm", 26}
	if r.r != expectOutput.r {
		t.Errorf("Expected string output, %s, but got %s", expectOutput.r, r.r)
	}

	if r.i != expectOutput.i {
		t.Errorf("Expected index output, %d, but got %d", expectOutput.i, r.i)
	}
}
