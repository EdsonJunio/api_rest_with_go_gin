package routes

import (
	"github.com/EdsonJunio/api_rest_with_go_gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleResquests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibirTodosOsAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriarNovoAluno)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	r.DELETE("aluno/:id", controllers.DeletaAluno)
	r.PATCH("/alunos/:id", controllers.EditarAluno)
	r.GET("/Alunos/:cpf", controllers.BuscarAlunoCPF)
	r.Run()
}
