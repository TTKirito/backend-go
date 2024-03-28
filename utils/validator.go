package utils

const (
	Active   = "Active"
	Inactive = "Inactive"
	Develop  = "Develop"
	Design   = "Design"
	Man      = "Man"
	Women    = "Women"
	Meeting  = "Meeting"
	Event    = "Event"
	Office   = "Office"
	Online   = "Online"
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

func IsSupportedEventType(eventType string) bool {
	switch eventType {
	case Meeting, Event:
		return true
	}
	return false
}

func IsSupportedVisitType(visitType string) bool {
	switch visitType {
	case Office, Online:
		return true
	}
	return false
}
