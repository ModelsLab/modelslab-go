# ModelsLab Go SDK

Official Go SDK for [ModelsLab API](https://modelslab.com) - Generate AI content including images, videos, audio, 3D models, and more.

## Features

**Text-to-Image & Image-to-Image** generation
**Text-to-Speech & Music** generation
**Face Swap & Deepfake** operations
**Interior Design & 3D** modeling
**Image Editing** (upscaling, background removal, etc.)
**Realtime** generation APIs
**Enterprise** features support

## Quick Start

### 1. Installation

```bash
go mod init your-project
go get github.com/modelslab/modelslab-go
```

### 2. Get Your API Key

1. Sign up at [ModelsLab.com](https://modelslab.com)
2. Go to your dashboard
3. get your API key

### 3. Basic Usage with community models

```go
import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelslab/modelslab-go/pkg/apis/community"
	"github.com/modelslab/modelslab-go/pkg/client"
	communitySchema "github.com/modelslab/modelslab-go/pkg/schemas/community"
)

func main() {
	c := client.New("your-api-key")
	api := community.New(c, false)

	model := "midjourney"
	req := &communitySchema.Text2ImageRequest{
		Prompt:  "a cat",
		ModelID: &model,
	}

	resp, err := api.TextToImage(context.Background(), &req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	out, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(out))
}
```

### 4. Run Your Code

```bash
go run main.go
```

## API Examples

### Text-to-Speech

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelslab/modelslab-go/pkg/apis/audio"
	"github.com/modelslab/modelslab-go/pkg/client"
	audioSchema "github.com/modelslab/modelslab-go/pkg/schemas/audio"
)

func main() {
	c := client.New("your-api-key")
	api := audio.New(c, false)

	voice_id := "madison"
	language := "english"
	req := audioSchema.Text2SpeechRequest{
		Prompt:   "a cat sitting on a mat",
		VoiceID:  &voice_id,
		Language: &language,
	}

	resp, err := api.TextToSpeech(context.Background(), &req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	out, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(out))
}
```

### Multiple Face Swap

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelslab/modelslab-go/pkg/apis/deepfake"
	"github.com/modelslab/modelslab-go/pkg/client"
	"github.com/modelslab/modelslab-go/pkg/schemas/base"
	deepfakeSchema "github.com/modelslab/modelslab-go/pkg/schemas/deepfake"
)

func main() {
	c := client.New("your-api-key")
	deepfakeAPI := deepfake.New(c, false)

	initImageURL := "https://i.pinimg.com/564x/4c/6a/d0/4c6ad0f74a3a251344cb115699a9a7c9.jpg"
	targetImageURL := "https://i.pinimg.com/564x/11/ac/0d/11ac0ddaf6962e395f30abc61043393e.jpg"

	req := deepfakeSchema.MultipleFaceSwapRequest{
		InitImage: base.FileInput{
			URL: &initImageURL,
		},
		TargetImage: base.FileInput{
			URL: &targetImageURL,
		},
	}

	resp, err := deepfakeAPI.MultipleFaceSwap(context.Background(), &req)
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
	videoSchema "github.com/modelslab/modelslab-go/pkg/schemas/video"
)

func main() {
	c := client.New("your-api-key")
	videoAPI := video.New(c, false)

	req := videoSchema.Text2VideoRequest{
		Prompt:  "A cat playing with a ball in a sunny garden",
		ModelID: "cogvideox",
	}

	resp, err := videoAPI.TextToVideo(context.Background(), &req)
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

## Configuration

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

## Response Format

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

## File Input Options

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

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Links

- [ModelsLab Website](https://modelslab.com)
- [API Documentation](https://docs.modelslab.com)
- [Get API Key](https://modelslab.com/dashboard)
- [Discord Community](https://discord.com/invite/modelslab-1033301189254729748)

## Support

- Email: mailto:support@modelslab.com
- Discord: [Join our community](https://discord.com/invite/modelslab-1033301189254729748)
- Documentation: [modelslab.com/docs](https://docs.modelslab.com)

---