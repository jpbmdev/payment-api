package services

import (
	"encoding/json"

	"github.com/tkanos/go-dtree"
)

// -----------------------------------------------
// -- Decision Tree service
// -----------------------------------------------
type DecisionTreeService interface {
	CreateAndExecuteDecisionTree(tree []dtree.Tree, cant float64, amountTotal float64) (string, error)
}

type decisionTreeService struct{}

func NewDecisionTreeService() DecisionTreeService {
	return &decisionTreeService{}
}

func (s *decisionTreeService) CreateAndExecuteDecisionTree(treeStruct []dtree.Tree, cant float64, amountTotal float64) (string, error) {
	//Parse the array of nodes into json
	j, _ := json.Marshal(treeStruct)

	//Create the tree
	tree, err := dtree.LoadTree([]byte(j))

	if err != nil {
		return "", err
	}

	//Generate tree request
	request := make(map[string]interface{})
	//FLOAT
	request["amount_total"] = amountTotal
	request["cant"] = cant

	//Pass the requesto to the tree
	result, err := tree.Resolve(request)

	if err != nil {
		return "", err
	}

	//Return result
	return result.Name, nil

}
