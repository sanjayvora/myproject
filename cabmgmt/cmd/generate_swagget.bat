swagger generate spec -m  -o swagger.json -c cabmgmt/spec 
swagger flatten swagger.json --with-flatten=remove-unused -o swagger.json
