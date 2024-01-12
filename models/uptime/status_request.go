package uptime

// StatusPageReqBody represents the request body for creating/updating a status page.
type StatusPageReqBody struct {
	// Required: CompanyName: Name of your company.
	CompanyName string `json:"company_name"`

	// Required: Timezone: The accepted values can be found in the Rails TimeZone documentation. https://api.rubyonrails.org/classes/ActiveSupport/TimeZone.html
	Timezone string `json:"timezone"`

	// Required: Subdomain: What subdomain should we use for your status page?
	// This needs to be unique across our entire application, so choose carefully. (only alphanumeric characters, dashes, and underscores are allowed)
	Subdomain string `json:"subdomain"`

	// History: Number of days to display on the status page. Minimum 90 days.
	History int `json:"history,omitempty"`

	// CompanyURL: URL of your company's website.
	CompanyURL string `json:"company_url,omitempty"`

	// ContactURL: URL that should be used for contacting you in case of an emergency.
	ContactURL string `json:"contact_url,omitempty"`

	// LogoURL: A direct link to your company's logo. The image should be under 20MB in size.
	LogoURL string `json:"logo_url,omitempty"`

	// CustomDomain: Do you want a custom domain on your status page?
	// Add a CNAME record that points your domain to status.betteruptime.com
	// Example: CNAME status.walmine.com statuspage.betteruptime.com
	CustomDomain string `json:"custom_domain,omitempty"`

	// MinIncidentLength: If you don't want to display short incidents on your status page, this attribute is for you. In seconds.
	MinIncidentLength int `json:"min_incident_length,omitempty"`

	// Subscribable: Do you want to allow users to subscribe to your status page changes?
	Subscribable bool `json:"subscribable,omitempty"`

	// HideFromSearchEngines: Hide your status page from search engines.
	HideFromSearchEngines bool `json:"hide_from_search_engines,omitempty"`

	// CustomCSS: Unleash your inner designer and tweak our status page design to fit your branding.
	CustomCSS string `json:"custom_css,omitempty"`

	// GoogleAnalyticsID: Specify your own Google Analytics ID if you want to receive hits on your status page.
	GoogleAnalyticsID string `json:"google_analytics_id,omitempty"`

	// Announcement: Add an announcement to your status page.
	Announcement string `json:"announcement,omitempty"`

	// AnnouncementEmbedVisible: Toggle this field if you want to show an announcement in your embed.
	AnnouncementEmbedVisible bool `json:"announcement_embed_visible,omitempty"`

	// AnnouncementEmbedLink: Point your embedded announcement to a specified URL.
	AnnouncementEmbedLink string `json:"announcement_embed_link,omitempty"`

	// AnnouncementEmbedCustomCSS: Modify the design of the announcement embed.
	AnnouncementEmbedCustomCSS string `json:"announcement_embed_custom_css,omitempty"`

	// PasswordEnabled: Do you want to enable password protection on your status page?
	PasswordEnabled bool `json:"password_enabled,omitempty"`

	// Password: Set a password for your status page (we won't store it as plaintext, promise).
	// Required when PasswordEnabled is true.
	// We will set PasswordEnabled to false automatically when you send us an empty password.
	Password string `json:"password,omitempty"`
}
