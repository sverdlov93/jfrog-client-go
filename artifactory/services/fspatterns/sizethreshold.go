package fspatterns

// ThresholdCondition represents whether the threshold is for files above or below a specified size.
type ThresholdCondition int

const (
	// GreaterThan is greater & equal
	GreaterEqualThan ThresholdCondition = iota
	// LessThan is only less
	LessThan
)

type SizeThreshold struct {
	SizeInBytes int64
	Condition   ThresholdCondition
}

func (st SizeThreshold) IsSizeWithinThreshold(actualSize int64) bool {
	switch st.Condition {
	case GreaterEqualThan:
		return actualSize >= st.SizeInBytes
	case LessThan:
		return actualSize < st.SizeInBytes
	default:
		return false
	}
}
