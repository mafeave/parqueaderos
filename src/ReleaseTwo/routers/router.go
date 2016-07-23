// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"ReleaseTwo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/tipo_propietario",
			beego.NSInclude(
				&controllers.TipoPropietarioController{},
			),
		),

		beego.NSNamespace("/propietario",
			beego.NSInclude(
				&controllers.PropietarioController{},
			),
		),

		beego.NSNamespace("/vehiculo",
			beego.NSInclude(
				&controllers.VehiculoController{},
			),
		),

		beego.NSNamespace("/grupo_isla",
			beego.NSInclude(
				&controllers.GrupoIslaController{},
			),
		),

		beego.NSNamespace("/isla",
			beego.NSInclude(
				&controllers.IslaController{},
			),
		),
	)
	beego.AddNamespace(ns)
}