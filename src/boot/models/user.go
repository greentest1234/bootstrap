package models

type User struct {
	AuthToken string `json:"auth_token"`
	ID        int    `json:"id"`
	Login     string `json:"login"`
	AvatarURL string `json:"avatar_url"`
	HTMLURL   string `json:"html_url"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	ApiToken  string `json:"api_token"`
}
