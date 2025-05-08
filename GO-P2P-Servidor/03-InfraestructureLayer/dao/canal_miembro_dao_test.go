package dao

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"model"
)

// TestCanalMiembroDAO_GuardarCanalMiembro verifica la funcionalidad básica del método GuardarCanalMiembro
func TestCanalMiembroDAO_GuardarCanalMiembro(t *testing.T) {
	// Inicializar conexión a la base de datos
	dbPool := setupTestDB(t)
	defer cleanupTestDB(t, dbPool)

	// Crear un canal primero ya que es necesario para la relación
	canalDAO := NuevoCanalDAO(dbPool)
	canalID := uuid.New()
	usuarioID := uuid.New() // ID del usuario que será miembro

	// Crear canal en la base de datos
	canal, err := model.NewCanalServidor(canalID, "Canal Prueba Miembros", "Canal para probar miembros", model.CanalPublico)
	if err != nil {
		t.Fatalf("Error creando objeto canal: %v", err)
	}

	err = canalDAO.Crear(canal)
	if err != nil {
		t.Fatalf("Error creando canal en la base de datos: %v", err)
	}

	// Crear instancia del DAO de miembros
	canalMiembroDAO := NuevoCanalMiembroDAO(dbPool)

	// Crear miembro de prueba
	rol := "ADMIN" // Roles posibles: ADMIN, MODERADOR, MIEMBRO, etc.
	miembro, err := model.NewCanalMiembro(canalID, usuarioID, rol)
	if err != nil {
		t.Fatalf("Error creando objeto miembro: %v", err)
	}

	// Ejecutar operación de guardar
	err = canalMiembroDAO.Guardar(miembro)
	if err != nil {
		t.Fatalf("Error guardando miembro: %v", err)
	}

	// Verificar que se guardó correctamente
	miembroRecuperado, err := canalMiembroDAO.BuscarPorIDs(canalID, usuarioID)
	if err != nil {
		t.Fatalf("Error recuperando miembro: %v", err)
	}

	if miembroRecuperado == nil {
		t.Fatal("No se encontró el miembro guardado")
	}

	if miembroRecuperado.CanalID() != canalID {
		t.Errorf("CanalID incorrecto. Esperado: %s, Obtenido: %s", canalID, miembroRecuperado.CanalID())
	}

	if miembroRecuperado.UsuarioID() != usuarioID {
		t.Errorf("UsuarioID incorrecto. Esperado: %s, Obtenido: %s", usuarioID, miembroRecuperado.UsuarioID())
	}

	if miembroRecuperado.Rol() != rol {
		t.Errorf("Rol incorrecto. Esperado: %s, Obtenido: %s", rol, miembroRecuperado.Rol())
	}

	// Limpiar datos de prueba
	_, err = dbPool.ExecContext(
		context.Background(),
		"DELETE FROM canal_miembros WHERE canal_id = ? AND usuario_id = ?",
		canalID.String(), usuarioID.String(),
	)
	if err != nil {
		t.Logf("Error limpiando miembro de prueba: %v", err)
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

// TestCanalMiembroDAO_ObtenerCanalMiembro verifica la funcionalidad básica del método ObtenerCanalMiembro
func TestCanalMiembroDAO_ObtenerCanalMiembro(t *testing.T) {
	// Inicializar conexión a la base de datos
	dbPool := setupTestDB(t)
	defer cleanupTestDB(t, dbPool)

	// Crear un canal primero ya que es necesario para la relación
	canalDAO := NuevoCanalDAO(dbPool)
	canalID := uuid.New()
	usuarioID := uuid.New()

	// Crear canal en la base de datos
	canal, err := model.NewCanalServidor(canalID, "Canal para Obtener Miembro", "Canal para probar ObtenerCanalMiembro", model.CanalPrivado)
	if err != nil {
		t.Fatalf("Error creando objeto canal: %v", err)
	}

	err = canalDAO.Crear(canal)
	if err != nil {
		t.Fatalf("Error creando canal en la base de datos: %v", err)
	}

	// Crear instancia del DAO de miembros
	canalMiembroDAO := NuevoCanalMiembroDAO(dbPool)

	// Crear y guardar miembro de prueba
	rol := "MODERADOR"
	miembro, err := model.NewCanalMiembro(canalID, usuarioID, rol)
	if err != nil {
		t.Fatalf("Error creando objeto miembro: %v", err)
	}

	err = canalMiembroDAO.Guardar(miembro)
	if err != nil {
		t.Fatalf("Error guardando miembro: %v", err)
	}

	// Ejecutar operación a probar
	miembroObtenido, err := canalMiembroDAO.BuscarPorIDs(canalID, usuarioID)
	if err != nil {
		t.Fatalf("Error en BuscarPorIDs: %v", err)
	}

	// Verificaciones
	if miembroObtenido == nil {
		t.Fatal("No se encontró el miembro")
	}

	if miembroObtenido.CanalID() != canalID {
		t.Errorf("CanalID incorrecto. Esperado: %s, Obtenido: %s", canalID, miembroObtenido.CanalID())
	}

	if miembroObtenido.UsuarioID() != usuarioID {
		t.Errorf("UsuarioID incorrecto. Esperado: %s, Obtenido: %s", usuarioID, miembroObtenido.UsuarioID())
	}

	if miembroObtenido.Rol() != rol {
		t.Errorf("Rol incorrecto. Esperado: %s, Obtenido: %s", rol, miembroObtenido.Rol())
	}

	// Limpiar datos de prueba
	_, err = dbPool.ExecContext(
		context.Background(),
		"DELETE FROM canal_miembros WHERE canal_id = ? AND usuario_id = ?",
		canalID.String(), usuarioID.String(),
	)
	if err != nil {
		t.Logf("Error limpiando miembro de prueba: %v", err)
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

// TestCanalMiembroDAO_ObtenerMiembrosCanal verifica la funcionalidad básica del método ObtenerMiembrosCanal
func TestCanalMiembroDAO_ObtenerMiembrosCanal(t *testing.T) {
	// Inicializar conexión a la base de datos
	dbPool := setupTestDB(t)
	defer cleanupTestDB(t, dbPool)

	// Crear un canal para las pruebas
	canalDAO := NuevoCanalDAO(dbPool)
	canalID := uuid.New()

	canal, err := model.NewCanalServidor(canalID, "Canal para Miembros", "Canal para probar ObtenerMiembrosCanal", model.CanalPublico)
	if err != nil {
		t.Fatalf("Error creando objeto canal: %v", err)
	}

	err = canalDAO.Crear(canal)
	if err != nil {
		t.Fatalf("Error creando canal en la base de datos: %v", err)
	}

	// Crear instancia del DAO de miembros
	canalMiembroDAO := NuevoCanalMiembroDAO(dbPool)

	// Crear varios miembros para el canal
	usuarioID1 := uuid.New()
	usuarioID2 := uuid.New()

	miembro1, err := model.NewCanalMiembro(canalID, usuarioID1, "ADMIN")
	if err != nil {
		t.Fatalf("Error creando miembro 1: %v", err)
	}

	miembro2, err := model.NewCanalMiembro(canalID, usuarioID2, "MIEMBRO")
	if err != nil {
		t.Fatalf("Error creando miembro 2: %v", err)
	}

	// Guardar los miembros
	err = canalMiembroDAO.Guardar(miembro1)
	if err != nil {
		t.Fatalf("Error guardando miembro 1: %v", err)
	}

	err = canalMiembroDAO.Guardar(miembro2)
	if err != nil {
		t.Fatalf("Error guardando miembro 2: %v", err)
	}

	// Ejecutar operación a probar
	miembros, err := canalMiembroDAO.BuscarPorCanalID(canalID)
	if err != nil {
		t.Fatalf("Error en BuscarPorCanalID: %v", err)
	}

	// Verificar que se obtuvieron los miembros esperados
	if len(miembros) != 2 {
		t.Fatalf("Número incorrecto de miembros. Esperado: 2, Obtenido: %d", len(miembros))
	}

	// Verificar que los IDs de los miembros son correctos
	found1, found2 := false, false
	for _, m := range miembros {
		if m.UsuarioID() == usuarioID1 {
			found1 = true
			if m.Rol() != "ADMIN" {
				t.Errorf("Rol incorrecto para miembro 1. Esperado: ADMIN, Obtenido: %s", m.Rol())
			}
		}
		if m.UsuarioID() == usuarioID2 {
			found2 = true
			if m.Rol() != "MIEMBRO" {
				t.Errorf("Rol incorrecto para miembro 2. Esperado: MIEMBRO, Obtenido: %s", m.Rol())
			}
		}
	}

	if !found1 {
		t.Error("No se encontró el miembro 1 en los resultados")
	}

	if !found2 {
		t.Error("No se encontró el miembro 2 en los resultados")
	}

	// Limpiar datos de prueba
	_, err = dbPool.ExecContext(
		context.Background(),
		"DELETE FROM canal_miembros WHERE canal_id = ?",
		canalID.String(),
	)
	if err != nil {
		t.Logf("Error limpiando miembros de prueba: %v", err)
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

// TestCanalMiembroDAO_ActualizarRol verifica la funcionalidad básica del método ActualizarRol
func TestCanalMiembroDAO_ActualizarRol(t *testing.T) {
	// Inicializar conexión a la base de datos
	dbPool := setupTestDB(t)
	defer cleanupTestDB(t, dbPool)

	// Crear un canal para las pruebas
	canalDAO := NuevoCanalDAO(dbPool)
	canalID := uuid.New()
	usuarioID := uuid.New()

	canal, err := model.NewCanalServidor(canalID, "Canal para ActualizarRol", "Canal para probar la actualización de rol", model.CanalPublico)
	if err != nil {
		t.Fatalf("Error creando objeto canal: %v", err)
	}

	err = canalDAO.Crear(canal)
	if err != nil {
		t.Fatalf("Error creando canal en la base de datos: %v", err)
	}

	// Crear instancia del DAO de miembros
	canalMiembroDAO := NuevoCanalMiembroDAO(dbPool)

	// Crear y guardar miembro con rol inicial
	rolInicial := "MIEMBRO"
	miembro, err := model.NewCanalMiembro(canalID, usuarioID, rolInicial)
	if err != nil {
		t.Fatalf("Error creando miembro: %v", err)
	}

	err = canalMiembroDAO.Guardar(miembro)
	if err != nil {
		t.Fatalf("Error guardando miembro: %v", err)
	}

	// Verificar que se guardó correctamente con el rol inicial
	miembroInicial, err := canalMiembroDAO.BuscarPorIDs(canalID, usuarioID)
	if err != nil {
		t.Fatalf("Error buscando miembro: %v", err)
	}

	if miembroInicial == nil {
		t.Fatal("No se encontró el miembro")
	}

	if miembroInicial.Rol() != rolInicial {
		t.Errorf("Rol inicial incorrecto. Esperado: %s, Obtenido: %s", rolInicial, miembroInicial.Rol())
	}

	// Crear objeto actualizado con nuevo rol
	nuevoRol := "ADMIN"
	miembroActualizado, err := model.NewCanalMiembro(canalID, usuarioID, nuevoRol)
	if err != nil {
		t.Fatalf("Error creando miembro actualizado: %v", err)
	}

	// Ejecutar operación de actualización
	err = canalMiembroDAO.Actualizar(miembroActualizado)
	if err != nil {
		t.Fatalf("Error actualizando rol: %v", err)
	}

	// Verificar que el rol se actualizó correctamente
	miembroFinal, err := canalMiembroDAO.BuscarPorIDs(canalID, usuarioID)
	if err != nil {
		t.Fatalf("Error buscando miembro actualizado: %v", err)
	}

	if miembroFinal == nil {
		t.Fatal("No se encontró el miembro después de actualizar")
	}

	if miembroFinal.Rol() != nuevoRol {
		t.Errorf("Rol actualizado incorrecto. Esperado: %s, Obtenido: %s", nuevoRol, miembroFinal.Rol())
	}

	// Limpiar datos de prueba
	_, err = dbPool.ExecContext(
		context.Background(),
		"DELETE FROM canal_miembros WHERE canal_id = ? AND usuario_id = ?",
		canalID.String(), usuarioID.String(),
	)
	if err != nil {
		t.Logf("Error limpiando miembro de prueba: %v", err)
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

// TestCanalMiembroDAO_EliminarMiembro verifica la funcionalidad básica del método EliminarMiembro
func TestCanalMiembroDAO_EliminarMiembro(t *testing.T) {
	// Inicializar conexión a la base de datos
	dbPool := setupTestDB(t)
	defer cleanupTestDB(t, dbPool)

	// Crear un canal para las pruebas
	canalDAO := NuevoCanalDAO(dbPool)
	canalID := uuid.New()
	usuarioID := uuid.New()

	canal, err := model.NewCanalServidor(canalID, "Canal para EliminarMiembro", "Canal para probar la eliminación de miembros", model.CanalPublico)
	if err != nil {
		t.Fatalf("Error creando objeto canal: %v", err)
	}

	err = canalDAO.Crear(canal)
	if err != nil {
		t.Fatalf("Error creando canal en la base de datos: %v", err)
	}

	// Crear instancia del DAO de miembros
	canalMiembroDAO := NuevoCanalMiembroDAO(dbPool)

	// Crear y guardar miembro 
	miembro, err := model.NewCanalMiembro(canalID, usuarioID, "MIEMBRO")
	if err != nil {
		t.Fatalf("Error creando miembro: %v", err)
	}

	err = canalMiembroDAO.Guardar(miembro)
	if err != nil {
		t.Fatalf("Error guardando miembro: %v", err)
	}

	// Verificar que se guardó correctamente
	miembroGuardado, err := canalMiembroDAO.BuscarPorIDs(canalID, usuarioID)
	if err != nil {
		t.Fatalf("Error buscando miembro: %v", err)
	}

	if miembroGuardado == nil {
		t.Fatal("No se encontró el miembro creado")
	}

	// Ejecutar operación de eliminación
	err = canalMiembroDAO.Eliminar(canalID, usuarioID)
	if err != nil {
		t.Fatalf("Error eliminando miembro: %v", err)
	}

	// Verificar que se eliminó correctamente
	miembroEliminado, err := canalMiembroDAO.BuscarPorIDs(canalID, usuarioID)
	if err != nil {
		t.Fatalf("Error buscando miembro eliminado: %v", err)
	}

	if miembroEliminado != nil {
		t.Error("El miembro no se eliminó correctamente, aún existe en la base de datos")
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
