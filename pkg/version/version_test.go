package version

import (
	"testing"
)

func TestCompare_Greater(t *testing.T) {
	res := Compare("3.0.8", "2.0.8")
	if res != GreaterThan {
		t.Errorf("expected 3.0.8 to be GreaterThan 2.0.8")
	}

	res = Compare("3.3.3", "3.3.2")
	if res != GreaterThan {
		t.Errorf("expected 3.3.3 to be GreaterThan 3.3.2")
	}
}

func TestCompare_Less(t *testing.T) {
	res := Compare("3.1.3", "3.2.0")
	if res != LessThan {
		t.Errorf("expected 3.1.3 to be LessThan 3.2.0")
	}
}

func TestCompare_Equal(t *testing.T) {
	res := Compare("3.3.0", "3.3.0")
	if res != EqualTo {
		t.Errorf("expected 3.3.0 to be EqualTo 3.3.0")
	}
}

func TestCompare_Off(t *testing.T) {
	res := Compare("0.2.0", "3.0")
	if res != LessThan {
		t.Errorf("expected 0.2.0 to be LessThan 3.0")
	}

	res = Compare("2.0", "5.2.0")
	if res != LessThan {
		t.Errorf("expected 2.0 to be LessThan 5.2.0")
	}
}

func TestGetConstraintType(t *testing.T) {
	res := GetConstraintType("~> 4.5.3")
	if res != BeApproxGreaterThan {
		t.Errorf("expected constraint ~> to be BeApproxGreaterTahn")
	}

	res = GetConstraintType("< 4.5.3")
	if res != BeLessThan {
		t.Errorf("expected constraint < to be BeLessThan")
	}
}

func TestMatchConstraint(t *testing.T) {
	v := "2.5.0"
	c := "= 2.5.0"
	if !MatchConstraint(v, c) {
		t.Errorf("expected MatchConstraint() to return true")
	}

	v = "1.5.0"
	c = "<= 2.5.0"
	if !MatchConstraint(v, c) {
		t.Errorf("expected MatchConstraint() to return true")
	}

	v = "1.5.0"
	c = ">= 2.5.0"
	if MatchConstraint(v, c) {
		t.Errorf("expected MatchConstraint() to return false")
	}

	v = "2.6.3"
	c = "~> 2.6.4"
	if MatchConstraint(v, c) {
		t.Errorf("expected MatchConstraint() to return false")
	}

	v = "2.6.7"
	c = "~> 2.4"
	if !MatchConstraint(v, c) {
		t.Errorf("expected MatchConstraint() to return true")
	}
}

func TestApproxGreaterThan(t *testing.T) {
	v := "2.4.3"
	c := "2.3"
	if !ApproxGreaterThan(v, c) {
		t.Errorf("expected 2.4.3 to be approximatly greater than 2.3")
	}

	v = "2.5.3"
	c = "2.6"
	if ApproxGreaterThan(v, c) {
		t.Errorf("expected 2.5.3 to be less than 2.6")
	}
}
