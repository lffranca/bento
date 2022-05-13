package entity

type UserAuthentication struct {
	User         User   `json:"user" yaml:"user"`
	IDToken      string `json:"id_token" yaml:"id_token"`
	AccessToken  string `json:"access_token" yaml:"access_token"`
	RefreshToken string `json:"refresh_token" yaml:"refresh_token"`
}
