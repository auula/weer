package model

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
// responseBody, err := UnmarshalResponseBody(bytes)
// bytes, err = responseBody.Marshal()

import "encoding/json"

// UnmarshalResponseBody https://www.alapi.net/doc/show/36.html
func UnmarshalResponseBody(data []byte) (ResponseBody, error) {
	var r ResponseBody
	err := json.Unmarshal(data, &r)
	return r, err
}

// Marshal https://www.alapi.net/doc/show/36.html
func (r *ResponseBody) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// ResponseBody https://www.alapi.net/doc/show/36.html
type ResponseBody struct {
	Code   int8   `json:"code"`
	Msg    string `json:"msg"`
	Data   Data   `json:"data"`
	Author Author `json:"author"`
}

// Author https://www.alapi.net/doc/show/36.html
type Author struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

// Data https://www.alapi.net/doc/show/36.html
type Data struct {
	Basic         Basic           `json:"basic"`
	Update        Update          `json:"update"`
	Status        string          `json:"status"`
	DailyForecast []DailyForecast `json:"daily_forecast"`
}

// Basic https://www.alapi.net/doc/show/36.html
type Basic struct {
	Cid        string `json:"cid"`
	Location   string `json:"location"`
	ParentCity string `json:"parent_city"`
	AdminArea  string `json:"admin_area"`
	Cnty       string `json:"cnty"`
	Lat        string `json:"lat"`
	Lon        string `json:"lon"`
	Tz         string `json:"tz"`
}

// DailyForecast https://www.alapi.net/doc/show/36.html
type DailyForecast struct {
	CondCodeD string `json:"cond_code_d"`
	CondCodeN string `json:"cond_code_n"`
	CondTxtD  string `json:"cond_txt_d"`
	CondTxtN  string `json:"cond_txt_n"`
	Date      string `json:"date"`
	Hum       string `json:"hum"`
	Mr        string `json:"mr"`
	MS        string `json:"ms"`
	Pcpn      string `json:"pcpn"`
	Pop       string `json:"pop"`
	Pres      string `json:"pres"`
	Sr        string `json:"sr"`
	Ss        string `json:"ss"`
	TmpMax    string `json:"tmp_max"`
	TmpMin    string `json:"tmp_min"`
	UvIndex   string `json:"uv_index"`
	Vis       string `json:"vis"`
	WindDeg   string `json:"wind_deg"`
	WindDir   string `json:"wind_dir"`
	WindSc    string `json:"wind_sc"`
	WindSpd   string `json:"wind_spd"`
}

// Update https://www.alapi.net/doc/show/36.html
type Update struct {
	LOC string `json:"loc"`
	UTC string `json:"utc"`
}
