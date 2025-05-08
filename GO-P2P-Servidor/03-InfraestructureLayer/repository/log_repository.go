package repository

import (
	"context"
	"database/sql"
	"pool"
	"time"

	"github.com/google/uuid"
	"model"
)

// LogRepository implementa ILogRepository
type LogRepository struct {
	dbPool *pool.DBConnectionPool
}

// NewLogRepository crea una nueva instancia del repositorio
func NewLogRepository(dbPool *pool.DBConnectionPool) *LogRepository {
	return &LogRepository{dbPool: dbPool}
}

// Save guarda un registro de log
func (r *LogRepository) Save(ctx context.Context, entry *model.LogEntry) error {
	query := `INSERT INTO logs (id, user_id, tipo_evento, descripcion, fecha)
              VALUES (?, ?, ?, ?, ?)`

	_, err := r.dbPool.ExecContext(ctx, query,
		entry.ID().String(),
		entry.UserID().String(),
		string(entry.EventType()),
		entry.Description(),
		entry.Date())

	return err
}

// Delete elimina un registro de log
func (r *LogRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM logs WHERE id = ?`
	_, err := r.dbPool.ExecContext(ctx, query, id.String())
	return err
}

// FindByID busca un registro de log por ID
func (r *LogRepository) FindByID(ctx context.Context, id uuid.UUID) (*model.LogEntry, error) {
	query := `SELECT id, user_id, tipo_evento, descripcion, fecha
              FROM logs WHERE id = ?`

	row := r.dbPool.QueryRowContext(ctx, query, id.String())

	var idStr, userIDStr, tipoEventoStr, descripcion string
	var fecha time.Time

	if err := row.Scan(&idStr, &userIDStr, &tipoEventoStr, &descripcion, &fecha); err != nil {
		return nil, err
	}

	logID, err := uuid.Parse(idStr)
	if err != nil {
		return nil, err
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, err
	}

	tipoEvento := model.EventoTipo(tipoEventoStr)

	return model.NewLogEntry(logID, userID, tipoEvento, descripcion, fecha)
}

// ListAll lista todos los registros de log
func (r *LogRepository) ListAll(ctx context.Context) ([]*model.LogEntry, error) {
	query := `SELECT id, user_id, tipo_evento, descripcion, fecha FROM logs`
	return r.executeScanQuery(ctx, query)
}

// ListByUser lista los registros de log de un usuario específico
func (r *LogRepository) ListByUser(ctx context.Context, userID uuid.UUID) ([]*model.LogEntry, error) {
	query := `SELECT id, user_id, tipo_evento, descripcion, fecha 
              FROM logs WHERE user_id = ?`

	return r.executeScanQueryWithParam(ctx, query, userID.String())
}

// ListByType lista los registros de log de un tipo específico
func (r *LogRepository) ListByType(ctx context.Context, t model.EventoTipo) ([]*model.LogEntry, error) {
	query := `SELECT id, user_id, tipo_evento, descripcion, fecha 
              FROM logs WHERE tipo_evento = ?`

	return r.executeScanQueryWithParam(ctx, query, string(t))
}

// ListByDateRange lista los registros de log en un rango de fechas
func (r *LogRepository) ListByDateRange(ctx context.Context, from, to time.Time) ([]*model.LogEntry, error) {
	query := `SELECT id, user_id, tipo_evento, descripcion, fecha 
              FROM logs WHERE fecha BETWEEN ? AND ?`

	rows, err := r.dbPool.QueryContext(ctx, query, from, to)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanRows(rows)
}

// Método auxiliar para ejecutar consultas y escanear resultados
func (r *LogRepository) executeScanQuery(ctx context.Context, query string) ([]*model.LogEntry, error) {
	rows, err := r.dbPool.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanRows(rows)
}

// Método auxiliar para ejecutar consultas con un parámetro y escanear resultados
func (r *LogRepository) executeScanQueryWithParam(ctx context.Context, query string, param string) ([]*model.LogEntry, error) {
	rows, err := r.dbPool.QueryContext(ctx, query, param)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanRows(rows)
}

// Método auxiliar para escanear filas de resultados
func (r *LogRepository) scanRows(rows *sql.Rows) ([]*model.LogEntry, error) {
	var logs []*model.LogEntry

	for rows.Next() {
		var idStr, userIDStr, tipoEventoStr, descripcion string
		var fecha time.Time

		if err := rows.Scan(&idStr, &userIDStr, &tipoEventoStr, &descripcion, &fecha); err != nil {
			return nil, err
		}

		logID, err := uuid.Parse(idStr)
		if err != nil {
			return nil, err
		}

		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			return nil, err
		}

		tipoEvento := model.EventoTipo(tipoEventoStr)

		logEntry, err := model.NewLogEntry(logID, userID, tipoEvento, descripcion, fecha)
		if err != nil {
			return nil, err
		}

		logs = append(logs, logEntry)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return logs, nil
}
