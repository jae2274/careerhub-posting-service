package dto

type JobPostingRes struct {
	Site        string   `json:"site"`
	PostingId   string   `json:"posting_id"`
	Title       string   `json:"title"`
	CompanyName string   `json:"company_name"`
	Skills      []string `json:"skills"`
	ImageUrl    string   `json:"image_url"`
}
