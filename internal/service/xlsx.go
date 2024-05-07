package service

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/xuri/excelize/v2"
)

const (
	sheet_name = "Sheet1"
    alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	min_x_coordinate = 5 // E = timesheet | patient | patient | patient | patient
	column_width = 25
	row_height = 25
)

type Input struct {
	Text string
	Time string
}

var STYLE_PROPERTIES = &excelize.Style{
	Border: []excelize.Border{
		{
			Type:  "left",
			Color: "000000",
			Style: 1,
		},
		{
			Type:  "right",
			Color: "000000",
			Style: 1,
		},
		{
			Type:  "top",
			Color: "000000",
			Style: 1,
		},
		{
			Type:  "bottom",
			Color: "000000",
			Style: 1,
		},
	},
	Alignment: &excelize.Alignment{
		Horizontal: "left",
		Vertical:   "center",
	},
	Font: &excelize.Font{
		Size: 10,
	},
}

type XlsxService interface {
	GenerateXlsxReport([]Input, time.Time) error
}

type xlsxService struct{}

func NewXlsxService() XlsxService {
	return &xlsxService{}
}



// Documentation: https://xuri.me/excelize/en
func (s *xlsxService) GenerateXlsxReport(input []Input, date time.Time) error {
	var (
		y_axis int = 1
		startHour int = 8
		startMinute int = 0
	)


	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln(err.Error())
		}
	}()

	// CREATING THE HEADER
	formattedDate := date.Format("02/01/2006")

	f.SetCellValue(sheet_name, fmt.Sprintf("A%d", y_axis), formattedDate)
	minCellName, err := excelize.CoordinatesToCellName(min_x_coordinate, y_axis)
	if (err != nil) {
		return err
	}

	err = f.MergeCell(sheet_name, "A1", minCellName)

	if err != nil {
		return err
	}
	y_axis++
	// END CREATING THE HEADER

	// GENERATING TIMESLOTS
	for {
		// if startHour is over 18 - break the loop as it's the last possible appointment
		if startHour >= 18 {
			f.SetCellValue(sheet_name, fmt.Sprintf("A%d", y_axis), fmt.Sprintf("%02d:00", startHour))
			break
		}
		// if startHour is 11 it's a full hour - no 20 minute increment
		if startHour == 11 {
			f.SetCellValue(sheet_name, fmt.Sprintf("A%d", y_axis), fmt.Sprintf("%02d:00", startHour))
			startMinute = 0
			startHour++
			y_axis++
			continue
		}
		// if startHour is 12 - it's break time. That means it's empty
		if startHour == 12 {
			f.SetCellValue(sheet_name, fmt.Sprintf("A%d", y_axis), "")
			startMinute = 0
			startHour++
			y_axis++
			continue
		}
		// if startMinute is over 40 - reset it and increment startHour as it will be a full hour done
		if startMinute > 40 {
			startMinute = 0
			startHour++
			continue
		}
		f.SetCellValue(sheet_name, fmt.Sprintf("A%d", y_axis), fmt.Sprintf("%02d:%02d", startHour, startMinute))
		y_axis++

		startMinute += 20

	} // END GENERATING TIMESLOTS


	// START INPUT PROCESSING
	y_axis = 2 // Start below the header

	cols, err := f.GetCols(sheet_name)
	if err != nil {
		return err
	}

	var patientsToWrite []string
	for columnIndex, col := range cols {
		patientsToWrite = []string{}

		for rowIndex, rowCell := range col {
			// columnIndex == 0 is A column
			// CellIndex == 0 is the first row (header)
			// I want to start writing only when the condition below is false - It must be Column A and not the Header
			if columnIndex != 0 && rowIndex == 0 {
				continue
			}

			for _, t := range input {
				if t.Time != rowCell {
					continue
				}
				patientsToWrite = append(patientsToWrite, t.Text)
				if len(patientsToWrite) == 0 {
					break
				}
			}

			for i, p := range patientsToWrite {
				// write input to the cell
				columnCoordinates := columnIndex + 2 + i
				rowCoordinates := rowIndex + 1
				cellName, err := excelize.CoordinatesToCellName(columnCoordinates, rowCoordinates)
				if err != nil {
					return err
				}
				f.SetCellValue(sheet_name, cellName, p)
			}
			patientsToWrite = []string{}
			

		}
	}

	// END INPUT PROCESSING

	// Apply style to the cells
	if err := s.applyStyle(f); err != nil {
		return err
	}

	fileName := fmt.Sprintf("Pacientes_%s.xlsx", date.Format("02-01-2006"))
	
	// Save workbook
	if err := f.SaveAs(fileName); err != nil {
		return err
	}

	return nil
}


// styling the cells
// it's very specific logic so I decided to keep it in a different function
func (s *xlsxService) applyStyle(f *excelize.File) error {
	// Getting latest worksheet sizes
	// Easir to understand if you think of it as a map with X coordinates and Y coordinates
	rows, err := f.GetRows(sheet_name)
	if err != nil {
		return err
	}

	cols, err := f.GetCols(sheet_name)
	if err != nil {
		return err
	}

	x_axis := len(cols)
	y_axis := len(rows)
	x_cell_name, err := excelize.CoordinatesToCellName(x_axis, 1)
	if err != nil {
		return err
	}
	// Merging first cell to last cell
	// If x_axis is bigger than min_x_coordinate - it means that I need to increase the size of the merge cell
	if x_axis > min_x_coordinate {

		err = f.MergeCell(sheet_name, "A1", x_cell_name)
	
		if err != nil {
			return err
		}
	} else {
		x_axis = min_x_coordinate
		x_cell_name, err = excelize.CoordinatesToCellName(x_axis, 1)
		if err != nil {
			return err
		}
	}

	// Column width
	regex := regexp.MustCompile(`\d+`)
	x_alpha := regex.ReplaceAllString(x_cell_name, "")
	err = f.SetColWidth(sheet_name, "B", x_alpha, column_width)
	if err != nil {
		return err
	}

	// column height
	for i := 1; i <= y_axis; i++ {
		err = f.SetRowHeight(sheet_name, i, row_height)
		if err != nil {
			return err
		}
	}

	lastCellName, err := excelize.CoordinatesToCellName(x_axis, y_axis)

	if err != nil {
		return err
	}

	style, err := f.NewStyle(STYLE_PROPERTIES)

	if err != nil {
		return err
	}

	err = f.SetCellStyle(sheet_name, "A1", lastCellName, style)

	if err != nil {
		return err
	}

	// next styles are all bold
	STYLE_PROPERTIES.Font.Bold = true
	// Centering header
	STYLE_PROPERTIES.Alignment.Horizontal = "center"
	style, err = f.NewStyle(STYLE_PROPERTIES)
	if err != nil {
		return err
	}
	
	err = f.SetCellStyle(sheet_name, "A1", x_cell_name, style)
	if err != nil {
		return err
	}
	
	// making timesheet bold
	STYLE_PROPERTIES.Alignment.Horizontal = "left"
	style, err = f.NewStyle(STYLE_PROPERTIES)
	if err != nil {
		return err
	}
	
	y_cell := fmt.Sprintf("A%d", y_axis)
	err = f.SetCellStyle(sheet_name, "A2", y_cell, style)
	if err != nil {
		return err
	}

	return nil

}