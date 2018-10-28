package common

import (
	"strconv"
	"time"
)

type JsonTime struct {
	time.Time
}

func (t *JsonTime) UnmarshalJSON(b []byte) error {
	s := string(b)
	if s == "null" {
		t.Time = time.Time{}
		return nil
	}

	local, _ := time.ParseInLocation(`"2006-01-02 15:04:05"`, s, time.Local)
	t.Time = local
	return nil
}

func (c *JsonTime) HourString() string {
	currentYear := time.Now().Year()
	year := c.Year()
	month := c.Month()
	day := c.Day()
	hour := c.Hour()

	if currentYear == year {
		return strconv.Itoa(int(month)) + "月" + strconv.Itoa(day) + "日" + strconv.Itoa(hour) + "时"
	}
	return strconv.Itoa(year) + "年" + strconv.Itoa(int(month)) + "月" + strconv.Itoa(day) + "日" + strconv.Itoa(hour) + "时"
}

func (c *JsonTime) DayString() string {
	currentYear := time.Now().Year()
	year := c.Year()
	month := c.Month()
	day := c.Day()

	if currentYear == year {
		return strconv.Itoa(int(month)) + "月" + strconv.Itoa(day) + "日"
	}
	return strconv.Itoa(year) + "年" + strconv.Itoa(int(month)) + "月" + strconv.Itoa(day) + "日"
}
