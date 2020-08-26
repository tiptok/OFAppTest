package constant

import "fmt"

const (
	Api         = `/api`
	ApiProtocol = `/api/protocol`

	Controller  = `/pkg/port/%v/controller`
	Router      = `/pkg/port/%v/router`
	Application = `/pkg/application/%v`
	Protocol    = `/pkg/protocol/%v/`
)

func WithController(lib string) string {
	if len(lib) == 0 {
		return fmt.Sprintf(Controller, "beego")
	}
	return fmt.Sprintf(Controller, lib)
}

func WithRouter(lib string) string {
	if len(lib) == 0 {
		return fmt.Sprintf(Router, "beego")
	}
	return fmt.Sprintf(Router, lib)
}

func WithApplication(method string) string {
	if len(method) == 0 {
		return fmt.Sprintf(Application, "test")
	}
	return fmt.Sprintf(Application, method)
}

func WithProtocol(method string) string {
	if len(method) == 0 {
		return fmt.Sprintf(Protocol, "test")
	}
	return fmt.Sprintf(Protocol, method)
}
