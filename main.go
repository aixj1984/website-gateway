package main

import (
	"website-gateway/server"

	"github.com/mkideal/cli"
)

type cliArgs struct {
	cli.Helper
	ConnectionStr string `cli:"*c,*conn" usage:"mysql connection str" dft:"$GATEWAY_CONN_STR"`
	ListenAddr    string `cli:"*l,*listen" usage:"gateway listen host and port" dft:"$GATEWAY_LS"`
	ResourceURL   string `cli:"*r,*resource" usage:"gateway resource url" dft:"$GATEWAY_RESOURCE_URL"`
}

func main() {
	cli.Run(new(cliArgs), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*cliArgs)
		server.NewGatewayServer(argv.ConnectionStr, argv.ResourceURL).Start(argv.ListenAddr)
		return nil
	})
}

/*
package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func (t *Template) AddTemplate() {

}

func main() {
	e := echo.New()

	t := &Template{
		templates: template.ParseFiles(template.ParseGlob("views/*")),
	}

	e.Renderer = t

	e.GET("/hello", Hello)

	e.Static("/static", "static")
	e.GET("/", func(c echo.Context) error {

		return c.File("views/index.html")
		//return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/test", Hello)

	e.File("/abc", "views/index.html")

	e.Logger.Fatal(e.Start(":1323"))
}

func Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "hello", nil)
}

*/
