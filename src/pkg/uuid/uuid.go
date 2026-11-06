package uuid

import "uuid"

func Generate() (uuid.UUID) {
	return uuid.NewV7()
}
