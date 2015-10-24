# flapjack-broker
Flapjack HTTP Broker, this currently supports notifications from New Relic Alerts (Beta), Flapjack and Cloudwatch.

Each endpoint is controlled with its own martini router, should you wish to add a new broker modify libexec/httpbroker.go and add a martini router then create the package and functions inside of src/flapjackfeeder/. A good starting point for a new hok would be the most simple feeder which is src/flapjackfeeder/flapjack.go

If the data type is different to those surrently supported make sure that you add the JSON struct to src/flapjackconfig/config.go

To build the broker simply run ./build.sh

This has been built and tested on Ubuntu and OS X

# Brokers 
/state - CloudWatch

/flapjack - Flapjack

/newrelic - New Relic

# Example
````
curl -w 'response: %{http_code} \n' -X POST \
  -H "Content-type: application/json" -d \
  '{
    "entity": "foo-app-01",
    "check": "PING",
    "type": "service",
    "tags": "apps",
    "state": "OK",
    "summary": "3 ms round trip",
    "ttl": 30
  }' http://localhost:3090/flapjack
````

