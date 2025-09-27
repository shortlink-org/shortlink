package main

import (
	"fmt"
	"sync"
	"time"
	"unique"
)

// Пример использования unique в многопоточной среде
func concurrentExample() {
	fmt.Println("=== Многопоточная работа с unique ===")
	
	var wg sync.WaitGroup
	results := make(chan unique.Handle[string], 100)
	
	// Создаем несколько горутин, которые канонизируют строки
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				str := fmt.Sprintf("Строка %d", j%3) // Только 3 уникальных строки
				handle := unique.Make(str)
				results <- handle
			}
		}(i)
	}
	
	// Закрываем канал после завершения всех горутин
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// Собираем результаты и проверяем уникальность
	handleMap := make(map[unique.Handle[string]]int)
	for handle := range results {
		handleMap[handle]++
	}
	
	fmt.Printf("Обработано %d строк\n", 100)
	fmt.Printf("Уникальных handle'ов: %d\n", len(handleMap))
	
	for handle, count := range handleMap {
		fmt.Printf("'%s': %d раз\n", handle.Value(), count)
	}
	fmt.Println()
}

// Пример системы интернирования с метриками
type InternPool[T comparable] struct {
	mu          sync.RWMutex
	handles     map[T]unique.Handle[T]
	hitCount    int64
	missCount   int64
	totalValues int64
}

func NewInternPool[T comparable]() *InternPool[T] {
	return &InternPool[T]{
		handles: make(map[T]unique.Handle[T]),
	}
}

func (p *InternPool[T]) Intern(value T) unique.Handle[T] {
	p.mu.RLock()
	if handle, exists := p.handles[value]; exists {
		p.hitCount++
		p.totalValues++
		p.mu.RUnlock()
		return handle
	}
	p.mu.RUnlock()
	
	p.mu.Lock()
	defer p.mu.Unlock()
	
	// Двойная проверка
	if handle, exists := p.handles[value]; exists {
		p.hitCount++
		p.totalValues++
		return handle
	}
	
	handle := unique.Make(value)
	p.handles[value] = handle
	p.missCount++
	p.totalValues++
	return handle
}

func (p *InternPool[T]) Stats() (hitRate float64, uniqueValues int, totalRequests int64) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	
	if p.totalValues == 0 {
		return 0, 0, 0
	}
	
	return float64(p.hitCount) / float64(p.totalValues), len(p.handles), p.totalValues
}

func internPoolExample() {
	fmt.Println("=== Пример системы интернирования с метриками ===")
	
	pool := NewInternPool[string]()
	
	// Симуляция работы с повторяющимися строками
	testData := []string{
		"user:123", "user:456", "user:123", "user:789", "user:123",
		"session:abc", "session:def", "session:abc", "user:456",
		"token:xyz", "token:uvw", "token:xyz", "session:abc",
	}
	
	fmt.Println("Интернирование строк:")
	for i, data := range testData {
		handle := pool.Intern(data)
		fmt.Printf("%d. %s -> Handle: %p\n", i+1, data, &handle)
	}
	
	hitRate, uniqueValues, totalRequests := pool.Stats()
	fmt.Printf("\nСтатистика пула:\n")
	fmt.Printf("Уникальных значений: %d\n", uniqueValues)
	fmt.Printf("Общих запросов: %d\n", totalRequests)
	fmt.Printf("Hit rate: %.2f%%\n", hitRate*100)
	fmt.Println()
}

// Пример с пользовательским типом данных
type UserID struct {
	Namespace string
	ID        int64
}

type UserData struct {
	Handle unique.Handle[UserID]
	Name   string
	Email  string
}

func customTypeExample() {
	fmt.Println("=== Пример с пользовательскими типами ===")
	
	// Создаем пользователей с дублирующимися ID
	users := []UserData{
		{Handle: unique.Make(UserID{"app", 123}), Name: "Алексей", Email: "alex@example.com"},
		{Handle: unique.Make(UserID{"app", 456}), Name: "Мария", Email: "maria@example.com"},
		{Handle: unique.Make(UserID{"app", 123}), Name: "Алексей", Email: "alex@example.com"}, // дубликат
		{Handle: unique.Make(UserID{"admin", 789}), Name: "Админ", Email: "admin@example.com"},
		{Handle: unique.Make(UserID{"app", 456}), Name: "Мария", Email: "maria@example.com"}, // дубликат
	}
	
	// Группируем пользователей по уникальным ID
	userGroups := make(map[unique.Handle[UserID]][]UserData)
	for _, user := range users {
		userGroups[user.Handle] = append(userGroups[user.Handle], user)
	}
	
	fmt.Printf("Всего пользователей: %d\n", len(users))
	fmt.Printf("Уникальных ID: %d\n", len(userGroups))
	
	fmt.Println("\nГруппировка по ID:")
	for handle, group := range userGroups {
		id := handle.Value()
		fmt.Printf("ID %s:%d - %d записей\n", id.Namespace, id.ID, len(group))
		for _, user := range group {
			fmt.Printf("  - %s (%s)\n", user.Name, user.Email)
		}
	}
	fmt.Println()
}

// Пример бенчмарка сравнения производительности
func benchmarkExample() {
	fmt.Println("=== Детальный бенчмарк производительности ===")
	
	// Подготавливаем тестовые данные
	testStrings := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		testStrings[i] = fmt.Sprintf("TestString_%d", i%50) // 50 уникальных строк
	}
	
	// Канонизируем строки
	handles := make([]unique.Handle[string], len(testStrings))
	for i, s := range testStrings {
		handles[i] = unique.Make(s)
	}
	
	// Бенчмарк обычного сравнения строк
	iterations := 10000
	
	start := time.Now()
	stringComparisons := 0
	for iter := 0; iter < iterations; iter++ {
		for i := 0; i < len(testStrings); i++ {
			for j := i + 1; j < len(testStrings); j++ {
				if testStrings[i] == testStrings[j] {
					stringComparisons++
				}
			}
		}
	}
	stringTime := time.Since(start)
	
	// Бенчмарк сравнения handle'ов
	start = time.Now()
	handleComparisons := 0
	for iter := 0; iter < iterations; iter++ {
		for i := 0; i < len(handles); i++ {
			for j := i + 1; j < len(handles); j++ {
				if handles[i] == handles[j] {
					handleComparisons++
				}
			}
		}
	}
	handleTime := time.Since(start)
	
	fmt.Printf("Итераций: %d\n", iterations)
	fmt.Printf("Строк в каждой итерации: %d\n", len(testStrings))
	fmt.Printf("Сравнений в каждой итерации: %d\n", len(testStrings)*(len(testStrings)-1)/2)
	fmt.Printf("\nРезультаты:\n")
	fmt.Printf("Сравнение строк: %v (%d совпадений)\n", stringTime, stringComparisons)
	fmt.Printf("Сравнение handle'ов: %v (%d совпадений)\n", handleTime, handleComparisons)
	fmt.Printf("Ускорение: %.2fx\n", float64(stringTime)/float64(handleTime))
	fmt.Printf("Время на одно сравнение (строки): %.2f нс\n", float64(stringTime)/float64(iterations*len(testStrings)*(len(testStrings)-1)/2))
	fmt.Printf("Время на одно сравнение (handle'ы): %.2f нс\n", float64(handleTime)/float64(iterations*len(testStrings)*(len(testStrings)-1)/2))
	fmt.Println()
}

func runAdvancedExamples() {
	fmt.Println("Продвинутые примеры использования пакета unique")
	fmt.Println("==============================================")
	fmt.Println()
	
	concurrentExample()
	internPoolExample()
	customTypeExample()
	benchmarkExample()
}