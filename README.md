# ModelsLab Go SDK

Official Go SDK for [ModelsLab API](https://modelslab.com) - Generate AI content including images, videos, audio, 3D models, and more.

## ✨ Features

- 🖼️ **Text-to-Image & Image-to-Image** generation
- 🎥 **Text-to-Video & Image-to-Video** creation  
- 🎵 **Text-to-Speech & Music** generation
- 🎭 **Face Swap & Deepfake** operations
- 🏠 **Interior Design & 3D** modeling
- 🎨 **Image Editing** (upscaling, background removal, etc.)
- ⚡ **Realtime** generation APIs
- 🔧 **Enterprise** features support

## 🚀 Quick Start

### 1. Installation

```bash
go mod init your-project
go get github.com/modelslab/modelslab-go
```

### 2. Get Your API Key

1. Sign up at [ModelsLab.com](https://modelslab.com)
2. Go to your dashboard
3. Copy your API key

### 3. Basic Usage

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/modelslab/modelslab-go/pkg/apis/community"
	"github.com/modelslab/modelslab-go/pkg/client"
	"github.com/modelslab/modelslab-go/pkg/schemas/community"
)

func main() {
	// Initialize client with your API key
	apiKey := "your-api-key-here"
	c := client.New(apiKey)

	// Create community API instance
	communityAPI := community.New(c, false)

	// Generate an image
	prompt := "A beautiful sunset over mountains"
	req := &community.Text2ImageRequest{
		Prompt: prompt,
	}

	// Call the API
	resp, err := communityAPI.TextToImage(context.Background(), req)
	if err != nil {
		log.Fatalf("API call failed: %v", err)
	}

	// Print the complete response
	prettyJSON, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(prettyJSON))
}
```

### 4. Run Your Code

```bash
go run main.go
```

## 📚 API Examples

### Text-to-Speech

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelslab/modelslab-go/pkg/apis/audio"
	"github.com/modelslab/modelslab-go/pkg/client"
	"github.com/modelslab/modelslab-go/pkg/schemas/audio"
)

func main() {
	c := client.New("your-api-key")
	audioAPI := audio.New(c, false)

	language := "english"
	voiceID := "madison"
	speed := 1.0

	req := &audio.Text2SpeechRequest{
		Prompt:   "Hello, this is a test of ModelsLab text-to-speech!",
		Language: &language,
		VoiceID:  &voiceID,
		Speed:    &speed,
	}

	resp, err := audioAPI.TextToSpeech(context.Background(), req)
	if err != nil {
		panic(err)
	}

	// Response contains audio URL, processing info, etc.
	prettyJSON, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(prettyJSON))
}
```

### Face Swap

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelslab/modelslab-go/pkg/apis/deepfake"
	"github.com/modelslab/modelslab-go/pkg/client"
	"github.com/modelslab/modelslab-go/pkg/schemas/base"
	"github.com/modelslab/modelslab-go/pkg/schemas/deepfake"
)

func main() {
	c := client.New("your-api-key")
	deepfakeAPI := deepfake.New(c, false)

	initImageURL := "https://example.com/person1.jpg"
	targetImageURL := "https://example.com/person2.jpg"

	req := &deepfake.MultipleFaceSwapRequest{
		InitImage: base.FileInput{
			URL: &initImageURL,
		},
		TargetImage: base.FileInput{
			URL: &targetImageURL,
		},
	}

	resp, err := deepfakeAPI.MultipleFaceSwap(context.Background(), req)
	if err != nil {
		panic(err)
	}

	prettyJSON, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(prettyJSON))
}
```

### Text-to-Video

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelslab/modelslab-go/pkg/apis/video"
	"github.com/modelslab/modelslab-go/pkg/client"
	"github.com/modelslab/modelslab-go/pkg/schemas/video"
)

func main() {
	c := client.New("your-api-key")
	videoAPI := video.New(c, false)

	req := &video.Text2VideoRequest{
		Prompt: "A cat playing with a ball in a sunny garden",
	}

	resp, err := videoAPI.TextToVideo(context.Background(), req)
	if err != nil {
		panic(err)
	}

	prettyJSON, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(prettyJSON))
}
```

## 🛠️ Available APIs

| API | Description | Package |
|-----|-------------|---------|
| **Community** | Text-to-Image, Image-to-Image, Inpainting, ControlNet | `github.com/modelslab/modelslab-go/pkg/apis/community` |
| **Audio** | Text-to-Speech, Music Generation, Voice Cloning | `github.com/modelslab/modelslab-go/pkg/apis/audio` |
| **Video** | Text-to-Video, Image-to-Video | `github.com/modelslab/modelslab-go/pkg/apis/video` |
| **Deepfake** | Face Swap, Video Face Swap | `github.com/modelslab/modelslab-go/pkg/apis/deepfake` |
| **Image Editing** | Super Resolution, Background Removal, Outpainting | `github.com/modelslab/modelslab-go/pkg/apis/image_editing` |
| **Interior** | Interior Design, Room Decoration | `github.com/modelslab/modelslab-go/pkg/apis/interior` |
| **3D** | Text-to-3D, Image-to-3D | `github.com/modelslab/modelslab-go/pkg/apis/threed` |
| **Realtime** | Real-time Image Generation | `github.com/modelslab/modelslab-go/pkg/apis/realtime` |

## 🔧 Configuration

### Basic Client

```go
import "github.com/modelslab/modelslab-go/pkg/client"

// Simple client
c := client.New("your-api-key")
```

### Custom Configuration

```go
import (
	"time"
	"github.com/modelslab/modelslab-go/pkg/client"
)

config := &client.Config{
	APIKey:       "your-api-key",
	BaseURL:      "https://modelslab.com/api/",
	FetchRetry:   10,
	FetchTimeout: 2 * time.Second,
	HTTPTimeout:  30 * time.Second,
}

c := client.NewWithConfig(config)
```

### Enterprise Mode

```go
// For enterprise users
communityAPI := community.New(c, true) // true = enterprise mode
```

## 📝 Response Format

The SDK returns **complete raw API responses** as `map[string]interface{}` to preserve all fields:

```json
{
  "status": "success",
  "message": "Image generated successfully",
  "output": ["https://example.com/generated-image.jpg"],
  "id": 12345,
  "meta": {
    "prompt": "A beautiful sunset",
    "model": "stable-diffusion",
    "steps": 20,
    "seed": 12345
  },
  "generationTime": 5.2,
  "proxy_links": ["https://cdn.example.com/image.jpg"]
}
```

You get **ALL** fields returned by the API, including metadata, generation time, proxy links, and any future fields.

## 🔍 File Input Options

The SDK supports multiple ways to provide images/audio:

```go
import "github.com/modelslab/modelslab-go/pkg/schemas/base"

// URL input
imageURL := "https://example.com/image.jpg"
fileInput := base.FileInput{
    URL: &imageURL,
}

// Base64 input
base64Data := "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQ..."
fileInput := base.FileInput{
    Base64: &base64Data,
}

// File path (for local files)
filePath := "/path/to/image.jpg"
fileInput := base.FileInput{
    FilePath: &filePath,
}
```

## ⚡ Async Operations

Many APIs return a processing status. Use the fetch API to get results:

```go
// Initial request returns processing status
resp, _ := videoAPI.TextToVideo(context.Background(), req)

// Extract fetch URL from response
fetchURL := (*resp)["fetch_result"].(string)

// Poll for completion (implement your own retry logic)
time.Sleep(10 * time.Second)
// Use base API fetch method or make direct HTTP calls
```

## 🚨 Error Handling

```go
resp, err := communityAPI.TextToImage(context.Background(), req)
if err != nil {
    log.Printf("API error: %v", err)
    return
}

// Check response status
if status, ok := (*resp)["status"].(string); ok && status == "error" {
    if message, ok := (*resp)["message"].(string); ok {
        log.Printf("API returned error: %s", message)
    }
}
```

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🔗 Links

- [ModelsLab Website](https://modelslab.com)
- [API Documentation](https://modelslab.com/docs)
- [Get API Key](https://modelslab.com/dashboard)
- [Discord Community](https://discord.gg/modelslab)

## 📞 Support

- 📧 Email: support@modelslab.com
- 💬 Discord: [Join our community](https://discord.gg/modelslab)
- 📖 Documentation: [modelslab.com/docs](https://modelslab.com/docs)

---

Made with ❤️ by the ModelsLab team