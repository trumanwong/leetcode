package squareIsWhite

func SquareIsWhite(coordinates string) bool {
	result := false
	for i := int('a'); i < int(coordinates[0]); i++ {
		result = !result
	}
	for i := int('1'); i < int(coordinates[1]); i++ {
		result = !result
	}
	return result
}
