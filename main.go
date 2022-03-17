package main

import (
	"DesignPatterns/Structural"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

func main() {
	e := echo.New()
	// Creational : Abstract Factory Pattern
	//dataAccessFactory := Creational.ReturnDataFactory(Creational.MongoFactory)
	//dataAccessFactory.CreateReadStore()
	//dataAccessFactory.CreateWriteStore()
	//
	//log.Info(dataAccessFactory.GetReadStore().GetReadPath())
	//log.Info(dataAccessFactory.GetWriteStore().GetWritePath())

	// Creational : Builder Pattern
	//builderDirector := Creational.Director{}
	//builderDirector.SetBuilder(Creational.MonolithProcess)
	//builderDirector.Builder.SetAsyncQueue("kafka").SetDatabaseType("SQL").SetFramework("Spring")
	//
	//details := builderDirector.Builder.GetApplicationDetails()
	//log.Info(details)

	//Creational : Factory Pattern
	//applicationCreator := Creational.GetGoogleApplicationCreator(Creational.DocumentIdentifier)
	//applicationCreator.CreateFile()
	//file := applicationCreator.GetFile()
	//log.Info(file.GetMeta())

	// Creational : Prototype Pattern
	//clonnablePhone := new(Creational.ClonnablePhone)
	//clonnablePhone.Name = "One Plus 7"
	//clonnablePhone.Brand = "One+"
	//clonnablePhone.Color = "Blue"
	//
	//clonedPhone := clonnablePhone.Clone().(*Creational.ClonnablePhone)
	//clonnablePhone.Name = "One Plus 8"
	//log.Info(clonedPhone)

	// Creational : Singleton Pattern
	//singletonPattern := Creational.GetDBConnection()
	//log.Info(singletonPattern)

	// Structural : Composite Pattern
	//lineComponent := new(Structural.Line)
	//lineComponent.Create()
	//
	//dotComponent := new(Structural.Dot)
	//dotComponent.Create()
	//
	//pictureComponent := new(Structural.Picture)
	//pictureComponent.Create()
	//
	//parentPictureComponent := new(Structural.Picture)
	//parentPictureComponent.Create()
	//
	//pictureComponent.AddComponent(lineComponent)
	//pictureComponent.AddComponent(dotComponent)
	//
	//parentPictureComponent.AddComponent(pictureComponent)
	//
	//parentPictureComponent.ShowAllComponents()

	// Structural : Adapter Pattern
	//data := "DataToPrint"
	//legacyPrinter := Structural.OldPrinter{}
	//legacyPrinter.Print(data)
	//printerAdapter := Structural.PrinterAdapter{
	//	&legacyPrinter,
	//	data,
	//}
	//printerAdapter.PrintStored()

	// Structural: Bridge Pattern
	//Structural.PublishNotification("Push")
	//Structural.PublishNotification("Text")

	// Structural : Proxy Pattern
	proxyService := Structural.ProxyService()
	log.Info(proxyService.Search("1"))
	log.Info(proxyService.Search("2"))

	//MeetingScheduler.TestEvent()
	//CarRentalService.CarRentalService()

	// VendingMachineDesign
	//VendingMachineDesign.VendingMachineService()
	e.GET("/", func(c echo.Context) error {
		log.Infof("heck it")
		return c.HTML(http.StatusOK, "This application is up!")
	})

	e.Logger.Fatal(e.Start(":32007"))
}
