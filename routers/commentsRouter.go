package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["account_service/controllers:Account_anomaliesController"] = append(beego.GlobalControllerRouter["account_service/controllers:Account_anomaliesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["account_service/controllers:Account_anomaliesController"] = append(beego.GlobalControllerRouter["account_service/controllers:Account_anomaliesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["account_service/controllers:Account_anomaliesController"] = append(beego.GlobalControllerRouter["account_service/controllers:Account_anomaliesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["account_service/controllers:Account_anomaliesController"] = append(beego.GlobalControllerRouter["account_service/controllers:Account_anomaliesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["account_service/controllers:Account_anomaliesController"] = append(beego.GlobalControllerRouter["account_service/controllers:Account_anomaliesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["account_service/controllers:AccountsController"] = append(beego.GlobalControllerRouter["account_service/controllers:AccountsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["account_service/controllers:AccountsController"] = append(beego.GlobalControllerRouter["account_service/controllers:AccountsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["account_service/controllers:AccountsController"] = append(beego.GlobalControllerRouter["account_service/controllers:AccountsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["account_service/controllers:AccountsController"] = append(beego.GlobalControllerRouter["account_service/controllers:AccountsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["account_service/controllers:AccountsController"] = append(beego.GlobalControllerRouter["account_service/controllers:AccountsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["account_service/controllers:ApplicationPropertiesController"] = append(beego.GlobalControllerRouter["account_service/controllers:ApplicationPropertiesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["account_service/controllers:ApplicationPropertiesController"] = append(beego.GlobalControllerRouter["account_service/controllers:ApplicationPropertiesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["account_service/controllers:ApplicationPropertiesController"] = append(beego.GlobalControllerRouter["account_service/controllers:ApplicationPropertiesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["account_service/controllers:ApplicationPropertiesController"] = append(beego.GlobalControllerRouter["account_service/controllers:ApplicationPropertiesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["account_service/controllers:ApplicationPropertiesController"] = append(beego.GlobalControllerRouter["account_service/controllers:ApplicationPropertiesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["account_service/controllers:Customer_account_anomaliesController"] = append(beego.GlobalControllerRouter["account_service/controllers:Customer_account_anomaliesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["account_service/controllers:Customer_account_anomaliesController"] = append(beego.GlobalControllerRouter["account_service/controllers:Customer_account_anomaliesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["account_service/controllers:Customer_account_anomaliesController"] = append(beego.GlobalControllerRouter["account_service/controllers:Customer_account_anomaliesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["account_service/controllers:Customer_account_anomaliesController"] = append(beego.GlobalControllerRouter["account_service/controllers:Customer_account_anomaliesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["account_service/controllers:Customer_account_anomaliesController"] = append(beego.GlobalControllerRouter["account_service/controllers:Customer_account_anomaliesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["account_service/controllers:Customer_accountsController"] = append(beego.GlobalControllerRouter["account_service/controllers:Customer_accountsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["account_service/controllers:Customer_accountsController"] = append(beego.GlobalControllerRouter["account_service/controllers:Customer_accountsController"],
        beego.ControllerComments{
            Method: "AccountHistory",
            Router: `/account-history/:accountNumber`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["account_service/controllers:Customer_accountsController"] = append(beego.GlobalControllerRouter["account_service/controllers:Customer_accountsController"],
        beego.ControllerComments{
            Method: "GetAccountByAccountNumber",
            Router: `/account/:accountNumber`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["account_service/controllers:Customer_accountsController"] = append(beego.GlobalControllerRouter["account_service/controllers:Customer_accountsController"],
        beego.ControllerComments{
            Method: "AddCustomerAccount",
            Router: `/add-account`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["account_service/controllers:Customer_accountsController"] = append(beego.GlobalControllerRouter["account_service/controllers:Customer_accountsController"],
        beego.ControllerComments{
            Method: "CreditAccount",
            Router: `/credit-account/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["account_service/controllers:Customer_accountsController"] = append(beego.GlobalControllerRouter["account_service/controllers:Customer_accountsController"],
        beego.ControllerComments{
            Method: "GetAccountsByCustomerId",
            Router: `/customer/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["account_service/controllers:Customer_accountsController"] = append(beego.GlobalControllerRouter["account_service/controllers:Customer_accountsController"],
        beego.ControllerComments{
            Method: "DebitAccount",
            Router: `/debit-account/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
