package structs

/* IG API User Timeline Structs */

// IGAPITimeline gives access to the timeline data.
type IGAPITimeline struct {
	Data igAPITimelineData `json:"data"`
}

type igAPITimelineData struct {
	User igAPITimelineUser `json:"user"`
}

type igAPITimelineUser struct {
	Media igAPITimelineMedia `json:"edge_owner_to_timeline_media"`
}

type igAPITimelineMedia struct {
	Count    int                        `json:"count"`
	Edges    []igAPITimelineMediaEdge   `json:"edges"`
	PageInfo igAPITimelineMediaPageInfo `json:"page_info"`
}

type igAPITimelineMediaPageInfo struct {
	HasNextPage bool   `json:"has_next_page"`
	EndCursor   string `json:"end_cursor"`
}

type igAPITimelineMediaEdge struct {
	Node igAPITimelineMediaEdgeNode `json:"node"`
}

type igAPITimelineMediaEdgeNode struct {
	Shortcode string `json:"shortcode"`
}

/* IGAPI User Metadata Structs */

// IGAPIUserMetadata gives access to the user metadata.
type IGAPIUserMetadata struct {
	GraphQL igAPIUserGraphQL `json:"graphql"`
}

type igAPIUserGraphQL struct {
	User igAPIUser `json:"user"`
}

type igAPIUser struct {
	ID string `json:"id"`
}

/* IGAPI Media Structs */

// IGAPIMediaDetail is the struct that is exported and converted to the target struct.
type IGAPIMediaDetail struct {
	GraphQL igAPIMediaGraphQL `json:"graphql"`
}

type igAPIMediaGraphQL struct {
	Media igAPIMediaShortcode `json:"shortcode_media"`
}

type igAPIMediaShortcode struct {
	ID             string                          `json:"id"`
	Shortcode      string                          `json:"shortcode"`
	DisplayURL     string                          `json:"display_url"`
	Timestamp      int                             `json:"taken_at_timestamp"`
	Location       igAPIMediaShortcodeLocation     `json:"location"`
	EdgesToCaption igAPIMediaCaptionEdgesToCaption `json:"edge_media_to_caption"`
}

type igAPIMediaShortcodeLocation struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type igAPIMediaCaptionEdgesToCaption struct {
	Edges []igAPIMediaCaptionEdges `json:"edges"`
}

type igAPIMediaCaptionEdges struct {
	Node igAPIMediaCaptionEdgeNode `json:"node"`
}

type igAPIMediaCaptionEdgeNode struct {
	Text string `json:"text"`
}
