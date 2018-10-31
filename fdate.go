package fdate

import (
	"errors"
	"log"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	MIN_YEAR func() int
	MAX_YEAR func() int
)

func init() {
	MIN_YEAR = func() int {
		return 1945
	}
	MAX_YEAR = time.Now().Year
}

var specialCase = regexp.MustCompile(`\d{4}(/|-|年)\d{1,2}(/|-|月)\d{1,2}`)
var delimitersExp = regexp.MustCompile(`(/|-|年|月)`)

func isSpecialCase(s string) bool {
	return specialCase.Copy().MatchString(s)
}

func pickSpecialDate(s string) (time.Time, error) {
	if !isSpecialCase(s) {
		return time.Time{}, errors.New("not find pattern")
	}

	dateStr := specialCase.Copy().FindString(s)
	a := delimitersExp.Copy().Split(dateStr, -1)
	y, _ := strconv.Atoi(a[0])
	m, _ := strconv.Atoi(a[1])
	d, _ := strconv.Atoi(a[2])
	if !ValidationDate(y, m, d) {
		return time.Time{}, errors.New("not find pattern")
	}

	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Now().Location()), nil
}

func isLeapYear(year int) bool {
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		return true
	}
	return false
}

func ValidationDate(y, m, d int) bool {
	if m <= 0 || 12 < m {
		return false
	} else if d <= 0 || 31 < d {
		return false
	}

	switch m {
	case 1, 3, 5, 7, 8, 10, 12:
		return true
	}

	if d == 31 {
		return false
	}

	if m == 2 {
		if isLeapYear(y) {
			return d <= 29
		}
		return d <= 28
	}

	return true
}

func isNumber(r rune) bool {
	return '0' <= r && r <= '9'
}

func isYear(s string) (int, bool) {
	ra := []rune(s)

	if len(ra) != 4 {
		return 0, false
	}

	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, false
	}

	return n, MIN_YEAR() <= n && n <= MAX_YEAR()
}

func isMonth(s string) (int, bool) {
	ra := []rune(s)

	if len(ra) == 0 || 3 <= len(ra) {
		return 0, false
	}

	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, false
	}

	return n, 1 <= n && n <= 12
}

func isDay(s string) (int, bool) {
	ra := []rune(s)

	if len(ra) == 0 || 3 <= len(ra) {
		return 0, false
	}

	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, false
	}

	return n, 1 <= n && n <= 31
}

func getDaySet(s string) []int {
	if len([]byte(s)) == 1 {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil
		}
		return []int{n}
	}
	ret := []int{}

	ra := []rune(s)
	for i := 0; i+1 < len(ra); i++ {
		s := string([]rune{ra[i], ra[i+1]})
		n, err := strconv.Atoi(s)
		if err != nil {
			continue
		}
		ret = append(ret, n)
	}

	for i := 0; i < len(ra); i++ {
		n, err := strconv.Atoi(string([]rune{ra[i]}))
		if err != nil {
			continue
		}
		ret = append(ret, n)
	}

	return ret
}

func pickPossibleDate(s string) ([]time.Time, error) {
	ra := []rune(s)
	ret := []time.Time{}

	for i := 0; i <= len(ra)-6; i++ {
		yearCandidate := ra[i : i+4]
		y, ok := isYear(string(yearCandidate))
		if !ok {
			continue
		}

		backDatasetMD := ra[i+4:]

		for j := 0; j <= len(backDatasetMD)-2; j++ {
			for monthStrLength := 1; monthStrLength <= 2; monthStrLength++ {
				monthCandidate := backDatasetMD[j : j+monthStrLength]
				m, err := strconv.Atoi(string(monthCandidate))
				if err != nil {
					continue
				}

				backDatasetD := backDatasetMD[j+monthStrLength:]
				for k := 0; k <= len(backDatasetD)-2; k++ {
					for dayStrLength := 1; dayStrLength <= 2; dayStrLength++ {
						dayCandidate := backDatasetD[k : k+dayStrLength]
						d, err := strconv.Atoi(string(dayCandidate))
						if err != nil {
							continue
						}
						if !ValidationDate(y, m, d) {
							continue
						}
						ret = append(ret, time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Now().Location()))
					}
				}
			}
		}
	}
	if len(ret) == 0 {
		log.Println("not find pattern")
		return nil, errors.New("not find pattern")
	}
	m := map[string]time.Time{}

	for _, d := range ret {
		m[d.String()] = d
	}

	buf := []time.Time{}

	for _, v := range m {
		buf = append(buf, v)
	}

	today := time.Now()
	sort.Slice(buf, func(i, j int) bool {
		return math.Abs(float64(today.Unix()-buf[i].Unix())) < math.Abs(float64(today.Unix()-buf[j].Unix()))
	})

	return buf, nil
}

func PickPossibleDate(s string) ([]time.Time, error) {
	if len([]byte(s)) > 32 {
		return nil, errors.New("too long string: Input needs to be 32 bytes or less")
	}
	if len([]byte(s)) < 6 {
		return nil, errors.New("not find pattern")
	}

	sd, err := pickSpecialDate(s)
	if err == nil {
		return []time.Time{sd}, nil
	}

	if len([]rune(s)) == 6 {
		for _, r := range []rune(s) {
			if !isNumber(r) {
				log.Println("not find pattern")
				return nil, errors.New("not find pattern")
			}
		}
		ys, ms, ds := s[:4], s[4:5], s[5:]
		y, _ := strconv.Atoi(ys)
		m, _ := strconv.Atoi(ms)
		d, _ := strconv.Atoi(ds)

		if !ValidationDate(y, m, d) {
			log.Println("not find pattern")
			return nil, errors.New("not find pattern")
		}

		return []time.Time{
			time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Now().Location()),
		}, nil
	}

	return pickPossibleDate(strings.Replace(s, " ", "", -1))
}
