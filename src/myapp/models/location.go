package models


type Location struct {
	Task_id int `json:"task_id"`
	Soldier_id int  `json:"soldier"`
	Longitude float64  `json:"longitude"`
	Latitude float64  `json:"latitude"`
}
var allLocations []Location;

func init(){
}

func GetAllLocations(id int) ([]*Location,bool) {
	var locations []*Location;
	flag := false;
	var i int;
	for i = 0; i < len(allLocations); i++ {
		if allLocations[i].Task_id == id {
			locations = append(locations, &allLocations[i]);
			flag = true
		}
	}
	return locations,flag
}
func GetLocationById(id1 int, id2 int) (Location,bool) {
	var loc Location;
	var i int;
	flag := false
	for i = 0; i < len(allLocations); i++ {
		if allLocations[i].Task_id == id1 && allLocations[i].Soldier_id == id2 {
			loc = allLocations[i];
			flag = true
			break;
		}
	}
	return loc,flag
}
func UpdateLocation(id1 int, id2 int, long1 float64, lati float64) {
	var i int;
	flag := false;
	index := 0;
	for i = 0; i < len(allLocations); i++ {
        if allLocations[i].Task_id == id1 && allLocations[i].Soldier_id == id2 {
			flag = true;
			index = i;
			break;
		}
	}
	if flag == false {
		loc := Location{id1, id2, long1, lati}
        allLocations = append(allLocations, loc);
	} else {
		allLocations[index].Longitude = long1;
		allLocations[index].Latitude = lati;
	}
}