package tokenizer

const (
	// templates
	TemplateOpen                      = "{{"
	TemplateParamSeparator            = "|"
	TemplateParamSeparatorWithEscapes = `\|`
	TemplateParamEquals               = "="
	TemplateClose                     = "}}"
	TemplateParamOpen                 = "{{{"
	TemplateParamClose                = "}}}"
	TemplateParamClose4               = "}}}}" // corner case, should be broken into }}, }}
	TemplateAsteriskInList            = "*"

	// wikilinks
	WikilinkOpen      = "[["
	WikilinkSeparator = "|"
	WikilinkClose     = "]]"
	ExternalLinkOpen  = `[`
	ExternalLinkClose = `]`

	// comments
	CommentStart = "<!--"
	CommentEnd   = "-->"

	// tags
	TagRefOpen     = "<ref>"
	TagRefOpen1    = "<ref"
	TagRefClose    = "</ref>"
	TagClose       = "/>"
	TagBreak       = "<br>"
	TagBreak1      = "<br/>"
	TagBreak2      = "<br />"
	TagGreaterThan = ">"
	TagLessThan    = "<"
	TagListItem    = "<li>"
	TagSmallOpen   = "<small>"
	TagSmallClose  = "</small>"
	// format
	Italic = `''`
	Bold   = `'''`

	// text
	TextWithEscapes = `[^{}\[\]\|='<>\*]*`
)
