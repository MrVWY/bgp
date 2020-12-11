package handle

import "encoding/json"

const (
	MessageTagOne = "BodyIsNull"
	MessageTagTwo = "ParsingJsonFalse"
)

type message struct {
	code string
	msg  string
}

type policy struct {
	PolicyName string
}

type createPolicy struct {
	PolicyName, StatementsName, PrefixSetName, NeighborSetName, CommunitySetName, CommunityAction, action string
	Community []string
}

type addStatementToPolicy struct {
	PolicyName, StatementName string
}

type createStatement struct {
	StatementsName, PrefixSetName, NeighborSetName, CommunitySetName, CommunityAction, action string
	Community []string
}

type Statements struct {
	StatementsName string
}

type createPrefixSet struct {
	Type, PrefixSetName, ipPrefix, MaskMin, MaskMax string
}

type createCommunitySet struct {
	CommunitySetName, Type string
	list []string
}

type createNeighborSet struct {
	NeighborSetName, Type string
	list []string
}

type deleteDefinedSet struct {
	DefinedSetName string
}

func Json(code, msg string) ([]byte, error) {
	message, err := json.Marshal(message{code: code, msg: msg})
	if err != nil {
		return nil, err
	}
	return message, nil
}
