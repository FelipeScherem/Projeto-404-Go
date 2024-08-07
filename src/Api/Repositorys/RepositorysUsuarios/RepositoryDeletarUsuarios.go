package repositoryUsuarios

import (
	db "projeto404/src/Api/Database"
	modelUsuario "projeto404/src/Api/Models/ModelUsers"
)

func DeletarUsuario(usuarioStruct modelUsuario.UsuarioStruct, id int) (string, error) {
	database := db.ConectaDB() // Abre a conexão com o banco de dados
	defer db.FechaDB(database) // Fecha conexão com o banco de dados no final da função

	if err := database.
		Delete(&usuarioStruct, id).
		Error; err != nil {
		return "Erro ao excluir o usuário:", err
	} else {
		return "Usuario excluido com sucesso", nil
	}
}
