# avro-compat-test

Test for gogen-avro backward/forward compliance.

This is to test gogen-avro ability to support backward or forward compatible change to an avro schema.

You first need to install gogen-avro:

```
go get gopkg.in/alanctgardner/gogen-avro.v4
```

You can then run the test with:

```
go generate
go test .
```

Test will fail with error:

```
--- FAIL: TestV1ToV2 (0.00s)
        validator_test.go:43: Cannot deserialize object1 with schema2: EOF
```

It shows that if last values are missing, default values are not applied.