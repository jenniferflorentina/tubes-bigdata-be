package controllers

import (
	"TubesBigData/dto"
	e "TubesBigData/err"
	"TubesBigData/mapper"
	"TubesBigData/model"
	"TubesBigData/response"
	"TubesBigData/services"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/dealancer/validate.v2"
)

func AddTodo(ctx *fiber.Ctx) error {

	createDto := new(dto.CreateTodoDTO)
	err := ctx.BodyParser(createDto)
	if err != nil {
		e.HandleErr(ctx, err)
		return nil
	}

	err = validate.Validate(&createDto)
	if err != nil {
		e.HandleErr(ctx, err)
		return nil
	}

	var toDo model.ToDo
	mapper.Map(createDto, &toDo)

	insertResult, err := services.AddTodo(&toDo)

	if err != nil {
		e.HandleErr(ctx, err)
		return nil
	}
	_ = ctx.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: insertResult,
	})
	return nil
}

type DeleteItem struct {
	ListToDelete []string `query:"listToDelete"`
}

func DeleteMultiple(ctx *fiber.Ctx) error {
	p := new(DeleteItem)
	if err := ctx.QueryParser(p); err != nil {
		log.Fatalln("DeleteMany receive list:", err)
		e.HandleErr(ctx, err)
		return nil
	}
	result, err := services.DeleteMultiple(p.ListToDelete)

	if err != nil {
		e.HandleErr(ctx, err)
		return nil
	}
	_ = ctx.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: result,
	})
	return nil
}

func DeleteOne(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	result, err := services.DeleteOne(id)

	if err != nil {
		e.HandleErr(ctx, err)
		return nil
	}
	_ = ctx.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: result,
	})
	return nil
}

func GetOne(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	result, err := services.GetOne(id)

	if err != nil {
		e.HandleErr(ctx, err)
		return nil
	}
	_ = ctx.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: result,
	})
	return nil
}

func UpdateOne(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	createDto := new(dto.CreateTodoDTO)
	err := ctx.BodyParser(createDto)
	if err != nil {
		e.HandleErr(ctx, err)
		return nil
	}

	err = validate.Validate(&createDto)
	if err != nil {
		e.HandleErr(ctx, err)
		return nil
	}

	var toDo model.ToDo
	mapper.Map(createDto, &toDo)

	updateResult, err := services.UpdateOne(id, toDo)

	if err != nil {
		e.HandleErr(ctx, err)
		return nil
	}
	_ = ctx.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: updateResult,
	})
	return nil
}

func GetAll(ctx *fiber.Ctx) error {
	result, err := services.GetAll()

	if err != nil {
		e.HandleErr(ctx, err)
		return nil
	}
	_ = ctx.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: result,
	})
	return nil
}
