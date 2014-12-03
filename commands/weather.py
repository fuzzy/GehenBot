#!/usr/bin/env python

import os, sys, json, urllib2

def have_data(nick):
    if not os.path.isfile('/tmp/gehenbot-weather.db'): return False
    data = json.loads(open('/tmp/gehenbot-weather.db', 'r').read())
    if nick in data.keys(): return data[nick]
    else: return False

def forget_data(nick):
    if have_data(nick):
        if os.path.isfile('/tmp/gehenbot-weather.db'):
            data = json.loads(open('/tmp/gehenbot-weather.db', 'r').read())
            data.pop(nick)
            open('/tmp/gehenbot-weather.db', 'w+').write(json.dumps(data))

def save_data(nick, zipcode):
    if not have_data(nick):
        if os.path.isfile('/tmp/gehenbot-weather.db'): data = json.loads(open('/tmp/gehenbot-weather.db', 'r').read())
        else: data = {}
        data[nick] = zipcode
        open('/tmp/gehenbot-weather.db', 'w+').write(json.dumps(data))

def usage():
    print 'Usage: !weather (zipcode|@nick) <save|forget>'
    sys.exit(0)

if sys.argv[1] == '--usage':
    usage()

nick = sys.argv[1]
mask = sys.argv[2]
save = False
frgt = False

if len(sys.argv) == 4 and sys.argv[3] != "":
    targs = sys.argv[3].split()
    if len(targs) == 2:
        zipc = targs[0]
        if targs[1] == 'save': save = True
        if targs[1] == 'forget': frgt = True
    elif len(targs) == 1: zipc = targs[0]
    else: usage()
else:
    if have_data(nick): zipc = json.loads(open('/tmp/gehenbot-weather.db', 'r').read())[nick]
    else: usage()

if zipc[0] == '@': 
    nick = zipc[1:]
    zipc = have_data(nick)
    save = False
    frgt = False

if save: save_data(nick, zipc)

if frgt or zipc == 'forget':
    zipc = have_data(nick)
    if zipc: forget_data(nick)

data = json.loads(urllib2.urlopen('http://api.wunderground.com/api/407583fb16ebf33c/forecast/conditions/q/%s.json' % zipc).read())
print '%s: %s: %s, %sF/%sC (H:%sF/%sC L:%sF/%sC), Humidity: %s, Wind: %smph/%skph' % (
    nick,
    data['current_observation']['display_location']['full'],
    data['forecast']['simpleforecast']['forecastday'][0]['conditions'],
    data['current_observation']['temp_f'],
    data['current_observation']['temp_c'],
    data['forecast']['simpleforecast']['forecastday'][0]['high']['fahrenheit'],
    data['forecast']['simpleforecast']['forecastday'][0]['high']['celsius'],
    data['forecast']['simpleforecast']['forecastday'][0]['low']['fahrenheit'],
    data['forecast']['simpleforecast']['forecastday'][0]['low']['celsius'],
    data['current_observation']['relative_humidity'],
    data['forecast']['simpleforecast']['forecastday'][0]['avewind']['mph'],
    data['forecast']['simpleforecast']['forecastday'][0]['avewind']['kph'],
)
