package Creational

import "fmt"

const (
	SpreadSheetIdentifier = 2
	DocumentIdentifier    = 1
)

type GoogleApplicationCreator interface {
	CreateFile()
	GetFile() File
}

type DocumentApplicationCreator struct {
	File *File
}

type SpreadSheetApplicationCreator struct {
	File *File
}

type File struct {
	Name string
	Type string
	Size string
}

func (application *DocumentApplicationCreator) CreateFile() {
	application.File = new(File)
	application.File.Name = "Document"
	application.File.Type = "doc"
	application.File.Size = "10MB"
}

func (application *DocumentApplicationCreator) GetFile() File {
	return *application.File
}

func (application *SpreadSheetApplicationCreator) CreateFile() {
	application.File = new(File)
	application.File.Name = "Spreadsheet"
	application.File.Type = "xls"
	application.File.Size = "10MB"
}

func (application *SpreadSheetApplicationCreator) GetFile() File {
	return *application.File
}

func (file *File) GetMeta() string {
	return fmt.Sprintf("File details are : File Name %v, File Type %v, File Size %v", file.Name, file.Type, file.Size)
}

func GetGoogleApplicationCreator(id int) GoogleApplicationCreator {
	switch id {
	case 1:
		return new(DocumentApplicationCreator)
	case 2:
		return new(SpreadSheetApplicationCreator)
	default:
		return nil
	}
}
