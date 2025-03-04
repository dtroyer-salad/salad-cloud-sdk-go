package inferenceendpoints

import "encoding/json"

// Represents a page from the collection of inference endpoints.
type InferenceEndpointList struct {
	// The list of inference endpoints.
	Items []InferenceEndpoint `json:"items,omitempty" required:"true" maxItems:"100"`
	// The page number.
	Page *int64 `json:"page,omitempty" required:"true" min:"1" max:"2147483647"`
	// The maximum number of items per page.
	PageSize *int64 `json:"page_size,omitempty" required:"true" min:"1" max:"100"`
	// The total number of items in the collection.
	TotalSize *int64 `json:"total_size,omitempty" required:"true" min:"0" max:"2147483647"`
}

func (i *InferenceEndpointList) GetItems() []InferenceEndpoint {
	if i == nil {
		return nil
	}
	return i.Items
}

func (i *InferenceEndpointList) SetItems(items []InferenceEndpoint) {
	i.Items = items
}

func (i *InferenceEndpointList) GetPage() *int64 {
	if i == nil {
		return nil
	}
	return i.Page
}

func (i *InferenceEndpointList) SetPage(page int64) {
	i.Page = &page
}

func (i *InferenceEndpointList) GetPageSize() *int64 {
	if i == nil {
		return nil
	}
	return i.PageSize
}

func (i *InferenceEndpointList) SetPageSize(pageSize int64) {
	i.PageSize = &pageSize
}

func (i *InferenceEndpointList) GetTotalSize() *int64 {
	if i == nil {
		return nil
	}
	return i.TotalSize
}

func (i *InferenceEndpointList) SetTotalSize(totalSize int64) {
	i.TotalSize = &totalSize
}

func (i InferenceEndpointList) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InferenceEndpointList to string"
	}
	return string(jsonData)
}
