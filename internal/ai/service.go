package ai

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// Service mendefinisikan antarmuka untuk layanan AI.
type Service interface {
	SummarizePDF(ctx context.Context, file *multipart.FileHeader) (string, error)
}

type service struct {
	genaiClient *genai.Client
}

// NewService sekarang menginisialisasi client Go SDK resmi dari Google.
func NewService(ctx context.Context) (Service, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("GEMINI_API_KEY environment variable not set")
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to create genai client: %w", err)
	}

	return &service{genaiClient: client}, nil
}

// SummarizePDF sekarang menggunakan metode yang benar sesuai dokumentasi resmi.
func (s *service) SummarizePDF(ctx context.Context, fileHeader *multipart.FileHeader) (string, error) {
	// 1. Buka file yang di-upload oleh user
	file, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open uploaded file: %w", err)
	}
	defer file.Close()

	// 2. Baca konten file ke dalam byte slice
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read file content: %w", err)
	}

	// 3. Upload file menggunakan SDK.
	uploadCtx, cancel := context.WithTimeout(ctx, time.Minute*5) // Timeout 5 menit untuk upload
	defer cancel()

	// PERBAIKAN: Membuat struct UploadFileOptions dan mengisinya.
	// Kita akan memberikan nama file yang unik dan tipe MIME yang benar.
	opts := &genai.UploadFileOptions{
		MIMEType:    "application/pdf",
		DisplayName: fileHeader.Filename,
	}

	// PERBAIKAN: Memanggil UploadFile dengan parameter yang benar.
	// Parameter 'name' bisa dikosongkan agar dibuat otomatis oleh Google.
	uploadedFile, err := s.genaiClient.UploadFile(uploadCtx, "", bytes.NewReader(fileBytes), opts)
	if err != nil {
		return "", fmt.Errorf("SDK failed to upload file: %w", err)
	}

	model := s.genaiClient.GenerativeModel("gemini-2.5-pro")

	prompt := genai.Text("Anda adalah asisten yang ahli dalam merangkum dokumen pengadaan/tender. Tolong rangkum proposal dalam file PDF yang terlampir ini dengan jelas, singkat, dan objektif. Fokus pada poin-poin utama seperti solusi yang ditawarkan, keunggulan, dan estimasi biaya jika ada. Jawab langsung, to the point, tanpa basa basi, tanpa kalimat pembuka. (Buat output nya menjadi paragraph biasa) ")

	fileData := genai.FileData{URI: uploadedFile.URI}

	// 6. Kirim permintaan ke Gemini untuk menghasilkan konten
	genCtx, cancel := context.WithTimeout(ctx, time.Minute*2) // Timeout 2 menit untuk generasi konten
	defer cancel()
	resp, err := model.GenerateContent(genCtx, prompt, fileData)
	if err != nil {
		return "", fmt.Errorf("failed to generate content: %w", err)
	}

	// 7. Ekstrak dan kembalikan hasil teks dari respons
	if len(resp.Candidates) > 0 && resp.Candidates[0].Content != nil && len(resp.Candidates[0].Content.Parts) > 0 {
		if textPart, ok := resp.Candidates[0].Content.Parts[0].(genai.Text); ok {
			return string(textPart), nil
		}
	}

	return "", fmt.Errorf("no text content found in Gemini response")
}
