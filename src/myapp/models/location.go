package models

// import (
// 	"fmt"
// )

type Location struct {
	Id int
	Task_id int
	Soldier_id int
	Longitude float64
	Latitude float64
}
func GetAllLocations() []*Location {
	// o := orm.NewOrm()
	// o.Using("default")
	// var locations []*Location
	// q := o.QueryTable("location")
	// q.All(&locations)
	var locations []*Location
    locations = []*Location{
		&Location{
			1,
			1,
			1,
			1.1,
			1.1,
		},
		&Location{
			1,
			2,
			3,
			4.4,
			4.5,
		},
	}
	return locations
}
func GetLocationById(id int) Location {
	// u := Location{Id:id}
	// o := orm.NewOrm()
	// o.Using("default")
	// err := o.Read(&u)
	// if err == orm.ErrNoRows {
	// 	fmt.Println("查询不到")
	// } else if err == orm.ErrMissPK {
	// 	fmt.Println("找不到主键")
	// }
	// return u
	loc := Location{
		      1,
		      1,
			  1,
			  1.1,
			  1.1,
			}
	return loc
}
func AddLocation(location *Location) int{
	//   o :=orm.NewOrm()
	//   o.Using("default")
	//   o.Insert(location)
	//   return location.Id
	return 1
}
func UpdateLocation(location *Location) {
	// o := orm.NewOrm()
	// o.Using("default")
	// o.Update(location)
}
// func DeleteLocation(id int) {
// 	o := orm.NewOrm()
// 	o.Using("default")
// 	o.Delete(&Location{Id:id})
// }
// func init() {
// 	orm.RegisterModel(new(Location))
// }