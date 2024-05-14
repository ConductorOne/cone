package output

import "github.com/pterm/pterm"

var InputNeeded = pterm.PrefixPrinter{
	MessageStyle: &pterm.Style{pterm.FgMagenta},
	Prefix: pterm.Prefix{
		Style: &pterm.Style{pterm.BgLightMagenta},
		Text:  "INPUT NEEDED",
	},
}
