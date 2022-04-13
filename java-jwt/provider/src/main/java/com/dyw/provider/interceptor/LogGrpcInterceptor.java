package com.dyw.provider.interceptor;

import com.dyw.provider.config.GlobalInterceptorConfiguration;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import io.grpc.Metadata;
import io.grpc.ServerCall;
import io.grpc.ServerCallHandler;
import io.grpc.ServerInterceptor;
import net.devh.boot.grpc.server.interceptor.GrpcGlobalServerInterceptor;

/**
 * Example {@link ServerInterceptor} that logs all called methods. In this example it is added to Spring's application
 * context via {@link GlobalInterceptorConfiguration}, but is also possible to directly annotate this class with
 * {@link GrpcGlobalServerInterceptor}.
 */
//@GrpcGlobalServerInterceptor
public class LogGrpcInterceptor implements ServerInterceptor {

    private static final Logger log = LoggerFactory.getLogger(LogGrpcInterceptor.class);

    @Override
    public <ReqT, RespT> ServerCall.Listener<ReqT> interceptCall(
            ServerCall<ReqT, RespT> serverCall,
            Metadata metadata,
            ServerCallHandler<ReqT, RespT> serverCallHandler) {

        log.info(serverCall.getMethodDescriptor().getFullMethodName());
        return serverCallHandler.startCall(serverCall, metadata);
    }

}