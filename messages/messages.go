package messages

const (
	CMDHelp            = "/help"
	CMDListProxies     = "/list_proxies"
	CMDCheckProxies    = "/check_proxies"
	CMDUpdatePasswords = "/update_passwords"
)

const (
	Help = `Hello there. I'm BeletVideo platform's manager bot.
I can do several jobs.
1. Check working socks5h proxies - /check_proxies
2. List socks5h proxies - /list_proxies
3. Update service passwords - /update_passwords

Have a good day!`
	UpdatePasswords = "Unfortunately, this functionality on development process"
	UnknownCMD      = "Unknown command 🤔"
)
