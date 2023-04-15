package ski

import (
	"strconv"
	"time"
)

type ski struct {
	datetime time.Time
	// Indicates whether the `ski` object is a repeated wall time in the current timezone.
	Ambiguous bool
	// Returns a datetime representation of the `ski` object.
	Datetime time.Time
	// Returns a naive datetime representation of the ski object.
	Naive time.Time
	// Returns a floating-point timestamp representation of the ski object, in UTC time.
	FloatTimestamp float64
	// Returns a int64 timestamp representation of the `ski` object, in local time.
	LocalizedFloatTimestamp float64
	// Returns a int64 representation of the `ski` object.
	IntTimestamp int64
	// Returns the fold value of the ski object.
	Fold int64
	// Indicates whether the ski object exists in the current timezone.
	Imaginary bool
	// The timezone for the `ski` object.
	TZInfo time.Location
}

type SkiOpt struct {
	Year        int
	Month       int
	Day         int
	Hour        int
	Minute      int
	Second      int
	Microsecond int
	TZ          string
	Fold        int
}

func NewSki(opt *SkiOpt) *ski {
	s := &ski{}
	s.datetime = time.Now()
	return s
}

// Returns a datetime object, converted to the specified timezone.
func (s *ski) ASTimezone(tz *time.Location) *time.Time {
	return nil
}

// Frame Literal
type Frame string

const (
	Microsecond  Frame = "microsecond"  //"µ"
	Microseconds Frame = "microseconds" //"µs"
	Second       Frame = "second"
	Seconds      Frame = "seconds"
	Minute       Frame = "minute"
	Minutes      Frame = "minutes"
	Hour         Frame = "hour"
	Hours        Frame = "hours"
	Day          Frame = "day"
	Days         Frame = "days"
	Week         Frame = "week"
	Weeks        Frame = "weeks"
	Month        Frame = "month"
	Months       Frame = "months"
	Quarter      Frame = "quarter"
	Quarters     Frame = "quarters"
	Year         Frame = "year"
	Years        Frame = "years"
	Auto         Frame = "auto"
)

// Returns a new ski object,
// representing the “ceiling” of the timespan of the ski object in a given timeframe.
// Equivalent to the second element in the 2-tuple returned by span.
func (s *ski) Ceil(frame Frame) *ski {
	return nil
}

func (s *ski) Clone() *ski {
	newS := ski{}
	newS = *s
	return &newS
}

func (s *ski) CTime() string {
	return ""
}

func (s *ski) Date() *time.Time {
	return nil
}

func (s *ski) DeHumanize(inputStr string, locale ...string) *ski {
	// localeName := "en_us"
	// if len(locale) > 0 {
	// 	localeName = locale[0]
	// }
	// localeObj, _ := time.LoadLocation(localeName)
	// normalLocalName := strings.ReplaceAll(strings.ToLower(localeName), "_", "-")

	return nil
}

func (s *ski) DST() *time.Duration {
	return nil
}

func (s *ski) Floor(frame Frame) *ski {
	return nil
}

func (s *ski) ForJSON() string {
	return ""
}

func (s *ski) Format(fmt string, local ...string) string {
	return ""
}

func (s *ski) Humanize() string {
	return ""
}

type Bounds string

const (
	LeftCloseRightOpen  Bounds = "[)" //"left-close-right-open"
	LeftOpenRightClose  Bounds = "(]" //"left-open-right-close"
	LeftCloseRightClose Bounds = "[]" //"left-close-right-close"
	LeftOpenRightOpen   Bounds = "()" //"left-open-right-open"

)

func (s *ski) IsBetween(start, end *ski, bounds ...Bounds) bool {
	return false
}

// Returns a 3-tuple, (ISO year, ISO week number, ISO weekday).
func (s *ski) ISOCalendar() (int, int, int) {
	return 0, 0, 0
}

func (s *ski) ISOFormat(sep string, tz time.Location) string {
	return ""
}

// Returns the ISO day of the week as an integer (1-7).
func (s *ski) ISOWeekDay() int {
	return int(s.datetime.UTC().Weekday())
}

type ReplaceOpt struct {
}

func (s *ski) Replace(opt *ReplaceOpt) *ski {
	return nil
}

type ShiftOpt struct {
	Year        int
	Quarter     int
	Month       int
	Day         int
	Hour        int
	Minute      int
	Second      int
	Microsecond int
}

const (
	MonthsPerQuarter = 3
	SecsPerMinute    = 60
	SecsPerHour      = 60 * 60
	SecsPerDay       = 60 * 60 * 24
	SecsPerWeek      = 60 * 60 * 24 * 7
	SecsPerMonth     = 60 * 60 * 24 * 30.5
	SecsPerQuarter   = 60 * 60 * 24 * 30.5 * 3
	SecsPerYear      = 60 * 60 * 24 * 365
)

func (s *ski) FromDatetime(t time.Time) *ski {
	s.datetime = t
	return s
}

func (s *ski) Shift(opt *ShiftOpt) *ski {
	var duration int64 = 0
	duration += 
	s.datetime.Add()
	return s.FromDatetime()
}

type SpanOpt struct {
	Frame     Frame
	Count     int
	Bounds    Bounds
	Exact     bool
	WeekStart int
	TZ        time.Location
	Limit     int
}

func (s *ski) Span(opt *SpanOpt) []*ski {
	return nil
}

func (s *ski) StrFormatTime(fmt string) string {
	return s.String(fmt)
}

func (s *ski) String(fmt ...string) string {
	return ""
}

func (s *ski) Time() *time.Time {
	return nil
}

func (s *ski) Timestamp() float64 {
	return float64(s.datetime.UnixNano()) / 1000
}

func (s *ski) TimeTuple() {

}

func (s *ski) TimeTZ() *time.Time {
	return nil
}

func (s *ski) To(tz *time.Location) *ski {
	return nil
}

// Returns the proleptic Gregorian ordinal of the date.
func (s *ski) ToOrdinal() int {
	return 0
}

func (s *ski) UTCOffset() *time.Duration {
	return nil
}

func (s *ski) UTCTimeTuple() {
}

func (s *ski) Weekday() int {
	return 0
}

func NewFromDate() *ski {
	return nil
}

func NewFromDatetime() *ski {
	return nil
}

func NewFromOrdinal(ordinal int) *ski {
	return nil
}

func NewFromTimestamp(ts int64, tz ...time.Location) *ski {
	return nil
}

func NewFromInterval(frame Frame, start, end time.Time, interval ...int) []*ski {
	return nil
}

func (s *ski) UTCFromTimestamp(timestamp string) *ski {
	// if timestamp is str
	i, _ := strconv.ParseInt(timestamp, 10, 64)
	s.datetime = time.Unix(int64(i/1000), int64(i%1000))
	return s
}

func StrPtime(dateStr, fmt string, tz time.Location) *ski {
	return NewFromStr(dateStr, fmt, tz)
}

func NewFromStr(dateStr, fmt string, tz time.Location) *ski {
	return nil
}
