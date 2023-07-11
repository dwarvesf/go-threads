package thread

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

const AndroidPrefix = "android"

func GenerateAndroidDeviceID() string {
	timestamp := strconv.FormatInt(time.Now().UnixMicro(), 10)
	hash := sha256.Sum256([]byte(timestamp))
	deviceID := hex.EncodeToString(hash[:16])[:16]
	return fmt.Sprintf("%s-%s", AndroidPrefix, deviceID)
}
