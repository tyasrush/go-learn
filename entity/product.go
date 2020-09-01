package entity

type Product struct {
	ID   int64 `json:"id"`
	Name int64 `json:"name"`
	Desc int64 `json:"description"`
	Typ  int64 `json:"type"`
}

type ProductPostRequest struct {
	Name int64 `json:"name"`
	Desc int64 `json:"description"`
	Typ  int64 `json:"type"`
}

type ProductGetRequest struct {
	ID int64 `json:"id"`
}

type ProductGetByParamRequest struct {
	Name int64 `json:"name"`
	Typ  int64 `json:"type"`
}
