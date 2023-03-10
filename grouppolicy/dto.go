package grouppolicy

import "time"

type AddPolicyToGroupPayload struct {
	GroupId  string `json:"groupId" validate:"required,uuid"`
	PolicyId string `json:"policyId" validate:"required,uuid"`
}

type GroupPolicyResponse struct {
	Id        string `json:"id"`
	PolicyId  string `json:"policyId"`
	GroupId   string `json:"groupId"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (ent GroupPolicyEntity) mapEntity() GroupPolicyResponse {
	return GroupPolicyResponse{
		Id:        ent.Id,
		PolicyId:  ent.PolicyId,
		GroupId:   ent.GroupId,
		CreatedAt: ent.CreatedAt,
		UpdatedAt: ent.UpdatedAt,
	}
}
