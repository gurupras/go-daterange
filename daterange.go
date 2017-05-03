package daterange

import (
	"time"

	"github.com/aodin/date"
	"github.com/pbnjay/strptime"
)

type DateRange struct {
	date.Range
}

func (dr *DateRange) AsList() []time.Time {
	start := dr.Start
	end := dr.End

	ret := make([]time.Time, 0)
	ret = append(ret, start.Time)
	idx := 1
	for {
		day := start.AddDays(idx)
		if day.Time.Sub(end.Time) >= 0 {
			break
		}
		ret = append(ret, day.Time)
		idx += 1
	}
	return ret
}

func (dr *DateRange) ContainsDate(t time.Time) bool {
	d := date.FromTime(t)
	if (dr.Start.Equals(d) || dr.Start.Before(d)) && dr.End.After(d) {
		return true
	}
	return false
}

func NewDateRange(from string, to string) *DateRange {
	str, err := strptime.Parse(from, "%Y%m%d")
	if err != nil {
		return nil
	}
	start := date.FromTime(str)
	str, err = strptime.Parse(to, "%Y%m%d")
	if err != nil {
		return nil
	}
	end := date.FromTime(str)

	if end.Time.Sub(start.Time) < 0 {
		return nil
	}

	return &DateRange{date.NewRange(start, end)}
}
