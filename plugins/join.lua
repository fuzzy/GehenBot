
register_handler("PRIVMSG", "joinChan")

function joinChan(nick, event, target, args)
	for token in string.gmatch(args, "[^%s]+") do
		print(token)
	end
	print(string.gmatch(args, "[^%s]+")[0])
	if string.gmatch(args, "[^%s]+")[0] == "!join" then
		print("Well, it matched.")
		join_channel(string.gmatch(args, "[^%s]+")[1])
	end
end
