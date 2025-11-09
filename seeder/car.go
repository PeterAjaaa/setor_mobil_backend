package seeder

import "github.com/PeterAjaaa/setor_mobil_backend/models"

var carSeed = []models.Cars{
	{
		RegistrationNum: "D 1234 EOF",
		Brand:           "Toyota",
		Model:           "Avanza",
		Year:            2025,
		PricePerDay:     750000,
		Status:          "Available",
		Description:     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis nec nulla leo. Fusce sed nibh ipsum. Nunc id ex ut ex suscipit laoreet id sit amet mi. Cras lobortis sodales ligula, ut fermentum justo faucibus vitae. Suspendisse potenti. Nunc tempor felis non commodo auctor. Vivamus congue justo non nisi vehicula interdum.",
		ImageURL:        "https://picsum.photos/seed/avanza/800/600",
	},
	{
		RegistrationNum: "D 5678 EOF",
		Brand:           "Honda",
		Model:           "Civic",
		Year:            2025,
		PricePerDay:     15000000,
		Status:          "Rented",
		Description:     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi cursus leo mollis augue aliquam, quis hendrerit nunc vehicula. Sed finibus est sit amet gravida consectetur. Phasellus convallis, purus eget luctus posuere, nibh felis suscipit urna, aliquet congue massa neque non massa. Donec vitae lacus sed elit bibendum sollicitudin. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut finibus congue libero. Fusce sed efficitur mauris. Phasellus gravida orci vel aliquet mollis. Donec sit amet maximus purus. Curabitur eget libero condimentum, consectetur ante sit amet, fringilla erat. Mauris a sapien ut ipsum posuere venenatis. Aenean vitae urna volutpat, commodo nibh et, viverra augue.",
		ImageURL:        "https://picsum.photos/seed/civic/800/600",
	},
}
