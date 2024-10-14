package javaui

import (
	"fmt"
	"reflect"

	"github.com/rivo/tview"
)

type LanguageInfo struct {
	El *tview.Table
}

// Convert a slice of structs to a slice of interface{}
func convertToInterfaces[T any](elems []T) []interface{} {
	interfaces := make([]interface{}, len(elems))
	for i, v := range elems {
		interfaces[i] = v
	}
	return interfaces
}

// Generic function to initialize the table
func (li *LanguageInfo) init(elems []any) *tview.Table {
	li.El = tview.NewTable()
	li.El.
		SetFixed(1, 1).
		SetSelectable(true, false).
		SetSeparator(tview.Borders.Vertical).SetBorderPadding(1, 1, 2, 0).
		SetBorder(true).SetTitle("Language Information")

	// Populate the table with the provided attributes and objects
	li.UpdateTable(elems)

	return li.El
}

// Function to update the table data with inferred attributes (column names) and objects
func (li *LanguageInfo) UpdateTable(objects []any) {
	// Clear the current table
	li.El.Clear()

	// Check if there are objects to infer from
	if len(objects) == 0 {
		return
	}

	// Infer the attributes (field names) from the first object
	val := reflect.ValueOf(objects[0])
	if val.Kind() != reflect.Struct {
		return // If the object is not a struct, we can't infer the attributes
	}

	var attributes []string

	// Loop through fields of the struct to get the field names based on the existence of the "table" tag
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		if _, ok := field.Tag.Lookup("table"); ok { // Check for the presence of the "table" tag
			attributes = append(attributes, field.Name) // Include the field if the tag is present
		}
	}

	// Add header row (attributes)
	for col, attr := range attributes {
		li.El.SetCell(0, col, tview.NewTableCell(attr).
			SetAlign(tview.AlignCenter).
			SetTextColor(tview.Styles.SecondaryTextColor).
			SetSelectable(false)) // Set header to bold
	}

	// Add rows for each object
	for rowIndex, obj := range objects {
		val := reflect.ValueOf(obj)

		// Handle struct types
		if val.Kind() == reflect.Struct {
			for colIndex, attr := range attributes {
				field := val.FieldByName(attr)

				if !field.IsValid() {
					li.El.SetCell(rowIndex+2, colIndex, tview.NewTableCell("N/A"))
					continue
				}

				// Convert the field value to string and insert it into the table
				li.El.SetCell(rowIndex+2, colIndex, tview.NewTableCell(fieldToString(field)))
			}
		}
	}
}

// Helper function to convert a reflect.Value to a string for table display
func fieldToString(field reflect.Value) string {
	switch field.Kind() {
	case reflect.String:
		return "  " + field.String() + "  "
	case reflect.Int, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d  ", field.Int())
	case reflect.Bool:
		if field.Bool() {
			return "True"
		}
		return "False"
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%.2f  ", field.Float())
	default:
		return "N/A"
	}
}
