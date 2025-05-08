package dao

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"model"
)

// TestChatPrivadoDAO_Guardar verifica que el método Guardar funcione correctamente
func TestChatPrivadoDAO_Guardar(t *testing.T) {
	// Configuración de la prueba
	dbPool := setupTestDB(t)
	defer cleanupTestDB(t, dbPool)

	// Crear el DAO
	chatDAO := NuevoChatPrivadoDAO(dbPool)

	// Crear un nuevo chat privado para la prueba
	chatID := uuid.New()
	chatPrivado, err := model.NewChatPrivado(chatID)
	require.NoError(t, err)

	// Ejecutar la operación que queremos probar
	err = chatDAO.Guardar(chatPrivado)

	// Verificar que no haya errores
	assert.NoError(t, err)

	// Verificar que el chat se haya guardado correctamente
	guardado, err := chatDAO.BuscarPorID(chatID)
	require.NoError(t, err)
	assert.NotNil(t, guardado)
	assert.Equal(t, chatID, guardado.ID())
}

// TestChatPrivadoDAO_BuscarPorID verifica que el método BuscarPorID funcione correctamente
func TestChatPrivadoDAO_BuscarPorID(t *testing.T) {
	// Configuración de la prueba
	dbPool := setupTestDB(t)
	defer cleanupTestDB(t, dbPool)

	// Crear el DAO
	chatDAO := NuevoChatPrivadoDAO(dbPool)

	// Crear un nuevo chat privado para la prueba
	chatID := uuid.New()
	chatPrivado, err := model.NewChatPrivado(chatID)
	require.NoError(t, err)

	// Guardarlo primero en la base de datos
	err = chatDAO.Guardar(chatPrivado)
	require.NoError(t, err)

	// Ejecutar la operación que queremos probar
	encontrado, err := chatDAO.BuscarPorID(chatID)

	// Verificar que no haya errores
	assert.NoError(t, err)

	// Verificar que el chat encontrado sea el correcto
	assert.NotNil(t, encontrado)
	assert.Equal(t, chatID, encontrado.ID())

	// Buscar un chat que no existe
	idInexistente := uuid.New()
	noEncontrado, err := chatDAO.BuscarPorID(idInexistente)

	// Verificar que no haya errores pero el resultado sea nil
	assert.NoError(t, err)
	assert.Nil(t, noEncontrado)
}

// TestChatPrivadoDAO_BuscarTodos verifica que el método BuscarTodos funcione correctamente
func TestChatPrivadoDAO_BuscarTodos(t *testing.T) {
	// Configuración de la prueba
	dbPool := setupTestDB(t)
	defer cleanupTestDB(t, dbPool)

	// Crear el DAO
	chatDAO := NuevoChatPrivadoDAO(dbPool)

	// Crear varios chats privados para la prueba
	chat1, err := model.NewChatPrivado(uuid.New())
	require.NoError(t, err)
	
	chat2, err := model.NewChatPrivado(uuid.New())
	require.NoError(t, err)
	
	chat3, err := model.NewChatPrivado(uuid.New())
	require.NoError(t, err)

	// Guardarlos en la base de datos
	err = chatDAO.Guardar(chat1)
	require.NoError(t, err)
	
	err = chatDAO.Guardar(chat2)
	require.NoError(t, err)
	
	err = chatDAO.Guardar(chat3)
	require.NoError(t, err)

	// Ejecutar la operación que queremos probar
	chats, err := chatDAO.BuscarTodos()

	// Verificar que no haya errores
	assert.NoError(t, err)

	// Verificar que se hayan encontrado al menos los tres chats creados
	assert.GreaterOrEqual(t, len(chats), 3)

	// Verificar que los chats creados estén en la lista
	encontrados := 0
	for _, c := range chats {
		if c.ID() == chat1.ID() || c.ID() == chat2.ID() || c.ID() == chat3.ID() {
			encontrados++
		}
	}

	assert.Equal(t, 3, encontrados)
}

// TestChatPrivadoDAO_Eliminar verifica que el método Eliminar funcione correctamente
func TestChatPrivadoDAO_Eliminar(t *testing.T) {
	// Configuración de la prueba
	dbPool := setupTestDB(t)
	defer cleanupTestDB(t, dbPool)

	// Crear el DAO
	chatDAO := NuevoChatPrivadoDAO(dbPool)

	// Crear un nuevo chat privado para la prueba
	chatID := uuid.New()
	chatPrivado, err := model.NewChatPrivado(chatID)
	require.NoError(t, err)

	// Guardarlo primero en la base de datos
	err = chatDAO.Guardar(chatPrivado)
	require.NoError(t, err)

	// Verificar que existe antes de eliminar
	encontrado, err := chatDAO.BuscarPorID(chatID)
	require.NoError(t, err)
	assert.NotNil(t, encontrado)

	// Ejecutar la operación que queremos probar
	err = chatDAO.Eliminar(chatID)

	// Verificar que no haya errores
	assert.NoError(t, err)

	// Verificar que el chat ya no existe
	eliminado, err := chatDAO.BuscarPorID(chatID)
	assert.NoError(t, err)
	assert.Nil(t, eliminado)
}
