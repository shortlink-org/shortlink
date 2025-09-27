package main

import (
	"fmt"
	"runtime"
	"time"
	"unique"
	"unsafe"
)

// Пример структуры для демонстрации канонизации сложных типов
type Person struct {
	Name string
	Age  int
	City string
}

// Демонстрация базовой канонизации строк
func basicStringExample() {
	fmt.Println("=== Базовая канонизация строк ===")
	
	// Создаем несколько одинаковых строк
	str1 := "Hello, World!"
	str2 := "Hello, World!"
	str3 := "Hello" + ", World!"
	
	// Канонизируем их
	handle1 := unique.Make(str1)
	handle2 := unique.Make(str2)
	handle3 := unique.Make(str3)
	
	fmt.Printf("Оригинальные строки равны: %v\n", str1 == str2 && str2 == str3)
	fmt.Printf("Handle'ы равны: %v\n", handle1 == handle2 && handle2 == handle3)
	fmt.Printf("Адреса оригинальных строк: %p, %p, %p\n", 
		unsafe.Pointer(unsafe.StringData(str1)),
		unsafe.Pointer(unsafe.StringData(str2)), 
		unsafe.Pointer(unsafe.StringData(str3)))
	
	// Получаем канонические значения
	canonical1 := handle1.Value()
	canonical2 := handle2.Value()
	canonical3 := handle3.Value()
	
	fmt.Printf("Адреса канонических строк: %p, %p, %p\n",
		unsafe.Pointer(unsafe.StringData(canonical1)),
		unsafe.Pointer(unsafe.StringData(canonical2)),
		unsafe.Pointer(unsafe.StringData(canonical3)))
	
	fmt.Println()
}

// Демонстрация канонизации структур
func structCanonizationExample() {
	fmt.Println("=== Канонизация структур ===")
	
	// Создаем одинаковые структуры
	person1 := Person{Name: "Алексей", Age: 30, City: "Москва"}
	person2 := Person{Name: "Алексей", Age: 30, City: "Москва"}
	person3 := Person{Name: "Мария", Age: 25, City: "СПб"}
	
	// Канонизируем их
	handle1 := unique.Make(person1)
	handle2 := unique.Make(person2)
	handle3 := unique.Make(person3)
	
	fmt.Printf("person1 == person2: %v\n", person1 == person2)
	fmt.Printf("handle1 == handle2: %v (одинаковые структуры)\n", handle1 == handle2)
	fmt.Printf("handle1 == handle3: %v (разные структуры)\n", handle1 == handle3)
	
	// Адреса в памяти
	fmt.Printf("Адреса оригинальных структур: %p, %p, %p\n", &person1, &person2, &person3)
	
	canonical1 := handle1.Value()
	canonical2 := handle2.Value()
	canonical3 := handle3.Value()
	
	fmt.Printf("Адреса канонических структур: %p, %p, %p\n", &canonical1, &canonical2, &canonical3)
	fmt.Println()
}

// Демонстрация производительности
func performanceComparison() {
	fmt.Println("=== Сравнение производительности ===")
	
	// Создаем множество дублирующихся строк
	strings := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		switch i % 5 {
		case 0:
			strings[i] = "Первая строка"
		case 1:
			strings[i] = "Вторая строка"
		case 2:
			strings[i] = "Третья строка"
		case 3:
			strings[i] = "Четвертая строка"
		case 4:
			strings[i] = "Пятая строка"
		}
	}
	
	// Тест обычного сравнения строк
	start := time.Now()
	normalComparisons := 0
	for i := 0; i < len(strings); i++ {
		for j := i + 1; j < len(strings); j++ {
			if strings[i] == strings[j] {
				normalComparisons++
			}
		}
	}
	normalTime := time.Since(start)
	
	// Канонизируем строки
	handles := make([]unique.Handle[string], len(strings))
	for i, s := range strings {
		handles[i] = unique.Make(s)
	}
	
	// Тест сравнения через handle'ы
	start = time.Now()
	handleComparisons := 0
	for i := 0; i < len(handles); i++ {
		for j := i + 1; j < len(handles); j++ {
			if handles[i] == handles[j] {
				handleComparisons++
			}
		}
	}
	handleTime := time.Since(start)
	
	fmt.Printf("Обычные сравнения: %d совпадений за %v\n", normalComparisons, normalTime)
	fmt.Printf("Handle сравнения: %d совпадений за %v\n", handleComparisons, handleTime)
	fmt.Printf("Ускорение: %.2fx\n", float64(normalTime)/float64(handleTime))
	fmt.Println()
}

// Демонстрация экономии памяти
func memoryOptimizationExample() {
	fmt.Println("=== Оптимизация использования памяти ===")
	
	// Функция для получения статистики памяти
	getMemStats := func() uint64 {
		var m runtime.MemStats
		runtime.GC()
		runtime.ReadMemStats(&m)
		return m.Alloc
	}
	
	allocBefore := getMemStats()
	
	// Создаем много дублирующихся строк БЕЗ канонизации
	var regularStrings []string
	for i := 0; i < 10000; i++ {
		value := fmt.Sprintf("Повторяющаяся строка %d", i%100) // только 100 уникальных значений
		regularStrings = append(regularStrings, value)
	}
	
	allocAfterRegular := getMemStats()
	
	// Создаем много дублирующихся строк С канонизацией
	var canonicalHandles []unique.Handle[string]
	for i := 0; i < 10000; i++ {
		value := fmt.Sprintf("Повторяющаяся строка %d", i%100) // только 100 уникальных значений
		handle := unique.Make(value)
		canonicalHandles = append(canonicalHandles, handle)
	}
	
	allocAfterCanonical := getMemStats()
	
	regularMemory := allocAfterRegular - allocBefore
	canonicalMemory := allocAfterCanonical - allocAfterRegular
	
	fmt.Printf("Память до создания строк: %d байт\n", allocBefore)
	fmt.Printf("Память после обычных строк: %d байт (использовано: %d)\n", allocAfterRegular, regularMemory)
	fmt.Printf("Память после канонических строк: %d байт (использовано: %d)\n", allocAfterCanonical, canonicalMemory)
	
	if regularMemory > canonicalMemory {
		fmt.Printf("Экономия памяти: %.2f%%\n", float64(regularMemory-canonicalMemory)/float64(regularMemory)*100)
	} else {
		fmt.Printf("Канонические строки используют на %.2f%% больше памяти\n", float64(canonicalMemory-regularMemory)/float64(regularMemory)*100)
	}
	
	// Предотвращаем оптимизацию компилятора
	_ = regularStrings
	_ = canonicalHandles
	
	fmt.Println()
}

// Пример кэширования с использованием unique
type StringCache struct {
	cache map[unique.Handle[string]]string
}

func NewStringCache() *StringCache {
	return &StringCache{
		cache: make(map[unique.Handle[string]]string),
	}
}

func (c *StringCache) ProcessString(input string) string {
	handle := unique.Make(input)
	
	if result, exists := c.cache[handle]; exists {
		return result
	}
	
	// Симуляция дорогой операции обработки
	result := fmt.Sprintf("ОБРАБОТАНО: %s", input)
	c.cache[handle] = result
	return result
}

func cachingExample() {
	fmt.Println("=== Пример кэширования с unique ===")
	
	cache := NewStringCache()
	
	// Тестовые данные с дублирующимися строками
	testStrings := []string{
		"hello",
		"world",
		"hello", // дубликат
		"golang",
		"world", // дубликат
		"unique",
		"hello", // еще один дубликат
	}
	
	fmt.Println("Обработка строк:")
	for i, s := range testStrings {
		result := cache.ProcessString(s)
		fmt.Printf("%d. %s -> %s\n", i+1, s, result)
	}
	
	fmt.Printf("\nРазмер кэша: %d уникальных записей\n", len(cache.cache))
	fmt.Println()
}

// Демонстрация использования в качестве ключей карт
func mapKeyExample() {
	fmt.Println("=== Использование Handle как ключей карт ===")
	
	// Карта для подсчета встречаемости строк
	counter := make(map[unique.Handle[string]]int)
	
	words := []string{
		"apple", "banana", "apple", "cherry", "banana", "apple",
		"date", "elderberry", "apple", "banana", "cherry",
	}
	
	for _, word := range words {
		handle := unique.Make(word)
		counter[handle]++
	}
	
	fmt.Println("Подсчет слов:")
	for handle, count := range counter {
		fmt.Printf("%s: %d раз\n", handle.Value(), count)
	}
	fmt.Println()
}

func main() {
	fmt.Println("Демонстрация пакета unique в Go")
	fmt.Println("================================")
	fmt.Println()
	
	basicStringExample()
	structCanonizationExample()
	performanceComparison()
	memoryOptimizationExample()
	cachingExample()
	mapKeyExample()
	
	fmt.Println()
	runAdvancedExamples()
	
	fmt.Println()
	runRealWorldExamples()
	
	fmt.Println("Заключение:")
	fmt.Println("Пакет unique предоставляет мощные возможности для:")
	fmt.Println("- Канонизации значений (интернирование)")
	fmt.Println("- Оптимизации сравнений (сравнение указателей)")
	fmt.Println("- Экономии памяти (дедупликация)")
	fmt.Println("- Эффективного кэширования")
	fmt.Println("- Использования в качестве ключей карт")
	fmt.Println("- Многопоточной работы")
	fmt.Println("- Пользовательских систем интернирования")
}