package router

import "github.com/gin-gonic/gin"

/***
O gin aqui está funcionando como um framework para criacão de rota http.
Um comparativo seria algo como fastfy ou express pro javascript
***/

func Initialize() {
	//Inicializa o Router utilizando as configs Default do gin
	router := gin.Default()

	//Initialize Routes
	initialize_routes(router)

	//para mudar a porta padrao é necessário utilizar o Run
	router.Run(":3333")
}
