// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package satellite_test

import (
	"strings"
	"testing"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/stretchr/testify/assert"

	"storj.io/common/testcontext"
	"storj.io/storj/integration/ui/uitest"
	"storj.io/storj/private/testplanet"
)

func TestLoginToAccount(t *testing.T) {
	uitest.Run(t, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet, browser *rod.Browser) {
		loginPageURL := "http://" + planet.Satellites[0].API.Console.Listener.Addr().String() + "/login"
		aliceEmail := "alice@mail.test"
		alicePassword := "123a123"

		page := browser.Timeout(25 * time.Second).MustPage(loginPageURL)
		page.MustSetViewport(1350, 600, 1, false)
		page.MustElement(".headerless-input").MustInput(aliceEmail)
		page.MustElement("[type=password]").MustInput(alicePassword)
		page.Keyboard.MustPress(input.Enter)

		assert.True(t, strings.Contains(page.MustElement(".dashboard-area__header-wrapper__title").MustText(), "Dashboard"))
	})
}