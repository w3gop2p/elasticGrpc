package domain

type Adv struct {
	ID         string   `json:"id"`
	Categories Category `json:"categories"`
	Title      Title    `json:"title"`
	Type       string   `json:"type"`
	Posted     float64  `json:"posted"`
}

type Category struct {
	Subcategory string `json:"subcategory"`
}

type Title struct {
	Ro string `json:"ro"`
	Ru string `json:"ru"`
}

func NewAdv(id string, subcategory string, titleRo string, titleRu string, adType string, posted float64) Adv {
	return Adv{
		ID: id,
		Categories: Category{
			Subcategory: subcategory,
		},
		Title: Title{
			Ro: titleRo,
			Ru: titleRu,
		},
		Type:   adType,
		Posted: posted,
	}
}
