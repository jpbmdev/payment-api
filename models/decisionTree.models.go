package models

// -----------------------------------------------
// -- Decision Tree models & dtos
// -----------------------------------------------

type DecisionTreeInputs struct {
	Cant        float64 `json:"cant" binding:"number"`
	AmountTotal float64 `json:"amountTotal" binding:"number"`
}

type Tree struct {
	ID       int                    `json:"id"`
	Name     string                 `json:"name"`
	ParentID int                    `json:"parentId"`
	Value    interface{}            `json:"value"`
	Operator string                 `json:"operator"`
	Key      string                 `json:"key"`
	Order    int                    `json:"order"`
	Content  interface{}            `json:"content"`
	Headers  map[string]interface{} `json:"headers"`
}
