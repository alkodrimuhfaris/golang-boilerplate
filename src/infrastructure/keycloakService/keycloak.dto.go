package keycloakService

type Credential struct {
	Password     string `form:"password"`
	Username     string `form:"username"`
	ClientID     string `form:"client_id"`
	GrantType    string `form:"grant_type"`
	Scope        string `form:"scope"`
	ClientSecret string `form:"client_secret"`
}

type Response struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	IDToken      string `json:"idToken"`
}
