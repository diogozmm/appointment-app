package enum

type UserType int

const (
	Client UserType = iota
	Artist
)

var userTypeStrings = [...]string{
	"Client",
	"Artist",
}

func (userType UserType) ToString() string {
	if userType < Client || userType > Artist {
		return "Unknown"
	}
	return userTypeStrings[userType]
}
