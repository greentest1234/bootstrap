package models

//import "strings"

type Service struct {
	ProjectID                string  `json:"project_id"`
	ServiceID                string  `json:"id"`
	BuildPackID              string  `json:"buildpack_id,omitempty"`
	IconClass                string  `json:"icon_class"`
	Name                     string  `json:"name"`
	ImageSource              string  `json:"image_source"`
	BuildImage               string  `json:"build_image"`
	DeployImage              string  `json:"deploy_image"`
	Repository               string  `json:"repository"`
	Organization             string  `json:"organization"`
	Public                   bool    `json:"public"`
	CPU                      float32 `json:"default_cpu"`
	RAM                      float32 `json:"default_cpu"`
	BuildCommand             string  `json:"build_command"`
	TestCommand              string  `json:"test_command"`
	ContainerPort            int     `json:"container_port"`
	ContainerSharedDirectory string  `json:"-"`
}

func NewService() *Service {
	return &Service{}
}

func (ref *Service) SshUrl() string {
	//return ref.Repository

	//for ssh repo git ops
	return strings.Replace(ref.Repository, "https://github.com/", "git@github.com:", -1)
}
