package helpers

type LoadSchema[M any] interface {
	Fill(*M) *ValidationError
}
