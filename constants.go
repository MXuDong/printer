package printer

type Color string

// Colors
var (
	CFBlack  Color = "\033[30m" // CFBlack  font color: black
	CFRed    Color = "\033[31m" // CFRed    font color: red
	CFGreen  Color = "\033[32m" // CFGreen  font color: green
	CFYellow Color = "\033[33m" // CFYellow font color: yellow
	CFBlue   Color = "\033[34m" // CFBlue   font color: blue
	CFWhite  Color = "\033[37m" // CFWhite  font color: white

	CBBlack  Color = "\033[40m" // CBBlack  background color: black
	CBRed    Color = "\033[41m" // CBRed    background color: red
	CBGreen  Color = "\033[42m" // CBGreen  background color: green
	CBYellow Color = "\033[43m" // CBYellow background color: yellow
	CBBlue   Color = "\033[44m" // CBBlue   background color: blue
	CBWhite  Color = "\033[47m" // CBWhite  background color: white

	CReset Color = "\033[0m" // CReset reset colors
)
