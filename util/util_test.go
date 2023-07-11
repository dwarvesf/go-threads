package thread

import (
	"encoding/hex"
	"fmt"
	"strings"
	"testing"
)

func TestGenerateAndroidDeviceID(t *testing.T) {
	t.Run("Generated device ID starts with 'android-'", func(t *testing.T) {
		deviceID := GenerateAndroidDeviceID()
		if !strings.HasPrefix(deviceID, "android-") {
			t.Errorf("Expected device ID to start with 'android-', got %s", deviceID)
		}
	})

	t.Run("Generated device ID has length 24", func(t *testing.T) {
		deviceID := GenerateAndroidDeviceID()
		fmt.Println(deviceID)
		fmt.Println(len(deviceID))
		if len(deviceID) != 24 {
			t.Errorf("Expected device ID length to be 24, got %d", len(deviceID))
		}
	})

	t.Run("Generated device ID is a valid SHA-256 hash", func(t *testing.T) {
		deviceID := GenerateAndroidDeviceID()
		deviceID = strings.TrimPrefix(deviceID, "android-")
		_, err := hex.DecodeString(deviceID)
		if err != nil {
			t.Errorf("Expected device ID to be a valid SHA-256 hash, got %s", deviceID)
		}
	})

	t.Run("Generated device IDs are unique", func(t *testing.T) {
		deviceID1 := GenerateAndroidDeviceID()
		deviceID2 := GenerateAndroidDeviceID()
		if deviceID1 == deviceID2 {
			t.Error("Expected generated device IDs to be unique")
		}
	})
}
