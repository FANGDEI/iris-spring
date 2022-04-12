package com.dyw.provider.service.impl;

import com.dyw.jwt.proto.code.Code;
import com.dyw.jwt.proto.code.CodeRequest;
import com.dyw.jwt.proto.code.CodeResponse;
import com.dyw.jwt.proto.code.getCodeGrpc;
import io.grpc.stub.StreamObserver;
import net.devh.boot.grpc.server.service.GrpcService;
import org.springframework.beans.factory.annotation.Autowired;

/**
 * @author Devil
 * @date 2022-04-12-14:24
 */
@GrpcService
public class CodeServiceImpl extends getCodeGrpc.getCodeImplBase {
    @Override
    public void getCode(CodeRequest request, StreamObserver<CodeResponse> responseObserver) {
        System.out.println("dfsa");
        CodeResponse response = CodeResponse.newBuilder().build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }
}
