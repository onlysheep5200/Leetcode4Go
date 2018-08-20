package main

import "fmt"

/**
Given a set of non-overlapping intervals, insert a new interval into the intervals (merge if necessary).

You may assume that the intervals were initially sorted according to their start times.

Example 1:

Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
Output: [[1,5],[6,9]]
Example 2:

Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
Output: [[1,2],[3,10],[12,16]]
Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].
 */

type Interval struct {
	Start int
	End   int
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	if len(intervals) == 0 {
		intervals = append(intervals, newInterval)
		return intervals
	}

	newIntervals := make([]Interval, 0, len(intervals)+1)


	//针对新的interval在第一个时的场景
	if newInterval.End < intervals[0].Start{
		newIntervals = append(newIntervals, newInterval)
		newIntervals = append(newIntervals, intervals...)
		return newIntervals
	}

	mergeState := 0
	inserted := false
	for i := 0; i < len(intervals); i++ {
		switch mergeState {
			case 0:
				if canMerge(intervals[i], newInterval) {
					intervals[i].Start = lessOne(intervals[i].Start, newInterval.Start)
					intervals[i].End = largerOne(intervals[i].End, newInterval.End)
					mergeState = 1
					inserted = true
					newIntervals = append(newIntervals, intervals[i])
					continue
				}
				newIntervals = append(newIntervals, intervals[i])
				if i < len(intervals)-1 && intervals[i].End < newInterval.Start && intervals[i+1].Start > newInterval.End{
					newIntervals = append(newIntervals, newInterval)
					inserted = true
					mergeState = 3
				}
			case 1:
				if len(newIntervals) != 0 && canMerge(intervals[i], newIntervals[len(newIntervals)-1]){
					mergeState = 2
				}else {
					mergeState = 3
				}
				i--
			case 2:
				//slice中的值若非指针，则取值后返回的是原始值的拷贝，对返回的值操作不会影响原有值
				latest := &newIntervals[len(newIntervals) - 1]
				if canMerge(*latest, intervals[i]){
					latest.Start = lessOne(latest.Start, intervals[i].Start)
					latest.End = largerOne(latest.End, intervals[i].End)
				}else{
					mergeState = 3
					i--
				}
			case 3:
				newIntervals = append(newIntervals, intervals[i])
		}
	}
	if mergeState == 0 &&  !inserted{
		newIntervals = append(newIntervals, newInterval)
	}
	return newIntervals
}

func canMerge(i, j Interval) bool {
	if i.Start >= j.Start && i.Start <= j.End{
		return true
	}

	if j.Start >= i.Start && j.Start <= i.End {
		return true
	}

	return false
}

func largerOne(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func lessOne(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {
	intervals := []Interval{Interval{1,2}, Interval{3,5}, Interval{6,7}, Interval{8,10},
	Interval{12,16}}

	newInterval := Interval{4,8}

	fmt.Println(insert(intervals, newInterval))
}
