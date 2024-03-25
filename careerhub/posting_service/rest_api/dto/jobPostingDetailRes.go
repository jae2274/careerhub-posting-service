package dto

type JobPostingDetailRes struct {
	Site           string   `json:"site"`
	PostingId      string   `json:"postingId"`
	Title          string   `json:"title"`
	Skills         []string `json:"skills"`
	Intro          string   `json:"intro"`
	MainTask       string   `json:"mainTask"`
	Qualifications string   `json:"qualifications"`
	Preferred      string   `json:"preferred"`
	Benefits       string   `json:"benefits"`
	RecruitProcess *string  `json:"recruitProcess"`
	CareerMin      *int32   `json:"careerMin"`
	CareerMax      *int32   `json:"careerMax"`
	Addresses      []string `json:"addresses"`
	CompanyId      string   `json:"companyId"`
	CompanyName    string   `json:"companyName"`
	CompanyImages  []string `json:"companyImages"`
	Tags           []string `json:"tags"`
}
