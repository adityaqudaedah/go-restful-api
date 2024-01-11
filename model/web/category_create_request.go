package web

type CategoryCreateRequest struct {
	Name string `validate:"min=2,max=200,required" json:"name"`
}
