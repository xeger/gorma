//************************************************************************//
// API "congo": Application Resource Href Factories
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/goadesign/gorma/example
// --design=github.com/goadesign/gorma/example/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "fmt"

// UserHref returns the resource href.
func UserHref(userID interface{}) string {
	return fmt.Sprintf("/api/users/%v", userID)
}
