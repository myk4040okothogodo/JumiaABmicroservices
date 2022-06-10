package server

import (
    service_bv1 "github.com/myk4040okothogodo/JumiaABmicroservices/gen/go/service_B"

)


var (
    testOrder1 = &service_bv1.Order{
        Email: "mikeokoth@gmail.com",
        Weight: "288.88",
        Phonenumber: "254717535877",
        Countrycode: "254",

    }

    testOrder2 = &service_bv1.Order{
        Email: "kevinotis@gmail.com",
        Weight: "43.09",
        Phonenumber: "343938473828",
        Countrycode: "343",
    }

)
