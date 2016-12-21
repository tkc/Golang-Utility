package DataBase

import (
		"time"
)

func ConvertDateMysqlDateTime(dateTime time.Time) string {
		convertTime, _ := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", dateTime.String())
		convertDate := convertTime.Format("2006年01月02日 15時04分")
		return convertDate
}
