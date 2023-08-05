package consts

import "regexp"

var TimeRegex, _ = regexp.Compile(`^([01]?[0-9]|2[0-3])\.\d{2}$`)
