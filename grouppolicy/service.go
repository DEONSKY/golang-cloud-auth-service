package grouppolicy

import (
	"fmt"

	"github.com/forfam/authentication-service/customerror"
	"github.com/forfam/authentication-service/group"
	"github.com/forfam/authentication-service/policy"
	"github.com/forfam/authentication-service/postgres"
	reslog "github.com/forfam/authentication-service/utils/resultlogger"
)

func CreatePolicy(data *AddPolicyToGroupPayload) (*GroupPolicyEntity, error) {

	item := GroupPolicyEntity{
		GroupId:  data.GroupId,
		PolicyId: data.PolicyId,
	}

	result := postgres.AuthenticationDb.First(&policy.PolicyEntity{Id: item.PolicyId})

	result, err := reslog.LogGormResult(result, logger, item.PolicyId, reslog.ErrorNotFoundLogMsg, "", "Policy")
	if result == nil {
		return nil, err
	}

	result = postgres.AuthenticationDb.First(&group.GroupEntity{Id: item.GroupId})

	result, err = reslog.LogGormResult(result, logger, item.GroupId, reslog.ErrorNotFoundLogMsg, "", "Group")
	if result == nil {
		return nil, err
	}

	result = postgres.AuthenticationDb.Create(&item)
	result, err = reslog.LogGormResult(result, logger, "undefined", reslog.ErrorDuringCreateLogMsg, "", "Group Policy")
	if result == nil {
		fmt.Println("here")
		return nil, customerror.NewInternalServerError("sadasd", &err, &customerror.Translatable{})
	}

	return &item, nil
}

func deleteGroupPolicyById(id string) (*GroupPolicyEntity, error) {
	item := GroupPolicyEntity{
		Id: id,
	}

	result := postgres.AuthenticationDb.First(&item)

	result, err := reslog.LogGormResult(result, logger, id, reslog.ErrorDuringFindLogMsg, reslog.NotFoundForDeleteLogMsg, "Group Policy")
	if result == nil {
		return nil, err
	}

	result = postgres.AuthenticationDb.Delete(&item)

	result, err = reslog.LogGormResult(result, logger, id, reslog.ErrorDuringDeleteLogMsg, reslog.NotDeletedLogMsg, "Group Policy")
	if result == nil {
		return nil, err
	}

	return &item, nil
}
