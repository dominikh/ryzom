# -*- coding: utf-8 -*-

# This script generates a JSON file, mapping Sheet IDs to item names.

require "json"
require "pp"

open("data/sheetids.html") do |f|
  code = f.read

  ids = {}
  code.scan(/<tr><td.+?>(.+?)<\/td><td.+?>(.+?)<\/td>/) do |(id, name)|
    ids[id] = name
  end

  puts ids.to_json
end
