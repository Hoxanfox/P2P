package dao

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"model"
)

// TestInvitacionCanalDAO_Crear verifica que el método Guardar funcione correctamente
func TestInvitacionCanalDAO_Crear(t *testing.T) {
	// Inicializar conexión a la base de datos
	dbPool := setupTestDB(t)
	defer cleanupTestDB(t, dbPool)

	// Crear un canal primero ya que es necesario para la relación
	canalDAO := NuevoCanalDAO(dbPool)
	canalID := uuid.New()
	destinatarioID := uuid.New()

	// Crear canal en la base de datos
	canal, err := model.NewCanalServidor(canalID, "Canal para Invitaciones", "Canal para probar invitaciones", model.CanalPublico)
	if err != nil {
		t.Fatalf("Error creando objeto canal: %v", err)
	}

	err = canalDAO.Crear(canal)
	if err != nil {
		t.Fatalf("Error creando canal en la base de datos: %v", err)
	}

	// Crear instancia del DAO de invitaciones
	invitacionDAO := NuevoInvitacionCanalDAO(dbPool)

	// Crear una invitación
	id := uuid.New()
	estado := model.InvitacionPendiente
	fechaEnvio := time.Now()

	invitacion, err := model.NewInvitacionCanal(id, canalID, destinatarioID, estado, fechaEnvio)
	if err != nil {
		t.Fatalf("Error creando objeto invitación: %v", err)
	}

	// Ejecutar operación a probar
	err = invitacionDAO.Guardar(invitacion)
	if err != nil {
		t.Fatalf("Error guardando invitación: %v", err)
	}

	// Verificar que se guardó correctamente
	invitacionRecuperada, err := invitacionDAO.BuscarPorID(id)
	if err != nil {
		t.Fatalf("Error recuperando invitación: %v", err)
	}

	if invitacionRecuperada == nil {
		t.Fatal("No se encontró la invitación guardada")
	}

	// Verificar que los datos son correctos
	if invitacionRecuperada.ID() != id {
		t.Errorf("ID incorrecto. Esperado: %s, Obtenido: %s", id, invitacionRecuperada.ID())
	}

	if invitacionRecuperada.CanalID() != canalID {
		t.Errorf("CanalID incorrecto. Esperado: %s, Obtenido: %s", canalID, invitacionRecuperada.CanalID())
	}

	if invitacionRecuperada.DestinatarioID() != destinatarioID {
		t.Errorf("DestinatarioID incorrecto. Esperado: %s, Obtenido: %s", destinatarioID, invitacionRecuperada.DestinatarioID())
	}

	if invitacionRecuperada.Estado() != estado {
		t.Errorf("Estado incorrecto. Esperado: %s, Obtenido: %s", estado, invitacionRecuperada.Estado())
	}

	// Limpiar datos de prueba
	_, err = dbPool.ExecContext(
		context.Background(),
		"DELETE FROM invitaciones_canal WHERE id = ?",
		id.String(),
	)
	if err != nil {
		t.Logf("Error limpiando invitación de prueba: %v", err)
	}

	_, err = dbPool.ExecContext(
		context.Background(),
		"DELETE FROM canales WHERE id = ?",
		canalID.String(),
	)
	if err != nil {
		t.Logf("Error limpiando canal de prueba: %v", err)
	}
}

// TestInvitacionCanalDAO_ObtenerPorID verifica que el método ObtenerPorID funcione correctamente
func TestInvitacionCanalDAO_ObtenerPorID(t *testing.T) {
	// Inicializar conexión a la base de datos
	dbPool := setupTestDB(t)
	defer cleanupTestDB(t, dbPool)

	// Crear un canal para las pruebas
	canalDAO := NuevoCanalDAO(dbPool)
	canalID := uuid.New()
	destinatarioID := uuid.New()

	// Crear canal en la base de datos
	canal, err := model.NewCanalServidor(canalID, "Canal para BuscarPorID", "Canal para probar BuscarPorID", model.CanalPublico)
	if err != nil {
		t.Fatalf("Error creando objeto canal: %v", err)
	}

	err = canalDAO.Crear(canal)
	if err != nil {
		t.Fatalf("Error creando canal en la base de datos: %v", err)
	}

	// Crear instancia del DAO de invitaciones
	invitacionDAO := NuevoInvitacionCanalDAO(dbPool)

	// Crear y guardar una invitación
	id := uuid.New()
	estado := model.InvitacionPendiente
	fechaEnvio := time.Now().Round(time.Second) // Redondeamos para evitar problemas con microsegundos

	invitacion, err := model.NewInvitacionCanal(id, canalID, destinatarioID, estado, fechaEnvio)
	if err != nil {
		t.Fatalf("Error creando objeto invitación: %v", err)
	}

	err = invitacionDAO.Guardar(invitacion)
	if err != nil {
		t.Fatalf("Error guardando invitación: %v", err)
	}

	// Ejecutar operación a probar
	invitacionRecuperada, err := invitacionDAO.BuscarPorID(id)
	if err != nil {
		t.Fatalf("Error buscando invitación por ID: %v", err)
	}

	// Verificar que se recuperó correctamente
	if invitacionRecuperada == nil {
		t.Fatal("No se encontró la invitación")
	}

	// Verificar que los datos son correctos
	if invitacionRecuperada.ID() != id {
		t.Errorf("ID incorrecto. Esperado: %s, Obtenido: %s", id, invitacionRecuperada.ID())
	}

	if invitacionRecuperada.CanalID() != canalID {
		t.Errorf("CanalID incorrecto. Esperado: %s, Obtenido: %s", canalID, invitacionRecuperada.CanalID())
	}

	if invitacionRecuperada.DestinatarioID() != destinatarioID {
		t.Errorf("DestinatarioID incorrecto. Esperado: %s, Obtenido: %s", destinatarioID, invitacionRecuperada.DestinatarioID())
	}

	if invitacionRecuperada.Estado() != estado {
		t.Errorf("Estado incorrecto. Esperado: %s, Obtenido: %s", estado, invitacionRecuperada.Estado())
	}

	// Limpiar datos de prueba
	_, err = dbPool.ExecContext(
		context.Background(),
		"DELETE FROM invitaciones_canal WHERE id = ?",
		id.String(),
	)
	if err != nil {
		t.Logf("Error limpiando invitación de prueba: %v", err)
	}

	_, err = dbPool.ExecContext(
		context.Background(),
		"DELETE FROM canales WHERE id = ?",
		canalID.String(),
	)
	if err != nil {
		t.Logf("Error limpiando canal de prueba: %v", err)
	}
}

// TestInvitacionCanalDAO_ObtenerPorDestinatario verifica que el método ObtenerPorDestinatario funcione correctamente
func TestInvitacionCanalDAO_ObtenerPorDestinatario(t *testing.T) {
	// Inicializar conexión a la base de datos
	dbPool := setupTestDB(t)
	defer cleanupTestDB(t, dbPool)

	// Crear un canal para las pruebas
	canalDAO := NuevoCanalDAO(dbPool)
	canalID := uuid.New()
	destinatarioID := uuid.New() // Este será el destinatario para todas las invitaciones

	// Crear canal en la base de datos
	canal, err := model.NewCanalServidor(canalID, "Canal para Destinatario", "Canal para probar BuscarPorDestinatarioID", model.CanalPublico)
	if err != nil {
		t.Fatalf("Error creando objeto canal: %v", err)
	}

	err = canalDAO.Crear(canal)
	if err != nil {
		t.Fatalf("Error creando canal en la base de datos: %v", err)
	}

	// Crear instancia del DAO de invitaciones
	invitacionDAO := NuevoInvitacionCanalDAO(dbPool)

	// Crear varias invitaciones para el mismo destinatario
	id1 := uuid.New()
	id2 := uuid.New()
	estado1 := model.InvitacionPendiente
	estado2 := model.InvitacionAceptada
	fechaEnvio := time.Now().Round(time.Second)

	invitacion1, err := model.NewInvitacionCanal(id1, canalID, destinatarioID, estado1, fechaEnvio)
	if err != nil {
		t.Fatalf("Error creando invitación 1: %v", err)
	}

	invitacion2, err := model.NewInvitacionCanal(id2, canalID, destinatarioID, estado2, fechaEnvio)
	if err != nil {
		t.Fatalf("Error creando invitación 2: %v", err)
	}

	// Guardar las invitaciones
	err = invitacionDAO.Guardar(invitacion1)
	if err != nil {
		t.Fatalf("Error guardando invitación 1: %v", err)
	}

	err = invitacionDAO.Guardar(invitacion2)
	if err != nil {
		t.Fatalf("Error guardando invitación 2: %v", err)
	}

	// Ejecutar operación a probar
	invitaciones, err := invitacionDAO.BuscarPorDestinatarioID(destinatarioID)
	if err != nil {
		t.Fatalf("Error buscando invitaciones por destinatario: %v", err)
	}

	// Verificar que se recuperaron las invitaciones correctamente
	if len(invitaciones) != 2 {
		t.Fatalf("Número incorrecto de invitaciones. Esperado: 2, Obtenido: %d", len(invitaciones))
	}

	// Verificar que los IDs de las invitaciones son correctos
	found1, found2 := false, false
	for _, inv := range invitaciones {
		if inv.ID() == id1 {
			found1 = true
			if inv.Estado() != estado1 {
				t.Errorf("Estado incorrecto para invitación 1. Esperado: %s, Obtenido: %s", estado1, inv.Estado())
			}
		}
		if inv.ID() == id2 {
			found2 = true
			if inv.Estado() != estado2 {
				t.Errorf("Estado incorrecto para invitación 2. Esperado: %s, Obtenido: %s", estado2, inv.Estado())
			}
		}
	}

	if !found1 {
		t.Error("No se encontró la invitación 1 en los resultados")
	}

	if !found2 {
		t.Error("No se encontró la invitación 2 en los resultados")
	}

	// Limpiar datos de prueba
	_, err = dbPool.ExecContext(
		context.Background(),
		"DELETE FROM invitaciones_canal WHERE id IN (?, ?)",
		id1.String(), id2.String(),
	)
	if err != nil {
		t.Logf("Error limpiando invitaciones de prueba: %v", err)
	}

	_, err = dbPool.ExecContext(
		context.Background(),
		"DELETE FROM canales WHERE id = ?",
		canalID.String(),
	)
	if err != nil {
		t.Logf("Error limpiando canal de prueba: %v", err)
	}
}

// TestInvitacionCanalDAO_Actualizar verifica que el método Actualizar funcione correctamente
func TestInvitacionCanalDAO_Actualizar(t *testing.T) {
	// Inicializar conexión a la base de datos
	dbPool := setupTestDB(t)
	defer cleanupTestDB(t, dbPool)

	// Crear un canal para las pruebas
	canalDAO := NuevoCanalDAO(dbPool)
	canalID := uuid.New()
	destinatarioID := uuid.New()

	// Crear canal en la base de datos
	canal, err := model.NewCanalServidor(canalID, "Canal para Actualizar", "Canal para probar Actualizar", model.CanalPublico)
	if err != nil {
		t.Fatalf("Error creando objeto canal: %v", err)
	}

	err = canalDAO.Crear(canal)
	if err != nil {
		t.Fatalf("Error creando canal en la base de datos: %v", err)
	}

	// Crear instancia del DAO de invitaciones
	invitacionDAO := NuevoInvitacionCanalDAO(dbPool)

	// Crear y guardar una invitación con estado inicial
	id := uuid.New()
	estadoInicial := model.InvitacionPendiente
	fechaEnvio := time.Now().Round(time.Second)

	invitacion, err := model.NewInvitacionCanal(id, canalID, destinatarioID, estadoInicial, fechaEnvio)
	if err != nil {
		t.Fatalf("Error creando objeto invitación: %v", err)
	}

	err = invitacionDAO.Guardar(invitacion)
	if err != nil {
		t.Fatalf("Error guardando invitación: %v", err)
	}

	// Verificar que se guardó con el estado inicial
	invitacionInicial, err := invitacionDAO.BuscarPorID(id)
	if err != nil {
		t.Fatalf("Error recuperando invitación inicial: %v", err)
	}

	if invitacionInicial.Estado() != estadoInicial {
		t.Errorf("Estado inicial incorrecto. Esperado: %s, Obtenido: %s", estadoInicial, invitacionInicial.Estado())
	}

	// Cambiar el estado de la invitación
	nuevoEstado := model.InvitacionAceptada
	invitacionActualizada, err := model.NewInvitacionCanal(id, canalID, destinatarioID, nuevoEstado, fechaEnvio)
	if err != nil {
		t.Fatalf("Error creando objeto invitación actualizada: %v", err)
	}

	// Ejecutar operación a probar
	err = invitacionDAO.Actualizar(invitacionActualizada)
	if err != nil {
		t.Fatalf("Error actualizando invitación: %v", err)
	}

	// Verificar que se actualizó correctamente
	invitacionFinal, err := invitacionDAO.BuscarPorID(id)
	if err != nil {
		t.Fatalf("Error recuperando invitación actualizada: %v", err)
	}

	if invitacionFinal.Estado() != nuevoEstado {
		t.Errorf("Estado actualizado incorrecto. Esperado: %s, Obtenido: %s", nuevoEstado, invitacionFinal.Estado())
	}

	// Limpiar datos de prueba
	_, err = dbPool.ExecContext(
		context.Background(),
		"DELETE FROM invitaciones_canal WHERE id = ?",
		id.String(),
	)
	if err != nil {
		t.Logf("Error limpiando invitación de prueba: %v", err)
	}

	_, err = dbPool.ExecContext(
		context.Background(),
		"DELETE FROM canales WHERE id = ?",
		canalID.String(),
	)
	if err != nil {
		t.Logf("Error limpiando canal de prueba: %v", err)
	}
}

// TestInvitacionCanalDAO_Eliminar verifica que el método Eliminar funcione correctamente
func TestInvitacionCanalDAO_Eliminar(t *testing.T) {
	// Inicializar conexión a la base de datos
	dbPool := setupTestDB(t)
	defer cleanupTestDB(t, dbPool)

	// Crear un canal para las pruebas
	canalDAO := NuevoCanalDAO(dbPool)
	canalID := uuid.New()
	destinatarioID := uuid.New()

	// Crear canal en la base de datos
	canal, err := model.NewCanalServidor(canalID, "Canal para Eliminar", "Canal para probar Eliminar", model.CanalPublico)
	if err != nil {
		t.Fatalf("Error creando objeto canal: %v", err)
	}

	err = canalDAO.Crear(canal)
	if err != nil {
		t.Fatalf("Error creando canal en la base de datos: %v", err)
	}

	// Crear instancia del DAO de invitaciones
	invitacionDAO := NuevoInvitacionCanalDAO(dbPool)

	// Crear y guardar una invitación
	id := uuid.New()
	estado := model.InvitacionPendiente
	fechaEnvio := time.Now().Round(time.Second)

	invitacion, err := model.NewInvitacionCanal(id, canalID, destinatarioID, estado, fechaEnvio)
	if err != nil {
		t.Fatalf("Error creando objeto invitación: %v", err)
	}

	err = invitacionDAO.Guardar(invitacion)
	if err != nil {
		t.Fatalf("Error guardando invitación: %v", err)
	}

	// Verificar que se guardó correctamente
	invitacionGuardada, err := invitacionDAO.BuscarPorID(id)
	if err != nil {
		t.Fatalf("Error recuperando invitación: %v", err)
	}

	if invitacionGuardada == nil {
		t.Fatal("No se encontró la invitación guardada")
	}

	// Ejecutar operación a probar
	err = invitacionDAO.Eliminar(id)
	if err != nil {
		t.Fatalf("Error eliminando invitación: %v", err)
	}

	// Verificar que se eliminó correctamente
	invitacionEliminada, err := invitacionDAO.BuscarPorID(id)
	if err != nil {
		t.Fatalf("Error buscando invitación eliminada: %v", err)
	}

	if invitacionEliminada != nil {
		t.Fatal("La invitación no se eliminó correctamente")
	}

	// Limpiar el canal de prueba
	_, err = dbPool.ExecContext(
		context.Background(),
		"DELETE FROM canales WHERE id = ?",
		canalID.String(),
	)
	if err != nil {
		t.Logf("Error limpiando canal de prueba: %v", err)
	}
}
