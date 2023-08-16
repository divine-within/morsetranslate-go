# MorseTranslate - Go

`MorseTranslate` is an API with speed and efficiency in mind;
written in Go using std library and a third party `Chi` router.

## Routes

- `/api/v1/dictionary` - returns a JSON response with the dictionary
- `/api/v1/translate/text` - takes a `text` query parameter and returns a JSON response
- `/api/v1/translate/morsecode` - takes a `morsecode` query parameter and returns a JSON response

## Efficiency

API utilizes Go's standard library and a router from Chi, which makes it very fast.
Hash map data structure is used for the models, because of the time complexity O(1) of searches.

## Authentication

So far only a sample JWT token is being generated for debugging purposes.

## Dependencies

- [Chi](https://github.com/go-chi/chi)
- [JWTAuth](https://github.com/go-chi/jwtauth)
- [CORS](https://github.com/go-chi/cors)

## License

Licensed under [MIT License](https://github.com/divine-within/morsetranslate-go/blob/main/LICENSE)
