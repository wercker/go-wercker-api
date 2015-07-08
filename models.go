package wercker

import "time"

type Build struct {
	ID          string              `json:"id"`
	URL         string              `json:"url"`
	Application *ApplicationSummary `json:"application"`
	Branch      string              `json:"branch"`
	CommitHash  string              `json:"commitHash"`
	CreatedAt   time.Time           `json:"createdAt"`
	FinishedAt  time.Time           `json:"finishedAt"`
	Message     string              `json:"message"`
	Progress    int                 `json:"progress"`
	Result      string              `json:"result"`
	StartedAt   time.Time           `json:"startedAt"`
	Status      string              `json:"status"`
}

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

type DeploySummary struct {
	ID         string    `json:"id"`
	URL        string    `json:"url"`
	Status     string    `json:"status"`
	Result     string    `json:"result"`
	CreatedAt  time.Time `json:"createdAt"`
	FinishedAt time.Time `json:"finishedAt"`
	Progress   int       `json:"progress"`
}

type SCM struct {
	Type       string `json:"type"`
	Owner      string `json:"owner"`
	Repository string `json:"repository"`
}

type Application struct {
	ID        string       `json:"id"`
	URL       string       `json:"url"`
	Name      string       `json:"name"`
	Owner     *UnifiedUser `json:"owner"`
	Builds    string       `json:"builds"`
	Deploys   string       `json:"deploys"`
	SCM       *SCM         `json:"scm"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt string       `json:"updatedAt"`
	Privacy   string       `json:"privacy"`
	Stack     int          `json:"stack"`
}

type ApplicationSummary struct {
	ID        string       `json:"id"`
	URL       string       `json:"url"`
	Name      string       `json:"name"`
	Owner     *UnifiedUser `json:"owner"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt string       `json:"updatedAt"`
	Privacy   string       `json:"privacy"`
	Stack     int          `json:"stack"`
}

type Token struct {
	ID             string    `json:"id"`
	URL            string    `json:"url"`
	Name           string    `json:"name"`
	Token          string    `json:"token"`
	HashedToken    string    `json:"hashedToken"`
	LastCharacters string    `json:"lastCharacters"`
	CreatedAt      time.Time `json:"createdAt"`
	LastUsedAt     time.Time `json:"lastUsedAt"`
}

type TokenSummary struct {
	ID             string    `json:"id"`
	URL            string    `json:"url"`
	Name           string    `json:"name"`
	HashedToken    string    `json:"hashedToken"`
	LastCharacters string    `json:"lastCharacters"`
	CreatedAt      time.Time `json:"createdAt"`
	LastUsedAt     time.Time `json:"lastUsedAt"`
}

type UnifiedUser struct {
	Type   string             `json:"type"`
	Name   string             `json:"name"`
	Avatar *UnifiedUserAvatar `json:"avatar"`
	UserID string             `json:"userId"`
	Meta   *UnifiedUserMeta   `json:"meta"`
}

type UnifiedUserAvatar struct {
	Gravatar string `json:"gravatar"`
}

type UnifiedUserMeta struct {
	Username        string `json:"username"`
	WerckerEmployee bool   `json:"werckerEmployee"`
}
