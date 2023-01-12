package filesig

import "regexp"

var (
	AVIF       = []byte{0x00, 0x00, 0x00}
	BMP        = []byte{0x42, 0x4D}
	DIB_0      = []byte{0x42, 0x4D}
	DIB_1      = []byte{0x28, 0x00}
	GIF        = []byte{0x47, 0x49, 0x46, 0x38}
	TIFF       = []byte{0x49, 0x49, 0x2A, 0x00}
	MP3        = []byte{0x49, 0x44, 0x33}
	MPG_0      = []byte{0x00, 0x00, 0x01, 0xB3}
	MPG_1      = []byte{0x00, 0x00, 0x01, 0xBA}
	FLV        = []byte{0x46, 0x4C, 0x56}
	APK        = []byte{0x50, 0x4B, 0x03, 0x04}
	MS_OFFICE  = []byte{0x50, 0x4B, 0x03, 0x04}
	JAR        = []byte{0x50, 0x4B, 0x03, 0x04}
	SWF_0      = []byte{0x43, 0x57, 0x53}
	SWF_1      = []byte{0x46, 0x57, 0x53}
	PNG        = []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}
	JPEG       = []byte{0xFF, 0xD8}
	MP4_0      = []byte{0x00, 0x00, 0x00, 0x14, 0x66, 0x74, 0x79, 0x70, 0x69, 0x73, 0x6F, 0x6D}
	MP4_1      = []byte{0x00, 0x00, 0x00, 0x18, 0x66, 0x74, 0x79, 0x70}
	MP4_2      = []byte{0x00, 0x00, 0x00, 0x1C, 0x66, 0x74, 0x79, 0x70}
	THREE_GP_0 = []byte{0x00, 0x00, 0x00, 0x20, 0x66, 0x74, 0x79, 0x70}
	THREE_GP_1 = []byte{0x00, 0x00, 0x00, 0x20, 0x66, 0x74, 0x79, 0x70}
	THREE_GP_2 = []byte{0x00, 0x00, 0x00, 0x14, 0x66, 0x74, 0x79, 0x70}
	MKV        = []byte{0x1A, 0x45, 0xDF, 0xA3}
	PDF        = []byte{0x25, 0x50, 0x44, 0x46}
	RAR        = []byte{0x52, 0x61, 0x72, 0x21, 0x1A, 0x07, 0x00}
	SVG        = []byte{0x3C, 0x3F, 0x78, 0x6D, 0x6C}
	GZIP       = []byte{0x1F, 0x8B, 0x08}
	ZIP_0      = []byte{0x50, 0x4B, 0x03, 0x04}
	ZIP_1      = []byte{0x50, 0x4B, 0x05, 0x06}
	ZIP_2      = []byte{0x50, 0x4B, 0x07, 0x08}
	WEBP       = []byte{0x52, 0x49, 0x46, 0x46}

	HtmlCommentRegex = regexp.MustCompile(`(?i)<!--([\s\S]*?)-->`)
	SvgRegex         = regexp.MustCompile(`(?i)^\s*(?:<\?xml[^>]*>\s*)?(?:<!doctype svg[^>]*>\s*)?<svg[^>]*>[^*]*<\/svg>\s*$`)
)
