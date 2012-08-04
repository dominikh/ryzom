package api

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/dominikh/ryzom/assets"
	"github.com/dominikh/ryzom/material"
	"io/ioutil"
	"net/http"
)

var ItemNames map[string]string

func init() {
	dec := json.NewDecoder(bytes.NewReader(assets.Item_names_json()))
	err := dec.Decode(&ItemNames)
	if err != nil {
		panic(err)
	}
}

type Character struct {
	APIKey      string
	Inventories map[string]Inventory
}

type Inventory struct {
	Name         string
	CreationTime uint // unix timestamp
	ExpireTime   uint // unix timestamp
	Items        []*Item
}

type Item struct {
	Slot     uint `xml:"slot,attr"`
	Name     string
	Quantity uint16 `xml:"s,attr"`
	Quality  uint16 `xml:"q,attr"`
	SheetID  string `xml:",chardata"`
	Uses     []material.Use
}

type MaterialItem struct {
	*Item
	Material  *material.Material
	Grade     material.Grade
	Ecosystem material.Ecosystem
}

func (i *MaterialItem) String() string {
	return fmt.Sprintf("%dx Q%d %s %s %s", i.Quantity, i.Quality, i.Grade, i.Ecosystem, i.Material.Name)
}

func (i *Item) IconURL() string {
	return fmt.Sprintf("http://atys.ryzom.com/api/item_icon.php?sheetid=%s&q=%d&s=%d", i.SheetID, i.Quality, i.Quantity)
}

func (i *Item) ToMaterialItem() (ret *MaterialItem, ok bool) {
	if i.SheetID[0] != 'm' {
		return ret, false
	}

	materialId := i.SheetID[1:5]
	typeId := i.SheetID[5]
	ecosystemId := i.SheetID[8]
	gradeId := i.SheetID[9]

	var ecosystem material.Ecosystem
	var grade material.Grade

	switch ecosystemId {
	case 'd':
		ecosystem = "Desert"
	case 'f':
		ecosystem = "Forest"
	case 'j':
		ecosystem = "Jungle"
	case 'l':
		ecosystem = "Lakes"
	case 'p':
		ecosystem = "Prime Roots"
	case 'c':
		ecosystem = "Generic"
	}

	if typeId == 'd' {
		switch gradeId {
		case 'b':
			grade = "Basic"
		case 'c':
			grade = "Fine"
		case 'd':
			grade = "Choice"
		case 'e':
			grade = "Excellent"
		case 'f':
			grade = "Supreme"
		}
	} else {
		switch gradeId {
		case 'a':
			grade = "Basic"
		case 'b':
			grade = "Fine"
		case 'c':
			grade = "Choice"
		case 'd':
			grade = "Excellent"
		case 'e':
			grade = "Supreme"
		}
	}

	material, ok := material.Materials.FindByID(materialId)

	if !ok {
		return ret, false
	}

	return &MaterialItem{i, material, grade, ecosystem}, true
}

func NewCharacter(key string) (*Character, error) {
	character := &Character{key, make(map[string]Inventory)}

	url := fmt.Sprintf("http://atys.ryzom.com/api/character.php?key=%s&part=items", key)
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	type xmlInventory struct {
		Items []*Item `xml:"item"`
	}

	var xmlInventories struct {
		CacheInfo struct {
			Created uint `xml:"created,attr"`
			Expires uint `xml:"expire,attr"`
		} `xml:"cache"`
		Bag  xmlInventory `xml:"inventories>bag"`
		Room xmlInventory `xml:"room"`
		Pet1 xmlInventory `xml:"inventories>pet_animal1"`
		Pet2 xmlInventory `xml:"inventories>pet_animal2"`
		Pet3 xmlInventory `xml:"inventories>pet_animal3"`
		Pet4 xmlInventory `xml:"inventories>pet_animal4"`
	}

	err = xml.Unmarshal(data, &xmlInventories)

	if err != nil {
		return nil, err
	}

	for name, inventory := range map[string]xmlInventory{
		"bag":  xmlInventories.Bag,
		"room": xmlInventories.Room,
		"pet1": xmlInventories.Pet1,
		"pet2": xmlInventories.Pet2,
		"pet3": xmlInventories.Pet3,
		"pet4": xmlInventories.Pet4,
	} {
		for _, item := range inventory.Items {
			item.Name = ItemNames[item.SheetID]

			matItem, ok := item.ToMaterialItem()
			if ok {
				item.Uses = make([]material.Use, len(matItem.Material.Uses))
				copy(item.Uses, matItem.Material.Uses)
			} else {
				item.Uses = make([]material.Use, 0)
			}

			if item.Quantity == 0 {
				item.Quantity = 1
			}
		}

		character.Inventories[name] = Inventory{name, xmlInventories.CacheInfo.Created, xmlInventories.CacheInfo.Expires, inventory.Items}
	}

	return character, nil
}
