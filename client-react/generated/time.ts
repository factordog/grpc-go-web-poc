/* eslint-disable */
import { grpc } from "@improbable-eng/grpc-web";
import { BrowserHeaders } from "browser-headers";
import * as _m0 from "protobufjs/minimal";
import { Observable } from "rxjs";
import { share } from "rxjs/operators";

export const protobufPackage = "smpl.time.api.v1";

export interface GetGreetRequest {
}

export interface GetCurrentTimeRequest {
}

export interface GetCurrentTimeResponse {
  currentTime: string;
}

export interface GetGreetResponse {
  greet: string;
}

function createBaseGetGreetRequest(): GetGreetRequest {
  return {};
}

export const GetGreetRequest = {
  encode(_: GetGreetRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetGreetRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetGreetRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): GetGreetRequest {
    return {};
  },

  toJSON(_: GetGreetRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetGreetRequest>, I>>(_: I): GetGreetRequest {
    const message = createBaseGetGreetRequest();
    return message;
  },
};

function createBaseGetCurrentTimeRequest(): GetCurrentTimeRequest {
  return {};
}

export const GetCurrentTimeRequest = {
  encode(_: GetCurrentTimeRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetCurrentTimeRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetCurrentTimeRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): GetCurrentTimeRequest {
    return {};
  },

  toJSON(_: GetCurrentTimeRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetCurrentTimeRequest>, I>>(_: I): GetCurrentTimeRequest {
    const message = createBaseGetCurrentTimeRequest();
    return message;
  },
};

function createBaseGetCurrentTimeResponse(): GetCurrentTimeResponse {
  return { currentTime: "" };
}

export const GetCurrentTimeResponse = {
  encode(message: GetCurrentTimeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.currentTime !== "") {
      writer.uint32(10).string(message.currentTime);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetCurrentTimeResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetCurrentTimeResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.currentTime = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetCurrentTimeResponse {
    return { currentTime: isSet(object.currentTime) ? String(object.currentTime) : "" };
  },

  toJSON(message: GetCurrentTimeResponse): unknown {
    const obj: any = {};
    message.currentTime !== undefined && (obj.currentTime = message.currentTime);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetCurrentTimeResponse>, I>>(object: I): GetCurrentTimeResponse {
    const message = createBaseGetCurrentTimeResponse();
    message.currentTime = object.currentTime ?? "";
    return message;
  },
};

function createBaseGetGreetResponse(): GetGreetResponse {
  return { greet: "" };
}

export const GetGreetResponse = {
  encode(message: GetGreetResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.greet !== "") {
      writer.uint32(10).string(message.greet);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetGreetResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetGreetResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.greet = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetGreetResponse {
    return { greet: isSet(object.greet) ? String(object.greet) : "" };
  },

  toJSON(message: GetGreetResponse): unknown {
    const obj: any = {};
    message.greet !== undefined && (obj.greet = message.greet);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetGreetResponse>, I>>(object: I): GetGreetResponse {
    const message = createBaseGetGreetResponse();
    message.greet = object.greet ?? "";
    return message;
  },
};

export interface TimeService {
  GetCurrentTime(
    request: DeepPartial<GetCurrentTimeRequest>,
    metadata?: grpc.Metadata,
  ): Promise<GetCurrentTimeResponse>;
  GetGreet(request: DeepPartial<GetGreetRequest>, metadata?: grpc.Metadata): Promise<GetGreetResponse>;
  StreamTimeUpdates(
    request: DeepPartial<GetCurrentTimeRequest>,
    metadata?: grpc.Metadata,
  ): Observable<GetCurrentTimeResponse>;
}

export class TimeServiceClientImpl implements TimeService {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.GetCurrentTime = this.GetCurrentTime.bind(this);
    this.GetGreet = this.GetGreet.bind(this);
    this.StreamTimeUpdates = this.StreamTimeUpdates.bind(this);
  }

  GetCurrentTime(
    request: DeepPartial<GetCurrentTimeRequest>,
    metadata?: grpc.Metadata,
  ): Promise<GetCurrentTimeResponse> {
    return this.rpc.unary(TimeServiceGetCurrentTimeDesc, GetCurrentTimeRequest.fromPartial(request), metadata);
  }

  GetGreet(request: DeepPartial<GetGreetRequest>, metadata?: grpc.Metadata): Promise<GetGreetResponse> {
    return this.rpc.unary(TimeServiceGetGreetDesc, GetGreetRequest.fromPartial(request), metadata);
  }

  StreamTimeUpdates(
    request: DeepPartial<GetCurrentTimeRequest>,
    metadata?: grpc.Metadata,
  ): Observable<GetCurrentTimeResponse> {
    return this.rpc.invoke(TimeServiceStreamTimeUpdatesDesc, GetCurrentTimeRequest.fromPartial(request), metadata);
  }
}

export const TimeServiceDesc = { serviceName: "smpl.time.api.v1.TimeService" };

export const TimeServiceGetCurrentTimeDesc: UnaryMethodDefinitionish = {
  methodName: "GetCurrentTime",
  service: TimeServiceDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return GetCurrentTimeRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = GetCurrentTimeResponse.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const TimeServiceGetGreetDesc: UnaryMethodDefinitionish = {
  methodName: "GetGreet",
  service: TimeServiceDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return GetGreetRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = GetGreetResponse.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const TimeServiceStreamTimeUpdatesDesc: UnaryMethodDefinitionish = {
  methodName: "StreamTimeUpdates",
  service: TimeServiceDesc,
  requestStream: false,
  responseStream: true,
  requestType: {
    serializeBinary() {
      return GetCurrentTimeRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = GetCurrentTimeResponse.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

interface UnaryMethodDefinitionishR extends grpc.UnaryMethodDefinition<any, any> {
  requestStream: any;
  responseStream: any;
}

type UnaryMethodDefinitionish = UnaryMethodDefinitionishR;

interface Rpc {
  unary<T extends UnaryMethodDefinitionish>(
    methodDesc: T,
    request: any,
    metadata: grpc.Metadata | undefined,
  ): Promise<any>;
  invoke<T extends UnaryMethodDefinitionish>(
    methodDesc: T,
    request: any,
    metadata: grpc.Metadata | undefined,
  ): Observable<any>;
}

export class GrpcWebImpl {
  private host: string;
  private options: {
    transport?: grpc.TransportFactory;
    streamingTransport?: grpc.TransportFactory;
    debug?: boolean;
    metadata?: grpc.Metadata;
    upStreamRetryCodes?: number[];
  };

  constructor(
    host: string,
    options: {
      transport?: grpc.TransportFactory;
      streamingTransport?: grpc.TransportFactory;
      debug?: boolean;
      metadata?: grpc.Metadata;
      upStreamRetryCodes?: number[];
    },
  ) {
    this.host = host;
    this.options = options;
  }

  unary<T extends UnaryMethodDefinitionish>(
    methodDesc: T,
    _request: any,
    metadata: grpc.Metadata | undefined,
  ): Promise<any> {
    const request = { ..._request, ...methodDesc.requestType };
    const maybeCombinedMetadata = metadata && this.options.metadata
      ? new BrowserHeaders({ ...this.options?.metadata.headersMap, ...metadata?.headersMap })
      : metadata || this.options.metadata;
    return new Promise((resolve, reject) => {
      grpc.unary(methodDesc, {
        request,
        host: this.host,
        metadata: maybeCombinedMetadata,
        transport: this.options.transport,
        debug: this.options.debug,
        onEnd: function (response) {
          if (response.status === grpc.Code.OK) {
            resolve(response.message!.toObject());
          } else {
            const err = new GrpcWebError(response.statusMessage, response.status, response.trailers);
            reject(err);
          }
        },
      });
    });
  }

  invoke<T extends UnaryMethodDefinitionish>(
    methodDesc: T,
    _request: any,
    metadata: grpc.Metadata | undefined,
  ): Observable<any> {
    const upStreamCodes = this.options.upStreamRetryCodes || [];
    const DEFAULT_TIMEOUT_TIME: number = 3_000;
    const request = { ..._request, ...methodDesc.requestType };
    const maybeCombinedMetadata = metadata && this.options.metadata
      ? new BrowserHeaders({ ...this.options?.metadata.headersMap, ...metadata?.headersMap })
      : metadata || this.options.metadata;
    return new Observable((observer) => {
      const upStream = (() => {
        const client = grpc.invoke(methodDesc, {
          host: this.host,
          request,
          transport: this.options.streamingTransport || this.options.transport,
          metadata: maybeCombinedMetadata,
          debug: this.options.debug,
          onMessage: (next) => observer.next(next),
          onEnd: (code: grpc.Code, message: string, trailers: grpc.Metadata) => {
            if (code === 0) {
              observer.complete();
            } else if (upStreamCodes.includes(code)) {
              setTimeout(upStream, DEFAULT_TIMEOUT_TIME);
            } else {
              const err = new Error(message) as any;
              err.code = code;
              err.metadata = trailers;
              observer.error(err);
            }
          },
        });
        observer.add(() => client.close());
      });
      upStream();
    }).pipe(share());
  }
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var tsProtoGlobalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

export class GrpcWebError extends tsProtoGlobalThis.Error {
  constructor(message: string, public code: grpc.Code, public metadata: grpc.Metadata) {
    super(message);
  }
}
