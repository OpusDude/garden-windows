package backend

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"strings"

	"github.com/cloudfoundry-incubator/garden/api"
	"github.com/pivotal-cf-experimental/garden-dot-net/container"
)

type dotNetBackend struct {
	containerizerURL url.URL
}

func NewDotNetBackend(containerizerURL string) (*dotNetBackend, error) {
	u, err := url.Parse(containerizerURL)
	if err != nil {
		return nil, err
	}
	return &dotNetBackend{
		containerizerURL: *u,
	}, nil
}

func (dotNetBackend *dotNetBackend) ContainerizerURL() string {
	return dotNetBackend.containerizerURL.String()
}

func (dotNetBackend *dotNetBackend) Start() error {
	return nil
}

func (dotNetBackend *dotNetBackend) Stop() {}

func (dotNetBackend *dotNetBackend) GraceTime(api.Container) time.Duration {
	return time.Second
}

func (dotNetBackend *dotNetBackend) Ping() error {
	return nil
}

func (dotNetBackend *dotNetBackend) Capacity() (api.Capacity, error) {
	capacity := api.Capacity{}
	return capacity, nil
}

func (dotNetBackend *dotNetBackend) Create(containerSpec api.ContainerSpec) (api.Container, error) {
	netContainer := container.NewContainer(dotNetBackend.containerizerURL, "containerhandle")
	url := dotNetBackend.containerizerURL.String() + "/api/containers"
	containerSpecJSON, err := json.Marshal(containerSpec)
	if err != nil {
		return nil, err
	}
	_, err = http.Post(url, "application/json", strings.NewReader(string(containerSpecJSON)))
	if err != nil {
		return netContainer, err
	}
	return netContainer, nil
}

func (dotNetBackend *dotNetBackend) Destroy(handle string) error {
	url := dotNetBackend.containerizerURL.String() + "/api/containers/" + handle

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	_, err = http.DefaultClient.Do(req)

	return err
}

func (dotNetBackend *dotNetBackend) Containers(api.Properties) ([]api.Container, error) {
	containers := []api.Container{
		container.NewContainer(dotNetBackend.containerizerURL, "containerhandle"),
		container.NewContainer(dotNetBackend.containerizerURL, "containerhandle"),
		container.NewContainer(dotNetBackend.containerizerURL, "containerhandle"),
		container.NewContainer(dotNetBackend.containerizerURL, "containerhandle"),
	}
	return containers, nil
}

func (dotNetBackend *dotNetBackend) Lookup(handle string) (api.Container, error) {
	netContainer := container.NewContainer(dotNetBackend.containerizerURL, handle)
	return netContainer, nil
}
