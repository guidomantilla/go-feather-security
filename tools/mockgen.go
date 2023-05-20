package tools

//go:generate mockgen -package=auth		-source ../pkg/auth/types.go 		-destination ../pkg/auth/mocks.go
//go:generate mockgen -package=password	-source ../pkg/password/types.go 	-destination ../pkg/password/mocks.go
