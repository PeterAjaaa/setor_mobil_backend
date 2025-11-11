package seeder

import "github.com/PeterAjaaa/setor_mobil_backend/models"

var motorcycleSeed = []models.Motorcycles{
	{
		RegistrationNum: "B 7890 XYZ",
		Brand:           "Kawasaki",
		Model:           "Ninja",
		Year:            2025,
		PricePerDay:     250000,
		Status:          "Available",
		Description:     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse ac elit finibus, pulvinar neque sed, tempor magna. Vestibulum nec nulla molestie, tristique tortor at, consequat orci. Fusce rutrum efficitur nulla, vel efficitur nibh sollicitudin in. Aenean quis ex id nunc facilisis mollis vitae quis quam. Vivamus vestibulum ipsum nec aliquam lacinia. Donec ut consectetur odio. Integer pretium varius lectus, vitae porta mauris scelerisque consectetur. Suspendisse ac laoreet ligula, sed hendrerit nibh. Vestibulum convallis augue sed egestas semper. Fusce cursus quam lectus, non condimentum lacus eleifend sit amet. Sed et lorem leo. Aliquam id mauris vitae libero consequat elementum.",
		ImageURL:        "https://picsum.photos/seed/ninja/800/600",
		CreatedByID:     1,
	},
	{
		RegistrationNum: "D 5678 EOF",
		Brand:           "Honda",
		Model:           "Supra",
		Year:            2025,
		PricePerDay:     100000,
		Status:          "Rented",
		Description:     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla dignissim est justo. Integer nec diam a enim eleifend dignissim eget in enim. Aenean pellentesque pulvinar lectus, vel efficitur libero hendrerit vitae. Donec sed dolor nec ante aliquam molestie. In hac habitasse platea dictumst. Ut diam metus, fermentum nec ex sit amet, ultricies faucibus elit. Praesent eu facilisis enim. In ultrices libero enim, quis suscipit purus pulvinar nec. Ut ac elementum dui, sit amet porta lorem. Duis rutrum ultrices placerat. Sed porta vel turpis sit amet pharetra. Nunc venenatis elementum maximus.",
		ImageURL:        "https://picsum.photos/seed/supra/800/600",
		CreatedByID:     2,
	},
}
