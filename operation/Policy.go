package operation

import (
	"context"
	"errors"
	"fmt"
	gobgpapi "github.com/osrg/gobgp/api"
	"log"
)

func AddStatementToPolicy(ctx context.Context, PolicyName, StatementName string) (string, error){
	Policy, err := ListPolicies(ctx, PolicyName)
	if err != nil {
		return "false", fmt.Errorf("ListPolicies is happend a err, err is %s", err)
	}
	if Policy == nil {
		return "false", errors.New("There is not exist this Policy ")
	}
	Statement, err := ListStatements(ctx, StatementName)
	if err != nil {
		return "false", fmt.Errorf("ListStatements is happend a err, err is %s", err)
	}
	if Statement == nil {
		return "false", errors.New("There is not exist this Statement ")
	}
	Policy.Statements = append(Policy.Statements, Statement)
	return "Successful", nil
}

func AddPolicies(ctx context.Context, PolicyName, StatementsName, PrefixSetName, NeighborSetName string) (string, error) {
	var err error
	pas, err := ListPolicies(ctx, PolicyName)
	if err != nil {
		return "false", fmt.Errorf("ListPolicies is happend a err, err is %s", err)
	}
	if pas != nil {
		return "false", errors.New("The policy is exist ")
	}
	has, err := ListStatements(ctx, StatementsName)
	if err != nil {
		return "false", fmt.Errorf("ListStatements is happend a err, err is %s", err)
	}
	if has != nil {
		return "false", errors.New("The statement is exist ")
	}
	newPolicy := newAddPolicyRequest(PolicyName, StatementsName, PrefixSetName, NeighborSetName)
	_, err = Client.AddPolicy(ctx, newPolicy)
	if err != nil {
		return "false", fmt.Errorf("AddPolicy happend a err, err is %s", err)
	}
	return "Successful", nil
}

func DeletePolicies(ctx context.Context, PolicyName string) (string, error) {
	var err error
	has, err := ListPolicies(ctx, PolicyName)
	if err != nil {
		return "false", fmt.Errorf("ListPolicies is happend a err, err is %s", err)
	}
	if has == nil {
		return "false", errors.New("Policy is not exist ")
	}
	Del := newDelPolicyRequest(has)
	_, err = Client.DeletePolicy(ctx, Del)
	if err != nil {
		return "false", fmt.Errorf("DeletePolicies happen a err, err is %s", err)
	}
	return "Successful", nil
}

func ListPolicies(ctx context.Context, PolicyName string) (*gobgpapi.Policy, error) {
	ListPolicyClient, err := Client.ListPolicy(ctx, &gobgpapi.ListPolicyRequest{Name: PolicyName})
	if err != nil {
		log.Fatalf("ListPolicies happen a err, err is %s", err)
	}
	PolicyResponse, err := ListPolicyClient.Recv()
	if err != nil {
		return nil, fmt.Errorf("ListPolicies happen a err, err is %s", err)
	}
	return PolicyResponse.GetPolicy(), nil
}

func SetPolicies() {

}