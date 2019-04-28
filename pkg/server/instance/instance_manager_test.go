// +-------------------------------------------------------------------------
// | Copyright (C) 2018 Yunify, Inc.
// +-------------------------------------------------------------------------
// | Licensed under the Apache License, Version 2.0 (the "License");
// | you may not use this work except in compliance with the License.
// | You may obtain a copy of the License in the LICENSE file, or at:
// |
// | http://www.apache.org/licenses/LICENSE-2.0
// |
// | Unless required by applicable law or agreed to in writing, software
// | distributed under the License is distributed on an "AS IS" BASIS,
// | WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// | See the License for the specific language governing permissions and
// | limitations under the License.
// +-------------------------------------------------------------------------

package instance

import (
	"os"
	"path"
	"runtime"
	"testing"
)

var getim = func() InstanceManager {
	// get storage class
	var filePath string
	if runtime.GOOS == "linux" {
		filePath = path.Join(os.Getenv("GOPATH"), "src/github.com/yunify/qingcloud-csi/deploy/disk/kubernetes/config.yaml")
	}
	if runtime.GOOS == "darwin" {
		filePath = path.Join(os.Getenv("GOPATH"), "src/github.com/yunify/qingcloud-csi/deploy/disk/kubernetes/config.yaml")
	}
	im, err := NewInstanceManagerFromFile(filePath)
	if err != nil {
		return nil
	}

	return im
}

func TestFindInstance(t *testing.T) {
	im := getim()
	testcases := []struct {
		name  string
		id    string
		found bool
	}{
		{
			name:  "Available",
			id:    "fake",
			found: true,
		},
		{
			name:  "Not found",
			id:    "instance-1234",
			found: false,
		},
		{
			name:  "By name",
			id:    "neonsan-test",
			found: false,
		},
	}
	for _, v := range testcases {
		ins, err := im.FindInstance(v.id)
		if err != nil {
			t.Errorf("name %s error: %s", v.name, err.Error())
		}
		if v.found && (ins == nil || *ins.InstanceID != v.id) {
			t.Errorf("name %s: find id error", v.name)
		}
	}
}
