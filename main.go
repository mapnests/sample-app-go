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

type apiCall func(ctx context.Context, client mapnests.Client) error

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

	mapClient := mapnests.NewClient(apiKey, pkgName)

	ctx := context.Background()

	// Comment/uncomment the lines below to choose which calls to run.
	run(ctx, mapClient, "DistanceMatrix", distanceMatrixCall)
	run(ctx, mapClient, "DistanceMatrixDetails", distanceMatrixDetailsCall)
	run(ctx, mapClient, "PairwiseRouteSummary", pairwiseRouteSummaryCall)
	run(ctx, mapClient, "MultiSourceRouteSummary", multiSourceRouteSummaryCall)
	run(ctx, mapClient, "Search", searchCall)
	run(ctx, mapClient, "Geocode", geocodeCall)
	run(ctx, mapClient, "Reverse Geocode", reverseCall)
	run(ctx, mapClient, "Autocomplete", autocompleteCall)
	run(ctx, mapClient, "Search By Radius", searchByRadiusCall)
	run(ctx, mapClient, "Details by Place ID", detailsByPlaceIDCall)
	run(ctx, mapClient, "Snap to Road", snapToRoadCall)
	run(ctx, mapClient, "MultiStopPoint", multiStopPointCall)
	run(ctx, mapClient, "MultiSourceRouteSummaryWithoutGeometry", multiSourceRouteSummaryWithoutGeometryCall)
	run(ctx, mapClient, "ETA Without Geometry By Modes", etaWithoutGeometryByModesCall)
	run(ctx, mapClient, "ETA With Geometry By Modes", etaWithGeometryByModesCall)
	run(ctx, mapClient, "ETA WithOut Stoppage", etaWithOutStoppageCall)
	run(ctx, mapClient, "ETA Multi Stoppage", etaMultipleStoppageCall)
	run(ctx, mapClient, "Pair Wise Distance Matrix", pairWiseDistanceMatrixCall)
}

// run prints a header for the named call, executes it, and exits on error.
func run(ctx context.Context, client mapnests.Client, name string, call apiCall) {
	fmt.Printf("\n📍 Testing %s...\n", name)
	if err := call(ctx, client); err != nil {
		log.Fatalf("❌ %s error: %v", name, err)
	}
}

func distanceMatrixCall(ctx context.Context, client mapnests.Client) error {
	res, err := client.DistanceMatrix(ctx, mapnests.DistanceMatrixRequest{
		OriginLat: 23.8103, // Dhaka
		OriginLon: 90.4125,
		DestLat:   23.8029, // Chattogram
		DestLon:   90.4226,
		Mode:      mapnests.TravelModeCar,
	})
	if err != nil {
		return err
	}
	fmt.Println("✅ DistanceMatrix result:\n", *res)
	return nil
}

func distanceMatrixDetailsCall(ctx context.Context, client mapnests.Client) error {
	res, err := client.DistanceMatrixDetails(ctx, mapnests.DistanceMatrixDetailsRequest{
		OriginLat: 23.7806,
		OriginLon: 90.3984,
		DestLat:   23.774,
		DestLon:   90.3681,
		Mode:      mapnests.TravelModeCar,
	})
	if err != nil {
		return err
	}
	fmt.Println("✅ Distance Matrix Details result:\n", res)
	return nil
}

func pairwiseRouteSummaryCall(ctx context.Context, client mapnests.Client) error {
	res, err := client.PairWiseRouteSummary(ctx, mapnests.PairWiseRouteSummaryRequest{
		Pairs: []mapnests.PairWiseRoute{
			{ID: 1, Src: mapnests.Coordinate{Lat: 23.8113, Lon: 90.4135}, Dest: mapnests.Coordinate{Lat: 23.7815, Lon: 90.4123}, Mode: mapnests.TravelModeBicycling},
			{ID: 2, Src: mapnests.Coordinate{Lat: 23.8123, Lon: 90.4145}, Dest: mapnests.Coordinate{Lat: 23.7825, Lon: 90.4133}, Mode: mapnests.TravelModeBicycling},
			{ID: 3, Src: mapnests.Coordinate{Lat: 23.8133, Lon: 90.4155}, Dest: mapnests.Coordinate{Lat: 23.7835, Lon: 90.4143}, Mode: mapnests.TravelModeBicycling},
		},
	})
	if err != nil {
		return err
	}
	resJSON, _ := json.MarshalIndent(res, "", "  ")
	fmt.Println("✅ Pairwise Route Summary Response result:\n" + string(resJSON))
	return nil
}

func multiSourceRouteSummaryCall(ctx context.Context, client mapnests.Client) error {
	res, err := client.MultiSourceRouteSummary(ctx, mapnests.MultiSourceRouteSummaryRequest{
		Sources: []mapnests.Source{
			{ID: 1, Lat: 23.7805733, Lon: 90.2792399, Mode: string(mapnests.TravelModeCar)},
			{ID: 2, Lat: 23.75, Lon: 90.36, Mode: string(mapnests.TravelModeCar)},
		},
		Destination: mapnests.Destination{Lat: 23.810332, Lon: 90.412518},
	})
	if err != nil {
		return err
	}
	resJSON, _ := json.MarshalIndent(res, "", "  ")
	fmt.Println("✅ MultiSourceRouteSummary Response result:\n" + string(resJSON))
	return nil
}

func searchCall(ctx context.Context, client mapnests.Client) error {
	var limit int64 = 5
	res, err := client.Search(ctx, mapnests.SearchRequest{
		Query: "Uttara, Dhaka",
		Limit: &limit,
	})
	if err != nil {
		return err
	}
	fmt.Println("✅ Search result:\n", *res)
	return nil
}

func geocodeCall(ctx context.Context, client mapnests.Client) error {
	var limit int64 = 5
	res, err := client.Geocode(ctx, mapnests.GeocodeRequest{
		Query: "Uttara, Dhaka",
		Limit: &limit,
	})
	if err != nil {
		return err
	}
	resJSON, _ := json.MarshalIndent(res, "", "  ")
	fmt.Println("✅ Geocode Response result:\n" + string(resJSON))
	return nil
}

func reverseCall(ctx context.Context, client mapnests.Client) error {
	res, err := client.Reverse(ctx, mapnests.ReverseRequest{
		Lat: 23.805675432193333,
		Lon: 90.42062140436256,
	})
	if err != nil {
		return err
	}
	resJSON, _ := json.MarshalIndent(res, "", "  ")
	fmt.Println("✅ Reverse Geocode Response result:\n" + string(resJSON))
	return nil
}

func autocompleteCall(ctx context.Context, client mapnests.Client) error {
	var limit int64 = 1
	var activeZone bool = false
	res, err := client.Autocomplete(ctx, mapnests.AutoCompleteRequest{
		Query:      "Mirpur",
		ActiveZone: &activeZone,
		Limit:      &limit,
	})
	if err != nil {
		return err
	}
	resJSON, _ := json.MarshalIndent(res, "", "  ")
	fmt.Println("✅ Auto Complete Response result:\n" + string(resJSON))
	return nil
}

func searchByRadiusCall(ctx context.Context, client mapnests.Client) error {
	res, err := client.SearchByRadius(ctx, mapnests.SearchByRadiusRequest{
		Query:  "Uttara, Dhaka",
		Radius: 1000,
		Lat:    23.8766874,
		Lon:    90.3576884,
	})
	if err != nil {
		return err
	}
	fmt.Println("✅ Search result:\n", *res)
	return nil
}

func detailsByPlaceIDCall(ctx context.Context, client mapnests.Client) error {
	res, err := client.DetailsByPlaceID(ctx, mapnests.DetailsByPlaceIDRequest{
		PlaceID: "5f50dd3902085e86d414427a41559222cdfd4d8169378a50ad6659d250601067",
	})
	if err != nil {
		return err
	}
	resJSON, _ := json.MarshalIndent(res, "", "  ")
	fmt.Println("✅ Place Details Response result:\n" + string(resJSON))
	return nil
}

func snapToRoadCall(ctx context.Context, client mapnests.Client) error {
	res, err := client.SnapToRoad(ctx, mapnests.SnapToRoadRequest{
		Mode:      "walking",
		Latitude:  23.8103,
		Longitude: 90.4125,
	})
	if err != nil {
		return err
	}
	resJSON, _ := json.MarshalIndent(res, "", "  ")
	fmt.Println("✅ Snap To Road Response result:\n" + string(resJSON))
	return nil
}

func multiStopPointCall(ctx context.Context, client mapnests.Client) error {
	res, err := client.MultiStopPoints(ctx, mapnests.MultiStopPointsRequest{
		Src: mapnests.Coordinate{Lat: 23.809973415982903, Lon: 90.35697149649764},
		StopPoints: []mapnests.StopPoint{
			{ID: 1, Lat: 23.81038738311683, Lon: 90.36203008256733},
			{ID: 2, Lat: 23.798308134287165, Lon: 90.43522641639149},
		},
		Mode: mapnests.TravelModeMotorcycle,
	})
	if err != nil {
		return err
	}
	resJSON, _ := json.MarshalIndent(res, "", "  ")
	fmt.Println("✅ MultiStopPoints Response result:\n" + string(resJSON))
	return nil
}

func multiSourceRouteSummaryWithoutGeometryCall(ctx context.Context, client mapnests.Client) error {
	res, err := client.MultiSourceRouteSummaryWithoutGeometry(ctx, mapnests.MultiSourceRouteSummaryRequest{
		Sources: []mapnests.Source{
			{ID: 1, Lat: 23.7805733, Lon: 90.2792399, Mode: string(mapnests.TravelModeCar)},
			{ID: 2, Lat: 23.75, Lon: 90.36, Mode: string(mapnests.TravelModeCar)},
		},
		Destination: mapnests.Destination{Lat: 23.810332, Lon: 90.412518},
	})
	if err != nil {
		return err
	}
	resJSON, _ := json.MarshalIndent(res, "", "  ")
	fmt.Println("✅ MultiSourceRouteSummary Response Without Geometry result:\n" + string(resJSON))
	return nil
}

func etaWithoutGeometryByModesCall(ctx context.Context, client mapnests.Client) error {
	var numberOfRoutes int64 = 3
	res, err := client.EtaWithoutGeometryByMode(ctx, mapnests.ETAByModeRequest{
		FromLat:    23.82440246492232,
		FromLon:    90.35823219240551,
		ToLat:      23.77683932043941,
		ToLon:      90.40510870368077,
		Mode:       mapnests.TravelModeCar,
		NoOfRoutes: &numberOfRoutes,
	})
	if err != nil {
		return err
	}
	resJSON, _ := json.MarshalIndent(res, "", "  ")
	fmt.Println("✅ ETA Service Response Without Geometry result:\n" + string(resJSON))
	return nil
}

func etaWithGeometryByModesCall(ctx context.Context, client mapnests.Client) error {
	res, err := client.EtaWithGeometryByMode(ctx, mapnests.ETAWithGeometryByModeRequest{
		FromLat: 23.7805733,
		FromLon: 90.2792399,
		ToLat:   23.810332,
		ToLon:   90.412518,
		TripId:  "550e8400-e29b-41d4-a716-446655440000",
		Mode:    mapnests.TravelModeCar,
	})
	if err != nil {
		return err
	}
	resJSON, _ := json.MarshalIndent(res, "", "  ")
	fmt.Println("✅ ETA Service Response With Geometry (Single Mode) result:\n" + string(resJSON))
	return nil
}

func etaWithOutStoppageCall(ctx context.Context, client mapnests.Client) error {
	res, err := client.EtaWithoutStoppage(ctx, mapnests.ETAsRequest{
		FromLat: 23.759248321873432,
		FromLon: 90.41691646513276,
		ToLat:   23.770148309864418,
		ToLon:   90.40842674496818,
	})
	if err != nil {
		return err
	}
	// Print all data
	// resJSON, _ := json.MarshalIndent(res, "", "  ")
	// fmt.Println("✅ ETA Service Response Without Stoppage (Multi Mode) result:\n" + string(resJSON))

	//Print specific mode data
	cngRouteData := res.Data["cng"]
	cngRouteDataJson, _:=json.MarshalIndent(cngRouteData, "", "  ")
	fmt.Println("✅ ETA Service Response Without Stoppage (Multi Mode) result:\n" + string(cngRouteDataJson))
	
	return nil
}

func etaMultipleStoppageCall(ctx context.Context, client mapnests.Client) error {
	req := mapnests.ETAMultipleStoppageRequest{
		Source:  mapnests.LatLng{Latitude: 23.8115, Longitude: 90.414},
		DropOff: mapnests.LatLng{Latitude: 23.8115, Longitude: 90.414},
		Modes: []mapnests.Mode{
			mapnests.TravelModeCar,
			mapnests.TravelModeCNG,
			mapnests.TravelModeMotorcycle,
		},
		Stoppages: []mapnests.ETAStoppage{
			{Latitude: 23.795, Longitude: 90.41, StoppageOrder: 1},
		},
		TripId: "550e8400-e29b-41d4-a716-446655440000",
	}

	res, err := client.EtaMultiStoppage(ctx, req)
	if err != nil {
		return err
	}

	// Dynamic access, driven by whatever modes were requested
	for _, mode := range req.Modes {
		modeData, ok := res.Data[mode]
		if !ok || modeData == nil {
			continue
		}
		modeDataJSON, _ := json.MarshalIndent(modeData, "", "  ")
		fmt.Printf("Mode: %s\n%s\n", mode, modeDataJSON)
	}

	return nil

}

func pairWiseDistanceMatrixCall(ctx context.Context, client mapnests.Client) error {
	res, err := client.PairwiseDistanceMatrix(ctx, mapnests.PairwiseDistanceMatrixRequest{
		Mode: mapnests.TravelModeWalking,
		Routes: []mapnests.PairwiseDistanceMatrixRoute{
			{ID: 1, Origin: mapnests.Coordinate{Lat: 23.81183142056928, Lon: 90.41966658048152}, Destination: mapnests.Coordinate{Lat: 23.81011947054907, Lon: 90.42197475156428}},
			{ID: 2, Origin: mapnests.Coordinate{Lat: 23.81284142056928, Lon: 90.41969658048151}, Destination: mapnests.Coordinate{Lat: 23.81041947054907, Lon: 90.42107475156428}},
		},
		Sort: mapnests.SortDistanceInMetresAsc,
	})
	if err != nil {
		return err
	}
	resJSON, _ := json.MarshalIndent(res, "", "  ")
	fmt.Println("✅ Pairwise Distance Matrix Response result:\n" + string(resJSON))
	return nil
}
