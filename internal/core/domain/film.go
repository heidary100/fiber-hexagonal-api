package domain

type Film struct {
	Name string
}

type FilmSearchResult struct {
	Id               int
	Title            string
	OriginalTitle    string
	VoteAverage      float64
	VoteCount        int64
	PosterPath       string
	Popularity       float64
	OriginalLanguage string
	ReleaseDate      string
	Overview         string
}
