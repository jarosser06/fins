// Package version provides a VersionComparison method to make
// comparing versions easy.
package version

import (
	"regexp"
	"strconv"
	"strings"
)

type (
	VersionComparison int
	ConstraintType    int
)

// Version Constrain Type
const (
	NoMatch ConstraintType = iota
	BeEqualTo
	BeGreaterThan
	BeLessThan
	BeGreaterThanOrEqual
	BeLessThanOrEqual
	BeApproxGreaterThan
)

// Version Comparison
const (
	GreaterThan VersionComparison = iota
	LessThan
	EqualTo
)

// VersionCompare takes a version and compares it to another
// version.  It returns GreaterThan, LessThan, or EqualTo
func Compare(ver string, compTo string) VersionComparison {
	verArray := strings.Split(ver, ".")
	compToArray := strings.Split(compTo, ".")

	diff := len(verArray) - len(compToArray)
	switch {
	case diff > 0:
		for i := 0; i < diff; i++ {
			if v, _ := strconv.Atoi(verArray[i]); v > 0 {
				return GreaterThan
			}
		}
	case diff < 0:
		for i, x := 0, diff; x < 0; x++ {
			i++
			if v, _ := strconv.Atoi(compToArray[i]); v > 0 {
				return LessThan
			}
		}
	}

	for i := 0; i < len(verArray); i++ {
		c, _ := strconv.Atoi(compToArray[i])
		v, _ := strconv.Atoi(verArray[i])
		if v > c {
			return GreaterThan
		} else if v < c {
			return LessThan
		}
	}

	return EqualTo
}

func GetConstraintType(verConstraint string) ConstraintType {
	var conType ConstraintType
	re := regexp.MustCompile("^[~><=]+")
	switch re.FindString(verConstraint) {
	case "=":
		conType = BeEqualTo
	case ">":
		conType = BeGreaterThan
	case "<":
		conType = BeLessThan
	case ">=":
		conType = BeGreaterThanOrEqual
	case "<=":
		conType = BeLessThanOrEqual
	case "~>":
		conType = BeApproxGreaterThan
	default:
		conType = NoMatch
	}

	return conType
}

func ApproxGreaterThan(ver string, compareTo string) bool {
	result := false
	verArray := strings.Split(ver, ".")
	compareToArray := strings.Split(compareTo, ".")

	for i := 0; i < len(compareToArray); i++ {
		v, _ := strconv.Atoi(verArray[i])
		c, _ := strconv.Atoi(compareToArray[i])

		// If we are at the last element of the verison
		// check if the corresponding element in ver is
		// greater than or equal to
		if i == (len(compareToArray)-1) && v >= c {
			result = true
		} else if v != c {
			break
		}
	}

	return result
}

func MatchConstraint(ver string, constraint string) bool {
	re := regexp.MustCompile("(([0-9]+).)+([0-9]+)")
	compVer := re.FindString(constraint)
	result := false
	comparison := Compare(ver, compVer)

	switch GetConstraintType(constraint) {
	case BeEqualTo:
		if comparison == EqualTo {
			result = true
		}
	case BeGreaterThan:
		if comparison == GreaterThan {
			result = true
		}
	case BeLessThan:
		if comparison == LessThan {
			result = true
		}
	case BeGreaterThanOrEqual:
		if comparison == GreaterThan || comparison == EqualTo {
			result = true
		}
	case BeLessThanOrEqual:
		if comparison == LessThan || comparison == EqualTo {
			result = true
		}
	case BeApproxGreaterThan:
		if ApproxGreaterThan(ver, compVer) {
			result = true
		}
	}

	return result
}
