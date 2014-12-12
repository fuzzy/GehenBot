
register_handler("PRIVMSG", "mycallback")

function mycallback(nick, event, target, args)
	print("nick:", nick)
	print("event:", event)
	print("target:", target)
	print("args:", args)
end
