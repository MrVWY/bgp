package operation

import (
	"context"
	"errors"
	"fmt"
	gobgpapi "github.com/osrg/gobgp/api"
)

func AddStatementsToPolicy(ctx context.Context, PolicyName, StatementsName string) (string, error) {
	policy, err := ListPolicies(ctx, PolicyName)
	if err != nil {
		return "false", fmt.Errorf("ListPolicies is happend a err, err is %s", err)
	}
	Statement, err := ListStatements(ctx, StatementsName)
	if err != nil {
		return "false", fmt.Errorf("ListStatements is happend a err, err is %s", err)
	}
	policy.Statements = append(policy.Statements, Statement)
	return "Successful", nil
}

func AddStatements(ctx context.Context, StatementsName, PrefixSetName, NeighborSetName, CommunitySetName, CommunityAction, action string, Community []string) (string, error) {
	var err error
	has, err := ListStatements(ctx, StatementsName)
	if err != nil {
		return "", fmt.Errorf("ListStatements is happend a err, err is %s", err)
	}
	if has == nil {
		return "", errors.New("The Statement is exist ")
	}
	_, err = Client.AddStatement(ctx, newAddStatementRequest(StatementsName, PrefixSetName, NeighborSetName, CommunitySetName, CommunityAction, action, Community))
	if err != nil {
		return "false", fmt.Errorf("DeleteStatements happen a err, err is %s", err)
	}
	return "Successful", nil
}

func DeleteStatements(ctx context.Context, StatementsName string) (string, error) {
	var err error
	Statements, err := ListStatements(ctx, StatementsName)
	if err != nil {
		return "", fmt.Errorf("ListStatements is happend a err, err is %s", err)
	}
	if Statements == nil {
		return "", errors.New("There is not exist statement ")
	}
	Del := newDelStatements(Statements)
	_, err = Client.DeleteStatement(ctx, Del)
	if err != nil {
		return "false", fmt.Errorf("DeleteStatements happen a err, err is %s", err)
	}
	return "Successful", nil
}

func ListStatements(ctx context.Context, StatementsName string) (*gobgpapi.Statement, error) {
	ListStatementClient, err := Client.ListStatement(ctx, &gobgpapi.ListStatementRequest{Name: StatementsName})
	if err != nil {
		return nil, fmt.Errorf("ListStatements happen a err, err is %s", err)
	}
	StatementResponse, err := ListStatementClient.Recv()
	if err != nil {
		return nil, fmt.Errorf("ListStatements happen a err, err is %s", err)
	}
	if StatementResponse == nil {
		return nil, errors.New("There is no Statement ")
	}
	return StatementResponse.GetStatement(), nil
}
