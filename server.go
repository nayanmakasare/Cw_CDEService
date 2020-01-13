package main

import (
	pb "Cw_CDEService/proto"
	"context"
	"errors"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
	"net/http"
	"reflect"
	"strings"
	"time"
)

const (
	developmentMongoHost = "mongodb://dev-uni.cloudwalker.tv:6592"
	schedularMongoHost = "mongodb://192.168.1.143:27017"
	schedularRedisHost = "redis:6379"
)

type nullawareStrDecoder struct{}

func (nullawareStrDecoder) DecodeValue(dctx bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Kind() != reflect.String {
		return errors.New("bad type or not settable")
	}
	var str string
	var err error
	switch vr.Type() {
	case bsontype.String:
		if str, err = vr.ReadString(); err != nil {
			return err
		}
	case bsontype.Null: // THIS IS THE MISSING PIECE TO HANDLE NULL!
		if err = vr.ReadNull(); err != nil {
			return err
		}
	default:
		return fmt.Errorf("cannot decode %v into a string type", vr.Type())
	}

	val.SetString(str)
	return nil
}



// private type for Context keys
type contextKey int

const (
	clientIDKey contextKey = iota
)

var targetArray TileArray

type TileArray []pb.ContentTile

func (e TileArray) String(i int) string  {
	return e[i].Title
}

func(e TileArray) Len() int {
	return len(e)
}


func credMatcher(headerName string) (mdName string, ok bool) {
	if headerName == "Login" || headerName == "Password" {
		return headerName, true
	}
	return "", false
}

// authenticateAgent check the client credentials
func authenticateClient(ctx context.Context, s *Server) (string, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		clientLogin := strings.Join(md["login"], "")
		clientPassword := strings.Join(md["password"], "")
		if clientLogin != "nayan" {
			return "", fmt.Errorf("unknown user %s", clientLogin)
		}
		if clientPassword != "makasare" {
			return "", fmt.Errorf("bad password %s", clientPassword)
		}
		log.Printf("authenticated client: %s", clientLogin)
		return "42", nil
	}
	return "", fmt.Errorf("missing credentials")
}

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	s, ok := info.Server.(*Server)
	if !ok {
		return nil, fmt.Errorf("unable to cast the server")
	}
	clientID , err := authenticateClient(ctx, s)
	if err != nil {
		return nil, err
	}
	ctx = context.WithValue(ctx, clientIDKey, clientID)
	return handler(ctx, req)
}

func startGRPCServer(address, certFile, keyFile string, server Server) error {
	// create a listener on TCP port
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}  // create a server instance
	if err != nil {
		return err
	}

	// Create the TLS credentials
	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		return fmt.Errorf("could not load TLS keys: %s", err)
	}  // Create an array of gRPC options with the credentials
	_ = []grpc.ServerOption{grpc.Creds(creds), grpc.UnaryInterceptor(unaryInterceptor)}


	// create a gRPC server object
	//grpcServer := grpc.NewServer(opts...)

	// attach the Ping service to the server
	grpcServer := grpc.NewServer()  // attach the Ping service to the server
	pb.RegisterCDEServiceServer(grpcServer, &server)  // start the server
	log.Printf("starting HTTP/2 gRPC server on %s", address)
	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %s", err)
	}
	return nil
}

func startRESTServer(address, grpcAddress, certFile string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux(runtime.WithIncomingHeaderMatcher(credMatcher))
	//creds, err := credentials.NewClientTLSFromFile(certFile, "")
	//if err != nil {
	//	return fmt.Errorf("could not load TLS certificate: %s", err)
	//}  // Setup the client gRPC options
	//
	//opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}  // Register ping


	opts := []grpc.DialOption{grpc.WithInsecure()}  // Register ping
	err := pb.RegisterCDEServiceHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		return fmt.Errorf("could not register service Ping: %s", err)
	}

	log.Printf("starting HTTP/1.1 REST server on %s", address)
	http.ListenAndServe(address, mux)
	return nil
}

func getMongoCollection(dbName, collectionName, mongoHost string )  *mongo.Collection {

	// Register custom codecs for protobuf Timestamp and wrapper types
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mongoClient, err :=  mongo.Connect(ctx, options.Client().ApplyURI(mongoHost), options.Client().SetRegistry(bson.NewRegistryBuilder().
		RegisterDecoder(reflect.TypeOf(""), nullawareStrDecoder{}).
		Build(),))

	if err != nil {
		log.Println("Error while making collection obj ")
		log.Fatal(err)
	}
	return mongoClient.Database(dbName).Collection(collectionName)
}

func main()  {
	//grpcAddress := fmt.Sprintf("%s:%d", "cloudwalker.services.tv", 7775)
	//restAddress := fmt.Sprintf("%s:%d", "cloudwalker.services.tv", 7776)

	initializeProcess();

	serverhandler := Server{
		Tiles:         targetArray ,
	}

	grpcAddress := fmt.Sprintf(":%d",  7771)
	restAddress := fmt.Sprintf(":%d",  7772)
	certFile := "cert/server.crt"
	keyFile := "cert/server.key"

	// fire the gRPC server in a goroutine
	go func() {
		err := startGRPCServer(grpcAddress, certFile, keyFile, serverhandler)
		if err != nil {
			log.Fatalf("failed to start gRPC server: %s", err)
		}
	}()

	// fire the REST server in a goroutine
	go func() {
		err := startRESTServer(restAddress, grpcAddress, certFile)
		if err != nil {
			log.Fatalf("failed to start gRPC server: %s", err)
		}
	}()

	// infinite loop
	log.Printf("Entering infinite loop")
	select {}
}

func initializeProcess()  {

	fmt.Println("Welcome to init() function")
	primeTiles := getMongoCollection("cwtx2devel", "tiles", developmentMongoHost)
	loadingInToArray(primeTiles)

	//shemrootiles := getMongoCollection("tiles", "shemrootiles", developmentMongoHost)
	//loadingInToArray(shemrootiles)
	//
	//sonylivtiles := getMongoCollection("tiles", "sonylivtiles", developmentMongoHost)
	//loadingInToArray(sonylivtiles)
	//
	//hungamatiles := getMongoCollection("tiles", "hungamatiles", developmentMongoHost)
	//loadingInToArray(hungamatiles)
}

func loadingInToArray(tileCollection *mongo.Collection){
	// creating pipes for mongo aggregation
	myStages := mongo.Pipeline{}
	myStages = append(myStages, bson.D{{"$match", bson.D{{"content.publishState", true}}}})

	myStages = append(myStages, bson.D{{"$project", bson.D{
		{"_id", 0},
		{"ref_id", 1},
		{"metadata.title", 1},
		{"posters.landscape", 1},
		{"posters.portrait", 1},
		{"content.package", 1},
		{"content.source", 1},
		{"content.target", 1},
		{"created_at", 1},
		{"content.detailPage", 1},
		{"metadata.releaseDate", 1}}}} )


	cur, err := tileCollection.Aggregate(context.Background(), myStages)
	if err != nil {
		log.Println("Error while find ")
		log.Fatal(err)
	}

	for cur.Next(context.Background()) {
		var movieTile pb.MovieTile
		// converting curors to movieTiles
		err := cur.Decode(&movieTile)
		if err != nil {
			log.Fatal("Error decoding ************* ", err)
		}
		var contentTile pb.ContentTile
		contentTile.Title = movieTile.Metadata.Title
		contentTile.IsDetailPage = movieTile.Content.DetailPage
		if len(movieTile.Posters.Portrait) > 0 {
			contentTile.Portrait = movieTile.Posters.Portrait
		}

		if len(movieTile.Posters.Landscape) > 0 {
			contentTile.Poster = movieTile.Posters.Landscape
		}

		if len(movieTile.RefId) == 0 {
			movieTile.RefId = cur.Current.Lookup("ref_id").StringValue()
		}
		contentTile.ContentId = movieTile.RefId
		contentTile.Target = movieTile.Content.Target
		contentTile.RealeaseDate = movieTile.Metadata.ReleaseDate
		contentTile.PackageName = movieTile.Content.Package

		contentTile.TileType = pb.TileType_ImageTile
		// filling the target Array.
		targetArray = append(targetArray, contentTile)
	}
}