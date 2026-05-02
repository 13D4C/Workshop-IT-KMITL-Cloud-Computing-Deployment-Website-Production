package database

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/itkmitl/workshop-registration/internal/config"
	"github.com/itkmitl/workshop-registration/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect(cfg *config.Config) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSSLMode,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("✅ Database connected successfully")
}

func Migrate() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Activity{},
		&models.Registration{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("✅ Database migrated successfully")
}

func Seed() {

	var count int64
	DB.Model(&models.Activity{}).Count(&count)
	if count > 0 {
		log.Println("ℹ️  Seed data already exists, skipping...")
		return
	}

	activities := []models.Activity{
		{
			ID:               uuid.New(),
			Title:            "Secure Web Coding 101",
			Description:      "เรียนรู้พื้นฐานการเขียนเว็บอย่างปลอดภัย ครอบคลุม OWASP Top 10, XSS, SQL Injection, CSRF และ Security Best Practices พร้อมทำ Hands-on Lab จริง",
			ImageURL:         "/Workshop/Secure-Web-Coding-101.jpg",
			Location:         "ห้อง IT Auditorium ชั้น 3 อาคารคณะเทคโนโลยีสารสนเทศ",
			StartDate:        time.Now().AddDate(0, 0, 14),
			EndDate:          time.Now().AddDate(0, 0, 14).Add(6 * time.Hour),
			RegisterDeadline: time.Now().AddDate(0, 0, 10),
			MaxParticipants:  50,
			Status:           "open",
			Category:         "Workshop",
		},
		{
			ID:               uuid.New(),
			Title:            "IT Clash Competition",
			Description:      "การแข่งขันทักษะด้าน IT ชิงแชมป์ระดับคณะ ทดสอบความสามารถด้าน Programming, Networking, Database และ Problem Solving ชิงเงินรางวัลรวมกว่า 50,000 บาท",
			ImageURL:         "/Compettion/ITclash.jpg",
			Location:         "IT Co-working Space ชั้น 1 อาคารคณะเทคโนโลยีสารสนเทศ",
			StartDate:        time.Now().AddDate(0, 1, 0),
			EndDate:          time.Now().AddDate(0, 1, 1),
			RegisterDeadline: time.Now().AddDate(0, 0, 25),
			MaxParticipants:  100,
			Status:           "open",
			Category:         "Competition",
		},
		{
			ID:               uuid.New(),
			Title:            "IT Camp",
			Description:      "ค่ายเตรียมความพร้อมสำหรับนักศึกษาใหม่คณะ IT เรียนรู้ทักษะพื้นฐาน Programming, Teamwork และ Creative Thinking พร้อมกิจกรรมสร้างสรรค์มากมาย",
			ImageURL:         "/Camp/ITcamp.jpg",
			Location:         "อาคารคณะเทคโนโลยีสารสนเทศ สจล.",
			StartDate:        time.Now().AddDate(0, 0, 21),
			EndDate:          time.Now().AddDate(0, 0, 23),
			RegisterDeadline: time.Now().AddDate(0, 0, 17),
			MaxParticipants:  80,
			Status:           "open",
			Category:         "Camp",
		},
		{
			ID:               uuid.New(),
			Title:            "To Be IT #69",
			Description:      "กิจกรรมแนะนำคณะ IT KMITL สำหรับน้อง ๆ ม.ปลาย ที่สนใจเข้าศึกษาต่อ พาทัวร์คณะ พบรุ่นพี่ และทดลองเรียน Workshop จริง",
			ImageURL:         "/Camp/TobeIT69.jpg",
			Location:         "ห้อง IT Auditorium ชั้น 3 อาคารคณะเทคโนโลยีสารสนเทศ",
			StartDate:        time.Now().AddDate(0, 0, 7),
			EndDate:          time.Now().AddDate(0, 0, 7).Add(6 * time.Hour),
			RegisterDeadline: time.Now().AddDate(0, 0, 5),
			MaxParticipants:  120,
			Status:           "open",
			Category:         "Camp",
		},
		{
			ID:               uuid.New(),
			Title:            "Unite Camp",
			Description:      "ค่ายรวมพลังนักศึกษา IT ทุกชั้นปี สร้างความสัมพันธ์ เรียนรู้การทำงานเป็นทีม และพัฒนา Soft Skills ที่จำเป็นสำหรับการทำงานในอนาคต",
			ImageURL:         "/Camp/Unite.jpg",
			Location:         "อาคารคณะเทคโนโลยีสารสนเทศ สจล.",
			StartDate:        time.Now().AddDate(0, 0, 28),
			EndDate:          time.Now().AddDate(0, 0, 29),
			RegisterDeadline: time.Now().AddDate(0, 0, 24),
			MaxParticipants:  60,
			Status:           "open",
			Category:         "Camp",
		},
		{
			ID:               uuid.New(),
			Title:            "Pre-Programming Camp",
			Description:      "ค่ายเตรียมความพร้อมด้าน Programming สำหรับนักศึกษาชั้นปีที่ 1 เรียนรู้พื้นฐาน Python, Algorithm และ Data Structure เพื่อเตรียมตัวสำหรับวิชาเรียน",
			ImageURL:         "/Camp/Prepro.jpeg",
			Location:         "Lab 2 ชั้น 4 อาคารคณะเทคโนโลยีสารสนเทศ",
			StartDate:        time.Now().AddDate(0, 0, 10),
			EndDate:          time.Now().AddDate(0, 0, 11),
			RegisterDeadline: time.Now().AddDate(0, 0, 9),
			MaxParticipants:  40,
			Status:           "open",
			Category:         "Camp",
		},
	}

	for _, a := range activities {
		DB.Create(&a)
	}

	log.Println("✅ Seed data created successfully")
}
