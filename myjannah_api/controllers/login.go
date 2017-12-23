package controllers

import (
  "myjannah_api/models"
  "myjannah_api/request"
  "myjannah_api/response"
  "github.com/astaxie/beego/validation"
  "github.com/astaxie/beego"
  "myjannah_api/utils"
  "time"
  "encoding/json"
)

// LoginController operations for Login
type LoginController struct {
	beego.Controller
}

type LoginToken struct {
	Username  string      `json:"username"`
	Token string          `json:"token"`
}

// URLMapping ...
func (c *LoginController) URLMapping() {
	c.Mapping("Post", c.Login)
}

// Login ...
// @Username
// @Password
// @TypeLogin
// @Param	body		body 	models.Login	true		"body for Login content"
// @Success 201 {int} models.Travelagent
// @Failure 403 body is empty
// @router / login [post]
func(l *LoginController) Login(){

  var loginReq request.LoginRequest

  if err := json.Unmarshal(l.Ctx.Input.RequestBody, &loginReq); err != nil {
     l.Ctx.ResponseWriter.WriteHeader(410)
      l.Data["json"] = ErrResponse{410000, "Invalid Request"}
      l.ServeJSON()
      return
  }

     valid := validation.Validation{}

     //validation
       valid.Required(loginReq.Username, "username").Message("username must be insert")
       valid.Required(loginReq.Password, "password").Message("password must be insert")
       valid.Required(loginReq.TypeLogin, "typeLogin").Message("Type Login must be insert")

   //If Has Error From validation
       if valid.HasErrors() {
         for _, err := range valid.Errors {
             l.Ctx.ResponseWriter.WriteHeader(403)
             l.Data["json"] = ErrResponse{403001, map[string]string{err.Key: err.Message}}
             l.ServeJSON()
             return
            }
       }else{
         if loginReq.TypeLogin == "TRAVELAGENT" {
           //Check Available Account
            w, err := models.CheckUserByEmailAndPassword(loginReq.Username, loginReq.Password);
            if err != nil {
              l.Data["json"] = ErrResponse{428001, "Account Not Be Found"}
              l.ServeJSON()
              return
            }else{
              //Check Status Account
                if w.Status != "ACTIVE" {
                  l.Data["json"] = ErrResponse{429001, "Your Account Not Be Activated"}
                  l.ServeJSON()
                  return
                } else{
                  //Create Token
                    et := utils.EasyToken{
                    Username: loginReq.Username,
                    Expires:  time.Now().Unix() + 3600,
                  }
                  token, err := et.GetToken()
                  if token == "" || err != nil {
                    l.Data["json"] = ErrResponse{-0, err.Error()}
                  } else {
                    l.Data["json"] = Response{0, "success.", response.LoginResponse{token}}
                  }
                  l.ServeJSON()
                }
            }
        }
      }
  }



  //validation
    // valid.Required(username, "username").Message("username must be insert")
    // valid.Required(password, "password").Message("password must be insert")
    // valid.Required(typeLogin, "typeLogin").Message("Type Login must be insert")

  //If Has Error From validation
  // if valid.HasErrors() {
  //     for _, err := range valid.Errors {
  //       l.Ctx.ResponseWriter.WriteHeader(403)
  //       l.Data["json"] = ErrResponse{403001, map[string]string{err.Key: err.Message}}
  //       l.ServeJSON()
  //       return
  //     }
  // }else{
  //   if typeLogin == "TRAVELAGENT" {

      //Check Available Account
      // w, err := models.CheckUserByEmailAndPassword(username, password);
      // if err != nil {
      //   l.Data["json"] = ErrResponse{428001, "Account Not Be Found"}
      //   l.ServeJSON()
      //   return
      // }

      //Check Status Account
    //   if w.Status != "ACTIVE" {
    //     l.Data["json"] = ErrResponse{429001, "Your Account Not Be Activated"}
    //     l.ServeJSON()
    //     return
    //   }
    // }

    //Create Token
    // et := utils.EasyToken{
    //     Username: username,
    //     Expires:  time.Now().Unix() + 3600,
    // }
    // token, err := et.GetToken()
    // if token == "" || err != nil {
    //     l.Data["json"] = ErrResponse{-0, err.Error()}
    // } else {
    //     l.Data["json"] = Response{0, "success.", LoginToken{username, token}}
    // }
    //   l.ServeJSON()
    // }
// }
