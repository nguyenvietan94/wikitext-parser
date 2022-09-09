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
	TagRefOpen  = "<ref>"
	TagRefClose = "</ref>"
	TagBreak    = "<br>"

	// text
	TextWithEscapes = `[^{}\[\]\|='<>\*]*`

	// format
	Italic = `''`
	Bold   = `'''`
)
