package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `[
			{
				"id": "_pb_users_auth_",
				"created": "2022-12-25 06:44:56.320Z",
				"updated": "2022-12-25 06:53:05.058Z",
				"name": "users",
				"type": "auth",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "users_name",
						"name": "name",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "users_avatar",
						"name": "avatar",
						"type": "file",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"maxSize": 5242880,
							"mimeTypes": [
								"image/jpg",
								"image/jpeg",
								"image/png",
								"image/svg+xml",
								"image/gif"
							],
							"thumbs": null
						}
					}
				],
				"listRule": "id = @request.auth.id",
				"viewRule": "id = @request.auth.id",
				"createRule": "",
				"updateRule": "id = @request.auth.id",
				"deleteRule": "id = @request.auth.id",
				"options": {
					"allowEmailAuth": true,
					"allowOAuth2Auth": false,
					"allowUsernameAuth": false,
					"exceptEmailDomains": null,
					"manageRule": null,
					"minPasswordLength": 8,
					"onlyEmailDomains": null,
					"requireEmail": true
				}
			},
			{
				"id": "vw19gz7x2lskx0v",
				"created": "2022-12-25 06:49:38.449Z",
				"updated": "2022-12-25 18:44:15.562Z",
				"name": "device",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "wrygpjuc",
						"name": "owner",
						"type": "relation",
						"required": true,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": true
						}
					},
					{
						"system": false,
						"id": "eno0rrim",
						"name": "name",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"listRule": "@request.auth.id = owner.id\n|| (id = @collection.device_permissions.device.id && @request.auth.id = @collection.device_permissions.user.id && @collection.device_permissions.read = true)",
				"viewRule": "@request.auth.id = owner.id",
				"createRule": "@request.auth.id = owner.id",
				"updateRule": "@request.auth.id = owner.id",
				"deleteRule": "@request.auth.id = owner.id",
				"options": {}
			},
			{
				"id": "ts9179qz5gnf0jg",
				"created": "2022-12-25 07:01:15.792Z",
				"updated": "2022-12-25 07:05:56.782Z",
				"name": "locations",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "efohn1fb",
						"name": "device",
						"type": "relation",
						"required": true,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"collectionId": "vw19gz7x2lskx0v",
							"cascadeDelete": true
						}
					},
					{
						"system": false,
						"id": "tmx4jbwn",
						"name": "lat",
						"type": "number",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null
						}
					},
					{
						"system": false,
						"id": "ctddfyvt",
						"name": "lon",
						"type": "number",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null
						}
					},
					{
						"system": false,
						"id": "o1j3hkvt",
						"name": "acc",
						"type": "number",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null
						}
					},
					{
						"system": false,
						"id": "cyuvzvqk",
						"name": "speed",
						"type": "number",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null
						}
					},
					{
						"system": false,
						"id": "zdhezo2i",
						"name": "timestamp",
						"type": "date",
						"required": false,
						"unique": false,
						"options": {
							"min": "",
							"max": ""
						}
					}
				],
				"listRule": "device.owner.id = @request.auth.id",
				"viewRule": "device.owner.id = @request.auth.id",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "4v953trrzfanl4i",
				"created": "2022-12-25 07:08:29.498Z",
				"updated": "2022-12-25 07:09:01.597Z",
				"name": "device_tokens",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "e8exaevy",
						"name": "device",
						"type": "relation",
						"required": true,
						"unique": true,
						"options": {
							"maxSelect": 1,
							"collectionId": "vw19gz7x2lskx0v",
							"cascadeDelete": true
						}
					},
					{
						"system": false,
						"id": "uciloz1v",
						"name": "token",
						"type": "text",
						"required": true,
						"unique": true,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"listRule": null,
				"viewRule": null,
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "c7pd5fce2jxypic",
				"created": "2022-12-25 18:19:23.471Z",
				"updated": "2022-12-25 18:27:02.458Z",
				"name": "device_permissions",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "arowrylh",
						"name": "user",
						"type": "relation",
						"required": true,
						"unique": true,
						"options": {
							"maxSelect": 1,
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": true
						}
					},
					{
						"system": false,
						"id": "iyh4zjkb",
						"name": "device",
						"type": "relation",
						"required": true,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"collectionId": "vw19gz7x2lskx0v",
							"cascadeDelete": true
						}
					},
					{
						"system": false,
						"id": "wxk15g0n",
						"name": "read",
						"type": "bool",
						"required": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "2k5oprda",
						"name": "write",
						"type": "bool",
						"required": false,
						"unique": false,
						"options": {}
					}
				],
				"listRule": "device.owner.id = @request.auth.id",
				"viewRule": "device.owner.id = @request.auth.id",
				"createRule": "device.owner.id = @request.auth.id",
				"updateRule": "device.owner.id = @request.auth.id",
				"deleteRule": "device.owner.id = @request.auth.id",
				"options": {}
			}
		]`

		collections := []*models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collections); err != nil {
			return err
		}

		return daos.New(db).ImportCollections(collections, true, nil)
	}, func(db dbx.Builder) error {
		return nil
	})
}
