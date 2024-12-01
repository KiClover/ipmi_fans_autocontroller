package ipmi

import (
	"bytes"
	"os/exec"
	"serverTemperature/model"
	"strconv"
)

// DecimalToHex 10转16 0-100
func DecimalToHex(decimal int) string {
	info := strconv.FormatInt(int64(decimal), 16)
	info = "0x" + info
	return info
}

func FansControl(speed int, conf model.Config) error {
	speeds := DecimalToHex(speed)
	cmd := exec.Command("ipmitool", "-I lanplus",
		"-H "+conf.Ipmi.Host,
		"-U "+conf.Ipmi.User,
		"-P "+conf.Ipmi.Pwd, "raw", "0x3a", "0x01",
		speeds, speeds, speeds, speeds, speeds, speeds, speeds, speeds)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return err
	}
	// output := out.String()
	return nil
}
