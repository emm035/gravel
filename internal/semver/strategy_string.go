// Code generated by "stringer -type=Strategy -linecomment"; DO NOT EDIT.

package semver

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[StrategyDate-0]
}

const _Strategy_name = "date"

var _Strategy_index = [...]uint8{0, 4}

func (i Strategy) String() string {
	if i >= Strategy(len(_Strategy_index)-1) {
		return "Strategy(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Strategy_name[_Strategy_index[i]:_Strategy_index[i+1]]
}
