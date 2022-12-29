import {grpc} from "@improbable-eng/grpc-web";
import {GrpcWebImpl, TimeServiceClientImpl} from "../generated/time";

const rpc = new GrpcWebImpl('http://localhost:8080', {
    debug: false,
    metadata: new grpc.Metadata({}),
});

const client = new TimeServiceClientImpl(rpc);

console.log(await client.GetGreet({}));
console.log(await client.GetCurrentTime({}));

