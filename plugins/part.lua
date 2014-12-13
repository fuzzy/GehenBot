
register_handler("PRIVMSG", "partChan")

function partChan(nick, event, target, args)
	local haveCmd = 0
	for token in string.gmatch(args, "[^%s]+") do
		if haveCmd == 0 then
			if token == "!part" then
				haveCmd = 1
			end
		elseif haveCmd == 1 then
			part_channel(token)
			haveCmd = 2
		end
	end
end

