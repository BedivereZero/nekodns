package v1alpha1

type Record struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type RecordList struct {
	Items []Record `json:"items"`
}
