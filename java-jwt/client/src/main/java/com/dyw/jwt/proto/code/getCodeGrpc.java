package com.dyw.jwt.proto.code;

import static io.grpc.MethodDescriptor.generateFullMethodName;
import static io.grpc.stub.ClientCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ClientCalls.asyncClientStreamingCall;
import static io.grpc.stub.ClientCalls.asyncServerStreamingCall;
import static io.grpc.stub.ClientCalls.asyncUnaryCall;
import static io.grpc.stub.ClientCalls.blockingServerStreamingCall;
import static io.grpc.stub.ClientCalls.blockingUnaryCall;
import static io.grpc.stub.ClientCalls.futureUnaryCall;
import static io.grpc.stub.ServerCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ServerCalls.asyncClientStreamingCall;
import static io.grpc.stub.ServerCalls.asyncServerStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnaryCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.8.0)",
    comments = "Source: code.proto")
public final class getCodeGrpc {

  private getCodeGrpc() {}

  public static final String SERVICE_NAME = "getcode.getCode";

  // Static method descriptors that strictly reflect the proto.
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  @java.lang.Deprecated // Use {@link #getGetCodeMethod()} instead. 
  public static final io.grpc.MethodDescriptor<com.dyw.jwt.proto.code.CodeRequest,
      com.dyw.jwt.proto.code.CodeResponse> METHOD_GET_CODE = getGetCodeMethod();

  private static volatile io.grpc.MethodDescriptor<com.dyw.jwt.proto.code.CodeRequest,
      com.dyw.jwt.proto.code.CodeResponse> getGetCodeMethod;

  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static io.grpc.MethodDescriptor<com.dyw.jwt.proto.code.CodeRequest,
      com.dyw.jwt.proto.code.CodeResponse> getGetCodeMethod() {
    io.grpc.MethodDescriptor<com.dyw.jwt.proto.code.CodeRequest, com.dyw.jwt.proto.code.CodeResponse> getGetCodeMethod;
    if ((getGetCodeMethod = getCodeGrpc.getGetCodeMethod) == null) {
      synchronized (getCodeGrpc.class) {
        if ((getGetCodeMethod = getCodeGrpc.getGetCodeMethod) == null) {
          getCodeGrpc.getGetCodeMethod = getGetCodeMethod = 
              io.grpc.MethodDescriptor.<com.dyw.jwt.proto.code.CodeRequest, com.dyw.jwt.proto.code.CodeResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "getcode.getCode", "getCode"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dyw.jwt.proto.code.CodeRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dyw.jwt.proto.code.CodeResponse.getDefaultInstance()))
                  .setSchemaDescriptor(new getCodeMethodDescriptorSupplier("getCode"))
                  .build();
          }
        }
     }
     return getGetCodeMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static getCodeStub newStub(io.grpc.Channel channel) {
    return new getCodeStub(channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static getCodeBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    return new getCodeBlockingStub(channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static getCodeFutureStub newFutureStub(
      io.grpc.Channel channel) {
    return new getCodeFutureStub(channel);
  }

  /**
   */
  public static abstract class getCodeImplBase implements io.grpc.BindableService {

    /**
     */
    public void getCode(com.dyw.jwt.proto.code.CodeRequest request,
        io.grpc.stub.StreamObserver<com.dyw.jwt.proto.code.CodeResponse> responseObserver) {
      asyncUnimplementedUnaryCall(getGetCodeMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getGetCodeMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                com.dyw.jwt.proto.code.CodeRequest,
                com.dyw.jwt.proto.code.CodeResponse>(
                  this, METHODID_GET_CODE)))
          .build();
    }
  }

  /**
   */
  public static final class getCodeStub extends io.grpc.stub.AbstractStub<getCodeStub> {
    private getCodeStub(io.grpc.Channel channel) {
      super(channel);
    }

    private getCodeStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected getCodeStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new getCodeStub(channel, callOptions);
    }

    /**
     */
    public void getCode(com.dyw.jwt.proto.code.CodeRequest request,
        io.grpc.stub.StreamObserver<com.dyw.jwt.proto.code.CodeResponse> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getGetCodeMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class getCodeBlockingStub extends io.grpc.stub.AbstractStub<getCodeBlockingStub> {
    private getCodeBlockingStub(io.grpc.Channel channel) {
      super(channel);
    }

    private getCodeBlockingStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected getCodeBlockingStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new getCodeBlockingStub(channel, callOptions);
    }

    /**
     */
    public com.dyw.jwt.proto.code.CodeResponse getCode(com.dyw.jwt.proto.code.CodeRequest request) {
      return blockingUnaryCall(
          getChannel(), getGetCodeMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class getCodeFutureStub extends io.grpc.stub.AbstractStub<getCodeFutureStub> {
    private getCodeFutureStub(io.grpc.Channel channel) {
      super(channel);
    }

    private getCodeFutureStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected getCodeFutureStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new getCodeFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.dyw.jwt.proto.code.CodeResponse> getCode(
        com.dyw.jwt.proto.code.CodeRequest request) {
      return futureUnaryCall(
          getChannel().newCall(getGetCodeMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_GET_CODE = 0;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final getCodeImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(getCodeImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_GET_CODE:
          serviceImpl.getCode((com.dyw.jwt.proto.code.CodeRequest) request,
              (io.grpc.stub.StreamObserver<com.dyw.jwt.proto.code.CodeResponse>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  private static abstract class getCodeBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    getCodeBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.dyw.jwt.proto.code.Code.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("getCode");
    }
  }

  private static final class getCodeFileDescriptorSupplier
      extends getCodeBaseDescriptorSupplier {
    getCodeFileDescriptorSupplier() {}
  }

  private static final class getCodeMethodDescriptorSupplier
      extends getCodeBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    getCodeMethodDescriptorSupplier(String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (getCodeGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new getCodeFileDescriptorSupplier())
              .addMethod(getGetCodeMethod())
              .build();
        }
      }
    }
    return result;
  }
}
