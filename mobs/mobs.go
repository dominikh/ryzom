package mobs

import (
	"errors"
)

type Mob int
type Name int
type Class int

func (mob Mob) String() string {
	return ReverseMobMap[mob]
}

func (name Name) String() string {
	return ReverseNameMap[name]
}

type Creature struct {
	Mob   Mob
	Level Level
}

type Level struct {
	Level int
	Name  Name
}

func (mob Mob) LevelFromName(name Name) (Level, error) {
	for _, level := range Levels[mob] {
		if level.Name == name {
			return level, nil
		}
	}

	// TODO return error
	return Level{}, errors.New("Unknown name")
}

func (mob Mob) LevelFromInt(lvl int) Level {
	for _, level := range Levels[mob] {
		if level.Level == lvl {
			return level
		}
	}

	// TODO return error
	return Level{}
}

const (
	Arana Mob = iota
	Arma
	Bawaab
	Bodoc
	Bolobi
	Capryni
	Clopper
	Cratcha
	Cray
	Cute
	Cuttler
	Frahar
	Frippo
	Gibbai
	Gingo
	Gnoof
	Goari
	Gubani
	Horncher
	Igara
	Izam
	Javing
	Jubla
	Jugula
	Kiban
	Kidinak
	Kincher
	Kinrey
	Kipee
	Kipesta
	Kipucka
	Kipucker
	Kirosta
	Kizarak
	Kizoar
	Lumper
	Madakam
	Mektoub
	Messab
	Najab
	Ocyx
	Ploderos
	Psykopla
	Ragus
	Raspal
	Rendor
	Shalah
	Shooki
	Slaveni
	Stinga
	Timari
	Torbak
	Tyrancha
	Varinx
	Vorax
	Wombai
	Yber
	Yelk
	Yetin
	Yubo
	Zerx
)

const (
	Unknown Name = iota
	// --
	Awesome
	Baying
	Belligerent
	Bloated
	Blooming
	Budding
	Carrion
	Crabby
	Creeping
	Dangerous
	Dehydrated
	Devastating
	Docile
	Dominant
	Dreaded
	Drowsy
	EliteOverlord
	Famished
	Fearsome
	Feral
	Ferocious
	Fierce
	Fledgling
	Frightening
	Furious
	Gluttonous
	Gorged
	Great
	GreatOverlord
	Growling
	Gruesome
	Gruff
	Grunting
	Hard
	Horrific
	Horrifying
	Huge
	Hungry
	Hunting
	Incensed
	Killer
	Lacerating
	Leering
	Lesser
	Lethal
	Lurking
	Malicious
	Malignant
	Marauding
	Master
	Mean
	Menacing
	Mighty
	Minor
	Mocking
	Moderate
	Morbid
	Nauseous
	Nettled
	NewBorn
	Noxious
	Obstinate
	Overlord
	Parched
	PowerOverlord
	Preying
	Prickly
	Prime
	Prodigeous
	Prowling
	Puny
	Raging
	Rapacious
	Rattled
	Roaming
	Roaring
	Robust
	Rooting
	Savage
	Scampering
	Scary
	Scavenging
	Scowling
	Scrounging
	Seasoned
	Sluggish
	Sprightly
	Stalking
	Strong
	Suckling
	Terrifying
	Timorous
	Unpredictable
	Veteran
	Vicious
	Vigilant
	Vigorous
	Vile
	Violent
	Voracious
	Vulgar
	Wary
	Weanling
	Weeny
	Placid
	Peaceful
)

const (
	Avian Class = iota
	Carnivore
	Flora
	Herbivore
	Javan
	Kitin
	KitinInvader
	PrimitiveTribe
)

var MobMap = map[string]Mob{
	"arana":    Arana,
	"arma":     Arma,
	"bawaab":   Bawaab,
	"bodoc":    Bodoc,
	"bolobi":   Bolobi,
	"capryni":  Capryni,
	"clopper":  Clopper,
	"cratcha":  Cratcha,
	"cray":     Cray,
	"cute":     Cute,
	"cuttler":  Cuttler,
	"frahar":   Frahar,
	"frippo":   Frippo,
	"gibbai":   Gibbai,
	"gingo":    Gingo,
	"gnoof":    Gnoof,
	"goari":    Goari,
	"gubani":   Gubani,
	"horncher": Horncher,
	"igara":    Igara,
	"izam":     Izam,
	"javing":   Javing,
	"jubla":    Jubla,
	"jugula":   Jugula,
	"kiban":    Kiban,
	"kidinak":  Kidinak,
	"kincher":  Kincher,
	"kinrey":   Kinrey,
	"kipee":    Kipee,
	"kipesta":  Kipesta,
	"kipucka":  Kipucka,
	"kipucker": Kipucker,
	"kirosta":  Kirosta,
	"kizarak":  Kizarak,
	"kizoar":   Kizoar,
	"lumper":   Lumper,
	"madakam":  Madakam,
	"mektoub":  Mektoub,
	"messab":   Messab,
	"najab":    Najab,
	"ocyx":     Ocyx,
	"ploderos": Ploderos,
	"psykopla": Psykopla,
	"ragus":    Ragus,
	"raspal":   Raspal,
	"rendor":   Rendor,
	"shalah":   Shalah,
	"shooki":   Shooki,
	"slaveni":  Slaveni,
	"stinga":   Stinga,
	"timari":   Timari,
	"torbak":   Torbak,
	"tyrancha": Tyrancha,
	"varinx":   Varinx,
	"vorax":    Vorax,
	"wombai":   Wombai,
	"yber":     Yber,
	"yelk":     Yelk,
	"yetin":    Yetin,
	"yubo":     Yubo,
	"zerx":     Zerx,
}

var NameMap = map[string]Name{
	"baying":        Baying,
	"blooming":      Blooming,
	"budding":       Budding,
	"dangerous":     Dangerous,
	"dehydrated":    Dehydrated,
	"devastating":   Devastating,
	"docile":        Docile,
	"rowsy":         Drowsy,
	"famished":      Famished,
	"feral":         Feral,
	"ferocious":     Ferocious,
	"fierce":        Fierce,
	"furious":       Furious,
	"gluttonous":    Gluttonous,
	"growling":      Growling,
	"gruff":         Gruff,
	"grunting":      Grunting,
	"horrifying":    Horrifying,
	"malicious":     Malicious,
	"menacing":      Menacing,
	"minor":         Minor,
	"morbid":        Morbid,
	"nauseous":      Nauseous,
	"nettled":       Nettled,
	"noxious":       Noxious,
	"obstinate":     Obstinate,
	"parched":       Parched,
	"prickly":       Prickly,
	"prodigeous":    Prodigeous,
	"puny":          Puny,
	"rapacious":     Rapacious,
	"rattled":       Rattled,
	"roaming":       Roaming,
	"robust":        Robust,
	"rooting":       Rooting,
	"scampering":    Scampering,
	"scary":         Scary,
	"scowling":      Scowling,
	"scrounging":    Scrounging,
	"sluggish":      Sluggish,
	"sprightly":     Sprightly,
	"stalking":      Stalking,
	"suckling":      Suckling,
	"timorous":      Timorous,
	"unpredictable": Unpredictable,
	"vicious":       Vicious,
	"vigorous":      Vigorous,
	"vile":          Vile,
	"violent":       Violent,
	"voracious":     Voracious,
	"wary":          Wary,
	"weanling":      Weanling,
	"weeny":         Weeny,
	// new below this line
	"dominant":    Dominant,
	"seasoned":    Seasoned,
	"veteran":     Veteran,
	"lurking":     Lurking,
	"creeping":    Creeping,
	"crabby":      Crabby,
	"belligerent": Belligerent,
	"frightening": Frightening,
	"preying":     Preying,
	"raging":      Raging,

	"fearsome":   Fearsome,
	"huge":       Huge,
	"terrifying": Terrifying,
	"awesome":    Awesome,
	"horrific":   Horrific,
	"great":      Great,
	"savage":     Savage,
	"prowling":   Prowling,

	"roaring":    Roaring,
	"marauding":  Marauding,
	"hunting":    Hunting,
	"prime":      Prime,
	"vigilant":   Vigilant,
	"hungry":     Hungry,
	"scavenging": Scavenging,
	"leering":    Leering,
	"malignant":  Malignant,
	"mocking":    Mocking,

	"carrion":   Carrion,
	"fledgling": Fledgling,
	"lesser":    Lesser,
	"gorged":    Gorged,
	"vulgar":    Vulgar,
	"bloated":   Bloated,
	"newborn":   NewBorn,

	"incensed":   Incensed,
	"mighty":     Mighty,
	"gruesome":   Gruesome,
	"lacerating": Lacerating,
	"killer":     Killer,
	"dreaded":    Dreaded,
	"master":     Master,
	"moderate":   Moderate,
	"hard":       Hard,
	"strong":     Strong,
	"mean":       Mean,
	"lethal":     Lethal,
	"placid":     Placid,
	"peaceful":   Peaceful,
}

var ReverseMobMap = map[Mob]string{
	Arana:    "Arana",
	Arma:     "Arma",
	Bawaab:   "Bawaab",
	Bodoc:    "Bodoc",
	Bolobi:   "Bolobi",
	Capryni:  "Capryni",
	Clopper:  "Clopper",
	Cratcha:  "Cratcha",
	Cray:     "Cray",
	Cute:     "Cute",
	Cuttler:  "Cuttler",
	Frahar:   "Frahar",
	Frippo:   "Frippo",
	Gibbai:   "Gibbai",
	Gingo:    "Gingo",
	Gnoof:    "Gnoof",
	Goari:    "Goari",
	Gubani:   "Gubani",
	Horncher: "Horncher",
	Igara:    "Igara",
	Izam:     "Izam",
	Javing:   "Javing",
	Jubla:    "Jubla",
	Jugula:   "Jugula",
	Kiban:    "Kiban",
	Kidinak:  "Kidinak",
	Kincher:  "Kincher",
	Kinrey:   "Kinrey",
	Kipee:    "Kipee",
	Kipesta:  "Kipesta",
	Kipucka:  "Kipucka",
	Kipucker: "Kipucker",
	Kirosta:  "Kirosta",
	Kizarak:  "Kizarak",
	Kizoar:   "Kizoar",
	Lumper:   "Lumper",
	Madakam:  "Madakam",
	Mektoub:  "Mektoub",
	Messab:   "Messab",
	Najab:    "Najab",
	Ocyx:     "Ocyx",
	Ploderos: "Ploderos",
	Psykopla: "Psykopla",
	Ragus:    "Ragus",
	Raspal:   "Raspal",
	Rendor:   "Rendor",
	Shalah:   "Shalah",
	Shooki:   "Shooki",
	Slaveni:  "Slaveni",
	Stinga:   "Stinga",
	Timari:   "Timari",
	Torbak:   "Torbak",
	Tyrancha: "Tyrancha",
	Varinx:   "Varinx",
	Vorax:    "Vorax",
	Wombai:   "Wombai",
	Yber:     "Yber",
	Yelk:     "Yelk",
	Yetin:    "Yetin",
	Yubo:     "Yubo",
	Zerx:     "Zerx",
}

var ReverseNameMap = map[Name]string{
	Baying:        "Baying",
	Blooming:      "Blooming",
	Budding:       "Budding",
	Dangerous:     "Dangerous",
	Dehydrated:    "Dehydrated",
	Devastating:   "Devastating",
	Docile:        "Docile",
	Drowsy:        "Drowsy",
	Famished:      "Famished",
	Feral:         "Feral",
	Ferocious:     "Ferocious",
	Fierce:        "Fierce",
	Furious:       "Furious",
	Gluttonous:    "Gluttonous",
	Growling:      "Growling",
	Gruff:         "Gruff",
	Grunting:      "Grunting",
	Horrifying:    "Horrifying",
	Malicious:     "Malicious",
	Menacing:      "Menacing",
	Minor:         "Minor",
	Morbid:        "Morbid",
	Nauseous:      "Nauseous",
	Nettled:       "Nettled",
	Noxious:       "Noxious",
	Obstinate:     "Obstinate",
	Parched:       "Parched",
	Prickly:       "Prickly",
	Prodigeous:    "Prodigeous",
	Puny:          "Puny",
	Rapacious:     "Rapacious",
	Rattled:       "Rattled",
	Roaming:       "Roaming",
	Robust:        "Robust",
	Rooting:       "Rooting",
	Scampering:    "Scampering",
	Scary:         "Scary",
	Scowling:      "Scowling",
	Scrounging:    "Scrounging",
	Sluggish:      "Sluggish",
	Sprightly:     "Sprightly",
	Stalking:      "Stalking",
	Suckling:      "Suckling",
	Timorous:      "Timorous",
	Unpredictable: "Unpredictable",
	Vicious:       "Vicious",
	Vigorous:      "Vigorous",
	Vile:          "Vile",
	Violent:       "Violent",
	Voracious:     "Voracious",
	Wary:          "Wary",
	Weanling:      "Weanling",
	Weeny:         "Weeny",
	Dominant:      "Dominant",
	Seasoned:      "Seasoned",
	Veteran:       "Veteran",
	Lurking:       "Lurking",
	Creeping:      "Creeping",
	Crabby:        "Crabby",
	Belligerent:   "Belligerent",
	Frightening:   "Frightening",
	Preying:       "Preying",
	Raging:        "Raging",
	Fearsome:      "Fearsome",
	Huge:          "Huge",
	Terrifying:    "Terrifying",
	Awesome:       "Awesome",
	Horrific:      "Horrific",
	Great:         "Great",
	Savage:        "Savage",
	Prowling:      "Prowling",
	Roaring:       "Roaring",
	Marauding:     "Marauding",
	Hunting:       "Hunting",
	Prime:         "Prime",
	Vigilant:      "Vigilant",
	Hungry:        "Hungry",
	Scavenging:    "Scavenging",
	Leering:       "Leering",
	Malignant:     "Malignant",
	Mocking:       "Mocking",
	Carrion:       "Carrion",
	Fledgling:     "Fledgling",
	Lesser:        "Lesser",
	Gorged:        "Gorged",
	Vulgar:        "Vulgar",
	Bloated:       "Bloated",
	NewBorn:       "New born",
	Incensed:      "Incensed",
	Mighty:        "Mighty",
	Gruesome:      "Gruesome",
	Lacerating:    "Lacerating",
	Killer:        "Killer",
	Dreaded:       "Dreaded",
	Master:        "Master",
	Moderate:      "Moderate",
	Hard:          "Hard",
	Strong:        "Strong",
	Mean:          "Mean",
	Lethal:        "Lethal",
	Placid:        "Placid",
	Peaceful:      "Peaceful",
	Unknown:       "Unknown name",
}

var ClassMap = map[string]Class{
	"avian":           Avian,
	"carnivore":       Carnivore,
	"flora":           Flora,
	"herbivore":       Herbivore,
	"javan":           Javan,
	"kitin":           Kitin,
	"kitin invader":   KitinInvader,
	"primitive tribe": PrimitiveTribe,
}

var Coefficients = map[Mob]float64{
	Arana:    1,
	Arma:     1,
	Bawaab:   1.5,
	Bodoc:    1.5,
	Bolobi:   7.5,
	Capryni:  1,
	Clopper:  1.5,
	Cratcha:  1,
	Cray:     1,
	Cute:     1,
	Cuttler:  1,
	Frahar:   1,
	Frippo:   1,
	Gibbai:   1,
	Gingo:    1.5,
	Gnoof:    1,
	Goari:    1,
	Gubani:   1,
	Horncher: 1,
	Igara:    1,
	Izam:     1,
	Javing:   1,
	Jubla:    1,
	Jugula:   1,
	Kiban:    1,
	Kidinak:  1.5,
	Kincher:  2.5,
	Kinrey:   2.5,
	Kipee:    2,
	Kipesta:  1,
	Kipucka:  1,
	Kipucker: 1,
	Kirosta:  1,
	Kizarak:  1,
	Kizoar:   1,
	Lumper:   1.5,
	Madakam:  5,
	Mektoub:  1.5,
	Messab:   1,
	Najab:    1,
	Ocyx:     1.5,
	Ploderos: 7.5,
	Psykopla: 1,
	Ragus:    1,
	Raspal:   1,
	Rendor:   1,
	Shalah:   5,
	Shooki:   1,
	Slaveni:  1,
	Stinga:   1,
	Timari:   1,
	Torbak:   1.5,
	Tyrancha: 1.5,
	Varinx:   1,
	Vorax:    1,
	Wombai:   1,
	Yber:     1,
	Yelk:     1.5,
	Yetin:    1,
	Yubo:     1,
	Zerx:     1.5,
	// 8,
	// 25,
}

var Levels = map[Mob][]Level{
	Arana: {{138, Nettled},
		{146, Vicious},
		{155, Feral},
		{163, Violent},
		{179, Furious},
		{187, Fierce},
		{196, Ferocious},
		{204, Voracious}},

	Arma: {{51, Robust},
		{60, Wary},
		{72, Gluttonous},
		{81, Grunting},
		{96, Gruff},
		{105, Obstinate},
		{114, Rooting},
		{123, Scrounging},
		{133, Nettled},
		{142, Vicious},
		{150, Feral},
		{159, Violent},
		{175, Furious},
		{183, Fierce},
		{192, Ferocious},
		{200, Voracious}},

	Bawaab: {{67, Robust},
		{78, Wary},
		{88, Gluttonous},
		{99, Grunting},
		{117, Gruff},
		{127, Obstinate},
		{137, Rooting},
		{147, Scrounging}},

	Bodoc: {{8, Suckling},
		{14, Weanling},
		{18, Docile},
		{10, Timorous},
		{21, Roaming},
		{30, Scampering},
		{41, Sprightly},
		{58, Robust},
		{69, Wary},
		{79, Gluttonous},
		{89, Grunting},
		{102, Gruff},
		{114, Obstinate},
		{124, Rooting},
		{133, Scrounging},
		{143, Nettled},
		{152, Vicious},
		{161, Feral},
		{170, Violent}},

	Bolobi: {{151, Nettled},
		{161, Vicious},
		{180, Feral},
		{190, Violent},
		{235, Ferocious},
		{250, Voracious}},

	Capryni: {{8, Suckling},
		{15, Weanling},
		{23, Nauseous},
		{14, Timorous},
		{24, Roaming},
		{33, Scampering},
		{43, Sprightly},
		{58, Robust},
		{68, Wary},
		{77, Gluttonous},
		{87, Grunting},
		{101, Gruff},
		{110, Obstinate},
		{118, Rooting},
		{127, Scrounging}},

	Clopper: {{11, Weeny},
		{20, Puny},
		{23, Vigorous},
		{35, Growling},
		{46, Scowling},
		{59, Baying},
		{72, Scary},
		{86, Malicious},
		{97, Dangerous},
		{109, Menacing}},

	Cratcha: {{107, Parched},
		{117, Stalking},
		{127, Rooting},
		{137, Noxious},
		{149, Vile},
		{159, Unpredictable},
		{168, Prickly},
		{178, Rattled},
		{196, Horrifying},
		{205, Rapacious},
		{215, Prodigeous},
		{224, Devastating}},

	Cray: {{101, Dominant},
		{110, Gruff},
		{119, Obstinate},
		{128, Rooting},
		{162, Scrounging}},

	Cute: {{112, Seasoned},
		{122, Veteran},
		{131, Lurking},
		{141, Creeping},
		{151, Crabby},
		{161, Belligerent},
		{179, Frightening}},

	Cuttler: {{172, Preying},
		{182, Raging},
		{192, Fearsome},
		{202, Huge},
		{221, Terrifying},
		{231, Awesome},
		{241, Horrific},
		{251, Great}},

	Frahar: {{110, Savage},
		{120, Seasoned},
		{130, Veteran},
		{139, Lurking},
		{151, Creeping},
		{161, Crabby},
		{170, Belligerent},
		{179, Frightening}},

	Frippo: {{60, Robust},
		{70, Wary},
		{80, Gluttonous},
		{91, Grunting}},

	Gibbai: {{111, Savage},
		{120, Seasoned},
		{130, Veteran},
		{140, Lurking},
		{151, Creeping},
		{161, Crabby},
		{170, Belligerent},
		{179, Frightening}},

	Gingo: {{15, Weeny},
		{24, Puny},
		{23, Vigorous},
		{35, Growling},
		{46, Scowling},
		{59, Baying},
		{72, Scary},
		{82, Malicious},
		{97, Dangerous},
		{109, Menacing},
		{123, Roaring},
		{133, Marauding},
		{143, Hunting},
		{153, Prowling},
		{166, Preying},
		{175, Raging},
		{185, Fearsome},
		{195, Huge}},

	Gnoof: {{64, Robust},
		{76, Wary},
		{86, Gluttonous},
		{97, Grunting},
		{150, Nettled},
		{159, Vicious},
		{168, Feral},
		{176, Violent}},

	Goari: {{11, Weeny},
		{20, Puny},
		{20, Vigorous},
		{32, Growling},
		{44, Scowling},
		{57, Baying},
		{70, Scary},
		{83, Malicious},
		{91, Dangerous},
		{106, Menacing},
		{150, Prowling},
		{247, Great}},

	Gubani: {{140, Nettled},
		{149, Vicious},
		{157, Feral},
		{166, Violent},
		{181, Furious},
		{189, Fierce},
		{198, Ferocious},
		{206, Voracious}},

	Horncher: {{119, Roaring},
		{129, Marauding},
		{140, Hunting},
		{150, Prowling},
		{217, Terrifying},
		{227, Awesome},
		{237, Horrific},
		{247, Great}},

	Igara: {{50, Prime},
		{65, Vigilant},
		{78, Hungry},
		{88, Scavenging},
		{99, Leering},
		{114, Vicious},
		{124, Malignant},
		{134, Mocking},
		{144, Carrion}},

	Izam: {{10, Fledgling},
		{17, Lesser},
		{15, Gorged},
		{26, Vulgar},
		{37, Bloated},
		{44, Prime},
		{59, Vigilant},
		{69, Hungry},
		{78, Scavenging},
		{89, Leering},
		{106, Vicious},
		{116, Malignant},
		{125, Mocking},
		{135, Carrion}},

	Javing: {{9, Fledgling},
		{9, Weeny},
		{15, NewBorn},
		{17, Vigorous},
		{29, Growling},
		{29, Vulgar},
		{41, Bloated},
		{51, Prime},
		{69, Scary},
		{80, Malicious},
		{91, Dangerous},
		{103, Menacing},
		{118, Vicious},
		{118, Roaring},
		{128, Malignant},
		{128, Marauding},
		{138, Hunting},
		{149, Prowling}},

	Jubla: {{109, Parched},
		{119, Stalking},
		{129, Rooting},
		{139, Noxious},
		{150, Vile},
		{159, Unpredictable},
		{169, Prickly},
		{178, Rattled},
		{195, Horrifying},
		{205, Rapacious},
		{214, Prodigeous},
		{224, Devastating}},

	Jugula: {{224, Terrifying},
		{235, Awesome},
		{245, Horrific},
		{255, Great}},

	Kiban: {{153, Nettled},
		{163, Incensed},
		{172, Mighty},
		{181, Awesome},
		{199, Gruesome},
		{208, Lacerating},
		{217, Killer},
		{227, Great}},

	Kidinak: {{172, Nettled},
		{182, Incensed},
		{193, Mighty},
		{203, Awesome},
		{232, Lacerating},
		{242, Killer},
		{253, Great}},

	Kincher: {{126, Dreaded},
		{136, Master},
		{147, Veteran},
		{157, Hungry},
		{169, Nettled},
		{179, Incensed},
		{189, Mighty},
		{199, Awesome},
		{218, Gruesome},
		{228, Lacerating},
		{238, Killer},
		{248, Great}},

	Kinrey: {{177, Nettled},
		{188, Incensed},
		{199, Mighty},
		{209, Awesome},
		{228, Gruesome},
		{239, Lacerating},
		{249, Killer}},

	Kipee: {{21, Weeny},
		{16, Vigorous},
		{28, Vulgar},
		{39, Moderate},
		{52, Hard},
		{64, Vigilant},
		{75, Strong},
		{90, Mean},
		{102, Vicious},
		{117, Dreaded},
		{127, Master},
		{137, Veteran},
		{148, Hungry},
		{160, Nettled},
		{170, Incensed},
		{180, Mighty},
		{190, Awesome},
		{208, Gruesome},
		{218, Lacerating},
		{228, Killer},
		{238, Great}},

	Kipesta: {{120, Dreaded},
		{130, Master},
		{140, Veteran},
		{150, Hungry},
		{163, Nettled},
		{172, Incensed},
		{182, Mighty},
		{192, Awesome},
		{210, Gruesome},
		{220, Lacerating},
		{230, Killer},
		{239, Great}},

	Kipucka: {{166, Nettled},
		{176, Incensed},
		{186, Mighty},
		{196, Awesome},
		{215, Gruesome},
		{225, Lacerating},
		{235, Killer},
		{245, Great}},

	Kipucker: {{215, Overlord},
		{225, PowerOverlord},
		{235, GreatOverlord},
		{245, EliteOverlord}},

	Kirosta: {{172, Nettled},
		{183, Incensed},
		{193, Mighty},
		{203, Awesome},
		{223, Gruesome},
		{233, Lacerating},
		{243, Killer},
		{253, Great}},

	Kizarak: {{172, Nettled},
		{203, Awesome},
		{253, Great}},

	Kizoar: {{65, Vigilant},
		{75, Strong},
		{85, Mean},
		{102, Vicious},
		{117, Dreaded},
		{127, Master},
		{137, Veteran},
		{147, Hungry},
		{160, Nettled},
		{169, Incensed},
		{179, Mighty},
		{189, Awesome},
		{207, Gruesome},
		{217, Lacerating},
		{227, Killer},
		{236, Great},
		{236, Lethal}},

	Lumper: {{110, Gruff},
		{120, Obstinate},
		{130, Rooting},
		{139, Scrounging},
		{151, Nettled},
		{161, Vicious},
		{170, Feral},
		{179, Violent},
		{206, Fierce},
		{215, Ferocious},
		{224, Voracious}},

	Madakam: {{142, Nettled},
		{151, Vicious},
		{160, Feral},
		{169, Violent},
		{186, Furious},
		{235, Ferocious},
		{250, Voracious}},

	Mektoub: {{11, Suckling},
		{19, Weanling},
		{26, Docile},
		{23, Timorous},
		{34, Roaming},
		{44, Scampering},
		{59, Sprightly},
		{73, Robust},
		{83, Wary},
		{94, Gluttonous},
		{106, Grunting},
		{120, Gruff},
		{130, Obstinate},
		{140, Rooting},
		{150, Scrounging},
		{162, Nettled},
		{172, Vicious},
		{181, Feral},
		{191, Violent},
		{209, Furious},
		{219, Fierce},
		{228, Ferocious},
		{238, Voracious}},

	Messab: {{8, Suckling},
		{15, Weanling},
		{15, Timorous},
		{24, Roaming},
		{34, Scampering},
		{45, Sprightly},
		{57, Robust},
		{66, Wary},
		{76, Gluttonous},
		{85, Grunting},
		{101, Gruff},
		{110, Obstinate},
		{118, Rooting},
		{127, Scrounging}},

	Najab: {{151, Prowling},
		{169, Preying},
		{179, Raging},
		{189, Fearsome},
		{199, Huge},
		{218, Terrifying},
		{228, Awesome},
		{238, Horrific},
		{248, Great}},

	Ocyx: {{157, 0},
		{167, 0},
		{176, 0},
		{186, 0},
		{204, 0},
		{213, 0},
		{223, 0},
		{232, 0}},

	Ploderos: {{142, 0},
		{151, 0},
		{180, 0},
		{190, 0},
		{235, 0},
		{250, 0}},

	Psykopla: {{9, Sluggish},
		{21, Drowsy},
		{32, Minor},
		{44, Budding},
		{60, Robust},
		{71, Blooming},
		{81, Dehydrated},
		{93, Famished},
		{107, Parched},
		{117, Stalking},
		{127, Rooting},
		{149, Vile},
		{159, Unpredictable},
		{168, Prickly},
		{178, Rattled}},

	Ragus: {{11, 0},
		{19, 0},
		{17, 0},
		{30, 0},
		{40, 0},
		{52, 0},
		{68, 0},
		{79, 0},
		{90, 0},
		{101, 0},
		{121, 0},
		{132, 0},
		{142, 0},
		{153, 0}},

	Raspal: {{72, 0},
		{82, 0},
		{93, 0},
		{104, 0},
		{117, 0},
		{127, 0},
		{137, 0},
		{147, 0}},

	Rendor: {{12, 0},
		{16, 0},
		{56, 0},
		{67, 0},
		{77, 0},
		{87, 0},
		{101, 0},
		{110, 0},
		{118, 0},
		{127, 0}},

	Shalah: {{158, 0},
		{167, 0},
		{180, 0},
		{190, 0},
		{235, 0},
		{250, 0}},

	Shooki: {{12, Sluggish},
		{23, Drowsy},
		{34, Minor},
		{46, Budding},
		{64, Robust},
		{74, Blooming},
		{85, Dehydrated},
		{97, Famished},
		{113, Parched},
		{123, Stalking},
		{133, Rooting},
		{144, Noxious}},

	Slaveni: {{15, Placid},
		{18, Peaceful},
		{14, Sluggish},
		{26, Drowsy},
		{38, Minor},
		{51, Budding},
		{63, Robust},
		{73, Blooming},
		{84, Dehydrated},
		{95, Famished},
		{155, Vile},
		{165, Unpredictable},
		{175, Prickly},
		{185, Rattled},
		{204, Horrifying},
		{214, Rapacious},
		{224, Prodigeous},
		{233, Devastating}},

	Stinga: {{13, Sluggish},
		{24, Drowsy},
		{36, Morbid},
		{36, Minor},
		{48, Budding},
		{61, Robust},
		{75, Blooming},
		{86, Dehydrated},
		{97, Famished},
		{115, Parched},
		{125, Stalking},
		{136, Rooting},
		{146, Noxious},
		{233, Devastating}},

	Timari: {{110, 0},
		{120, 0},
		{130, 0},
		{140, 0},
		{151, 0},
		{161, 0},
		{170, 0},
		{179, 0},
		{197, 0},
		{206, 0},
		{215, 0},
		{225, 0}},

	Torbak: {{75, 0},
		{86, 0},
		{97, 0},
		{109, 0},
		{127, 0},
		{138, 0},
		{148, 0},
		{159, 0},
		{171, 0},
		{181, 0},
		{191, 0},
		{201, 0},
		{220, 0},
		{230, 0},
		{240, 0}},

	Tyrancha: {{163, 0},
		{173, 0},
		{183, 0},
		{193, 0},
		{210, 0},
		{220, 0},
		{230, 0},
		{240, 0}},

	Varinx: {{127, 0},
		{138, 0},
		{148, 0},
		{159, 0},
		{172, 0},
		{182, 0},
		{192, 0},
		{202, 0}},

	Vorax: {{219, 0},
		{229, 0},
		{239, 0},
		{249, 0}},

	Wombai: {{147, 0},
		{156, 0},
		{165, 0},
		{173, 0},
		{190, 0},
		{199, 0},
		{208, 0}},

	Yber: {{14, 0},
		{21, 0},
		{20, 0},
		{31, 0},
		{41, 0},
		{52, 0},
		{64, 0},
		{75, 0},
		{82, 0},
		{95, 0}},

	Yelk: {{20, 0},
		{103, 0},
		{113, 0},
		{122, 0},
		{132, 0},
		{142, 0},
		{151, 0},
		{160, 0},
		{169, 0},
		{186, 0},
		{195, 0},
		{204, 0},
		{213, 0}},

	Yetin: {{205, 0},
		{214, 0},
		{224, 0},
		{233, 0}},

	Yubo: {{5, 0},
		{11, 0},
		{13, 0},
		{24, 0},
		{34, 0},
		{46, 0}},

	Zerx: {{113, 0},
		{123, 0},
		{133, 0},
		{143, 0},
		{155, 0},
		{165, 0},
		{174, 0},
		{184, 0},
		{202, 0},
		{211, 0},
		{221, 0},
		{230, 0}},
}

var Classes = []Class{
	Herbivore, Herbivore, Herbivore, Herbivore, Herbivore, Herbivore, Carnivore,
	Flora, Herbivore, PrimitiveTribe, Carnivore, PrimitiveTribe, Herbivore,
	PrimitiveTribe, Carnivore, Herbivore, Carnivore, Herbivore, Carnivore,
	Avian, Avian, Javan, Flora, Carnivore, Kitin, Kitin, Kitin, Kitin, Kitin,
	Kitin, Kitin, KitinInvader, Kitin, Kitin, Kitin, Herbivore, Herbivore,
	Herbivore, Herbivore, Carnivore, Carnivore, Herbivore, Flora, Carnivore,
	Herbivore, Herbivore, Herbivore, Flora, Flora, Flora, Herbivore,
	Carnivore, Carnivore, Carnivore, Carnivore, Herbivore, Avian, Herbivore,
	Carnivore, Herbivore, Carnivore,
}

func StringToMob(mobName string) Mob {
	return MobMap[mobName]
}

func StringToName(name string) Name {
	return NameMap[name]
}

func StringToClass(class string) Class {
	return ClassMap[class]
}

func (mob Mob) Coefficient() float64 {
	return Coefficients[mob]
}

func (creature Creature) Coefficient() float64 {
	return creature.Mob.Coefficient()
}

func RoundLevel(mob Mob, lvl int) (Level, error) {
	for _, level := range Levels[mob] {
		if level.Level >= lvl {
			return level, nil
		}
	}

	// TODO return error
	return Level{}, errors.New("No suitable level found")
}
