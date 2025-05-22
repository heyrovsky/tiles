package utils

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
	"time"
)

type Writer struct {
	mu          sync.Mutex
	ticker      *time.Ticker
	stopChan    chan bool
	frameIndex  int
	currentLine string
	isActive    bool
	prefix      string
	style       Style
}

type Style struct {
	Frames      []string
	Interval    time.Duration
	SuccessIcon string
	ErrorIcon   string
	WorkingIcon string
}

var (
	// SpinnerStyle - Default spinning dots
	SpinnerStyle = Style{
		Frames:      []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
		Interval:    100 * time.Millisecond,
		SuccessIcon: "✓",
		ErrorIcon:   "✗",
		WorkingIcon: "⠋",
	}

	// CircleStyle - Spinning circle
	CircleStyle = Style{
		Frames:      []string{"◐", "◓", "◑", "◒"},
		Interval:    150 * time.Millisecond,
		SuccessIcon: "●",
		ErrorIcon:   "○",
		WorkingIcon: "◐",
	}

	// BarStyle - Classic bar spinner
	BarStyle = Style{
		Frames:      []string{"|", "/", "-", "\\"},
		Interval:    120 * time.Millisecond,
		SuccessIcon: "✓",
		ErrorIcon:   "✗",
		WorkingIcon: "|",
	}

	// DotsStyle - Pulsing dots
	DotsStyle = Style{
		Frames:      []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"},
		Interval:    80 * time.Millisecond,
		SuccessIcon: "●",
		ErrorIcon:   "○",
		WorkingIcon: "⣾",
	}

	// ArrowStyle - Arrow spinner
	ArrowStyle = Style{
		Frames:      []string{"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"},
		Interval:    150 * time.Millisecond,
		SuccessIcon: "→",
		ErrorIcon:   "×",
		WorkingIcon: "←",
	}
)

// Options for creating a new progress writer
type Options struct {
	Prefix string
	Style  Style
}

// New creates a new progress writer with the given options
func New(opts Options) *Writer {
	if opts.Style.Frames == nil {
		opts.Style = DotsStyle
	}

	return &Writer{
		prefix:   opts.Prefix,
		style:    opts.Style,
		stopChan: make(chan bool, 1),
	}
}

// NewDefault creates a new progress writer with default spinner style
func NewDefault() *Writer {
	return New(Options{Style: DotsStyle})
}

// NewWithPrefix creates a new progress writer with a prefix and default style
func NewWithPrefix(prefix string) *Writer {
	return New(Options{Prefix: prefix, Style: DotsStyle})
}

// NewWithStyle creates a new progress writer with a custom style
func NewWithStyle(style Style) *Writer {
	return New(Options{Style: style})
}

// Write implements io.Writer interface
func (w *Writer) Write(p []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	text := string(p)
	lines := strings.Split(strings.TrimSpace(text), "\n")

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		// Clean the line from any ANSI escape sequences
		cleanLine := w.cleanLine(line)

		if cleanLine != "" {
			w.currentLine = cleanLine
			if !w.isActive {
				w.startAnimation()
			}
			w.updateDisplay()
		}
	}

	return len(p), nil
}

// Start begins the progress animation with an initial message
func (w *Writer) Start(message string) {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.currentLine = message
	if !w.isActive {
		w.startAnimation()
	}
	w.updateDisplay()
}

// Update changes the current progress message
func (w *Writer) Update(message string) {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.currentLine = message
	if !w.isActive {
		w.startAnimation()
	}
	w.updateDisplay()
}

// Success stops the animation and shows a success message
func (w *Writer) Success(message string) {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.stop()
	w.printFinal(w.style.SuccessIcon, message)
}

// Error stops the animation and shows an error message
func (w *Writer) Error(message string) {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.stop()
	w.printFinal(w.style.ErrorIcon, message)
}

// Stop stops the animation and shows the current message with success icon
func (w *Writer) Stop() {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.stop()
	if w.currentLine != "" {
		w.printFinal(w.style.SuccessIcon, w.currentLine)
	}
}

// StopWithMessage stops the animation and shows a custom final message
func (w *Writer) StopWithMessage(message string) {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.stop()
	w.printFinal(w.style.SuccessIcon, message)
}

// Internal methods

func (w *Writer) cleanLine(line string) string {
	// Remove ANSI escape sequences
	ansiRegex := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	cleaned := ansiRegex.ReplaceAllString(line, "")

	// Remove carriage returns and extra whitespace
	cleaned = strings.ReplaceAll(cleaned, "\r", "")
	cleaned = strings.TrimSpace(cleaned)

	return cleaned
}

func (w *Writer) startAnimation() {
	if w.isActive {
		return
	}

	w.isActive = true
	w.ticker = time.NewTicker(w.style.Interval)

	go func() {
		for {
			select {
			case <-w.ticker.C:
				w.mu.Lock()
				w.frameIndex = (w.frameIndex + 1) % len(w.style.Frames)
				w.updateDisplay()
				w.mu.Unlock()
			case <-w.stopChan:
				return
			}
		}
	}()
}

func (w *Writer) updateDisplay() {
	// Clear current line and move cursor to beginning
	fmt.Print("\r\033[K")

	// Build display string
	spinner := w.style.Frames[w.frameIndex]
	display := fmt.Sprintf("%s %s", spinner, w.currentLine)

	if w.prefix != "" {
		display = fmt.Sprintf("%s %s", w.prefix, display)
	}

	// Print without newline
	fmt.Print(display)
}

func (w *Writer) stop() {
	if !w.isActive {
		return
	}

	w.isActive = false

	if w.ticker != nil {
		w.ticker.Stop()
	}

	select {
	case w.stopChan <- true:
	default:
	}
}

func (w *Writer) printFinal(icon, message string) {
	// Clear the line
	fmt.Print("\r\033[K")

	finalMsg := fmt.Sprintf("%s %s", icon, message)
	if w.prefix != "" {
		finalMsg = fmt.Sprintf("%s %s", w.prefix, finalMsg)
	}
	fmt.Println(finalMsg)
}

// Utility functions

// WriteFunc is a helper function that creates a temporary progress writer for a function
func WriteFunc(message string, fn func() error) error {
	w := NewDefault()
	w.Start(message)
	defer w.Stop()

	err := fn()
	if err != nil {
		w.Error(fmt.Sprintf("Failed: %v", err))
		return err
	}

	w.Success("Completed successfully")
	return nil
}

// WriteFuncWithStyle is like WriteFunc but with custom style
func WriteFuncWithStyle(message string, style Style, fn func() error) error {
	w := NewWithStyle(style)
	w.Start(message)
	defer w.Stop()

	err := fn()
	if err != nil {
		w.Error(fmt.Sprintf("Failed: %v", err))
		return err
	}

	w.Success("Completed successfully")
	return nil
}
