package i2pcontrol

import (
  "math/rand"
  "time"
)

const charset = "1234567890"

var seededRand *rand.Rand = rand.New(
  rand.NewSource(time.Now().UnixNano()))

func randomStringWithCharset(length int, charset string) string {
  b := make([]byte, length)
  for i := range b {
    b[i] = charset[seededRand.Intn(len(charset))]
  }
  return string(b)
}

func randomString(length int) string {
  return randomStringWithCharset(length, charset)
}
