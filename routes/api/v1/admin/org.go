// Copyright 2015 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package admin

import (
	api "github.com/gogs/go-gogs-client"

	"github.com/c3systems/c3-sdk-go-example-gogs/pkg/context"
	"github.com/c3systems/c3-sdk-go-example-gogs/routes/api/v1/org"
	"github.com/c3systems/c3-sdk-go-example-gogs/routes/api/v1/user"
)

// https://github.com/gogs/go-gogs-client/wiki/Administration-Organizations#create-a-new-organization
func CreateOrg(c *context.APIContext, form api.CreateOrgOption) {
	org.CreateOrgForUser(c, form, user.GetUserByParams(c))
}
