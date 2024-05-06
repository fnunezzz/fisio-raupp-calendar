package service

import (
	"fmt"
	"time"

	"github.com/xuri/excelize/v2"
)

const (
	sheet_name = "Sheet1"
    alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	min_x_coordinate = 5 // E
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
		Horizontal: "center",
		Vertical:   "center",
	},
}

type XlsxService interface {
	GenerateXlsxReport([]Input, time.Time) error
}

type xlsxService struct{}

func NewXlsxService() XlsxService {
	return &xlsxService{}
}




func (s *xlsxService) GenerateXlsxReport(input []Input, date time.Time) error {
	var (
		y_axis int = 1
		startHour int = 8
		startMinute int = 0
	)


	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// CREATING THE HEADER
	formattedDate := date.Format("02/01/2006")

	f.SetCellValue(sheet_name, fmt.Sprintf("A%d", y_axis), formattedDate)
	minCellName, err := excelize.CoordinatesToCellName(min_x_coordinate, y_axis)
	if (err != nil) {
		return err
	}

	err = f.MergeCell("Sheet1", "A1", minCellName)

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
			// I want to start writing only when the condition below is true. It's Column A and not the Header
			if columnIndex == 0 && rowIndex > 0 {
				for _, t := range input {
					if t.Time != rowCell {
						continue
					}
					patientsToWrite = append(patientsToWrite, t.Text)
					if len(patientsToWrite) == 0 {
						break
					}
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
		fmt.Println(err)
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

	x_axis := min_x_coordinate
	y_axis := len(rows)

	// Merging first cell to last cell
	// If x_axis is bigger than min_x_coordinate - it means that I need to increase the size of the merge cell
	if x_axis > min_x_coordinate {
		x_axis = len(cols)
		x_merge_coordinate, err := excelize.CoordinatesToCellName(x_axis, 1)
		if err != nil {
			return err
		}
	
		err = f.MergeCell("Sheet1", "A1", x_merge_coordinate)
	
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

	err = f.SetCellStyle("Sheet1", "A1", lastCellName, style)

	if err != nil {
		return err
	}

	return nil

}