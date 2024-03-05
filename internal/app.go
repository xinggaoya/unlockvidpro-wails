package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

// GetAppVersion returns a greeting for the given name
func (a *App) GetAppVersion() string {
	// 获取 wails.json 版本号
	return fmt.Sprintf("Hello %s, It's show time!")
}

// DailyQuote returns a daily quote
func (a *App) DailyQuote() ResBody {
	resp, err := http.Get("https://apiv3.shanbay.com/weapps/dailyquote/quote")
	if err != nil {
		fmt.Errorf("http.Get() failed with '%s'\n", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Errorf("io.ReadAll() failed with '%s'\n", err)
	}

	var res ResBody
	err = json.Unmarshal(body, &res)

	fmt.Printf("res: %v\n", res.Content)
	if err != nil {
		fmt.Errorf("json.Unmarshal() failed with '%s'\n", err)
	}

	return res
}

type ResBody struct {
	Id         string `json:"id"`
	Content    string `json:"content"`
	Author     string `json:"author"`
	AssignDate string `json:"assign_date"`
	AdUrl      string `json:"ad_url"`
	ShareUrl   string `json:"share_url"`
	ShareUrls  struct {
	} `json:"share_urls"`
	OriginImgUrls []string `json:"origin_img_urls"`
	ShareImgUrls  []string `json:"share_img_urls"`
	JoinNum       int      `json:"join_num"`
	Translation   string   `json:"translation"`
	PosterImgUrls []string `json:"poster_img_urls"`
	TrackObject   struct {
	} `json:"track_object"`
	DailyAudioUrls string `json:"daily_audio_urls"`
}
