package util

import (
	"bytes"
	"encoding/json"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hiroki-Fukumoto/matching-app-api/api/error_handler"
	"github.com/hiroki-Fukumoto/matching-app-api/api/model"
)

// ミドルウェアからセットされたログイン中のユーザー情報を取得
func GetLoginUser(c *gin.Context) (user *model.User, err error) {
	authorizedUser, ok := c.Get("AuthorizedUser")
	if !ok {
		return nil, error_handler.ErrUnauthorized
	}

	return authorizedUser.(*model.User), nil
}

// structをmapに変換
func StructToJsonTagMap(data interface{}) (map[string]interface{}, error) {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	out := new(bytes.Buffer)
	err = json.Indent(out, jsonStr, "", "    ")
	if err != nil {
		return nil, err
	}
	var mapData map[string]interface{}
	if err := json.Unmarshal([]byte(out.String()), &mapData); err != nil {
		return nil, err
	}
	return mapData, err
}

// 配列の中に指定された値が含まれるかチェック
func Contains(list interface{}, elem interface{}) bool {
	listV := reflect.ValueOf(list)

	if listV.Kind() == reflect.Slice {
		for i := 0; i < listV.Len(); i++ {
			item := listV.Index(i).Interface()

			if !reflect.TypeOf(elem).ConvertibleTo(reflect.TypeOf(item)) {
				continue
			}

			target := reflect.ValueOf(elem).Convert(reflect.TypeOf(item)).Interface()

			if ok := reflect.DeepEqual(item, target); ok {
				return true
			}
		}
	}

	return false
}

// 年齢から誕生月を算出
func CalcBirthdayMonthFromAge(age int) string {
	now := time.Now()
	m := age
	year := now.Year() - m
	month := now.Month()
	day := now.Day()
	newMonth := int(month)
	newLastDay := getLastDay(year, newMonth)
	var newDay int
	if day > newLastDay {
		newDay = newLastDay
	} else {
		newDay = day
	}

	d := time.Date(year, time.Month(newMonth), newDay, now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), now.Location())
	date := FormatDate(d)
	return date

}

func getLastDay(year, month int) int {
	t := time.Date(year, time.Month(month+1), 1, 0, 0, 0, 0, time.Local)
	t = t.AddDate(0, 0, -1)
	return t.Day()
}
