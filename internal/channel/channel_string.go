// Code generated by "stringer -type=Channel"; DO NOT EDIT.

package channel

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Telegram-1]
	_ = x[Email-2]
	_ = x[Mock-3]
}

const _Channel_name = "TelegramEmailMock"

var _Channel_index = [...]uint8{0, 8, 13, 17}

func (i Channel) String() string {
	i -= 1
	if i < 0 || i >= Channel(len(_Channel_index)-1) {
		return "Channel(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Channel_name[_Channel_index[i]:_Channel_index[i+1]]
}
