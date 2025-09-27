package main

import (
	"testing"
	"unique"
)

// Тесты для демонстрации корректности работы unique
func TestBasicUnique(t *testing.T) {
	s1 := "test string"
	s2 := "test string"
	s3 := "different string"
	
	h1 := unique.Make(s1)
	h2 := unique.Make(s2)
	h3 := unique.Make(s3)
	
	// Одинаковые строки должны давать одинаковые handle'ы
	if h1 != h2 {
		t.Error("Expected equal handles for equal strings")
	}
	
	// Разные строки должны давать разные handle'ы
	if h1 == h3 {
		t.Error("Expected different handles for different strings")
	}
	
	// Значения должны сохраняться
	if h1.Value() != s1 {
		t.Error("Handle value doesn't match original string")
	}
}

func TestStructUnique(t *testing.T) {
	type TestStruct struct {
		A int
		B string
	}
	
	s1 := TestStruct{A: 1, B: "test"}
	s2 := TestStruct{A: 1, B: "test"}
	s3 := TestStruct{A: 2, B: "test"}
	
	h1 := unique.Make(s1)
	h2 := unique.Make(s2)
	h3 := unique.Make(s3)
	
	if h1 != h2 {
		t.Error("Expected equal handles for equal structs")
	}
	
	if h1 == h3 {
		t.Error("Expected different handles for different structs")
	}
}

// Бенчмарки для демонстрации производительности
func BenchmarkStringComparison(b *testing.B) {
	strings := make([]string, 100)
	for i := 0; i < 100; i++ {
		strings[i] = "test string with some length to make comparison slower"
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(strings); j++ {
			for k := j + 1; k < len(strings); k++ {
				_ = strings[j] == strings[k]
			}
		}
	}
}

func BenchmarkHandleComparison(b *testing.B) {
	strings := make([]string, 100)
	handles := make([]unique.Handle[string], 100)
	
	for i := 0; i < 100; i++ {
		strings[i] = "test string with some length to make comparison slower"
		handles[i] = unique.Make(strings[i])
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(handles); j++ {
			for k := j + 1; k < len(handles); k++ {
				_ = handles[j] == handles[k]
			}
		}
	}
}

func BenchmarkUniqueCreation(b *testing.B) {
	testString := "test string for unique creation benchmark"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = unique.Make(testString)
	}
}

func BenchmarkMapWithStringKeys(b *testing.B) {
	m := make(map[string]int)
	testStrings := make([]string, 1000)
	
	for i := 0; i < 1000; i++ {
		testStrings[i] = "test string number " + string(rune(i%100))
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, s := range testStrings {
			m[s]++
		}
	}
}

func BenchmarkMapWithHandleKeys(b *testing.B) {
	m := make(map[unique.Handle[string]]int)
	testStrings := make([]string, 1000)
	handles := make([]unique.Handle[string], 1000)
	
	for i := 0; i < 1000; i++ {
		testStrings[i] = "test string number " + string(rune(i%100))
		handles[i] = unique.Make(testStrings[i])
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, h := range handles {
			m[h]++
		}
	}
}