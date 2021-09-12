package types

type ImageContent struct {
	Photos Photos `json:"photos"`
}

type Photos struct {
	Total int            `json:"total"`
	Photo []PhotoContent `json:"photo"`
}
type PhotoContent struct {
	ID     string `json:"id"`
	Secret string `json:"secret"`
	Server string `json:"server"`
}
