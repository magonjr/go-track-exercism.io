package leap

// IsLeapYear given a year if it's a leap year
/*
on every year that is evenly divisible by 4
  except every year that is evenly divisible by 100
    unless the year is also evenly divisible by 400
*/
func IsLeapYear(year int) bool {
	leap := year%4 == 0 && (!(year%100 == 0) || (year%400 == 0))
	return leap
}
