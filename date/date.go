package date

import (
	"fmt"
	"strings"
	"time"
)

type Date time.Time

func New(year, month, day int) Date {
	return Date(time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC))
}

func Parse(input string) (date Date, err error) {
	if tim, err := time.Parse(time.DateOnly, input); err == nil {
		return Date(tim), nil
	} else {
		return date, err
	}
}

func Today() Date {
	year, month, day := time.Now().Date()

	return New(year, int(month), day)
}

func (date Date) IsZero() bool {
	return time.Time(date).IsZero()
}

func (date Date) String() string {
	return time.Time(date).Format(time.DateOnly)
}

func (date Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", date.String())), nil
}

func (date *Date) UnmarshalJSON(bts []byte) error {
	str := strings.Trim(string(bts), `"`)

	tim, err := time.Parse(time.DateOnly, str)
	if err != nil {
		return err
	}

	*date = Date(tim)

	return nil
}
