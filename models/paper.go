package models

type Paper struct {
	PaperNumber         string `json:"paper_number"`
	PaperName           string `json:"paper_name"`
	PaperTopic          string `json:"paper_topic"`
	PaperContent        string `json:"paper_content"`
	RelatedCode         string `json:"related_code"`
	DeviceNumberModelId string `json:"device_number_model_id"`
}
