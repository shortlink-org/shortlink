/**
 * Инициализация proto3.util для @bufbuild/protobuf
 * Этот файл должен быть загружен ПЕРВЫМ перед любыми proto файлами
 * 
 * Workaround для проблемы с @bufbuild/protobuf в Vitest:
 * proto3.util не инициализируется автоматически в некоторых окружениях
 */
// proto3 is not exported from @bufbuild/protobuf v2
// The runtime is initialized automatically when proto files are imported

// In @bufbuild/protobuf v2, the runtime is initialized automatically
// when proto files are imported. No manual initialization needed.

