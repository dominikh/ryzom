package material

import (
	"encoding/json"
	"io"
	"github.com/dominikh/ryzom/assets"
	"bytes"
)

var Materials MaterialSlice

func init() {
	r := bytes.NewReader(assets.Materials_json())
	Materials = ReadMaterials(r)
}

// A material is something like "Adrial Bark" – It describes the name
// (Adrial Bark), the underlying material type (Bark), its uses (Ammo
// Bullet, ...) and its colors for all grades and ecosystems.
// Note that the grade (e.g. Fine Adrial Bark) is NOT part of the
// material. The grade is an attribute of an Ingredient.

type Material struct {
	Name   Name
	Type   Type
	Stats  map[Use]StatsMap
	Uses   []Use
	Colors ColorMap
	ID     string
}

type jsonMaterial struct {
	Name   string                       `json:"name"`
	Stats  map[string]map[string][]int8 `json:"stats"`
	Uses   []string                     `json:"uses"`
	Colors map[string]map[string]string `json:"colors"`
	ID     string                       `json:"id"`
}

// Adrial Bark
type Name string

// Bark
type Type string

// Basic, Fine, Choice, Excellent, Supreme
type Grade string

func (g Grade) Index() int {
	switch g {
	case "Basic":
		return 0
	case "Fine":
		return 1
	case "Choice":
		return 2
	case "Excellent":
		return 3
	case "Supreme":
		return 4
	}

	panic("Unknown grade")
}

// Forest, Jungle, ...
type Ecosystem string
type Color string
type ColorMap map[Grade]map[Ecosystem]Color
type Use string
type StatsMap map[Grade][]int8
type MaterialSlice []*Material

var StatNames = map[Use][]string{
	"Ammo bullet":   {"Durability", "Lightness", "Damage", "Speed", "Range"},
	"Ammo jacket":   {"Durability", "Lightness", "Speed", "Range"},
	"Armor clip":    {"Durability", "Lightness", "Dodge modifier", "Parry modifier", "Protection factor", "Max slashing protection", "Max smashing protection", "Max piercing protection"},
	"Armor shell":   {"Durability", "Lightness", "Dodge modifier", "Parry modifier", "Protection factor", "Max slashing protection", "Max smashing protection", "Max piercing protection"},
	"Barrel":        {"Durability", "Lightness", "Sap load", "Damage", "Speed", "Range", "Dodge modifier", "Parry modifier", "Adversary dodge modifier", "Adversary parry modifier"},
	"Blade":         {"Durability", "Lightness", "Sap load", "Damage", "Speed", "Dodge modifier", "Parry modifier", "Adversary dodge modifier", "Adversary parry modifier"},
	"Clothes":       {"Durability", "Lightness", "Dodge modifier", "Parry modifier", "Protection factor", "Max slashing protection", "Max smashing protection", "Max piercing protection"},
	"Counterweight": {"Durability", "Lightness", "Sap load", "Speed", "Dodge modifier", "Parry modifier", "Adversary dodge modifier", "Adversary parry modifier"},
	"Explosive":     {"Durability", "Lightness", "Damage", "Speed", "Range"},
	"Firing pin":    {"Durability", "Lightness", "Sap load", "Damage", "Speed", "Range", "Dodge modifier", "Parry modifier", "Adversary dodge modifier", "Adversary parry modifier"},
	"Grip":          {"Durability", "Lightness", "Sap load", "Speed", "Dodge modifier", "Parry modifier", "Adversary dodge modifier", "Adversary parry modifier"},
	"Hammer":        {"Durability", "Lightness", "Sap load", "Damage", "Speed", "Dodge modifier", "Parry modifier", "Adversary dodge modifier", "Adversary parry modifier"},
	"Jewel setting": {"Durability", "Lightness", "Desert resistance", "Forest resistance", "Jungle resistance", "Lakes resistance", "Prime Roots resistance"},
	"Jewel":         {"Durability", "Lightness", "Acid protection", "Cold protection", "Rot protection", "Fire protection", "Shockwave protection", "Poison protection", "Electric protection"},
	"Lining":        {"Durability", "Lightness", "Dodge modifier", "Parry modifier", "Protection factor", "Max slashing protection", "Max smashing protection", "Max piercing protection"},
	"Magic focus":   {"Durability", "Lightness", "Sap load", "Elemental cast speed", "Elemental power", "Offensive affliction cast speed", "Offensive affliction power", "Defensive affliction cast speed", "Defensive affliction power", "Heal cast speed", "Heal power"},
	"Point":         {"Durability", "Lightness", "Sap load", "Damage", "Speed", "Dodge modifier", "Parry modifier", "Adversary dodge modifier", "Adversary protection modifier"},
	"Shaft":         {"Durability", "Lightness", "Sap load", "Damage", "Speed", "Dodge modifier", "Parry modifier", "Adversary dodge modifier", "Adversary parry modifier"},
	"Stuffing":      {"Durability", "Lightness", "Dodge modifier", "Parry modifier", "Protection factor", "Max slashing protection", "Max smashing protection", "Max piercing protection"},
	"Trigger":       {"Durability", "Lightness", "Sap load", "Speed", "Dodge modifier", "Parry modifier", "Adversary dodge modifier", "Adversary parry modifier"},
}

func (m *Material) String() string {
	return string(m.Name)
}

func (materials MaterialSlice) Index(material *Material) int {
	for k, mat := range materials {
		if mat == material {
			return k
		}
	}

	panic("Could not find material.")
}

func (materials MaterialSlice) FindByID(id string) (*Material, bool) {
	for _, material := range materials {
		if material.ID == id {
			return material, true
		}
	}

	return nil, false
}

func (materials MaterialSlice) FindByUse(use Use) MaterialSlice {
	found := make(MaterialSlice, 0, 15)
	for _, material := range materials {
		if material.HasUse(use) {
			found = append(found, material)
		}
	}

	return found
}

func (materials MaterialSlice) FindByName(name Name) (*Material, bool) {
	for _, material := range materials {
		if material.Name == name {
			return material, true
		}
	}

	return nil, false
}

func ReadMaterials(r io.Reader) MaterialSlice {
	dec := json.NewDecoder(r)
	materials := make(MaterialSlice, 0, 100)
	for {
		var jm jsonMaterial

		err := dec.Decode(&jm)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil // TODO error
		}
		materials = append(materials, jm.ToMaterial())
	}

	return materials
}

func (m *Material) HasUse(use Use) bool {
	for _, mUse := range m.Uses {
		if mUse == use {
			return true
		}
	}

	return false
}

func (jm *jsonMaterial) ToMaterial() *Material {
	material := &Material{}

	material.Name = Name(jm.Name)
	material.Uses = make([]Use, len(jm.Uses))
	for index, value := range jm.Uses {
		material.Uses[index] = Use(value)
	}

	material.Stats = make(map[Use]StatsMap)
	for use, statsMap := range jm.Stats {
		material.Stats[Use(use)] = make(StatsMap)
		for grade, stats := range statsMap {
			material.Stats[Use(use)][Grade(grade)] = stats
		}
	}

	material.Colors = make(ColorMap)
	for grade, colorMap := range jm.Colors {
		material.Colors[Grade(grade)] = make(map[Ecosystem]Color)
		for ecosystem, color := range colorMap {
			material.Colors[Grade(grade)][Ecosystem(ecosystem)] = Color(color)
		}
	}

	material.ID = jm.ID

	return material
}

func GradeIndexToGrade(grade int64) (res Grade) {
	switch grade {
	case 0:
		res = "Basic"
	case 1:
		res = "Fine"
	case 2:
		res = "Choice"
	case 3:
		res = "Excellent"
	case 4:
		res = "Supreme"
	}

	return
}
