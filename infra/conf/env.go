package conf

// type Environment int

var env = "DEV"

func IsDev() bool {
	return env == "DEV"
}

func IsFP() bool {
	return env == "FP"
}

func IsBeta() bool {
	return env == "BETA"
}

func IsProd() bool {
	return env == "PROD"
}

func IsTest() bool {
	return env == "TEST"
}
