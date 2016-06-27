package stow

type errUnknownKind string

func (e errUnknownKind) Error() string {
	s := string(e)
	if len(s) > 0 {
		return "stow: unknown kind \"" + string(e) + "\""
	}
	return "stow: unknown kind"
}
