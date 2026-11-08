package links

type Link struct {
	Code string `json:"code"`
	URL  string `json:"url"`
	Hits int    `json:"hits"`
}
