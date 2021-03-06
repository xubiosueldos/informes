package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xubiosueldos/informes/cargassocialesf931"
	"github.com/xubiosueldos/informes/librosueldos"
	"github.com/xubiosueldos/informes/librosueldosdigital"
	"github.com/xubiosueldos/informes/liquidacionfinalanual"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

type Routes []Route

func newRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandleFunc)

	}

	return router
}

var routes = Routes{
	Route{
		"Healthy",
		"GET",
		"/api/informe/healthy",
		Healthy,
	},
	Route{
		"InformeF931",
		"GET",
		"/api/informe/informes/cargas-sociales-f931",
		cargassocialesf931.InformeF931,
	},
	Route{
		"LibroSueldos",
		"GET",
		"/api/informe/informes/libro-sueldos",
		librosueldos.LibroSueldos,
	},
	Route{
		"InformeF931ExportarTxt",
		"GET",
		"/api/informe/informes/cargas-sociales-f931-exportartxt",
		cargassocialesf931.InformeF931ExportarTxt,
	},
	Route{
		"ImpresionEncabezado",
		"GET",
		"/api/informe/informes/libro-sueldos/impresion-encabezado",
		librosueldos.ImpresionEncabezado,
	},
	Route{
		"ImpresionLiquidaciones",
		"GET",
		"/api/informe/informes/libro-sueldos/impresion-liquidaciones",
		librosueldos.ImpresionLiquidaciones,
	},
	Route{
		"LiquidacionFinalAnualF1357",
		"GET",
		"/api/informe/informes/liquidacion-final-anual-f1357",
		liquidacionfinalanual.LiquidacionFinalAnualF1357,
	},
	Route{
		"LiquidacionFinalAnualF1357ExportarTxt",
		"GET",
		"/api/informe/informes/liquidacion-final-anual-f1357-exportartxt",
		liquidacionfinalanual.LiquidacionFinalAnualF1357ExportarTxt,
	},
	Route{
		"LibroSueldosDigital",
		"GET",
		"/api/informe/informes/afip-libro-sueldos-digital",
		librosueldosdigital.LibroSueldosDigital,
	},
	Route{
		"LibroSueldosDigitalExportarTxtConceptosAFIP",
		"GET",
		"/api/informe/informes/afip-libro-sueldos-digital-exportartxt-conceptosafip",
		librosueldosdigital.LibroSueldosDigitalExportarTxtConceptosAFIP,
	},
	Route{
		"LibroSueldosDigitalExportarTxtLiquidacionesPeriodo",
		"GET",
		"/api/informe/informes/afip-libro-sueldos-digital-exportartxt-liquidacionesperiodo",
		librosueldosdigital.LibroSueldosDigitalExportarTxtLiquidacionesPeriodo,
	},
}
