/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-08 13:08:43
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-10 16:00:20
 */
package model

import "encoding/json"

type Profile struct {
	Name       string
	Gender     string
	Age        int
	Height     int
	Weight     int
	Income     string
	Marriage   string
	Education  string
	Occupation string
	Hokou      string
	Xinzuo     string
	House      string
	Car        string
}

func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(s, &profile)
	return profile, err
}
