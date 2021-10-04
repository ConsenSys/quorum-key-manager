package placeholder

import (
	"context"
	"regexp"

	"github.com/consensys/quorum-key-manager/pkg/errors"
	aliasent "github.com/consensys/quorum-key-manager/src/aliases/entities"
)

type Parser struct {
	regex *regexp.Regexp
}

func New() (*Parser, error) {
	const aliasParseFormat = `{{(?m)(?P<registry>[a-zA-Z0-9-_+]+):(?P<alias>[a-zA-Z0-9-_+]+)}}$`
	regex, err := regexp.Compile(aliasParseFormat)
	if err != nil {
		return nil, errors.ConfigError("bad regexp format '%v': %v", aliasParseFormat, err)
	}
	return &Parser{
		regex: regex,
	}, nil
}

func (p *Parser) ParseAlias(alias string) (regName, aliasKey string, isAlias bool, err error) {
	submatches := p.regex.FindStringSubmatch(alias)
	if len(submatches) < 3 {
		return "", "", false, nil
	}

	regName = submatches[1]
	aliasKey = submatches[2]

	return regName, aliasKey, true, nil
}

func (p *Parser) ReplaceAliases(ctx context.Context, aliasBackend aliasent.AliasBackend, addrs []string) ([]string, error) {
	var values []string
	for _, addr := range addrs {
		regName, aliasKey, isAlias, err := p.ParseAlias(addr)
		if err != nil {
			return nil, err
		}

		// it is not an alias
		if !isAlias {
			values = append(values, addr)
			continue
		}

		alias, err := aliasBackend.GetAlias(ctx, regName, aliasKey)
		if err != nil {
			return nil, err
		}

		values = append(values, alias.Value...)
	}
	return values, nil
}
