package controller

import (
	"GDSC-PROJECT/controller/validation"
	"GDSC-PROJECT/database"
	"GDSC-PROJECT/models/entity"
	"bytes"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jung-kurt/gofpdf"
)

func GetAllWarehouse(ctx *fiber.Ctx) error {
	var warehouses []entity.Warehouse

	result := database.DB.Preload("Stocks").Preload("Stocks.Product").Preload("Stocks.Product.Category").Find(&warehouses)
	if err := validation.QueryResultValidation(result, "warehouse"); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"warehouses": warehouses,
	})
}

func GetWarehouseByID(ctx *fiber.Ctx) error {
	warehouseID, err := validation.ParseAndIDValidation(ctx, "id", "warehouse")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var warehouse entity.Warehouse
	result := database.DB.Preload("Stocks").Preload("Stocks.Product").Preload("Stocks.Product.Category").First(&warehouse, warehouseID)
	if err = validation.EntityByIDValidation(result, "warehouse"); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"warehouse": warehouse,
	})
}

func GetWarehouseByProductID(ctx *fiber.Ctx) error {
	productID, err := validation.ParseAndIDValidation(ctx, "product_id", "product")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var warehouses []entity.Warehouse
	result := database.DB.Preload("Stocks").Preload("Stocks.Product").Preload("Stocks.Product.Category").Joins("JOIN product_warehouses ON warehouses.id = product_warehouses.warehouse_id").Where("product_id = ?", productID).Find(&warehouses)
	if err = validation.EntityByIDValidation(result, "warehouse"); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"warehouses": warehouses,
	})
}

func CreateWarehouse(ctx *fiber.Ctx) error {
	warehouse := new(entity.Warehouse)

	if err := ctx.BodyParser(warehouse); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := validation.WarehouseValidation(warehouse); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	database.DB.Debug().Create(&warehouse)

	return ctx.JSON(fiber.Map{
		"message":   "created successfully",
		"warehouse": warehouse,
	})
}

func UpdateWarehouse(ctx *fiber.Ctx) error {
	requestWarehouse := new(entity.WarehouseResponse)
	warehouse := new(entity.Warehouse)

	warehouseID, err := validation.ParseAndIDValidation(ctx, "id", "warehouse")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	result := database.DB.First(&warehouse, warehouseID)
	if err = validation.EntityByIDValidation(result, "warehouse"); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := ctx.BodyParser(requestWarehouse); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := validation.WarehouseUpdateValidation(requestWarehouse, warehouse); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	errUpdate := database.DB.Debug().Save(&warehouse).Error
	if errUpdate != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errUpdate.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message":   "updated successfully",
		"warehouse": warehouse,
	})
}

func DeleteWarehouse(ctx *fiber.Ctx) error {
	warehouse := new(entity.Warehouse)

	warehouseID, err := validation.ParseAndIDValidation(ctx, "id", "warehouse")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	result := database.DB.First(&warehouse, warehouseID)
	if err = validation.EntityByIDValidation(result, "warehouse"); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	errUpdate := database.DB.Debug().Delete(&warehouse).Error
	if errUpdate != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errUpdate.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message":   "delete successfully",
		"warehouse": warehouse,
	})
}

func GetWarehouseReport(ctx *fiber.Ctx) error {

	warehouseID, err := validation.ParseAndIDValidation(ctx, "id", "warehouse")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var warehouse entity.Warehouse
	result := database.DB.Preload("Stocks.Product.Category").First(&warehouse, warehouseID)
	if err = validation.EntityByIDValidation(result, "warehouse"); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 12)

	pdf.MultiCell(0, 12, "WAREHOUSE REPORT", "", "C", false)
	pdf.Ln(10)
	pdf.SetFont("Arial", "", 12)
	info := map[string]string{
		"Warehouse Name": warehouse.Name,
		"Warehouse ID":   strconv.FormatUint(uint64(warehouse.ID), 10),
		"Location":       warehouse.Location,
		"Capacity":       strconv.FormatUint(uint64(warehouse.Capacity), 10),
	}

	for key, value := range info {
		pdf.CellFormat(40, 10, key, "", 0, "L", false, 0, "")
		pdf.CellFormat(0, 10, fmt.Sprintf(": %s", value), "", 1, "L", false, 0, "")
	}
	pdf.Ln(10)

	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(40, 10, "Stock ID", "1", 0, "C", false, 0, "")
	pdf.CellFormat(40, 10, "Product Name", "1", 0, "C", false, 0, "")
	pdf.CellFormat(40, 10, "Category", "1", 0, "C", false, 0, "")
	pdf.CellFormat(40, 10, "Quantity", "1", 0, "C", false, 0, "")
	pdf.Ln(-1)

	pdf.SetFont("Arial", "", 12)
	for _, stock := range warehouse.Stocks {
		pdf.CellFormat(40, 10, fmt.Sprintf("%d", stock.ID), "1", 0, "", false, 0, "")
		pdf.CellFormat(40, 10, stock.Product.Name, "1", 0, "", false, 0, "")
		pdf.CellFormat(40, 10, stock.Product.Category.Name, "1", 0, "", false, 0, "")
		pdf.CellFormat(40, 10, fmt.Sprintf("%d", stock.Quantity), "1", 0, "", false, 0, "")
		pdf.Ln(-1)
	}

	var buf bytes.Buffer
	pdfError := pdf.Output(&buf)
	if pdfError != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": pdfError.Error(),
		})
	}

	ctx.Set("Content-Type", "application/pdf")
	ctx.Set("Content-Disposition", "attachment; filename=warehouse_report.pdf")
	return ctx.Send(buf.Bytes())
}
