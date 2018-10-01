package dson

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func isNum(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}

func getRandom(s ...string) string {
	rand.Seed(time.Now().UnixNano())
	return s[rand.Intn(len(s)-1)]
}

func getOctal(nums string) string {
	n, _ := strconv.Atoi(nums)
	return strconv.FormatInt(int64(n), 8)
}

func getDecimal(nums string) string {
	n, _ := strconv.ParseInt(nums, 8, 64)
	return strconv.Itoa(int(n))
}

func generateVery(number string) string {
	oct := getOctal(number)
	if strings.HasSuffix(oct, "0") {
		if zeros := len(regexp.MustCompile("0").FindAllStringIndex(oct, -1)); zeros >= 1 {
			return strings.Split(oct, "0")[0] + "very" + strconv.Itoa(zeros)
		}
	}
	return oct
}

func parseNumber(number string) string {
	if !strings.Contains(number, "very") {
		return number
	}
	s := strings.Split(strings.ToLower(number), "very")
	ret := s[0]
	if len(s) >= 2 {
		if t, err := strconv.ParseInt(s[1], 10, 64); err == nil {
			ret += strings.Repeat("0", int(t))
		}
	}
	return ret
}

func trimSpace(str string) string { return strings.TrimSpace(str) }
func toLower(str string) string   { return strings.ToLower(str) }
func contains(s, str string) bool { return strings.Contains(s, str) }
func isSeparator(r rune) bool     { return r == ',' || r == '!' || r == '?' || r == '.' }
func exists(s string) bool        { return len(s) > 0 }
