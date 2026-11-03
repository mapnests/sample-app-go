package main

import (
	"context"
	"encoding/base64"
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

	//------------Geo Map --------------//
	run(ctx, mapClient, "Search", searchCall)
	// run(ctx, mapClient, "Geocode", geocodeCall)
	// run(ctx, mapClient, "Reverse Geocode", reverseCall)
	// run(ctx, mapClient, "Autocomplete", autocompleteCall)
	// run(ctx, mapClient, "Search By Radius", searchByRadiusCall)
	// run(ctx, mapClient, "Details by Place ID", detailsByPlaceIDCall)

	//----------Route Map ------------//

	// run(ctx, mapClient, "DistanceMatrix", distanceMatrixCall)
	// run(ctx, mapClient, "DistanceMatrixDetails", distanceMatrixDetailsCall)
	// run(ctx, mapClient, "PairwiseRouteSummary", pairwiseRouteSummaryCall)
	// run(ctx, mapClient, "MultiSourceRouteSummary", multiSourceRouteSummaryCall)
	// run(ctx, mapClient, "Snap to Road", snapToRoadCall)
	// run(ctx, mapClient, "MultiStopPoint", multiStopPointCall)
	// run(ctx, mapClient, "MultiSourceRouteSummaryWithoutGeometry", multiSourceRouteSummaryWithoutGeometryCall)
	// run(ctx, mapClient, "Pair Wise Distance Matrix", pairWiseDistanceMatrixCall)
	// run(ctx, mapClient, "MultiStopPoint With Geometry", multiStopPointWithGeometryCall)

	// //-----------ETA-------------
	// run(ctx, mapClient, "ETA Without Geometry By Modes", etaWithoutGeometryByModesCall)
	// run(ctx, mapClient, "ETA With Geometry By Modes", etaWithGeometryByModesCall)
	// run(ctx, mapClient, "ETA WithOut Stoppage", etaWithOutStoppageCall)
	// run(ctx, mapClient, "ETA Multi Stoppage", etaMultipleStoppageCall)
	// run(ctx, mapClient, "ETA With MultiSource MultiStopPoint", etaWithMultiSourceMultiStopPointCall)

	//Static Map
	// run(ctx, mapClient, "Static Map", staticMapCall)

}

// run prints a header for the named call, executes it, and exits on error.
func run(ctx context.Context, client mapnests.Client, name string, call apiCall) {
	fmt.Printf("\n📍 Testing %s...\n", name)
	if err := call(ctx, client); err != nil {
		log.Fatalf("❌ %s error: %v", name, err)
	}
}

func distanceMatrixCall(ctx context.Context, client mapnests.Client) error {
	// traceparent := "00-0af7651916cd43dd8448eb211c80319c-b7ad6b7169203331-01"
	res, err := client.DistanceMatrix(ctx, mapnests.DistanceMatrixRequest{
		OriginLat:   23.8103, // Dhaka
		OriginLon:   90.4125,
		DestLat:     23.8029, // Chattogram
		DestLon:     90.4226,
		Mode:        mapnests.TravelModeCar,
		// Traceparent: &traceparent,
	})
	if err != nil {
		return err
	}
	fmt.Println("\n✅ DistanceMatrix result:\n", *res)
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
		PlaceID: "a55cbbb89e9a8f91fe1a08da9665b1811473a8fe020290765418769eaf195155",
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
	var numberOfRoutes int64 = 1
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
		Modes: []mapnests.Mode{
			mapnests.TravelModeCar,
			mapnests.TravelModeCNG,
		},
	})
	if err != nil {
		return err
	}
	cngRouteData := res.Data
	cngRouteDataJson, _ := json.MarshalIndent(cngRouteData, "", "  ")
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

func multiStopPointWithGeometryCall(ctx context.Context, client mapnests.Client) error {
	res, err := client.MultiStopPointDetailsWithGeometry(ctx, mapnests.MultiStopPointDetailsWithGeometryRequest{
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
	fmt.Println("✅ MultiStopPoints with geometry Response result:\n" + string(resJSON))
	return nil
}

func etaWithMultiSourceMultiStopPointCall(ctx context.Context, client mapnests.Client) error {
	res, err := client.EtaWithMultiSourceMultiStopPoint(ctx, mapnests.EtaWithMultiSourceMultiStopPointRequest{
		PickupPoint: mapnests.Coordinate{Lat: 23.809973415982903, Lon: 90.35697149649764},
		Riders: []mapnests.EtaWithMultiSourceMultiStopPointRider{
			{
				RiderID:         "rider-1",
				CurrentLocation: mapnests.Coordinate{Lat: 23.81038738311683, Lon: 90.36203008256733},
				DropOffPoint:    mapnests.Coordinate{Lat: 23.798308134287165, Lon: 90.43522641639149},
			},
			{
				RiderID:         "rider-2",
				CurrentLocation: mapnests.Coordinate{Lat: 23.770348, Lon: 90.407785},
				DropOffPoint:    mapnests.Coordinate{Lat: 23.794249, Lon: 90.414536},
			},
		},
		Mode: mapnests.TravelModeMotorcycle,
	})
	if err != nil {
		return err
	}
	resJSON, _ := json.MarshalIndent(res, "", "  ")
	fmt.Println("✅ ETA With MultiSource MultiStopPoint Response result:\n" + string(resJSON))
	return nil
}

func staticMapCall(ctx context.Context, client mapnests.Client) error {

	res, err := client.StaticMap(ctx, mapnests.StaticMapRequest{
		Polyline: "ylqxq@iklsgD{ELwGq@_M}A_PyCgQuCiB[eKgBh@oDfLpBhANvYnE~NvBvLzBnIxCxn@vLnq@pMdB\\|[zGtE`AxTvEjTvErTfFhd@rKnt@dN`NlCdFvAvi@rLha@hGxb@pIl^lHld@pJvx@lQdRnDlg@vKxk@vL~~@rSjf@jKli@fLjy@dRhj@xLzQbE`RbEdZtGl]xHrc@bK|^nIb[|H~FjBvb@bJzPfE~X~FpTfGrQ~EnM`EtSjHbWlJlVjJvUzJ~i@bUx^rOr_@fQl\\lNxV`Ljh@~Ujh@fUtmAli@xGpDtGhDhl@h[rM~Fh@Tjb@lRjVxK`y@|]~c@vRpv@r\\hZrMp[rM|BbAlx@d^~zAbp@ppAdk@bzD|cBvbAnc@nyAlo@fCfAzg@zTrmAvh@dnBn|@`Ab@j[nNxwAdq@`Bv@vQbIhe@xSly@t^rJhEd_@lPjo@fYxDbBbIpDhD|A|n@hYru@~\\fw@z]xg@lU~@`@tr@zYzr@`\\vhBrw@lF~BlAh@xPdHt\\dOph@dUdb@xQ~yAro@jGnC|DbB~lBxy@lrBz}@tPvHlYjM`HxC|FbC`\\`OjSbJtSfJb_@nPxX`M|c@bSxa@|Qj^bPtrApl@fMvFdDzAnQtHvp@jY~aAnb@nhBzv@vOzG~~@|`@hx@h]t`Bdu@nNrGhu@d]~Ar@lg@~T|PjI`m@zWre@xS~M`GdeBzt@bd@dQrf@bUhbBv{@jw@hd@ncAjl@xRrLtfKvzFtf@nXhBbAn}@vg@bg@jWra@`O`i@bOtpA`Yni@vKxFxAdoAfYroApYriD`zA`n@bXhfB`w@ry@z^|mBn_Al`Apc@~uA`n@`]tQt^rUh^rVlh@~\\n}@hl@lW`PdHtBxJjB`Fn@xFTfGInJcAx^uJx]sKfgAc\\ngAi[dkA_]rd@mMrd@yMt|Aid@z|A{b@f|DmhAtp@uRltA{`@nK}C~]iKvrAe`@|b@cLnEkAli@cPli@{Nni@sPns@iSxAa@dxBin@rB]pj@wQv_@eRxTgXfLsUhHiZxFoZzDeXnXofBfR{iApF}\\~}@}xFjOwdAdCiSdCgUhE{u@j@_KfEw}@v@sh@h@_^PaLpDacCx@i[hF}l@t@wOa@wwDGq}BE{{AM_uB^yv@f@gOn@iSpFsgBlE}vArAgc@nA}e@JsD`C}~@nD_uAlBqe@vEeg@xHmk@|CsQfNo~@bHmg@nL{xAl@}`@bA{rAhAsu@z@iIzn@gaCbt@giCfAeEtH}Yzb@_lBJiAxXgaD`i@igG~Gax@LqAjIi`AbKmf@dl@cmAx]y}@hh@icExD_}@`Hm`BpEcKpO}]pO_g@~Iwm@rEcb@nD{SpDaKnEaKfFqLvFgJxy@g{@xLeJ|c@op@hOqUrSq[fJmNnNcVrJ{VhDek@bA{fCh@_[aA{x@TmRGyVdDmeC@mI?sD@aRf@qZb@mYzC}TjK_ZtM__@zQ}d@vLkXtDaJfPk`@dNm\\fBmEzs@kmBtYso@lB{EvOm`@dNa^hBmRt@}HnOy_BbM{fArMovArDc^jj@umFn@kIlHmt@~AoUDwVwP{~A}@ip@f@yt@rBwv@~E}sBbHyzDnLoeG|@i\\`Eu`BtEy_B|HihBpA}g@bBgOpCwNzBoGzCsFlDiElDyDzZeYjM{Itn@mf@~McKfLgOvBoHpDiMnAqOCgOmBeQoE}Oyi@us@mtAwbBaHkLoLwTwEeN_n@adBmUqm@wcA_qAatB}hCaHyL}PmZk^yxAwFsZwcA}~F{p@giEal@kxDmUk|AkRazAuGgj@sNadA{RavAiFg`@_AyGoBqRoA}OeBca@k@eSqBqv@sAeo@s@iWcA}UeA{NmAkQsAaQaCk\\sFku@{Few@s@aLKqHN{In@yJ~@qHfB_IlCuI`E}IhKwQhVs_@fd@is@bq@udAr^uk@|`@kn@lf@qu@|b@ur@tD{F~_AoyAfRyY|l@c_Anx@aqAjoAeoB`dB_jCxNwUle@{s@bc@sq@djBavCdDaFh~Aq`CfJgNxb@uo@t{@esA|r@gfA`LiQhf@ct@`r@geAffAqaBzYuc@raA{zAh~@evA`l@{}@|Zke@vIqKxHaHtJeGtJaDlJ_BzK_@dQnAbNnB`^pFbWhD|Jj@|GAnEg@bEuA|DuBnEmE`DoFdB{E~AcJxEuu@tD_m@vEyt@|HuxA~A}VhB}O`DmQxCmJpE_KfFsIbGaIzHkI~JeHbJwEzJ{DpPoD|R{@h]Djd@D~wA~@z`An@nyBjBblA~Bpj@jDv}@fKhZzClPxA~HfAxHf@zKT`O_@fMy@dWqB|UiBnT{AjSuAzPiAnRuApM[`Le@lM?~_@p@be@r@bnAvCb{BdDnrAnCxuDzGt{@d@d~@Nrx@^nRnAxLtBvGtAxMpDrb@|QnyAjt@dMlEld@`PxBv@h~CtgAp`@jNvzAdi@~aAz]jmBvq@zq@pVjeD|kAtOtCtP`Bz@B|i@zB~DHnAD`g@bAleAjCjNg@`NeBrj@kLz}@yVx`A_g@fQgIzCwAz\\gJzM}BbPgC~UuD`hF}v@niCg_@j|@{LvwBmZrd@oG~z@sLhgAqN|iAyL~|@gJxFg@n`@kCv_@Xh_AhC`v@xCzaAtDpk@l@p_@^p\\\\bA@j}BcCdi@k@|yFcMlw@eBprAwChh@uBjCGnd@gA`rAq@dUAfp@kCn^E|Qd@hb@tGjXzDbkDzf@dJjA|Gx@hpAjRbANptCra@vp@fI|mCxX`c@lE`dAhKbCXtpA`OlYdD`xAnPlZpDrvAfR~i@tJri@xLdo@tR`e@nNrqAt`@jVpHnx@xNlGfAz\\zFdoAbTxXxE~XvCraAxOvn@rL|{D`x@`rAzUvrCrg@fFh@hF\\hNb@jkAoH~N}@twA_JdTq@tBLvBd@`B`@vAbAvA|AfHvH~JxM`Yf_@hd@xm@bc@vk@n@v@nNjQ|G~GhI~G~LdJzFpFxFnGbGvH`KxMjc@xk@nZf`@dJlNdDtAhBt@`AP~@FXAj@?~@OTIn@Q`Aq@tGoFYzBaBhHeA~Dc@dBc@tAk@pAqBbIcB~EuB`IqFtQsZjaAcIxVqLx^q@xB_T`q@sAjEmAxD{DhLgEnKsFvLcFlLgGfNmErJ{B`F}@pBuKpVs@`BgB|D_Pd^wKnVwKlVmMfY{E|KwBjF[x@yAvEsAvGOrBSlCMrFDnEb@rEp@lEnApFlHjUhInVbB|EzGbRpH`UzKf_@pLv`@tPph@`Tlt@xI|YzFlRjYzaAtEiBpAy@j]i\\zBaDn@oBLuB`BEdGcClFyAbGgB|By@hBiB`BmCpKaKv@}AbAaCdCaG_@gAaAwImAoLw@kK[mJMqJReI^yH|@sIvAsJxAiIfCyK~CuI|DoJvEsJvD_HhF{HrDkEjGmFzGsExJoF|k@m\\df@iYldA{m@m\\oDyL|@oMn@ZkBR{B@{DTiCl@kChC_F~@uC\\kDNeDYaGi@iK]yP_A_ZQ{B_BsIyF}Zm@mDEiBFeBh@_DhGq\\xDcYvByPt@}N`@iPZ_[HaDToCf@oCvE{M`DgNp@yEb@gFLqJXwN\\yI[oEiAqFsDuNwFeQoEgHeB}DcEgJyB}JcAqFu@}H]aISaJGcDCoLEsKRkPS_IOkCa@w@k@a@uDs@cF}BgDs@kFUmECy@]m@{@g@_BBaHxAm\\FyO\\iM^k@t@Ch@a@J{@|@wWRePCcNKaFy@kKnNw@|e@mCvO{@zO_Ard@}BvQe@tYQnNCxf@Oje@e@ji@]|q@Wnv@i@zPOjIYfLo@jS{Ajh@mEtj@aFxUuBjUuBp`@eDrvAeMl@G`e@wEtHa@~^gDb\\iCpFc@r\\}ChFg@vGo@bI_AfFeAlF}A`@QdEgBvGuCl[oPp]eQr@_@`IcEvZ}O|FmCxGiC`HcCfGaB`FcArF_AjGw@zJaAfTyAx\\yCnUoBjg@aEjd@mDdGe@fTiBnGo@xGaAlHwAzG{AjOyDhJaCzUqGvF{AjK}BdMwBjGw@jGk@hH[`ISvDCrFClOG|IClBArQHxDBfJ?zTF`QIrHIrCErCIbCEpLWhFUtQ{@rf@eB|@BrEGpGQpOy@|g@qB`Ee@nFa@fEGrF]dF[dFa@pO}BhAQx\\cEv]cH|[iFhr@qJfn@qGzEYfJk@rjCcPhmBoLbmAqGxeBiJtsAiH~bCoKpESjkAoFzi@?j|@rOzi@rMz~AhSfeAzMpkAxNfjAzL~JnAjXjDjd@vGflChZx^bElOvA`OlBx^vEz_AbKl\\~D`]tBl[XxZ}@zMe@nh@iBzwBkGvbAaBb}BuDtU_@|Wc@vBEz_CyDnHKpzAkGhoAgKl\\}Jrm@wTzWsJbIuC`IaDvSgHjG}@hAQfI}@`XoAfDOzVkA~WiAvPsAvFe@`E_@fEg@lKwAh]sDtLYjLFxY`@ft@|@jIa@bYsAzv@oGj[}Bd[y@lIZleB~TtlBnWhMnBbVdC~g@`IfnAhRbzAlTpMlB~_@hBhDT~\\hBzYlBjh@hEfdBpPtG^`HBbHDxHKpoAoE`jAyC~EMx\\dBlAFjEl@xX|DrEh@~M~ApUtAnE\\lv@|B~`@bAz^|@foAzXhoA|XtBb@v`@|HzDv@|GrAfOx@nQbA`b@h@zmBbC~]uAj_@kCh[v@hMjCjnAhWfv@jQxLzAjG`@zMn@j^sCvPiAtLm@`Lo@vJ?lLn@xHd@vKrAtC^lZhFlM`B``@~FhL~B~`@lAvFPrv@SzNRjJLbb@vAb]jAzl@hCvc@zBfJ?x[dAfOEpy@uCzG]ZCzHc@rNu@lq@sDbd@cCdH_@hx@kEzG_@nbCuMhPQtGI~TrAfrApUfIpEvHzH`GxO|J~\\^nAxDjJxA`D`J|Y~EfLzCzEnHrEhGlCvThGplE`bAfa@~Hf_A|SneAlWdZzFlV`GzzAra@zDfAtaLpfCn~@bUv|C~u@vq@vPpjBzb@n@LtrAhZveEvaA~TbGrm@bNpDv@tr@dPfIlBrUrFtNhDnP~Dt_@nIj_@dK|b@jKrb@jK|aAnUf~@rN~}@fMnJvAjiBfUjlB|Uj\\rEjc@fGv`AfL~_BjU~uDxd@rWbDf`CtYj{@nKha@bF~Ej@p\\|Dx|Cxb@zq@pJvIlA|[~@pbADvz@d@lzCtBbNJjYTjbDZjl@x@~Vd@hVrCzr@hNvr@jNnUrEvTjEbk@pLhRtDhdAjSxVxE`WrFpn@bMha@tFzHn@l`@bDb[~Gbm@|Kj\\xG`_AtY~\\bHleAtTdxBbc@rgBd^nMhCta@zHlDp@bh@pKxpA|Wh{@vP~mG~pAhuEz_AjmDptAhBr@pMxFdw@b\\bh@jUpnAzi@t_DnuAlwArn@~gD~yAxAl@`c@jRhzJpjEx`FhxB`fDxzAzsTj_E`m@bLjdDbn@piG~jAxNnCvqChi@|gCdg@nPbDz_@|PfdAfr@~oF`lDfvClkBpuAh}@x_Alm@vsBlsAnlAvw@pFjHpk@hv@lyAp{BhOnUx~AzgCbuCnuExZtf@peEhoGxG`Kj~@|qAn]t`@l[t\\ziBrnBzu@hu@|EhDbMjM~AdB~h@ji@jCbDjB]hBZrAnAf@pBIxBbFlFt^t_@rCpG~`BxbB`lAznAxpBztBrK`MlbCrjCxv@h~@pFnHxNrOjLpJvCjDnJnLdC|CbJlEvI`DzMzArHMdHuAfKgF|IiIrD}GpCqGrEwRtQerAn\\ghBbxAwuC`@y@tqBcaEnm@cz@jhCutCdaHe{H~nD{eEp`E_nC~nAox@d[wRtuAk}@h`@{VbUoNtuCwjBny@ol@jAy@|f@i\\|bBkuAl{GqoGx}@e{@`xBugB`v@sn@d@_@`kAm`Ah{@at@`\\uYzYi\\lkAmzAvMsPz|AykBrh@am@zV{Xr\\}WvnBqlA~ByAhlBsmAhz@}u@tj@ak@hqCklCpiDckCl]aXbIiGnPeLtDyDbFaIroBouEnVsc@vb@mo@~QqU`z@ybA`aA_lAjz@e}@vLuLhEmEnbAcaAhc@ab@dUcUdPiOrNgMrxB__Bze@u_@hGaFdf@}_@dUuQhNgRplAm`BnlA{eAnc@k`@dSgOxR}Oj`@a\\va@e]zc@w]tKyH|T{M|OgGz\\}Ij^qHzNwCfn@wMdn@}LvX}Ff]iInvAwYz`AaThc@{Jza@}OtFiCrHsElo@oc@nZwTdMcJbmByvAj]eWxVsOfd@iUbh@aUp{Aer@rhAug@tUeK`_Aua@hFaC|k@{Udh@}Szi@wTbX_IdnAi]ld@eKlkBwh@xsCow@dLmCrHsA|H}@dGYfYWdfDQtUAte@Fhv@Jzn@LjGWtGu@dEmA`FeChFyDfEeFpDqGzEsS~OmpAdDcZ`UeqBdHys@x@cMQkLaBaLqPq}@iHm|@_Ds^kTodC_B{UHcKpBe_@|^sqCbFs_@`Sq{AnTcbBvG_d@l^qcCj]av@lXsn@nFk`@dGcxCx@{[~L}|EfEwlALqDhz@m}DfKsp@xd@qxCrXifBxb@{oCfKop@nV_fBv@aJt@cKjGq{@fAiOtEmo@zLycBlDuf@bBuO`DaZpKc^lHwKpGoJxk@au@hGaIdZw_@n{@cgAh_@oe@hbAcpAnDuElo@ky@bsCwoDhhBe{BxmAywArt@}t@`lA_wAl]i_@|e@{m@pE{Fht@whAvwCgrEjy@ceApPkRnL{Mh{A{dBhN}ObCoExjB{iDhBiDxcAimBlDmGvNoWtNoQfNqKpO_LjtAcbAxVw]xVuc@zHkIbReL|cAcj@|aB{|@lf@i[rFwErsAojAjd@_[|f@mYhKcGz\\_UpHkG`j@_p@xDsEvfA{tApWeVbC{BdVwN|_@kUhg@_\\|LwLjGgOjBoPn@iNL_CvAgZ|B}U`FeUlGyPrGaOxCmGnLoV~Wql@vRkc@jX}k@rb@c~@xB{E|KiVzNs[hIqQtJaJ~a@iPf@SnXwJdQaDpVcAdFSfEIjfBkCrHKho@aAja@k@x~CyGndDmGxO[dS_@tg@cAf^q@xxAgC~c@o@pNGv`@SbOItZMz[mDb\\aL~J{DrdAka@zoBgt@llBax@pH_DtgCugAdJ}Df_Ay`@bJ{DhHwCbNyFnaAga@b]{DdlBsH|x@cDzfA_ElpAyE`Ig@n_@cClwByf@zq@yOx`Bq^rg@cLbvAa[bbIenBbMmJdGsEdFqIlH{LrNie@bfAybDtEaOpGmSnl@skBvSsp@nTir@|Jq[d`@onA|L_`@~O}SvNkIpJcDjK_CnH[~NBzy@GjKAfiAt@xgD{B`z@v@pNLdNE`}@WzhBqSbiCqXlpAyNvJiAnaA{KtOgBxToDhiAmMlvA}Ol\\kDbbAeK~`@_AbWrDf`@zKlI`DtYxKdq@dW`n@fTze@~LheDd{@t|@nTjYdHx]rHbe@|JxpAvS|KhBrUpDbCZdEp@``ApOfoCxc@d^dIdyBbg@znBtc@nDd@~Fn@n^}AjSmC`BSvpAuZbxBaj@no@yL`WBxkBXva@d@jF@bXg@rvBPn{@{BdnAiD~c@[p]]f\\qAvMk@nz@A`}@dEv`AbH`SdBdD_A|D{@bDiB|BaDdC{FxNw{A|OszA``@qbFrNgjBjg@}qF`QqrBtJahA`AiVXwEtG}iBf@uKrImlBpGcyBlG{mBx@sf@^kNpRy`FtJ}eCnFqpAzI}sBbCqk@dAaZ`AeYjAyo@UsK[oOIuDiIqn@kCoOwC}P_e@ecCeXkvAsEaT}AyHcg@ofAoq@ezA}h@acA}DsHeQg_@_w@gaBsr@kvA_k@ygAgjC{lC{kAkmA_MgK{GyFiMcJ{Aw@{PoGyyDweAg`Ckx@qZgKyHsEaDkByTaOkuDwpCk~@qu@aTkP}e@ca@oIqHoIqMoLySgJwWgGqSkB{Jy@{Jy@sMGoNv@{H`CyIzQib@hFwLjoAq_BxSyVnGwHdM_O~DiEpMiNdOaRxPiWjGaNxB_LbBuMFkP_@_QwDaRwEoO}e@ut@kE{Gkc@il@eQ{RaLsPuCuEkB{FsAgJK}KPq\\pBu[fGaiA`IklAwHgyGa@aYE{AqBcyBaJsy@uBmSkFwn@oSkiBqBwRIaF]{MAgKJaNl@kOtl@_pHpMo`BpFwu@PkCnFal@nAoN`E}[xEim@?kXmI{aAqKuaA{C}XiRizAsU_fEsSgaEyFswAHuOvDcu@tAqg@rBsc@dA{T~Bcv@X_KvFgbBvBwf@fBce@tBwb@`Ci|@z@aR|EaaB|Aig@rAoZbA{U~Fgb@xFuUlZsnArlA{jFNs@Nq@lAyFfL{c@rRew@r_@i`BhQm`Ada@obF~^oeF`H{|@nWufDzJqxAVyD~N_xBxAyQx[c`EnCo^nIshApC}^hDiw@K_IYeUy@kp@hBwn@dFav@fJucB~PooCxDsb@tJg|AVc_@yBce@}@kG}J{u@}AaOgDag@gNw_CiFmdAyBcWyAoO{TuwA{g@kkD{Lat@wFub@qN}yAy@cMyLstAeDe_@mBaSiMupByD{s@oXskF}Xc{EsBga@s@_NeNewBk]c{D_b@suEoSsdCwAcO}Gwo@}CqV{Y}~@gIuViSuk@wAsDVaWqIym@mEm\\yc@k`D{Eg^cGy]e@oBmCiNk\\ghBeEmVgA_FiOev@g`@wuBc\\{aBcWsnAqMgm@iKgh@oA{FeBsIsKuj@kOst@aBiHuEuUqIma@e@mBsDsRgL_n@o@gD}D_TyU{hA}DyNeAaDxc@sWjXcI]kCEiFmAgJi@uDoIiU{@_E?_Gt@k]~@ep@v@kt@LqvAm@yLo@eG}Hw\\yEaWoEid@mNswA_Gsl@uQgjBsJsaAeZ_yCaA}JyAcN}CoYoAiLoL_gAiN}uA_Dq[e@wGsB_U_AeJ_@}DWsCeDy[aWkfCgC{VgEsi@oB}OuFqb@{T_lCaRmoBuf@kqEYqCw@aI_@uDaLylAoEi^_Hgg@i[}nBa_@ebBq^kmAcTsq@kXez@uAyDce@mnAqWcm@chAyzBoe@e~@ijAwcCaTa_@qS_d@sGgNkDaIiY{n@kk@saAip@sqAeEeFyl@muAyBuFemAkcCkdAwsBoH}NwUie@{NgZks@mrAei@ggAs`AmoBcuD}vHwI{PiB}D}kCmmFc{B_pEi|Ag{C_e@mfAq\\{aA_b@o_BaBmGyDcRu_@wiBcXgwB}Gsy@mJw`BuDqhAc@ik@MenBCiU{Cmo_@o@uuGAwJM}nACkRGas@Muo@?iB?gAeAwpOqAiwFKwhIo@_lCm@ioHaAueLe@}gAkAanH}@}jCTu~CmDsb\\{@}vIE_yCe@wtK?w@iAutPkCs`WAeR|@}}DYqc@Ikg@DihIN}ke@@epDBsdF?cSOsm@j@_\\d@wRz@{R~AqQtCkVhIkf@tDaSpNwp@dj@kuC~bBqxIxDi]fO}rAp@gGlo@kaFxYg|B~@mH~i@gdEpdAs|HngBe|MbGge@lGyUrJgYtO}UzYil@lLuXfEmMhDwMzCkOjCkOvCm]tBo[z@wOjCcp@lIkmBlCa\\|Kwh@tMqe@jJ_V|IgS~Qo]hKkSltAalClPy[p~AeyCn@mAbgA_rB|GiM|l@}iAnqD{nGvaAmcBrmAmcCttAa}B|o@e}@~iEs|FtyAgjBtb@gi@dpBcaCtYoX`jAqfAl@k@j~CmsDj_@uc@vV}Y`U_X~OgQpUiQzM{Iv|@{`@rV_O~KiL|HqM~HeVlF{SrJ}f@|HaPtLkNr@a@nJwFrYwIn{@kZdLsDbXuIpmAw\\zm@_OfV{GpU_LrDkDrEkE~[_[dL}JnEsD`FaDdDmClDmCbCuBbFcBhAaBfBw@|M_P|GwKb_@yi@n_@ij@f\\}e@dj@my@fk@cXjCoAbDaC|Rsb@|Qig@jFeTpJk`@tL}t@fC}i@vGyj@o@q@YaACgAReAh@w@v@a@nIi{@~@oGtAeBhCaBzGwBhG}Ahj@eMpX}G~TuFt`@uGxh@{Adl@@jA_A|AU|AVfQ}m@nWmpAdH_^hNao@rKmf@nCsP|X{~AlCuUvEsZ`DoTl@oC~F{VrGwXn@iCjDcN~Kyg@bNyj@pMsh@nC{JhGmSnFiOpJaO~QeQjK}AvJDdOeBtg@c_@xBoBpSwQj`Bq^f^~Ipj@xM|lAjZ`dAvUjd@fKnWrFfWhFhWrCpITrADlHV|MqA`ZgG|YkLdu@mVvLkFdh@aU~Lh@`|@gNtXwD|d@cDxh@cB|r@|C`jAa@dq@GbcA|Av_AYpo@b@`KFdz@m@rKDrpA`@fx@ItA?|lARfoAsBnBCzXc@fWuDnCcCnHyHxPs\\nkCm|Fr@_BdMmUnOgR`R_KvSyFhj@wS|QkG`PoF`gB}l@~X_LxZcI`JaCdC_@|E]~CK~DBpDd@hDt@x^fLnWlN`cAv]|JjDvShIn|@l_@h`Al^|WmC|MuBzu@sVjt@cRtj@mQxTsHbx@gXfNgEdVsItS_I|lDmdA``C}t@~Ai@`eAi\\xw@oVvMkEdjDyeArl@oRzyB}p@`Cu@fsAea@pUgIj^uKfaGwiBnGqBl~DcnAjiCsy@plAgf@tMsEn}@_[za@yMnc@yJp]cIb_AeDvm@nCriAzDtH`@dlBxJjH^flJ|`@bwBbJpeEoExYMlg@dAdhD~GnrCzFzYlA|pChL~v@`DhENxaApDtbArDvxAjF|iAhNnWrCp]`E`|Bx]bqBb[|yDdZrdAnF|mBnOdIHfSvCfkBbL~^|Dd{DxT`p@lA~aCrEdVF|pAVhtCl@j{Cn@hzH`BzxDbFjKNpl@o@no@ZjpB~@xPHf{BdAjtFjC`f@u@p_PuiBv|BmYlmG_v@~{@aJvcCa\\`cAsI~cAeOlbDy\\lD_@rtAwNt~@gK~sBiX`u@uG|v@j@dtEjDra@lAf`@\\~c@}@t^cH`i@wN~zAqa@pdBcf@fzAyf@n~JcoD~zBcu@rxPiyAzr@sJzQaCdfImfAdv@qInq@yHzZ{Ezm@wHxPoBzn@oGv_AgJjtAiNrGy@v{AqRliBsRvgFqg@v_XqjCxeC{UvtDy^v}Di^rnN_sAzn\\avC`pXyuBt|Ke|@jsE_]lsDqd@f~Ekp@pcFqq@d_Fqs@nvCc`@jbB{YjViE|d@gIfWwKhh@u_@byBq~An|@im@bgA{m@f`@}Sbn@c\\baBe|@djGaiDzRsKxAw@zfBw`AnbBg_ApuBofA`z@qg@jv@ea@vi@o[|g@iZlsA}r@vWiThUiUnM{NbSmZpSq^rd@_x@rjAopB`i@c}@~Wab@bVaWva@cXhPyGx\\qG|tRavAleD_a@zmNc`AbfHqy@jxHq_AzvAaQnkSqkCjaEsg@fmDcb@v~@eSlo@}Wfi@m^vj@c^jlAwv@ztBkuApuAq}@vzA}cA~oAmz@`mBsqA~]w\\rN_PxKkNdYeZbYw]t_@{e@jNoQvNsO~FyEjFyDbG_D|FeCzGeBzGgAhHu@xHqAvPr@fHp@fHnAlgAzZrc@jMpYlHdYrBt]hB`{@`@xrBmD|xAaGtvAqCr|@sDnoAaDtx@Mhb@h@bVbFjs@~Jh{@xSnf@dKhyBri@j^xNna@xIruBb_@`\\lD`i@jAvl@aAlpCoSxiHkb@~Hg@zaFa[`x@sBvoB_KloCgS|mAgFvr@kHtsAe`@haAk]rwAsj@ngBew@`rAuk@twAcr@bxA_p@~l@aOtXcJxa@cNjo@{Yh_@uK~NkDnqB}`@zQaEzZuHbxC}s@pJaC~sBwb@~l@wNn}A}\\rUyG`{@sf@jd@s]lt@uV~_A}AjHh@|a@eCvq@{Jr}Agg@p_AgXtHyBbbAiYz\\wMf[}L~cB}e@br@iSnfAm[ffAs[jXiJfZeKjDiAvtCi~@`RmIdpBel@xH}BnP_FjmByc@~Dg@byEcm@xDg@bmAeSnyBu^`mBcWfk@wHv]wE~jAgJ|VoBzyBeQxT}@rQs@tr@dAffBfIpH@lp@LlG@nnCyTvzByQ`s@}Flo@aQhKqDvlDqmAdzAeu@jg@_Wru@s_@fp@{\\tuAus@|\\eQdReP~zB{mBnc@u_@zh@e]jb@qTlp@yXzmByx@jUqJ~u@eYr`@yIfEu@x^mGx{AuQ~p@aInhC{WtNyApfCe^nZsJxYuLnc@gXfMqKnyAsnAd_Acu@dG{EvAiAjSiPzCaCnhB}zAfj@gc@`j@wf@xe@aj@l[u\\~mA_tA|RmRpf@i_@z]{Tx[mNxMeGpj@iSduAci@bd@mPdc@eOtfAs_@j[oI~[kGxQsBlU{AjoFaMf`A_AvnAeBveAFjeAnC~{CpNxq@~E|x@hEpnDjF|_H`D|kAmNdlAaVfbAaUrk@{K|d@wGl]Pj\\rDfz@lUt}@nXpd@hEvxAa@lwAKnm@jAnp@tCxs@vC|m@pCpm@|Crq@lArPAjPiAtb@uFjpAqSpo@oJj[qEpz@_Mb{@mMrf@_Jju@wFbSrAtf@|FdbAtLdxD`g@xpCj\\fRjCny@hI`CThk@tD|j@`EbDJtw@rCvOlAhXtBfUhB`y@nFf\\xB~a@|B~`@jC~H\\rbB|Hd\\GzSwDlsAak@~PkGzOqFjF_@l^wAtk@q@fCFng@rArHNh]tF`~@`Lj\\lFnOhCvRhCx[jGfnD~q@lyApK`m@vBn{@tD|Wp@fVrAnaAnDpg@`Adn@|EfaB~[vm@tMlv@xNj`Ehh@p{Bd]x~Bls@ls@zSva@dN|fArQ~\\`Hxp@lHrc@nA~fIuc@r{D`Or_@vBhn@tCzoAnGxu@vCnv@fHjy@vMzdBr\\vo@lNxdA|TveAnUtwBje@vr@tQrbBr]ttCxv@pyHh~BjtAla@boA~a@`r@|Rv^~MnpBl_An_@xNt^rKdZtFb^rDdn@tD`_AdKteCxd@~fC`i@`cAd^t]jNzWxL`RjHp}Ahb@x[hHlm@vJpd@nE|c@fE|k@`EvpAbEhvArErcAxDlv@fHts@hKvx@~Lt{@fMlhAzJvmAfJ`hAhBzkA`@feA_ClnA_F|k@sEnlAkNzaAyKlcB{Yl_AkQzi@wH~bAyIdeAcF|yA_DdjAZbuAwAnxA_@bxAQpl@I~lCa@tlCrAdrBh@liC`Bh`At@nb@n@rGXzi@K~|@sChv@wI`z@kLdoAoO`~AwWrY}DhS}EhSgHhg@ySrc@eSl[kMt[gJvtAoVnoA}RjYcClZwAna@_@|e@VvsArF`vBxLhRtBzhBrTzaAbLrxAhLzlB`KldCjBfr@|Bn_@rBbz@fFtdBjMhwAvJh{AhNziAfLb}@~LneAbQll@jRviCtfAxc@~UlmBjwA`bBxoAjc@x]ziAjnAp~@lgAhR~K~e@fPn}B~Yjl@zI|m@bJ`c@lNzZnR`]zXtx@bx@ry@jt@|c@lc@xChD|u@feAjG~I~D|FlPzL`DhDhLfN|KfNvMfOvJvKzKtKpOfOhMdKbO|JdKrGlN|HxLrGfi@zV`XdMrTfKrOdInPrHpRlKtNlItFtDbXxR~RfPbXzRd]bXbOdLdT`QdQ~MbO~LbUvQlQvMhMdKzNfLdN|JhMrHpO|H~JpDdKjDlPdE~FdAlI~AnHdAjWxE|bApOr{@lNtJdB`GdApLnAhKdAnQ~@bJZpQRzO_@rWiA~VaBr|AaJn{@oFrReBnZkDz[qDb]eDv^_Cdj@uC`k@eDxr@qElp@{Fxa@qDtSiBtmAaL~t@wGvg@kEz]eDl^uA~_@eB~j@qDdQeAlUo@~KQvM?jw@N`o@p@fMO~Qa@zLiAjO_BbOeCrPqDraA}UvrAe[rdAiVhgAyVvo@sMxf@kJ|T{E~RwE`f@uJfh@mJ|NqCpLiBdK}Ajo@aKtr@sKx`A}O~\\{Fnc@wIpUuEzXcGlGaBtGcBpxAmb@xq@eTdn@oQng@qNlp@_Nxf@eKd`@iKtf@cPjp@yRtf@sM`y@yStGcBjU_Hve@mNrSkIpo@mZt\\_Tb`@oVzf@e^hRkK~MeEnZkGjm@{Kpm@oJnh@sIbi@}F|zAyMddBkOb_@wEd\\_E~AQdAKhBc@pA_@lBi@dFoB|EqBdEcBzEgCpB_Bb@S^Kr@An@DdAVr@RfBbAPBR@P?N@p@Gp@EpA]xB{AhAq@lAk@lEeBhNeFvIyBlOoCx^cDz\\mAl[}A|XmDh`@kQrYoQx_@gYxpAsr@rR_Otk@sc@ja@kRhSqIl]oIfe@mMvx@aYzJ_Ebe@kRlf@_S|q@i`@j_@eRz\\}Tbr@in@pXgVvZaQlf@{SzX{Ml@YvXaHd`@cFxf@aIvjAkQ`Em@`\\}K|s@wZ~^{Sp`@cOpa@aNhDiA`Bi@ph@eRzUqIh[aLrEsAlMqD~R}A`NKP@jO\\X?|Kz@lVvBvn@d@v\\uF~U}B~f@qHB?hR{@xh@gA~Sc@``@kAd`@mAtp@g@bMAhSXft@nAhi@bAhJD|~@dCzM^`N@lTc@rTiArKgC|P}GzFyBjXqIvKeA`Lc@xCEp[Mdm@Cf[e@dUp@d~@~Bn`ArBpJBjJ?J}Af@sA|@_AjAg@tAIrARjAr@v@jAZzA?`BjIx@nYfChPpBhOz@hp@mDxVmAjAAnf@g@bEo@nC[zJiBhTqDpQeEf_@aJfDgAnCuAbC_BbEeD|DcDhD{BpCoAnFgBnIiCvJ}BzHkBpFs@xF}@pHWzGWhv@Ul_@C`IOvG[|O}@hb@gDpPuA~Ly@`TaApd@wB|@Izd@wEtm@gHfQiCbGcAbEcAfFwAjUkFvNwEny@_\\nD{Aze@yRfUkHvf@}Krx@aOth@mMp`@cL\\K`A[`RoG|VcKjNyGtNyGrCqAjBu@jE{@|Ca@`D_@~Nu@hAG|G]|XqBr^iBtQ_AhWiBnP{Ajm@oHjBQz}@aJlmAePnz@{Ifw@mGf@?|`@]tq@t@dPa@po@aBhy@sJrf@_HrMoB|e@_HbXcFnR{BlFm@dFk@hNcApc@{FtHqBpD{@`RqDzP{ClZeG~e@qIhFcAlXsFf_@sJfn@eT`a@sPzj@sVj}BkaAf}@w[rw@_WfIkCfDeA|NyExC_An[}KxU{IvHwCpMaHlOqKxGsGbQ{P|]}]`QmQlJcItKgHjFmCjL}DvM}CvNmBpMg@vP@ntA`B|m@HnVXpSa@xXwAhuAmJnLcAfOoApa@iDrR_DdVu@fb@yIpdAoWdDgAbtAgc@tk@uTdY_LbM}E`oAqf@dOkIx\\eRxz@oe@rn@aYjA[vNwD~`@gDf}@cI|SoBvPgEpMcI`F_DbXwUdq@sk@va@mSdiAgUz{Bub@xi@iKhnC_v@rm@iTvh@yRhTsGtA]pQeErCg@~RcDxWyBnp@}FnG_@jo@wBjz@wBla@i@x`@dAjx@`NjV|F~^bHjZfBtSWza@mB`e@_Djf@cEtc@cD`d@{Cht@kDni@s@jy@Kjn@_BbuB{]bf@{En]iAzIKdS]lHI~MGxgA{@vgAK`vCqChq@sBhq@f@|u@pB|u@Z~jBmAlk@i@nn@u@xn@{@~n@{@za@eAx^cCjq@wEbUy@bUA~p@XpQTpQzAfe@pIld@nIjh@zGxt@jF~o@~Exf@DjdAgAvf@uB|f@cDrg@{DflAwHxPaAtNYfN^n_@jDrU`FvTpElT|ChUz@lfAkBzb@qApb@Iv[hAdR`BdHn@rRrBzQjDxc@lKt}@zWppAtNt\\dBpb@G~[gDf[y@xa@hA~z@xH`sApL~uBd\\`XfApo@XneBwAdkAqEliAuLzfAsP|F{An~@yUfw@mRhAY~]}Nl^uQ~c@}\\zzBg`Bf~A{eA~{@ui@`QqMlsAacAjxAw~@nwCuvAzoBwaAd{@eb@`hAeo@p[sPhdCo|Af}@qe@n}@u\\hv@gY`ZcGlb@Y`}@~@j{@Xt_@wCvQwDbg@aMxcAm[n|Agm@rf@iQno@_YngA}j@txCuyAldDiyAh_Buo@ls@_Xji@aVfwA{n@vt@k\\le@mR`o@uYhTeJrw@}\\vq@m[diA}g@jB{@zxBibAdrAgm@zl@_WxU}IfMuB|L}ApRg@hYFdq@W`Q_@tPu@nLgAxOmCtb@eNdm@gQ`l@cRxJcD|LcG`NkJnSiRbo@_p@v|A}~A~YwZbR{OrXgQnW{Mjo@wYztAuh@tg@kS~i@qSd`Au]lpB}r@hqD{uAvoCebBho@u]th@}ZfNcHrK}DhCe@jHqAvNeAfLD|QvApd@tD|k@lEfS~@lUGbPw@~N_Bzn@cIniC{[`sBqVvl@uGv[kFxd@sJn`Ds|@thFiyAnuBgl@fu@ySlZ}Fn]uDxvD{Njd@kB~[wDjPuD`[{I~N{Ffp@_YvdAkd@pCiAxZgLxIsDz}BsbAfe@ySx_@oR~f@_[vg@qf@bYe]vZmf@xl@whAxRa[fNkNnJ{GpYcNru@}ZtaCicAbUeJxqC_iAzkDg|Av]_O|_@gNl]iIpo@yNfe@kK~Z}KpXuSf_EmnDdi@{e@bk@qi@te@}r@baAqzAxYsXbR}Iv[mIfX}B|U{BxtAmMfkA_Lvh@iDxRmC`LoBpP{H~QiPxHiN~GiUzBmYA{lDfCci@pGe_@lHeZvVomAfx@a}C`z@ymD~FqZnCuZx@cf@tAcoAtDwhCtCqXxDyPnLu[vLoTfQaWl^{c@jY_^vfAouAllAaxApvCipDtTqVrTu_@`Y}s@bjBe{EnkA{sDvNoVfOiO~}@mt@bDkChEkDpdAaz@~eD_nChEsDpEuD|bAuz@`Y{UbsA}hAvPsJrGsDpY}Eng@gC`r@{EnK{AhZkGj[gNvrAom@|fCelAvmF}fC|a@_LrX}E~g@mBzyA_BhcD}DzjCaDdUbAr~GoF~`DgCpl@sBjvA{IxpEqWf\\cArXOhw@jFr~AlK~_Hxf@jNzCh|Cvp@bbAhYdxF|jApUdD~N}@bSwClKyBhRyCbmAiJb~AwI|Tc@~N^pJbB|T~DbzFxhAhd@nIbh@pIjVZpJ}AxKcE`K}G|IcJnFcH~G}Mfc@iiAlQad@fa@qdAxLiRlTuYjXqSzP_JdIkEli@eVlmEqvBpRwKhTmPlPaRjWed@jv@yyBbKqYbAuCfN}]vS{`@pc@qi@~y@oaA`Sw[hp@a`AnOwQltBwdCho@{{@bm@ap@~w@gv@hbB{}ArPaOfM}Kf]mVbg@cW~tA{m@vn@uSjPaDv@O`a@}ChgAoItLw@zDYrRuAzd@}Dx[iJ|QmJlY_Tlt@wo@lTue@fJcXx\\ai@fHoWtJ}Xv[s~@r~AonEnO}g@|M{X~KyMpCcDdOaNpOqJba@aUze@aXv_@cTleA}h@dr@cZbi@mR~o@oPvs@oGxFi@viCmRn_Es^d~CeYftBgKtr@s@tZWjQOli@e@xmBoApmCeCfq@o@zxDgAvwDrBtUS|QwAx[uDzEeAlPuDl]{Knh@wP~vD}jAdjAea@h_Byj@hsCiaAfGqDxGeDjEmHv@oAbBwD\\uBH}D}AyZi@yLH}J]wDsEsx@iCyb@yIiyAuDsn@uC}e@eIuvAcIgdAwKi{@{DsZeH_r@}BuTieB_zNu@qG{L_fA{KgjAwYqdCyBcR{PyzAcm@c_GmAul@R}T@gW^{VXsV~Ai\\|Ae^v@mPb@}I`@eIbAoTrAy\\dB{c@jAcRJoB`BcWhDyj@vCuf@bCqb@z@sQVeF~Ay\\rAuWzH}rAtF_cAvCcc@zCqh@pEeu@f@uKj@iM`AqTlCch@nDaf@tB}X`B}NlL}h@bJs\\dAcDvO{e@vJsYrLy[dC{GdHkTbHkThNua@fPwh@nKe]tJm]lMsg@lCwJ~Jo^fFmT`CcKzGo[`Sw{@`BgHbMui@v@kDfPat@lPqt@hUccAFm@fH}[tHiZjp@{xCtk@seC|UedAlc@}bB|CoLjUk|@vEwO`^elAvq@}|BvwBmaG|tAybDbtAovCdmDqrGhM}RprAisB~Xqm@`Ngg@fQ}}@f]m{B`PufA|WmiB|PmmAlHoa@jFa_@bXe`Bt]}pBdA}Fl]khBZ}AzLgr@lAwHLu@PeAnLit@bNyw@dVwaB|SsuAhCiR|Key@r[wtCp_@wrCbL_}@bPuoAhHe_@lHm[~Kcf@lh@}tBxJma@fR_y@rF_]bGqg@x@}PVcRh@eY@kl@eAq`ACiIQuSAiPz@cWa@yVaBqz@cA}}@c@aeD{AqjDKmqBYukAEe_ARio@p@u`@HeCP{If@yOt@oGfAcGnD}WpGe_@|DyWrIwi@lJel@tCqKhCmG|BkHlBePxHaq@pDg_@`Feh@jI_w@jMcqA`Eof@bEq_@pFkn@lHymAbEsp@vEkmAdEw}Af@wZhCg{Bd@qhASkc@j@kxDWgEsAeIMaFgB}~BeDkmD{HopCiAiy@a@ua@uM{gBcLgiA}TagBwNieAuYggB{^_yBwg@yfD}KibAuIkx@gKmkAmHaw@kAcJeBeHuBaDcBo@uAgAaA_Bi@kBQuBJuBb@oB|@cBhBwIx@yI^_NZaV?wf@[}g@w@kd@eFwgAwEal@gA_LyFip@wHoq@}G{_@sHk`@eFyXqAgGcJwe@oBgMoAqLoC_TkAkH_A_Io@uFqAwO}Hgv@yN{kBgGaw@uNawBsHkeAyDap@aE_r@iEcp@}D{t@aH}|AiFizAmFuaBcCw{@yBq{@_A{g@iAik@_@i`@Wei@eAafAgAeuA[q{AZu~AvAieCdAocA~AsgAbBaeAxBybA|A_l@rBen@zFszA`Cam@rDsu@`Hu}AhDak@pDwp@`Kw~AdYiyDpF}l@|F}o@hOsxAjNquAzPgxAjIkr@tJ{z@vImp@jJoq@|Jgw@`Kov@pMweAbMw`AvHyl@tJkr@bM{q@vDeYhEwX`Piv@~Ogm@nOuf@f]c}@fRe`@`Umc@~So]hVg^dGcJvAeCbAiCj@_CRcCGgCXcCx@uBvA_BhBaAxB]xCcBjCyBvGsIxGgJ|g@egAbgBa|DtSwk@lOok@jQ{eA|PmaAj^usB`f@krClUcpAzw@g{EpXqtB|RgfBdNy}ArOa}B|IgjBtEgrApCayBpCclClDwnGrAkgBdHcbFpCkXdKcp@pSct@xs@ekBz|C{pH`xBguFfv@chBbyAilD~`BwzDxpBedEzcA{hBte@oz@|aC_lDvoCkhDbzD{mE|nEaaFptDatE~tCylE`q@{jA|eBilDl}A_lDfqA_uD~m@wpBdl@{uBlF{S|fAcrEhKe]lSed@pOyXfOeYlU_Y|UqWlD}BnH{EjIoFhP{KvLwHdOcFnc@uV|tCi`BvPaNd[_SnDcBbFqClAk@^SbBo@lAi@rFqBxCeAnQaFfYiA~ViCxQcAhXy@d\\_@pSwArq@Dno@z@lIGpe@Tt\\PrwBq@`bBR|~Aq@jd@Oll@Zbh@~Bpa@hE|`A|Ijq@fIjl@tEvQn@rNDda@_AhlBgKlqB_IxJi@xHWtEMlGGvF?nDLtDLxAJx@BhEXdD\\xE|@lHtB`RlGzl@jWzkAtg@t^fQjXvJri@lLl~A|Qrl@xMxn@xRdsA~d@ljAtZ~{A|]dd@bHpSbB`TpA~]vCna@xEjf@pEb|AdLnKjA~i@~D|AN~eDpUt`@lB``@lCt~@xJdeAzFtz@vDlyBvDz_AbCj}@bGvi@lCj|@~Fnv@pI~b@lE~NxA|WrDxU|CnYpE|Q~@t\\Th[q@|]uEhA{AlA}@xAa@~ACxAZxC{BnTeMhK_I|Q}OpPcP~ReXfRm^`La\\xKwe@lOu}@bR}dAxEaYxIal@bMyn@jKy`@pKc\\tPug@dRud@r^}r@`f@ev@dw@uhA|~@qsAjnBknC|}@moAtz@cjAfTmRtUgQpWuLl_@uMvYkFnYuDb{@eEjnBgGjp@uBj{AcFxkDoJrfHoSb|@kDdo@eB~l@oA|l@{Alx@yBxv@cEde@}Cff@eDnq@kHjq@mHzn@uGxl@_HxmBoQ`{@{J~y@kJd_@gErTaCvMkBtLkDrNyFrNsHjEiE`G}FxE{GpDgG|HiObJiYfHy^dTcaAbTecArVyfAl_@agBt[wzAzFm[zCwMhD}LfNqe@`H}RrHiPdNcYhNaYrK{TdKiU|Ysi@pXgh@zHwNtFmKnKiN~MqMbJyHv\\aUb[mNtw@ai@vf@ub@lf@k_@`pCa{BtrCmeCn@kDzAsPvDa[zV}p@p^ebA|V}m@|n@g_Bfl@o_Bvn@iiBvCcVlA_\\jBciA~@u_Af@aHt@wFzBaOxDcPzEeQnGkPfHuZfCiJj@cCj@yDV{EHeDLsCHmEJ_D@oCDoAHqBNyA^sBtAyDnD{HfKySxg@wdArg@e~@z@SvLiE|g@wFdPaBdAO`i@uH~X_Fvo@eLpHqAzmAwSpg@{Hl_@mEl_@qCbk@qDrnAkIbfBkK~bBaN|aAkHpP{AtIw@fKiApVgJvW}Udi@}c@hoA{|@~iAcz@`xAy_AztBg~@fhCijA~aDgvA`eAyd@hl@w\\fo@ke@rh@mc@paA_{@~u@yo@zx@ao@xz@ee@b|@aa@nVaMhQmNxLeRhVi`@v_@ew@tk@ehAf`CgwEduGwgMlaCcsEn|@q`Bjy@idBd]{r@|R}XvTaTlVwNxSoIzIcCdMkC~HqAdGa@fQiAnuD_IteHePjs@kDx|Hk_@jeCyIthFaM||BoGve@uAlF]zQiBlT}Dnd@gQpf@wR|\\cPtMgGzAo@p\\}N`OsGvOoEnJ}BbI}AvJkAbKq@`N]hgCtDjY`@~nAt@hs@OrBGrRq@dJo@fK{AhLcBjFqArDkAnHuBjJkDtJkEtLmGzHsG|IeHhWuSpPgPvPcS`IaQnJiWjg@ccB~Zo~@rqAw~DrZybA|Q_o@nMgb@xD}L|Qwk@vTmt@nHiWxPcn@d[kvA~j@maDpIef@rQybAvJ_j@lNuv@vJgf@fFaTdGyPzFeKlH}MpGuHhP{RpVmV|GaGpKsNtPuV`MkW|K{\\rXe_AvSaq@xLwc@nFsStKc[tKu^nLuj@lIui@hC}j@xRufBpJgf@bKe`@xKkb@fGoQ`JaRrLcUv^ei@vf@{n@bFcGrWk[bOoOnO{LhYsRr[oTpb@c^bRcXvKcSxMu\\zLcZtYqt@phAe~Bzl@cbAtGcJdQeVhVkZvgByeC|GcJbh@mr@vf@up@be@_r@tRg]zHkPdGuMnIyUhFoRlDwN|DiW`DmVbCaPnKoq@fK{n@tLez@nHue@dIig@|Fy\\lGq[bIc[|ImYnIsPpJuSbZem@niAgzBpY_h@rN{WlIqOlVqe@vT_c@|MqVxMwWvFgOpFcPnMmm@jE}YpEue@pEsp@|Cyn@`B_]jHseBpK}lCfHydBj^ggD|Ewf@jLabA|Gac@dI_]vMed@xVcs@dIuSxkBc~ExeAqtCnGaN`S_j@lSii@bSoh@p[uw@nM{YbKgUvRq[zYec@tCsDhGyJfH}JjLaNpXi\\fM}PlPmUfPu[~DcLn@yC^sCJgDJyHD_IDgIGao@i@wx@?ul@wAqoANe`@gBgr@_@eOo@_XaAsZeCo[SkYlA}m@|DgbAfDoy@x@_\\dGywAzD}fApCuo@tDq|@~Awj@dCm_A^unA?yFgBwaAsAql@uC}i@{Bqd@aGaxAQkE_Cqh@eHwrAaGilA_Bg`@s@wQKae@v@e[vDa^hGoc@|B_Nf@uCZuAv@kE`@sFLyAXmBPoBj@eDr@uCpAcKnAkHf@wEv@mHbCcPbAcIb@mEZ}BNmDN{BFsDJkCNeDDgDJ_FNsZm@wj@{@_v@{Au_@gDss@kCy^cDsh@}Ia`AyJwo@sJyo@aGo_@cAmFuKqp@eEa[sWyoBiImz@yDqv@qAq_AGq]FuYTiQ~@cNpB{MnHy^lKmb@ld@ocB~S}s@p[e}@bb@qy@p_@au@rLgV~a@iw@bHeLtQ{TrLoMxLkLdFiE~D}Czs@ue@jo@u_@dMwIrTiSdK_MbDcGzCqFrDaHrF}OnDyN|BmN|AwNpCk^bHqdAlB}^hCah@h@wR|@e_@jAub@vCkf@t@gJ`NujA~Hc_A~Di}@vBadB`@ie@Isp@}Aqc@MmEcAw]sAq}A]oVkAm`@w@se@kA_c@qCmq@u@uNc@wJOmJM{HImJDiHRaHRoD^mHbBeI~@}EzCgJfA_EdBgFv@kBjA{BzHyNpJmOtCeEbJsMhGsKneAoxApl@ms@zhA_tA`nAcxAhgA}vAjM{Nve@yi@vVuXrIsIdLcL~_@{a@tTeShVsTz^m\\fFiFmBcDgAoBkA_CgByDs@gDa@iCI_B@}CXwCrB_IzA_FfB}GhDwOlIse@~DkUxBqInCiJ|BeG`CiFjCsDzDsEbDmCxGwEpp@aUvmAyb@j{Bcs@df@qOrsBcp@|eBcn@hCwCzT{VrJ{^fH}x@zO}mAxDulA}E}_AoBot@`Cub@jLy^nTsThh@oQje@kLjWkL|l@y^bg@_[b`@a[lWy^p_@oXtg@gNzy@qNfEs@dyAmV|d@gEp]uB`Y{XhPuJls@{YzQ_KxRmK~N}I~NiNhJ{JdJoL`\\en@~[gn@jNsXhIgPrXyi@f_@c{@tb@q|@bU}`@bLiNjJuJtUeRb~@gp@jcAmr@t`@_WbKqI|EqF|EeG|GwL~FgMbFmO~GwL`d@mr@pc@ev@pYuf@tCaF|_@wo@~Zgj@~Xmk@tYyp@zKcO`Uee@~Qge@|Oo]xDqIvZwk@`y@yuAtZqm@jy@m{AjLs\\vJkb@nPqhA~Akc@jN{ZlVwb@b}@czA~g@a|@`c@ip@lYy_@~mFkgH``@si@znAcfBdh@et@|O{TfbAutAt^om@bOob@x\\}oA~|B{lI`IyXlGySnZ_fAnNgf@na@axA|_@gcAj@{@`CkDz[ywAjA{Gxx@}mDrBaJ`EwOlD{SzC}SrAuQb@mRGuNW_OaA{_@kDc_AgIorCwCetAeBmcAgA_x@z@oTpC{XpBeMbHiUzSge@|e@gw@rwBosDreAshBh~@siBfj@wgAvZum@r}ByoEb`@su@xwAquCdh@wdA`FoM~D}KtEyMtG__@nDq_@nG_aAnF{~@nCg_@jCe]|AeY|JgvAlNkzB~Fc{@~D{g@nEwl@rGku@tCe`@hD}j@`AgMxIyiAdI}jAlPcyA~Lew@`Hg`@jOm{@nQkdAtCySx@iOpLis@hQm}@jSwp@rJgVxK_Pzh@qv@byA_jBnpAwmAjz@_}@l_@ab@|fBghBv{EqzEnh@ih@rx@yx@~g@uk@|@aAne@_r@hIsLxd@mq@jm@ir@fe@ej@rD{Dbr@}v@|O}T`NsU`I{NbHoObMw[~Pws@pRcfArOsu@~GoVfMg\\hKiRbn@ehArfBwuCzfAqeBhSyZ~a@wj@hW}WtSgUj_@k`@tcAw_Aph@qg@tb@sb@blBkfBlDmCd`As}@rQkVtNeZvFuRnCaP`B_WY_k@a@ojAk@a~AEqOFadAaE{{C|EaMsBs~@Hgc@tAmh@rC__@`BaT~Gkn@x@yHRkB~a@{xA`Jc]`BeGrIqXfI}WfS}o@tJm[xD_Mfl@qvBvM}e@|EeRjDiN|j@}yBlJ__@lVit@|RwTxMsPbWmS|W{M`b@_JxoAgZhS{EzeCat@xBo@z`C_q@fQ}EreAcV`o@kM|NgCpYkE~_A_LlYqDjnBoVnx@eJraBkSj{AaRfw@kLnK}AdXiL~V}NtQoTfOoT~Tqg@z@kBzL_iAh`@qlCzFo`@rTyzA~_@o}Bji@saDvHu]z@kG`@oCjB{NdIu}@rMq~BrUuqD|LqrBdIsp@dIm^bYckAfXkbAns@{cCvVcu@t@sCvf@quAp~@g_Cj\\uv@lYiu@baAcyBjIaOpC{E~D}GfLeQf@u@vJ_Oz[qb@fgAqtApV{Z~eAwqAleAuvAvYu`@`Zaa@vjAicBb}@owA|X{c@feAccBjKsPnt@giAdMqRhw@cmA`m@}{@vU_YnJsJxLiI~_@kWvJgEhKmEnVsJtBy@pYkHrh@eHhTuG~^cVbTgX|NwXjgAwhDvyBq|FfS}h@p}@u`CnDkJ`q@ywAbOs[hNoRlMsJlMgExBu@|`@kJhf@sKzzAk[zd@qL`c@kNxPkLb{@al@p[}_@|iAgbBvPqVju@afA|nAycBjE_GrVy\\frAaiBjTuVlYkR`ZoK`~@iW~mBch@riEi{@tY{EnkAsUjn@oMzlAwVvb@yFhf@qN|e@{Q|`Ccy@f`AwSrj@wKfaAyB`IWb|@_Bv\\e@`TYzZgEpWmIbM}I|L}PhGuSlG{c@hH_m@vOgqAx[yoC`HyxBrI{mCxGmr@r@e{@t@i}@z@}JzAuNfNerAdHyPtt@yv@bmBsv@xQoHp`Bkp@ti@yNvqAaYjm@sWja@wItd@yJbtAs[h{@qWhkB}c@dLuE|QcNlIcIhYybArm@a{AjQs`@xp@aaAvMkVda@qm@pUe_@~bCk{DvSuSz[ySpyDeeBnFkCzvAqr@|E_C~hC}oAhdBut@bf@k\\~w@}s@hCaCtP}NvZuXpKsJfQ{O|PkPlj@uc@tXuYfCcDdn@ix@pLgNjJoI|IsGz\\oSbs@gg@jo@}i@zw@cv@~CoDj_BewAprBmlBf]}ZnKqJjrCegCvc@w^tCcCbgAm}@l_@aa@nXeZ`LmKnRqNpB{AVQhMyIx_A}t@zi@ad@`Z_UhDaCfNkTfMaMvIqIjd@yc@vEeHnDaHxGiOhBmIvC_NlBoWi@q]]aGw@}Fad@mzAkNce@mTgt@gYg`Ayf@waBgd@wzAoGwW}HiTkJy]iHeb@kFgc@q@aSL_UO_rAn@eYVwR_Aaj@aDof@yBac@mBgb@cAwc@MoOVyO`AcTpAgRlAmMdDiTpDoLpEyInDsFpPgPzOgKxKsEhOsDre@wIvg@}FhNaBfv@cJtW{Dn_@{FvTiDtWiG`U{GtOqFpg@qQfc@iPt^eNbbAg]fs@uXxPiGnJaDx@K~HmAvGu@bzA}IlZyAlEo@tGcAbk@mHdz@uKzOgFrPcHv_@cV`OqCh|@iEdToAhQsDbVeJp[sQ~PgJrYkGj_@UpbAiBfl@g@lWqArRuAnX}C~KaBpQ{Ehn@ePrIcC~DyApEqC`J{IbIwHxF{E`TiNfIcFdHqF`FkGdNqPxHcJdFoErI{D~UqKfj@wQrVaG`SwC`L_B|HmAnPq@rWQ~Lc@zOaBzQyDpJoC|~@o_@pSiDlZvA~c@xFfLxAdUlAhi@uGftAy\\lZN~[xC~`@zLlUbN|c@t`@~PjLnJfE|SvFfZpAtuB|EbFKn`@oBdUsCbU}Ehn@eN|eA{Lnu@oF|k@k@~eB_D~\\UfI\\dItAxN|Dpk@hWhLxCpQjG`NvFnWtIpOzEvDdA~KjCpKnA|LzAxJ^~LHhRu@nL_BdLuBnXaFtY}F~AUzU}FbK{B`HyA~FeBfS{EpO_BdUwAdECnRKlReBdL_BfBg@hJgCfWmJt[aOdImCjEcBxLoCbQeBfIq@~QUfPmCfMcBtIc@bm@yBlb@mCdYgA|B?ld@JjYJpSFl_@o@pY}AbY_F~f@{LrWsBfj@_F`z@{M~j@iJpSuDnp@qHjwAmLtd@cNze@uR~QuH|bAgf@fvBy`At`Auc@xxAqq@dzAyr@bmAkj@vp@}Xl`@qQju@_^r`@}PlVoKl`Aqd@dn@wXlu@{]zn@gZ|p@oYtq@uZ|o@_Zbm@eYzq@kZbUaLly@g\\tt@}\\nx@o^tv@i^pq@_[lh@cTl]cNvGeBfL}AlJ_AbKc@d]iCvYwB`X}B~\\wBzb@oDb^iDdYkBff@uDxSwAvZkCfd@uBph@yD`[wBl]wC`[_Cre@eDfYeC|l@iEv{@aFbsAuJbObRreAvrAfiAxsAhq@`y@dNhR`^r`@trA``BhmEjnFxsB`dCrNvRx`ExaFpjA~tAbnAt{Ad~@hhAn|@veApSzVduAvdB~u@baAj_BtnBtSnX|l@`t@pRpUlp@`x@zDfFbJbLr[p`@l_@hc@`_@ne@ny@xbAj]tb@|l@hu@hw@j_A|b@d`@ta@rZjQrL|^lR~u@`a@~XdMd^bPvqAhl@dg@xUd`Abb@vLlFhpAbl@hV|Kh{BlcAbo@`Z`IkJlLmLdCkBzIcLrJgKrg@if@fn@ul@fl@ah@zl@mh@`[yWtAkAjBwAdi@{d@rv@ir@fZ{Xzd@wa@vO}KfS_K~lAmh@fyAsl@`EyBjn@oX`lA_h@roAej@tvA}m@jWkKzhAwe@lQiIzpAij@~u@o[vn@iW~h@wTrXgL|P_IffDctAzi@_Thl@}WhQsH|nA{h@la@aRlb@_Wf\\uXdZq]`Zic@hS{`@lPy[`a@sjAbQwd@jQy`@zNaX~b@cv@jc@oz@hwAg|B~OwWde@{|@nj@}}@n]gk@tx@uxA`LmTfh@ejAblAodCt`A{iBrh@mfAdb@_x@|FiK`^ur@tcAamBhXm[b@e@lAoAdVaPfb@}UdoIgrDvn@oYhJyDjCeAxMuFpq@cYxp@{XfU}Hta@qQlRkI~PgJdSuL|MuLpP{Pp[_c@jJuLvi@eo@ps@_}@l^sh@dyA{oBlu@o`AzQyRlRuNrPcKpWcO|gEc}Bho@m\\jo@q]xa@oThd@uWt{@if@fJ}FpWgP~_@sS`g@sWfm@u\\zm@{]fFyDzWaRtvA{g@tw@o_@ji@k[``@sHb[qI`SiDtTmA|Yc@jYm@zUwA`Vi@v[oBpPeBxSqC`QgFj]uJ|]yLh_@yMt}@uZz}@oZb{@i\\|SqHhh@cRjTsJnXyKll@_Vt{Aum@bv@u\\t^kQfJ}GfKuIrOeQlMkQfHsLvFuMpG_QtGmPfFuOdSck@rV{n@pX}x@hZkv@|Qy^jPiWdAaBhb@{j@vd@gl@lf@um@z~AiqBfw@}}@nPoQbhAqdAde@oc@pcBe|AppBkoBpHcHno@_m@x[mZf]y[rl@}i@`[qXx`@s_@dg@sg@thA}fAby@uv@vn@ol@`IqHxt@mq@bEyCva@gYv|@ka@~IoE~ZoM~`@aQvlAek@xcByu@jx@q_@tr@iZpFyBps@y]ddBmu@va@oQtGsC~j@sWho@{YzVgLv|A{s@r`@sR~QqM`IgFdb@ue@hDiExz@kfAvKcNze@{l@jf@cj@fq@}x@xSeVpWqZd`@kf@pUyZpG}IrDgFrm@kt@p~AikBxbBosBrOoRvoAk{At_@{d@nh@_o@~D_FlOmSlg@uj@l~@}jAhiBa{BfvBilCdQoRhJ{Lph@cr@lu@e}@zFyGjm@es@tX_\\hDaEh\\u^nS_VrwBgfCdhAmmAdeAumAj~@ycArz@o`Atw@m|@hv@k~@vr@ew@vp@wx@zOsPvSsT||@{dAj`@sa@`NaPtM_RlQ}^vFsOzEgQzEoXrAoMr@eNd@__@i@c`@m@eXwEabBgEscBeKw~GaJapDqHqyCqDevAgD{hAiDmdAsNqlDiF}x@oDsk@sAmi@HmTg@{Tg@aHsDsdAmHgtAe@kDmC{U{BkNsImd@{Okw@aSsw@yZobAuOqg@__@ojAWwA{AcIcAaOGy[d@iYnC}b@lEuQbHiRtk@_}AnQ{e@`x@onBv{@kzBfYo{@`_@{mA`d@eaBfJie@~D{i@n@iv@?uv@gBocAsFwcAyBgf@_Dof@aA}N_Cwh@{DiqFs@mrAy@s}AiEqgBiBg^uDst@sDkt@{KmtBmG_hAu@aaAw@miA?mSBww@eBspDsDkmBgK}oBSiFkDeaA?eDWgpAVonCBcMsFctCGwA_Je~BaAqWcAe|@gAo_A]aYWeReAsr@_Bmq@uEq|@yDml@yBcYgBeWiEuh@aJkeAeEqf@oAm`@d@ut@hCok@lF{y@zFo`ApAsUnDuc@~Iis@dNe|@PcBxDm_@p@uRIqSmM}~B_@qcAQkh@fEqrGLsQqIctEkIwpA}`@_eEmAeM_Eg`@}Tk`Cg[qqCwEyi@sHm{@kK}cAoIogAgFcq@eQeoAcVuoAod@ucCoB{JiDsQaEcUY{AmDiWg@sDgBkXEkNt@o]jCcYxFqYnRqu@vC_Ll[mlAzBuIlH{ZbEiRdJmTdGmLnBwDfhAiqA|_@ua@|`@cc@|^wb@z]cf@nu@m|@rWwZlA_BvqAydB~mDcuE`tAwhBraAyqA`Y_s@h[yyA~Lgp@nL_s@~Mav@~Oy|@nXsqAzYcbBti@kdDvd@adCjp@ybChWa~@nXu~@pGaXxL_g@~n@wwBXaAfSsu@vT}s@hc@cnAtKiUhg@efAhN}YnJiStTa`@|h@}gApZ}n@bs@iuAne@a~@~GuM|sAygC`g@{eAdb@{eAlcAi{CvLgf@hEak@Cif@Cyl@sCo~AfDu`BnBsw@dGsdCzFaaDs@cnAIyP[wg@cGyfDmB}fAiKkxCeBkf@aEkjAsD{rBUm}@bFs|AvLyyBzDgfBNcGhFoaDxDqnAxKokAnR{lAjBa\\]c\\wDyi@sFsh@aGeg@iDqg@q@wc@@ac@WkkBCyqBbBisBlBuf@`L}uA~OcuA`i@s`Dfd@ejCdCsOpEi[pGyc@x@}F`Fa^zMs`AbEy\\jDyWb@aCpKkt@tN}~@fFma@pb@ugDpj@qtElBwRzA}O`Fqx@b@qGdDyi@XeJdFyaBnKqxAnDuf@LqNz@we@UsdAjBsw@tJctCfDwx@b@uKdGo{@lAsJtPgkApW}cBfc@mwCzE_\\fq@}iEbr@{}EdGmb@~BoOzYcnBbKeaAdAeNnAyPBUdBm^~BsRjAkIbCeQjGwYbLoh@pDuPd_@{dBrW}lAlWslArCwLbIw\\p^ulAxOco@jKgb@~]gwAlYgjAbEqQ|B{Il]gxApOop@~eA_sEbKyq@pCgWdCmVt@cMHcFt@}TY}n@mAch@gCg`@aBeNeFq`@yDoZyAyNsOkwAwDe]cBs][}OTwO|@y]pB_RpFkg@rH_\\z~@}cE~Rkw@f|@okDxFoRtA}EfF{Q|gAqmD`Voo@fcCilGrZiq@dKqWhGoShEgUlCoVrAcZwCenA}D{lAwCktAx@cTvA}RjD}VvIw]`GyPlIyRrn@etAvTkk@l@eDpKuk@~Ak\\H{^gB}k@OqBoGg_AqAiRmUagBqNcgAiEe\\mOwdBuD{aBWq~Eq@uv@eCwvCMkNPgW~@eSpA{QtVipB|Jgu@dBwYj@u\\y@i[uBm\\aAyl@}@wi@UqNrCwqDtCysD`@ynAtFyg@lJef@di@gmAzs@ulAr|@wsAt@}@nA{An^wc@`PgT~KsUrN_b@vIwn@|AmzAXuYxAcdBfDoyBdDcvBnBixARy]UeHk@kHaAeKU}BeEsVqGgV_Xsy@q@oBaJcZ}G{YaQ_v@iNox@}BoYeBee@lEa~@nNax@pB}KjJmj@xFuh@pE}fA`Cc~@JeDzBgmBQyk@_Awk@sA_h@qFmsBwLebEo@wNaGasAaBw]cAoT{D}n@qF{m@wFak@_A}NyOonBuBiY_@oKIyI@em@@uXb@mc@FsIPs\\W{b@Bke@k@gc@{Acj@OyFi@mLOcBsAmc@k@eb@Eo[jBgk@nAmTt@mWIiOUsKEcCw@}PiBeToMkgA{OcqAaWenBeFcb@yAuRiAqMM_KX_SZ_JbBoQlHo_@hPyu@~Keh@vFsWhLcf@vEyRvKmd@dIwZvCqL|GaXfHwYbGeTxGgQfGcNhImNfN}UrP{U|IyHlGeFhKoIjF}DnWwMlG{C`JkDzDyA`KuDvFuAvKuBfVwEzq@wNloA{X~NcD~Q_EvMsCtxCeo@~}@iRndA}TjnBmb@rQ}Dhv@yOru@_P||@eRx~@ySb{@sQnw@qOhXaGrKeDlIiCrVaIvpF{sBrx@yZxeBwp@~s@qXxn@mVzgAgb@|kAid@ddAc_@|Ai@pe@uOpYiJpWoIvh@yPfgDufAj`@yLlcA_\\veAy\\tiC_z@taAoZn}Aqg@hyAyc@vfAa[pf@oKz_@aIlr@eP|{A_^zXaJbXwL|SqMrRaOtNyNrNiRvLwOlJiRdQaa@j\\u}@vKgZ~Xkv@`m@i_B~Xsw@pO}b@`Pec@|JiXr]w~@f[{y@b^ibAl]a~@zx@exB`\\k{@zRsg@tLgZxDcKnJ{TjTeh@|S{e@r^ax@ly@okBlG}NpxAyfDjz@umBp]ueApIe^nI_a@~I_t@nBsOtLwlBzd@}nHxj@kiJf@cJfBg\\tRssCpg@}lIfAmPfGc~@zMgqBjD}WtDiUlCeOnAcH~F_WzIm\\~Lq]lY_l@zUa`@rVe\\tt@sv@tz@a_A~\\e_@j_CceCdR}RzfBckBzFgGfCmCro@ur@rp@{r@dm@ko@xx@}{@xFcG`UyU~b@wd@hg@ee@t]yXfHyFnv@cm@ht@kk@x`BgnAj]sWvLcJxzAmiAvQ{NpjBwyA`cB_qAfiDojClfBiuAhp@{f@rKcIj_@eYxf@y^ps@wk@xViS|e@_]tb@}[dTuNlz@_i@lMqGx`Akm@nbDarBxy@wg@|eAcp@p}C}nBlUyNvg@i[rc@uXhhAyr@ppAst@|C{AbFiCvkAul@leCanA~i@kXzUoLdn@mZpPeItLwF|GcDxXqNvn@mWjt@yW`c@cR|c@cUfcAoo@hh@eb@bd@qa@p^yW`TsObd@}ZfIoFpx@wi@fZ_RrUyNzvAw{@dGoEr{B}`BvgAor@t`@{Sbd@aV`Ai@hY_OxYwNtFaDpPsJtY_SbQ{N`}@ku@bd@_b@zuAypAvrBaiBhXeVtBkBx]a[hLcKlhD{vC|m@qe@jeBqhAvVaOtgC{wAd}Ae}@jcA_p@pjDaiClcDuiB~mAun@deA_b@d_@mLj\\gKhxCkw@nu@sPlh@kLzdAwSzDw@zqAmUbCc@`q@cLnt@gOlBa@`ZeJ~r@gZxaBat@ho@sYpnAyj@vlAai@`m@qQhh@_LhNyCnuD}o@~mAqOfDYjcCeR~SqBtj@wHzfAkVpmB{c@hbCuj@bFiAnk@{M~t@sTxNkEzk@gS`a@gM|[sLpTqKjSqMvW}StPyRr{A_yB`]kg@fJsInSkRz@w@d`@sVxp@{XjmAcg@dp@gXlz@m]tq@eY~gBmu@dIgDzv@_Xl_@uJzb@{KbL}BdV}E`O{C``@oJ`b@sNtg@gTbr@_]~t@w[x[aOjXqIv[uFfOuAbo@}FflByMrmB{NzlAkJjaBkMtg@}DbTaBbzBuFhwAoDfu@aHfLoCdgB_b@nY}Gpa@cNdfBcl@vKqDlOgFxEeBz`@}Nza@cM`c@sMhh@qMvxAs^|o@gPtgAyZbb@oLbJgC|t@qQra@yJf`@kFnBMbi@oDzd@y@zgABf}AeAz@AlZSj}ClAlLD~q@cBdO?b@AdCE|]}BzSuAbk@eEziAqNlf@kGjs@gKt_@eHff@mFtm@{AlrAd@hiCfA~Sd@`[^tk@i@lUm@|Z{Brk@uHh{@{Mh|@yOhcBqb@xi@uNvQcDb[kEtiAiDzd@iDzj@cH~xDy^b^{Dt_AgK`kBmJ~_CkU|xEyf@fkAyEjrAgLzf@cDrDWn|De`@hMwAroAiGdKsC~k@sGlc@}IvMwCz|@uXvc@ePxc@mOb\\eMf\\oLlTcGzIuArPiEl[uBxSk@nc@p@tWbBxHnA|QpDla@rKfLhCpNxBhMvDv]lGb[~C`L`A~ZnAxYT|x@iCdGY~cB}HfGYfaBiCxiCcEfzBmDrACxhBsCxc@_AfO[`dBLfo@o@zm@k@`d@kDht@m@|o@i@v{EmHpn@gBdo@yHrkAa\\dfDueA~GcBrv@iR|z@wSxyAo^ve@mO`UiH`k@qYxl@oZbYgQvYsLjIyBjEu@dU{Dvf@gDfSc@ll@Tl[D`cB]tt@pBbLXxrApD~SxAbw@rFtg@`Etd@`AbP?xPg@he@cEz]{EfaCy[rwCu^vEw@~FaAxbAiP`e@mNr@StVsHxcA{ZtmCcy@~FqBz[aL`kBuo@jb@}NjGsBfbCwv@zfAo]deA{a@|c@iQhl@}Xzh@}V|`Ami@h{BynAvn@i\\fZaNtDmB`Ae@`K{FtoB_eAnhCkkA~[cOxt@m\\h~DugBvzA{q@|u@m]jvEioBrl@yV~KeFv}@}a@~fAag@jaDisAtCmA`o@sZnpBe`Arv@y^h_Awc@jCqAfbBuy@p\\oPjjCipApw@_`@fcAii@fVmMdXsMxaAwe@pZ_OzH}DhbAmg@nu@m_@la@uS~t@m]lM_Hn_A_g@jd@kUl[iPtf@iVvVuK~VkI|NkCdLi@fo@{Fxd@uDjY{CbUgE|]mHzu@wTvBm@zs@iRdeAmYt_@wHzRmC~S]lCEpB?h]?bmAxBni@vAzZDvd@oBpc@eDp_@sE~NyC~tDco@rvAkVr_GkiA~~@{RpkAwQho@cHzXaCra@iDpMm@hMm@dKe@bPc@xPe@hJUzcAaBzwC{ExHCbKH`OBhMQf]w@`_@aAbPy@dRc@xZs@dACxl@y@tFKdo@}A|OYnHOvq@qAraBgCfFItgAaCpkErFprClEf}C`Gzu@l@ta@kBj\\gEx`@iHhf@uNte@iTrWcO~]oW|UkVtc@oh@hw@ibA|hAivA|RoVto@ox@h}A_oB`~ByqCptAwcBxw@qbA|KkNpr@q}@toAiyAdu@s{@p|@_iAh_@ok@vLcUjGuK|`B_lC`nBidDbp@yfAjQqZn|AohC~Zui@fnGekK`GyKvoAmtBpg@ax@df@{q@zYi^viA}mAvn@on@xa@ua@hkCcoC`\\a\\~P{P~n@aq@bm@co@jSsQdZyWx@{@jAkAPSxpAapAjVgXbeD_dDde@_e@tMsMjqAknA~_Ao`AfuGowGvg@gk@bc@qm@p|@{yA~l@gbAnfAqgBrq@kjAza@g}@pUwq@ldBatEbUcf@vPo^~Se_@va@_o@drBmxC|b@qn@~`@{m@fNmUlMaVjSkd@na@akAfOs_@jPo^rOm\\hO}XjR{\\fYkg@tl@ccA~lEmzHd\\cq@bJaR|i@arAvwAoqDtW{o@pk@awApLwYxjA}vC`Q{m@dIob@hK{u@~Dsm@Xwk@[ij@yBux@oDe_CqJ}{EwEy}BkEkoCkBmaAkBot@gCqs@yEiq@sFon@_H{m@iVcpB_UyeBmHy}@kBmo@@udA`A{n@bAkp@jBynAtFcxCnIsaFjA}~@~Aww@`BmoAbCw\\~Du]dAkFrDkR`Nah@hHaQ|KsTdOeYnHkKvHkI~YoYlWgUlhAc`AlD}Cru@al@`LyH~LkGfa@iQnxA_l@zcDcuA~qBiz@xuA{o@rzA_}@z`B{`Ab_BebAzp@u]dn@iX~yLyqEboN{vFrx@w[ly@c\\rbBup@diHysCzxHuzCjAe@t~Aof@j]aIvQaHnd@gJhhAqQtlA}NdJu@fcBcNzMaBnaAwLhh@{Kh[uHln@aJdcEu~@xsEohAje@{Ib`D{l@tjFihAbrBgd@~mEshAnlCmr@lhAgYbtKwoCdr@qNhr@gJrpAaJtuAuDtsC_I|jBgFxaAoCxqAiDtuA}C|d@qBdm@sAth@aBvb@mDf\\gEjh@uKnjAgZzgB}d@tnAc]xy@}RzZwJjp@yNnT_DzScAnNSx{AqBh^g@nbByDfl@wDdWyBvVaDlOwBbC]zHkAbjAyPpSoCliCa_@dKqAvlB}Wj_@{GlZaGhq@yQlYuJraAe]x_AqX~d@oMfu@oQvN_DzHiBlWcGtZkHvYiHd`AqTvq@gPfM{CjDi@lI{BpNkCbN_CdP{Bv`AaL|ZgDxUiCju@kJngBeY~`Eaq@bN}Bdd@uK`d@cPfbA_h@rUoMbS{K|OeItG{CpN}JfQeJlp@_Qfj@{HvkB}PtNqApa@sGfDu@`k@{Lj]}Hbs@aPtKkCnb@cKxk@uMrgA{Vng@kOhZiKvQkG~tCafA~_Aoc@v`A{e@h]}PfZcNxCuA~b@wRzL{EzOaGjd@iQfhBkp@zEgBxFuBf{CkkAdI{CnIiDz_@kOvx@y[rpB{y@zH{CfSmInXiKtu@eXt\\eKpO{E|MqDd_@aK|^eJbP_FtYiKnMkFnLaGjXiOjYwObZeO~NyGhN_G|Ak@|CcA`M_F`SeHxLcE~ZmKvt@qVhv@uXpCqAf_@eQla@iU|~@sj@xQyKpg@qZna@yWrd@}\\~}BibBvxAifAhDgC`_BkjAny@uq@v]i[|OsNnk@ah@bk@og@bx@qt@h[oYdQ}Pz`@s[`OcL|PgMna@cZhUsQzE{DxDwCfg@y^xKeInr@ah@lbBokAhvAkcAfOiK~SeOpU{OdL}GhE_CpF_DdVgNbs@k[bp@oXd|A}p@rd@sSff@}SfnAod@zJqDtYqKjg@mPxoAs\\jL{Cbd@{Lzk@gOxJkC`M{CfUiGlSmFxV}G~~@mUzYeG~Dy@zBc@`KsAbo@}H|RkDlTmEh~@qT`NmCfp@oM~x@wOn[oH~[sHpnE}hAphAmUrPkDlK}CnDcAt^aMx^gPxuAeu@fJ_Fh_@_SjMqGvUqLfd@yXvWqPpYiQhj@wYnEaC`gBez@rUgMfdAqj@zHqE~uAix@tc@iUd_@gQng@iVzSgKr]{OtLuFxbB}w@jMeGhh@sVjlCmoAf~BogAbJiEnNuGhVaLb\\sOtx@o_@jeAcc@|m@wU`eEoaBfbEa`BhkA{d@`j@gUfNkFp_@wNhfB}k@pbAsY~jBah@p[yInKaEpj@aThJmDrUmJxp@eXnVqJzq@aXtRcI`SyKdS_MxRkKhCwAt`@gTnNwGxOoGbMyEdL_EzEwAx]}KpP_FvOsDvb@aKt`@qLv}@kYvc@ePfe@yNpa@aNna@kL~c@sLlUeHde@wMvN{Dj{@oUjvCou@t{As_@hs@cRht@}Tzc@}Ora@sPbXyK`h@{Sx`A{`@x`@wPza@kQxx@qb@~kA{p@vDwBfi@wZ`XmOdYgObm@}\\fCyApeA{l@rzEymC`p@u_@t\\mOpMuFrN}FpUkK~TqJvY}K`XqKdl@sUpg@mTzL}EvvAuj@feBqv@f|@i`@nf@mUxrA}o@xg@yT|s@_]by@{_@t[aOb\\}Njx@m_@~ZwOlJcFld@iVhd@}Tbe@sSbw@u\\`w@k\\~]aNd_@qN|{@k\\vZqM|[cQnx@qe@v^}TpSoLhK_GpMkH|NgIzSyKx~@gd@dyBicAl`@aQlLeFpMiGfe@iOpa@wLXIbh@yNjXkHl`@aMzb@wLrf@}NhYaI`ZkJp`@sQzZqMbd@_Vd\\sUh]_[vO_NdUmSdOwOpVkVlWiWtTiTfOwLbOgKvBsAlPqKfZsPhOoJtb@{QzZaNpL{Fv\\}NzTiJtMoHfYeLxVkLvMsG|^qOpI}DjNsG~\\oO~KgFxS_J`YyMjZeM`HwCrSgJ|^aSrg@aZh\\qT|g@s\\ja@eXlWqP`[iRjYuOf]eRr^{RfRcLbI_Fh\\eSr_@yU~O}Jpf@wYld@_Ynb@yWt_@kUze@sZ|g@g[pTyOnd@cXv]_Rp`@{Ple@}Sfe@iTpZsQtYwPtRsKhk@aPz_@mJhg@cM|SuGbk@sQlg@}OtI}Cdu@aVvv@gXjY}Izf@oOt\\{KpZ{Jd\\aMv[iL|SaHvWoIbWeIxe@{OxE}Arb@yQtw@s\\xc@kOtZaKzt@mVlUwH~f@eQ|o@uTho@{Q``@cLbV{Iza@sMps@oRji@mN~d@kOh~@mVft@qO|ZyFvJiBvl@qKvs@gM~r@}Ndp@qMxkA_Vr_AcTrkAiUx}@}Pvb@uGt[uFdXwEreAcS~PiCnZqEznAsVrtAmSbnAcRn{@wNfp@sO~e@aOtp@sTzl@oT~GqCvoA{k@lOwHvYqQlg@iY|KgJje@gVzf@eVzd@iYn_@oYrKkIxN}Khg@cYd^uP~d@qSbh@cTvq@qVvl@aRjs@{Ufz@aUffAeZhf@iMz~@cVnf@wNja@gOnd@iPfe@gR`UiI|S{Hde@cQ|z@i\\bz@sY|r@eYzq@mXze@{Qt]mNhLoE`VeJr{@}\\dJmD`a@mOr}@iY~s@aWni@mRjDiAt\\wKfo@eTveE}oAxhCkw@b~Bsy@j}Beu@`L}C`z@kUnyCgbAhlB{s@fO}GzuEkgB|lBew@bGeCnm@aWfm@{SbkAq_@jqA{_@lu@aX`LuFtScKjQwIbv@{d@jz@sd@ne@ySn[mNh}@uW|PaF|l@gL|cB}Rvc@kFncCyWxyAePx_AgKnyAeUxHyAnr@yMvWgKnRuHnMeFnSoInCeAjAc@bVaJhZuKdVeIlWaIzYeIpLoCzKiCfWmFvScFzQwEdO}EvXiKf[}LtTqJbVoKtL{ElMeFpO}D~WkGlV{EpTqEtVgErAUl]iGnReDlZwF`nA{U|l@mLj\\}Ibb@cPjY_LhZcLpFuBvf@}Opi@_Rbk@cRxb@oNbe@cPbfBcr@vr@kXri@{SlaA{`@dVaKbFuB~h@aVjO}G|UwL~NkFl\\gNzh@oQ~j@cR|CeAba@oLrx@eWllAi^rg@mNrx@}Rdd@wObZqJf`@}Nf[}Lxj@sWfRkIbLqEdFsAxBk@fw@cTln@wLze@mJp_@aI|YuHxd@aLzOaDbX{FtT_DfVuCjZ{Cl\\qEnKuBfJaB`OcDdHaAjScGlReFfUmHjSmHbTsG|U{H`\\kId_@uK~^aKtI{C`SkGt]{JdXgGxPwExKyC`U}Hr]gMfVoKlTwIfLoEtEaBjVyH~PoEzXcH`_@gL`T_HzVcHfKoDhMaE~PqFnPuFxKuD`GuB~IgEhMeFzHoDnJwDhJ{E|J{FtI{EfH{ErHkFvHuE`IqEtDuBfAo@dHkEvFgDtEuCnHqE~FuDhIsFtGcEhEeC`F{BzB_BzImFrG}DjBaAhCwAlFkC~H{EpGkDd@_@~@s@|HcElEeChGoDhBm@~HoE`MmG`F}CvDkCpEqC`Ay@tBgAhDkCtGaGfIkHdGqFfFkFtEgF|EmG~E{EnCmCtFeG|DiE~AeBtGcG`EmEpDiE`CqBnEyEdHiHjGkH~GcHjJ{IhUgVp\\_]dIiH~MgL|NsJ~QeLzLaHnSmKdXwMdZcNvQ{GjUkIlUaJdKeEb]gQtGiDjd@uXxHeFxVgOnDeCzSyNff@q[n_@eXj[_TtTwMhVsMfOuHtIeD`F{Bzc@{Nta@wMtMyClLgC~S}Cv\\iEf\\}ElNiBdEa@bEu@pNiB~PoCpM_BjGu@rHiA`Go@rL{AhHuAhEe@pMoBbMeBpOuB~GiAfMwB~IcBnNwCnLgBfMwBxI{A~H_BtKoBtCe@tGgAfJaBzLqChTiDpMmB|LkBxKiBjHcArM_BfKgBrIqArNiBlKkBdN_BzLeBpOkB`PyCfMcBjN}BtIoAxIkBbK{AxPoD~KoBfMuBnJkBlbAsQzl@qKtZaFxYeF~`@aHnj@aJtZyF~`@}F~p@{GjIcAbh@}GbP}A|VcDxB_@jEu@lQ{B|FsAfBk@nR_DfDm@lGcApZ{Eb_@cJxp@{U~[eMf`@}Nd[wJf]sKpXsI|w@cVlj@wQ~HqDbOmGpOkKn\\{Sph@m\\r[aUd]kW|v@mh@lj@_]ls@ya@b\\aUjZ{VlU{TxAqBfP}T`Ywc@zP}WvPuWrb@mk@f^ia@fYcUd\\}TfG{D`g@i]v[g[xYua@h_@_h@tXu]`OuTvNiRb[kYtK}JxVcUrLyKjIwHtFiExFkE|F_EjPcLbI{Epg@y\\fOkLhG_FlGaFtXsTtHwHxC}CdT{VfRi\\fLaLjDgEvCeBnLiHhd@{XrWwRnSeMvKyG`TeNzMcJdJ}Gdd@{[lTyOxQeNbAw@fWsSlCyAnCaAdHyB`pAs\\bDi@h@yCd@oDf@cHp@kI_BwiA}@gR{AsQeBmGyDcNuAiHU}INuD\\sEh@eF\\eCr@yExByMd@aBr@_BzA{B~A_BbEoD~IgHzTyPfv@yi@hO{LdIsHtAmB~@aCTuAFaBMyCiJwe@_AwEeAyEkAqCmBwDqIsKeLyLeFgGcA_BwAyCm@_CKsAJsA^wBnOs_@xz@{{ApA}Dh@{CV}EgAcZuA{_@wAgYaHce@w@{GOaC@iCZuCbAqCnAwCjAsC`AyCd@}BDeDcBwZwBeJoCiH}AqCqBqBiKsKgGkHaEmIoBmFqAeFmFoY_AmDeAmC}NyTs@eBi@qC{@mIsBeXo@kFwA_Ne@wGEiJRqKn@mXrAst@BuADuCfAal@UmV_@y_@FuQXkMd@sKtEmr@f@{JrB}Vn@sJn@yOVoMO_Pm@u`@qAwe@oBiq@_AaSG_AMeBaBmU{AkPuD_^g@eF]eMCuGI{MF_KXef@HsNX_Ev@mIzB}SrAaQ`AkVhAmPzC{VbF{]lAuHdEaYVcDHoDZeI[a^cQeqByA}WImHD}Ej@{ExAyCxDsFfEmEzYoRdR{KpFkEpEyDlGmHnJwLnKiNpI_NzIiXlAaGbAuHbA_FfAeDfBsDxBkCjEeFfCqFfBcG|@oH|B}S|@kJhBm`@g@wHgCaLcCsAuFyCcI{AwLY}s@hFcOxAgRrCe`@nGsR|D{SbEaSvCsTdBoQ^aCPcYaA}HsAq@Kq@MwWoEi]wLsAu@{@e@mScLiEqEiOaPy]ga@iQeS{CwDkFmHePe[qVqm@}Iic@yBqKoIil@sHgl@o@uKmNg}@qFm\\{B_NcS}z@wCoOwCwNsFiWKc@uBgPiDo^q@cHcBgb@CeOEa`@`Ag]LgE@o@`@kOpAev@ZwRjCwl@lA_X`Cq[|@qLxBqYzF_e@tFmc@jD}ThHae@p@gGtDyYVkETiEl@eHbCc`ArAyn@eAod@qAsMmA}He@mDScAaHm]_DiT_Hge@g@w@Wy@G_ABu@ViAmBu`@}I_jAi@_Ha@{[PqC\\cF`AyHl@uEbCoLp@yCfPc]xN{PzPyMxIqFnMeGbE{AfI}CjyAin@l@WpOiMfWcZhe@kh@~D_DlDwBlNcHtFiBzIkB|J_B~VeEjQaCdHaAj`@yFlwBy\\|pCkc@|_@mFpX{En@GnAMxImBfUqBxVc@`ShClFr@jBXp[rEtvA|SzX|CrFj@vFn@j`@rKvPnFlT~Drd@bD??hGx@jb@~I~Bb@|O|ArDb@x_@fFxEx@lIn@vHr@bBF|Xe@rZkFbg@qRvj@qS`n@aT~d@uObCa@jL_B`Lw@dS_@pHMdNWja@yAzAG|g@gBfOo@b`@yA`IUlo@aCz`AgCdiAyCzIw@jG_AnJeC|IgFpViMpWkOvFmDxX_Sr@e@fXkNjN_EnPm@lRnEdIlEdIfGtGlFzM|K`FvEne@pb@~KzK`r@nm@pLpKtXtUrRzRlPfLpAd@pAPlBXhCFhCAx@gAfAq@pA[tACrAThAj@z@|@hLgAz~Ac`@bw@cOp_BiZl{A_TfmBoNdlB_KlcAyAbe@rDhn@c@~p@kAnk@{Hdi@}Rxl@kXne@kT~d@uOlo@_KzqA{RzT}Djd@gNr{@k\\bOqF`EsEdEoBtIkFxEmChAsCxBgBrCm@bAFtAH`FYtDuA|DkC~BmAzAmAda@e\\tX_UfmAk{@jd@e]z[iVf[eTt]iYlJeIhJmK|IeInHyGfYcQ|c@c\\`k@ac@hDiCd]cX~\\kVjo@of@xc@cYxU_OnMmMzTcQvBcB|]kOjTyFxo@_QxSkD~FcAjd@qHbj@oJfi@}Mv\\eJds@oRlBo@lPuFnf@sQfKwDhUcLfVkOpd@sa@zP{RvQ{Tz@_AtJgKxFeJbGwZ`@qYqBuFaCyEcYuj@mKyS{V_o@uVyn@}lAy{C_DaIqHeRekAkxCmD_J_Yyp@ypAuiD}DoMcYk_A}SkbAkRafAczAanIsGs^{]moBodAa~FgoAkgHql@ylDku@emDeV}aBcQqm@gn@ckAerFmhGsxBakCsn@yq@iTeUy^a`@{]a`@{JqLkAuAaOmOyNqQkViYiPySwHgMgEsI_CeJgAsGeAoJ}C}Si@yGGcFK}FdAgGtBqM|B}PhAeNdBqS|Em}@n@uLhAiLfBaMrAkMbA_M\\_LTwP\\qHv@iGpA_HtAeG`DeLlCmHnEwIdGeJjFcHfCsEpBoEtAaFpAgGxBiNnJkq@vCqSlAsGfAqE|A_EvEyJ`C_GlAmFp@cGdA}Un@_Kv@uGjAaIvC}TlAeM|@aOR}N?oMe@}H{@aFaBoGeEmKcAwDq@qFMwF`@ia@EkP_@iGm@kEeAuCcBiCcDqCwGcFeB}BiAwC}@{D]uEWwHYeSJiYPiU^eHv@yGfCuRnD{UfA}KTeHIwIa@mI}@mLOkFLeEh@iEbF{Q|CcO|AiMdAcNlBm^n@kSNsOUkRg@aTUoQQoPJ_If@{HjAsHxFiZpAeJz@aDlAoCvB}BvHoGbGeGrFmHpFaJdD}GpBoGxA{GhAeK~@_IhAyEfBaDlBiCfKiHv\\_Rrk@mYf[wNz[kM`JgDlKgCbSmDdK{DpKyDhK{GpHoE|MgI`OwIvOqIbKgEzEqAhFuA`S{EbWyFdIg@nH_@bEbArOz@j\\bAj\\d@`_@JxZM~Mk@pNsA`IkAdDgAzBcBjA{Ar@}@rA}DdB}J~AuPdKsp@lDy^l@}^cAk\\mCmQ@sHtFkl@xGag@tFy^|Ii{@lAwLbFag@NwJf@a\\ZcVlX_nBlAgKl@eDvFw[`EoPnBePp@kd@jBcc@nIeiAt@gJrKkgApGci@tHe[p@aD`BkEzGqZx~BojKvV{gAtFwVzMc_AlWmdAhEsNhEwMnNc[zSel@~IwRn@u@xFaJ`IeLnF}G`FqEdDiBxB}DlDgOjO_eA|AaJ~CeJ~LgNtDmCnEiCdMyCvQmAtRyB~KuFvFaKxDgMvDmOnAuE~@sBvCqEpBuA`C}@~D_AzFy@l_@GhZmBzFaAbIiDtIaL~HkYrAmFfGmVnq@gqAv\\wn@fNaVtPmThRiUpSqRlT{XtIcJjJiIrPuZ`i@}fBnK{^|@kC|JuZpDoIzR_a@lI}LjQk`@dNea@|@uCrNcd@zPki@jQsf@~Oq^bB{Dru@{tAnTg\\dCuD|]y_@`YsZlEcH~J_PlP{Xdb@wq@tWwg@`Wgi@np@opAt_@wb@~FoFxGiElCw@rCg@bHi@fOm@xVcB|IqBfK}Ctb@wTbRkK~UuFfUuCrTsEbKgEpImGxKcLjL{IlOqGhPqElfA}Cvl@c@pWSjBK|XuAf\\eDny@wGnA?fc@?j_@C|MGbvAi@v[_@|p@w@lMOvk@cBl^aGfr@aZdViKhUeJhViElVwBfbAcIz`@uHzH}AjS_E|e@wQp_Bas@~q@iTfOeDxb@mJbWyI|RgLjb@kZxb@e_@h@i@ftAokA`e@{h@la@ka@|RcOrb@y[vb@{[rjBq{ApKsLhJwMjI{MhF}ItGuHhF{CrL{D|J}CjUwHzUuInm@}YxBkBxUgSrPuN~HiLnD{Nd@mHdD}f@~Fwp@p@{F|AcF~GgMzKaHjLuErYoLfWkHvW{Bfz@r@zu@Zz@?tA@~[Yt[yAfe@aBhScAtHk@bIoBbKeDjJ{EhYaTjn@af@hNmJ|MiIpT_EpTsDth@iIbsA}Z~t@kPnPcF~OyE|_@eKfa@gJda@mIv_AgM~[kEpm@uLzVqE~JqBxIcB`vDew@nhBe\\`m@aH~s@yGltAeM~Jm@~j@iDtW}AxgAwGl~@sFhlAgHhb@iJj^sNbYoZbd@_k@t[}Y~g@ad@fg@kc@vlAynA|e@{f@xg@u`@xFaG~HwFhI}FtRcN~VuN`i@kZtv@{l@fN}KnE_FpDiFvCuF~cAczBxt@e_Bn_@gp@~\\qZji@yf@dJqFdFgBrFk@hEOjE\\hi@bHpEJzFSjE[nFeA|cBq{@bFgBbH_BhNgDjk@kJ`ZyFlMeE|c@oTbN}ElFwAhEa@nYeBn`@RjFu@lEqAhE}BnEcD`J}KhDkC|BwA~G}CdS_EfVkClk@oEpZiBzKaBb]aMfMwDvDi@xEq@rDYbm@{CfUo@bGq@`DgA`CsAxAqCl@oFjBgUFkDxAiHpDeGvFyHhQiM~GaHhFaHxHwMpPs]vGqInF}EfGaDj\\yLpRaGjP}DlTqBrDo@|Dq@tCmAjC}Avj@qk@zIgJ|LgLlLaLbLiKhGoCpIcB|OaBpMmAzScD`LoCpb@mLr]wL~KwB|SqBx^cDrReDxKgClU}HrKwBrJuAfXyCna@iCju@eHl_@yCpP_BzNmAz\\iBde@eHj[uEh[qGng@oIr[}JtSoG`S{FxQeDfSe@dk@FzX[`WsAnMmAfP?xX`A|QrAfQnAtc@rEhe@bDnf@pB~Qc@zQyBlYqFtKkAle@eKje@qLxGgIfKcXbIsa@~KiqBlDwn@`Die@vCaRlE_Q|Viy@\\gA~V_w@lIe^zA_MdBkTzDaeAdAyYj@aGbCyJ`FiKbd@sr@hO_VdSe]jL_PlKmKpAqAjb@y]dtAoy@dgAgq@noAmu@dy@{s@rXiVlw@mw@~^aBrc@dHfnAvo@d~Azj@naKjCdGIjnPmSxfFqLrqA}N`Y_DbuIk{@`dMsu@dcDcDf`A}HzyFkm@|jHgg@d`@i@dP]vgAiBnm@cAnS]~BIhPUzDG~d@i@dIY|YaAziAaInkAsUdUiOnSoZdLsPhUi_@bn@aeAjVsa@zU}a@lPaVd\\we@xt@wr@jo@}o@zl@qr@fe@wiA~_@}|@bj@keAhu@_bAxWu]fvEshG~yB_tCtg@mkAhq@aeCje@saAnRmQ|`@__@rbA_`@n`DwU|~@wObv@iSh{HqtE~kC{fDtLkKtTqIhTwDp}Gwd@hO{@dSoB|RkDdaDicAbtAmYjgAqLjb@o@|m@eAtEDjLi@fKb@tKDly@qAjj@pBtt@dLhd@|Gn}@rMhYhEp]vF`El@bP`Cxt@nJllC|AfZ`CjYvG|SnIxQdPfw@xaAxZla@pNxNxFlHlF`GrNjOp@d@nNnJvMrHfUnMzKxCtYzH`Gx@|_@lFzcApAbzDxEdf@l@fa@L|]LrbA_CfAKjAMl`@}DfjBmRr]c@xFdAnt@hXbCnAzv@da@bb@vSja@fP~RpEhTjAjcBaGvp@{Btq@i@tY`ClTtGxi@~QhDvAlNbFxfAxc@pXnSpCnErDhEtR|]pf@pfAt_@~z@~b@b|@xO|TdSnQntBxjBre@h]~k@zRfv@Q`XeD|JoAnOqD~d@_M|QmDzVoAxIZ`OjCpWrHp|@fj@`^dStm@r_@zc@rX`ClAhr@~`@`ClA`j@pXzaAxf@bClAzx@ra@drDfjBx|@vc@xa@rO~a@nBh]iBz\\eDtcAk@nsA`Zfy@tQjgA`[`jAbXhwCzi@vJrGdVbPhC|AfbB~bAd|@~Y~k@dT`mA`_@d]jK`ShJ`b@tRjaBzzAhnAxhBz`AvoCzHf_A|Ghk@~^bi@pp@p\\jbAvx@dTfZxKxg@oChkA~Azs@xL|a@lN~NnRhLj_Bxz@di@r_@`~@lvAhUvYfcBpz@jj@z[jd@bNbu@jI``@fEx[vNdf@~[vWpFhjBvG|b@yAtdBko@lcB{~@p}@aj@|H{EdpAcpAdDeDx_Ay|@~A{ArZcTl^_Mlp@eNnZYhJh@lIh@hDRzhCdQ|j@xDjTxA|QnApw@nFniB`Mnx@pFd]~Bfk@~C`l@`HxShF`KdCvvAnd@r^lQ`[hX~[dc@fQpIzDdBnGjBfS~FfYbFlPtC|T~FdV`Jhk@hTbl@bVrj@fYdsB`e@~g@hJx_@dFl_Blz@zyBxjAdq@rWfq@hLbjDjOxt@~EdhAlOnQ`AdV`A|MZrQmAzVsExSiHlXuNfMiGtMoH|ImDrMe@tLh@lc@vCpTh@rTg@|OkAvUkC|r@iLvLaClPsA`s@wClg@}Hre@}Ixd@iJxd@wEhcAqFv`Ay@~Xg@b]f@fhAxIhh@tFbe@E`FSfeDcXrAKjbEq\\vo@wCxNo@fZ`Cpg@`O|^bLbj@iDjmAcW|hAk@pxA|OdcBvo@j@Rz_Blm@jn@zC~tAmR|sAyNti@cFjbAjGlhBdNnlC~U~fApG`eAvKfb@jExb@Hfz@cOrYqG~QuBzRo@lV[nYl@b_AtFzm@`D|q@pFjo@nF|l@jBtfAeBjgAyH|XwDdo@m@`zAbQzlApHr\\cGtgAoQht@gKl_AkEt`A_D|k@aA|k@aJjeA{Ztq@cRbl@mNdzBmFdYsA`g@cCf[H|iFdDrf@V`l@{@lMuAdMcDpSqLxSsMpQsKxSqMbLyGhLqEvU}F|ScApSh@zHRvNn@~[x@dUOpViDhr@sXn{@id@bw@yi@|a@eWz_@}PpnAwSzd@iI`SoFbHsCzGsD`KqJfIwKrS{_@jL{OxMcMxbAgWz_Cmp@~hAs[naAqYzzBop@pHyBh{@gj@`y@ah@lp@c_@pu@qSjnDg~@hqDi~@zo@sEdsDhVd`@b@vZeAnTiDxNuFbKaH|MoKnr@miAtfCioBdX}PzJuEdJ{D|LiBv`DeEfl@jDhm@fPrg@fLjo@tFbu@nCfTbB~g@pAhc@sBvq@cLnd@yJjSyDt[kB~`@fEd]~Fni@|Fj[GnXcAzViFznBu[n[iEjx@gGnlAoO|}@sQbTuCn^qCtZ|Cr\\dCpcA|HnU~@`V]htDmW|cA{Gh`CeFbKCbL{A|IkEbGoGnOiQtq@a_A|p@c|@pl@ez@dSkTpOiNxXoS~XaN`u@u[xs@qYf\\kJbKaAxJOrTd@lKdArM|ApPtFzPfIrp@bi@vi@|b@`o@~j@`g@ng@rTbX|mBhdCff@vo@duApcAbz@vy@`^p]~KnGjR`G~vAhE~g@L|h@iFzT{BrgDq\\nlA}Jrb@{Bjg@qChiA{Ir]kBto@lC`p@lBrYwBfXeDf[cFlZsH`YeI|IgCzPyFxO}FzdDcnBhpAez@llAus@|GcDpJmDzJmCry@mJbd@gErd@wCdcAoLd`@oHrOaCxOmA~T_@nTd@|QEpJ{@zGkBnQ}ItOgJtc@kV`J_D|JwBvOiAdHWvEQdvDcG|STfaBoDr_BwD~tAuBdIEho@h@pm@`@`iDyAb[O`[rDxSh@nR`Dh`@pF`j@pHbtA`R|b@dYl]l]vb@`Tl[jJrYVfMg@dSkEjH{ClD{AzH_Dhq@qX`GoC|IaEba@uPhBu@r[{MjGa@vKs@xF_@p\\yBnAEbo@j@~]nDj\\zMbRtMjc@vTjx@fc@`T|R|a@|c@p]hY`n@tOvWnBrb@BhEC|EElE@d}@qA~UnAbNp@`CjBpIvAzOrOpCnClGhJ`OpN|LlGtNvF|LjAh\\b@xLPzCT|g@lCdGbCnWpDxCd@zQjD`WdC~Vn@xAC|NI`TUlOY|OsAlAg@l_@oOdKyGzUsO~]}YhIgJ`BiB@CdDqDpLyMdFqB|QmH~Kw@fSwAh_@t@zFLpe@gEl[sClhA}JtQwAr]oCnh@cExk@sEjJItHGrLKpNMrCZdm@nGvc@vFja@rJzdAjTl|@jHlY`CrSdBpCT~fArEjr@lCtNbD|t@j_@~UnO`@TjQfMxn@~c@rVzBzHvAji@zJdg@nP`|@j[j~AtIj[lBzb@jCpoBlKdb@tCjl@hEzBCnk@w@zMQvY^nh@r@xt@jLfxA`UhSzCj_AbOlsB`ZhMBtM|A`\\fEny@nGrr@Ex\\c@xf@}A~_AaFlx@aDxGYrJI|x@kAhBAxAIrB?bE?tJIxFCfJ~@hF^jKnDtFhCvG`DjEbBlEfBrOvFzPrG|IhCxHlArMr@tZgAx]aBvTMxRd@f}@rMjv@hLtx@|MrHj@zHhApQdEvI`CxLlDnLrB~YjCxVb@`DIld@^vbEm@jaHzEvXgDlWsJrzC_wB`lAk{@lQiMfHwEzHwC~SwGbUmE|b@aAnf@lDff@tGrgAvFf\\~@hf@uAtoD{i@x_AmMb_@i@fa@tCts@`Rxl@z^jYj\\jQ|U|h@vs@dS`Y`d@dn@da@hh@nqA`dB`tAjeB|G|IbGhHrKtNhLfOlMlPhFvFfHxE|H|FjKdGfMpHdNdGvm@nWlP`H`nA|l@d`@lQ|uAfq@heCplA|k@zXzw@``@jjApf@buAro@|WfH|TxBbyJ~@tq@h@thAZjo@zFtk@dHhvAhTlNjCnO~ApLTtQKbSs@vk@cCvO{AvPsDlIiDbLmIla@}g@`OgObQwJxK}DzNiBdU@fY`CxRKbSaD|J{ExDgCvOqPfg@aeAbJsQ`OcOdLeEjPaAlb@vEzZtCjTZrTkDtRsHp}A{i@fiCkd@l[wFdY{@nv@{D`O_@jQaGdR{MjJ}L|KcZ`Ju]|K{`@bGaKpIsNvNwLdRkPji@g`@`OaL`RgMtDwChFmBhSgGd[mIbn@uMpF}@dcAuP`X{ExYaAz\\\\j@@xPtAvOpChR|ExPlG~d@zR~PpItQhIt[rK|MzA|C\\dkAnEdC@xi@tAxTh@fJVlIBlmBjGxh@fElZdGxg@lNh`A|U`{@rUdx@vIjaB|NpfAlRx_A|Ong@`JvCpAfd@nJbWpAdo@hCzdAjCf]|Avo@vCrc@fC~G^bMp@fTlAzBLxOz@pQl@xv@jDvj@rBf\\dCj\\|DdWbIrRdJvf@dYbj@x]j_@bVti@j]dPbIlj@lO~tBbYpf@bAheAkJte@sEzeBmSp{@oIfxCyNzmAkB~lBeDjaE_Irp@iDbvCqZlzAiLvy@uKpnFseBfAa@laBoh@zaA}^jWaM`w@k^t[qOrY{Mz[{K`m@oKla@mHvr@eNzPuHrVoK~[gUfAw@rFqEf\\mYjq@yh@bH_Fxm@}f@~lAu~@nJiHbQyMli@kKbmBnHfLd@rj@rCtOXnRi@~Fk@xBUtFaAbSqFnv@i_@bt@i`@nbAgf@fVoLleAsg@lPuI`r@g^~`Bcl@jcAw\\|ZoI`TmKtNoNbHwJvGmOzJod@fFu_@hCmTzAsT_Aim@i@_]gAse@`@yJx@uJnFsP`HeLh|@{u@b]}Yxk@ib@jNkMzFmLhGsOdDyNf@gNUkk@x@yh@jGyXxMeWtHkKbQ_Njr@iV|UeI~yA}d@lsF{iB~]_PjZsTxk@se@jb@yWht@}`@pj@g]``A_|@jc@aUfd@aOja@uRbcA}t@xi@wg@pl@ke@x}AmmArb@qWhQgFxScDv_D_XvgDkYfOgDdK{D~EeB`NcK`QsV`c@yk@|Ye`@xDeFlLwKrHgFdJcEfF_@bCQdIo@bNjBdkAfq@`tA`w@nk@x^rLhLrI|N~JnUtQ~_A`Mj[RPbZvXrd@tSlWdKdRvHrLlD|HvDbOlEba@lGhRhBjInBrLt@hTpA~Pp@`P?~IB`FSpF]vLw@vLqB~YwFpTuF~W_IzRyGr]wGfZeFjL{BjQeDnXoGbYsGzV}Gv[eHnYmH``@yGxQoC|Bc@zAElIsAdWkIhXeEt[qHrf@{J`m@wLvUaG|N{CpN_B|UeAdW_Aja@aBp]s@bKOhS[jh@yAbUaBtEUrJi@b_@{DzS_Bl]oCfCWbO{AxAMlFi@tOuC|P{CrF_B~Ae@`JaFpJkGfLqEdMoDxHo@pHKzOWbDG|FKtd@u@hx@cCzVgDrUgE`W}@nWeBlW{ClJ_G`H{HvEsGdJkMlSgS~P_QvUoKbNqEzO{ApQNva@`Bz^eBnZgDvm@_K`W}GbR_Sv[a~@vGgOjOkM`Ay@j`@}Ib^cKv[iX|^gBvUrCpKpAvc@vHtf@dMlRnFzY~CpZjA`fAwDxgAoDnyAqFfkAqAj{@qLrdA_Ezw@yHbQuBvi@c@te@hDhHf@xZdFtj@rCdZoD`c@uDhaAkKnXyI|HwDhj@cXnz@uJfQiBhw@iI|]mHj]uMty@}Tv{@qPhi@}U`PoKldAst@hOaLt[iR~XoShRiJf`@qOdVeFxrAeYxRgLhWaWnMmFva@gEnYeBlSgC|L{A`a@aEhPVnQVplALrkBNnzAL~YDvI}AhLqB`Y{Q`H{IpIuKhk@yd@j]kUdHyE|p@__@jj@aRtl@eMp^mJjCq@`a@s\\x^cP~YaKn[f@~Z~BnW?xUwAdVmFpZoCli@aG~UsB`VyDxKyDfPaDrr@kBpTsA~p@_Cjv@tPvm@bI`Me@`i@wJz]jG`m@bD`ZlPn_@fa@zPzPjOjOfStUz\\|Mxh@vPjSJna@OnKrKnVbi@|IrEdWgE`VsHhq@uLfUaH|O}Ef`@sGfMuBph@kB|fAQ|VEph@fMjk@|J~[_WbOgAnXdCrKhDx[dKnSnLlPtNvcA~z@nu@dm@tw@|k@hOpF|RdHtX|Mry@ba@ju@f^fi@hWxSdDliA~ShCd@jsCzh@|y@xW`s@|TbQrFtd@~Kph@hc@xXjPrh@~Mdy@jZ|]xUj}@py@l]hUrRdMzj@vi@p[~T`c@py@nj@`t@pUlRn_@xa@ts@bq@`PdSpOnRps@zh@h_A~d@h_@lRnjAnv@rMvUbHj~@`RbStBjDfFjIvQnW`AbEtNnl@tSzbAtTrL`T`LxPbLtT`R~TxQpg@`e@dDzApEfBxKpCtFh@zIXvQJbmDNrjA?l]qGdWuB`OaNhRwKd^iNhXeJjBo@|FcGr@mERyFF_MDiJzCoLbOuJdm@eLjHwA|E_AphAiTvCk@nWaH|QmS`Xo[ld@}]jVsI``Am@bRjBpOvD|LrGjDrB`@VvWjPbNbFfDnAtZrG~_@`J|]zHrz@dOpDn@`BN|AA`BO`B_@jT}NdVaSrJwDrI{@bXuB|]eBnj@qBrOkDhDw@lTkFbo@uOfXcD|D?~DT`MhAnUvCtEt@lI\\zLYv`@uFjRgHzQiF~H}Bvj@wPzOmOtl@qj@|MoLzG}F~EiB`F_AjU_ApUS|j@XvjDlAbR{Azb@aD`b@wFjmB_M|iBsJryB}NfGUbkBcPvd@eEhPeApw@}ErGo@lDm@jHsAnNmDx[gMvy@mZv|Bay@lwBgv@|q@uWzaCm_Arn@{WzlAwg@xpAki@rGoEtLgK|N_Sb^{j@pk@wcApL{SjZki@fFcJ|dAefBdSkY~A{BfSmYrUq\\~k@_v@zWm[lHuJt}A_tBdPcSldAcpAjlA{_B|j@uu@bPiRby@o`Avn@oj@xTeWnQoS`QkNbWoNxF{InLc^rCeDhZu]zJwVzJ_\\rKmSvW_g@dq@q|@ds@c|@pR_O~EwEhSyd@zTia@lN}UpWk]rm@o}@`l@w|@hTe\\~`@ah@vWs_@pt@ghAl@_Anb@am@`R{YlJgQxo@_mA`BcClLwMdJiI|Y{PjM{F`MaFfDqBzC{ChD{FzCyGpGaLrPsTnDoJrEkLlHeT\\eA`AwCb@aDFmJo@oRIiIEeFEaD@eE\\oDjA_D~C}ExEcDpMyIjMkKn@eAl@}A~AyEhFqTdDsKbEcJzBgDpAyAtAiA`OsGnOkElJwCdKmDt[w[lWwY~DwCbEaCfH_CxHcBj[IdDD~D[lAWh@Qj@Yv@_@~@o@pC}Cx@qBTm@Ro@Z{B`@ik@gA{OkBaUuCa`@sRwrBqVwpA{Pee@wCyG{FqMmBoEqEcWEge@`Age@fBk{@qIwu@|@qRlMwe@jC{K|I{MxPoFb[{F|TmL`O{L`OyLjHiDdPcE~_@wJjWgHjVqInb@iR~C_CtVsTjMiPpM_OtSmZlH}L~HwLvK_JxX_Ntl@yRnGqAlIq@dp@Fl`@t@j`BfGrs@hCbYpIjz@rM`Ij@pYpExVvBtSGjBAzYoEnYiHrG}BnBs@zK}Dfc@uOdVsIvDeBrPsH~Go@bNH`J[vf@cBfj@xBhJ^`Hz@z@LjS|FvRvFjn@fS`MpCbGF`GyA|MuJdNuMjIkFv]}MdXcIjf@gOtDqA|K{DzIaMfZ{`@vAuBxDoHbEgNnDaR|EoRfBkHhBsGbBcFbE{JlCoG~C}HrBgG~CoK|@cDlAmE`DaLvByHtAyFpC_MlAyEdAiC~AiCfBiA~ScUfCuB|DeD~EcEtKuEdOyAdFBxLB`SeC~JoAzGcAfJcAtVoBhIK`NUjKOlCC`DC`BAfMmA`M_B~FoAhJ{BtGeBxVcEpSqBvBSxFg@lHwAnHiBhM}ExGqF~E}HnG_MvIaQjCyD|FsIvCkEpRkNnEcDbBoAxAm@tKiBlLeArPe@fOhAxD^~QjB`Hp@~F_@lBiAzBsA~A_AjGwAhJo@jHLpH\\~CZhEf@bK`BbKjAzJfAxFiAdX}LfCwAzFsCbCy@tHUtFz@|EfCnI|GbIfDvJzAdTwA|OcAzH@~I@zF@jGJ|FLtK\\fGPlP`@lY?pRIfKI|MSnCEjFFdXAnFSb@CbDWrGk@r^kGvDgAxMcCxSkEpDw@jJoBpFiA|NmC`Dm@lPaDz\\mCvSHfTzBzIvClD|AvAn@vE`DrFtDr_@dRhg@~WxP~IfGlE`SnLdAj@~EzCdb@j[~RlOxR~KxE`FhDhEbD~I~Wxt@t_@tb@tAxA`Xd]p`@z`@jFvEfZ~WvXzVjChBza@rZ~M|K`HvElWpRtFlEzBfCbArDp@vEbEba@h@hFRpEGfEYnHs@jI_AnHe@tDa@dBKjADjG?fCv@jF|ApE|BlFxBxDr@bBv@tApB~ClD`GxC~FzE|G~D~ElDhD~DrD`F|Df[|VhKhJlIzGhIzFhKhGhIvFhMvIpFhC`GnBdGpAxItApKlAjC~@xEdBdCvAlDhChJbHjI|GlEdClFxB|Dn@jJq@xJuChFgAfF]pE\\tDlArDlBrG`Exm@l`@pDtA`Ch@lBFdQy@`AAxv@mEvM~Bv\\pD~WpGrPlEzZ`Kn\\~JpJbEfStGbn@bVxbAl]ld@bM``@nDj^rBdYfCnANjYpCnLPzNRtm@nCjBn@hZtTj@h@zQxOnc@j`@tC`C`ErDxo@zk@`ClBbJ~KnB|@rEtBrPzJdBz@dTbKhAJvLbCtNkAzNuAdOwAtQeAzGXrFHxSf@vKFdX@fJH~EJfSe@jNUlUk@p\\wA~UyAx^yBvK[nOo@rH]x^aCdYoBzLyAjM[dI@bNNdM?tA@zBT|CNnFt@FB|KxD|HrBtAP~Hd@vIS`IOdRc@fJEX?dPC~TSlDO~CI`Gm@lB[xGaBvJgCnAe@zE}AfFgBlHoAzC_@nFk@zDa@rC[zGq@zJ{@rFe@dIKtKl@dVxAlNx@fM\\vMh@~Gg@fDo@`Dy@rBi@|Bw@pCq@|CuAjDmD|IaIzIoIfOyNhi@ie@|AwA`B_AxE}A|PcD~JcAjW}C`E{@bCmBxA_DjA}JNgF^wNpB{QbPy\\lYoc@nWqMlPiFfGqCnGwBlKaClq@sJdSsCvPeC|F}@hGkAxEeBjQuGvOqFxEaB`n@aXbg@gQrLsAvAe@d_@aNfs@eT`m@cSrL}GlI{Eza@mUvq@}`@bPoH`CcExeAakCbLoXhOub@xEkS`AoOvAaHxCuXbA{Kb@oKbDsx@rAeR^wFFgA`BgVnPw_A`D}\\lA{Ql@_J`EchAdCuy@f@yUzBsm@n@{Ub@cRPmGv@mGzA}Bv@kAnL_QjMeQvBcErAgGd@o_Ab@wQpBuMnSsVtXqVxp@wk@~eA{`AdB}AbBgCbHkKy@oKwGqR~C_}@r@iFpBsJRgAj[okBx]}jBvBuD`CaClf@gPlKmJ|A}BxDcGteAy_BpOwUb]_h@tHeMdPqXzMmRj@_AhNgQ`l@}a@d\\uUtPuL~RmLnUwPpX{Ov}@}e@|HmEnn@_X`Bk@xCcA`]mOnQcMzKaKtReMfVqJn\\cPz[sI|ReHnLaHxFmL`Ky[lBeG|@qCLa@hE_LvFqKlF}FjFsEtNsK`KyHnJuFt\\_Np]gPhJkD``@wGfl@}Ij\\iGpWaFfNkE|c@wWlc@cYpEuC`^uUv^cXbQaOjT_SfRkSvd@wn@rWw_@lbBqaCtCiEr\\se@fAcCpH_M|Pif@tGeJ|HyI~GyEpq@yl@vRePxI_H|QsLvgAmn@bHuA|Gq@|NcAvn@eFnu@wFdw@}DfMUt^}AlWeHvgAgl@veAcn@p`Bm{@|k@iZvv@iSxv@cQxkAcXpTsEbn@kLt`@oGrWyEhVgE|QwCbPcDxP{DxRmFnYqK|NiGhBw@dQ_HthAyj@nV}L~LaGlMeGdd@{UrGgDlAk@jHgDb^mPz^yPzBUdCYli@mJhaA}PvFcAzYoFfoAeVdd@qJfSwKpTqLrL}JfKsIvAw@hKyFfP}IdGeCvB{@zKqEnWmKzO{GfNwGhNwMlQiT|D{E~OaUrL{RvY_u@jHsMlCkCdIsDnr@{Rzy@cWfZcJjZ_MtPcPrBkBd`Aq~@t]uXfg@uPnCQzb@mBjwA_Ehb@eAng@\\hp@|AxFPzBAnfAfAhZv@fJVH?~RCxw@_Zn_@sN~MgHfXu`@|JsHfJiEvNeCrh@_Dbq@qDxb@s@dQ[`DGb@?tr@cBlbAaDvuAoExe@}A|gCgQzc@}CnZqIpS}M|WiYzMa]nEe_@lByiAd@ir@`CkrA`DcdArF{LpOqM~L{EbeA}RpkA_TtKiA`JmBhu@eSzmCwm@rlAu[lq@cWlCgAt}Bs}@zdA_b@r^{NvqAmh@bbAaa@hh@sSpqAah@~u@sJnv@{Ghq@qG~b@{CflAeHpyBmKhcA{EvTkAb`DcMlEi@tcAsCfT_@dmB}F`NwBjKeCbeAwa@j~Agn@`|@wc@fx@yh@|IyD|GiEzf@yUnz@{d@nEaCjMkGp\\qTzTwPfi@sb@lNaNfK}Jbx@gq@fWyTnv@_q@zPcK|kAo{@|tAy|@tk@u^dVcQrIaElG{BpHqEfOoCdU_Bj`AaC`m@}AlJy@dw@iEth@wElOcBxrAiM`PcBfWcC`[wEpY{DzIkA~d@gEvOoAjUgBfZ_BrIg@p~CiUpnA{Iz_A{C~Xw@`YQ~`@c@fm@bAvU~BzTlEdHnBdOvEhPbFzPzDvz@|X~_Cdr@|}@vUjaAdVvmBth@\\Jzk@|PdTvFp_@nKzWtIbLbFjQjJxHlEbKzEv\\pQfYnN|n@r[hCrA``Ajh@dSrJzSvHfW|HvMfA`H\\rPtA~Wa@pBIre@qAjf@sC~d@eDl@G~Ii@bNu@`L_BloAaJva@eDtGi@|JiBlJcE`KsHlIiJxLgM~XmX~eAucAdOwNx[yRf\\_Ovg@kRfi@sTn]{PbHgDrOwHxLaGxYkGpSiCZElNgBvVaDrC]j`@aHh^_KdNcFff@cK|Q{Ere@iMhPqEnUgGtv@sSfhBuk@do@eUto@aSbb@qLzh@gO|W_IvEwA~QmFpWsHl@Q|H{BnB_H`JsC|PuFhHaCfWsJnUyI~GkEZg@~CyE|BeFpFoOjKaZnI{UvGsPnC{EvBwD~DyD~EkBrGeCtKkCrDi@lGO|Xl@zADzWhAxP^hGHxYM`DHzDAzUlA`DDlNNb]`@dB@`{@~@`]WlLa@zI_@rPiAlB[lZmG|Q{DbLmCdMyDhOgGz`Agc@`mB_t@~UuJde@sNta@cIth@cGng@{Fpu@mIrt@gMv_AgPro@cE`p@qD|^yBpUwAzLu@ju@sEjlDiYlTiCfQoFxPcI|Y_PxJ_Fpc@uQ|o@{[j^qQv]gRbOqJdMiIxOeL~RcQt\\m]xYiSvGqBn~Aie@xn@mXduDytBbbEosCtgCgqAnZePf_@wQboGmeEz_Aud@jYeKpaAu]zc@}^|XiSlBaCfzAwjB~\\kd@vTua@v\\y`AjrAyaC~a@{g@fO}OpT}OjVgL|QmPlrAsxAjS{VtPo[xI}[ff@ghErI{`@lAoNzCsXlOcf@`FkLxHsJns@yl@zlCwzBbPcCdRfA`NtD|NfLfSxDxcCza@`~Bz}@x^jErgEsM|k@_ApZNhn@{@hq@uBj][z}@{@ri@i@bnAiA|MIhFCzyACxKGh\\eApe@c@bv@sCd[_@d[_@dt@{@|X]jl@fBjAB|YTfOtAxs@jEbc@bBnPp@|NvLvQvHjq@l_@||@zEdJn_@fEjS`u@ff@p]n]jNvO`KtTzNbn@zK~d@hGxOvCrLpD`_@vFhf@nJl|@rBfRrKxz@bCpRnPxqAtAzHhSjhAp^lfBtGbHja@bT~h@~Wrx@pa@tWvMdgAf`@lt@vUrT|Ipa@lPhU`OdJ`Lxn@h|Ab`@dqAxYzlAlGlW~Mnj@~F~]bPldAjNr_AdK|l@nBhLtgA}f@x_Asl@lsImxDxwBmbAfsBgbA`lBubAlAm@xrAko@vIeEthD}`BdaCk{@fn@yU~xEchBfwCkuAfzCcpAbe@wRlBy@zn@{Wje@}Rfa@gQ~zAcm@h~@o^lbAc]bLwEbBs@",
		Markers: []mapnests.Coordinate{
			{Lat: 23.770348, Lon: 90.407785},
			{Lat: 23.794249, Lon: 90.414536},
			{Lat: 23.816310, Lon: 90.421117},
		},
	})
	if err != nil {
		return err
	}
	img, err := base64.StdEncoding.DecodeString(res.ImageBase64)
	if err != nil {
		return fmt.Errorf("decode static map image: %w", err)
	}
	const outFile = "static_map.png"
	if err := os.WriteFile(outFile, img, 0o644); err != nil {
		return fmt.Errorf("write %s: %w", outFile, err)
	}
	fmt.Printf("✅ Static Map image saved to %s (%d bytes), distance: %.2f m\n", outFile, len(img), res.DistanceMeters)
	return nil
}
