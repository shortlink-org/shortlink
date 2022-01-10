package repl

import (
	"github.com/c-bata/go-prompt"
)

var suggestions = []prompt.Suggest{
	// Command =========================================================================================================
	{Text: ".help", Description: "Help snippet"},
	{Text: ".open", Description: "Select database"},
	{Text: ".save", Description: "Save payload from this session in database"},
	{Text: ".close", Description: "Close this session"},

	// SQL =============================================================================================================
	// Table -----------------------------------------------------------------------------------------------------------
	{Text: "create table", Description: "create new table"},
	{Text: "drop table", Description: "drop table"},
}
