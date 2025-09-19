// Package utils provides utility functions for the ModelsLab SDK
package utils

import (
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// ImageToBase64 converts an image file to base64 string
func ImageToBase64(imagePath string) (string, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return "", fmt.Errorf("failed to open image file: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read image file: %w", err)
	}

	return base64.StdEncoding.EncodeToString(data), nil
}

// Base64ToImage converts a base64 string to an image and saves it to a file
func Base64ToImage(base64Str, outputPath string) error {
	data, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return fmt.Errorf("failed to decode base64 string: %w", err)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write image data: %w", err)
	}

	return nil
}

// SaveImageToFile saves an image.Image to a file
func SaveImageToFile(img image.Image, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer file.Close()

	ext := strings.ToLower(filepath.Ext(outputPath))
	switch ext {
	case ".png":
		err = png.Encode(file, img)
	case ".jpg", ".jpeg":
		err = jpeg.Encode(file, img, &jpeg.Options{Quality: 95})
	default:
		return fmt.Errorf("unsupported image format: %s", ext)
	}

	if err != nil {
		return fmt.Errorf("failed to encode image: %w", err)
	}

	return nil
}

// ReadImageFromFile reads an image from a file
func ReadImageFromFile(imagePath string) (image.Image, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open image file: %w", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %w", err)
	}

	return img, nil
}
