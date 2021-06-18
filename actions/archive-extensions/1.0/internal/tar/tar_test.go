// Copyright (c) 2021 Terminus, Inc.
//
// This program is free software: you can use, redistribute, and/or modify
// it under the terms of the GNU Affero General Public License, version 3
// or later ("AGPL"), as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package tar_test

import (
	"path/filepath"
	"testing"

	"github.com/erda-project/erda-actions/actions/archive-extensions/1.0/internal/tar"
)

func TestTar(t *testing.T) {
	var src = ".."
	abs, err := filepath.Abs(src)
	if err != nil {
		t.Fatalf("failed to filepath.Abs: %v", err)
	}
	dst, err := tar.Tar(abs, "")
	if err != nil {
		t.Fatalf("failed to tar.Tar, src: %s: %v", src, err)
	}
	t.Log(dst)
}
