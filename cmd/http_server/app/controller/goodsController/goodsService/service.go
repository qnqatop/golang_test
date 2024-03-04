package goodsService

import (
	"context"
	"encoding/json"
	"golang_testovoe/cmd/http_server/app/controller/goodsController/goodsService/dto"
	"golang_testovoe/cmd/http_server/config/initModules"
	"strconv"
	"time"
)

type GoodsService struct {
}

func (g *GoodsService) List() {
}

func (g *GoodsService) Create(request dto.GoodsCreateRequest, projectId string) (dto.GoodsResponse, error) {
	var res dto.GoodsResponse

	priority, err := g.getLastPriority(projectId)
	newPriority := priority + 1

	insertRowToGoods := `
    INSERT INTO goods (name, project_id, description, removed, created_at, priority)
    VALUES ($1, $2, $3, $4, $5, $6)
    RETURNING id, project_id, name, description, priority, removed, created_at;
`
	row := initModules.ApplicationConfig.GetPostgresDb().QueryRow(insertRowToGoods, request.Name, projectId, "", false, time.Now(), newPriority)

	err = row.Scan(
		&res.ID,
		&res.ProjectID,
		&res.Name,
		&res.Description,
		&res.Priority,
		&res.Removed,
		&res.CreateAt,
	)
	if err != nil {
		return res, err
	}
	err = g.redisSet(res)
	if err != nil {
		return dto.GoodsResponse{}, err
	}

	return res, nil
}

func (g *GoodsService) Update(request dto.GoodsUpdateRequest, projectId, Id string) (dto.GoodsResponse, error) {
	var res dto.GoodsResponse

	tx, err := initModules.ApplicationConfig.GetPostgresDb().Begin()
	if err != nil {
		return res, err
	}
	defer tx.Rollback()

	updateGoodsQuery := `
        UPDATE goods SET 
            name = $1,
            description = $2
        WHERE id = $3 AND project_id = $4;
    `
	_, err = tx.Exec(updateGoodsQuery, request.Name, request.Description, Id, projectId)
	if err != nil {
		return res, err
	}

	selectGoodsQuery := `
        SELECT id, project_id, name, description, priority, removed, created_at
        FROM goods
        WHERE id = $1 AND project_id = $2;
    `
	row := tx.QueryRow(selectGoodsQuery, Id, projectId)
	err = row.Scan(
		&res.ID,
		&res.ProjectID,
		&res.Name,
		&res.Description,
		&res.Priority,
		&res.Removed,
		&res.CreateAt,
	)
	if err != nil {
		return res, err
	}

	err = tx.Commit()
	if err != nil {
		return res, err
	}

	err = g.redisSet(res)
	if err != nil {
		return dto.GoodsResponse{}, err
	}

	return res, nil
}

func (g *GoodsService) Delete(id, projectId string) (dto.GoodsDeleteResponse, error) {
	var res dto.GoodsDeleteResponse

	tx, err := initModules.ApplicationConfig.GetPostgresDb().Begin()
	if err != nil {
		return res, err
	}
	defer tx.Rollback()

	deleteGoodsQuery := `
        DELETE FROM goods 
        WHERE id = $1 ;
    `

	_, err = tx.Exec(deleteGoodsQuery, id)
	if err != nil {
		return res, err
	}

	err = tx.Commit()
	if err != nil {
		return res, err
	}

	res.ID, _ = strconv.Atoi(id)
	res.ProjectId, _ = strconv.Atoi(projectId)
	res.Removed = true

	g.redisDel(id)

	return res, nil
}

func (g *GoodsService) getLastPriority(projectId string) (int, error) {

	getHighestPriorityQuery := `
    SELECT priority FROM goods
    WHERE project_id = $1
    ORDER BY priority DESC
    LIMIT 1;
    `

	var highestPriority int
	err := initModules.ApplicationConfig.GetPostgresDb().QueryRow(getHighestPriorityQuery, projectId).Scan(&highestPriority)
	if err != nil {
		return 0, err
	}
	return highestPriority, nil
}

func (g *GoodsService) redisSet(res dto.GoodsResponse) error {
	goodsJSON, err := json.Marshal(res)
	if err != nil {
		return err
	}
	initModules.ApplicationConfig.GetRDb().Set(context.Background(), strconv.Itoa(res.ID), string(goodsJSON), 0)
	return nil
}

func (g *GoodsService) redisDel(id string) {
	initModules.ApplicationConfig.GetRDb().Del(context.Background(), id)
}

func (g *GoodsService) GetValueRedis(id string) (string, error) {
	return initModules.ApplicationConfig.GetRDb().Get(context.Background(), id).Result()
}
