# Authentication Service

## DB Diagram
![image](https://user-images.githubusercontent.com/56198330/215271917-ef87085c-3266-4f67-87c0-19862e194fcf.png)
```
Table auth.users {
  id uuid
  organizationId uuid [ref: > auth.organizations.id]
  firstName string
  middleName string
  lastName string
  email string
  username string
  passwordHash string
  passwordSalt string
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
  userId uuid [ref: > auth.users.id]
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

Table auth.organizations {
  id uuid
  name string
  description string
  //  ownerId uuid [ref: > auth.users.id]
  subscription string // it would be enum. I don't think about it right now.
  createdAt datetime
  updatedAt datetime
  deletedAt datetime
}

```


## Installation

```sh
cp sample.env .env
```

