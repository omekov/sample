package model

type Error struct {
	Error string `json:"error"`
}

// Credential ...
type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Token ...
type Token struct {
	AccessToken  string `json:"accessToken"`
	Refreshtoken string `json:"refreshToken"`
}
