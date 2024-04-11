package model

import "strings"

// colocar seters
func (md *movieDomain) TransformNameToUpperCase() {
	md.name = strings.ToUpper(md.GetName())
}
func (md *movieDomain) TransformDescriptionToUpperCase() {
	md.description = strings.ToUpper(md.GetDescription())
}
func (md *movieDomain) TransformGenreToUpperCase() {
	md.genre = strings.ToUpper(md.GetGenre())
}
