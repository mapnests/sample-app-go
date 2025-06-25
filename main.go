package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	mapnests "github.com/mapnests/sdk-go"
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

	mapClient := mapnests.NewClient(
		apiKey,
		pkgName,
	)

	ctx := context.Background()

	// Search
	fmt.Println("\n📍 Testing Search...")
	searchRes, err := mapClient.Search(ctx, mapnests.SearchRequest{
		Query: "Uttara Sector 3, Dhaka",
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

	// Distance Matrix Details
	fmt.Println("\n📍 Testing DistanceMatrixDetails...")
	detailsRes, err := mapClient.DistanceMatrixDetails(ctx, mapnests.DistanceMatrixDetailsRequest{
		OriginLat: 23.8103,
		OriginLon: 90.4125,
		DestLat:   22.3475,
		DestLon:   91.8123,
		Mode:      mapnests.TravelModeBicycling,
	})
	if err != nil {
		log.Fatal("❌ DistanceMatrixDetails error:", err)
	}
	fmt.Println("✅ DistanceMatrixDetails result:\n", *detailsRes)
}
