package constants

const (
	ErrorDuringFindLogMsg   string = `Something went wrong during find "%s"!`
	ErrorNotFoundLogMsg     string = `%s not found!`
	NotFoundForDeleteLogMsg string = `%s not found for delete!`
	NotFoundForUpdateLogMsg string = `%s not found for update!`
	ErrorDuringDeleteLogMsg string = `Something went wrong during delete "%s"!`
	NotDeletedLogMsg        string = `%s not deleted!`
	ErrorDuringUpdateLogMsg string = `Something went wrong during update "%s"!`
	NotUpdatedLogMsg        string = `%s not updated!`
	ErrorDuringCreateLogMsg string = `Something went wrong during creation of "%s"!`
	ConflictExistedIn       string = `Relationship of %s existed. New entry causes conflict. Please try to add not existed relationship`
	Id                      string = ` - Id: %s`
	Error                   string = ` - Error: %s`
	ReqItem                 string = ` - Request Item: %s`
)
