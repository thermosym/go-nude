# go-nude

Nudity detection with Go.

(Go porting from https://github.com/pa7/nude.js)
(Fork from https://github.com/koyachi/go-nude)

## Install
```bash
go get github.com/thermosym/go-nude
```

## Build
```bash
cd server # go into server subfolder
go build # it will generate an executable file 'server'
```

## Usage

```bash
./server [-p port_number]
```

## HTTP Access
```bash
curl --request POST \
  --url http://localhost:8080/check \
  --header 'content-type: multipart/form-data' \
  --form 'file=@/Path/To/Photo'
```

Response:
```json
{
  "is_nude": true,
  "success": true,
  "message": "ok"
}
```


## License
[MIT License](LICENSE.txt)