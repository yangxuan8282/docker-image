// Package base declares the major types for static site
package base

type (
	// Config is basic options for compiling site
	Config struct {
		PostDir      string `toml:"post_dir"`
		PageDir      string `toml:"page_dir"`
		MediaDir     string `toml:"media_dir"`
		LangDir      string `toml:"lang_dir"`
		ThemeDir     string `toml:"theme_dir"`
		OutputDir    string `toml:"output_dir"`
		PostPageSize int    `toml:"post_pagesize"`
	}
)
