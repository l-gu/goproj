package main

import (
	"net/http"

	"github.com/l-gu/goproject/internal/app/web/mvc"
)

func initWebMVC() {
	initWebMVCTemplates()
	initWebMVCControllers()
}

func initWebMVCTemplates() {
	log("Templates initialization - Start")
	mvc.InitTemplates()
	log("Templates initialization - End")
}

func initWebMVCControllers() {

	log("Controllers initialization - Start")

	mvc.LogFlag = true

	// Specific Paths with specific controllers

	languageController := mvc.NewLanguageController()
	http.HandleFunc("/language/", languageController.HttpHandler) // "/language/*"

	studentController := mvc.NewStudentController()
	http.HandleFunc("/student/", studentController.HttpHandler) // "/student/*"

	foobarController := mvc.NewFooBarController()
	http.HandleFunc("/foobar/", foobarController.HttpHandler) // "/foobar/*"

	log("Controllers initialization - End")
}
