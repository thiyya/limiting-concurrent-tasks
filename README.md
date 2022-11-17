# limiting-concurent-tasks
## Description
It is a golang tool with the ability to run in parallel that provides MD5 hashes of the responses of the given urls.

 - Default limit of the number of parallel request is 10.
 - With the -parallel flag, this limit can be set by the user.
 - Max limit of the number of parallel request is 10. 

## Usage
#### Execution

```bash
go build -o myhttp
./myhttp google.com
./myhttp -parallel 3 adjust.com google.com facebook.com yahoo.com yandex.com twitter.com reddit.com/r/funny reddit.com/r/notfunny baroquemusiclibrary.com
```

#### Output
```bash
https://baroquemusiclibrary.com - 9836aea9192bae4b59f7425c043d255b
https://adjust.com - e16184ca90fb37e5c058584f46e42eed
https://reddit.com/r/notfunny - 5ab2daa35aa8b87b2eb6d57aeec2b7f3
https://google.com - 1abf4f69854d4205aacb27fcde3270a5
https://reddit.com/r/funny - 2964d0a3bd21b75b1d376bb6c0c4d876
https://yandex.com - bbd2aa52d1f05940d7ef1a4e0b5dc6a5
https://yahoo.com - f8326ea8bc7a7bb82f4ee5c82a880514
https://facebook.com - 7c78352e1c69b93bb689b7524a9c8507
https://twitter.com - f0a2f77da745334fd4b2e819fa4bbf34
```