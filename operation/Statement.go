package operation

import (
	"context"
	"errors"
	gobgpapi "github.com/osrg/gobgp/api"
	"log"
)

func AddStatementsToPolicy(ctx context.Context, PolicyName, StatementsName string) (bool, error) {
	policy, err := ListPolicies(ctx, PolicyName)
	if err != nil {
		return false, err
	}
	Statement, err := ListStatements(ctx, StatementsName)
	if err != nil {
		return false, err
	}
	policy.Statements = append(policy.Statements, Statement)
	return true, nil
}

func AddStatements(ctx context.Context, StatementsName, PrefixSetName string) (string, error) {
	var err error
	has, err := ListStatements(ctx, StatementsName)
	if err != nil {
		return "", err
	}
	if has == nil {
		return "", errors.New("The Statement is exist ")
	}
	response, err := Client.AddStatement(ctx, newAddStatementRequest(StatementsName, PrefixSetName))
	if err != nil {
		log.Fatalf("DeleteStatements happen a err, err is %s", err)
	}
	return response.String(), nil
}

func DeleteStatements(ctx context.Context, StatementsName string) (string, error) {
	var err error
	Statements, err := ListStatements(ctx, StatementsName)
	if err != nil {
		return "", err
	}
	if Statements == nil {
		return "", errors.New("There is not exist statement ")
	}
	Del := newDelStatements(Statements)
	response, err := Client.DeleteStatement(ctx, Del)
	if err != nil {
		log.Fatalf("DeleteStatements happen a err, err is %s", err)
	}
	return response.String(), nil
}

func ListStatements(ctx context.Context, StatementsName string) (*gobgpapi.Statement, error) {
	ListStatementClient, err := Client.ListStatement(ctx, &gobgpapi.ListStatementRequest{Name: StatementsName})
	if err != nil {
		log.Fatalf("ListStatements happen a err, err is %s", err)
	}
	StatementResponse, err := ListStatementClient.Recv()
	if err != nil {
		log.Fatalf("ListStatements happen a err, err is %s", err)
	}
	if StatementResponse == nil {
		return nil, errors.New("There is no Statement ")
	}
	return StatementResponse.GetStatement(), nil
}
