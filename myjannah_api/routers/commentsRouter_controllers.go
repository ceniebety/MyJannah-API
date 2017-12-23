package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["myjannah_api/controllers:CategoryFacilityController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:CategoryFacilityController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:CategoryFacilityController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:CategoryFacilityController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:CategoryFacilityController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:CategoryFacilityController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:CategoryFacilityController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:CategoryFacilityController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:CategoryFacilityController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:CategoryFacilityController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:FasilitasController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:FasilitasController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:FasilitasController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:FasilitasController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:FasilitasController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:FasilitasController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:FasilitasController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:FasilitasController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:FasilitasController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:FasilitasController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:JamaahController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:JamaahController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:JamaahController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:JamaahController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:JamaahController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:JamaahController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:JamaahController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:JamaahController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:JamaahController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:JamaahController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:LoginController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:ManasikScheduleController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:ManasikScheduleController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:ManasikScheduleController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:ManasikScheduleController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:ManasikScheduleController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:ManasikScheduleController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:ManasikScheduleController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:ManasikScheduleController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:ManasikScheduleController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:ManasikScheduleController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:MenuController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:MenuController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:MenuController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:MenuController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:MenuController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:MenuController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:MenuController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:MenuController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:MenuController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:MenuController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:PaketTravelController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:PaketTravelController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:PaketTravelController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:PaketTravelController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:PaketTravelController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:PaketTravelController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:PaketTravelController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:PaketTravelController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:PaketTravelController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:PaketTravelController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:RegistrationController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:RegistrationController"],
		beego.ControllerComments{
			Method: "RegisterTravel",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:RoleController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:RoleController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:RoleController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:RoleController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:RoleController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:RoleController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:RoleController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:RoleController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:RoleController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:RoleController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:ScheduleController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:ScheduleController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:ScheduleController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:ScheduleController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:ScheduleController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:ScheduleController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:ScheduleController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:ScheduleController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:ScheduleController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:ScheduleController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:SubPaketTravelController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:SubPaketTravelController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:SubPaketTravelController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:SubPaketTravelController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:SubPaketTravelController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:SubPaketTravelController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:SubPaketTravelController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:SubPaketTravelController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:SubPaketTravelController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:SubPaketTravelController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:TravelagentController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:TravelagentController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:TravelagentController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:TravelagentController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:TravelagentController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:TravelagentController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:TravelagentController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:TravelagentController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myjannah_api/controllers:TravelagentController"] = append(beego.GlobalControllerRouter["myjannah_api/controllers:TravelagentController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
