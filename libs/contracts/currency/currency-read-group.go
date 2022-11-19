package currency

type ReadGroupsRequest struct{}

type Group struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

type ReadGroupsResponse []Group
