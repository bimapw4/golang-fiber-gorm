package types

type WorksDTO struct {
	Work        string `json:"work"`
	Description string `json:"description"`
}

type WorksUpdateDTO struct {
	Work string `json:"work"`
}

type WorksListDB struct {
	ID          int    `json:"id"`
	Work        string `json:"work"`
	Description string `json:"description"`
}

type WorksListResp struct {
	Status bool          `json:"status"`
	Data   []WorksListDB `json:"data"`
}
