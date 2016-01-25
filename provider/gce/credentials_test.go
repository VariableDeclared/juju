// Copyright 2016 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package gce_test

import (
	"github.com/juju/errors"
	"github.com/juju/testing"
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/environs"
	envtesting "github.com/juju/juju/environs/testing"
)

type credentialsSuite struct {
	testing.IsolationSuite
	provider environs.EnvironProvider
}

var _ = gc.Suite(&credentialsSuite{})

func (s *credentialsSuite) SetUpTest(c *gc.C) {
	s.IsolationSuite.SetUpTest(c)

	var err error
	s.provider, err = environs.Provider("gce")
	c.Assert(err, jc.ErrorIsNil)
}

func (s *credentialsSuite) TestCredentialSchemas(c *gc.C) {
	envtesting.AssertProviderAuthTypes(c, s.provider, "oauth2", "jsonfile")
}

var sampleCredentialAttributes = map[string]string{
	"client-id":    "123",
	"client-email": "test@example.com",
	"project-id":   "fourfivesix",
	"private-key":  "sewen",
}

func (s *credentialsSuite) TestOAuth2CredentialsValid(c *gc.C) {
	envtesting.AssertProviderCredentialsValid(c, s.provider, "oauth2", map[string]string{
		"client-id":    "123",
		"client-email": "test@example.com",
		"project-id":   "fourfivesix",
		"private-key":  "sewen",
	})
}

func (s *credentialsSuite) TestOAuth2HiddenAttributes(c *gc.C) {
	envtesting.AssertProviderCredentialsAttributesHidden(c, s.provider, "oauth2", "private-key")
}

func (s *credentialsSuite) TestJSONFileCredentialsValid(c *gc.C) {
	envtesting.AssertProviderCredentialsValid(c, s.provider, "jsonfile", map[string]string{
		// For now at least, the contents of the file are not validated
		// by the credentials schema. That is left to the provider.
		"file": "whatever",
	})
}

func (s *credentialsSuite) TestDetectCredentialsNotFound(c *gc.C) {
	credentials, err := s.provider.DetectCredentials()
	c.Assert(err, jc.Satisfies, errors.IsNotFound)
	c.Assert(credentials, gc.HasLen, 0)
}
