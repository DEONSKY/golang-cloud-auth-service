# Authentication Service

## DB Diagram
![image](https://user-images.githubusercontent.com/56198330/215278354-dc0a4268-982e-43f2-93c4-72422f276f3f.png)

```
Table auth.organizations {
  id uuid
  name string
  description string
  createdAt datetime
  updatedAt datetime
  deletedAt datetime
}

Table auth.userMails {
  id uuid
  email string
  createdAt datetime
  updatedAt datetime
  deletedAt datetime
}

Table auth.organizationUsers {
  id uuid
  userId uuid [ref: > auth.userMails.id]
  organizationId uuid [ref: > auth.organizations.id]
  firstName string
  middleName string
  lastName string
  username string
  passwordHash string
  passwordSalt string
  state string // active, pending_invite_confirmation, disabled
  createdAt datetime
  updatedAt datetime
  deletedAt datetime
}

Table auth.groups {
  id uuid
  organizationId uuid [ref: > auth.organizations.id] 
  name string
  description string
}

Table auth.groupPolicies {
  id uuid
  groupId uuid [ref: > auth.groups.id]
  policyId uuid [ref: > auth.policies.id]
  createdAt datetime
  updatedAt datetime
  deletedAt datetime
}

Table auth.userGroups {
  id uuid
  userId uuid [ref: > auth.organizationUsers.id]
  groupId uuid [ref: > auth.groups.id]
  createdAt datetime
  updatedAt datetime
  deletedAt datetime
}

Table auth.policies {
  id uuid
  name string
  organizationId uuid [ref: > auth.organizations.id]
  createdAt datetime
  updatedAt datetime
  deletedAt datetime
}

Table auth.clients {
  id uuid
  organizationId uuid [ref: > auth.organizations.id]
  secret string
}


Ref: "auth"."groups"."id" < "auth"."groups"."description"
```


## Installation

```sh
cp sample.env .env
```

