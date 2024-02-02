package field

// Fields Type to pass when we want to call WithFields for structured logging
// TODO: maybe we can use zap.Field instead of map[string]any
type Fields map[string]any
