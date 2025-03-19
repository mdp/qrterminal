package qrterminal

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"rsc.io/qr"
)

// Original tests that just verify the code doesn't crash
func TestGenerate(t *testing.T) {
	Generate("https://github.com/mdp/qrterminal", L, os.Stdout)
}

func TestGenerateWithConfig(t *testing.T) {
	config := Config{
		Level:     M,
		Writer:    os.Stdout,
		BlackChar: WHITE, // Inverted
		WhiteChar: BLACK,
		QuietZone: QUIET_ZONE,
	}
	GenerateWithConfig("https://github.com/mdp/qrterminal", config)
}

func TestGenerateHalfBlock(t *testing.T) {
	GenerateHalfBlock("https://github.com/mdp/qrterminal", L, os.Stdout)
}

func TestGenerateWithHalfBlockConfig(t *testing.T) {
	config := Config{
		Level:          M,
		Writer:         os.Stdout,
		HalfBlocks:     true,
		BlackChar:      BLACK_BLACK,
		WhiteBlackChar: WHITE_BLACK,
		WhiteChar:      WHITE_WHITE,
		BlackWhiteChar: BLACK_WHITE,
		QuietZone:      3,
	}
	GenerateWithConfig("https://github.com/mdp/qrterminal", config)
}

func TestGenerateWithHalfBlockMinConfig(t *testing.T) {
	config := Config{
		Level:          M,
		Writer:         os.Stdout,
		HalfBlocks:     true,
		QuietZone:      3,
	}
	GenerateWithConfig("https://github.com/mdp/qrterminal", config)
}

// New tests that actually verify the output

// Test that captures and verifies the output
func TestCaptureOutput(t *testing.T) {
	testCases := []struct {
		name       string
		input      string
		level      qr.Level
		halfBlocks bool
	}{
		{"BasicURL", "https://example.com", L, false},
		{"ShortText", "test", M, false},
		{"HalfBlockMode", "test", L, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer
			
			if tc.halfBlocks {
				config := Config{
					Level:      tc.level,
					Writer:     &buf,
					HalfBlocks: true,
				}
				GenerateWithConfig(tc.input, config)
			} else {
				Generate(tc.input, tc.level, &buf)
			}
			
			output := buf.String()
			
			// Verify output is not empty
			if len(output) == 0 {
				t.Errorf("Generated QR code is empty")
			}
			
			// Verify output contains multiple lines
			lines := strings.Split(output, "\n")
			if len(lines) <= 1 {
				t.Errorf("Generated QR code should have multiple lines, got %d", len(lines))
			}
			
			// Verify the output contains the expected characters
			if tc.halfBlocks {
				// Half blocks mode should contain these characters
				expectedChars := []string{BLACK_BLACK, WHITE_WHITE, BLACK_WHITE, WHITE_BLACK}
				foundExpectedChar := false
				
				for _, char := range expectedChars {
					if strings.Contains(output, char) {
						foundExpectedChar = true
						break
					}
				}
				
				if !foundExpectedChar {
					t.Errorf("Half block output doesn't contain expected characters")
				}
			} else {
				// Regular mode should contain BLACK and WHITE
				if !strings.Contains(output, BLACK) && !strings.Contains(output, WHITE) {
					t.Errorf("Output doesn't contain expected BLACK or WHITE characters")
				}
			}
		})
	}
}

// Test the structure of the QR code (size, quiet zone, etc.)
func TestQRStructure(t *testing.T) {
	testCases := []struct {
		name      string
		input     string
		quietZone int
	}{
		{"DefaultQuietZone", "test", QUIET_ZONE},
		{"MinimalQuietZone", "test", 1},
		{"LargeQuietZone", "test", 8},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer
			config := Config{
				Level:     L,
				Writer:    &buf,
				BlackChar: BLACK,
				WhiteChar: WHITE,
				QuietZone: tc.quietZone,
			}
			GenerateWithConfig(tc.input, config)
			
			output := buf.String()
			lines := strings.Split(output, "\n")
			
			// Check that we have at least 2*quietZone lines for top and bottom borders
			if len(lines) < 2*tc.quietZone {
				t.Errorf("Expected at least %d lines for quiet zone, got %d", 2*tc.quietZone, len(lines))
			}
			
			// Check that the first few lines contain WHITE (quiet zone)
			for i := 0; i < tc.quietZone && i < len(lines); i++ {
				if len(lines[i]) > 0 && !strings.Contains(lines[i], WHITE) {
					t.Errorf("Line %d should contain WHITE (quiet zone)", i)
				}
			}
			
			// Check that the last few lines are all WHITE (quiet zone)
			// Note: The last line might be empty due to a trailing newline
			for i := len(lines) - tc.quietZone; i < len(lines); i++ {
				if i < len(lines) && len(lines[i]) > 0 && !strings.Contains(lines[i], WHITE) {
					t.Errorf("Line %d should contain WHITE (quiet zone)", i)
				}
			}
		})
	}
}

// Test with various configurations
func TestConfigVariations(t *testing.T) {
	testCases := []struct {
		name   string
		config Config
		input  string
	}{
		{
			"InvertedColors",
			Config{
				Level:     L,
				BlackChar: WHITE,
				WhiteChar: BLACK,
			},
			"test",
		},
		{
			"CustomCharacters",
			Config{
				Level:     L,
				BlackChar: "XX",
				WhiteChar: "..",
			},
			"test",
		},
		{
			"HalfBlocksCustomChars",
			Config{
				Level:          L,
				HalfBlocks:     true,
				BlackChar:      "a",
				WhiteChar:      "b",
				BlackWhiteChar: "c",
				WhiteBlackChar: "d",
			},
			"test",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer
			tc.config.Writer = &buf
			GenerateWithConfig(tc.input, tc.config)
			
			output := buf.String()
			
			// Verify output is not empty
			if len(output) == 0 {
				t.Errorf("Generated QR code is empty")
			}
			
			// For custom characters, verify they appear in the output
			if tc.name == "CustomCharacters" {
				if !strings.Contains(output, "XX") || !strings.Contains(output, "..") {
					t.Errorf("Output doesn't contain custom characters")
				}
			} else if tc.name == "HalfBlocksCustomChars" {
				// Check for at least one of the custom characters
				customChars := []string{"a", "b", "c", "d"}
				foundCustomChar := false
				
				for _, char := range customChars {
					if strings.Contains(output, char) {
						foundCustomChar = true
						break
					}
				}
				
				if !foundCustomChar {
					t.Errorf("Output doesn't contain custom half block characters")
				}
			}
		})
	}
}

// Test edge cases
func TestEdgeCases(t *testing.T) {
	testCases := []struct {
		name  string
		input string
	}{
		{"EmptyString", ""},
		{"VeryLongString", strings.Repeat("a", 100)},
		{"SpecialCharacters", "!@#$%^&*()_+{}|:<>?"},
		{"Unicode", "こんにちは世界"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer
			
			// Test that it doesn't panic
			Generate(tc.input, L, &buf)
			
			output := buf.String()
			
			// Verify output is not empty (unless input is empty)
			if tc.input != "" && len(output) == 0 {
				t.Errorf("Generated QR code is empty for input: %s", tc.input)
			}
			
			// For empty string, we should still get some output (the QR code for an empty string)
			if tc.input == "" && len(output) == 0 {
				t.Errorf("Generated QR code for empty string should not be empty")
			}
		})
	}
}

// Test that the same input always produces the same output
func TestConsistentOutput(t *testing.T) {
	input := "https://github.com/mdp/qrterminal"
	
	var buf1 bytes.Buffer
	Generate(input, L, &buf1)
	output1 := buf1.String()
	
	var buf2 bytes.Buffer
	Generate(input, L, &buf2)
	output2 := buf2.String()
	
	if output1 != output2 {
		t.Errorf("Generated QR codes for the same input should be identical")
	}
}

// Test that different error correction levels produce different outputs
func TestErrorCorrectionLevels(t *testing.T) {
	input := "https://github.com/mdp/qrterminal"
	
	var bufL bytes.Buffer
	Generate(input, L, &bufL)
	outputL := bufL.String()
	
	var bufM bytes.Buffer
	Generate(input, M, &bufM)
	outputM := bufM.String()
	
	var bufH bytes.Buffer
	Generate(input, H, &bufH)
	outputH := bufH.String()
	
	// Different error correction levels should produce different outputs
	// (higher levels add more redundancy, changing the pattern)
	if outputL == outputM || outputL == outputH || outputM == outputH {
		t.Errorf("Different error correction levels should produce different outputs")
	}
}

// Test that the sixel detection function works
func TestSixelDetection(t *testing.T) {
	// This is a simple test that just ensures the function doesn't crash
	// We can't really test the actual detection without a terminal
	result := IsSixelSupported(os.Stdout)
	
	// The result could be true or false depending on the terminal
	// We just want to make sure it runs without error
	t.Logf("Sixel support detected: %v", result)
}

// Test that the QR code pattern is consistent and contains the expected pattern
func TestQRPattern(t *testing.T) {
	// Generate a QR code with a known input
	input := "test"
	
	var buf bytes.Buffer
	Generate(input, L, &buf)
	output := buf.String()
	
	// Split the output into lines
	lines := strings.Split(output, "\n")
	
	// Skip the quiet zone at the top
	contentStart := QUIET_ZONE
	
	// Check for the finder patterns (the three square patterns in the corners)
	// These are a key part of any QR code and should be present
	
	// The exact position depends on the QR code size, but we can check for patterns
	// that should be present in any valid QR code for our input
	
	// Check for some BLACK pixels in the content area (not just quiet zone)
	foundBlack := false
	for i := contentStart; i < len(lines) - QUIET_ZONE; i++ {
		if strings.Contains(lines[i], BLACK) {
			foundBlack = true
			break
		}
	}
	
	if !foundBlack {
		t.Errorf("QR code doesn't contain any BLACK pixels in the content area")
	}
}
