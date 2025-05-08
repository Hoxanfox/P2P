package dao

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"model"
)

// TestNotificacionDAO_Guardar verifica que el método Guardar funcione correctamente
func TestNotificacionDAO_Guardar(t *testing.T) {
	// Inicializar conexión a la base de datos
	dbPool := setupTestDB(t)
	defer cleanupTestDB(t, dbPool)

	// Crear instancia del DAO
	notificacionDAO := NuevoNotificacionDAO(dbPool)

	// Crear notificación de prueba
	id := uuid.New()
	usuarioID := uuid.New()
	contenido := "Notificación de prueba para test"
	fecha := time.Now()
	// UUID vacío para invitacionID ya que no necesitamos asociarlo a una invitación
	invitacionID := uuid.Nil

	// Crear la notificación usando el constructor del modelo
	notificacion, err := model.NewNotificacion(id, usuarioID, contenido, fecha, invitacionID)
	if err != nil {
		t.Fatalf("Error al crear objeto notificación: %v", err)
	}

	// Guardar notificación en la base de datos
	err = notificacionDAO.Guardar(notificacion)
	if err != nil {
		t.Fatalf("Error al guardar notificación: %v", err)
	}

	// Verificar que se guardó correctamente reobteniendo por ID
	reobtenida, err := notificacionDAO.BuscarPorID(id)
	if err != nil {
		t.Fatalf("Error al obtener notificación guardada: %v", err)
	}

	if reobtenida == nil {
		t.Fatalf("No se encontró la notificación guardada con ID: %s", id)
	}

	// Verificar que los datos sean correctos
	if reobtenida.ID() != id {
		t.Errorf("ID incorrecto. Esperado: %s, Obtenido: %s", id, reobtenida.ID())
	}

	if reobtenida.UsuarioID() != usuarioID {
		t.Errorf("UsuarioID incorrecto. Esperado: %s, Obtenido: %s", usuarioID, reobtenida.UsuarioID())
	}

	if reobtenida.Contenido() != contenido {
		t.Errorf("Contenido incorrecto. Esperado: %s, Obtenido: %s", contenido, reobtenida.Contenido())
	}

	// Eliminar la notificación después de la prueba
	_, err = dbPool.ExecContext(
		context.Background(),
		"DELETE FROM notificaciones WHERE id = ?",
		notificacion.ID().String(),
	)
	if err != nil {
		t.Logf("Error al limpiar notificación de prueba: %v", err)
	}
}

// TestNotificacionDAO_BuscarPorID verifica que el método BuscarPorID funcione correctamente
func TestNotificacionDAO_BuscarPorID(t *testing.T) {
	// Inicializar conexión a la base de datos
	dbPool := setupTestDB(t)
	defer cleanupTestDB(t, dbPool)

	// Crear instancia del DAO
	notificacionDAO := NuevoNotificacionDAO(dbPool)

	// Crear notificación de prueba
	id := uuid.New()
	usuarioID := uuid.New()
	contenido := "Notificación para buscar por ID"
	fecha := time.Now()
	invitacionID := uuid.Nil

	// Crear la notificación
	notificacion, err := model.NewNotificacion(id, usuarioID, contenido, fecha, invitacionID)
	if err != nil {
		t.Fatalf("Error al crear objeto notificación: %v", err)
	}

	// Guardar notificación en la base de datos
	err = notificacionDAO.Guardar(notificacion)
	if err != nil {
		t.Fatalf("Error al guardar notificación: %v", err)
	}

	// Buscar la notificación por ID
	reobtenida, err := notificacionDAO.BuscarPorID(id)
	if err != nil {
		t.Fatalf("Error al buscar notificación por ID: %v", err)
	}

	if reobtenida == nil {
		t.Fatalf("No se encontró la notificación con ID: %s", id)
	}

	// Verificar que los datos sean correctos
	if reobtenida.ID() != id {
		t.Errorf("ID incorrecto. Esperado: %s, Obtenido: %s", id, reobtenida.ID())
	}

	if reobtenida.UsuarioID() != usuarioID {
		t.Errorf("UsuarioID incorrecto. Esperado: %s, Obtenido: %s", usuarioID, reobtenida.UsuarioID())
	}

	if reobtenida.Contenido() != contenido {
		t.Errorf("Contenido incorrecto. Esperado: %s, Obtenido: %s", contenido, reobtenida.Contenido())
	}

	// Limpiar después de la prueba
	_, err = dbPool.ExecContext(
		context.Background(),
		"DELETE FROM notificaciones WHERE id = ?",
		id.String(),
	)
	if err != nil {
		t.Logf("Error al limpiar notificación de prueba: %v", err)
	}
}

// TestNotificacionDAO_BuscarPorUsuarioID verifica que el método BuscarPorUsuarioID funcione correctamente
func TestNotificacionDAO_BuscarPorUsuarioID(t *testing.T) {
	// Inicializar conexión a la base de datos
	dbPool := setupTestDB(t)
	defer cleanupTestDB(t, dbPool)

	// Crear instancia del DAO
	notificacionDAO := NuevoNotificacionDAO(dbPool)

	// Generar un ID de usuario al azar para esta prueba
	usuarioID := uuid.New()

	// Crear varias notificaciones para el mismo usuario
	notificaciones := make([]*model.Notificacion, 0)
	for i := 0; i < 3; i++ {
		id := uuid.New()
		contenido := fmt.Sprintf("Notificación %d para usuario", i+1)
		fecha := time.Now().Add(time.Duration(i) * time.Hour) // Diferentes horas
		invitacionID := uuid.Nil

		// Crear la notificación
		notificacion, err := model.NewNotificacion(id, usuarioID, contenido, fecha, invitacionID)
		if err != nil {
			t.Fatalf("Error al crear objeto notificación %d: %v", i+1, err)
		}

		// Guardar notificación en la base de datos
		err = notificacionDAO.Guardar(notificacion)
		if err != nil {
			t.Fatalf("Error al guardar notificación %d: %v", i+1, err)
		}

		notificaciones = append(notificaciones, notificacion)
	}

	// Buscar las notificaciones por usuario ID
	resultados, err := notificacionDAO.BuscarPorUsuarioID(usuarioID)
	if err != nil {
		t.Fatalf("Error al buscar notificaciones por usuario ID: %v", err)
	}

	// Verificar que se encontraron las notificaciones
	if len(resultados) != len(notificaciones) {
		t.Errorf("Número incorrecto de notificaciones. Esperado: %d, Obtenido: %d", 
			len(notificaciones), len(resultados))
	}

	// Limpiar después de la prueba
	for _, notif := range notificaciones {
		_, err = dbPool.ExecContext(
			context.Background(),
			"DELETE FROM notificaciones WHERE id = ?",
			notif.ID().String(),
		)
		if err != nil {
			t.Logf("Error al limpiar notificación de prueba: %v", err)
		}
	}
}

// TestNotificacionDAO_ActualizarEstadoLeido verifica que el método ActualizarEstadoLeido funcione correctamente
func TestNotificacionDAO_ActualizarEstadoLeido(t *testing.T) {
	// Inicializar conexión a la base de datos
	dbPool := setupTestDB(t)
	defer cleanupTestDB(t, dbPool)

	// Crear instancia del DAO
	notificacionDAO := NuevoNotificacionDAO(dbPool)

	// Crear notificación de prueba (inicialmente no leída)
	id := uuid.New()
	usuarioID := uuid.New()
	contenido := "Notificación para actualizar estado"
	fecha := time.Now()
	invitacionID := uuid.Nil

	// Crear la notificación
	notificacion, err := model.NewNotificacion(id, usuarioID, contenido, fecha, invitacionID)
	if err != nil {
		t.Fatalf("Error al crear objeto notificación: %v", err)
	}

	// Por defecto debería estar como no leída
	if notificacion.Leido() {
		t.Fatalf("La notificación debería estar marcada como no leída por defecto")
	}

	// Guardar notificación en la base de datos
	err = notificacionDAO.Guardar(notificacion)
	if err != nil {
		t.Fatalf("Error al guardar notificación: %v", err)
	}

	// Marcar como leída
	err = notificacionDAO.ActualizarEstadoLeido(id, true)
	if err != nil {
		t.Fatalf("Error al actualizar estado leído: %v", err)
	}

	// Verificar que se actualizó correctamente
	reobtenida, err := notificacionDAO.BuscarPorID(id)
	if err != nil {
		t.Fatalf("Error al obtener notificación actualizada: %v", err)
	}

	if !reobtenida.Leido() {
		t.Errorf("La notificación debería estar marcada como leída")
	}

	// Limpiar después de la prueba
	_, err = dbPool.ExecContext(
		context.Background(),
		"DELETE FROM notificaciones WHERE id = ?",
		id.String(),
	)
	if err != nil {
		t.Logf("Error al limpiar notificación de prueba: %v", err)
	}
}

// TestNotificacionDAO_Eliminar verifica que el método Eliminar funcione correctamente
func TestNotificacionDAO_Eliminar(t *testing.T) {
	// Inicializar conexión a la base de datos
	dbPool := setupTestDB(t)
	defer cleanupTestDB(t, dbPool)

	// Crear instancia del DAO
	notificacionDAO := NuevoNotificacionDAO(dbPool)

	// Crear notificación de prueba
	id := uuid.New()
	usuarioID := uuid.New()
	contenido := "Notificación para eliminar"
	fecha := time.Now()
	invitacionID := uuid.Nil

	// Crear la notificación
	notificacion, err := model.NewNotificacion(id, usuarioID, contenido, fecha, invitacionID)
	if err != nil {
		t.Fatalf("Error al crear objeto notificación: %v", err)
	}

	// Guardar notificación en la base de datos
	err = notificacionDAO.Guardar(notificacion)
	if err != nil {
		t.Fatalf("Error al guardar notificación: %v", err)
	}

	// Eliminar la notificación
	err = notificacionDAO.Eliminar(id)
	if err != nil {
		t.Fatalf("Error al eliminar notificación: %v", err)
	}

	// Verificar que ya no existe
	reobtenida, err := notificacionDAO.BuscarPorID(id)
	if err != nil {
		t.Fatalf("Error al buscar notificación eliminada: %v", err)
	}

	if reobtenida != nil {
		t.Errorf("La notificación no debería existir después de eliminarla")
	}
}
