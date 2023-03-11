package grouppolicy

import (
	"github.com/forfam/authentication-service/genericrepo"
	"github.com/forfam/authentication-service/group"
	"github.com/forfam/authentication-service/policy"
)

func CreatePolicy(data *AddPolicyToGroupPayload) (*GroupPolicyEntity, error) {

	item := GroupPolicyEntity{
		GroupId:  data.GroupId,
		PolicyId: data.PolicyId,
	}

	if err := genericrepo.Take(&policy.PolicyEntity{Id: item.PolicyId}, "Policy", *logger); err != nil {
		return nil, err
	}

	if err := genericrepo.Take(&group.GroupEntity{Id: item.GroupId}, "Group", *logger); err != nil {
		return nil, err
	}

	if err := genericrepo.IsRelationNotExists(&item, []string{"Policy", "Group"}, *logger); err != nil {
		return nil, err
	}

	if err := genericrepo.Create(&item, "Group Policy", *logger); err != nil {
		return nil, err
	}

	return &item, nil
}

func DeleteGroupPolicyById(id string) (*GroupPolicyEntity, error) {
	item := GroupPolicyEntity{
		Id: id,
	}

	if err := genericrepo.Take(&item, "Group Policy", *logger); err != nil {
		return nil, err
	}

	if err := genericrepo.Delete(&item, "Group Policy", *logger); err != nil {
		return nil, err
	}

	return &item, nil
}
