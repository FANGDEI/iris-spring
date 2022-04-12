package com.dyw.consumer.service.impl;

import com.dyw.jwt.proto.code.CodeRequest;
import com.dyw.jwt.proto.code.CodeResponse;
import com.dyw.jwt.proto.code.getCodeGrpc;
import net.devh.boot.grpc.client.inject.GrpcClient;
import org.springframework.stereotype.Service;

/**
 * @author Devil
 * @date 2022-04-12-14:46
 */
@Service
public class CodeServiceImpl {
    @GrpcClient("cloud-grpc-server")
    private getCodeGrpc.getCodeBlockingStub blockingStub;

    public void getCode(){
        CodeRequest request = CodeRequest.newBuilder().setEmail("123").build();
        CodeResponse code = blockingStub.getCode(request);
        String message = code.getMessage();
        System.out.println(message);
    }
}
