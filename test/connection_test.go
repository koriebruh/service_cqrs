package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"koriebruh/cqrs/config"
	"koriebruh/cqrs/pkg"
	"testing"
)

var cnf = config.GetConfig()

func TestElasticClient(t *testing.T) {
	client := pkg.ElasticClient(cnf)

	info, err := client.Info()
	assert.Nil(t, err)
	defer info.Body.Close()

	fmt.Println(info)
	assert.NotNil(t, info)
}

func TestMysqlClient(t *testing.T) {
	client := pkg.MysqlClient(cnf)

	db, err := client.DB()
	assert.Nil(t, err)

	err = db.Ping()
	assert.Nil(t, err)
}
