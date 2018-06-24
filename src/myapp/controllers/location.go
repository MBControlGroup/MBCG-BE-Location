package controllers

import "github.com/astaxie/beego"
import (
       "myapp/models"
       "strconv"
)

type LocationController struct {
       beego.Controller
}
// @Title 获得所有地址
// @Description 返回所有的地址数据
// @Success 200 {object} models.Location
// @router /getAllLocs [post]
func (u *LocationController) GetAll() {
       id := u.Input().Get("task_id");
       _id,err1 := strconv.Atoi(id); 
       if (err1 != nil){
        u.Data["json"] = map[string]interface{}{"code":400, "enmsg":"no", "cnmsg":"失败"};
        u.ServeJSON()
        return
       }
       ss,flag := models.GetAllLocations(_id);
       if (flag == true){
        u.Data["json"] = map[string]interface{}{"code":200, "enmsg":"ok", "cnmsg":"成功", "data":ss};
        u.ServeJSON()
       } else {
        u.Data["json"] = map[string]interface{}{"code":400, "enmsg":"not found", "cnmsg":"未找到"};
        u.ServeJSON()
       }
}
// @Title 获得一个地址
// @Description 返回某地址数据
// @Param      id            path   int    true          "The key for staticblock"
// @Success 200 {object} models.Location
// @router /getOneLocs [post]
func (u *LocationController) GetById() {
       id1 := u.Input().Get("task_id");
       id2 := u.Input().Get("soldier_id");
       _id1,err1 := strconv.Atoi(id1); 
       _id2,err2 := strconv.Atoi(id2); 
       if (err1 != nil || err2 != nil){
        u.Data["json"] = map[string]interface{}{"code":400, "enmsg":"no", "cnmsg":"失败"};
        u.ServeJSON()
        return
      }
       s,flag := models.GetLocationById(_id1,_id2);
       if (flag == true){
           u.Data["json"] = map[string]interface{}{"code":200, "enmsg":"ok", "cnmsg":"成功", "data":s};
           u.ServeJSON()
       } else {
           u.Data["json"] = map[string]interface{}{"code":400, "enmsg":"not found", "cnmsg":"未找到"};
           u.ServeJSON()
       }
}
// @Title 有则改之，无则建之
// @Description 修改地址的内容
// @Param      body          body   models.Location true          "body for location content"
// @Success 200 {int} models.Location
// @Failure 403 body is empty
// @router /upOneLocs [post]
func (u *LocationController) Update() {
       id1 := u.Input().Get("task_id");
       id2 := u.Input().Get("soldier_id");
       long1 := u.Input().Get("longitude");
       lati := u.Input().Get("latitude");
       _id1,err1:= strconv.Atoi(id1);
       _id2,err2:= strconv.Atoi(id2); 
       _long,err3:= strconv.ParseFloat(long1, 64)
       _lati,err4:= strconv.ParseFloat(lati, 64)
       if (err1 != nil || err2 != nil || err3 != nil || err4 != nil){
           u.Data["json"] = map[string]interface{}{"code":400, "enmsg":"no", "cnmsg":"失败"};
           u.ServeJSON()
           return
       }
       models.UpdateLocation(_id1, _id2, _long, _lati);
       u.Data["json"] = map[string]interface{}{"code":200, "enmsg":"ok", "cnmsg":"成功"};
       u.ServeJSON()
}