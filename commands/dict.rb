#!/usr/bin/env ruby

nick = ARGV[0]
mask = ARGV[1]

data = `dict -h all.dict.org -d wn #{ARGV[2]}`.split
start = false
retv = Array.new

data.each {|v|
  if not start
    if v =~ /1:/
      retv.push v
      start = true
    end
  else
    if v !~ /2:/
      retv.push v
    end
  end
}

puts retv.join(' ')
