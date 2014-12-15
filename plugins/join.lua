
register("PRIVMSG", "joinChan")

function joinChan(nick, event, target, args)
	local haveCmd = 0
	for token in string.gmatch(args, "[^%s]+") do
		if haveCmd == 0 then
			if token == "!join" then
				haveCmd = 1
			end
		elseif haveCmd == 1 then
			join(token)
			haveCmd = 2
		end
	end
end
