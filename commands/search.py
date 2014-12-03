#!/usr/bin/env python

import sys, json, urllib2, random

if sys.argv[1] == '--usage':
    print 'Usage: !search <search terms>'
    sys.exit(0)

nick, host, args = sys.argv[1], sys.argv[2], sys.argv[3]
data = json.loads(urllib2.urlopen(
        'http://api.duckduckgo.com/?q=%s&format=json' % urllib2.quote(args)
    ).read())

if 'AbstractURL' in data.keys():
    if len(data['AbstractURL']) != 0:
        print '%s Related: %s' % (
            data['AbstractURL'],
            data['RelatedTopics'][random.randint(0, len(data['RelatedTopics'])-1)]['FirstURL']
        )
    else:
        print 'No results found.'
else:
    print 'No results found.'
