## Error struct
Custom errors structure and commonly used errors are created

customerror/customerror.go
```go
type CoreError struct {
	Err         error
	Description string
	CreatedAt   time.Time
	CauseError   error
	HttpCode    int
	I18n        *i18n.Translatable
}
```
This errors used inside service logic. But we may dont want to expose more details about error to client. For this reason I convert errors before sending client with fiber error handler. 
I am controling error type inside this handler. Assertions in go is pretty fast. We dont have extend for go. For this reason I defined error types inside code (like ```BadRequestError```) and automaticly using this error as Err parameter inside constructor of any type of ```CoreError```. This provides a feature like extending. All sub error types has coreError fields. If we want to add more specific details about error, we can add ```CauseError``` for that. ```CauseError``` can be gorm errors, validation errors, etc. We can add more information about error with this way. 

## Error Handling
And we can define which field will send to client by error strut to dto convertion functions. We can call different convertions for different ```CoreError``` types. If there is not custom convertion method for ```CoreError sub type```, default ```CoreError``` convertion runs, and hides unecessary/private data from client.

config/fiber.go
```go
		if re, ok := err.(*customerror.CoreError); ok {

			if errors.Is(re.Err, customerror.ValidationError) {
				return ctx.Status(re.HttpCode).JSON(re.MapToValidationErrorResponse())
			}

			return ctx.Status(re.HttpCode).JSON(re.MapToCoreErrorResponse())
		}
```

## We Made Code Much More Clean With Generic Repository
This type of errors should be created for every possible error in service logic. And this can be repetative. Also we need to log information about every error. For this reason I created a ```generic wrapper``` which automaticly creates common logs and CoreError's for common gorm functions. This wrapper sees repository duty. For this reason I named as ```genericrepo```. This is an example function:

genericrepo/genericrepo.go
```go
func Take[T Entity](item *T, targetName string, logger log.Logger) error {

	result := postgres.AuthenticationDb.Take(item)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		description := fmt.Sprintf(constants.ErrorNotFoundLogMsg, targetName)
		logger.Error(fmt.Sprintf(description+constants.ReqItem, (*item).String()))
		return customerror.NewNotFoundError(description, &result.Error, nil)
	}

	if result.Error != nil {
		description := fmt.Sprintf(constants.ErrorDuringFindLogMsg, targetName)
		logger.Error(fmt.Sprintf(description+constants.ReqItem+constants.Error, (*item).String(), result.Error))
		return customerror.NewInternalServerError(description, &result.Error, nil)
	}

	return nil
}
```

## Modular Log Messages
Also with this pr, i implemented AUT-43 gormlogger inside that. We dont need gorm logger anymore. Because we are logging errors like that way. At the same time error messages much more modular. We can create our log mesages by constant message blocks.

We are creating first part of our error message like this way.
```go
const (
	ErrorDuringFindLogMsg   string = `Something went wrong during find "%s"!`
	Id                      string = ` - Id: %s`
	Error                   string = ` - Error: %s`
	ReqItem                 string = ` - Request Item: %s`
)
)
description := fmt.Sprintf(constants.ErrorDuringFindLogMsg, targetName)
```
After that we can our message extra information like this. We are added request item, and cause error to log this way.
```go
logger.Error(fmt.Sprintf(description+constants.ReqItem+constants.Error, (*item).String(), result.Error))
```
But we dont want to show user extre information. For this reason we are using only description when creating error. We added result.Error as cause error. But InternalServerError will not include CauseError when converted to DTO
```go
return customerror.NewInternalServerError(description, &result.Error, nil)
```

## I18n Errors
Also new error structure provides i18n support. But not implemented. We can create ```Translatable```struct like this way.
utils/fiber/fiber.go
```go
	if validationErrs := server.ValidateStruct(*body); validationErrs != nil {
		return nil, customerror.NewValidationError("Body not fitting validation rules", &validationErrs,
			i18n.NewTranslatable("error.validation", &map[string]interface{}{
				"Code": "c",
				"Number": 2,
			}),
		)
	}
```
First we are giving i18n keyword as first parameter. After that we can give variables for Translation's variable fields by &map[string]interface{}

## Much More Clean

All this changes provides to writing pretty clean services and controller like that:

grouppolicy/service.go
```go
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
```