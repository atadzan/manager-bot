package messages

const (
	CMDStart           = "start"
	CMDHelp            = "help"
	CMDListProxies     = "list_proxies"
	CMDCheckProxies    = "check_proxies"
	CMDUpdatePasswords = "update_passwords"
)

const (
	Help = `Hello 👋.
I'm BeletVideo's manager bot. I can do several jobs.

1. Check working socks5h proxies - /check_proxies
2. List socks5h proxies - /list_proxies
3. Update service passwords - /update_passwords

You can access commands in the lower left corner

Good luck 👍`
	UpdatePasswords = "Unfortunately, this functionality on development process"
	UnknownCMD      = "Unknown command 🤔"
)
