package common

type Artist struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Popularity  int      `json:"popularity"`
	Recommended []string `json:"recommended"`
}
