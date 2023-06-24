package database

import (
	"Capstone/models"
	"context"
)

func CreateReport(ctx context.Context, Report models.Report) (models.Report, error) {
	err := DB.WithContext(ctx).Create(&Report).Error
	if err != nil {
		return models.Report{}, err
	}

	// Preload user data for the created Report
	err = DB.WithContext(ctx).Preload("ReportThread").Preload("ReportUser").First(&Report).Error
	if err != nil {
		return models.Report{}, err
	}

	return Report, nil
}

func DeleteReports(ctx context.Context, id int) error {
	var Report models.Report

	result := DB.WithContext(ctx).Where("id = ?", id).Delete(&Report)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrIDNotFound
	}

	return nil
}
func GetReportByID(ctx context.Context, id int) (models.Report, error) {
	var Report models.Report

	err := DB.WithContext(ctx).Preload("ReportThread").Preload("ReportUser").Where("id = ?", id).Find(&Report).Error
	if err != nil {
		return models.Report{}, err
	}

	return Report, nil
}

func GetReports(ctx context.Context) ([]models.Report, error) {
	var Report []models.Report

	err := DB.WithContext(ctx).Preload("ReportUser").Preload("ReportThread").Find(&Report).Error
	if err != nil {
		return nil, err
	}

	return Report, nil
}
