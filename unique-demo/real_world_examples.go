package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"unique"
)

// Пример 1: Система разрешений с канонизацией ролей
type Permission struct {
	Resource string
	Action   string
}

type Role struct {
	Name string
	// Используем строку для упрощения - в реальности можно использовать другие подходы
	PermissionKeys string
}

type User struct {
	ID       int64
	Name     string
	RoleKeys []string // Упрощаем для демонстрации
}

func rbacExample() {
	fmt.Println("=== Система RBAC с канонизацией ===")
	
	// Создаем базовые разрешения
	readUsers := unique.Make(Permission{"users", "read"})
	readPosts := unique.Make(Permission{"posts", "read"})
	writePosts := unique.Make(Permission{"posts", "write"})
	adminAll := unique.Make(Permission{"*", "*"})
	
	// Создаем роли
	viewer := unique.Make(Role{
		Name:           "viewer",
		PermissionKeys: "users:read,posts:read",
	})
	
	editor := unique.Make(Role{
		Name:           "editor", 
		PermissionKeys: "users:read,posts:read,posts:write",
	})
	
	admin := unique.Make(Role{
		Name:           "admin",
		PermissionKeys: "*:*",
	})
	
	// Карта разрешений для демонстрации
	permissions := map[string]unique.Handle[Permission]{
		"users:read":  readUsers,
		"posts:read":  readPosts,
		"posts:write": writePosts,
		"*:*":         adminAll,
	}
	
	// Создаем пользователей
	users := []User{
		{ID: 1, Name: "Иван", RoleKeys: []string{"viewer"}},
		{ID: 2, Name: "Мария", RoleKeys: []string{"editor", "viewer"}}, // дублирование ролей
		{ID: 3, Name: "Админ", RoleKeys: []string{"admin"}},
		{ID: 4, Name: "Анна", RoleKeys: []string{"viewer"}}, // та же роль что у Ивана
	}
	
	// Карта ролей для поиска
	roles := map[string]unique.Handle[Role]{
		"viewer": viewer,
		"editor": editor,
		"admin":  admin,
	}
	
	// Функция проверки разрешений
	hasPermission := func(user User, resource, action string) bool {
		for _, roleKey := range user.RoleKeys {
			if roleHandle, exists := roles[roleKey]; exists {
				role := roleHandle.Value()
				// Упрощенная проверка разрешений
				if role.PermissionKeys == "*:*" {
					return true
				}
				permKey := resource + ":" + action
				if strings.Contains(role.PermissionKeys, permKey) {
					return true
				}
			}
		}
		return false
	}
	
	// Тестируем разрешения
	testCases := []struct {
		user     User
		resource string
		action   string
	}{
		{users[0], "users", "read"},   // Иван читает пользователей
		{users[0], "users", "write"},  // Иван пишет пользователей
		{users[1], "posts", "write"},  // Мария пишет посты
		{users[2], "anything", "do"},  // Админ делает что угодно
	}
	
	fmt.Println("Проверка разрешений:")
	for _, tc := range testCases {
		result := hasPermission(tc.user, tc.resource, tc.action)
		fmt.Printf("%s может %s %s: %v\n", tc.user.Name, tc.action, tc.resource, result)
	}
	
	// Статистика канонизации
	roleMap := make(map[unique.Handle[Role]]int)
	for _, user := range users {
		for _, roleKey := range user.RoleKeys {
			if roleHandle, exists := roles[roleKey]; exists {
				roleMap[roleHandle]++
			}
		}
	}
	
	fmt.Printf("\nСтатистика ролей (уникальные handle'ы):\n")
	for roleHandle, count := range roleMap {
		fmt.Printf("Роль '%s': используется %d раз\n", roleHandle.Value().Name, count)
	}
	
	// Предотвращаем предупреждение о неиспользуемой переменной
	_ = permissions
	
	fmt.Println()
}

// Пример 2: Система конфигурации с интернированием строк
type DatabaseConfig struct {
	Host     unique.Handle[string]
	Port     int
	Database unique.Handle[string]
	Username unique.Handle[string]
}

type ServiceConfig struct {
	Name     unique.Handle[string]
	Database DatabaseConfig
	LogLevel unique.Handle[string]
}

func configExample() {
	fmt.Println("=== Система конфигурации с интернированием ===")
	
	// Часто используемые значения
	localhost := unique.Make("localhost")
	prodDB := unique.Make("production_db")
	devDB := unique.Make("development_db")
	apiUser := unique.Make("api_user")
	adminUser := unique.Make("admin_user")
	infoLevel := unique.Make("info")
	debugLevel := unique.Make("debug")
	
	// Конфигурации сервисов
	services := []ServiceConfig{
		{
			Name: unique.Make("user-service"),
			Database: DatabaseConfig{
				Host: localhost, Port: 5432, Database: prodDB, Username: apiUser,
			},
			LogLevel: infoLevel,
		},
		{
			Name: unique.Make("order-service"),
			Database: DatabaseConfig{
				Host: localhost, Port: 5432, Database: prodDB, Username: apiUser, // те же значения
			},
			LogLevel: infoLevel,
		},
		{
			Name: unique.Make("notification-service"),
			Database: DatabaseConfig{
				Host: localhost, Port: 5432, Database: devDB, Username: adminUser,
			},
			LogLevel: debugLevel,
		},
	}
	
	// Анализируем повторяющиеся конфигурации
	hostUsage := make(map[unique.Handle[string]][]string)
	dbUsage := make(map[unique.Handle[string]][]string)
	userUsage := make(map[unique.Handle[string]][]string)
	
	for _, service := range services {
		serviceName := service.Name.Value()
		hostUsage[service.Database.Host] = append(hostUsage[service.Database.Host], serviceName)
		dbUsage[service.Database.Database] = append(dbUsage[service.Database.Database], serviceName)
		userUsage[service.Database.Username] = append(userUsage[service.Database.Username], serviceName)
	}
	
	fmt.Println("Анализ повторяющихся конфигураций:")
	
	fmt.Println("\nХосты БД:")
	for handle, services := range hostUsage {
		fmt.Printf("  %s: %s\n", handle.Value(), strings.Join(services, ", "))
	}
	
	fmt.Println("\nБазы данных:")
	for handle, services := range dbUsage {
		fmt.Printf("  %s: %s\n", handle.Value(), strings.Join(services, ", "))
	}
	
	fmt.Println("\nПользователи БД:")
	for handle, services := range userUsage {
		fmt.Printf("  %s: %s\n", handle.Value(), strings.Join(services, ", "))
	}
	fmt.Println()
}

// Пример 3: Парсер JSON с дедупликацией ключей
type JSONKey struct {
	Path string
	Key  string
}

type JSONAnalyzer struct {
	keys   map[unique.Handle[JSONKey]]int
	values map[unique.Handle[string]]int
}

func NewJSONAnalyzer() *JSONAnalyzer {
	return &JSONAnalyzer{
		keys:   make(map[unique.Handle[JSONKey]]int),
		values: make(map[unique.Handle[string]]int),
	}
}

func (ja *JSONAnalyzer) analyzeValue(value interface{}, path string) {
	switch v := value.(type) {
	case string:
		handle := unique.Make(v)
		ja.values[handle]++
	case map[string]interface{}:
		for key, val := range v {
			keyHandle := unique.Make(JSONKey{Path: path, Key: key})
			ja.keys[keyHandle]++
			newPath := path + "." + key
			if path == "" {
				newPath = key
			}
			ja.analyzeValue(val, newPath)
		}
	case []interface{}:
		for i, val := range v {
			arrayPath := fmt.Sprintf("%s[%d]", path, i)
			ja.analyzeValue(val, arrayPath)
		}
	}
}

func (ja *JSONAnalyzer) Analyze(jsonData string) error {
	var data interface{}
	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		return err
	}
	
	ja.analyzeValue(data, "")
	return nil
}

func (ja *JSONAnalyzer) Report() {
	fmt.Println("Статистика ключей JSON:")
	for keyHandle, count := range ja.keys {
		key := keyHandle.Value()
		fmt.Printf("  %s.%s: %d раз\n", key.Path, key.Key, count)
	}
	
	fmt.Println("\nСтатистика строковых значений:")
	for valueHandle, count := range ja.values {
		if count > 1 { // показываем только повторяющиеся
			fmt.Printf("  '%s': %d раз\n", valueHandle.Value(), count)
		}
	}
}

func jsonAnalysisExample() {
	fmt.Println("=== Анализ JSON с дедупликацией ===")
	
	// Пример JSON с повторяющимися структурами
	sampleJSON := `{
		"users": [
			{
				"id": 1,
				"name": "Алексей",
				"status": "active",
				"role": "admin"
			},
			{
				"id": 2,
				"name": "Мария", 
				"status": "active",
				"role": "user"
			},
			{
				"id": 3,
				"name": "Иван",
				"status": "inactive",
				"role": "user"
			}
		],
		"settings": {
			"theme": "dark",
			"language": "ru",
			"notifications": {
				"email": "active",
				"push": "active"
			}
		},
		"metadata": {
			"version": "1.0",
			"language": "ru"
		}
	}`
	
	analyzer := NewJSONAnalyzer()
	if err := analyzer.Analyze(sampleJSON); err != nil {
		fmt.Printf("Ошибка анализа JSON: %v\n", err)
		return
	}
	
	analyzer.Report()
	fmt.Println()
}

func runRealWorldExamples() {
	fmt.Println("Практические примеры использования unique")
	fmt.Println("=========================================")
	fmt.Println()
	
	rbacExample()
	configExample()
	jsonAnalysisExample()
}