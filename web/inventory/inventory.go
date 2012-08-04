package main

import (
	"encoding/json"
	"fmt"
	"github.com/dominikh/ryzom/api"
	"github.com/dominikh/ryzom/material"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// TODO: Should probably break out the whole IconStore into a service
// of its own, independent from the inventory one.

var iconStore *IconStore = NewIconStore("cache/icons/")

type IconStore struct {
	sync.RWMutex
	cacheDir string
	cached   map[string]bool
}

func NewIconStore(cacheDir string) *IconStore {
	store := &IconStore{sync.RWMutex{}, cacheDir, make(map[string]bool)}
	store.populate()
	return store
}

func (store *IconStore) populate() {
	err := filepath.Walk(store.cacheDir, func(path string, _ os.FileInfo, err error) error {
		if err == nil {
			_, file := filepath.Split(path)
			log.Println("Found cached icon", file)
			store.markAsCached(file)

		}
		return nil
	})

	if err != nil {
		log.Panic(err)
	}
}

func (store *IconStore) markAsCached(id string) {
	store.Lock()
	defer store.Unlock()

	store.cached[id] = true
}

func (store *IconStore) unmarkAsCached(id string) {
	store.Lock()
	defer store.Unlock()

	store.cached[id] = false
}

func (store *IconStore) IsCached(sheetid, quality, quantity string) bool {
	store.RLock()
	defer store.RUnlock()

	return store.cached[store.attributesToID(sheetid, quality, quantity)]
}

func (store *IconStore) GetIcon(sheetid, quality, quantity string) ([]byte, error) {
	id := store.attributesToID(sheetid, quality, quantity)

	if store.IsCached(sheetid, quality, quantity) {
		// Load file from disk and serve it
		file, err := os.Open(store.cacheDir + id)
		if err != nil {
			if os.IsNotExist(err) {
				// Icon was marked as cached but cache cannot be found → refetch
				log.Println("Cached version missing, refetching:", id)
				store.unmarkAsCached(id)
				data, err := store.GetIcon(sheetid, quality, quantity)
				return data, err
			}
			// Serving the image failed for some other reason
			return []byte{}, err
		}

		data, err := ioutil.ReadAll(file)
		if err != nil {
			return []byte{}, err
		}

		return data, nil
	}

	// Download image, write to disk, mark as cached
	log.Println("Downloading icon", id)
	url := fmt.Sprintf("http://atys.ryzom.com/api/item_icon.php?sheetid=%s&q=%s&s=%s", sheetid, quality, quantity)
	resp, err := http.Get(url)

	if err != nil {
		return []byte{}, err
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return []byte{}, err
	}

	// Even if we downloaded the image multiple times due to racing,
	// they'll be written properly to disk
	store.Lock()
	file, err := os.Create(store.cacheDir + id)
	if err != nil {
		return []byte{}, err
	}
	defer file.Close()

	_, err = file.Write(data)
	store.Unlock()
	if err != nil {
		return []byte{}, err
	}

	go store.markAsCached(id)

	return data, nil
}

func (store *IconStore) attributesToID(sheetid, quality, quantity string) string {
	return strings.Join([]string{sheetid, quality, quantity}, "_")
}

func itemIconHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	qual := r.FormValue("qual")
	quan := r.FormValue("quan")

	data, err := iconStore.GetIcon(id, qual, quan)
	if err != nil {
		log.Println("Error serving icon:", err)
		return
	}

	w.Write(data)
}

func inventoryHandler(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("key")

	log.Println("Loading inventory for", key)

	character, _ := api.NewCharacter(key)
	json, _ := json.Marshal(character.Inventories)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(json)
}

func materialItemHandler(w http.ResponseWriter, r *http.Request) {
	sheetid := r.FormValue("id")

	item := &api.Item{0, "", 0, 0, sheetid, make([]material.Use, 0)}
	materialItem, _ := item.ToMaterialItem()

	json, _ := json.Marshal(materialItem)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(json)
}

func main() {
	err := os.MkdirAll("cache/icons/", 0700)
	if err != nil {
		log.Fatal("Could not create icon cache directory:", err)
	}

	// Serve item icons
	http.HandleFunc("/item_icon/", itemIconHandler)

	// Serve entire inventory representations as JSON
	http.HandleFunc("/inventory/", inventoryHandler)

	// Serve material information for a given Sheet ID. Note: Not yet
	// used in the UI.
	http.HandleFunc("/material_item/", materialItemHandler)

	port := os.Getenv("GO_HTTP_PORT")
	if len(port) == 0 {
		port = "8080"
	}

	port = ":" + port
	log.Println("Server Listening on", port)
	http.ListenAndServe(port, nil)
}
