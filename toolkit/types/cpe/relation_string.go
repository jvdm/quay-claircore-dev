// Code generated by "stringer -type Relation -linecomment"; DO NOT EDIT.

package cpe

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Superset-1]
	_ = x[Subset-2]
	_ = x[Equal-3]
	_ = x[Disjoint-4]
}

const _Relation_name = "⊃⊂=≠"

var _Relation_index = [...]uint8{0, 3, 6, 7, 10}

func (i Relation) String() string {
	i -= 1
	if i >= Relation(len(_Relation_index)-1) {
		return "Relation(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Relation_name[_Relation_index[i]:_Relation_index[i+1]]
}
