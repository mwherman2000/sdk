package trinsic.services;

import com.google.protobuf.ByteString;
import com.google.protobuf.InvalidProtocolBufferException;
import com.google.protobuf.Message;
import io.grpc.stub.MetadataUtils;
import trinsic.okapi.DidException;
import trinsic.okapi.Oberon;
import trinsic.okapi.security.v1.Security;
import trinsic.services.account.v1.AccountOuterClass;
import trinsic.services.account.v1.AccountGrpc;
import trinsic.services.common.v1.CommonOuterClass;

public class AccountService extends ServiceBase {
    private final AccountGrpc.AccountBlockingStub stub;

    public AccountService(AccountOuterClass.AccountProfile accountProfile, CommonOuterClass.ServerConfig serverConfig) {
        super(accountProfile, serverConfig);
        this.stub = AccountGrpc.newBlockingStub(this.getChannel());
    }

    public AccountOuterClass.SignInResponse signIn(AccountOuterClass.AccountDetails details) {
        if (details == null)
            details = AccountOuterClass.AccountDetails.newBuilder().build();
        var request = AccountOuterClass.SignInRequest.newBuilder().setDetails(details).build();
        return this.stub.signIn(request);
    }

    public AccountOuterClass.AccountProfile unprotect(AccountOuterClass.AccountProfile profile, String securityCode) throws InvalidProtocolBufferException, DidException {
        var request = Security.UnBlindOberonTokenRequest.newBuilder()
                .setToken(profile.getAuthToken())
                .addBlinding(ByteString.copyFromUtf8(securityCode))
                .build();
        var result = Oberon.unBlindToken(request);

        return AccountOuterClass.AccountProfile.newBuilder(profile)
                .setAuthToken(result.getToken())
                .setProtection(AccountOuterClass.TokenProtection.newBuilder()
                        .setMethod(AccountOuterClass.ConfirmationMethod.None)
                        .setEnabled(false)
                        .build()
                )
                .build();
    }

    public AccountOuterClass.AccountProfile protect(AccountOuterClass.AccountProfile profile, String securityCode) throws InvalidProtocolBufferException, DidException {
        var request = Security.BlindOberonTokenRequest.newBuilder()
                .setToken(profile.getAuthToken())
                .addBlinding(ByteString.copyFromUtf8(securityCode))
                .build();
        var result = Oberon.blindToken(request);

        return AccountOuterClass.AccountProfile.newBuilder(profile)
                .setAuthToken(result.getToken())
                .setProtection(AccountOuterClass.TokenProtection.newBuilder()
                        .setMethod(AccountOuterClass.ConfirmationMethod.Other)
                        .setEnabled(true)
                        .build()
                )
                .build();
    }

    public AccountOuterClass.InfoResponse getInfo() throws InvalidProtocolBufferException, DidException {
        var request = AccountOuterClass.InfoRequest.newBuilder().build();
        return clientWithMetadata(request).info(request);
    }

    private AccountGrpc.AccountBlockingStub clientWithMetadata(Message message) throws InvalidProtocolBufferException, DidException {
        return this.stub.withInterceptors(
                MetadataUtils.newAttachHeadersInterceptor(this.buildMetadata(message)));
    }
}