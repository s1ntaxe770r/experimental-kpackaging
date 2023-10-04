package main

import (
	"context"

	"github.com/vmware-tanzu/kpack-cli/pkg/image"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/vmware-tanzu/kpack-cli/pkg/k8s"
	"golang.org/x/exp/slog"
)

var (
	kclientSet = k8s.ClientSet{}
)

func main() {

	clientSet := k8s.DefaultClientSetProvider{}

	cs, err := clientSet.GetClientSet("default")
	if err != nil {
		slog.Error("unable to create default client set")
	}

	ctx := context.Background()
	factory := image.Factory{
		Builder:        "my-builder",
		GitRepo:        "https://github.com/s1ntaxe770r/pong-server",
		ServiceAccount: "tutorial-service-account",
		GitRevision:    "38cd8469e289a0f8b3d220b228b0f7365af9fcf0",
	}

	img, err := factory.MakeImage("kp-test", cs.Namespace, "nullchannel/nc-test-random")

	if err != nil {
		slog.Error("unable to create image factory", err.Error())

	}

	img, err = cs.KpackClient.KpackV1alpha2().Images(cs.Namespace).Create(ctx, img, metav1.CreateOptions{})
	if err != nil {

		slog.Error("unable to create image ", err.Error())
	}
}
