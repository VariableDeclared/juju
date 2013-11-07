// Copyright 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package instance_test

import (
	"testing"

	gc "launchpad.net/gocheck"

	"launchpad.net/juju-core/instance"
)

func TestPackage(t *testing.T) {
	gc.TestingT(t)
}

type InstanceSuite struct{}

var _ = gc.Suite(&InstanceSuite{})

func (s *InstanceSuite) TestParseSupportedContainerType(c *gc.C) {
	ctype, err := instance.ParseSupportedContainerType("lxc")
	c.Assert(err, gc.IsNil)
	c.Assert(ctype, gc.Equals, instance.LXC)

	ctype, err = instance.ParseSupportedContainerType("kvm")
	c.Assert(err, gc.IsNil)
	c.Assert(ctype, gc.Equals, instance.KVM)

	ctype, err = instance.ParseSupportedContainerType("none")
	c.Assert(err, gc.ErrorMatches, `invalid container type "none"`)

	ctype, err = instance.ParseSupportedContainerType("omg")
	c.Assert(err, gc.ErrorMatches, `invalid container type "omg"`)
}

func (s *InstanceSuite) TestParseSupportedContainerTypeOrNone(c *gc.C) {
	ctype, err := instance.ParseSupportedContainerTypeOrNone("lxc")
	c.Assert(err, gc.IsNil)
	c.Assert(ctype, gc.Equals, instance.LXC)

	ctype, err = instance.ParseSupportedContainerTypeOrNone("kvm")
	c.Assert(err, gc.IsNil)
	c.Assert(ctype, gc.Equals, instance.KVM)

	ctype, err = instance.ParseSupportedContainerTypeOrNone("none")
	c.Assert(err, gc.IsNil)
	c.Assert(ctype, gc.Equals, instance.NONE)

	ctype, err = instance.ParseSupportedContainerTypeOrNone("omg")
	c.Assert(err, gc.ErrorMatches, `invalid container type "omg"`)
}
