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

func NewDateRange(from string, to string) *DateRange {
	start := date.FromTime(strptime.MustParse(from, "%Y%m%d"))
	end := date.FromTime(strptime.MustParse(to, "%Y%m%d"))

	if end.Time.Sub(start.Time) < 0 {
		return nil
	}

	return &DateRange{date.NewRange(start, end)}
}
