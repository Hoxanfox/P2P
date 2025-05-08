package dao

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"model"
)

// TestCanalDAO_Crear verifica que el método Crear funcione correctamente
func TestCanalDAO_Crear(t *testing.T) {
	// Inicializar conexión a la base de datos
	dbPool := setupTestDB(t)
	defer cleanupTestDB(t, dbPool)

	// Crear instancia del DAO
	canalDAO := NuevoCanalDAO(dbPool)

	// Datos de prueba
	id := uuid.New()
	nombre := "Canal de Prueba"
	descripcion := "Este es un canal de prueba para el test"
	tipo := model.CanalPublico

	// Crear objeto canal
	canal, err := model.NewCanalServidor(id, nombre, descripcion, tipo)
	if err != nil {
		t.Fatalf("Error creando objeto canal: %v", err)
	}

	// Ejecutar operación
	err = canalDAO.Crear(canal)
	if err != nil {
		t.Fatalf("Error creando canal en la base de datos: %v", err)
	}

	// Verificar que se creó correctamente consultándolo
	canalObtenido, err := canalDAO.BuscarPorID(id)
	if err != nil {
		t.Fatalf("Error consultando canal creado: %v", err)
	}

	if canalObtenido == nil {
		t.Fatal("No se pudo obtener el canal recién creado")
	}

	// Verificar que los datos coincidan
	if canalObtenido.ID() != id {
		t.Errorf("ID incorrecto. Esperado: %s, Obtenido: %s", id, canalObtenido.ID())
	}

	if canalObtenido.Nombre() != nombre {
		t.Errorf("Nombre incorrecto. Esperado: %s, Obtenido: %s", nombre, canalObtenido.Nombre())
	}

	if canalObtenido.Descripcion() != descripcion {
		t.Errorf("Descripcion incorrecta. Esperado: %s, Obtenido: %s", descripcion, canalObtenido.Descripcion())
	}

	if canalObtenido.Tipo() != tipo {
		t.Errorf("Tipo incorrecto. Esperado: %s, Obtenido: %s", tipo, canalObtenido.Tipo())
	}

	// Limpiar después de la prueba
	_, err = dbPool.ExecContext(
		context.Background(),
		"DELETE FROM canales WHERE id = ?",
		id.String(),
	)
	if err != nil {
		t.Logf("Error al limpiar canal de prueba: %v", err)
	}
}

// TestCanalDAO_BuscarPorID verifica que el método BuscarPorID funcione correctamente
func TestCanalDAO_BuscarPorID(t *testing.T) {
	// Inicializar conexión a la base de datos
	dbPool := setupTestDB(t)
	defer cleanupTestDB(t, dbPool)

	// Crear instancia del DAO
	canalDAO := NuevoCanalDAO(dbPool)

	// Crear canal de prueba
	id := uuid.New()
	nombre := "Canal para BuscarPorID"
	descripcion := "Este es un canal para probar BuscarPorID"
	tipo := model.CanalPrivado

	// Crear objeto canal
	canal, err := model.NewCanalServidor(id, nombre, descripcion, tipo)
	if err != nil {
		t.Fatalf("Error creando objeto canal: %v", err)
	}

	// Guardarlo en la base de datos
	err = canalDAO.Crear(canal)
	if err != nil {
		t.Fatalf("Error creando canal en la base de datos: %v", err)
	}

	// Ejecutar operación que estamos probando
	canalObtenido, err := canalDAO.BuscarPorID(id)
	if err != nil {
		t.Fatalf("Error en BuscarPorID: %v", err)
	}

	// Verificar que se haya encontrado
	if canalObtenido == nil {
		t.Fatal("No se encontró el canal")
	}

	// Verificar que los datos son correctos
	if canalObtenido.ID() != id {
		t.Errorf("ID incorrecto. Esperado: %s, Obtenido: %s", id, canalObtenido.ID())
	}

	if canalObtenido.Nombre() != nombre {
		t.Errorf("Nombre incorrecto. Esperado: %s, Obtenido: %s", nombre, canalObtenido.Nombre())
	}

	if canalObtenido.Descripcion() != descripcion {
		t.Errorf("Descripcion incorrecta. Esperado: %s, Obtenido: %s", descripcion, canalObtenido.Descripcion())
	}

	if canalObtenido.Tipo() != tipo {
		t.Errorf("Tipo incorrecto. Esperado: %s, Obtenido: %s", tipo, canalObtenido.Tipo())
	}

	// Limpiar datos de prueba
	_, err = dbPool.ExecContext(
		context.Background(),
		"DELETE FROM canales WHERE id = ?",
		id.String(),
	)
	if err != nil {
		t.Logf("Error al limpiar canal de prueba: %v", err)
	}
}

// TestCanalDAO_BuscarTodos verifica que el método BuscarTodos funcione correctamente
func TestCanalDAO_BuscarTodos(t *testing.T) {
	// Inicializar conexión a la base de datos
	dbPool := setupTestDB(t)
	defer cleanupTestDB(t, dbPool)

	// Crear instancia del DAO
	canalDAO := NuevoCanalDAO(dbPool)

	// Crear canales de prueba
	id1 := uuid.New()
	id2 := uuid.New()
	
	canal1, err := model.NewCanalServidor(id1, "Canal 1", "Descripción 1", model.CanalPublico)
	if err != nil {
		t.Fatalf("Error creando objeto canal 1: %v", err)
	}
	
	canal2, err := model.NewCanalServidor(id2, "Canal 2", "Descripción 2", model.CanalPrivado)
	if err != nil {
		t.Fatalf("Error creando objeto canal 2: %v", err)
	}

	// Guardar canales en la base de datos
	err = canalDAO.Crear(canal1)
	if err != nil {
		t.Fatalf("Error creando canal 1: %v", err)
	}

	err = canalDAO.Crear(canal2)
	if err != nil {
		t.Fatalf("Error creando canal 2: %v", err)
	}

	// Ejecutar operación
	canales, err := canalDAO.BuscarTodos()
	if err != nil {
		t.Fatalf("Error en BuscarTodos: %v", err)
	}

	// Verificar que se hayan encontrado al menos los dos canales que creamos
	// Nota: pueden haber más canales si otras pruebas ejecutadas anteriormente
	// crearon canales y no los eliminaron
	if len(canales) < 2 {
		t.Fatalf("Se esperaban al menos 2 canales, se obtuvieron %d", len(canales))
	}

	// Verificar que nuestros canales estén en los resultados
	enc1, enc2 := false, false
	for _, c := range canales {
		if c.ID() == id1 {
			enc1 = true
			if c.Nombre() != "Canal 1" {
				t.Errorf("Nombre incorrecto para canal 1. Esperado: %s, Obtenido: %s", "Canal 1", c.Nombre())
			}
		}
		if c.ID() == id2 {
			enc2 = true
			if c.Nombre() != "Canal 2" {
				t.Errorf("Nombre incorrecto para canal 2. Esperado: %s, Obtenido: %s", "Canal 2", c.Nombre())
			}
		}
	}

	if !enc1 {
		t.Error("El canal 1 no fue encontrado en los resultados")
	}

	if !enc2 {
		t.Error("El canal 2 no fue encontrado en los resultados")
	}

	// Limpiar datos de prueba
	_, err = dbPool.ExecContext(
		context.Background(),
		"DELETE FROM canales WHERE id IN (?, ?)",
		id1.String(), id2.String(),
	)
	if err != nil {
		t.Logf("Error al limpiar canales de prueba: %v", err)
	}
}

// TestCanalDAO_Actualizar verifica que el método Actualizar funcione correctamente
func TestCanalDAO_Actualizar(t *testing.T) {
	// Inicializar conexión a la base de datos
	dbPool := setupTestDB(t)
	defer cleanupTestDB(t, dbPool)

	// Crear instancia del DAO
	canalDAO := NuevoCanalDAO(dbPool)

	// Crear canal de prueba
	id := uuid.New()
	nombre := "Canal Original"
	descripcion := "Descripción original"
	tipo := model.CanalPublico

	canal, err := model.NewCanalServidor(id, nombre, descripcion, tipo)
	if err != nil {
		t.Fatalf("Error creando objeto canal: %v", err)
	}

	// Guardarlo en la base de datos
	err = canalDAO.Crear(canal)
	if err != nil {
		t.Fatalf("Error creando canal en la base de datos: %v", err)
	}

	// Modificar canal para actualizarlo
	nuevoNombre := "Canal Actualizado"
	nuevaDescripcion := "Descripción actualizada"
	nuevoTipo := model.CanalPrivado

	canalActualizado, err := model.NewCanalServidor(id, nuevoNombre, nuevaDescripcion, nuevoTipo)
	if err != nil {
		t.Fatalf("Error creando objeto canal actualizado: %v", err)
	}

	// Ejecutar operación de actualización
	err = canalDAO.Actualizar(canalActualizado)
	if err != nil {
		t.Fatalf("Error actualizando canal: %v", err)
	}

	// Verificar que se actualizó correctamente
	canalRecuperado, err := canalDAO.BuscarPorID(id)
	if err != nil {
		t.Fatalf("Error recuperando canal actualizado: %v", err)
	}

	if canalRecuperado == nil {
		t.Fatal("No se encontró el canal actualizado")
	}

	if canalRecuperado.Nombre() != nuevoNombre {
		t.Errorf("Nombre no actualizado. Esperado: %s, Obtenido: %s", nuevoNombre, canalRecuperado.Nombre())
	}

	if canalRecuperado.Descripcion() != nuevaDescripcion {
		t.Errorf("Descripción no actualizada. Esperado: %s, Obtenido: %s", nuevaDescripcion, canalRecuperado.Descripcion())
	}

	if canalRecuperado.Tipo() != nuevoTipo {
		t.Errorf("Tipo no actualizado. Esperado: %s, Obtenido: %s", nuevoTipo, canalRecuperado.Tipo())
	}

	// Limpiar datos de prueba
	_, err = dbPool.ExecContext(
		context.Background(),
		"DELETE FROM canales WHERE id = ?",
		id.String(),
	)
	if err != nil {
		t.Logf("Error al limpiar canal de prueba: %v", err)
	}
}

// TestCanalDAO_Eliminar verifica que el método Eliminar funcione correctamente
func TestCanalDAO_Eliminar(t *testing.T) {
	// Inicializar conexión a la base de datos
	dbPool := setupTestDB(t)
	defer cleanupTestDB(t, dbPool)

	// Crear instancia del DAO
	canalDAO := NuevoCanalDAO(dbPool)

	// Crear canal de prueba
	id := uuid.New()
	nombre := "Canal para Eliminar"
	descripcion := "Canal que será eliminado"
	tipo := model.CanalPublico

	canal, err := model.NewCanalServidor(id, nombre, descripcion, tipo)
	if err != nil {
		t.Fatalf("Error creando objeto canal: %v", err)
	}

	// Guardarlo en la base de datos
	err = canalDAO.Crear(canal)
	if err != nil {
		t.Fatalf("Error creando canal en la base de datos: %v", err)
	}

	// Verificar que se creó correctamente
	canalVerificacion, err := canalDAO.BuscarPorID(id)
	if err != nil {
		t.Fatalf("Error verificando creación de canal: %v", err)
	}
	if canalVerificacion == nil {
		t.Fatal("El canal no se creó correctamente")
	}

	// Ejecutar operación de eliminación
	err = canalDAO.Eliminar(id)
	if err != nil {
		t.Fatalf("Error eliminando canal: %v", err)
	}

	// Verificar que se eliminó correctamente
	canalEliminado, err := canalDAO.BuscarPorID(id)
	if err != nil {
		t.Fatalf("Error consultando canal eliminado: %v", err)
	}

	if canalEliminado != nil {
		t.Fatal("El canal no se eliminó correctamente")
	}
}
