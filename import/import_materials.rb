# -*- coding: utf-8 -*-

# This script generates a JSON material database. Each line is a
# record, consisting of a material name, its stats, colors, uses and
# ID.

require "json"

mats = {}

files = {
  "Ammo bullet"   => "data/ammo_bullet.html",
  "Ammo jacket"   => "data/ammo_jacket.html",
  "Armor clip"    => "data/armor_clip.html",
  "Armor shell"   => "data/armor_shell.html",
  "Barrel"        => "data/barrel.html",
  "Blade"         => "data/blade.html",
  "Clothes"       => "data/clothes.html",
  "Counterweight" => "data/counterweight.html",
  "Explosive"     => "data/explosive.html",
  "Firing pin"    => "data/firing_pin.html",
  "Grip"          => "data/grip.html",
  "Hammer"        => "data/hammer.html",
  "Jewel"         => "data/jewel.html",
  "Jewel setting" => "data/jewel_setting.html",
  "Lining"        => "data/lining.html",
  "Magic focus"   => "data/magic_focus.html",
  "Point"         => "data/point.html",
  "Shaft"         => "data/shaft.html",
  "Stuffing"      => "data/stuffing.html",
  "Trigger"       => "data/trigger.html",
}

files.each do |use, file|
  open(file) do |f|
    html = f.read
    html.scan(/<tr class=".+?">(.+?)<\/tr>/m) do |(block)|
      id, name, grade, ecosystem, color, *stats = block.scan(/<td>(.+?)<\/td>/).flatten
      color = color[/color\/(.+?)\.gif/, 1]

      mats[name] ||= {:stats => {}, :colors => {}, :uses => []}
      mats[name][:uses] << use unless mats[name][:uses].include?(use)
      mats[name][:stats][use] ||= {}
      mats[name][:colors] ||= {}
      mats[name][:colors][grade] ||= {}
      mats[name][:stats][use][grade] = stats.map(&:to_i)
      mats[name][:colors][grade][ecosystem] = color
    end
  end
end

json_mats = []
mats.each do |mat, data|
  json_mats << {:name => mat}.merge(data)
end

# We use the sheetid->name mapping to later determine material IDs, by
# reducing all names to their common name.
#
# For example
# - "Bundle of Basic Abhaya Wood"
# - "Bundle of Choice Desert Abhaya Wood"
# - "Bundle of Excellent Desert Abhaya Wood"
# will be reduced to "Bundle of Abhaya Wood".
# This will then later be used to extract "Abhaya Wood" from.

ids = Hash.new {|h,k| h[k]=[]}

open("data/sheetids.html") do |f|
  code = f.read

  code.scan(/<tr><td.+?>(.+?)<\/td><td.+?>(.+?)<\/td>/) do |(id, name)|
    if id =~ /^m(\d{4})(\w{5})(\d{2})/
      ids[$1] << name
    end
  end
end

ids.each_key do |index|
  names = ids[index]

  # Fossil and Yetin Bone share one ID and screw up the common part
  names_parts = names.map{|name| name.split(" ")} - [["Fossil"]]
  common = names_parts.inject{|part, memo| memo & part}

  ids[index] = {:type => :material, :name => common.join(" ")}
end

# For each material, find its ID
json_mats.each do |material|
  id, _ = ids.find {|id, data|
    mname = material[:name].tr("ï", "i").sub("Trunck", "Trunk")
    data[:name].include?(mname)
  }

  material[:id] = id
end

# Add florist materials that are missing in our import source.
json_mats << {:name => "Cratcha Pistil",  :stats  => {}, :colors => {}, :uses => [], :id => "0163"}
json_mats << {:name => "Psykopla Pistil", :stats => {},  :colors => {}, :uses => [], :id => "0169"}
json_mats << {:name => "Slaveni Pistil",  :stats  => {}, :colors => {}, :uses => [], :id => "0171"}

json_mats.each do |mat|
  puts mat.to_json
end
