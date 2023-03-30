package cars

type Car struct {
	Id    string `json:"id,omitempty"`
	Brand string `json:"brand"`
	Model string `json:"model"`
}
