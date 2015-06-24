package wercker

type Application struct {
}

type Build struct {
	Id         string
	Url        string
	CreatedAt  string
	FinishedAt string
	Progress   float64
	Result     string
	Status     string
	UpdatedAt  string
}

type Deploy struct {
}

type Token struct {
}
