package utilities

import "time"

func GetNowDDHHMMSS() (string, string, string, string) {
	timestamp := time.Now()
	return timestamp.Format("2006-01-02"), timestamp.Format("15"), timestamp.Format("04"), timestamp.Format("05")
}

func GetLastMinuteDDHHMMSS() (string, string, string, string) {
	timestamp := time.Now().Add(-time.Minute)
	return timestamp.Format("2006-01-02"), timestamp.Format("15"), timestamp.Format("04"), timestamp.Format("05")
}

func GetLastTenMinuteDDHHMMSS() (string, string, string, string) {
	timestamp := time.Now().Add(-time.Minute * 10)
	return timestamp.Format("2006-01-02"), timestamp.Format("15"), timestamp.Format("04"), timestamp.Format("05")
}

func GetYesterdayDDHHMMSS() (string, string, string, string) {
	timestamp := time.Now().Add(-time.Hour * 24)
	return timestamp.Format("2006-01-02"), timestamp.Format("15"), timestamp.Format("04"), timestamp.Format("05")
}
