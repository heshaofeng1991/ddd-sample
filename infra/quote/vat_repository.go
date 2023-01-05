/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    vat
	@Date    2022/5/12 10:09
	@Desc
*/

package quote

import (
	"context"

	vat "github.com/heshaofeng1991/entgo/ent/gen/valueaddedtax"
	domainEntity "github.com/heshaofeng1991/logistics-backend/domain/entity"
	"github.com/pkg/errors"
)

func (c ShippingOptionRepository) GetVatInfoByCountryCode(
	ctx context.Context,
	countryCode string,
) (*domainEntity.Vat, error) {
	result, err := c.entClient.ValueAddedTax.Query().Where(vat.CountryCodeEQ(countryCode)).First(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "")
	}

	return domainEntity.NewVat(
		result.ID,
		result.CountryCode,
		result.StandardRate,
		result.WithoutIossRate,
		result.ExemptionInUsd,
	), nil
}
