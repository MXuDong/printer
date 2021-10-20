package Printer

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

var defaultPrinter = DefaultStdoutPrinter()

type ColorKey string

var (
	ErrorColorKey   ColorKey = "error"
	SuccessColorKey ColorKey = "success"
	WarnColorKey    ColorKey = "warn"
	InfoColorKey    ColorKey = "info"
	DebugColorKey   ColorKey = "debug"
)

type PrefixStrFunc func(bePrintStr string) string
type PostfixStrFunc func(bePrintStr string) string
type BeautyStrFunc func(bePrintStr string) string

// DateTimePrefixFunc will return now time with format in '2006-01-02 15:04:05'
func DateTimePrefixFunc(bePrintStr string) string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// EmptyFunc will do nothing
func EmptyFunc(bePrintStr string) string {
	return ""
}

// DefaultBeautyStrFunc will do nothing
func DefaultBeautyStrFunc(bePrintStr string) string {
	return bePrintStr
}

// Printer is core struct of printer
// It save the main colors for output in different kinds.
type Printer struct {
	PrefixStrFunc  PrefixStrFunc
	PostfixStrFunc PostfixStrFunc
	BeautyStrFunc  BeautyStrFunc
	Out            io.Writer

	ColorMap map[ColorKey]Color
}

func (p *Printer) basePrint(value string, preColor, endColor Color) {
	if p.PostfixStrFunc == nil {
		p.PostfixStrFunc = EmptyFunc
	}
	if p.PrefixStrFunc == nil {
		p.PrefixStrFunc = EmptyFunc
	}
	if p.BeautyStrFunc == nil {
		p.BeautyStrFunc = DefaultBeautyStrFunc
	}

	lines := strings.Split(p.BeautyStrFunc(value), "\n")

	for _, line := range lines {
		_, _ = fmt.Fprintf(p.Out, string(preColor)+p.PrefixStrFunc(line)+line+p.PostfixStrFunc(line)+string(endColor)+"\n")
	}
}

// PrintWithColor will append the value of color
func PrintWithColor(key ColorKey, format string, a ...interface{}) {
	defaultPrinter.PrintWithColor(key, format, a...)
}
func (p *Printer) PrintWithColor(key ColorKey, format string, a ...interface{}) {
	value := fmt.Sprintf(format, a...)
	p.basePrint(value, p.ColorMap[key], CReset)
}

// Print the value without color
func Print(format string, a ...interface{}) {
	defaultPrinter.Print(format, a...)
}
func (p *Printer) Print(format string, a ...interface{}) {
	p.PrintWithColor(InfoColorKey, format, a...)
}

// Error will print the error str with ErrorColorKey
func Error(format string, a ...interface{}) {
	defaultPrinter.Error(format, a...)
}
func (p *Printer) Error(format string, a ...interface{}) {
	p.PrintWithColor(ErrorColorKey, format, a...)
}

// Success will print the info with SuccessColorKey
func Success(format string, a ...interface{}) {
	defaultPrinter.Success(format, a...)
}
func (p *Printer) Success(format string, a ...interface{}) {
	p.PrintWithColor(SuccessColorKey, format, a...)
}

// Debug will print the info with DebugColorKey
func Debug(format string, a ...interface{}) {
	defaultPrinter.Debug(format, a...)
}
func (p *Printer) Debug(format string, a ...interface{}) {
	p.PrintWithColor(DebugColorKey, format, a...)
}

// ErrorE will output the error struct in Error
func ErrorE(e error) {
	defaultPrinter.ErrorE(e)
}
func (p *Printer) ErrorE(e error) {
	err := fmt.Sprintf("%v", e)
	p.Error(err)
}

// DefaultStdoutPrinter return default printer
func DefaultStdoutPrinter() *Printer {
	return &Printer{
		PrefixStrFunc:  DateTimePrefixFunc,
		PostfixStrFunc: EmptyFunc,
		Out:            os.Stdout,
		ColorMap: map[ColorKey]Color{
			ErrorColorKey:   CFRed,
			SuccessColorKey: CFGreen,
			WarnColorKey:    CFYellow,
			DebugColorKey:   CFBlue,
			InfoColorKey:    CReset,
		},
	}
}
