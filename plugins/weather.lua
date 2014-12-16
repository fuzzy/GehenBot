
function parseWeather(rawdata)
	local json = dofile(string.format("%s/dkjson.lua", scriptDir()))
	local data, pos, err = json.decode(rawdata, 1, nil)
	if err then
		print("Error:", err)
	end
	return(string.format("%s: %s, %sF/%sC (H:%sF/%sC L:%sF/%sC), Humidity: %s, Wind: %s %smph/%skph", 
		data.current_observation.display_location.full,
		data.forecast.simpleforecast.forecastday[1].conditions,
		data.current_observation.temp_f,
		data.current_observation.temp_c,
		data.forecast.simpleforecast.forecastday[1].high.fahrenheit,
		data.forecast.simpleforecast.forecastday[1].high.celsius,
		data.forecast.simpleforecast.forecastday[1].low.fahrenheit,
		data.forecast.simpleforecast.forecastday[1].low.celsius,
		data.current_observation.relative_humidity,
		data.forecast.simpleforecast.forecastday[1].avewind.dir,
		data.forecast.simpleforecast.forecastday[1].avewind.mph,
		data.forecast.simpleforecast.forecastday[1].avewind.kph
	))	
end

function weather(nick, event, target, args)
	local http          = require("socket.http")
	local buri          = "http://api.wunderground.com/api"
	local akey          = "407583fb16ebf33c"
	local wtype         = "forecast/conditions/q"
	local uri           = string.format("%s/%s/%s", buri, akey, wtype)
	local haveCmd       = 0
	
	for token in string.gmatch(args, "[^%s]+") do
		if haveCmd == 0 then
			if token == string.format("%swe", commandChar()) then
				haveCmd = 1
			end
		elseif haveCmd == 1 then
			zipc        = token
			haveCmd     = 2
			uri         = uri .. "/" .. zipc .. ".json"
			rd, e       = http.request(uri)
			say(target, nick .. ": " .. parseWeather(rd))
		end
	end
	
end

register("PRIVMSG", "weather")