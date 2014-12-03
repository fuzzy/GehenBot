#!/usr/bin/env python

import os, sys, urllib2
import xml.etree.ElementTree as ET


if sys.argv[1].split()[0] == '--usage':
    print 'Usage: !math <math expression eg: (64**2)/1024>'
    sys.exit(0)

nick, mask, args = sys.argv[1], sys.argv[2], sys.argv[3]

data = ET.fromstring(urllib2.urlopen(
        'http://api.wolframalpha.com/v2/query?&appid=T9GW74-E93KK2YW2A&input=%s' % urllib2.quote(args)
).read())

for child in data:
    if child.attrib['title'] == 'Result':
        print child[0][0].text
        sys.exit(0)

