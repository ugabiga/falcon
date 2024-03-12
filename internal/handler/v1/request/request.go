package request

type UpdateUserRequest struct {
	Name     string `json:"name"`
	Timezone string `json:"timezone"`
}
