package controller

import (
	"github.com/splicemaahs/orderedhttp-operator/pkg/controller/orderedhttp"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, orderedhttp.Add)
}
