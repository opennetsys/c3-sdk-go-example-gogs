// Copyright 2016 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package context

import (
	"github.com/c3systems/c3-sdk-go-example-gogs/models"
)

type APIOrganization struct {
	Organization *models.User
	Team         *models.Team
}
