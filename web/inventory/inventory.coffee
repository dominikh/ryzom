parseSheetID = (id) ->
    # For now, we only parse materials
    if id[0] == 'm'
        return parseSheetIDMaterial(id)
    if id[0..1] == 'tp'
        return parseSheetIDTeleport(id)
    return ["other"]


parseSheetIDMaterial = (id) ->
    type = id[5]
    ecosystemId = id[8]
    gradeId = id[9]

    ecosystems = []
    ecosystems["d"] = "desert"
    ecosystems["f"] = "forest"
    ecosystems["j"] = "jungle"
    ecosystems["l"] = "lakes"
    ecosystems["p"] = "primeroots"
    ecosystems["c"] = "generic"

    grades = []
    if type == 'd'
        grades["b"] = "basic"
        grades["c"] = "fine"
        grades["d"] = "choice"
        grades["e"] = "excellent"
        grades["f"] = "supreme"
    else
        grades["a"] = "basic"
        grades["b"] = "fine"
        grades["c"] = "choice"
        grades["d"] = "excellent"
        grades["e"] = "supreme"

    return ["material", ecosystems[ecosystemId], grades[gradeId]]

parseSheetIDTeleport = (id) ->
    lands = {
        "avalae": "forest",
        "avendale": "lakes",
        "bountybeaches": "lakes",
        "crystabell": "lakes",
        "davae": "forest",
        "dewdrops": "lakes",
        "dyron": "desert",
        "enchantedisle": "lakes",
        "fairhaven": "lakes",
        "fleetinggarden": "forest",
        "forbiddendepths": "pr",
        "frahartowers": "desert",
        "gateofobscurity": "pr",
        "groveofconfusion": "forest",
        "groveofumbra": "jungle",
        "havenofpurity": "jungle",
        "hereticshovel": "forest",
        "hiddensource": "forest",
        "knollofdissent": "forest",
        "knotofdementia": "jungle",
        "lagoonsofloria": "lakes",
        "maidengrove": "jungle",
        "natae": "forest",
        "nexusterre": "pr",
        "oflovaksoasis": "desert",
        "outlawcanyon": "desert",
        "pyr": "desert",
        "restingwater": "lakes",
        "sawdustmines": "desert",
        "theabyssofichormatis": "pr",
        "theelusiveforest": "pr",
        "thelandofcontinuity": "pr",
        "thesunkencity": "pr",
        "thetrenchoftrialszorai": "pr",
        "theunderspringfyros": "pr",
        "thewindygate": "pr",
        "thefount": "lakes",
        "thescorchedcorridor": "desert",
        "thesos": "desert",
        "thevoid": "jungle",
        "upperbog": "forest",
        "windermeer": "lakes",
        "windsofmuse": "lakes",
        "yrkanis": "forest",
        "zora": "jungle",
        "jenlai": "jungle",
        "hoicho": "jungle",
    }

    split = id[3..].split("_")
    faction = split[0]
    destination = split[1..].join("")
    # [faction, destination] = id[3..].split("_")
    destination = destination.split(".")[0]

    return ["teleport", faction, lands[destination], destination]


parseFilterString = (text) ->
    words = []
    quality = null
    ecosystem = null
    grade = null
    use = null

    text = text.toLowerCase()
    text = text.replace("ammo bullet", "ammo_bullet").replace("ammo jacket", "ammo_jacket")
    text = text.replace("armor clip", "armor_clip").replace("armor shell", "armor_shell")
    text = text.replace("firing pin", "firing_pin").replace("jewel setting", "jewel_setting")
    text = text.replace("magic focus", "magic_focus")

    for index, word of text.toLowerCase().split(" ")
        match = /^q(\d+)$/i.exec(word)
        if match
            quality = match[1]
        else if word in ["forest", "desert", "jungle", "lakes", "generic", "primeroots"]
            ecosystem = word
        else if word in ["basic", "fine", "choice", "excellent", "supreme"]
            grade = word
        else if word in ["ammo_bullet", "ammo_jacket", "armor_clip", "armor_shell",
            "barrel", "blade", "clothes", "counterweight", "explosive", "firing_pin",
            "grip", "hammer", "jewel", "jewel_setting", "lining", "magic_focus", "point",
            "shaft", "stuffing", "trigger"]
                use = word
        else
            words.push(word)

    return [words.join(" "), quality, ecosystem, grade, use]

parseFilterStrings = (text) ->
    text.split("|").map (t) -> parseFilterString(t)

filterInventory = (inventorySelection) ->
    text = $('#filter_input').val()
    # TODO support other filtering than material based filtering

    # skip = []
    for obj in $('.inventory ul li')
        origObj = obj
        obj = $(obj)

        for filter in $(parseFilterStrings(text))
            [name, quality, ecosystem, grade, use] = filter
            # return if origObj in skip

            [type, data...] = parseSheetID(obj.attr("data-sheetid"))
            if type == "material"
                mEcosystem = data[0]
                mGrade = data[1]
            else
                mEcosystem = ""
                mGrade = ""

            if !inventorySelection[obj.attr("data-inventory")] ||
                (quality != null && obj.attr("data-quality") != quality) ||
                (ecosystem != null && mEcosystem != ecosystem) ||
                (grade != null && mGrade != grade) ||
                (use != null && use not in obj.attr("data-uses").split(",")) ||
                (name != "" && obj.attr("data-name").toLowerCase().indexOf(name) == -1)
                    obj.hide()
            else
            # skip.push(origObj)
                obj.show()
                break

$ ->
    inventories = null
    inventorySelection = {
        "bag": true,
        "room": true,
        "pet1": true,
        "pet2": true,
        "pet3": true,
        "pet4": true,
    }
    $('#filter_input').bind 'keyup', (event) ->
        if $("#incremental_filter")[0].checked || event.which == 13
            filterInventory(inventorySelection)
    $('#example_filter_quality').bind 'click', (event) ->
        $('#filter_input').val("q160")
        filterInventory(inventorySelection)
    $('#example_filter_name').bind 'click', (event) ->
        $('#filter_input').val("dzao fiber")
        filterInventory(inventorySelection)
    $('#example_filter_grade').bind 'click', (event) ->
        $('#filter_input').val("excellent")
        filterInventory(inventorySelection)
    $('#example_filter_ecosystem').bind 'click', (event) ->
        $('#filter_input').val("forest")
        filterInventory(inventorySelection)
    $('#example_filter_craft').bind 'click', (event) ->
        $('#filter_input').val("point")
        filterInventory(inventorySelection)
    $('#example_filter_combination').bind 'click', (event) ->
        $('#filter_input').val("q160 excellent forest")
        filterInventory(inventorySelection)
    $('#example_filter_priority').bind 'click', (event) ->
        $('#filter_input').val("excellent capryni skin|excellent gubani skin|excellent igara leather|excellent kirosta mandible|excellent madakam pelvis|excellent madakam skin|excellent mektoub skin|excellent najab leather|excellent ploderos nail|excellent ploderos skin")
        filterInventory(inventorySelection)

    $("#load_inventory").bind "click", (event) ->
        $.getJSON "/inventory/?key=" + $('#api_key_input').val(), (data) ->
            inventories = data

            creationDate = new Date(inventories["bag"].CreationTime * 1000)
            expireDate = new Date(inventories["bag"].ExpireTime * 1000)

            $("#inventories_information .last_update").html(creationDate.toLocaleTimeString())
            $("#inventories_information .next_update").html(expireDate.toLocaleTimeString())

            displayInventories(inventories)
            filterInventory(inventorySelection)

    $(".inventory_selector li").bind "click", (event) ->
        toggleSelection(this, inventorySelection)
        filterInventory(inventorySelection)

toggleSelection = (el, selections) ->
    selection = $(el).attr("data-inventory")
    $(el).toggleClass("active")
    selections[$(el).attr("data-inventory")] = !selections[selection]

setItemWidth = ->
    $(".inventory ul").each (index, obj) ->
        widest = 0
        thisWidth = 0

        $(obj).children().each (index2, obj2) ->
            thisWidth = parseInt($(obj2).css('width'))

            if (thisWidth > widest)
                widest = thisWidth

        widest += 20 # padding + margin
        widest += 'px'

        $(obj).children().css('width', widest);


displayInventories = (inventories) ->
    # TODO remove existing elements
    list = $(".inventory ul")
    list.empty()
    displayInventory(inventories.bag)
    displayInventory(inventories.room)
    displayInventory(inventories.pet1)
    displayInventory(inventories.pet2)
    displayInventory(inventories.pet3)
    displayInventory(inventories.pet4)
    setItemWidth()

inventoryNameToDisplayName = (bagName) ->
    {
        "bag": "Bag",
        "room": "Apartment",
        "pet1": "Animal 1",
        "pet2": "Animal 2",
        "pet3": "Animal 3",
        "pet4": "Animal 4",
    }[bagName]


displayInventory = (inventory) ->
    return unless inventory.Items?
    list = $(".inventory ul")
    inventory.Items.sort (a, b) ->
        # TODO Proper sorting:
        # - If material: use, grade, quality, ecosystem (TODO check if this is like in game or not)
        # - If teleport: Land, name (TODO this will need a map of tp name to land)
        # - Otherwise sort by SheetID, Quality
        aID = a.SheetID
        bID = b.SheetID

        aParsed = parseSheetID(aID)
        bParsed = parseSheetID(bID)

        # Compare type
        s0 = aParsed[0].localeCompare(bParsed[0])
        if s0 != 0
            return s0

        # Default sorting for non teleports
        if aParsed[0] != "teleport"
            if aParsed[0] == "material"
                # Sort by uses
                s4 = (a.Uses[0] || "").localeCompare(b.Uses[0] || "")
                if s4 != 0
                    return s4

                s5 = a.Name.localeCompare(b.Name)
                if s5 != 0
                    return s5

            # Sort by SheetID
            if aParsed[0] != "material"
                s3 = aID.localeCompare(bID)
                if s3 != 0
                    return s3

            # sort by quality
            if a.Quality == b.Quality
                return 0
            else if a.Quality > b.Quality
                return 1
            else
                return -1

        # tp faction
        s1 = aParsed[1].localeCompare(bParsed[1])
        if s1 != 0
            return s1

        # tp land
        s2 = aParsed[2].localeCompare(bParsed[2])
        if s2 != 0
            return s2

        # tp destination
        return aParsed[3].localeCompare(bParsed[3])

    $(inventory.Items).each (index, item) ->
        if item.SheetID[0] == '#'
            # No idea what those weird items are
            return

        uses = "<span class='uses'>&nbsp;</span>"
        if item.Uses.length > 0
            uses = "<span class='uses'>Can craft: " + item.Uses.join(", ") + "</span>"

        list.append($("<li data-name='" + item.Name + "' data-quality='" + item.Quality + "' data-sheetid='" + item.SheetID + "'
            data-inventory='" + inventory.Name + "' data-uses='" + item.Uses.join(",").replace(/\s/g, "_").toLowerCase() + "'>
            <img src='/item_icon/?id=" + item.SheetID + "&amp;qual=" + item.Quality + "&amp;quan=" + item.Quantity + "' />
            <span class='name'>" + item.Name + "</span>
            <span class='location'>Stored in: " + inventoryNameToDisplayName(inventory.Name) + "</span>" + uses))
