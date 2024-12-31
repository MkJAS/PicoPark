package ParkManager

import (
	"sync"
)

var lock = &sync.Mutex{}

type Manager struct {
	catID int
	dogID int
}

var managerInstance *Manager

func GetInstance() *Manager {
	if managerInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if managerInstance == nil {
			//fmt.Println("Creating single instance now.")
			managerInstance = &Manager{}
		}
	}

	return managerInstance
}

func (pm *Manager) GiveId(animalType string) int {

	switch animalType {
	case "cat":
		pm.catID++
		return pm.catID
	case "dog":
		pm.dogID++
		return pm.dogID
	}
	return 0
}
