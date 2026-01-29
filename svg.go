/*
Filename: svg.go
Description: handles generating the SVG badge
Created by: osh
        at: 16:46 on Wednesday, the 28th of January, 2026.
Last edited 13:35 on Thursday, the 29th of January, 2026
*/

package main

import (
	"encoding/base64"
	"fmt"
)

// character widths for Verdana at 110px
// generated using https://github.com/metabolize/anafanafo/tree/main/packages/char-width-table-builder
var charWidths = [384]float64{
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // 0-15
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // 16-31
	38.67, 43.29, 50.49, 90.02, 69.93, 118.38, 79.92, 29.54, 49.95, 49.95, 69.93, 90.02, 40.01, 49.95, 40.01, 49.95, // ' '-'/'
	69.93, 69.93, 69.93, 69.93, 69.93, 69.93, 69.93, 69.93, 69.93, 69.93, 49.95, 49.95, 90.02, 90.02, 90.02, 60, // '0'-'?'
	110, 75.2, 75.41, 76.81, 84.76, 69.56, 63.22, 85.29, 82.66, 46.3, 50, 76.22, 61.23, 92.71, 82.29, 86.58, // '@'-'O'
	66.33, 86.58, 76.48, 75.2, 67.78, 80.51, 75.2, 108.76, 75.36, 67.68, 75.36, 49.95, 49.95, 49.95, 90.02, 69.93, // 'P'-'_'
	69.93, 66.06, 68.54, 57.31, 68.54, 65.53, 38.67, 68.54, 69.61, 30.19, 37.87, 65.1, 30.19, 106.99, 69.61, 66.76, // '`'-'o'
	68.54, 68.54, 46.94, 57.31, 43.34, 69.61, 65.1, 90.02, 65.1, 65.1, 57.79, 69.82, 49.95, 69.82, 90.02, 41.36, // 'p'-127
	61.18, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, // 128-143
	110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, // 144-159
	38.67, 43.29, 69.93, 69.93, 69.93, 69.93, 49.95, 69.93, 69.93, 110, 60, 70.9, 90.02, 0, 110, 69.93, // 160-175
	59.62, 90.02, 59.62, 59.62, 69.93, 70.58, 69.93, 40.01, 69.93, 59.62, 60, 70.9, 110, 110, 110, 60, // 176-191
	75.2, 75.2, 75.2, 75.2, 75.2, 75.2, 108.28, 76.81, 69.56, 69.56, 69.56, 69.56, 46.3, 46.3, 46.3, 46.3, // 192-207
	85.29, 82.29, 86.58, 86.58, 86.58, 86.58, 86.58, 90.02, 86.58, 80.51, 80.51, 80.51, 80.51, 67.68, 66.6, 68.21, // 208-223
	66.06, 66.06, 66.06, 66.06, 66.06, 66.06, 105.06, 57.31, 65.53, 65.53, 65.53, 65.53, 30.19, 30.19, 30.19, 30.19, // 224-239
	67.3, 69.61, 66.76, 66.76, 66.76, 66.76, 66.76, 90.02, 66.76, 69.61, 69.61, 69.61, 69.61, 65.1, 68.54, 65.1, // 240-255
	75.2, 66.06, 75.2, 66.06, 75.2, 66.06, 76.81, 57.31, 76.81, 57.31, 76.81, 57.31, 76.81, 57.31, 84.76, 71.22, // 256-271
	85.29, 68.54, 69.56, 65.53, 69.56, 65.53, 69.56, 65.53, 69.56, 65.53, 69.56, 65.53, 85.29, 68.54, 85.29, 68.54, // 272-287
	85.29, 68.54, 85.29, 68.54, 82.66, 69.61, 82.66, 69.61, 46.3, 30.19, 46.3, 30.19, 46.3, 30.19, 46.3, 30.19, // 288-303
	46.3, 30.19, 95.77, 67.51, 50, 37.87, 76.22, 65.1, 65.1, 61.23, 30.19, 61.23, 30.19, 61.23, 32.55, 61.23, // 304-319
	50.43, 61.77, 31.26, 82.29, 69.61, 82.29, 69.61, 82.29, 69.61, 80.35, 82.29, 69.61, 86.58, 66.76, 86.58, 66.76, // 320-335
	86.58, 66.76, 117.68, 107.96, 76.48, 46.94, 76.48, 46.94, 76.48, 46.94, 75.2, 57.31, 75.2, 57.31, 75.2, 57.31, // 336-351
	75.2, 57.31, 67.78, 43.34, 67.78, 43.34, 67.78, 43.34, 80.51, 69.61, 80.51, 69.61, 80.51, 69.34, 80.51, 69.61, // 352-367
	80.51, 69.61, 80.51, 69.34, 108.76, 90.02, 67.68, 65.1, 67.68, 75.36, 57.79, 75.36, 57.79, 75.36, 57.79, 33.03, // 368-383
}

func calculateTextWidth(text string) float64 {
	var width float64
	for _, char := range text {
		if char >= 0 && char < 384 {
			width += charWidths[rune(char)]
		} else {
			width += charWidths[rune('?')]
		}
	}

	// convert to 11px scale, truncate, round up to odd, then back to 110px scale (like shields.io)
	width11px := int(width / 10)
	if width11px%2 == 0 {
		width11px++
	}

	fmt.Println("calculated width:", width11px*10)
	return float64(width11px * 10)
}

const (
	height          = "20"
	labelBg         = "#000000"
	contentBg       = "#1bd96a"
	padding         = 50.0
	iconWidth       = 14.0
	iconX           = 5.0
	iconY           = 3.0
	iconPadding     = 3.0
	modrinthLogoSVG = `<svg fill="#1bd96a" role="img" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path d="M12.252.004a11.78 11.768 0 0 0-8.92 3.73 11 10.999 0 0 0-2.17 3.11 11.37 11.359 0 0 0-1.16 5.169c0 1.42.17 2.5.6 3.77.24.759.77 1.899 1.17 2.529a12.3 12.298 0 0 0 8.85 5.639c.44.05 2.54.07 2.76.02.2-.04.22.1-.26-1.7l-.36-1.37-1.01-.06a8.5 8.489 0 0 1-5.18-1.8 5.34 5.34 0 0 1-1.3-1.26c0-.05.34-.28.74-.5a37.572 37.545 0 0 1 2.88-1.629c.03 0 .5.45 1.06.98l1 .97 2.07-.43 2.06-.43 1.47-1.47c.8-.8 1.48-1.5 1.48-1.52 0-.09-.42-1.63-.46-1.7-.04-.06-.2-.03-1.02.18-.53.13-1.2.3-1.45.4l-.48.15-.53.53-.53.53-.93.1-.93.07-.52-.5a2.7 2.7 0 0 1-.96-1.7l-.13-.6.43-.57c.68-.9.68-.9 1.46-1.1.4-.1.65-.2.83-.33.13-.099.65-.579 1.14-1.069l.9-.9-.7-.7-.7-.7-1.95.54c-1.07.3-1.96.53-1.97.53-.03 0-2.23 2.48-2.63 2.97l-.29.35.28 1.03c.16.56.3 1.16.31 1.34l.03.3-.34.23c-.37.23-2.22 1.3-2.84 1.63-.36.2-.37.2-.44.1-.08-.1-.23-.6-.32-1.03-.18-.86-.17-2.75.02-3.73a8.84 8.839 0 0 1 7.9-6.93c.43-.03.77-.08.78-.1.06-.17.5-2.999.47-3.039-.01-.02-.1-.02-.2-.03Zm3.68.67c-.2 0-.3.1-.37.38-.06.23-.46 2.42-.46 2.52 0 .04.1.11.22.16a8.51 8.499 0 0 1 2.99 2 8.38 8.379 0 0 1 2.16 3.449 6.9 6.9 0 0 1 .4 2.8c0 1.07 0 1.27-.1 1.73a9.37 9.369 0 0 1-1.76 3.769c-.32.4-.98 1.06-1.37 1.38-.38.32-1.54 1.1-1.7 1.14-.1.03-.1.06-.07.26.03.18.64 2.56.7 2.78l.06.06a12.07 12.058 0 0 0 7.27-9.4c.13-.77.13-2.58 0-3.4a11.96 11.948 0 0 0-5.73-8.578c-.7-.42-2.05-1.06-2.25-1.06Z"/></svg>`
)

var modrinthLogoHref = "data:image/svg+xml;base64," + base64.StdEncoding.EncodeToString([]byte(modrinthLogoSVG))

// GenerateBadge builds an SVG string with an optional logo, optional label, and optional content
func GenerateBadge(logo bool, label string, content string) (svg string) {
	labelTextWidth := calculateTextWidth(label)
	contentTextWidth := calculateTextWidth(content)

	/// calculate widths using the generated table, and rescaling
	// label and content
	var leftWidth float64
	if logo && label == "" {
		leftWidth = 2*padding/10 + iconWidth
	} else if logo {
		leftWidth = iconX + iconWidth + iconPadding + labelTextWidth/10 + padding/10
	} else if label != "" {
		leftWidth = (labelTextWidth + padding*2) / 10
	}

	var rightWidth float64
	if content != "" {
		rightWidth = (contentTextWidth + padding*2) / 10
	}
	totalWidth := leftWidth + rightWidth

	// logo
	var leftCentre float64
	if logo && label != "" {
		leftCentre = (iconWidth+iconPadding+1)*10 + labelTextWidth/2 + padding
	} else if label != "" {
		leftCentre = padding + labelTextWidth/2
	}

	var rightCentre float64
	if content != "" {
		if leftWidth == 0 {
			rightCentre = padding + contentTextWidth/2
		} else {
			rightCentre = (leftWidth-1)*10 + padding + contentTextWidth/2
		}
	}

	title := label + ": " + content

	/// build the svgo
	// aria labelling + the gradient + rounded corners
	svg += fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" width="%.0f" height="%s" role="img" aria-label="%s">`, totalWidth, height, title)
	svg += `<title>` + title + `</title>`
	svg += `<linearGradient id="s" x2="0" y2="100%"><stop offset="0" stop-color="#bbb" stop-opacity=".1"/><stop offset="1" stop-opacity=".1"/></linearGradient>`
	svg += fmt.Sprintf(`<clipPath id="r"><rect width="%.0f" height="%s" rx="3" fill="#fff"/></clipPath>`, totalWidth, height)

	// left and right boxes
	svg += `<g clip-path="url(#r)">`
	if leftWidth > 0 {
		svg += fmt.Sprintf(`<rect width="%.0f" height="%s" fill="%s"/>`, leftWidth, height, labelBg)
	}
	if content != "" {
		svg += fmt.Sprintf(`<rect x="%.0f" width="%.0f" height="%s" fill="%s"/>`, leftWidth, rightWidth, height, contentBg)
	}
	svg += fmt.Sprintf(`<rect width="%.0f" height="%s" fill="url(#s)"/>`, totalWidth, height)
	svg += `</g>`

	// contains the logo, and both texts
	svg += `<g fill="#fff" text-anchor="middle" font-family="Verdana,Geneva,DejaVu Sans,sans-serif" text-rendering="geometricPrecision" font-size="110">`

	// add the logo if enabled
	if logo {
		svg += fmt.Sprintf(`<image x="%.0f" y="%.0f" width="%.0f" height="%.0f" href="%s"/>`, iconX, iconY, iconWidth, iconWidth, modrinthLogoHref)
	}

	// label text (shadow + main)
	if label != "" {
		svg += fmt.Sprintf(`<text aria-hidden="true" x="%.0f" y="150" fill="#010101" fill-opacity=".3" transform="scale(.1)" textLength="%.0f">%s</text>`, leftCentre, labelTextWidth, label)
		svg += fmt.Sprintf(`<text x="%.0f" y="140" transform="scale(.1)" fill="#fff" textLength="%.0f">%s</text>`, leftCentre, labelTextWidth, label)
	}

	// content text (shadow + main)
	if content != "" {
		svg += fmt.Sprintf(`<text aria-hidden="true" x="%.0f" y="150" fill="#010101" fill-opacity=".3" transform="scale(.1)" textLength="%.0f">%s</text>`, rightCentre, contentTextWidth, content)
		svg += fmt.Sprintf(`<text x="%.0f" y="140" transform="scale(.1)" fill="#fff" textLength="%.0f">%s</text>`, rightCentre, contentTextWidth, content)
	}

	svg += `</g></svg>`

	return svg
}
