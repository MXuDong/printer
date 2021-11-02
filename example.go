package printer

import "strconv"

// PreFuncWithLineNumber example: For deal with the line number
// If you need log rotate, you can reset prefunc by this func, to implement: to write different file from zero line
func PreFuncWithLineNumber() PrefixStrFunc {
	// if you use the line number mod, the pre and hou func can't have any \n
	count := 0
	return func(bePrintStr string) string {
		count++
		return strconv.Itoa(count)
	}
}

