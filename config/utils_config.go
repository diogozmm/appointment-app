package config

import "appointment-app/app/shared"

type Utils struct {
	Utils shared.Utils
}

func InitUtils() *Utils {
	return &Utils{
		Utils: shared.InitUtils(),
	}
}
