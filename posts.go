package ghostapi

// Posts struct
type Posts struct {
	Posts []struct {
		// ID                     string      `json:"id"`
		// UUID                   string      `json:"uuid"`
		Title string `json:"title"`
		Slug  string `json:"slug"`
		// HTML                   string      `json:"html"`
		// CommentID              string      `json:"comment_id"`
		FeatureImage string `json:"feature_image"`
		Featured     bool   `json:"featured"`
		Visibility   string `json:"visibility"`
		CreatedAt    string `json:"created_at"`
		UpdatedAt    string `json:"updated_at"`
		PublishedAt  string `json:"published_at"`
		// CustomExcerpt          string `json:"custom_excerpt"`
		// CodeinjectionHead      string `json:"codeinjection_head"`
		// CodeinjectionFoot      string `json:"codeinjection_foot"`
		// CustomTemplate         string `json:"custom_template"`
		// CanonicalURL           string `json:"canonical_url"`
		// SendEmailWhenPublished bool        `json:"send_email_when_published"`
		URL string `json:"url"`
		// Excerpt                string      `json:"excerpt"`
		// ReadingTime            int         `json:"reading_time"`
		// Access                 bool        `json:"access"`
		OgImage            string `json:"og_image"`
		OgTitle            string `json:"og_title"`
		OgDescription      string `json:"og_description"`
		TwitterImage       string `json:"twitter_image"`
		TwitterTitle       string `json:"twitter_title"`
		TwitterDescription string `json:"twitter_description"`
		MetaTitle          string `json:"meta_title"`
		MetaDescription    string `json:"meta_description"`
		// EmailSubject           string `json:"email_subject"`
	} `json:"posts"`
	Meta struct {
		Pagination struct {
			Page  int `json:"page"`
			Limit int `json:"limit"`
			Pages int `json:"pages"`
			Total int `json:"total"`
			Next  int `json:"next"`
			Prev  int `json:"prev"`
		} `json:"pagination"`
	} `json:"meta"`
}
