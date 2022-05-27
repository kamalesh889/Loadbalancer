# Loadbalancer

ENDPOINT USE:-

*)Use this Url (http://localhost:8080/urls/register) to register your urlstring

*)Then use (http://localhost:8080/proxy) in your chrome browser To make request and get all the response for url one my one (The urls should be alive )


Loadbalancer:-

Here the requests are distributed to healthy urls. This Process is performed in a round robin way to make the use of all alive servers .


Healthchk:-
Here all the urls health is being tested that they are able to handle request or not

I had some confusion regarding the health check conditions 
  
    *)If one url is dead , then it should be alive or not after handleing succesfull request
    *)I could not able thougth of way to recreate the the condition that you have given I.e
            ( A healthy endpoint is when the previous 3 responses are 200ok)

    *)So to counter that i have written a func which will update the status of urls in every 5secs,
    in that way the healthy url will  get the requets..
    It also solves the problem of first point


