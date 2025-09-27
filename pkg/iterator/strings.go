// Package iterator provides string iterator functions similar to those added in Go 1.24.
package iterator

import (
	"iter"
	"strings"
	"unicode"
)

// Lines returns an iterator over the newline-terminated lines in a string.
// The string s is split after each UTF-8-encoded newline byte (\n).
// The yielded strings do not include the newline character.
func Lines(s string) iter.Seq[string] {
	return func(yield func(string) bool) {
		if s == "" {
			return
		}
		
		start := 0
		for i, r := range s {
			if r == '\n' {
				if !yield(s[start:i]) {
					return
				}
				start = i + 1
			}
		}
		
		// Yield the last line (even if empty due to trailing newline)
		if start <= len(s) {
			yield(s[start:])
		}
	}
}

// SplitSeq returns an iterator over substrings of s separated by sep.
// If sep is empty, SplitSeq splits after each UTF-8 sequence.
// The substrings yielded by the iterator are the same as those
// returned by strings.Split(s, sep), but without creating the slice.
func SplitSeq(s, sep string) iter.Seq[string] {
	return func(yield func(string) bool) {
		if sep == "" {
			// Split after each UTF-8 sequence
			for _, r := range s {
				if !yield(string(r)) {
					return
				}
			}
			return
		}
		
		if sep == s {
			if !yield("") {
				return
			}
			yield("")
			return
		}
		
		start := 0
		for {
			i := strings.Index(s[start:], sep)
			if i < 0 {
				// No more separators found
				yield(s[start:])
				return
			}
			
			if !yield(s[start : start+i]) {
				return
			}
			start += i + len(sep)
		}
	}
}

// SplitAfterSeq returns an iterator over substrings of s split after each
// instance of sep. The substrings yielded by the iterator include sep.
// If sep is empty, SplitAfterSeq splits after each UTF-8 sequence.
func SplitAfterSeq(s, sep string) iter.Seq[string] {
	return func(yield func(string) bool) {
		if sep == "" {
			// Split after each UTF-8 sequence
			for _, r := range s {
				if !yield(string(r)) {
					return
				}
			}
			return
		}
		
		start := 0
		for {
			i := strings.Index(s[start:], sep)
			if i < 0 {
				// No more separators found
				if start < len(s) {
					yield(s[start:])
				}
				return
			}
			
			end := start + i + len(sep)
			if !yield(s[start:end]) {
				return
			}
			start = end
		}
	}
}

// FieldsSeq returns an iterator over substrings of s split around runs
// of whitespace characters, as defined by unicode.IsSpace.
func FieldsSeq(s string) iter.Seq[string] {
	return FieldsFuncSeq(s, unicode.IsSpace)
}

// FieldsFuncSeq returns an iterator over substrings of s split around runs
// of Unicode code points c satisfying f(c).
func FieldsFuncSeq(s string, f func(rune) bool) iter.Seq[string] {
	return func(yield func(string) bool) {
		// Skip leading separators
		start := 0
		for start < len(s) {
			r, size := decodeRune(s[start:])
			if !f(r) {
				break
			}
			start += size
		}
		
		for start < len(s) {
			// Find the end of the current field
			end := start
			for end < len(s) {
				r, size := decodeRune(s[end:])
				if f(r) {
					break
				}
				end += size
			}
			
			if !yield(s[start:end]) {
				return
			}
			
			// Skip separators
			start = end
			for start < len(s) {
				r, size := decodeRune(s[start:])
				if !f(r) {
					break
				}
				start += size
			}
		}
	}
}

// decodeRune is a helper function to decode a rune from the beginning of a string.
// It returns the rune and its size in bytes.
func decodeRune(s string) (rune, int) {
	if len(s) == 0 {
		return 0, 0
	}
	
	b := s[0]
	if b < 0x80 {
		return rune(b), 1
	}
	
	// For UTF-8 sequences, we need to handle them properly
	for i, r := range s {
		if i == 0 {
			continue
		}
		return r, len(s) - i + 1
	}
	
	// If we get here, s contains only one rune
	for _, r := range s {
		return r, len(s)
	}
	
	return 0, 0
}

// WordsSeq returns an iterator over words in the string, where words are
// defined as sequences of non-whitespace characters.
func WordsSeq(s string) iter.Seq[string] {
	return FieldsSeq(s)
}

// CharsSeq returns an iterator over each character (rune) in the string.
func CharsSeq(s string) iter.Seq[rune] {
	return func(yield func(rune) bool) {
		for _, r := range s {
			if !yield(r) {
				return
			}
		}
	}
}

// BytesSeq returns an iterator over each byte in the string.
func BytesSeq(s string) iter.Seq[byte] {
	return func(yield func(byte) bool) {
		for i := 0; i < len(s); i++ {
			if !yield(s[i]) {
				return
			}
		}
	}
}

// ReverseLinesSeq returns an iterator over the lines in reverse order.
func ReverseLinesSeq(s string) iter.Seq[string] {
	return func(yield func(string) bool) {
		lines := strings.Split(s, "\n")
		for i := len(lines) - 1; i >= 0; i-- {
			if !yield(lines[i]) {
				return
			}
		}
	}
}

// ParagraphsSeq returns an iterator over paragraphs separated by blank lines.
func ParagraphsSeq(s string) iter.Seq[string] {
	return func(yield func(string) bool) {
		lines := strings.Split(s, "\n")
		var paragraph strings.Builder
		
		for _, line := range lines {
			trimmed := strings.TrimSpace(line)
			if trimmed == "" {
				// Empty line - end of paragraph
				if paragraph.Len() > 0 {
					if !yield(strings.TrimSpace(paragraph.String())) {
						return
					}
					paragraph.Reset()
				}
			} else {
				if paragraph.Len() > 0 {
					paragraph.WriteString("\n")
				}
				paragraph.WriteString(line)
			}
		}
		
		// Yield the last paragraph if it exists
		if paragraph.Len() > 0 {
			yield(strings.TrimSpace(paragraph.String()))
		}
	}
}

// TokensSeq returns an iterator over tokens separated by any of the given separators.
func TokensSeq(s string, separators string) iter.Seq[string] {
	return func(yield func(string) bool) {
		if s == "" {
			return
		}
		
		isSeparator := func(r rune) bool {
			for _, sep := range separators {
				if r == sep {
					return true
				}
			}
			return false
		}
		
		start := 0
		for i, r := range s {
			if isSeparator(r) {
				if start < i {
					if !yield(s[start:i]) {
						return
					}
				}
				start = i + 1
			}
		}
		
		// Yield the last token if it exists
		if start < len(s) {
			yield(s[start:])
		}
	}
}