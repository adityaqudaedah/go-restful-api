package web

type CategoryCreateRequest struct {
	Name string `validate:"min=1,max=200,required" json:"name"`
}
