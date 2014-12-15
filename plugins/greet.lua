
register("JOIN", "greet")

function greet(nick, event, target, args)
	mnick = string.format("%s", nick)
	bnick = string.format("%s", mynick())

	if mnick ~= bnick then
		say(target, string.format("Hello %s, you...you have a sister as well.", nick))
	end
	
end
