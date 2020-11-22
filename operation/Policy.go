package operation

import (
	"context"
	"errors"
	gobgpapi "github.com/osrg/gobgp/api"
	"log"
)

func AddStatementToPolicy(ctx context.Context, PolicyName, StatementName string) (bool, error){
	Policy, err := ListPolicies(ctx, PolicyName)
	if err != nil {
		return false, err
	}
	if Policy == nil {
		return false, errors.New("There is not exist this Policy ")
	}
	Statement, err := ListStatements(ctx, StatementName)
	if err != nil {
		return false, err
	}
	if Statement == nil {
		return false, errors.New("There is not exist this Statement ")
	}
	Policy.Statements = append(Policy.Statements, Statement)
	return true, nil
}

func AddPolicies(ctx context.Context, PolicyName, StatementsName, PrefixSetName string) (string, error) {
	var err error
	pas, err := ListPolicies(ctx, PolicyName)
	if err != nil {
		return "false", err
	}
	if pas != nil {
		return "false", errors.New("The policy is exist ")
	}
	has, err := ListStatements(ctx, StatementsName)
	if err != nil {
		return "false", err
	}
	if has != nil {
		return "false", errors.New("The statement is exist ")
	}
	newPolicy := newAddPolicyRequest(PolicyName, StatementsName, PrefixSetName)
	response, err := Client.AddPolicy(ctx, newPolicy)
	if err != nil {
		log.Fatalf("AddPolicy happend a err, err is %s", err)
	}
	return response.String(), nil
}

func DeletePolicies(ctx context.Context, PolicyName string) (string, error) {
	var err error
	has, err := ListPolicies(ctx, PolicyName)
	if err != nil {
		return "", err
	}
	if has == nil {
		return "", errors.New("Policy is not exist ")
	}
	Del := newDelPolicyRequest(has)
	response, err := Client.DeletePolicy(ctx, Del)
	if err != nil {
		log.Fatalf("DeletePolicies happen a err, err is %s", err)
	}
	return response.String(), nil
}

func ListPolicies(ctx context.Context, PolicyName string) (*gobgpapi.Policy, error) {
	ListPolicyClient, err := Client.ListPolicy(ctx, &gobgpapi.ListPolicyRequest{Name: PolicyName})
	if err != nil {
		log.Fatalf("ListPolicies happen a err, err is %s", err)
	}
	PolicyResponse, err := ListPolicyClient.Recv()
	if err != nil {
		log.Fatalf("ListPolicies happen a err, err is %s", err)
	}
	return PolicyResponse.GetPolicy(), nil
}

func SetPolicies() {

}