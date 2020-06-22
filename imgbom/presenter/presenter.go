package presenter

import (
	"io"

	"github.com/anchore/imgbom/imgbom/pkg"
	"github.com/anchore/imgbom/imgbom/presenter/json"
	"github.com/anchore/imgbom/imgbom/presenter/text"
	"github.com/anchore/stereoscope/pkg/image"
)

type Presenter interface {
	Present(io.Writer, *image.Image, *pkg.Catalog) error
}

func GetPresenter(option Option) Presenter {
	switch option {
	case JSONPresenter:
		return json.NewPresenter()
	case TextPresenter:
		return text.NewPresenter()
	default:
		return nil
	}
}
