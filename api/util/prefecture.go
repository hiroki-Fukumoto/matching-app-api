package util

import (
	pref "github.com/diverse-inc/jp_prefecture"
)

func GetPrefectureCodeList() []uint16 {
	prefs := pref.List()
	var code []uint16
	for _, p := range prefs {
		code = append(code, uint16(p.Code()))
	}
	return code
}

type prefecture struct {
	Code uint16
	Name string
}

func GetPrefectureCodeAndNameList() []prefecture {
	prefs := pref.List()
	var results []prefecture
	for _, pre := range prefs {
		p := prefecture{uint16(pre.Code()), pre.KanjiShort()}

		results = append(results, p)
	}

	return results
}

func GetPrefectureName(code int) (name *string, err error) {
	prefInfo, ok := pref.FindByCode(code)
	if !ok {
		return nil, err
	} else {
		var n string
		n = prefInfo.KanjiShort()
		return &n, nil
	}
}
