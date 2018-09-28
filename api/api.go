package api

import (
	"github.com/bojand/ghz-web/config"
	"github.com/bojand/ghz-web/service"
	"github.com/labstack/echo"
)

// @title ghz-web API
// @version 1.0
// @description ghz-web REST API

// Setup sets up the application API
func Setup(
	config *config.Config,
	info *config.Info,
	g *echo.Group,
	ps service.ProjectService,
	ts service.TestService,
	rs service.RunService,
	ds service.DetailService) {

	SetupInfoAPI(info, g)

	projectGroup := g.Group("/projects")
	SetupProjectAPI(projectGroup, ps)

	testsGroup := projectGroup.Group("/:pid/tests")
	SetupTestAPI(testsGroup, ts, rs)

	runsGroup := testsGroup.Group("/:tid/runs")
	SetupRunAPI(runsGroup, rs, ds)

	detailGroup := runsGroup.Group("/:rid/details")
	SetupDetailAPI(detailGroup, ds)

	SetupRawAPI(g, ps, ts, rs, ds)
}
