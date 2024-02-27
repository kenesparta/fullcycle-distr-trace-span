# fullcycle-distr-trace-span
Distributed Trace &amp; Span

# ⚠️ Create your API key from weather API

- You need to **create your own API key** by accessing this URL [🌦️weather API](https://www.weatherapi.com/signup.aspx)
- After that, you can use your api_key, and you can put this into the json config below.
- You can use this API key `dca102b972c84ce989261931242701` if you do not have one, but we do not guarantee if that key
  is valid at the moment you run the project, so we highly recommend you create your own API key.

- Testing:
```shell
curl -i -X POST http://127.0.0.1:8080/temperature -d '{"cep":"6407552"}'
curl -i -X POST http://127.0.0.1:8080/temperature -d '{"cep":"64075525"}'
curl -i -X POST http://127.0.0.1:8080/temperature -d '{"cep":"00001033"}'
```
