package builtin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	oneFS = `Label: 'two'  uuid: 5efab9c9-55d8-4f0f-b5b9-4c521b567c70
	Total devices 2 FS bytes used 256.00KiB
	devid    1 size 1000.00MiB used 208.00MiB path /dev/loop1
	devid    2 size 1000.00MiB used 208.00MiB path /dev/loop2
	
Label: 'single'  uuid: 74595911-0f79-4c2e-925f-105d1279fb48
	Total devices 1 FS bytes used 192.00KiB
	devid    1 size 1000.00MiB used 132.00MiB path /dev/loop3`
)

func TestParseFS(t *testing.T) {
	fss, err := btrfsParseList([]byte(oneFS))
	assert.Equal(t, 2, len(fss))

	fs := fss[0]
	assert.Nil(t, err)
	assert.Equal(t, "two", fs.Label)
	assert.Equal(t, "5efab9c9-55d8-4f0f-b5b9-4c521b567c70", fs.UUID)
	assert.Equal(t, 2, fs.TotalDevices)
	assert.Equal(t, "256.00KiB", fs.Used)
	assert.Equal(t, 2, len(fs.Devices))
	dev := fs.Devices[0]
	assert.Equal(t, "1", dev.DevID)
	assert.Equal(t, "1000.00MiB", dev.Size)
	assert.Equal(t, "208.00MiB", dev.Used)
	assert.Equal(t, "/dev/loop1", dev.Path)
}