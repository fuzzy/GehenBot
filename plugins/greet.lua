
register_handler("JOIN", "greet")

function greet(nick, event, target, args)
	privmsg(target, string.format("Hello %s, you fucknut.", nick))
end
