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

    public String getCode(String email){
        CodeRequest request = CodeRequest.newBuilder().setEmail(email).build();
        CodeResponse code = blockingStub.getCode(request);
        System.out.println(code.getMessage());
        return code.getMessage();
    }
}
