package main

import (
	"context"
	"encoding/json"
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
	fmt.Println("\n📍 Testing Distance Matrix Details...")
	distanceDetailsRes, err := mapClient.DistanceMatrixDetails(ctx, mapnests.DistanceMatrixDetailsRequest{
		OriginLat	: 23.7806,
    	OriginLon	: 90.3984,
    	DestLat		: 23.774,
    	DestLon		: 90.3681,
		Mode:      mapnests.TravelModeCar,
	})
	if err != nil {
		log.Fatal("❌ Distance Matrix Details error:", err)
	}
	fmt.Println("✅ Distance Matrix Details result:\n", distanceDetailsRes)

	//Pairwise Route Summary
	fmt.Println("\n📍 Testing PairWiseRouteSummary...")
	pairwiseRes, err := mapClient.PairWiseRouteSummary(ctx, mapnests.PairWiseRouteSummaryRequest{
		Pairs: []mapnests.PairWiseRoute{
			{ID: 1, Src: mapnests.Coordinate{Lat: 23.8113, Lon: 90.4135}, Dest: mapnests.Coordinate{Lat: 23.7815, Lon: 90.4123}, Mode: mapnests.TravelModeBicycling},
			{ID: 2, Src: mapnests.Coordinate{Lat: 23.8123, Lon: 90.4145}, Dest: mapnests.Coordinate{Lat: 23.7825, Lon: 90.4133}, Mode: mapnests.TravelModeBicycling},
			{ID: 3, Src: mapnests.Coordinate{Lat: 23.8133, Lon: 90.4155}, Dest: mapnests.Coordinate{Lat: 23.7835, Lon: 90.4143}, Mode: mapnests.TravelModeBicycling},
		},
	})
	if err != nil {
		log.Fatal("PairwiseRouteSummary error:", err)
	}
	fmt.Println("PairwiseRouteSummary result:", *pairwiseRes)
	pairwiseResJSON, _ := json.MarshalIndent(pairwiseRes, "", "  ")
	fmt.Println("✅ Pairwise Route Summary Response result:\n" + string(pairwiseResJSON))

	// Multi Source Route Summary
	fmt.Println("\n📍 Testing MultiSourceRouteSummary...")
	multiRes, err := mapClient.MultiSourceRouteSummary(ctx, mapnests.MultiSourceRouteSummaryRequest{
		Sources: []mapnests.Source{
			{ID: 1, Lat: 23.7805733, Lon: 90.2792399, Mode: string(mapnests.TravelModeCar)},
			{ID: 2, Lat: 23.75, Lon: 90.36, Mode: string(mapnests.TravelModeCar)},
			{ID: 3, Lat: 23.7, Lon: 90.42, Mode: string(mapnests.TravelModeCar)},
			{ID: 4, Lat: 23.7654321, Lon: 90.3456789, Mode: string(mapnests.TravelModeCar)},
			{ID: 5, Lat: 23.7123456, Lon: 90.3765432, Mode: string(mapnests.TravelModeCar)},
		},
		Destination: mapnests.Destination{Lat: 23.810332, Lon: 90.412518},
	})
	if err != nil {
		log.Fatal("MultiSourceRouteSummary error:", err)
	}
	multiResJSON, _ := json.MarshalIndent(multiRes, "", "  ")
	fmt.Println("✅ MultiSourceRouteSummary Response result:\n" + string(multiResJSON))

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
		log.Fatal("❌ Reverse Geocode error:", err)
	}
	revResJSON, _ := json.MarshalIndent(revRes, "", "  ")
	fmt.Println("✅ Reverse Geocode Response result:\n" + string(revResJSON))

	//Autocomplete
	var limit int64 = 1
	var activeZone bool = false
	autocompleteRes, err := mapClient.Autocomplete(ctx, mapnests.AutoCompleteRequest{
		Query: "Mirpur",
		ActiveZone: &activeZone,
		Limit: &limit,
	})
	if err != nil {
		log.Fatal("❌ Auto Complete Response error:", err)
	}
	autocompleteResJSON, _ := json.MarshalIndent(autocompleteRes, "", "  ")
	fmt.Println("✅ Auto Complete Response result:\n" + string(autocompleteResJSON))

	// Search by Radius
	fmt.Println("\n📍 Testing Search by Radius...")
	searchByRadiusRes, err := mapClient.SearchByRadius(ctx, mapnests.SearchByRadiusRequest{
		Query: "Uttara, Dhaka",
		Radius: 1000,
		Lat: 23.8766874,
		Lon: 90.3576884,
	})
	if err != nil {
		log.Fatal("❌ Search by Radius error:", err)
	}
	fmt.Println("✅ Search result:\n", *searchByRadiusRes)

	// Details Search By PlaceID
	placeDetailsRes, err := mapClient.DetailsByPlaceID(ctx, mapnests.DetailsByPlaceIDRequest{
		PlaceID: "86645dbb7d0cf3d466fa169d4f66560de61a63aacdc355ee870c16d3b4feb0b2",
	})
	if err != nil {
		log.Fatal("❌ Place Details Response error:", err)
	}
	placeDetailsResJSON, _ := json.MarshalIndent(placeDetailsRes, "", "  ")
	fmt.Println("✅ Place Details Response result:\n" + string(placeDetailsResJSON))

	
	

}
