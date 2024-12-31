package ParkManager

type Manager struct {
	catID int
	dogID int
}

func (pm Manager) GiveId(animalType string) int {

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
