package zeus

import "regexp"

var problemWords = map[string]int{
	"simile":    3,
	"forever":   3,
	"shoreline": 2,
	"forest":    2,
}

var (
	consonantsRegexp = regexp.MustCompile("[^aeiouy]+")
)

var subSyllables = [...]*regexp.Regexp{
	regexp.MustCompile("cial"),
	regexp.MustCompile("tia"),
	regexp.MustCompile("cius"),
	regexp.MustCompile("cious"),
	regexp.MustCompile("giu"),
	regexp.MustCompile("ion"),
	regexp.MustCompile("ise"),
	regexp.MustCompile("iou"),
	regexp.MustCompile("sia$"),
	regexp.MustCompile("[^aeiouyt]{2,}ed$"),
	regexp.MustCompile(".ely$"),
	regexp.MustCompile("[cg]h?e[rsd]?$"),
	regexp.MustCompile("rved?$"),
	regexp.MustCompile("[aeiouy][dt]es?$"),
	regexp.MustCompile("[aeiouy][^aeiouydt]e[rsd]?$"),
	regexp.MustCompile("^[dr]e[aeiou][^aeiou]+$"),
	regexp.MustCompile("[aeiouy]rse$"),
}

// AddSyllables are syllables that would be counted as one but should be two
var addSyllables = [...]*regexp.Regexp{
	regexp.MustCompile("ia"),
	regexp.MustCompile("riet"),
	regexp.MustCompile("dien"),
	regexp.MustCompile("iu"),
	regexp.MustCompile("io"),
	regexp.MustCompile("ii"),
	regexp.MustCompile("[aeiouym]bl$"),
	regexp.MustCompile("[aeiou]{3}"),
	regexp.MustCompile("^mc"),
	regexp.MustCompile("ism$"),
	regexp.MustCompile("[^aeiouy]{2}l$"),
	regexp.MustCompile("[^l]lien"),
	regexp.MustCompile("^coa[dglx]."),
	regexp.MustCompile("[^gq]ua[^auieo]"),
	regexp.MustCompile("dnt$"),
	regexp.MustCompile("uity$"),
	regexp.MustCompile("ie(r|st)$"),
	regexp.MustCompile("yee$"),
}

// PrefixSuffixes are single syllable prefixes and suffixes
var prefixSuffixes = [...]*regexp.Regexp{
	regexp.MustCompile("^un"),
	regexp.MustCompile("^fore"),
	regexp.MustCompile("ly$"),
	regexp.MustCompile("less$"),
	regexp.MustCompile("ful$"),
	regexp.MustCompile("ers?$"),
	regexp.MustCompile("ings?$"),
}
