package dto

type TemperatureOutput struct {
	City  string  `json:"city"`
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

type TemperatureAPIOutput struct {
	Location string  `json:"location"`
	TempC    float64 `json:"temp_C"`
	TempF    float64 `json:"temp_F"`
	TempK    float64 `json:"temp_K"`
}
