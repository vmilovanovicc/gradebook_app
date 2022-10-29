package registry

// Hold all of the information about what a service registration looks like.

type Registration struct {
	ServiceName ServiceName
	ServiceURL  string
}
type ServiceName string

const (
	LogService = ServiceName("LogService")
)
