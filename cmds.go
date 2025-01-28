package wine

const (
	ServerDebug      = "--debug"
	ServerForeground = "--foreground"
	ServerKill       = "--kill"
	ServerPersistent = "--persistent"
	ServerWait       = "--wait"
)

const (
	BootEndSession = "--end-session"
	BootForceExit  = "--force"
	BootInit       = "--init"
	BootKill       = "--kill"
	BootRestart    = "--restart"
	BootShutdown   = "--shutdown"
	BootUpdate     = "--update"
)

// Tricks returns a [Cmd] for wineserver.
func (p *Prefix) Server(args ...string) *Cmd {
	return p.Command("wineserver", args...)
}

// Tricks returns a [Cmd] for wineboot.
func (p *Prefix) Boot(args ...string) *Cmd {
	return p.Wine("wineboot", args...)
}

// Kill returns a [Cmd] for killing the Wineprefix.
func (p *Prefix) Kill() *Cmd {
	return p.Server(ServerKill)
}

// Init returns a [Cmd] for initializating the Wineprefix.
func (p *Prefix) Init() *Cmd {
	return p.Boot(BootInit)
}

// Update returns a [Cmd] for updating the Wineprefix.
func (p *Prefix) Update() *Cmd {
	return p.Boot(BootUpdate)
}

// Tricks returns a [Cmd] for winetricks.
func (p *Prefix) Tricks() *Cmd {
	if p.IsProton() {
		// umu-run [winetricks [ARG...]]
		cmd := p.Wine("winetricks")
		if cmd.Args[0] == "umu-run" {
			return cmd
		}
		// fallback to regular winetricks
	}

	cmd := p.Command("winetricks")
	cmd.Env = append(cmd.Environ(),
		"WINE="+p.bin("wine64"),
		"WINESERVER="+p.bin("wineserver"),
	)

	return cmd
}
