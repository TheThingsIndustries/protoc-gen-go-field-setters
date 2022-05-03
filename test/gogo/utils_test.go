// Copyright © 2022 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

package test_test

import (
	"testing"

	proto "github.com/gogo/protobuf/proto"
	"github.com/google/go-cmp/cmp"
)

func expectMessage(t *testing.T, expected, got proto.Message) {
	t.Helper()

	expectedMsgText := proto.MarshalTextString(expected)
	gotMsgText := proto.MarshalTextString(got)

	generatedDiff := cmp.Diff(expectedMsgText, gotMsgText)

	if generatedDiff != "" {
		t.Errorf("expected : %s", expectedMsgText)
		t.Errorf("got: %s", gotMsgText)
		if generatedDiff != "" {
			t.Errorf("  diff   : %s", generatedDiff)
		}
	}
}
