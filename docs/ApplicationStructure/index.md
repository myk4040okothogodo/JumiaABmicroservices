Structure

The application will have two services

    - Service_A - for parsing our csv data and preping it for transmission
    - Service_B - for persisting our data  that is made up or orders and implementing logic for our controlles


we Then connect one of the services , Service_B that persisist our orders in an arrangoDB to GraphQL,
Both service_A and service_b are built on gRPC protocol, however service_A doesnt mantain a database instance only service_B does.

   - The GrapgQL gateway Service

The GraphQL layer will serve as the client side of our services , sepecifically service_B, since our services should be loosely coupled, we make this
gateway service receive all requests associated with Orders, and make one service to persist this data.
