// Code generated by "stringer -type=Segment -linecomment"; DO NOT EDIT.

package semver

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SegmentPatch-0]
	_ = x[SegmentMinor-1]
	_ = x[SegmentMajor-2]
}

const _Segment_name = "patchminormajor"

var _Segment_index = [...]uint8{0, 5, 10, 15}

func (i Segment) String() string {
	if i >= Segment(len(_Segment_index)-1) {
		return "Segment(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Segment_name[_Segment_index[i]:_Segment_index[i+1]]
}