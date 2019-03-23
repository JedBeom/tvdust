package main

func setGrade(a, b int) Data {
	var (
		data Data
		mixed int
	)

	if a > b {
		mixed = a
	} else {
		mixed = b
	}

	switch mixed {
	case 1:
		data.Grade = "최고"
		data.Bg = "#4b74b8"
	case 2:
		data.Grade = "좋음"
		data.Bg = "#5c98d1"
	case 3:
		data.Grade = "양호"
		data.Bg = "#63abc0"
	case 4:
		data.Grade = "보통"
		data.Bg = "#5b8f4a"
	case 5:
		data.Grade = "나쁨"
		data.Bg = "#de8b2f"
	case 6:
		data.Grade = "상당히 나쁨"
		data.Bg = "#c94e2c"
	case 7:
		data.Grade = "매우 나쁨"
		data.Bg = "#b83134"
	case 8:
		data.Grade = "최악"
		data.Bg = "#262626"
	}

	return data
}