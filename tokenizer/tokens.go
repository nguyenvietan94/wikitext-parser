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
	TemplateAsteriskInList            = "*"

	// wikilinks
	WikilinkOpen                 = "[["
	WikilinkOpenWithEscapes      = `\[\[`
	WikilinkSeparator            = "|"
	WikilinkSeparatorWithEscapes = `\|`
	WikilinkClose                = "]]"
	WikilinkCloseWithEscapes     = `\]\]`

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
