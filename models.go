package wercker

import "time"

// Build is a detailed api representation
type Build struct {
	ID          string              `json:"id"`
	URL         string              `json:"url"`
	Application *ApplicationSummary `json:"application"`
	Branch      string              `json:"branch"`
	CommitHash  string              `json:"commitHash"`
	CreatedAt   time.Time           `json:"createdAt"`
	EnvVars     []EnvVar            `json:"envVars"`
	FinishedAt  time.Time           `json:"finishedAt"`
	Message     string              `json:"message"`
	Progress    int                 `json:"progress"`
	Result      string              `json:"result"`
	StartedAt   time.Time           `json:"startedAt"`
	Status      string              `json:"status"`
}

// BuildSummary is a summary api representation
type BuildSummary struct {
	ID         string    `json:"id"`
	URL        string    `json:"url"`
	Branch     string    `json:"branch"`
	CommitHash string    `json:"commitHash"`
	CreatedAt  time.Time `json:"createdAt"`
	FinishedAt time.Time `json:"finishedAt"`
	Message    string    `json:"message"`
	Progress   int       `json:"progress"`
	Result     string    `json:"result"`
	StartedAt  time.Time `json:"startedAt"`
	Status     string    `json:"status"`
}

// Deploy is a detailed api representation
type Deploy struct {
	ID          string              `json:"id"`
	URL         string              `json:"url"`
	Status      string              `json:"status"`
	Result      string              `json:"result"`
	CreatedAt   time.Time           `json:"createdAt"`
	FinishedAt  time.Time           `json:"finishedAt"`
	Progress    int                 `json:"progress"`
	Application *ApplicationSummary `json:"application"`
	Build       *BuildSummary       `json:"build"`
}

// DeploySummary is a summary api representation
type DeploySummary struct {
	ID         string    `json:"id"`
	URL        string    `json:"url"`
	Status     string    `json:"status"`
	Result     string    `json:"result"`
	CreatedAt  time.Time `json:"createdAt"`
	FinishedAt time.Time `json:"finishedAt"`
	Progress   int       `json:"progress"`
}

// EnvVar represents a environment variable key value pair
type EnvVar struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// SCM is a detailed source control manager api representation
type SCM struct {
	Type       string `json:"type"`
	Owner      string `json:"owner"`
	Repository string `json:"repository"`
}

// Application is a detailed api representation
type Application struct {
	ID        string       `json:"id"`
	URL       string       `json:"url"`
	Name      string       `json:"name"`
	Owner     *UnifiedUser `json:"owner"`
	Builds    string       `json:"builds"`
	Deploys   string       `json:"deploys"`
	SCM       *SCM         `json:"scm"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
	Privacy   string       `json:"privacy"`
	Stack     int          `json:"stack"`
}

// ApplicationSummary is a summary api representation
type ApplicationSummary struct {
	ID        string       `json:"id"`
	URL       string       `json:"url"`
	Name      string       `json:"name"`
	Owner     *UnifiedUser `json:"owner"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
	Privacy   string       `json:"privacy"`
	Stack     int          `json:"stack"`
}

// Token is a detailed api representation
type Token struct {
	ID             string     `json:"id"`
	URL            string     `json:"url"`
	Name           string     `json:"name"`
	Token          string     `json:"token"`
	HashedToken    string     `json:"hashedToken"`
	LastCharacters string     `json:"lastCharacters"`
	CreatedAt      *time.Time `json:"createdAt"`
	LastUsedAt     *time.Time `json:"lastUsedAt"`
}

// TokenSummary is a summary api representation
type TokenSummary struct {
	ID             string     `json:"id"`
	URL            string     `json:"url"`
	Name           string     `json:"name"`
	HashedToken    string     `json:"hashedToken"`
	LastCharacters string     `json:"lastCharacters"`
	CreatedAt      *time.Time `json:"createdAt"`
	LastUsedAt     *time.Time `json:"lastUsedAt"`
}

// UnifiedUser is a flexible user representation. Not all fields have to be set
type UnifiedUser struct {
	Type   string             `json:"type"`
	Name   string             `json:"name"`
	Avatar *UnifiedUserAvatar `json:"avatar"`
	UserID string             `json:"userId"`
	Meta   *UnifiedUserMeta   `json:"meta"`
}

// UnifiedUserAvatar is the avatar property of the UnifiedUser
type UnifiedUserAvatar struct {
	Gravatar string `json:"gravatar"`
}

// UnifiedUserMeta is the meta property of the UnifiedUser
type UnifiedUserMeta struct {
	Username        string `json:"username"`
	WerckerEmployee bool   `json:"werckerEmployee"`
}

// Run is a detailed api representation
type Run struct {
	ID         string           `json:"id"`
	URL        string           `json:"url"`
	Branch     string           `json:"branch"`
	CommitHash string           `json:"commitHash"`
	CreatedAt  time.Time        `json:"createdAt"`
	EnvVars    []RunEnvVar      `json:"envVars"`
	FinishedAt time.Time        `json:"finishedAt"`
	Message    string           `json:"message"`
	Commits    []string         `json:"commits"`
	Progress   int              `json:"progress"`
	Result     string           `json:"result"`
	StartedAt  time.Time        `json:"startedAt"`
	Status     string           `json:"status"`
	User       *UnifiedUser     `json:"user"`
	Pipeline   *PipelineSummary `json:"pipeline"`
}

// RunEnvVar is the EnvVar property of the Run
type RunEnvVar struct {
}

// RunSummary is a summary api representation
type RunSummary struct {
	ID         string           `json:"id"`
	URL        string           `json:"url"`
	Branch     string           `json:"branch"`
	CommitHash string           `json:"commitHash"`
	CreatedAt  time.Time        `json:"createdAt"`
	FinishedAt time.Time        `json:"finishedAt"`
	Message    string           `json:"message"`
	Progress   int              `json:"progress"`
	Result     string           `json:"result"`
	StartedAt  time.Time        `json:"startedAt"`
	Status     string           `json:"status"`
	User       *UnifiedUser     `json:"user"`
	Pipeline   *PipelineSummary `json:"pipeline"`
}

// PipelineSummary is a summary api representation
type PipelineSummary struct {
	ID                   string    `json:"id"`
	URL                  string    `json:"url"`
	CreatedAt            time.Time `json:"createdAt"`
	Name                 string    `json:"name"`
	Permissions          string    `json:"permissions"`
	PipelineName         string    `json:"pipelineName"`
	SetScmProviderStatus bool      `json:"setScmProviderStatus"`
	Type                 string    `json:"type"`
}

// Pipeline is a detailed api representation
type Pipeline struct {
	ID                   string              `json:"id"`
	URL                  string              `json:"url"`
	CreatedAt            time.Time           `json:"createdAt"`
	Name                 string              `json:"name"`
	Permissions          string              `json:"permissions"`
	PipelineName         string              `json:"pipelineName"`
	SetScmProviderStatus bool                `json:"setScmProviderStatus"`
	Type                 string              `json:"type"`
	AllowedActions       []string            `json:"allowedActions"`
	Application          *ApplicationSummary `json:"application"`
}
