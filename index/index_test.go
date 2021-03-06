// Copyright (C) 2014 Frank Guan
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
// 
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package index

import (
	"testing"
)

func TestAddPageToIndex(t *testing.T) {
	index := map[string][]string{}
	AddPageToIndex(index, "fake.test", "This is a test")
	t.Logf("%v", index)

	AddPageToIndex(index, "not.test", "This is not a test")
	t.Logf("%v", index)
}
