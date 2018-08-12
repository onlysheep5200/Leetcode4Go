package main

import "sort"

type Interval struct {
	Start int
	End   int
}

type Intervals []Interval;

func (self Intervals) Len() int {
	return len(self)
}

func (self Intervals) Less(i, j int) bool {
	return self[i].Start < self[j].Start
}

func (self Intervals) Swap(i, j int)  {
	self[i], self[j] = self[j], self[i]
}

/**
Given a collection of intervals, merge all overlapping intervals.

Example 1:

Input: [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
Example 2:

Input: [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considerred overlapping.
 */
func merge(intervals []Interval) []Interval {
	if len(intervals) <= 0 {
		return intervals
	}

	sort.Sort(Intervals(intervals))

	result := make([]Interval, 0, len(intervals))
	for i := 1; i< len(intervals); i++{
		if intervals[i].Start <= intervals[i-1].End {
			intervals[i].Start = intervals[i-1].Start
			end := intervals[i].End
			if end < intervals[i-1].End{
				end = intervals[i-1].End
			}
			intervals[i].End = end
		}else{
			result = append(result, intervals[i-1])
		}
	}

	result = append(result, intervals[len(intervals) - 1])
	return result
}

func main() {

}