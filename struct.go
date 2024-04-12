package dbkit

type (
	QueryInfo struct {
		Query string `json:"query" yml:"query" xml:"query"`
		Args  []any  `json:"args" yml:"args" xml:"args"`
	}

	ResponseQuery struct {
		Success   bool                     `json:"success" xml:"success"`
		Message   string                   `json:"message" xml:"message"`
		ItemFound bool                     `json:"itemFound" xml:"itemFound"`
		Items     []map[string]interface{} `json:"items" xml:"items"`
	}

	ArgsUpdate struct {
		Columns []any
		Where   []any
	}

	UpdateInfo struct {
		Script string
		Args   ArgsUpdate
	}
)
