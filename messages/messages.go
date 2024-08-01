package messages

const (
	CMDStart           = "start"
	CMDHelp            = "help"
	CMDListProxies     = "list_proxies"
	CMDCheckProxies    = "check_proxies"
	UpdateProxies      = "update_proxies"
	CMDUpdatePasswords = "update_passwords"
	CMDClearProxyList  = "clear_list"
)

const (
	Help = `Hello 👋.
I'm BeletVideo's manager bot. I can do several jobs.

1. Check working socks5h proxies - /check_proxies
2. List socks5h proxies - /list_proxies
3. Update proxies - /update_proxies
4. Clear proxy list - /clear_list
3. Update service passwords - /update_passwords

You can access commands in the lower left corner

Good luck 👍`
	UpdatePasswords              = "Unfortunately, this functionality on development process 😢"
	UpdateProxiesMsg             = "Please provide proxies ⬇️"
	UnknownCMD                   = "Unknown command 🤔"
	ProxiesSuccessfullyUpdated   = "Proxies successfully updated ✅"
	ProxyListSuccessfullyCleared = "The proxy list has been successfully cleared ✅"
)

type Proxy struct {
	URL         string `json:"URL"`
	CountryCode string `json:"countryCode"`
}
