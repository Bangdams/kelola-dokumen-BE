package model

type RwListResponse struct {
	ID     uint   `json:"id" validate:"required"`
	UserId uint   `json:"user_id" validate:"required"`
	NameRw string `json:"name_rw" validate:"required"`
}

type RwListRequest struct {
	NameRw string `json:"name_rw"`
}

type RwListUpdateRequest struct {
	RwListID uint   `json:"rw_list_id"`
	UserId   uint   `json:"user_id" validate:"required"`
	NameRw   string `json:"name_rw" validate:"required"`
}
