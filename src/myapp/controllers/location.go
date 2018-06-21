package controllers

import "github.com/astaxie/beego"
import (
       "myapp/models"
       "encoding/json"
)

type LocationController struct {
       beego.Controller
}
// @Title 获得所有地址
// @Description 返回所有的地址数据
// @Success 200 {object} models.Location
// @router / [get]
func (u *LocationController) GetAll() {
       ss := models.GetAllLocations()
       u.Data["json"] = ss
       u.ServeJSON()
}
// @Title 获得一个地址
// @Description 返回某地址数据
// @Param      id            path   int    true          "The key for staticblock"
// @Success 200 {object} models.Location
// @router /:id [get]
func (u *LocationController) GetById() {
       id ,_:= u.GetInt(":id")
       s := models.GetLocationById(id)
       u.Data["json"] = s
       u.ServeJSON()
}
// @Title 创建地址
// @Description 创建地址的描述
// @Param      body          body   models.Location true          "body for location content"
// @Success 200 {int} models.Location.Id
// @Failure 403 body is empty
// @router / [post]
func (u *LocationController) Post() {
       var s models.Location
       json.Unmarshal(u.Ctx.Input.RequestBody, &s)
       uid := models.AddLocation(&s)
       u.Data["json"] = uid
       u.ServeJSON()
}
// @Title 修改地址
// @Description 修改地址的内容
// @Param      body          body   models.Location true          "body for location content"
// @Success 200 {int} models.Location
// @Failure 403 body is empty
// @router / [put]
func (u *LocationController) Update() {
       var s models.Location
       json.Unmarshal(u.Ctx.Input.RequestBody, &s)
       models.UpdateLocation(&s)
       u.Data["json"] = s
       u.ServeJSON()
}
// // @Title 删除一个地址
// // @Description 删除某地址数据
// // @Param      id            path   int    true          "The key for staticblock"
// // @Success 200 {object} models.Location
// // @router /:id [delete]
// func (u *LocationController) Delete() {
//        id ,_:= u.GetInt(":id")
//        models.DeleteLocation(id)
//        u.Data["json"] = true
//        u.ServeJSON()
// }