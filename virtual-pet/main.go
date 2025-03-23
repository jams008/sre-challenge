package main

import (
	"context"
	"os"
	"virtual-pet/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	happinessMetric = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pet_happiness",
		Help: "Current happiness level of pets",
	}, []string{"pet_id"})
	hungerMetric = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pet_hunger",
		Help: "Current hunger level of pets",
	}, []string{"pet_id"})
	energyMetric = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pet_energy",
		Help: "Current energy level of pets",
	}, []string{"pet_id"})
	getSinglePetCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "get_single_pet_requests_total",
		Help: "Total number of requests to get a single pet",
	})
	getAllPetsCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "get_all_pets_requests_total",
		Help: "Total number of requests to get all pets",
	})
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// Get MongoDB URI, database, and collection from environment variables
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		panic("MONGODB_URI not set in .env file")
	}
	mongoDatabase := os.Getenv("MONGODB_DATABASE")
	if mongoDatabase == "" {
		panic("MONGODB_DATABASE not set in .env file")
	}
	mongoCollection := os.Getenv("MONGODB_COLLECTION")
	if mongoCollection == "" {
		panic("MONGODB_COLLECTION not set in .env file")
	}

	// MongoDB connection
	client, _ := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	collection := client.Database(mongoDatabase).Collection(mongoCollection)

	// Initialize metrics for all pets
	initializeMetrics(collection)

	app := fiber.New()

	// Add logging middleware
	app.Use(logger.New())

	// Add recover middleware
	app.Use(recover.New())

	// Create pet
	app.Post("/pets", func(c *fiber.Ctx) error {
		pet := new(models.Pet)
		if err := c.BodyParser(pet); err != nil {
			return err
		}
		pet.ID = uuid.New().String()

		_, err := collection.InsertOne(context.Background(), pet)
		if err != nil {
			return err
		}

		// Update metrics
		happinessMetric.WithLabelValues(pet.ID).Set(pet.Happiness)
		hungerMetric.WithLabelValues(pet.ID).Set(pet.Hunger)
		energyMetric.WithLabelValues(pet.ID).Set(pet.Energy)

		return c.JSON(pet)
	})

	// Get single pet stats
	app.Get("/pets/:id", func(c *fiber.Ctx) error {
		getSinglePetCounter.Inc()
		id := c.Params("id")
		var pet models.Pet
		err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&pet)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Pet not found"})
		}

		// Update metrics
		happinessMetric.WithLabelValues(pet.ID).Set(pet.Happiness)
		hungerMetric.WithLabelValues(pet.ID).Set(pet.Hunger)
		energyMetric.WithLabelValues(pet.ID).Set(pet.Energy)

		return c.JSON(pet)
	})

	// Get all pets stats
	app.Get("/pets", func(c *fiber.Ctx) error {
		getAllPetsCounter.Inc()
		cursor, err := collection.Find(context.Background(), bson.D{})
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Database error"})
		}
		defer cursor.Close(context.Background())

		var pets []models.Pet
		if err = cursor.All(context.Background(), &pets); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Error parsing pets"})
		}

		// Update metrics for all pets
		for _, pet := range pets {
			happinessMetric.WithLabelValues(pet.ID).Set(pet.Happiness)
			hungerMetric.WithLabelValues(pet.ID).Set(pet.Hunger)
			energyMetric.WithLabelValues(pet.ID).Set(pet.Energy)
		}

		return c.JSON(pets)
	})

	// Update pet stats
	app.Put("/pets/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var updatePet models.Pet
		if err := c.BodyParser(&updatePet); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
		}

		update := bson.M{
			"$set": bson.M{
				"happiness": updatePet.Happiness,
				"hunger":    updatePet.Hunger,
				"energy":    updatePet.Energy,
			},
		}

		result, err := collection.UpdateOne(
			context.Background(),
			bson.M{"_id": id},
			update,
		)

		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Update failed"})
		}

		if result.ModifiedCount == 0 {
			return c.Status(404).JSON(fiber.Map{"error": "Pet not found"})
		}

		// Retrieve the updated pet
		err = collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&updatePet)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Error retrieving updated pet"})
		}

		// Update metrics
		happinessMetric.WithLabelValues(id).Set(updatePet.Happiness)
		hungerMetric.WithLabelValues(id).Set(updatePet.Hunger)
		energyMetric.WithLabelValues(id).Set(updatePet.Energy)

		return c.JSON(updatePet)
	})

	// Delete pet
	app.Delete("/pets/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		result, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Delete failed"})
		}

		if result.DeletedCount == 0 {
			return c.Status(404).JSON(fiber.Map{"error": "Pet not found"})
		}

		// Remove metrics
		happinessMetric.DeleteLabelValues(id)
		hungerMetric.DeleteLabelValues(id)
		energyMetric.DeleteLabelValues(id)

		return c.SendStatus(204)
	})

	// Prometheus metrics endpoint
	app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))

	// Panic test endpoint
	app.Get("/panic", func(c *fiber.Ctx) error {
		panic("This is a test panic")
	})

	app.Listen(":3000")
}

func initializeMetrics(collection *mongo.Collection) {
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		panic("Failed to initialize metrics: " + err.Error())
	}
	defer cursor.Close(context.Background())

	var pets []models.Pet
	if err = cursor.All(context.Background(), &pets); err != nil {
		panic("Failed to parse pets for metrics initialization: " + err.Error())
	}

	// Update metrics for all pets
	for _, pet := range pets {
		happinessMetric.WithLabelValues(pet.ID).Set(pet.Happiness)
		hungerMetric.WithLabelValues(pet.ID).Set(pet.Hunger)
		energyMetric.WithLabelValues(pet.ID).Set(pet.Energy)
	}
}
