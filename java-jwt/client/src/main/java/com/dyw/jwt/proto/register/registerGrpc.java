package com.dyw.jwt.proto.register;

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
    comments = "Source: register.proto")
public final class registerGrpc {

  private registerGrpc() {}

  public static final String SERVICE_NAME = "register.register";

  // Static method descriptors that strictly reflect the proto.
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  @java.lang.Deprecated // Use {@link #getRegisterMethod()} instead. 
  public static final io.grpc.MethodDescriptor<com.dyw.jwt.proto.register.RegisterRequest,
      com.dyw.jwt.proto.register.RegisterResponse> METHOD_REGISTER = getRegisterMethod();

  private static volatile io.grpc.MethodDescriptor<com.dyw.jwt.proto.register.RegisterRequest,
      com.dyw.jwt.proto.register.RegisterResponse> getRegisterMethod;

  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static io.grpc.MethodDescriptor<com.dyw.jwt.proto.register.RegisterRequest,
      com.dyw.jwt.proto.register.RegisterResponse> getRegisterMethod() {
    io.grpc.MethodDescriptor<com.dyw.jwt.proto.register.RegisterRequest, com.dyw.jwt.proto.register.RegisterResponse> getRegisterMethod;
    if ((getRegisterMethod = registerGrpc.getRegisterMethod) == null) {
      synchronized (registerGrpc.class) {
        if ((getRegisterMethod = registerGrpc.getRegisterMethod) == null) {
          registerGrpc.getRegisterMethod = getRegisterMethod = 
              io.grpc.MethodDescriptor.<com.dyw.jwt.proto.register.RegisterRequest, com.dyw.jwt.proto.register.RegisterResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "register.register", "Register"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dyw.jwt.proto.register.RegisterRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dyw.jwt.proto.register.RegisterResponse.getDefaultInstance()))
                  .setSchemaDescriptor(new registerMethodDescriptorSupplier("Register"))
                  .build();
          }
        }
     }
     return getRegisterMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static registerStub newStub(io.grpc.Channel channel) {
    return new registerStub(channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static registerBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    return new registerBlockingStub(channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static registerFutureStub newFutureStub(
      io.grpc.Channel channel) {
    return new registerFutureStub(channel);
  }

  /**
   */
  public static abstract class registerImplBase implements io.grpc.BindableService {

    /**
     */
    public void register(com.dyw.jwt.proto.register.RegisterRequest request,
        io.grpc.stub.StreamObserver<com.dyw.jwt.proto.register.RegisterResponse> responseObserver) {
      asyncUnimplementedUnaryCall(getRegisterMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getRegisterMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                com.dyw.jwt.proto.register.RegisterRequest,
                com.dyw.jwt.proto.register.RegisterResponse>(
                  this, METHODID_REGISTER)))
          .build();
    }
  }

  /**
   */
  public static final class registerStub extends io.grpc.stub.AbstractStub<registerStub> {
    private registerStub(io.grpc.Channel channel) {
      super(channel);
    }

    private registerStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected registerStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new registerStub(channel, callOptions);
    }

    /**
     */
    public void register(com.dyw.jwt.proto.register.RegisterRequest request,
        io.grpc.stub.StreamObserver<com.dyw.jwt.proto.register.RegisterResponse> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getRegisterMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class registerBlockingStub extends io.grpc.stub.AbstractStub<registerBlockingStub> {
    private registerBlockingStub(io.grpc.Channel channel) {
      super(channel);
    }

    private registerBlockingStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected registerBlockingStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new registerBlockingStub(channel, callOptions);
    }

    /**
     */
    public com.dyw.jwt.proto.register.RegisterResponse register(com.dyw.jwt.proto.register.RegisterRequest request) {
      return blockingUnaryCall(
          getChannel(), getRegisterMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class registerFutureStub extends io.grpc.stub.AbstractStub<registerFutureStub> {
    private registerFutureStub(io.grpc.Channel channel) {
      super(channel);
    }

    private registerFutureStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected registerFutureStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new registerFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.dyw.jwt.proto.register.RegisterResponse> register(
        com.dyw.jwt.proto.register.RegisterRequest request) {
      return futureUnaryCall(
          getChannel().newCall(getRegisterMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_REGISTER = 0;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final registerImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(registerImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_REGISTER:
          serviceImpl.register((com.dyw.jwt.proto.register.RegisterRequest) request,
              (io.grpc.stub.StreamObserver<com.dyw.jwt.proto.register.RegisterResponse>) responseObserver);
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

  private static abstract class registerBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    registerBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.dyw.jwt.proto.register.Register.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("register");
    }
  }

  private static final class registerFileDescriptorSupplier
      extends registerBaseDescriptorSupplier {
    registerFileDescriptorSupplier() {}
  }

  private static final class registerMethodDescriptorSupplier
      extends registerBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    registerMethodDescriptorSupplier(String methodName) {
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
      synchronized (registerGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new registerFileDescriptorSupplier())
              .addMethod(getRegisterMethod())
              .build();
        }
      }
    }
    return result;
  }
}
