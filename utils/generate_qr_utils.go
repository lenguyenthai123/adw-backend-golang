package utils

import (
	"bytes"
	"github.com/skip2/go-qrcode"
	"image"
)

func GenerateQRCode(text string) (image.Image, error) {
	// Generate QR code
	qr, err := qrcode.Encode(text, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}

	// Create image from QR code bytes
	img, _, err := image.Decode(bytes.NewReader(qr))
	if err != nil {
		return nil, err
	}

	return img, nil
}
