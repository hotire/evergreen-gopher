package test

import "testing"

func TestSum(t *testing.T)  {
	result := Sum(1, 2)
	if result != 3 {
		t.Errorf("expected:%d actual:%d", 3, result)
	}
}