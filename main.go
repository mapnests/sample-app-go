package main

import (
	"context"
	"fmt"
	"log"
	"os"

	mapnests "github.com/mapnests/sdk-go"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("API_KEY")
	pkgName := os.Getenv("PACKAGE_NAME")

	if apiKey == "" || pkgName == "" {
		log.Fatal("Api key or Package name is empty")
	}

	mapClient := mapnests.NewClient(apiKey,pkgName,3000)

	ctx := context.Background()

	// Search
	fmt.Println("\n📍 Testing Search...")
	searchRes, err := mapClient.Search(ctx, mapnests.SearchRequest{
		Query: "Uttara, Dhaka",
	})
	if err != nil {
		log.Fatal("❌ Search error:", err)
	}
	fmt.Println("✅ Search result:\n", *searchRes)

	// Reverse
	fmt.Println("\n📍 Testing Reverse...")
	revRes, err := mapClient.Reverse(ctx, mapnests.ReverseRequest{
		Lat: 23.805675432193333,
		Lon: 90.42062140436256,
	})
	if err != nil {
		log.Fatal("❌ ReverseGeocode error:", err)
	}
	fmt.Println("✅ ReverseGeocode result:\n", revRes)

	// Distance Matrix
	fmt.Println("\n📍 Testing DistanceMatrix...")
	distanceRes, err := mapClient.DistanceMatrix(ctx, mapnests.DistanceMatrixRequest{
		OriginLat: 23.8103, // Dhaka
		OriginLon: 90.4125,
		DestLat:   23.8029, // Chattogram
		DestLon:   90.4226,
		Mode:      mapnests.TravelModeCar,
	})
	if err != nil {
		log.Fatal("❌ DistanceMatrix error:", err)
	}
	fmt.Println("✅ DistanceMatrix result:\n", *distanceRes)

	
	//Autocomplete
	limit:= int64(1) 
	lat:=   22.8029  
	lon:=  	92.5908

	autocompleteRes, err := mapClient.Autocomplete(ctx, mapnests.AutoCompleteRequest{
		Query: "Uttara, Dhaka",
		Lat: &lat, //optional
		Lon: &lon, //optional
		Limit: &limit,	//optional
	})
	if err != nil {
		log.Fatal("❌ Auto Complete Response error:", err)
	}
	fmt.Println("✅ Auto Complete Response result:\n", *autocompleteRes)

	//Autocomplete Without Zone
	
	autocompleteWithOutZoneRes, err := mapClient.AutocompleteWithoutZone(ctx, mapnests.AutoCompleteRequest{
		Query: "Uttara, Dhaka",
		
	})
	if err != nil {
		log.Fatal("❌ Auto Complete  Without Zone Response error:", err)
	}
	fmt.Println("✅ Auto Complete Without Zone Response result:\n", *autocompleteWithOutZoneRes)

}
