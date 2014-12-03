#!/usr/bin/env python

import os, re, sys, json, urllib, urllib2
import xml.etree.ElementTree as ET

if sys.argv[1] == '--usage':
    print "Usage: !eztv <search term eg: Adventure Time The dentist>"
    sys.exit(0)

nick, mask, args = sys.argv[1], sys.argv[2], sys.argv[3]

uris = {
    'eztv': 'http://kickass.to/json.php',
    'goo.gl': '',
}

# First lets get our eztv data
values   = urllib.urlencode({'field': 'seeders', 'order': 'desc', 'q': args})
headers  = {'User-Agent': 'Mozilla/4.0 (compatible; MSIE 5.5; Windows NT)'}
request  = urllib2.Request('%s?%s' % (uris['eztv'], values), None, headers)
response = urllib2.urlopen(request)
data     = json.loads(response.read())

print 'magnet:?xt=urn:btih:%s&dn=%s&tr=udp://open.demonii.com:1337/announce' % (
    data['list'][0]['hash'],
    re.sub('(\[|\])', '', re.sub(' ', '+', data['list'][0]['title'].lower())),
    )
