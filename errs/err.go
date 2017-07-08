package errs

import (
	"errors"
)

var LibraryNotFound = errors.New("library not found")

var NotInAnyLibrary = errors.New("have pictures not in any library")
