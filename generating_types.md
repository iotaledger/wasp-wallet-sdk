The IOTA SDK offers a message bus which interacts via JSON messages.

The Go library is a PoC and mainly concentrates on the wallet/ledger functionality.

To bootstrap the library quicker, all typed JSON models were taken from the TypeScript binding.

The models were exported as a JSON schema, which can be used to generate Go structs.

Generated structs are not 100% true to the original and require refactoring once they are needed.

If new API functions are added to the library, the generated structs can be used as a reference - but they still need to be adjusted properly. (Use the nodejs/python iota-sdk bindings as a reference)


# Generating the Go types

## Create a JSON Schema from the TypeScript types

```bash
cd <root>/bindings/nodejs/
npm i
cd <root>/bindings/native/iota-sdk-go
```

### Create the schema

`npx ts-json-schema-generator --path '<iotasdk>/bindings/nodejs/lib/types/**/*.ts' -f "<iotasdk>/bindings/nodejs/tsconfig.json" -o types/schema.json`

Patch the `schema.json` file. Remove:

```
"nativeTokens": {
    "items": {
      "items": [
        {
          "type": "string"
        },
        {
          "$ref": "#/definitions/HexEncodedAmount"
        }
      ],
      "maxItems": 2,
      "minItems": 2,
      "type": "array"
    },
    "type": "array"
  },

```

**in line ~4019** at `SendNativeTokensParams`

otherwise the next step will fail.

The nativeTokens object needs to be added manually to the `types.go`

## Create Go structs

```bash
go install github.com/atombender/go-jsonschema/cmd/gojsonschema@latest
gojsonschema types/schema.json -p iota_sdk_go -o types/types.go_generated
```

It will warn about duplicates, but the types are complete.