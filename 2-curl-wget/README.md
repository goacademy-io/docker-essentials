## Get request
```bash
 docker run -d curlimages/curl -fsSL https://google.com
```

## Post request
```bash
 docker run --rm -it -v "$PWD:/work" curlimages/curl -d@/work/test.json https://httpbin.org/post
```

## Wget

```bash
 docker run -it cirrusci/wget wget http://ftp.gnu.org/gnu/wget/wget2-2.0.0.tar.gz
```
