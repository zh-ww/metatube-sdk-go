package xslistplus

import (
	"github.com/javtube/javtube-sdk-go/model"
	"github.com/javtube/javtube-sdk-go/provider"
	"github.com/javtube/javtube-sdk-go/provider/gfriends"
	"github.com/javtube/javtube-sdk-go/provider/xslist"
)

type xsListPlus struct {
	*xslist.XsList
	gf *gfriends.GFriends
}

func newXsListPlus() *xsListPlus {
	return &xsListPlus{
		XsList: xslist.New(),
		gf:     gfriends.New(),
	}
}

func (xsl *xsListPlus) GetActorInfoByID(id string) (info *model.ActorInfo, err error) {
	defer func() {
		if err == nil && info.Valid() {
			if gInfo, gErr := xsl.gf.GetActorInfoByID(info.Name); gErr == nil && gInfo.Valid() {
				info.Images = append(gInfo.Images, info.Images...)
			}
		}
	}()
	return xsl.XsList.GetActorInfoByID(id)
}

func (xsl *xsListPlus) SearchActor(keyword string) (results []*model.ActorSearchResult, err error) {
	defer func() {
		if err == nil {
			for _, result := range results {
				if result.Valid() {
					if gInfo, gErr := xsl.gf.GetActorInfoByID(result.Name); gErr == nil && gInfo.Valid() {
						result.Images = append(gInfo.Images, result.Images...)
					}
				}
			}
		}
	}()
	return xsl.XsList.SearchActor(keyword)
}

func init() {
	// We use a little hack here to override the original
	// xslist provider factory.
	provider.RegisterActorFactory("xslist", newXsListPlus)
}
