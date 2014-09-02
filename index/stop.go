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
	"strings"
)

var stopWords = createStopWordsMap()

func createStopWordsMap() map[string]bool {
	m := make(map[string]bool)
	for _, word := range strings.Fields(stopWordsText) {
		m[word] = true
	}
	return m
}

// taken from http://www.ranks.nl/stopwords
const stopWordsText = `
I
a
about
an
are
as
at
be
by
com
for
from
how
in
is
it
of
on
or
that
the
this
to
was
what
when
where
who
will
with
the
www
`
