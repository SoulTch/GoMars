// Code generated by "stringer -type=payWith"; DO NOT EDIT.

package core

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[withIron-0]
	_ = x[withTitanium-1]
	_ = x[withHeat-2]
}

const _payWith_name = "withIronwithTitaniumwithHeat"

var _payWith_index = [...]uint8{0, 8, 20, 28}

func (i payWith) String() string {
	if i < 0 || i >= payWith(len(_payWith_index)-1) {
		return "payWith(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _payWith_name[_payWith_index[i]:_payWith_index[i+1]]
}
