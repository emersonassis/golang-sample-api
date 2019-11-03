package web

//Object ...
type Object struct {
	ID   int    `json:"code"`
	Name string `json:"name"`
}

//ObjectRequest ...
type ObjectRequest struct {
	Object
}

//ObjectResponse ...
type ObjectResponse struct {
	ResponseBodyJSONDefault
	Object *Object `json:"object"`
}

//ObjectsResponse ...
type ObjectsResponse struct {
	ResponseBodyJSONDefault
	Objects []*Object `json:"objects"`
}
