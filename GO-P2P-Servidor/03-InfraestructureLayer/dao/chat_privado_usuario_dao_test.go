package dao

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"model"
)

// TestChatPrivadoUsuarioDAO_Guardar verifica que el método Guardar funcione correctamente
func TestChatPrivadoUsuarioDAO_Guardar(t *testing.T) {
	// Configuración de la prueba
	dbPool := setupTestDB(t)
	defer cleanupTestDB(t, dbPool)

	// Crear el DAO
	chatDAO := NuevoChatPrivadoDAO(dbPool)
	chatUsuarioDAO := NuevoChatPrivadoUsuarioDAO(dbPool)

	// Crear un chat privado para la prueba
	chatID := uuid.New()
	chatPrivado, err := model.NewChatPrivado(chatID)
	require.NoError(t, err)

	// Guardar el chat privado
	err = chatDAO.Guardar(chatPrivado)
	require.NoError(t, err)

	// Crear un ChatPrivadoUsuario
	usuarioID := uuid.New()
	chatUsuario, err := model.NewChatPrivadoUsuario(chatID, usuarioID)
	require.NoError(t, err)

	// Ejecutar la operación que queremos probar
	err = chatUsuarioDAO.Guardar(chatUsuario)
	
	// Verificar que no haya errores
	assert.NoError(t, err)

	// Verificar que se haya guardado correctamente
	guardado, err := chatUsuarioDAO.BuscarPorIDs(chatID, usuarioID)
	require.NoError(t, err)
	assert.NotNil(t, guardado)
	assert.Equal(t, chatID, guardado.ChatPrivadoID())
	assert.Equal(t, usuarioID, guardado.UsuarioID())
}

// TestChatPrivadoUsuarioDAO_BuscarPorIDs verifica que el método BuscarPorIDs funcione correctamente
func TestChatPrivadoUsuarioDAO_BuscarPorIDs(t *testing.T) {
	// Configuración de la prueba
	dbPool := setupTestDB(t)
	defer cleanupTestDB(t, dbPool)

	// Crear el DAO
	chatDAO := NuevoChatPrivadoDAO(dbPool)
	chatUsuarioDAO := NuevoChatPrivadoUsuarioDAO(dbPool)

	// Crear un chat privado para la prueba
	chatID := uuid.New()
	chatPrivado, err := model.NewChatPrivado(chatID)
	require.NoError(t, err)

	// Guardar el chat privado
	err = chatDAO.Guardar(chatPrivado)
	require.NoError(t, err)

	// Crear y guardar un ChatPrivadoUsuario
	usuarioID := uuid.New()
	chatUsuario, err := model.NewChatPrivadoUsuario(chatID, usuarioID)
	require.NoError(t, err)
	err = chatUsuarioDAO.Guardar(chatUsuario)
	require.NoError(t, err)

	// Ejecutar la operación que queremos probar
	encontrado, err := chatUsuarioDAO.BuscarPorIDs(chatID, usuarioID)

	// Verificar que no haya errores
	assert.NoError(t, err)

	// Verificar que el chat usuario encontrado sea el correcto
	assert.NotNil(t, encontrado)
	assert.Equal(t, chatID, encontrado.ChatPrivadoID())
	assert.Equal(t, usuarioID, encontrado.UsuarioID())

	// Buscar un chat usuario que no existe
	idInexistente := uuid.New()
	noEncontrado, err := chatUsuarioDAO.BuscarPorIDs(chatID, idInexistente)

	// Verificar que no haya errores pero el resultado sea nil
	assert.NoError(t, err)
	assert.Nil(t, noEncontrado)
}

// TestChatPrivadoUsuarioDAO_BuscarPorChatPrivadoID verifica que el método BuscarPorChatPrivadoID funcione correctamente
func TestChatPrivadoUsuarioDAO_BuscarPorChatPrivadoID(t *testing.T) {
	// Configuración de la prueba
	dbPool := setupTestDB(t)
	defer cleanupTestDB(t, dbPool)

	// Crear el DAO
	chatDAO := NuevoChatPrivadoDAO(dbPool)
	chatUsuarioDAO := NuevoChatPrivadoUsuarioDAO(dbPool)

	// Crear un chat privado para la prueba
	chatID := uuid.New()
	chatPrivado, err := model.NewChatPrivado(chatID)
	require.NoError(t, err)

	// Guardar el chat privado
	err = chatDAO.Guardar(chatPrivado)
	require.NoError(t, err)

	// Crear y guardar varios usuarios para el chat
	usuarioID1 := uuid.New()
	usuarioID2 := uuid.New()
	usuarioID3 := uuid.New()

	chatUsuario1, err := model.NewChatPrivadoUsuario(chatID, usuarioID1)
	require.NoError(t, err)

	chatUsuario2, err := model.NewChatPrivadoUsuario(chatID, usuarioID2)
	require.NoError(t, err)

	chatUsuario3, err := model.NewChatPrivadoUsuario(chatID, usuarioID3)
	require.NoError(t, err)

	// Guardar los usuarios del chat
	err = chatUsuarioDAO.Guardar(chatUsuario1)
	require.NoError(t, err)

	err = chatUsuarioDAO.Guardar(chatUsuario2)
	require.NoError(t, err)

	err = chatUsuarioDAO.Guardar(chatUsuario3)
	require.NoError(t, err)

	// Ejecutar la operación que queremos probar
	usuarios, err := chatUsuarioDAO.BuscarPorChatPrivadoID(chatID)

	// Verificar que no haya errores
	assert.NoError(t, err)

	// Verificar que se hayan encontrado los tres usuarios
	assert.Equal(t, 3, len(usuarios))

	// Verificar que cada usuario pertenezca al chat correcto
	for _, u := range usuarios {
		assert.Equal(t, chatID, u.ChatPrivadoID())
	}

	// Verificar que estén todos los usuarios
	usuariosIDs := make(map[uuid.UUID]bool)
	for _, u := range usuarios {
		usuariosIDs[u.UsuarioID()] = true
	}

	assert.True(t, usuariosIDs[usuarioID1])
	assert.True(t, usuariosIDs[usuarioID2])
	assert.True(t, usuariosIDs[usuarioID3])
}

// TestChatPrivadoUsuarioDAO_BuscarPorUsuarioID verifica que el método BuscarPorUsuarioID funcione correctamente
func TestChatPrivadoUsuarioDAO_BuscarPorUsuarioID(t *testing.T) {
	// Configuración de la prueba
	dbPool := setupTestDB(t)
	defer cleanupTestDB(t, dbPool)

	// Crear el DAO
	chatDAO := NuevoChatPrivadoDAO(dbPool)
	chatUsuarioDAO := NuevoChatPrivadoUsuarioDAO(dbPool)

	// Crear varios chats privados para la prueba
	chatID1 := uuid.New()
	chatID2 := uuid.New()
	chatID3 := uuid.New()

	chatPrivado1, err := model.NewChatPrivado(chatID1)
	require.NoError(t, err)

	chatPrivado2, err := model.NewChatPrivado(chatID2)
	require.NoError(t, err)

	chatPrivado3, err := model.NewChatPrivado(chatID3)
	require.NoError(t, err)

	// Guardar los chats privados
	err = chatDAO.Guardar(chatPrivado1)
	require.NoError(t, err)

	err = chatDAO.Guardar(chatPrivado2)
	require.NoError(t, err)

	err = chatDAO.Guardar(chatPrivado3)
	require.NoError(t, err)

	// Crear un usuario para la prueba
	usuarioID := uuid.New()

	// Asociar el usuario a los chats
	chatUsuario1, err := model.NewChatPrivadoUsuario(chatID1, usuarioID)
	require.NoError(t, err)

	chatUsuario2, err := model.NewChatPrivadoUsuario(chatID2, usuarioID)
	require.NoError(t, err)

	chatUsuario3, err := model.NewChatPrivadoUsuario(chatID3, usuarioID)
	require.NoError(t, err)

	// Guardar las asociaciones
	err = chatUsuarioDAO.Guardar(chatUsuario1)
	require.NoError(t, err)

	err = chatUsuarioDAO.Guardar(chatUsuario2)
	require.NoError(t, err)

	err = chatUsuarioDAO.Guardar(chatUsuario3)
	require.NoError(t, err)

	// Ejecutar la operación que queremos probar
	chats, err := chatUsuarioDAO.BuscarPorUsuarioID(usuarioID)

	// Verificar que no haya errores
	assert.NoError(t, err)

	// Verificar que se hayan encontrado los tres chats
	assert.Equal(t, 3, len(chats))

	// Verificar que todos pertenezcan al mismo usuario
	for _, c := range chats {
		assert.Equal(t, usuarioID, c.UsuarioID())
	}

	// Verificar que estén todos los chats
	chatsIDs := make(map[uuid.UUID]bool)
	for _, c := range chats {
		chatsIDs[c.ChatPrivadoID()] = true
	}

	assert.True(t, chatsIDs[chatID1])
	assert.True(t, chatsIDs[chatID2])
	assert.True(t, chatsIDs[chatID3])
}

// TestChatPrivadoUsuarioDAO_BuscarChatEntreUsuarios verifica que el método BuscarChatEntreUsuarios funcione correctamente
func TestChatPrivadoUsuarioDAO_BuscarChatEntreUsuarios(t *testing.T) {
	// Configuración de la prueba
	dbPool := setupTestDB(t)
	defer cleanupTestDB(t, dbPool)

	// Crear el DAO
	chatDAO := NuevoChatPrivadoDAO(dbPool)
	chatUsuarioDAO := NuevoChatPrivadoUsuarioDAO(dbPool)

	// Crear un chat privado para la prueba
	chatID := uuid.New()
	chatPrivado, err := model.NewChatPrivado(chatID)
	require.NoError(t, err)

	// Guardar el chat privado
	err = chatDAO.Guardar(chatPrivado)
	require.NoError(t, err)

	// Crear dos usuarios para la prueba
	usuarioID1 := uuid.New()
	usuarioID2 := uuid.New()

	// Asociar los usuarios al chat
	chatUsuario1, err := model.NewChatPrivadoUsuario(chatID, usuarioID1)
	require.NoError(t, err)

	chatUsuario2, err := model.NewChatPrivadoUsuario(chatID, usuarioID2)
	require.NoError(t, err)

	// Guardar las asociaciones
	err = chatUsuarioDAO.Guardar(chatUsuario1)
	require.NoError(t, err)

	err = chatUsuarioDAO.Guardar(chatUsuario2)
	require.NoError(t, err)

	// Ejecutar la operación que queremos probar
	chatEncontrado, err := chatUsuarioDAO.BuscarChatEntreUsuarios(usuarioID1, usuarioID2)

	// Verificar que no haya errores
	assert.NoError(t, err)

	// Verificar que se haya encontrado el chat
	assert.NotNil(t, chatEncontrado)
	assert.Equal(t, chatID, chatEncontrado.ID())

	// Probar con usuarios que no tienen un chat en común
	usuarioSinChat := uuid.New()
	chatNoEncontrado, err := chatUsuarioDAO.BuscarChatEntreUsuarios(usuarioID1, usuarioSinChat)

	// Verificar que no haya errores pero no se encontró chat
	assert.NoError(t, err)
	assert.Nil(t, chatNoEncontrado)
}

// TestChatPrivadoUsuarioDAO_Eliminar verifica que el método Eliminar funcione correctamente
func TestChatPrivadoUsuarioDAO_Eliminar(t *testing.T) {
	// Configuración de la prueba
	dbPool := setupTestDB(t)
	defer cleanupTestDB(t, dbPool)

	// Crear el DAO
	chatDAO := NuevoChatPrivadoDAO(dbPool)
	chatUsuarioDAO := NuevoChatPrivadoUsuarioDAO(dbPool)

	// Crear un chat privado para la prueba
	chatID := uuid.New()
	chatPrivado, err := model.NewChatPrivado(chatID)
	require.NoError(t, err)

	// Guardar el chat privado
	err = chatDAO.Guardar(chatPrivado)
	require.NoError(t, err)

	// Crear un usuario para la prueba
	usuarioID := uuid.New()

	// Asociar el usuario al chat
	chatUsuario, err := model.NewChatPrivadoUsuario(chatID, usuarioID)
	require.NoError(t, err)

	// Guardar la asociación
	err = chatUsuarioDAO.Guardar(chatUsuario)
	require.NoError(t, err)

	// Verificar que existe antes de eliminar
	encontrado, err := chatUsuarioDAO.BuscarPorIDs(chatID, usuarioID)
	require.NoError(t, err)
	assert.NotNil(t, encontrado)

	// Ejecutar la operación que queremos probar
	err = chatUsuarioDAO.Eliminar(chatID, usuarioID)

	// Verificar que no haya errores
	assert.NoError(t, err)

	// Verificar que ya no existe la asociación
	eliminado, err := chatUsuarioDAO.BuscarPorIDs(chatID, usuarioID)
	assert.NoError(t, err)
	assert.Nil(t, eliminado)
}

// TestChatPrivadoUsuarioDAO_EliminarPorChatPrivadoID verifica que el método EliminarPorChatPrivadoID funcione correctamente
func TestChatPrivadoUsuarioDAO_EliminarPorChatPrivadoID(t *testing.T) {
	// Configuración de la prueba
	dbPool := setupTestDB(t)
	defer cleanupTestDB(t, dbPool)

	// Crear el DAO
	chatDAO := NuevoChatPrivadoDAO(dbPool)
	chatUsuarioDAO := NuevoChatPrivadoUsuarioDAO(dbPool)

	// Crear un chat privado para la prueba
	chatID := uuid.New()
	chatPrivado, err := model.NewChatPrivado(chatID)
	require.NoError(t, err)

	// Guardar el chat privado
	err = chatDAO.Guardar(chatPrivado)
	require.NoError(t, err)

	// Crear varios usuarios para el chat
	usuarioID1 := uuid.New()
	usuarioID2 := uuid.New()
	usuarioID3 := uuid.New()

	// Asociar los usuarios al chat
	chatUsuario1, err := model.NewChatPrivadoUsuario(chatID, usuarioID1)
	require.NoError(t, err)

	chatUsuario2, err := model.NewChatPrivadoUsuario(chatID, usuarioID2)
	require.NoError(t, err)

	chatUsuario3, err := model.NewChatPrivadoUsuario(chatID, usuarioID3)
	require.NoError(t, err)

	// Guardar las asociaciones
	err = chatUsuarioDAO.Guardar(chatUsuario1)
	require.NoError(t, err)

	err = chatUsuarioDAO.Guardar(chatUsuario2)
	require.NoError(t, err)

	err = chatUsuarioDAO.Guardar(chatUsuario3)
	require.NoError(t, err)

	// Verificar que existen antes de eliminar
	usuarios, err := chatUsuarioDAO.BuscarPorChatPrivadoID(chatID)
	require.NoError(t, err)
	assert.Equal(t, 3, len(usuarios))

	// Ejecutar la operación que queremos probar
	err = chatUsuarioDAO.EliminarPorChatPrivadoID(chatID)

	// Verificar que no haya errores
	assert.NoError(t, err)

	// Verificar que ya no existen las asociaciones
	usuariosEliminados, err := chatUsuarioDAO.BuscarPorChatPrivadoID(chatID)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(usuariosEliminados))
}
