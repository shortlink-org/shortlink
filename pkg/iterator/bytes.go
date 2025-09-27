// Package iterator provides byte slice iterator functions similar to those added in Go 1.24.
package iterator

import (
	"bytes"
	"iter"
	"unicode"
)

// LinesBytes returns an iterator over the newline-terminated lines in a byte slice.
// The byte slice b is split after each newline byte (\n).
// The yielded byte slices do not include the newline character.
func LinesBytes(b []byte) iter.Seq[[]byte] {
	return func(yield func([]byte) bool) {
		if len(b) == 0 {
			return
		}
		
		start := 0
		for i, c := range b {
			if c == '\n' {
				if !yield(b[start:i]) {
					return
				}
				start = i + 1
			}
		}
		
		// Yield the last line if it doesn't end with newline
		if start < len(b) {
			yield(b[start:])
		}
	}
}

// SplitSeqBytes returns an iterator over subslices of b separated by sep.
// If sep is empty, SplitSeqBytes splits after each byte.
// The subslices yielded by the iterator are the same as those
// returned by bytes.Split(b, sep), but without creating the slice.
func SplitSeqBytes(b, sep []byte) iter.Seq[[]byte] {
	return func(yield func([]byte) bool) {
		if len(sep) == 0 {
			// Split after each byte
			for _, c := range b {
				if !yield([]byte{c}) {
					return
				}
			}
			return
		}
		
		if bytes.Equal(sep, b) {
			if !yield([]byte{}) {
				return
			}
			yield([]byte{})
			return
		}
		
		start := 0
		for {
			i := bytes.Index(b[start:], sep)
			if i < 0 {
				// No more separators found
				yield(b[start:])
				return
			}
			
			if !yield(b[start : start+i]) {
				return
			}
			start += i + len(sep)
		}
	}
}

// SplitAfterSeqBytes returns an iterator over subslices of b split after each
// instance of sep. The subslices yielded by the iterator include sep.
// If sep is empty, SplitAfterSeqBytes splits after each byte.
func SplitAfterSeqBytes(b, sep []byte) iter.Seq[[]byte] {
	return func(yield func([]byte) bool) {
		if len(sep) == 0 {
			// Split after each byte
			for _, c := range b {
				if !yield([]byte{c}) {
					return
				}
			}
			return
		}
		
		start := 0
		for {
			i := bytes.Index(b[start:], sep)
			if i < 0 {
				// No more separators found
				if start < len(b) {
					yield(b[start:])
				}
				return
			}
			
			end := start + i + len(sep)
			if !yield(b[start:end]) {
				return
			}
			start = end
		}
	}
}

// FieldsSeqBytes returns an iterator over subslices of b split around runs
// of whitespace bytes, as defined by unicode.IsSpace.
func FieldsSeqBytes(b []byte) iter.Seq[[]byte] {
	return FieldsFuncSeqBytes(b, func(r rune) bool {
		return unicode.IsSpace(r)
	})
}

// FieldsFuncSeqBytes returns an iterator over subslices of b split around runs
// of UTF-8-encoded code points c satisfying f(c).
func FieldsFuncSeqBytes(b []byte, f func(rune) bool) iter.Seq[[]byte] {
	return func(yield func([]byte) bool) {
		// Skip leading separators
		start := 0
		for start < len(b) {
			r, size := decodeRuneBytes(b[start:])
			if !f(r) {
				break
			}
			start += size
		}
		
		for start < len(b) {
			// Find the end of the current field
			end := start
			for end < len(b) {
				r, size := decodeRuneBytes(b[end:])
				if f(r) {
					break
				}
				end += size
			}
			
			if !yield(b[start:end]) {
				return
			}
			
			// Skip separators
			start = end
			for start < len(b) {
				r, size := decodeRuneBytes(b[start:])
				if !f(r) {
					break
				}
				start += size
			}
		}
	}
}

// decodeRuneBytes is a helper function to decode a rune from the beginning of a byte slice.
// It returns the rune and its size in bytes.
func decodeRuneBytes(b []byte) (rune, int) {
	if len(b) == 0 {
		return 0, 0
	}
	
	if b[0] < 0x80 {
		return rune(b[0]), 1
	}
	
	// For UTF-8 sequences, we use the standard library
	r, size := decodeUTF8(b)
	return r, size
}

// decodeUTF8 decodes a UTF-8 sequence from a byte slice.
func decodeUTF8(b []byte) (rune, int) {
	if len(b) == 0 {
		return 0, 0
	}
	
	c0 := b[0]
	
	// ASCII
	if c0 < 0x80 {
		return rune(c0), 1
	}
	
	// Determine the length of the UTF-8 sequence
	var length int
	if c0 < 0xE0 {
		length = 2
	} else if c0 < 0xF0 {
		length = 3
	} else {
		length = 4
	}
	
	if len(b) < length {
		return 0xFFFD, 1 // replacement character
	}
	
	// Decode the rune
	var r rune
	switch length {
	case 2:
		if b[1]&0xC0 != 0x80 {
			return 0xFFFD, 1
		}
		r = rune(c0&0x1F)<<6 | rune(b[1]&0x3F)
	case 3:
		if b[1]&0xC0 != 0x80 || b[2]&0xC0 != 0x80 {
			return 0xFFFD, 1
		}
		r = rune(c0&0x0F)<<12 | rune(b[1]&0x3F)<<6 | rune(b[2]&0x3F)
	case 4:
		if b[1]&0xC0 != 0x80 || b[2]&0xC0 != 0x80 || b[3]&0xC0 != 0x80 {
			return 0xFFFD, 1
		}
		r = rune(c0&0x07)<<18 | rune(b[1]&0x3F)<<12 | rune(b[2]&0x3F)<<6 | rune(b[3]&0x3F)
	}
	
	return r, length
}

// WordsSeqBytes returns an iterator over words in the byte slice, where words are
// defined as sequences of non-whitespace bytes.
func WordsSeqBytes(b []byte) iter.Seq[[]byte] {
	return FieldsSeqBytes(b)
}

// CharsSeqBytes returns an iterator over each character (rune) in the byte slice.
func CharsSeqBytes(b []byte) iter.Seq[rune] {
	return func(yield func(rune) bool) {
		for len(b) > 0 {
			r, size := decodeRuneBytes(b)
			if !yield(r) {
				return
			}
			b = b[size:]
		}
	}
}

// BytesSeqBytes returns an iterator over each byte in the byte slice.
func BytesSeqBytes(b []byte) iter.Seq[byte] {
	return func(yield func(byte) bool) {
		for _, c := range b {
			if !yield(c) {
				return
			}
		}
	}
}

// ReverseLinesSeqBytes returns an iterator over the lines in reverse order.
func ReverseLinesSeqBytes(b []byte) iter.Seq[[]byte] {
	return func(yield func([]byte) bool) {
		lines := bytes.Split(b, []byte{'\n'})
		for i := len(lines) - 1; i >= 0; i-- {
			if !yield(lines[i]) {
				return
			}
		}
	}
}

// ParagraphsSeqBytes returns an iterator over paragraphs separated by blank lines.
func ParagraphsSeqBytes(b []byte) iter.Seq[[]byte] {
	return func(yield func([]byte) bool) {
		lines := bytes.Split(b, []byte{'\n'})
		var paragraph []byte
		
		for _, line := range lines {
			trimmed := bytes.TrimSpace(line)
			if len(trimmed) == 0 {
				// Empty line - end of paragraph
				if len(paragraph) > 0 {
					if !yield(bytes.TrimSpace(paragraph)) {
						return
					}
					paragraph = nil
				}
			} else {
				if len(paragraph) > 0 {
					paragraph = append(paragraph, '\n')
				}
				paragraph = append(paragraph, line...)
			}
		}
		
		// Yield the last paragraph if it exists
		if len(paragraph) > 0 {
			yield(bytes.TrimSpace(paragraph))
		}
	}
}

// TokensSeqBytes returns an iterator over tokens separated by any of the given separator bytes.
func TokensSeqBytes(b []byte, separators []byte) iter.Seq[[]byte] {
	return func(yield func([]byte) bool) {
		if len(b) == 0 {
			return
		}
		
		isSeparator := func(c byte) bool {
			for _, sep := range separators {
				if c == sep {
					return true
				}
			}
			return false
		}
		
		start := 0
		for i, c := range b {
			if isSeparator(c) {
				if start < i {
					if !yield(b[start:i]) {
						return
					}
				}
				start = i + 1
			}
		}
		
		// Yield the last token if it exists
		if start < len(b) {
			yield(b[start:])
		}
	}
}