// Copyright 2014 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package action_test

import (
	"github.com/juju/juju/cmd/juju/action"
	gc "gopkg.in/check.v1"
)

type StatusSuite struct {
	BaseActionSuite
	UndefinedActionCommandSuite
	subcommand *action.StatusCommand
}

func (s *StatusSuite) SetUpTest(c *gc.C) {
	s.BaseActionSuite.SetUpTest(c)
	s.subcommand = &action.StatusCommand{}
}

func (s *StatusSuite) TestHelp(c *gc.C) {
	s.checkHelp(c, s.subcommand)
}
