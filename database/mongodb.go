package database

import (
    "context"
    "log"
    "os"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var (
    MongoDB          *mongo.Database
    MongoClient      *mongo.Client
    AchievementsCol  *mongo.Collection
)

func ConnectMongo() {
    mongoURI := os.Getenv("MONGO_URI")
    if mongoURI == "" {
        mongoURI = "mongodb://localhost:27017"
    }

    // connect
    clientOptions := options.Client().ApplyURI(mongoURI)
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal("MongoDB connection failed:", err)
    }

    // ping
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    if err := client.Ping(ctx, nil); err != nil {
        log.Fatal("MongoDB ping failed:", err)
    }

    MongoClient = client
    MongoDB = client.Database("prestasi_mahasiswa")

    // collections
    AchievementsCol = MongoDB.Collection("achievements")

    log.Println("✅ MongoDB connected successfully")
}
