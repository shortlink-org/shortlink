/*
Copyright 2020 The Kubernetes Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package csi_driver

import (
	"testing"
)

func TestGetSnapshotID(t *testing.T) {
	testCases := []struct {
		name               string
		inputPath          string
		expectedIsSnapshot bool
		expectedSnapshotID string
	}{
		{
			name:               "should recognize foo.snap as a valid snapshot with ID foo",
			inputPath:          "foo.snap",
			expectedIsSnapshot: true,
			expectedSnapshotID: "foo",
		},
		{
			name:               "should recognize baz.tar.gz as an invalid snapshot",
			inputPath:          "baz.tar.gz",
			expectedIsSnapshot: false,
			expectedSnapshotID: "",
		},
		{
			name:               "should recognize baz.tar.snap as a valid snapshot with ID baz.tar",
			inputPath:          "baz.tar.snap",
			expectedIsSnapshot: true,
			expectedSnapshotID: "baz.tar",
		},
		{
			name:               "should recognize baz.tar.snap.snap as a valid snapshot with ID baz.tar.snap",
			inputPath:          "baz.tar.snap.snap",
			expectedIsSnapshot: true,
			expectedSnapshotID: "baz.tar.snap",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualIsSnapshot, actualSnapshotID := getSnapshotID(tc.inputPath)
			if actualIsSnapshot != tc.expectedIsSnapshot {
				t.Errorf("unexpected result for path %s, Want: %t, Got: %t", tc.inputPath, tc.expectedIsSnapshot, actualIsSnapshot)
			}
			if actualSnapshotID != tc.expectedSnapshotID {
				t.Errorf("unexpected snapshotID for path %s, Want: %s; Got :%s", tc.inputPath, tc.expectedSnapshotID, actualSnapshotID)
			}
		})
	}
}
