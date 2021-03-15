package build

import (
	"context"
	"github.com/containers/buildah"
	"github.com/containers/storage"
	"github.com/containers/storage/pkg/unshare"
	"log"
)

var (
	builderOpts = buildah.BuilderOptions{
		FromImage:        "golang:latest",
	}

)

func buildContainer()  {
	buildStoreOptions, err := storage.DefaultStoreOptions(unshare.IsRootless(), unshare.GetRootlessUID())
	if err != nil {
		log.Fatal("error on building", err)
		return
	}
	buildStore, err := storage.GetStore(buildStoreOptions)
	if err != nil {
		log.Fatal("error on building", err)
		return
	}
	builder, err:= buildah.NewBuilder(context.TODO(), buildStore, builderOpts)
	if err != nil {
		log.Fatal("error on building", err)
		return
	}
	builder.SetCmd([]string{"./password-manager $ARGS"})
}