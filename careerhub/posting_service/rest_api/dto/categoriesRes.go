package dto

type CategoriesRes struct {
	CategoriesBySite []CategoryRes `json:"categoriesBySite"`
}

type CategoryRes struct {
	Site       string   `json:"site"`
	Categories []string `json:"categories"`
}
