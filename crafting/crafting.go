package crafting

import (
	"fmt"
	"github.com/dominikh/ryzom/item"
	"github.com/dominikh/ryzom/material"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type InvalidityReason uint

const (
	Valid InvalidityReason = iota
	UnrequiredMaterial
	WrongMaterialUse
	TooManyMaterials
	NotEnoughMaterials
)

type Plan struct {
	Quality      item.Quality
	ItemType     item.Type
	Requirements map[material.Use]uint8
}

type Recipe struct {
	Plan        *Plan
	Ingredients []Ingredient
}

type Ingredient struct {
	Material *material.Material
	Use      material.Use
	Grade    material.Grade
	Quantity uint8
}

var Plans = map[string]Plan{
	// Axe
	"Basic Axe": Plan{"Basic", "Axe", map[material.Use]uint8{
		"Blade":         5,
		"Shaft":         3,
		"Grip":          3,
		"Counterweight": 3,
	}},
	"Medium Axe": Plan{"Medium", "Axe", map[material.Use]uint8{
		"Blade":         5,
		"Shaft":         5,
		"Grip":          5,
		"Counterweight": 4,
	}},
	"High Axe": Plan{"High", "Axe", map[material.Use]uint8{
		"Blade":         6,
		"Shaft":         6,
		"Grip":          6,
		"Counterweight": 5,
	}},

	// Dagger
	"Basic Dagger": Plan{"Basic", "Dagger", map[material.Use]uint8{
		"Blade":         2,
		"Shaft":         2,
		"Grip":          1,
		"Counterweight": 1,
	}},
	"Medium Dagger": Plan{"Medium", "Dagger", map[material.Use]uint8{
		"Blade":         2,
		"Shaft":         2,
		"Grip":          2,
		"Counterweight": 2,
	}},
	"High Dagger": Plan{"High", "Dagger", map[material.Use]uint8{
		"Blade":         3,
		"Shaft":         3,
		"Grip":          2,
		"Counterweight": 2,
	}},

	// Lance
	"Basic Lance": Plan{"Basic", "Lance", map[material.Use]uint8{
		"Point": 5,
		"Shaft": 5,
		"Grip":  4,
	}},
	"Medium Lance": Plan{"Medium", "Lance", map[material.Use]uint8{
		"Point": 7,
		"Shaft": 6,
		"Grip":  6,
	}},
	"High Lance": Plan{"High", "Lance", map[material.Use]uint8{
		"Point": 8,
		"Shaft": 8,
		"Grip":  7,
	}},

	// Mace
	"Basic Mace": Plan{"Basic", "Mace", map[material.Use]uint8{
		"Hammer": 5,
		"Shaft":  3,
		"Grip":   3,
	}},
	"Medium Mace": Plan{"Medium", "Mace", map[material.Use]uint8{
		"Hammer": 5,
		"Shaft":  5,
		"Grip":   5,
	}},
	"High Mace": Plan{"High", "Mace", map[material.Use]uint8{
		"Hammer": 6,
		"Shaft":  6,
		"Grip":   6,
	}},

	// Staff
	"Basic Staff": Plan{"Basic", "Staff", map[material.Use]uint8{
		"Shaft": 7,
		"Grip":  7,
	}},
	"Medium Staff": Plan{"Medium", "Staff", map[material.Use]uint8{
		"Shaft": 10,
		"Grip":  9,
	}},
	"High Staff": Plan{"High", "Staff", map[material.Use]uint8{
		"Shaft": 12,
		"Grip":  11,
	}},

	// Sword
	"Basic Sword": Plan{"Basic", "Sword", map[material.Use]uint8{
		"Blade":         4,
		"Shaft":         4,
		"Grip":          3,
		"Counterweight": 3,
	}},
	"Medium Sword": Plan{"Medium", "Sword", map[material.Use]uint8{
		"Blade":         6,
		"Shaft":         5,
		"Grip":          5,
		"Counterweight": 4,
	}},
	"High Sword": Plan{"High", "Sword", map[material.Use]uint8{
		"Blade":         6,
		"Shaft":         6,
		"Grip":          6,
		"Counterweight": 5,
	}},

	// Pike
	"Basic Pike": Plan{"Basic", "Pike", map[material.Use]uint8{
		"Point": 7,
		"Shaft": 6,
		"Grip":  7,
	}},
	"Medium Pike": Plan{"Medium", "Pike", map[material.Use]uint8{
		"Point": 9,
		"Shaft": 9,
		"Grip":  9,
	}},
	"High Pike": Plan{"High", "Pike", map[material.Use]uint8{
		"Point": 11,
		"Shaft": 11,
		"Grip":  11,
	}},

	// Long axe
	"Basic Long axe": Plan{"Basic", "Long axe", map[material.Use]uint8{
		"Blade":         5,
		"Shaft":         5,
		"Grip":          5,
		"Counterweight": 5,
	}},
	"Medium Long axe": Plan{"Medium", "Long axe", map[material.Use]uint8{
		"Blade":         7,
		"Shaft":         7,
		"Grip":          7,
		"Counterweight": 6,
	}},
	"High Long axe": Plan{"High", "Long axe", map[material.Use]uint8{
		"Blade":         9,
		"Shaft":         8,
		"Grip":          8,
		"Counterweight": 8,
	}},

	// Long mace
	"Basic Long mace": Plan{"Basic", "Long mace", map[material.Use]uint8{
		"Hammer":        5,
		"Shaft":         5,
		"Grip":          5,
		"Counterweight": 5,
	}},
	"Medium Long mace": Plan{"Medium", "Long mace", map[material.Use]uint8{
		"Hammer":        7,
		"Shaft":         7,
		"Grip":          7,
		"Counterweight": 6,
	}},
	"High Long mace": Plan{"High", "Long mace", map[material.Use]uint8{
		"Hammer":        9,
		"Shaft":         8,
		"Grip":          8,
		"Counterweight": 8,
	}},

	// Long sword
	"Basic Long sword": Plan{"Basic", "Long sword", map[material.Use]uint8{
		"Blade":         5,
		"Shaft":         5,
		"Grip":          5,
		"Counterweight": 5,
	}},
	"Medium Long sword": Plan{"Medium", "Long sword", map[material.Use]uint8{
		"Blade":         7,
		"Shaft":         7,
		"Grip":          7,
		"Counterweight": 6,
	}},
	"High Long sword": Plan{"High", "Long sword", map[material.Use]uint8{
		"Blade":         9,
		"Shaft":         8,
		"Grip":          8,
		"Counterweight": 8,
	}},

	// Autolauncher
	"Basic Autolauncher": Plan{"Basic", "Autolauncher", map[material.Use]uint8{
		"Barrel":     3,
		"Trigger":    3,
		"Shaft":      3,
		"Firing pin": 3,
	}},
	"Medium Autolauncher": Plan{"Medium", "Autolauncher", map[material.Use]uint8{
		"Barrel":     4,
		"Trigger":    4,
		"Shaft":      4,
		"Firing pin": 4,
	}},
	"High Autolauncher": Plan{"High", "Autolauncher", map[material.Use]uint8{
		"Barrel":     5,
		"Trigger":    5,
		"Shaft":      5,
		"Firing pin": 5,
	}},

	// Bowpistol
	"Basic Bowpistol": Plan{"Basic", "Bowpistol", map[material.Use]uint8{
		"Barrel":     2,
		"Trigger":    1,
		"Shaft":      1,
		"Firing pin": 1,
	}},
	"Medium Bowpistol": Plan{"Medium", "Bowpistol", map[material.Use]uint8{
		"Barrel":     2,
		"Trigger":    2,
		"Shaft":      1,
		"Firing pin": 2,
	}},
	"High Bowpistol": Plan{"High", "Bowpistol", map[material.Use]uint8{
		"Barrel":     2,
		"Trigger":    2,
		"Shaft":      2,
		"Firing pin": 2,
	}},

	// Bowrifle
	"Basic Bowrifle": Plan{"Basic", "Bowrifle", map[material.Use]uint8{
		"Barrel":     4,
		"Trigger":    2,
		"Shaft":      2,
		"Firing pin": 2,
	}},
	"Medium Bowrifle": Plan{"Medium", "Bowrifle", map[material.Use]uint8{
		"Barrel":     2,
		"Trigger":    2,
		"Shaft":      1,
		"Firing pin": 2,
	}},
	"High Bowrifle": Plan{"High", "Bowrifle", map[material.Use]uint8{
		"Barrel":     4,
		"Trigger":    4,
		"Shaft":      4,
		"Firing pin": 4,
	}},

	// Launcher
	"Basic Launcher": Plan{"Basic", "Launcher", map[material.Use]uint8{
		"Barrel":     3,
		"Trigger":    3,
		"Shaft":      3,
		"Firing pin": 3,
	}},
	"Medium Launcher": Plan{"Medium", "Launcher", map[material.Use]uint8{
		"Barrel":     4,
		"Trigger":    4,
		"Shaft":      4,
		"Firing pin": 4,
	}},
	"High Launcher": Plan{"High", "Launcher", map[material.Use]uint8{
		"Barrel":     5,
		"Trigger":    5,
		"Shaft":      5,
		"Firing pin": 5,
	}},

	// Pistol
	"Basic Pistol": Plan{"Basic", "Pistol", map[material.Use]uint8{
		"Barrel":     1,
		"Trigger":    1,
		"Shaft":      1,
		"Firing pin": 1,
	}},
	"Medium Pistol": Plan{"Medium", "Pistol", map[material.Use]uint8{
		"Barrel":     2,
		"Trigger":    1,
		"Shaft":      1,
		"Firing pin": 1,
	}},
	"High Pistol": Plan{"High", "Pistol", map[material.Use]uint8{
		"Barrel":     2,
		"Trigger":    2,
		"Shaft":      1,
		"Firing pin": 2,
	}},

	// Rifle
	"Basic Rifle": Plan{"Basic", "Rifle", map[material.Use]uint8{
		"Barrel":     2,
		"Trigger":    2,
		"Shaft":      2,
		"Firing pin": 2,
	}},
	"Medium Rifle": Plan{"Medium", "Rifle", map[material.Use]uint8{
		"Barrel":     3,
		"Trigger":    3,
		"Shaft":      2,
		"Firing pin": 3,
	}},
	"High Rifle": Plan{"High", "Rifle", map[material.Use]uint8{
		"Barrel":     4,
		"Trigger":    3,
		"Shaft":      3,
		"Firing pin": 3,
	}},

	// Launcher ammo
	"Basic Launcher ammo": Plan{"Basic", "Launcher ammo", map[material.Use]uint8{
		"Ammo bullet": 4,
		"Ammo jacket": 4,
		"Explosive":   4,
	}},

	// Bowpistol ammo
	"Basic Bowpistol ammo": Plan{"Basic", "Bowpistol ammo", map[material.Use]uint8{
		"Ammo bullet": 2,
		"Ammo jacket": 2,
		"Explosive":   2,
	}},

	// Bowrifle ammo
	"Basic Bowrifle ammo": Plan{"Basic", "Bowrifle ammo", map[material.Use]uint8{
		"Ammo bullet": 2,
		"Ammo jacket": 2,
		"Explosive":   2,
	}},

	// Autolauncher ammo
	"Basic Autolauncher ammo": Plan{"Basic", "Autolauncher ammo", map[material.Use]uint8{
		"Ammo bullet": 2,
		"Ammo jacket": 2,
		"Explosive":   2,
	}},

	// Pistol ammo
	"Basic Pistol ammo": Plan{"Basic", "Pistol ammo", map[material.Use]uint8{
		"Ammo bullet": 1,
		"Ammo jacket": 2,
		"Explosive":   1,
	}},

	// Rifle ammo
	"Basic Rifle ammo": Plan{"Basic", "Rifle ammo", map[material.Use]uint8{
		"Ammo bullet": 1,
		"Ammo jacket": 2,
		"Explosive":   1,
	}},

	// Magic amplifiers
	"Basic Magic amplifiers": Plan{"Basic", "Magic amplifiers", map[material.Use]uint8{
		"Shaft":       5,
		"Grip":        5,
		"Magic focus": 10,
	}},
	"Medium Magic amplifiers": Plan{"Medium", "Magic amplifiers", map[material.Use]uint8{
		"Shaft":       6,
		"Grip":        6,
		"Magic focus": 15,
	}},
	"High Magic amplifiers": Plan{"High", "Magic amplifiers", map[material.Use]uint8{
		"Shaft":       7,
		"Grip":        6,
		"Magic focus": 20,
	}},

	// Jewels
	"Basic Anklet": Plan{"Basic", "Anklet", map[material.Use]uint8{
		"Jewel setting": 3,
		"Jewel":         3,
	}},
	"Medium Anklet": Plan{"Medium", "Anklet", map[material.Use]uint8{
		"Jewel setting": 4,
		"Jewel":         4,
	}},
	"High Anklet": Plan{"High", "Anklet", map[material.Use]uint8{
		"Jewel setting": 5,
		"Jewel":         5,
	}},

	"Basic Bracelet": Plan{"Basic", "Bracelet", map[material.Use]uint8{
		"Jewel setting": 3,
		"Jewel":         3,
	}},
	"Medium Bracelet": Plan{"Medium", "Bracelet", map[material.Use]uint8{
		"Jewel setting": 4,
		"Jewel":         4,
	}},
	"High Bracelet": Plan{"High", "Bracelet", map[material.Use]uint8{
		"Jewel setting": 5,
		"Jewel":         5,
	}},

	"Basic Diadem": Plan{"Basic", "Diadem", map[material.Use]uint8{
		"Jewel setting": 3,
		"Jewel":         3,
	}},
	"Medium Diadem": Plan{"Medium", "Diadem", map[material.Use]uint8{
		"Jewel setting": 4,
		"Jewel":         4,
	}},
	"High Diadem": Plan{"High", "Diadem", map[material.Use]uint8{
		"Jewel setting": 5,
		"Jewel":         5,
	}},

	"Basic Earring": Plan{"Basic", "Earring", map[material.Use]uint8{
		"Jewel setting": 3,
		"Jewel":         3,
	}},
	"Medium Earring": Plan{"Medium", "Earring", map[material.Use]uint8{
		"Jewel setting": 4,
		"Jewel":         4,
	}},
	"High Earring": Plan{"High", "Earring", map[material.Use]uint8{
		"Jewel setting": 5,
		"Jewel":         5,
	}},

	"Basic Pendant": Plan{"Basic", "Pendant", map[material.Use]uint8{
		"Jewel setting": 3,
		"Jewel":         3,
	}},
	"Medium Pendant": Plan{"Medium", "Pendant", map[material.Use]uint8{
		"Jewel setting": 4,
		"Jewel":         4,
	}},
	"High Pendant": Plan{"High", "Pendant", map[material.Use]uint8{
		"Jewel setting": 5,
		"Jewel":         5,
	}},

	"Basic Ring": Plan{"Basic", "Ring", map[material.Use]uint8{
		"Jewel setting": 3,
		"Jewel":         3,
	}},
	"Medium Ring": Plan{"Medium", "Ring", map[material.Use]uint8{
		"Jewel setting": 4,
		"Jewel":         4,
	}},
	"High Ring": Plan{"High", "Ring", map[material.Use]uint8{
		"Jewel setting": 5,
		"Jewel":         5,
	}},

	// Buckler
	"Basic Buckler": Plan{"Basic", "Buckler", map[material.Use]uint8{
		"Armor clip":  3,
		"Armor shell": 4,
	}},
	"Medium Buckler": Plan{"Medium", "Buckler", map[material.Use]uint8{
		"Armor clip":  4,
		"Armor shell": 5,
	}},
	"High Buckler": Plan{"High", "Buckler", map[material.Use]uint8{
		"Armor clip":  5,
		"Armor shell": 6,
	}},

	// Shield
	"Basic Shield": Plan{"Basic", "Shield", map[material.Use]uint8{
		"Armor clip":  6,
		"Armor shell": 7,
	}},
	"Medium Shield": Plan{"Medium", "Shield", map[material.Use]uint8{
		"Armor clip":  8,
		"Armor shell": 9,
	}},
	"High Shield": Plan{"High", "Shield", map[material.Use]uint8{
		"Armor clip":  10,
		"Armor shell": 11,
	}},
}

func (r *Recipe) IsValid() (bool, InvalidityReason) {
	used := make(map[material.Use]uint8)

	for requirement, _ := range r.Plan.Requirements {
		used[requirement] = 0
	}

	for _, ingredient := range r.Ingredients {
		_, ok := r.Plan.Requirements[ingredient.Use]
		if !ok {
			// This type isn't required (allowed) in this plan
			return false, UnrequiredMaterial
		}
		if !ingredient.Material.HasUse(ingredient.Use) {
			// The recipe is lying about a material's type
			return false, WrongMaterialUse
		}

		used[ingredient.Use] += ingredient.Quantity

		if used[ingredient.Use] > r.Plan.Requirements[ingredient.Use] {
			// Using too many
			return false, TooManyMaterials
		}
	}

	for use, amount := range used {
		if amount < r.Plan.Requirements[use] {
			return false, NotEnoughMaterials
		}
	}

	return true, Valid
}

func (r *Recipe) PrecraftStats() map[string]float64 {
	statWeights := make(map[string]uint8)
	endStats := make(map[string]float64)

	// Count how many mats we have per stat
	for materialName, amount := range r.Plan.Requirements {
		statNames := material.StatNames[materialName]
		for _, stat := range statNames {
			statWeights[stat] += amount
		}
	}

	for _, ingredient := range r.Ingredients {
		stats := ingredient.Material.Stats[ingredient.Use][ingredient.Grade]
		for index, statValue := range stats {
			stat := material.StatNames[ingredient.Use][index] // Look up which stat it is
			factor := (float64(ingredient.Quantity) / float64(statWeights[stat]))
			endStats[stat] += factor * float64(statValue)
		}
	}

	return endStats
}

func (r *Recipe) PrecraftStatsBoosted() map[string]float64 {
	stats := r.PrecraftStats()

	var bestStatName string
	var meanStat, bestStat float64
	for stat, value := range stats {
		if value > bestStat {
			bestStatName, bestStat = stat, value
		}

		meanStat += value
	}

	meanStat /= float64(len(stats))

	stats[bestStatName] = bestStat

	maxDelta := bestStat - meanStat
	if maxDelta < 30 {
		stretchFactor := 30 / maxDelta
		if stretchFactor > 2 {
			stretchFactor = 2
		}

		for stat, value := range stats {
			statDelta := value - meanStat
			newValue := meanStat + statDelta*stretchFactor
			// FIXME this screws the mean
			if newValue < 0 {
				newValue = 0
			} else if newValue > 100 {
				newValue = 100
			}
			stats[stat] = newValue
		}
	}

	// TODO check if the boosting applies for the best stat before or after stretching
	if bestStat-meanStat >= 35 {
		// boost the best value

		bestStat += 10
		if bestStat > 100 {
			bestStat = 100
		}
	}

	// TODO: KC rounds to one digit after the comma, at least for
	// display. But to my understanding, Ryzom rounds down, so 99.99 =
	// 99, not 100?
	return stats
}

func (r *Recipe) PostcraftStats( ) (postcraft map[string]float64) {
	precraft := r.PrecraftStatsBoosted()
	postcraft = make(map[string]float64)

	switch(r.Plan.ItemType) {
	case "Light armor vest": // TODO this prolly applies to all LA parts?
		postcraft["Protection factor"] = precraft["Protection factor"] / 100.0 * 20 + 5
		postcraft["Parry modifier"] = precraft["Parry modifier"] / 100.0 * 1
		postcraft["Dodge modifier"] = precraft["Dodge modifier"] / 100.0 * 1 + 1
		postcraft["Max slashing protection"] = 0 // This depends on Quality, so we have to consider that. Also some other part I don't know
		postcraft["Max smashing protection"] = 0
		postcraft["Max piercing protection"] = 0
	}

	return
}

func (i Ingredient) genome(materials material.MaterialSlice) string {
	materialIndex := materials.Index(i.Material)
	gradeS := fmt.Sprintf("%d", i.Grade.Index())
	amountS := fmt.Sprintf("%d", i.Quantity)
	numLength := strconv.FormatInt(int64(len(strconv.FormatInt(int64(len(materials)), 10))), 10)
	materialIndexS := fmt.Sprintf("%0"+numLength+"d", materialIndex)

	useIndex := 0
	for k, use := range i.Material.Uses {
		if use == i.Use {
			useIndex = k
			break
		}
	}

	useIndexS := fmt.Sprintf("%02d", useIndex)

	return materialIndexS + gradeS + amountS + useIndexS
}

func (r *Recipe) genome(materials material.MaterialSlice) string {
	genomes := make([]string, len(r.Ingredients))
	for k, ingredient := range r.Ingredients {
		genomes[k] = ingredient.genome(materials)
	}

	return strings.Join(genomes, "")
}

func (p *Plan) FindParent(materials material.MaterialSlice) string {
	seed := time.Now().UnixNano()
	fmt.Println("Seed:", seed)
	rand.Seed(seed) // FIXME use own PRNG instance

	ingredients := make([]Ingredient, 0)

	for use, amount := range p.Requirements {
		possibleMaterials := materials.FindByUse(use)
		remainingAmount := amount
		for remainingAmount > 0 {
			randomAmount := rand.Int63n(int64(remainingAmount)) + 1
			randomIndex := rand.Int63n(int64(len(possibleMaterials)))
			randomGrade := rand.Int63n(5)

			mat := possibleMaterials[randomIndex]
			ingredient := Ingredient{mat, use, material.GradeIndexToGrade(randomGrade), uint8(randomAmount)}
			ingredients = append(ingredients, ingredient)

			remainingAmount -= uint8(randomAmount)
		}
	}

	recipe := &Recipe{p, ingredients}
	genes := recipe.genome(materials)
	fmt.Println("Genome:", genes)
	return genes
}

func (p *Plan) RoundStat(stat string, value float64) (ret float64) {
	// TODO check if float rounding is a problem
	switch stat {
	case "Parry modifier":
		switch p.ItemType {
		case "Heavy armor":
			switch {
			case value >= 100.0:
				ret = 100.0
			case value >= 75.0:
				ret = 75.0
			case value >= 50.0:
				ret = 50.0
			default:
				ret = 0
			}
		case "Medium armor":
			switch {
			case value >= 100.0:
				ret = 100.0
			case value >= 67.0:
				ret = 67.0
			default:
				ret = 0
			}
		case "Light armor":
			switch {
			case value >= 100.0:
				ret = 100.0
			default:
				ret = 0.0
			}
		default:
			ret = value
		}
	default:
		ret = value
	}

	return
}
