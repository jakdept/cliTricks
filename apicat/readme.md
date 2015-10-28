apicat
======

This utility is used to beat a JSON api endpoint with input from your json input - incrementing the location at the `-requestedPage` until `-currentPage` matches `-totalPage`.

Usage
-----

```bash
$ apicat --help
Usage of apicat:
  -currentPage string
    	location in the response of the page returned
  -pageIncrement int
    	number to increase location request by (default 1)
  -password string
    	username to use for authentication
  -requestedPage string
    	location in the request of the page
  -totalPage string
    	location in the response of the total pages
  -url string
    	url location to direct POSt
  -username string
    	username to use for authentication
```

Example
-------

Unfortunatly, as this time, I have no example API endpoint that I can send this at - nothing that uses JSON as the input and has pagination. If you have a suggestion, please let me know - I would love to find one.