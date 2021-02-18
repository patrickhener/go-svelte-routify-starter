# go-svelte-routify-starter
This is a good starter to have an embedded multi-page routify svelte app within a go binary

# To test it

```bash
cd frontend
yarn install
yarn run build

cd ..
go build -o test main.go
./test
```

Now browse to `http://localhost:8080` or `http://localhost:8080/other` and see the magic happen.
