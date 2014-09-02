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

package pagerank

const (
	d = 0.8  // damping factor
	numLoops = 10
)


func ComputeRanks(graph map[string][]string) map[string]float64 {
	ranks := make(map[string]float64)
	npages := len(graph)
	for page, _ := range graph {
		ranks[page] = 1.0 / float64(npages)
	}

	for i := 0; i < numLoops; i++ {
		newRanks := make(map[string]float64)
		for page, _ := range graph {
			newRank := (1 - d) / float64(npages)
			// summation of inlinks
			for node, _ := range graph {
				newRank = newRank + d * (ranks[node] / float64(len(graph[node])))
			}

			newRanks[page] = newRank
		}
		ranks = newRanks

	}
	return ranks
}

func Search(index map[string][]string, ranks map[string]float64, keyword string) string {
	highestPage := ""
	highestRank := float64(0)
	if allPages, ok := index[keyword]; ok {
		for _, page := range allPages {
			if rank, ok := ranks[page]; ok {
				if rank > highestRank {
					highestRank = rank
					highestPage = page
				}
			}
		}
	}
	return highestPage
}
