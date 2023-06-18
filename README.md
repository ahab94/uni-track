# UniTrack

## API

# UniTrack API Documentation

Just execute the following command to get started with the local process
  ```bash
  make demo
  ```

## Endpoints

### Pool Data Endpoint

- **Purpose**: Retrieve information about a particular pool.
- **Path**: `/v1/api/pool/{pool-id}?block={block-number}`
- **Method**: `GET`
- **Possible Response Statuses**: `200`, `500`
- **CURL Example**:
  ```bash
  curl --location 'http://localhost:8080/v1/api/pool/0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640?block=17508123'
  ```
- **Expected Response**:
  ```json
  HTTP/1.1 200 OK
  Date: Mon, 30 Mar 2020 08:09:59 GMT
  Content-Length: 39
  Content-Type: application/json
  { "blockNumber": "17508270", "tick": "27909511162785237640782267391806507290588793534566025261598951948802966127860740510212802240046443963845555310749548015061718179056933732858792329275980206750485191840279446457359273270411716482975464989999922936560592179435546883713809330966118575833262340694441858082749724271126957839724737569927501186274251157574649757101796043748130848608082903397210529192699452081643431126075948038683562866810347421589873022711184374000395653223266804133765442579580489395603247201068323260386034597039715479712951921643365726682932500419942843154433", "token0Balance": "80501073118474", "token1Balance": "78903989236759639708054" }  ```

### Pool Historic Data Endpoint

- **Purpose**: Retrieve historical data for a particular pool.
- **Path**: `/v1/api/pool/{pool-id}/historic`
- **Method**: `GET`
- **Possible Response Statuses**: `200`, `500`
- **CURL Example**:
  ```bash
  curl --location 'http://localhost:8080/v1/api/pool/0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640/historic'
  ```
- **Expected Response**:
  ```json
  HTTP/1.1 200 OK
  Date: Mon, 30 Mar 2020 08:09:59 GMT
  Content-Length: 39
  Content-Type: application/json
  [ { "blockNumber": "17508270", "token0Balance": "80501073118474", "token0Delta": 0, "token1Balance": "78903989236759639708054", "token1Delta": 0 }, { "blockNumber": "17508268", "token0Balance": "80511324979125", "token0Delta": 10251860651, "token1Balance": "78898109385803929766726", "token1Delta": 0 }, { "blockNumber": "17508267", "token0Balance": "80511203979125", "token0Delta": -121000000, "token1Balance": "78898178713691634200099", "token1Delta": 0 }, { "blockNumber": "17508259", "token0Balance": "80511291158346", "token0Delta": 87179221, "token1Balance": "78898128713691634200099", "token1Delta": 0 } ] 
  ```

## Having Issues?

Please feel free to reach out at 'abdshah94@gmail.com'.
