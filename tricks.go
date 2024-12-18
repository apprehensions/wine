package wine

import (
	"strconv"
)

func (p *Prefix) Winetricks() error {
	return p.Command("winetricks").Run()
}

func (p *Prefix) SetDPI(dpi int) error {
	return p.RegistryAdd("HKEY_CURRENT_USER\\Control Panel\\Desktop", "LogPixels", REG_DWORD, strconv.Itoa(dpi))
}
