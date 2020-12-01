package handle

import "encoding/json"

const (
	MessageTagOne = "BodyIsNull"
	MessageTagTwo = "ParsingJsonFalse"
)

type message struct {
	code string
	msg string
}

type createPolicy struct {
	PolicyName, StatementsName, PrefixSetName, NeighborSetName, action string
}

type deletePolicy struct {
	PolicyName string
}

type listPolicy struct {
	PolicyName string
}

type addStatementToPolicy struct {
	PolicyName, StatementName string
}

type createStatement struct {
	StatementsName, PrefixSetName, NeighborSetName string
}

type deleteStatement struct {
	StatementsName string
}

type createPrefixSet struct {
	Type, PrefixSetName, ipPrefix, MaskMin, MaskMax string
}

type deletePrefixSet struct {
	PrefixSetName string
}

func Json(code, msg string) ([]byte, error) {
	message, err := json.Marshal(message{code: code, msg: msg})
	if err != nil {
		return nil ,err
	}
	return message, nil
}