package validator

import (
	"errors"
	"strings"
)

// ProductCode validates the XXXX-XXXX-XXXX-XXXX format.
//   - Exactly 19 characters (16 alphanumeric + 3 dashes)
//   - Dashes at positions 4, 9, 14
//   - Only uppercase A-Z and digits 0-9
func ProductCode(code string) error {
	code = strings.ToUpper(strings.TrimSpace(code))

	if len(code) != 19 {
		return errors.New("รหัสสินค้าต้องมี 16 หลัก รูปแบบ XXXX-XXXX-XXXX-XXXX")
	}
	if code[4] != '-' || code[9] != '-' || code[14] != '-' {
		return errors.New("รูปแบบไม่ถูกต้อง ต้องเป็น XXXX-XXXX-XXXX-XXXX")
	}

	for i, ch := range code {
		if i == 4 || i == 9 || i == 14 {
			continue
		}
		if !isAlphanumUpper(ch) {
			return errors.New("ใช้ได้เฉพาะตัวเลข (0-9) และตัวอักษรพิมพ์ใหญ่ (A-Z)")
		}
	}
	return nil
}

func isAlphanumUpper(r rune) bool {
	return (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}
