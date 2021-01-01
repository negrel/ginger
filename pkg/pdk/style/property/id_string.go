// Code generated by "stringer -type=ID -trimprefix=ID"; DO NOT EDIT.

package property

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[IDWidth-1]
	_ = x[IDMinWidth-2]
	_ = x[IDMaxWidth-3]
	_ = x[IDHeight-4]
	_ = x[IDMinHeight-5]
	_ = x[IDMaxHeight-6]
	_ = x[IDMarginLeft-7]
	_ = x[IDMarginTop-8]
	_ = x[IDMarginRight-9]
	_ = x[IDMarginBottom-10]
	_ = x[IDPaddingLeft-11]
	_ = x[IDPaddingTop-12]
	_ = x[IDPaddingRight-13]
	_ = x[IDPaddingBottom-14]
}

const _ID_name = "WidthMinWidthMaxWidthHeightMinHeightMaxHeightMarginLeftMarginTopMarginRightMarginBottomPaddingLeftPaddingTopPaddingRightPaddingBottom"

var _ID_index = [...]uint8{0, 5, 13, 21, 27, 36, 45, 55, 64, 75, 87, 98, 108, 120, 133}

func (i ID) String() string {
	i -= 1
	if i < 0 || i >= ID(len(_ID_index)-1) {
		return "ID(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _ID_name[_ID_index[i]:_ID_index[i+1]]
}