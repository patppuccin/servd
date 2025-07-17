package models

import "time"

type RestartConfig struct {
	MaxRetries   int           `json:"max_retries" yaml:"max_retries"`
	BackoffDelay time.Duration `json:"backoff_delay" yaml:"backoff_delay"`           // e.g., 10s
	BackoffMax   time.Duration `json:"backoff_max" yaml:"backoff_max"`               // e.g., 1m
	BackoffMult  float64       `json:"backoff_multiplier" yaml:"backoff_multiplier"` // e.g., 2.0
}

type ServiceSpec struct {
	// Identity
	Name        string   `json:"name" yaml:"name"`                 // system-safe identifier
	DisplayName string   `json:"display_name" yaml:"display_name"` // pretty label
	Description string   `json:"description" yaml:"description"`
	Tags        []string `json:"tags" yaml:"tags"`

	// Execution
	Exec     string   `json:"exec" yaml:"exec"`
	Dir      string   `json:"dir" yaml:"dir"`
	Env      []string `json:"env" yaml:"env"` // key=value
	EnvFiles []string `json:"env_files" yaml:"env_files"`

	// Behavior
	StartupType   string         `json:"startup_type" yaml:"startup_type"`     // auto|manual
	RestartPolicy string         `json:"restart_policy" yaml:"restart_policy"` // always|on-failure|unless-stopped
	RestartConfig *RestartConfig `json:"restart_config,omitempty" yaml:"restart_config,omitempty"`
	StartNow      bool           `json:"start_now,omitempty" yaml:"start_now,omitempty"`

	// Internal
	CreatedAt time.Time `json:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `json:"updated_at" yaml:"updated_at"`
}
