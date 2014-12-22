
function isFile(path)
	 local f = io.open(path, "r")
	 if f ~= nil then
			io.close(f)
			return true
	 else
			return false
	 end
end

function haveLink(link)
	 if isFile("/tmp/gehenbot.links") == true then
			for line in io.lines("/tmp/gehenbot.links") do
				 if line == link then
						return true
				 end
			end
	 end
	 return false
end

function findLink(tg, words)
	 for line in io.lines("/tmp/gehenbot.links") do
			for token in string.gmatch(words, "[^%s]") do
				 if string.match(line, "(.*" .. token .. ".*)") then
						say(tg, line)
				 end
			end
	 end
end

function storeLink(link)
	 file = io.open("/tmp/gehenbot.links", "a+")
	 io.output(file)
	 io.write(link .. "\n")
	 io.close(file)
end

--function mynick()
--	 return "GehenBoto"
--end

--function commandChar()
--	 return "!"
--end

--function say(a, b)
--	 print(a, b)
--end

function linkGrab(nick, event, target, args)
	 first  = false
	 second = false
	 search = false
	 add    = false
	 del    = false

	 if target ~= mynick() then
			out = target
	 else
			out = nick
	 end
	 
	 for token in string.gmatch(args, "[^%s]+") do
			if first == false then
				 if token == string.format("%slink", commandChar()) then
						first = true
				 else
						if string.match(token, "[(http|https)]://.*") then
							 if haveLink(token) == false then
									storeLink(token)
									say(out, nick .. ": ok, I've added: " .. token)
							 end
						end
				 end
			elseif first == true and second == false then
				 if token == "search" then
						second = true
						search = true
				 elseif token == "del" then
						second = true
						del    = true
				 end
			elseif first == true and second == true then
				 if del == true then
						say(out, "To be honest, I haven't gotten this part done yet.")
				 elseif search == true then
						findLink(out, token)
				 end
			end
	 end
end

register("PRIVMSG", "linkGrab")
