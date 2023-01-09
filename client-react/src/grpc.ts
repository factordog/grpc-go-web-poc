import {grpc} from "@improbable-eng/grpc-web";
import {GrpcWebImpl, TimeServiceClientImpl} from "../generated/time";
import {Observable, of} from "rxjs";

const myTransport = grpc.CrossBrowserHttpTransport({ withCredentials: false });
const rpc = new GrpcWebImpl('http://localhost:8080', {
    debug: false,
    metadata: new grpc.Metadata({}),
    transport: myTransport
});

const client = new TimeServiceClientImpl(rpc);

// NOTE: uni directional data. Calls once and ends requests based on server response.
console.log(await client.GetGreet({}));
console.log(await client.GetCurrentTime({}));

// NOTE: client calls server and gets data stream
const stream = client.StreamTimeUpdates({});
stream.subscribe((data: any) => {
    console.log(data);
});

// const emitter : Observable<{message: string}> = of({message: "Hello"});
// const convoStream = client.BackAndForth(emitter);
// convoStream.subscribe((data: any) => {
//     console.log(data);
// });
