package utils

const (
	Active   = "Active"
	Inactive = "Inactive"
	Develop  = "Develop"
	Design   = "Design"
	Man      = "Man"
	Women    = "Women"
)

func IsSupportedStatus(status string) bool {
	switch status {
	case Active, Inactive:
		return true
	}
	return false
}

func IsSupportedPosition(position string) bool {
	switch position {
	case Design, Develop:
		return true
	}

	return false
}

func IsSupportedGender(gender string) bool {
	switch gender {
	case Man, Women:
		return true
	}

	return false
}
