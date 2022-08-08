package cpmath

import (
	//"log"
	"testing"
)

func TestPagination(t *testing.T) {
	limit, offset := LimitOffset(1, 50)

	if limit != 50 {
		t.Errorf("Limit test failed")
	}

	if offset != 0 {
		t.Errorf("Offset test failed")
	}

	limit, offset = LimitOffset(3, 10)

	if limit != 10 {
		t.Errorf("Limit test failed")
	}

	if offset != 20 {
		t.Errorf("Offset test failed")
	}

}
