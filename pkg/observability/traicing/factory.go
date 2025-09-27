// Package traicing provides a factory for creating FlightRecorder instances.
// This layer handles dependency injection and component wiring using clean architecture principles.
package traicing

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/shortlink-org/go-sdk/logger"
	"github.com/shortlink-org/shortlink/pkg/observability/traicing/application"
	"github.com/shortlink-org/shortlink/pkg/observability/traicing/domain"
	"github.com/shortlink-org/shortlink/pkg/observability/traicing/infrastructure"
)

// Factory creates and configures FlightRecorder components.
// It follows the Factory pattern to encapsulate complex object creation.
type Factory struct {
	config *FactoryConfig
}

// FactoryConfig contains configuration for the FlightRecorder factory.
// This configuration determines which implementations to use for each layer.
type FactoryConfig struct {
	// Recorder configuration
	Enabled  bool
	MinAge   time.Duration
	MaxBytes uint64

	// Storage configuration
	StorageType     StorageType
	StorageBasePath string
	StorageMaxAge   time.Duration

	// Event handling configuration
	Logger logger.Logger
}

// StorageType defines the available storage backend implementations.
type StorageType string

const (
	// StorageTypeFileSystem uses local filesystem storage.
	StorageTypeFileSystem StorageType = "filesystem"
	// StorageTypeMemory uses in-memory storage (for testing).
	StorageTypeMemory StorageType = "memory"
)

// NewFactory creates a new FlightRecorder factory with the provided configuration.
// It validates the configuration and sets appropriate defaults.
func NewFactory(config *FactoryConfig) (*Factory, error) {
	if config == nil {
		return nil, fmt.Errorf("factory configuration is required")
	}

	// Set defaults
	if config.StorageType == "" {
		config.StorageType = StorageTypeFileSystem
	}

	if config.StorageBasePath == "" {
		homeDir, _ := os.UserHomeDir()
		config.StorageBasePath = filepath.Join(homeDir, ".shortlink", "traces")
	}

	if config.StorageMaxAge == 0 {
		config.StorageMaxAge = 24 * time.Hour // Default: keep traces for 24 hours
	}

	if config.Logger == nil {
		return nil, fmt.Errorf("logger is required")
	}

	return &Factory{
		config: config,
	}, nil
}

// CreateRecorderService creates a fully configured RecorderService.
// This method wires together all the dependencies using dependency injection.
func (f *Factory) CreateRecorderService() (*application.RecorderService, error) {
	// Create domain configuration
	domainConfig, err := domain.NewConfiguration(
		f.config.Enabled,
		f.config.MinAge,
		f.config.MaxBytes,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create domain configuration: %w", err)
	}

	// Create recorder implementation
	recorder, err := f.createRecorder(domainConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create recorder: %w", err)
	}

	// Create repository implementation
	repository, err := f.createRepository()
	if err != nil {
		return nil, fmt.Errorf("failed to create repository: %w", err)
	}

	// Create event handler implementation
	eventHandler := f.createEventHandler()

	// Create and return the service
	service := application.NewRecorderService(
		recorder,
		repository,
		eventHandler,
		domainConfig,
	)

	return service, nil
}

// createRecorder creates the appropriate recorder implementation.
func (f *Factory) createRecorder(config *domain.Configuration) (domain.Recorder, error) {
	return infrastructure.NewGoFlightRecorder(config)
}

// createRepository creates the appropriate repository implementation.
func (f *Factory) createRepository() (domain.Repository, error) {
	switch f.config.StorageType {
	case StorageTypeFileSystem:
		return infrastructure.NewFileSystemRepository(
			f.config.StorageBasePath,
			f.config.StorageMaxAge,
		)
	case StorageTypeMemory:
		// For testing purposes - you would implement this
		return nil, fmt.Errorf("memory storage not implemented yet")
	default:
		return nil, fmt.Errorf("unsupported storage type: %s", f.config.StorageType)
	}
}

// createEventHandler creates the appropriate event handler implementation.
func (f *Factory) createEventHandler() domain.EventHandler {
	loggingHandler := infrastructure.NewLoggingEventHandler(f.config.Logger)
	metricsHandler := infrastructure.NewMetricsEventHandler()

	return infrastructure.NewCompositeEventHandler(
		loggingHandler,
		metricsHandler,
	)
}

// DefaultFactoryConfig creates a default factory configuration.
// This provides sensible defaults for most use cases.
func DefaultFactoryConfig(log logger.Logger) *FactoryConfig {
	return &FactoryConfig{
		Enabled:         true,
		MinAge:          1 * time.Minute,
		MaxBytes:        3 << 20, // 3MB
		StorageType:     StorageTypeFileSystem,
		StorageBasePath: "", // Will use default
		StorageMaxAge:   24 * time.Hour,
		Logger:          log,
	}
}