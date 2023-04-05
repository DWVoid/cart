package chrono

import (
	"errors"
	"time"

	"github.com/DWVoid/cart/tool"
)

func BeginningOfThisHour() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, now.Location())
}

func BeginningOfToday() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

func BeginningOfThisWeek() time.Time {
	today := BeginningOfToday()
	return today.AddDate(0, 0, int(-today.Weekday()))
}

func BeginningOfThisMonth() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
}

func BeginningOfThisYear() time.Time {
	now := time.Now()
	return time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
}

// Period Represents a time period of [Start(), End())
type Period struct {
	start, end time.Time
}

func MakePeriod(start, end time.Time) Period {
	return Period{start: start, end: end}
}

func MakePeriodDuration(start time.Time, duration time.Duration) Period {
	return Period{start: start, end: start.Add(duration)}
}

func MakePeriodDate(start time.Time, years, months, days int) Period {
	return Period{start: start, end: start.AddDate(years, months, days)}
}

func ThisHour() Period {
	return MakePeriodDuration(BeginningOfThisHour(), time.Hour)
}

func Today() Period {
	return MakePeriodDate(BeginningOfToday(), 0, 0, 1)
}

func ThisWeek() Period {
	return MakePeriodDate(BeginningOfThisWeek(), 0, 0, 7)
}

func ThisMonth() Period {
	return MakePeriodDate(BeginningOfThisMonth(), 0, 1, 0)
}

func ThisYear() Period {
	return MakePeriodDate(BeginningOfThisYear(), 1, 0, 0)
}

func (p *Period) Start() time.Time {
	return p.start
}

func (p *Period) End() time.Time {
	return p.end
}

func (p *Period) Bounds() (time.Time, time.Time) {
	return p.start, p.end
}

func (p *Period) Add(duration time.Duration) Period {
	start, end := p.Bounds()
	return MakePeriod(start.Add(duration), end.Add(duration))
}

func (p *Period) AddDate(years, months, days int) Period {
	start, end := p.Bounds()
	return MakePeriod(start.AddDate(years, months, days), end.AddDate(years, months, days))
}

func (p *Period) AddSeconds(seconds int) Period {
	return p.Add(time.Duration(seconds) * time.Second)
}

func (p *Period) AddMinutes(minutes int) Period {
	return p.Add(time.Duration(minutes) * time.Minute)
}

func (p *Period) AddHours(hours int) Period {
	return p.Add(time.Duration(hours) * time.Hour)
}

func (p *Period) AddDays(days int) Period {
	return p.AddDate(0, 0, days)
}

func (p *Period) AddMonths(months int) Period {
	return p.AddDate(0, months, 0)
}

func (p *Period) AddYears(years int) Period {
	return p.AddDate(years, 0, 0)
}

func (p *Period) Duration() time.Duration {
	return p.End().Sub(p.Start())
}

func (p *Period) Contains(time time.Time) bool {
	start, end := p.Bounds()
	return (!time.Before(start)) && time.Before(end)
}

func (p *Period) ExtendBefore(duration time.Duration) Period {
	return MakePeriod(p.Start().Add(-duration), p.End())
}

func (p *Period) ExtendAfter(duration time.Duration) Period {
	return MakePeriod(p.Start(), p.end.Add(duration))
}

func (p *Period) ExtendDateBefore(years, months, days int) Period {
	return MakePeriod(p.Start().AddDate(-years, -months, -days), p.End())
}

func (p *Period) ExtendDateAfter(years, months, days int) Period {
	return MakePeriod(p.Start(), p.End().AddDate(years, months, days))
}

func (p *Period) TrimBefore(duration time.Duration) Period {
	if duration > p.Duration() {
		return MakePeriod(p.End(), p.End())
	}
	return MakePeriod(p.Start().Add(duration), p.End())
}

func (p *Period) TrimAfter(duration time.Duration) Period {
	if duration > p.Duration() {
		return MakePeriod(p.Start(), p.Start())
	}
	return MakePeriod(p.Start(), p.end.Add(-duration))
}

func (p *Period) TrimDateBefore(years, months, days int) Period {
	newStart := p.Start().AddDate(years, months, days)
	if newStart.After(p.End()) {
		return MakePeriod(p.End(), p.End())
	}
	return MakePeriod(newStart, p.End())
}

func (p *Period) TrimDateAfter(years, months, days int) Period {
	newEnd := p.End().AddDate(-years, -months, -days)
	if newEnd.Before(p.Start()) {
		return MakePeriod(p.Start(), p.Start())
	}
	return MakePeriod(p.Start(), newEnd)
}

func TryUnionPeriod(left, right Period) (Period, bool) {
	if left.Start().After(right.Start()) {
		tool.Swap(&left, &right)
	}
	if left.End().Before(right.Start()) {
		return Period{}, false
	}
	return MakePeriod(left.Start(), right.End()), true
}

func TryJointPeriod(left, right Period) (Period, bool) {
	if left.Start().After(right.Start()) {
		tool.Swap(&left, &right)
	}
	if left.End().Before(right.Start()) {
		return Period{}, false
	}
	return MakePeriod(right.Start(), left.End()), true
}

func UnionPeriod(left, right Period) (period Period, err error) {
	period, success := TryUnionPeriod(left, right)
	if !success {
		err = errors.New("periods to union does not overlap")
	}
	return
}

func JointPeriod(left, right Period) (period Period, err error) {
	period, success := TryJointPeriod(left, right)
	if !success {
		err = errors.New("periods to joint does not overlap")
	}
	return
}
