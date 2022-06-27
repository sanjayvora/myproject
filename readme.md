Cab Management Service

Service provides REST APIs to manage cabs.
Service operates on 3 different resources, Cab, City and Trip

Provided very generic and extensible solution with limited error handling. 
Proper error handlind and mapping to correct http error code is required.

Service is devided into multiple layers and each layer works on interface. 
Because of interface usage, each layer can be unit tested my mocking interface.
Not added any tests because of limited time.

DL layer is not implemented. Only interface is provided. 
Actual implementation would require interacting with database and doing queries on database.

Authorization and authentication support should be added on top of this.