/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    channel_cost_batch
	@Date    2022/5/12 10:08
	@Desc
*/

package quote

import (
	"context"
	"time"

	ent "github.com/heshaofeng1991/entgo/ent/gen"
	"github.com/heshaofeng1991/entgo/ent/gen/channelcostbatche"
	domainEntity "github.com/heshaofeng1991/logistics-backend/domain/entity"
	"github.com/pkg/errors"
)

func (c ShippingOptionRepository) GetChannelCostBatch(ctx context.Context, ids []int64,
	date time.Time,
) ([]*domainEntity.ChannelCostBatch, error) {
	var (
		err              error
		channelCostBatch []*ent.ChannelCostBatche
	)

	channelCostBatch, err = c.entClient.ChannelCostBatche.Query().Where(
		channelcostbatche.And(
			channelcostbatche.ChannelIDIn(ids...),
			channelcostbatche.Status(true),
			channelcostbatche.EffectiveDateLT(date),
		),
	).Order(ent.Desc(channelcostbatche.FieldID)).All(ctx)

	return domainEntity.CovertDBToChannelCostBatch(channelCostBatch), errors.Wrap(err, "")
}
